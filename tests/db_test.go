package tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	batis "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/gobatis/gobatis/logger"
	"github.com/gozelle/fastjson"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var db *batis.DB

const (
	driverPostgres   = "postgres"
	driverSQLite     = "sqlite"
	driverMySQL      = "mysql"
	driverClickhouse = "clickhouse"
	driverTiDB       = "tidb"
	driverSQLServer  = "sqlserver"
)

func init() {
	var err error
	defer func() {
		if err != nil {
			log.Fatalf("init db error: %s", err)
		}
	}()
	dsn := os.Getenv("GOBATIS_TEST_DSN")
	if dsn == "" {
		dsn = "postgresql://test:test@127.0.0.1:8432/gobatis-test-db?connect_timeout=10&sslmode=disable&TimeZone=Asia/Shanghai"
	}
	db, err = batis.Open(postgres.Open(dsn))
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}

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
	if err != nil {
		return
	}
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

type ProductJ struct {
	JId              *int            `json:"id"`
	JProductName     string          `json:"product_name"`
	JDescription     string          `json:"description"`
	JPrice           decimal.Decimal `json:"price"`
	JWeight          float32         `json:"weight"`
	JStockQuantity   int64           `json:"stock_quantity"`
	JIsAvailable     bool            `json:"is_available"`
	JManufactureDate time.Time       `json:"manufacture_date"`
	JAddedDateTime   time.Time       `json:"added_datetime"`
}

const (
	Laptop              = "Laptop"
	Smartphone          = "Smartphone"
	Smartwatch          = "Smartwatch"
	Chair               = "Chair"
	BluetoothHeadphones = "Bluetooth Headphones"
	TV                  = "TV"
)

func getProductsList() []*Product {
	m := getProductsMap()
	l := make([]*Product, 0)
	for _, v := range m {
		l = append(l, v)
	}
	return l
}

func getProductsMap() map[string]*Product {
	return map[string]*Product{
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

func compareProductJ(t *testing.T, v1, v2 *ProductJ) {
	p1 := &Product{
		Id:              v1.JId,
		ProductName:     v1.JProductName,
		Description:     v1.JDescription,
		Price:           v1.JPrice,
		Weight:          v1.JWeight,
		StockQuantity:   v1.JStockQuantity,
		IsAvailable:     v1.JIsAvailable,
		ManufactureDate: v1.JManufactureDate,
		AddedDateTime:   v1.JAddedDateTime,
	}
	p2 := &Product{
		Id:              v2.JId,
		ProductName:     v2.JProductName,
		Description:     v2.JDescription,
		Price:           v2.JPrice,
		Weight:          v2.JWeight,
		StockQuantity:   v2.JStockQuantity,
		IsAvailable:     v2.JIsAvailable,
		ManufactureDate: v2.JManufactureDate,
		AddedDateTime:   v2.JAddedDateTime,
	}
	compareProduct(t, p1, p2)
}

func compareProduct(t *testing.T, v1, v2 *Product) {
	// TODO check time format
	v1.AddedDateTime = time.Time{}
	v2.AddedDateTime = time.Time{}
	d1, err := json.Marshal(v1)
	require.NoError(t, err)
	d2, err := json.Marshal(v2)
	require.NoError(t, err)
	if err = fastjson.EqualsBytes(d1, d2); err != nil {
		t.Fatal(fmt.Errorf("compare products v1(%s) != v2(%s), err: %s", v1.ProductName, v2.ProductName, err))
	}
	return
}

func cleanProducts(t *testing.T) {
	require.NoError(t, db.Exec(`delete from products`).Error)
}

func prepareProducts(t *testing.T) {
	l := getProductsList()
	require.NoError(t, db.InsertBatch("products", 2, l).Error)
}

func expectExecutorConflictError(t *testing.T, err error) {
	require.True(t, errors.Is(err, batis.ErrExecutorConflict))
}

func expectAffectConstrictError(t *testing.T, err error) {
	require.True(t, errors.Is(err, batis.ErrAffectConstrict))
}

func expectContextDeadlineExceeded(t *testing.T, err error) {
	require.True(t, errors.Is(err, context.Canceled))
}

var _ logger.Writer = (*traceWriter)(nil)

func newTraceWriter() *traceWriter {
	return &traceWriter{}
}

type traceWriter struct {
	logs []string
}

func (tw *traceWriter) Printf(s string, i ...interface{}) {
	tw.logs = append(tw.logs, fmt.Sprintf(s, i...))
}

func (tw *traceWriter) expectTraceId(t *testing.T, traceId string) {
	require.True(t, len(tw.logs) > 0)
	for _, v := range tw.logs {
		require.True(t, strings.HasPrefix(v, fmt.Sprintf("[%s]", traceId)))
	}
}

func traceLogger(w *traceWriter) logger.Logger {
	return logger.NewtLogger(w)
}
