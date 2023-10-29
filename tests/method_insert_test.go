package tests

import (
	"context"
	"testing"
	"time"

	batis "github.com/gobatis/gobatis"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

// Test common insertion scenarios, including ordinary insertion,
// returning auto-incremented ID, handling conflict insertion,
// and handling conflict insertion while returning row fields
// 1. insert into ... returning ...
// 2. insert into ... on conflict ...
// 3. insert into ... on conflict ... returning ...
func TestInsert(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()

	// perform ordinary insertion operation and
	// return the auto-increment primary key
	m := getProductsMap()
	affected, err := db.Debug().Insert("products", m[Smartwatch], batis.Returning("id")).Scan(&m[Smartwatch].Id).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)
	require.True(t, m[Smartwatch].Id != nil && *m[Smartwatch].Id > 0)

	// test insertion conflict
	m[Smartwatch].ManufactureDate = time.Date(2023, time.April, 12, 0, 0, 0, 0, time.UTC)
	m[Smartwatch].AddedDateTime = time.Now()
	affected, err = db.Debug().Affect(1).Insert("products",
		&Product{
			ProductName:     "Smartwatch",
			ManufactureDate: m[Smartwatch].ManufactureDate,
			AddedDateTime:   m[Smartwatch].AddedDateTime,
		},
		batis.OnConflict("product_name", `do update set manufacture_date = excluded.manufacture_date`),
	).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)

	// test insertion conflict update and
	// return the specified field
	var productName string
	m[Smartwatch].Price = decimal.NewFromFloat(300.00)
	m[Smartwatch].ManufactureDate = time.Now()
	m[Smartwatch].AddedDateTime = time.Now()
	err = db.Debug().Affect(1).Insert("products",
		&Product{
			ProductName:     "Smartwatch",
			Price:           m[Smartwatch].Price,
			ManufactureDate: m[Smartwatch].ManufactureDate,
			AddedDateTime:   m[Smartwatch].AddedDateTime,
		},
		batis.OnConflict("product_name", `do update set price = excluded.price`),
		batis.Returning("product_name")).Scan(&productName).Error
	require.NoError(t, err)
	require.Equal(t, "Smartwatch", productName)

	// test query operation and
	// compare the data after changes
	var product *Product
	err = db.Query(`select * from products where id = #{id}`, batis.Param("id", *m[Smartwatch].Id)).Scan(&product).Error
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

func TestInsertAffect(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()
	m := getProductsMap()
	expectAffectConstrictError(t, db.Affect(0).Insert("products", m[Smartwatch]).Error)
	expectAffectConstrictError(t, db.Affect(0).Insert("products", m[Smartwatch], batis.Returning("id")).Scan(&m[Smartwatch].Id).Error)

	require.NoError(t, db.Affect(1).Insert("products", m[Smartwatch]).Error)
	require.NoError(t, db.Affect(1).Insert("products", m[Chair]).Error)

	var count int
	require.NoError(t, db.Query(`select count(1) from products`).Scan(&count).Error)
	require.Equal(t, 2, count)
}

func TestInsertExecutorConflict(t *testing.T) {
	m := getProductsMap()
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).Insert("products", &Product{}).Error)
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).Update("products", nil, batis.Where("")).Error)
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).Query(``).Error)
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).Exec(``).Error)
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).InsertBatch(``, 2, nil).Error)
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).ParallelQuery(batis.ParallelQuery{}).Error)
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).PagingQuery(batis.PagingQuery{}).Error)
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).FetchQuery(batis.FetchQuery{}).Error)
	expectExecutorConflictError(t, db.Insert("products", m[Smartwatch], batis.Returning("*")).AssociateQuery(batis.AssociateQuery{}).Error)
}

func TestInsertRowsAffected(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()
	m := getProductsMap()
	{
		rowsAffected, err := db.Insert("products", m[Smartwatch]).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(1), rowsAffected)
	}
	{
		rowsAffected, err := db.Insert("products", m[Chair], batis.Returning("id")).Scan(&m[Chair].Id).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(1), rowsAffected)
	}
}

func TestInsertLastInsertId(t *testing.T) {

}

func TestInsertContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	expectContextCanceled(t, db.WithContext(ctx).Insert("products", &Product{}).Error)
}

func TestInsertTraceId(t *testing.T) {
	m := getProductsMap()
	{
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithTraceId("id").Insert("products2", m[TV])
		w.expectTraceId(t, "id")
	}
	{
		ctx := batis.WithTraceId(context.Background(), "ctx")
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithContext(ctx).Insert("products2", m[TV])
		w.expectTraceId(t, "ctx")
	}
}

func TestInsertTx(t *testing.T) {

}

func TestInsertDebug(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()

	m := getProductsMap()
	db.Debug().Insert("products", m[TV])
}

func TestInsertTrace(t *testing.T) {

}

func TestInsertColumnTag(t *testing.T) {

	defer func() {
		cleanProducts(t)
	}()

	m := getProductsMap()
	p := m[TV]
	i := &ProductJ{
		JId:              p.Id,
		JProductName:     p.ProductName,
		JDescription:     p.Description,
		JPrice:           p.Price,
		JWeight:          p.Weight,
		JStockQuantity:   p.StockQuantity,
		JIsAvailable:     p.IsAvailable,
		JManufactureDate: p.ManufactureDate,
		JAddedDateTime:   p.AddedDateTime,
	}

	s := db.Session(&batis.Session{ColumnTag: "json"})

	require.NoError(t, s.Insert("products", i).Error)

	var n *ProductJ

	require.NoError(t, s.Query(`select * from products where product_name = #{ name }`,
		batis.Param("name", i.JProductName)).Scan(&n).Error)

	require.True(t, *n.JId > 0)
	n.JId = nil

	compareProductJ(t, i, n)
}

func TestInsertClone(t *testing.T) {

}
