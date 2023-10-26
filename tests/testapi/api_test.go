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
	testInsert(t)
	TestInsertBatch(t)
	TestUpdate(t)
	TestParallelQuery(t)
	TestPagingQuery(t)
	TestFetchQuery(t)
	testContext(t)
	testLooseScan(t)
	testDynamicSQL(t)
	testDelete(t)
	testRecoverValueWhenNoRows(t)
	testInParameter(t)
	testExec(t)
	TestAssociateQuery(t)
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

func TestQuery(t *testing.T) {
	initDB(t)
	var products []*Product
	err := db.Query(`select * from products`).Scan(&products).Error
	require.NoError(t, err)

	spew.Json(products)
}

func TestUpdate(t *testing.T) {
	initDB(t)
	memProducts[Smartphone].Price = decimal.NewFromFloat(900)
	memProducts[Smartphone].StockQuantity = 30

	affected, err := db.Debug().Affect(1).Update("products",
		map[string]any{
			"price":          memProducts[Smartphone].Price,
			"stock_quantity": memProducts[Smartphone].StockQuantity,
		},
		batis.Where("product_name = #{name}", batis.Param("name", Smartphone)),
	).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)

	var product *Product
	err = db.Query(`select * from products where product_name = #{name}`,
		batis.Param("name", Smartphone)).Scan(&product).Error
	require.NoError(t, err)

	memProducts[Smartphone].Id = product.Id
	memProducts[Smartphone].AddedDateTime = product.AddedDateTime

	compareProduct(t, memProducts[Smartphone], product)
}

func testParameter(t *testing.T) {

}

func testScan(t *testing.T) {

}

func TestParallelQuery(t *testing.T) {
	initDB(t)
	var products []*Product
	var count int64
	err := db.Debug().ParallelQuery(
		batis.ParallelQuery{
			SQL: `select * from products where price <= #{ price }`,
			Params: map[string]any{
				"price": 300,
			},
			Scan: func(s batis.Scanner) error {
				return s.Scan(&products)
			},
		},
		batis.ParallelQuery{
			SQL: `select count(1) from products where price <= #{ price }`,
			Params: map[string]any{
				"price": 300,
			},
			Scan: func(s batis.Scanner) error {
				return s.Scan(&count)
			},
		},
	).Error
	require.NoError(t, err)
	require.Equal(t, int64(3), count)
	require.Equal(t, 3, len(products))

	for _, v := range products {
		vv := memProducts[v.ProductName]
		vv.Id = v.Id
		vv.AddedDateTime = v.AddedDateTime
		compareProduct(t, v, vv)
	}
}

func TestPagingQuery(t *testing.T) {

	initDB(t)

	m := map[int]string{
		0: Chair,
		1: BluetoothHeadphones,
		2: Smartwatch,
		3: Smartphone,
	}

	for i := 0; i <= 4; i++ {
		var products []*Product
		var count int64

		q := batis.PagingQuery{
			Select: "*",
			Count:  "1",
			From:   "products",
			Where:  "price <= #{price}",
			Order:  "price asc",
			Page:   int64(i),
			Limit:  1,
			Params: map[string]any{
				"price": decimal.NewFromInt(900),
			},
			Scan: func(s batis.PagingScanner) error {
				return s.Scan(&count, &products)
			},
		}

		err := db.Debug().PagingQuery(q).Error
		require.NoError(t, err)
		require.Equal(t, int64(4), count)

		if i < 4 {
			require.Equal(t, 1, len(products))
			for _, v := range products {
				vv := memProducts[v.ProductName]
				vv.Id = v.Id
				vv.AddedDateTime = v.AddedDateTime
				require.Equal(t, m[i], v.ProductName)
				compareProduct(t, v, vv)
			}
		} else {
			require.Equal(t, 0, len(products))
		}
	}
}

func TestFetchQuery(t *testing.T) {
	initDB(t)
	var products []*Product
	err := db.Debug().FetchQuery(batis.FetchQuery{
		SQL: "select * from products where price < #{price} order by price asc",
		Params: map[string]any{
			"price": 1000,
		},
		Batch: 2,
		Scan: func(scanner batis.Scanner) error {
			var items []*Product
			e := scanner.Scan(&items)
			if e != nil {
				return e
			}
			products = append(products, items...)
			return nil
		},
	})
	require.NoError(t, err)
	expect := []*Product{
		memProducts[Chair],
		memProducts[BluetoothHeadphones],
		memProducts[Smartwatch],
		memProducts[Smartphone],
	}
	compareProducts(t, expect, products)
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

func testCustomizeDataType(t *testing.T) {

}

func testPlugin(t *testing.T) {

}

func testDynamicSQL(t *testing.T) {
	var products []*Product
	q := `
		select * from products 
		<where>
		    <if test="price > 0"> price > #{price} </if>
		    <if test="isAvailable">and is_available is true</if>
		</where>		                            		                            
	`
	err := db.Debug().Query(q, batis.Param("price", 1), batis.Param("isAvailable", true)).Scan(&products).Error
	require.NoError(t, err)

	spew.Json(products)

}

func testDelete(t *testing.T) {
	affected, err := db.Debug().Affect(1).Delete("products",
		batis.Where("product_name = #{name}", batis.Param("name", Smartphone))).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)
}

func testRecoverValueWhenNoRows(t *testing.T) {
	var product *Product
	err := db.Affect(0).Query(`select * from products where id = 0`).Scan(&product).Error
	require.NoError(t, err)
	require.True(t, product == nil)
}

func testInParameter(t *testing.T) {
	var products []*Product
	err := db.Debug().Query(`
   select * from products where id in 
	<foreach item="item" index="index" collection="ids" open="(" separator="," close=")">
		#{item}
	</foreach>
`, batis.Param("ids", []int64{1, 2, 3})).Scan(&products).Error
	require.NoError(t, err)
}

func TestAssociateQuery(t *testing.T) {

	initDB(t)

	type ProductWrap struct {
		Name    string
		Product *Product
		Age     int64
	}

	wraps := []ProductWrap{
		{Name: "Laptop"},
		{Name: "TV"},
	}

	err := db.Debug().AssociateQuery(batis.AssociateQuery{
		SQL: "select * from products where product_name in #{ids}",
		Params: map[string]any{
			//"ids": batis.Extract(wraps, "$.Name"),
			"ids": []string{"Laptop", "TV"},
		},
		Scan: func(scanner batis.AssociateScanner) error {
			return scanner.Scan(&wraps, "product_name => $.Name", "$.Product")
		},
	}).Error

	require.NoError(t, err)

	spew.Json(wraps)
}
