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
	"github.com/gobatis/gobatis/driver/postgres"
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
	if db != nil {
		return
	}
	host = os.Getenv(client.EnvOverrideHost)
	if host == "" {
		host = "127.0.0.1"
	}
	var err error
	db, err = batis.Open(postgres.Open(fmt.Sprintf("postgresql://test:test@%s:8432/gobatis-test-db?connect_timeout=10&sslmode=disable", host)))
	if err != nil {
		return
	}
	err = db.Ping()
	require.NoError(t, err)

	err = db.Exec(`
		create schema if not exists gobatis;
		create table if not exists products (
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
		create unique index if not exists products_product_name_uindex on products (product_name);
	`).Error
	require.NoError(t, err)

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

const (
	Laptop              = "Laptop"
	Smartphone          = "Smartphone"
	Smartwatch          = "Smartwatch"
	Chair               = "Chair"
	BluetoothHeadphones = "Bluetooth Headphones"
	TV                  = "TV"
)

var memProducts = map[string]*Product{
	Smartwatch: {
		ProductName:     Smartwatch,
		Description:     "Advanced health and fitness tracking smartwatch",
		Price:           decimal.NewFromFloat(299.99),
		Weight:          0.05,
		StockQuantity:   5,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	Laptop: {
		ProductName:     Laptop,
		Description:     "A high-end laptop",
		Price:           decimal.NewFromFloat(1200.50),
		Weight:          1.5,
		StockQuantity:   10,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.January, 20, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	Smartphone: {
		ProductName:     Smartphone,
		Description:     "Latest model smartphone",
		Price:           decimal.NewFromFloat(800.00),
		Weight:          0.2,
		StockQuantity:   20,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.February, 10, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	Chair: {
		ProductName:     Chair,
		Description:     "Comfortable office chair",
		Price:           decimal.NewFromFloat(150.00),
		Weight:          8.0,
		StockQuantity:   15,
		IsAvailable:     false,
		ManufactureDate: time.Date(2022, time.December, 15, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	BluetoothHeadphones: {
		ProductName:     BluetoothHeadphones,
		Description:     "Noise-cancelling over the ear headphones",
		Price:           decimal.NewFromFloat(250.00),
		Weight:          0.3,
		StockQuantity:   100,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.March, 5, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
	TV: {
		ProductName:     TV,
		Description:     "65 inch 4K LED TV",
		Price:           decimal.NewFromFloat(1000.00),
		Weight:          25.0,
		StockQuantity:   120,
		IsAvailable:     true,
		ManufactureDate: time.Date(2023, time.January, 28, 0, 0, 0, 0, time.UTC),
		AddedDateTime:   time.Now(),
	},
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
	//testInsert(t)
	//testInsertBatch(t)
	//testUpdate(t)
	//testParallelQuery(t)
	//testPagingQuery(t)
	//testFetchQuery(t)
	//testContext(t)
	//testLooseScan(t)
	//testDynamicSQL(t)
	//testDelete(t)
	//testRecoverValueWhenNoRows(t)
	//testInParameter(t)
	//testExec(t)
	TestAssociateQuery(t)
	//testNestedTx(t)
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
	affected, err := db.Debug().Insert("products", memProducts[Smartwatch], batis.Returning("id")).Scan(&memProducts[Smartwatch].Id).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)
	require.True(t, memProducts[Smartwatch].Id != nil && *memProducts[Smartwatch].Id > 0)

	// test insertion conflict
	memProducts[Smartwatch].ManufactureDate = time.Date(2023, time.April, 12, 0, 0, 0, 0, time.UTC)
	memProducts[Smartwatch].AddedDateTime = time.Now()
	affected, err = db.Debug().Affect(1).Insert("products",
		&Product{
			ProductName:     "Smartwatch",
			ManufactureDate: memProducts[Smartwatch].ManufactureDate,
			AddedDateTime:   memProducts[Smartwatch].AddedDateTime,
		},
		batis.OnConflict("product_name", `do update set manufacture_date = excluded.manufacture_date`),
	).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)

	// test insertion conflict update and
	// return the specified field
	var productName string
	memProducts[Smartwatch].Price = decimal.NewFromFloat(300.00)
	memProducts[Smartwatch].ManufactureDate = time.Now()
	memProducts[Smartwatch].AddedDateTime = time.Now()
	err = db.Debug().Affect(1).Insert("products",
		&Product{
			ProductName:     "Smartwatch",
			Price:           memProducts[Smartwatch].Price,
			ManufactureDate: memProducts[Smartwatch].ManufactureDate,
			AddedDateTime:   memProducts[Smartwatch].AddedDateTime,
		},
		batis.OnConflict("product_name", `do update set price = excluded.price`),
		batis.Returning("product_name")).Scan(&productName).Error
	require.NoError(t, err)
	require.Equal(t, "Smartwatch", productName)

	// test query operation and
	// compare the data after changes
	var product *Product
	err = db.Query(`select * from products where id = #{id}`, batis.Param("id", *memProducts[Smartwatch].Id)).Scan(&product).Error
	require.NoError(t, err)
	require.True(t, product.Id != nil && *product.Id > 0)
	require.Equal(t, "Smartwatch", product.ProductName)
	require.Equal(t, "Advanced health and fitness tracking smartwatch", product.Description)
	require.Equal(t, "300", product.Price.String())
	require.Equal(t, float32(0.05), product.Weight)
	require.Equal(t, int64(5), product.StockQuantity)
	require.Equal(t, true, product.IsAvailable)
	require.Equal(t, "2023-04-12", product.ManufactureDate.Format("2006-01-02"))
	require.Equal(t, true, product.AddedDateTime.Unix() > 0)
}

// Testing batch insertion of data, including checking the number of affected rows,
// verifying the inserted data, returning the last insert ID, and scanning all auto-incremented IDs.
func TestInsertBatch(t *testing.T) {
	initDB(t)
	affected, err := db.Debug().Affect(5).InsertBatch("products", 2, extractMemProducts(Smartwatch)).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(5), affected)

	var products []*Product
	err = db.Debug().Query(`select * from products where stock_quantity >= 10`).Scan(&products).Error
	require.NoError(t, err)

	compareProducts(t, extractMemProducts(Smartwatch), products)

	affected, err = db.Debug().Affect(5).
		Delete("products", batis.Where("stock_quantity >= #{ v }", batis.Param("v", 10))).RowsAffected()
	require.NoError(t, err)

	products = []*Product{}
	affected, err = db.Debug().Affect(5).InsertBatch("products", 2, extractMemProducts(Smartwatch),
		batis.Returning("*")).Scan(&products).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(5), affected)

	compareProducts(t, extractMemProducts(Smartwatch), products)
}

func compareProducts(t *testing.T, r, c []*Product) {
	m := map[string]*Product{}
	for _, v := range c {
		m[v.ProductName] = v
	}
	for _, v := range r {
		vv, ok := m[v.ProductName]
		require.True(t, ok)
		v.Id = vv.Id
		compareProduct(t, v, vv)
	}
}

func compareProduct(t *testing.T, v1, v2 *Product) {
	//d1, err := json.Marshal(v1)
	//require.NoError(t, err)
	//d2, err := json.Marshal(v2)
	//require.NoError(t, err)
	//if err = fastjson.EqualsBytes(d1, d2); err != nil {
	//	t.Fatal(fmt.Errorf("compare products v1 != v2, err: %s", err))
	//}
	return
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

func testUpdate(t *testing.T) {

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
	spew.Json(count, products)
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
