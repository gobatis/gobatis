package testapi

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	batis "github.com/gobatis/gobatis"
	"github.com/gozelle/spew"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var db *batis.DB
var containerID string
var cli *client.Client
var host string

func initEnv(t *testing.T) {
	ctx := context.Background()
	host = os.Getenv(client.EnvOverrideHost)
	if host == "" {
		host = "127.0.0.1"
	}
	t.Log(host)
	var err error
	cli, err = client.NewClientWithOpts(
		client.WithAPIVersionNegotiation(),
		client.WithHost(fmt.Sprintf("tcp://%s:2375", host)),
	)
	require.NoError(t, err)

	pwd, err := os.Getwd()
	require.NoError(t, err)

	const containerName = "gobatis-test-postgres-13-10"
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	require.NoError(t, err)

	for _, v := range containers {
		for _, name := range v.Names {
			if name == "/"+containerName {
				containerID = v.ID
				t.Logf("container: %s exists", name)
				return
			}
		}
	}

	//_, err = cli.ImagePull(ctx, "registry.cn-hangzhou.aliyuncs.com/tashost/timescaledb:13.10", types.ImagePullOptions{})
	_, err = cli.ImagePull(ctx, "adminium/postgres:13.10", types.ImagePullOptions{})
	require.NoError(t, err)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		//Image: "registry.cn-hangzhou.aliyuncs.com/tashost/timescaledb:13.10",
		Image: "adminium/postgres:13.10",
		Env: []string{
			"POSTGRES_PASSWORD=test",
			"POSTGRES_USER=test",
			"POSTGRES_DB=gobatis-test-db",
		},
		ExposedPorts: nat.PortSet{
			"5432/tcp": struct{}{},
		},
		WorkingDir:   pwd,
		AttachStdout: false,
		AttachStderr: false,
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8432",
				},
			},
		},
	}, nil, nil, containerName)
	require.NoError(t, err)
	containerID = resp.ID
	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	require.NoError(t, err)
}

func initDB(t *testing.T) {

	//err = db.Exec(`truncate products`).Error
	//require.NoError(t, err)

}

type Product struct {
	Id              *int            `db:"id"`
	ProductName     string          `db:"product_name"`
	Description     string          `db:"description"`
	Price           decimal.Decimal `db:"price"`
	Weight          float32         `db:"weight"`
	StockQuantity   int64           `db:"stock_quantity"`
	IsAvailable     bool            `db:"is_available"`
	ManufactureDate time.Time       `db:"manufacture_date"`
	AddedDateTime   time.Time       `db:"added_datetime"`
}

func extractMemProducts(exclude ...string) (products []*Product) {
	m := map[string]struct{}{}
	for _, v := range exclude {
		m[v] = struct{}{}
	}
	for k, v := range memProducts {
		if _, ok := m[k]; !ok {
			products = append(products, v)
		}
	}
	return
}

func TestAPIFeatures(t *testing.T) {
	defer func() {
		if db != nil {
			require.NoError(t, db.Close())
		}
		//require.NoError(t, cli.ContainerStop(context.Background(), containerID, container.StopOptions{}))
		//require.NoError(t, cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{}))
	}()

	initEnv(t)
	initDB(t)

	testContext(t)
	testLooseScan(t)

	testRecoverValueWhenNoRows(t)

	testExec(t)
	
	//testNestedTx(t)
}

func testExec(t *testing.T) {
	//rows, err := db.Exec(`INSERT INTO gobatis. (id, name) VALUES (2, 'tom');`).RowsAffected()
	//require.NoError(t, err)
	//t.Log(rows)
}

//func testNestedTx(t *testing.T) {
//	tx1 := db.Begin()
//	require.NoError(t, tx1.Error)
//
//	tx2 := tx1.Begin()
//	require.Error(t, tx2.Error)
//}

func testParameter(t *testing.T) {

}

func testScan(t *testing.T) {

}

func testContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer func() {
		cancel()
	}()
	err := db.WithContext(ctx).Query(`select pg_sleep(3);`).Scan(nil).Error
	require.Error(t, err)
	t.Log(err, errors.Is(err, net.ErrWriteToConnected))
}

type ProductPlus struct {
	Product
	Offline bool
}

func testLooseScan(t *testing.T) {
	var pb *Product
	err := db.Debug().Query(`select * from products order by id asc limit 1`).Scan(&pb).Error
	require.NoError(t, err)

	spew.Json(pb)
}

func testRecoverValueWhenNoRows(t *testing.T) {
	var product *Product
	err := db.Affect(0).Query(`select * from products where id = 0`).Scan(&product).Error
	require.NoError(t, err)
	require.True(t, product == nil)
}
