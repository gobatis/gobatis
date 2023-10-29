package tests

import (
	"context"
	"errors"
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {

	l := getProductsList()
	err := db.InsertBatch("products", 3, l).Error
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Exec("delete from products").Error)
	}()

	m := getProductsMap()
	m[Smartphone].Price = decimal.NewFromFloat(900)
	m[Smartphone].StockQuantity = 30

	affected, err := db.Debug().Affect(1).Update("products",
		map[string]any{
			"price":          m[Smartphone].Price,
			"stock_quantity": m[Smartphone].StockQuantity,
		},
		batis.Where("product_name = #{name}", batis.Param("name", Smartphone)),
	).RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), affected)

	var product *Product
	err = db.Query(`select * from products where product_name = #{name}`,
		batis.Param("name", Smartphone)).Scan(&product).Error
	require.NoError(t, err)

	m[Smartphone].Id = product.Id
	m[Smartphone].AddedDateTime = product.AddedDateTime

	compareProduct(t, m[Smartphone], product)
}

func TestUpdateAffect(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()

	u := &Product{
		ProductName: TV,
	}

	var p1 *Product
	require.NoError(t, db.Query(`select * from products where product_name = #{name}`,
		batis.Param("name", u.ProductName)).Scan(&p1).Error)
	require.True(t, p1.Price.GreaterThan(decimal.Zero))

	err := db.Affect(0).Update("products", map[string]any{
		"price": u.Price,
	}, batis.Where("product_name = #{name}", batis.Param("name", u.ProductName))).Error
	require.True(t, errors.Is(err, batis.ErrAffectConstrict))

	var p2 *Product
	require.NoError(t, db.Query(`select * from products where product_name = #{name}`,
		batis.Param("name", u.ProductName)).Scan(&p2).Error)
	require.True(t, p2.Price.GreaterThan(decimal.Zero))
	require.True(t, p2.Price.GreaterThan(decimal.Zero))
	require.True(t, p1.Price.Equal(p2.Price))

	p1.Price = p1.Price.Add(decimal.NewFromFloat(100))
	u.Price = p1.Price
	require.NoError(t, db.Affect(1).Update("products", map[string]any{
		"price": u.Price,
	}, batis.Where("product_name = #{name}", batis.Param("name", u.ProductName))).Error)
	require.True(t, u.Price.Equal(p1.Price))
}

func TestUpdateExecutorConflict(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()

	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).Insert("products", &Product{}).Error)
	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).Update("products", nil, batis.Where("")).Error)
	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).Query(``).Error)
	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).Exec(``).Error)
	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).InsertBatch(``, 2, nil).Error)
	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).ParallelQuery(batis.ParallelQuery{}).Error)
	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).PagingQuery(batis.PagingQuery{}).Error)
	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).FetchQuery(batis.FetchQuery{}).Error)
	expectExecutorConflictError(t, db.Update("products", map[string]any{"price": 0}, batis.Where("product_name= 'TV'")).AssociateQuery(batis.AssociateQuery{}).Error)
}

func TestUpdateRowsAffected(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()
	l := getProductsList()
	{
		rowsAffected, err := db.Update("products", map[string]any{"price": 0},
			batis.Where("product_name = #{name}", batis.Param("name", "TV"))).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(1), rowsAffected)
	}
	{
		rowsAffected, err := db.Update("products", map[string]any{"price": 0},
			batis.Where("1 = 1")).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(len(l)), rowsAffected)
	}
}

func TestUpdateContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	expectContextCanceled(t, db.WithContext(ctx).Update("products", map[string]any{"price": 0}, batis.Where("1=1")).Error)
}

func TestUpdateTraceId(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()
	{
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithTraceId("id").Update("products", map[string]any{"price2": 0}, batis.Where("1=1"))
		w.expectTraceId(t, "id")
	}
	{
		ctx := batis.WithTraceId(context.Background(), "ctx")
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithContext(ctx).Update("products", map[string]any{"price2": 0}, batis.Where("1=1"))
		w.expectTraceId(t, "ctx")
	}
}
