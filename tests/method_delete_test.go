package tests

import (
	"context"
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()
	{
		affected, err := db.Affect(1).Delete("products",
			batis.Where("product_name = #{name}", batis.Param("name", Smartphone))).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(1), affected)
	}
	{
		m := getProductsMap()
		var products []*Product
		err := db.Delete("products",
			batis.Where("product_name in #{names}", batis.Param("names", []string{Chair, BluetoothHeadphones})),
			batis.Returning("*"),
		).Scan(&products).Error
		require.NoError(t, err)
		compareProducts(t, []*Product{m[Chair], m[BluetoothHeadphones]}, products)
	}
}

func TestDeleteAffect(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()
	l := getProductsList()
	{
		expectAffectConstrictError(t, db.Affect(0).Delete("products", batis.Where("product_name = #{name}", batis.Param("name", TV))).Error)
	}
	{
		var count int
		require.NoError(t, db.Query(`select count(1) from products`).Scan(&count).Error)
		require.Equal(t, len(l), count)
	}
	{
		require.NoError(t, db.Affect(2).Delete("products", batis.Where("product_name in #{names}", batis.Param("names", []string{TV, Smartwatch}))).Error)
	}
	{
		var count int
		require.NoError(t, db.Query(`select count(1) from products`).Scan(&count).Error)
		require.Equal(t, len(l)-2, count)
	}
}

func TestDeleteExecutorConflict(t *testing.T) {
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).Insert("products", &Product{}).Error)
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).Update("products", nil, batis.Where("")).Error)
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).Query(``).Error)
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).Exec(``).Error)
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).InsertBatch(``, 2, nil).Error)
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).ParallelQuery(batis.ParallelQuery{}).Error)
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).PagingQuery(batis.PagingQuery{}).Error)
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).FetchQuery(batis.FetchQuery{}).Error)
	expectExecutorConflictError(t, db.Delete("products", batis.Where("1=1")).AssociateQuery(batis.AssociateQuery{}).Error)
}

func TestDeleteRowsAffected(t *testing.T) {
	prepareProducts(t)
	defer func() {
		cleanProducts(t)
	}()
	{
		rowsAffected, err := db.Delete("products", batis.Where("product_name in #{names}",
			batis.Param("names", []string{TV, Smartwatch}))).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(2), rowsAffected)
	}
	{
		rowsAffected, err := db.Delete("products",
			batis.Where("product_name in #{names}", batis.Param("names", []string{Chair, BluetoothHeadphones})),
			batis.Returning("*"),
		).Scan(nil).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(2), rowsAffected)
	}
}

func TestDeleteContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	expectContextCanceled(t, db.WithContext(ctx).Delete("products2", batis.Where("")).Error)
}

func TestDeleteTraceId(t *testing.T) {
	{
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithTraceId("id").Delete("products2", batis.Where(""))
		w.expectTraceId(t, "id")
	}
	{
		ctx := batis.WithTraceId(context.Background(), "ctx")
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithContext(ctx).Delete("products2", batis.Where(""))
		w.expectTraceId(t, "ctx")
	}
}

func TestDeleteColumnTag(t *testing.T) {
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
	require.NoError(t, s.Delete("products",
		batis.Where("product_name = #{ name }", batis.Param("name", i.JProductName)),
		batis.Returning("*"),
	).Scan(&n).Error)
	require.True(t, *n.JId > 0)
	n.JId = nil
	compareProductJ(t, i, n)
}
