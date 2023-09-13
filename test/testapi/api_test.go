package testapi

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	batis "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/gozelle/spew"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var db *batis.DB
var containerID string
var cli *client.Client

func TestInsert(t *testing.T) {
	//defer func() {
	//	if db != nil {
	//		require.NoError(t, db.Close())
	//	}
	//	//require.NoError(t, cli.ContainerStop(context.Background(), containerID, container.StopOptions{}))
	//	//require.NoError(t, cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{}))
	//}()
	//
	//initEnv(t)
	//initDB(t)
	//testInsert(t)
	//testInsertBatch(t)
	////testExec(t)
	//testNestedTx(t)
}

func testInsert(t *testing.T) {
	var id int64
	affected, err := db.Debug().Insert("gobatis.products", &Product{
		ProductName:     "Smartwatch",
		Description:     "Advanced health and fitness tracking smartwatch",
		Price:           decimal.NewFromFloat(299.99),
		Weight:          0.05,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	}, batis.Returning("id")).Scan(&id).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)

	var product *Product
	err = db.Query(`select * from gobatis.products where id = #{id}`, batis.Param("id", id)).Scan(&product).Error
	require.NoError(t, err)
	spew.Json(product)
}

func testInsertBatch(t *testing.T) {
	err := db.Debug().InsertBatch("gobatis.products", 2, exampleData).Error
	require.NoError(t, err)
}

func testExec(t *testing.T) {
	rows, err := db.Exec(`INSERT INTO gobatis. (id, name) VALUES (2, 'tom');`).RowsAffected()
	require.NoError(t, err)
	t.Log(rows)
}

func testNestedTx(t *testing.T) {
	tx1 := db.Begin()
	require.NoError(t, tx1.Error)

	tx2 := tx1.Begin()
	require.Error(t, tx2.Error)
}

func initEnv(t *testing.T) {
	ctx := context.Background()
	var err error
	cli, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
		client.WithHost("tcp://127.0.0.1:2375"),
		//client.WithVersion("1.43"),
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
				return
			}
		}
	}

	_, err = cli.ImagePull(ctx, "registry.cn-hongkong.aliyuncs.com/adminium/postgres:13.10", types.ImagePullOptions{})
	require.NoError(t, err)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "registry.cn-hongkong.aliyuncs.com/adminium/postgres:13.10",
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
					HostIP:   "127.0.0.1",
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
	var err error
	db, err = batis.Open(postgres.Open("postgresql://test:test@127.0.0.1:8432/gobatis-test-db?connect_timeout=10&sslmode=disable"))
	if err != nil {
		return
	}
	err = db.Ping()
	require.NoError(t, err)

	err = db.Exec(`
		CREATE SCHEMA IF NOT EXISTS gobatis;
		CREATE TABLE IF NOT EXISTS gobatis.products (
		    id serial PRIMARY KEY,
		    product_name VARCHAR(255) NOT NULL,
		    description TEXT,
		    price DECIMAL NOT NULL, 
		    weight FLOAT,
		    is_available BOOLEAN,
		    manufacture_date DATE,
		    added_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		TRUNCATE gobatis.products; 
	`).Error
	require.NoError(t, err)
}

type Product struct {
	Id              *int            `db:"id"`
	ProductName     string          `db:"product_name"`
	Description     string          `db:"description"`
	Price           decimal.Decimal `db:"price"`
	Weight          float32         `db:"weight"`
	IsAvailable     bool            `db:"is_available"`
	ManufactureDate time.Time       `db:"manufacture_date"`
	AddedDateTime   time.Time       `db:"added_datetime"`
}

var exampleData = []Product{
	{
		ProductName:     "Laptop",
		Description:     "A high-end laptop",
		Price:           decimal.NewFromFloat(1200.50),
		Weight:          1.5,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.January, 20, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	{
		ProductName:     "Smartphone",
		Description:     "Latest model smartphone",
		Price:           decimal.NewFromFloat(800.00),
		Weight:          0.2,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.February, 10, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	{
		ProductName:     "Desk Chair",
		Description:     "Comfortable office chair",
		Price:           decimal.NewFromFloat(150.00),
		Weight:          8.0,
		IsAvailable:     false,
		ManufactureDate: time.Date(2022, time.December, 15, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	{
		ProductName:     "Bluetooth Headphones",
		Description:     "Noise-cancelling over the ear headphones",
		Price:           decimal.NewFromFloat(250.00),
		Weight:          0.3,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.March, 5, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	{
		ProductName:     "4K TV",
		Description:     "65 inch 4K LED TV",
		Price:           decimal.NewFromFloat(1000.00),
		Weight:          25.0,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.January, 28, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
}
