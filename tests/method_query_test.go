package tests

import (
	"context"
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/gozelle/spew"
	"github.com/stretchr/testify/require"
)

func TestQuery(t *testing.T) {
	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()
	var products []*Product
	err = db.Query(`select * from products`).Scan(&products).Error
	require.NoError(t, err)

	spew.Json(products)
}

func TestQueryAffect(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()
	expectAffectConstrictError(t, db.Affect(0).Query(`select * from products`).Scan(nil).Error)
	require.NoError(t, db.Affect(1).Query(`select * from products where product_name = #{name}`, batis.Param("name", TV)).Scan(nil).Error)
}

func TestQueryExecutorConflict(t *testing.T) {
	expectExecutorConflictError(t, db.Query(`select 1`).Insert("products", &Product{}).Error)
	expectExecutorConflictError(t, db.Query(`select 1`).Update("products", nil, batis.Where("")).Error)
	expectExecutorConflictError(t, db.Query(`select 1`).Query(``).Error)
	expectExecutorConflictError(t, db.Query(`select 1`).Exec(``).Error)
	expectExecutorConflictError(t, db.Query(`select 1`).InsertBatch(``, 2, nil).Error)
	expectExecutorConflictError(t, db.Query(`select 1`).ParallelQuery(batis.ParallelQuery{}).Error)
	expectExecutorConflictError(t, db.Query(`select 1`).PagingQuery(batis.PagingQuery{}).Error)
	expectExecutorConflictError(t, db.Query(`select 1`).FetchQuery(batis.FetchQuery{}).Error)
	expectExecutorConflictError(t, db.Query(`select 1`).AssociateQuery(batis.AssociateQuery{}).Error)
}

func TestQueryRowsAffected(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()
	l := getProductsList()
	{
		rowsAffected, err := db.Query(`select * from products`).Scan(nil).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(len(l)), rowsAffected)
	}
	{
		rowsAffected, err := db.Query(`select * from products where product_name = #{name}`, batis.Param("name", TV)).Scan(nil).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(1), rowsAffected)
	}
	{
		rowsAffected, err := db.Query(`select 1`, batis.Param("name", TV)).Scan(nil).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(1), rowsAffected)
	}
}

func TestQueryContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	expectContextCanceled(t, db.WithContext(ctx).Query(`select 1`).Error)
}

func TestQueryTraceId(t *testing.T) {
	{
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithTraceId("id").Query(`select * from products2`).Scan(nil)
		w.expectTraceId(t, "id")
	}
	{
		ctx := batis.WithTraceId(context.Background(), "ctx")
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithContext(ctx).Query(`select * from products2`).Scan(nil)
		w.expectTraceId(t, "ctx")
	}
}

func TestQueryColumnTag(t *testing.T) {
	prepareProducts(t)
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

	var n *ProductJ

	require.NoError(t, s.Query(`select * from products where product_name = #{ name }`,
		batis.Param("name", TV)).Scan(&n).Error)

	require.True(t, *n.JId > 0)
	n.JId = nil

	compareProductJ(t, i, n)
}
