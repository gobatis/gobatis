package testapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gozelle/fastjson"
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

func initEnv(t *testing.T) {
	ctx := context.Background()
	host := os.Getenv(client.EnvOverrideHost)
	if host == "" {
		host = "tcp://127.0.0.1:2375"
	}
	t.Log(host)
	var err error
	cli, err = client.NewClientWithOpts(
		client.WithAPIVersionNegotiation(),
		client.WithHost(host),
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
	
	_, err = cli.ImagePull(ctx, "registry.cn-hangzhou.aliyuncs.com/tashost/timescaledb:13.10", types.ImagePullOptions{})
	require.NoError(t, err)
	
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "registry.cn-hangzhou.aliyuncs.com/tashost/timescaledb:13.10",
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
	var err error
	db, err = batis.Open(postgres.Open("postgresql://test:test@127.0.0.1:8432/gobatis-test-db?connect_timeout=10&sslmode=disable"))
	if err != nil {
		return
	}
	err = db.Ping()
	require.NoError(t, err)
	
	err = db.Exec(`
		create schema if not exists gobatis;
		create table if not exists gobatis.products (
		    id serial PRIMARY KEY,
		    product_name VARCHAR(255) NOT NULL,
		    description TEXT,
		    price DECIMAL NOT NULL, 
		    weight FLOAT,
            stock_quantity BIGINT,
		    is_available BOOLEAN,
		    manufacture_date DATE,
		    added_datetime TIMESTAMPTZ
		);
		create unique index if not exists products_product_name_uindex on gobatis.products (product_name);
		truncate gobatis.products; 
	`).Error
	require.NoError(t, err)
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

var exampleData = []*Product{
	{
		ProductName:     "Laptop",
		Description:     "A high-end laptop",
		Price:           decimal.NewFromFloat(1200.50),
		Weight:          1.5,
		StockQuantity:   10,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.January, 20, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	{
		ProductName:     "Smartphone",
		Description:     "Latest model smartphone",
		Price:           decimal.NewFromFloat(800.00),
		Weight:          0.2,
		StockQuantity:   20,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.February, 10, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	{
		ProductName:     "Desk Chair",
		Description:     "Comfortable office chair",
		Price:           decimal.NewFromFloat(150.00),
		Weight:          8.0,
		StockQuantity:   15,
		IsAvailable:     false,
		ManufactureDate: time.Date(2022, time.December, 15, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	{
		ProductName:     "Bluetooth Headphones",
		Description:     "Noise-cancelling over the ear headphones",
		Price:           decimal.NewFromFloat(250.00),
		Weight:          0.3,
		StockQuantity:   100,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.March, 5, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	{
		ProductName:     "4K TV",
		Description:     "65 inch 4K LED TV",
		Price:           decimal.NewFromFloat(1000.00),
		Weight:          25.0,
		StockQuantity:   120,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.January, 28, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
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
	testInsertBatch(t)
	//testExec(t)
	testNestedTx(t)
}

// Test common insertion scenarios, including ordinary insertion,
// returning auto-incremented ID, handling conflict insertion,
// and handling conflict insertion while returning row fields
// 1. insert into ... returning ...
// 2. insert into ... on conflict ...
// 3. insert into ... on conflict ... returning ...
func testInsert(t *testing.T) {
	// perform ordinary insertion operation and
	// return the auto-increment primary key
	var id int64
	affected, err := db.Debug().Insert("gobatis.products", &Product{
		ProductName:     "Smartwatch",
		Description:     "Advanced health and fitness tracking smartwatch",
		Price:           decimal.NewFromFloat(299.99),
		Weight:          0.05,
		StockQuantity:   5,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	}, batis.Returning("id")).Scan(&id).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)
	
	// test insertion conflict
	affected, err = db.Debug().Affect(1).Insert("gobatis.products",
		&Product{
			ProductName:     "Smartwatch",
			ManufactureDate: time.Date(2023, time.April, 12, 0, 0, 0, 0, time.UTC),
		},
		batis.OnConflict("product_name", `do update set manufacture_date = excluded.manufacture_date`),
	).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)
	
	// test insertion conflict update and
	// return the specified field
	var productName string
	err = db.Debug().Affect(1).Insert("gobatis.products",
		&Product{
			ProductName: "Smartwatch",
			Price:       decimal.NewFromFloat(300.00),
		},
		batis.OnConflict("product_name", `do update set price = excluded.price`),
		batis.Returning("product_name")).Scan(&productName).Error
	require.NoError(t, err)
	require.Equal(t, "Smartwatch", productName)
	
	// test query operation and
	// compare the data after changes
	var product *Product
	err = db.Query(`select * from gobatis.products where id = #{id}`, batis.Param("id", id)).Scan(&product).Error
	require.NoError(t, err)
	spew.Json(product)
	require.True(t, product.Id != nil && *product.Id > 0)
	require.Equal(t, "Smartwatch", product.ProductName)
	require.Equal(t, "Advanced health and fitness tracking smartwatch", product.Description)
	require.Equal(t, "300", product.Price.String())
	require.Equal(t, float32(0.05), product.Weight)
	require.Equal(t, int64(5), product.StockQuantity)
	require.Equal(t, true, product.IsAvailable)
	//require.Equal(t, "2023-04-12", product.ManufactureDate.Format("2006-01-02"))
	//require.Equal(t, true, product.AddedDateTime.Unix() > 0)
}

// Testing batch insertion of data, including checking the number of affected rows, 
// verifying the inserted data, returning the last insert ID, and scanning all auto-incremented IDs.
func testInsertBatch(t *testing.T) {
	affected, err := db.Debug().Affect(5).InsertBatch("gobatis.products", 2, exampleData).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(5), affected)
	
	var products []*Product
	err = db.Debug().Query(`select * from gobatis.products where stock_quantity >= 10`).Scan(&products).Error
	require.NoError(t, err)
	
	compareProducts(t, exampleData, products)
	
	affected, err = db.Debug().Affect(5).
		Delete("gobatis.products", batis.Where("stock_quantity >= #{ v }", batis.Param("v", 10))).RowsAffected()
	require.NoError(t, err)
	
	// TODO add .Affect(5)
	products = []*Product{}
	affected, err = db.Debug().InsertBatch("gobatis.products", 2, exampleData,
		batis.Returning("*")).Scan(&products).RowsAffected()
	require.NoError(t, err)
	//require.Equal(t, int64(5), affected)
	
	compareProducts(t, exampleData, products)
}

func compareProducts(t *testing.T, r, c []*Product) {
	cMap := map[string]*Product{}
	for _, v := range c {
		cMap[v.ProductName] = v
	}
	for _, v := range r {
		vv, ok := cMap[v.ProductName]
		require.True(t, ok)
		v.Id = vv.Id
		compareProduct(t, v, vv)
	}
}

func compareProduct(t *testing.T, v1, v2 *Product) bool {
	d1, err := json.Marshal(v1)
	require.NoError(t, err)
	d2, err := json.Marshal(v2)
	require.NoError(t, err)
	if err = fastjson.EqualsBytes(d1, d2); err != nil {
		t.Fatal(fmt.Errorf("compare products v1 != v2, err: %s", err))
	}
	return false
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

func testAffect(t *testing.T) {
	
}

func testParameter(t *testing.T) {
	
}

func testScan(t *testing.T) {
	
}

func testDelete(t *testing.T) {
	
}

func testPagingQuery(t *testing.T) {
	
}

func testParallelQuery(t *testing.T) {
	
}

func testFetchQuery(t *testing.T) {
	
}

func testContext(t *testing.T) {
	
}

func testAssociateQuery(t *testing.T) {
	
}

func testLooseScan(t *testing.T) {
	
}

func testCustomizeDataType(t *testing.T) {
	
}

func testPlugin(t *testing.T) {
	
}

func testDynamicSQL(t *testing.T) {
	
}
