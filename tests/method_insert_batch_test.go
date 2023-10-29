package tests

import (
	"context"
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

// Testing batch insertion of data, including checking the number of affected rows,
// verifying the inserted data, returning the last insert ID, and scanning all auto-incremented IDs.
func TestInsertBatch(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()

	{
		l := getProductsList()

		err := db.Affect(len(l)-1).InsertBatch("products", 2, l).Error
		require.Error(t, err)

		affected, err := db.Affect(6).InsertBatch("products", 2, l).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(6), affected)

		var result []*Product
		err = db.Query(`select * from products`).Scan(&result).Error
		require.NoError(t, err)

		compareProducts(t, result, l)

		err = db.Exec("delete from products").Error
		require.NoError(t, err)
	}

	{
		cleanProducts(t)
		l := getProductsList()
		var products []*Product
		affected, err := db.Affect(len(l)).InsertBatch("products", 2, l,
			batis.Returning("*")).Scan(&products).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(len(l)), affected)

		compareProducts(t, l, products)
	}
}

func TestInsertBatchAffect(t *testing.T) {

	defer func() {
		cleanProducts(t)
	}()

	l := getProductsList()

	expectAffectConstrictError(t, db.Affect(len(l)-1).InsertBatch("products", 2, l).Error)

	{
		var count int64
		require.NoError(t, db.Query(`select count(1) from products`).Scan(&count).Error)
		require.Equal(t, int64(0), count)

		require.NoError(t, db.Affect(len(l)).InsertBatch("products", 2, l).Error)
		require.NoError(t, db.Query(`select count(1) from products`).Scan(&count).Error)
		require.Equal(t, int64(len(l)), count)
		cleanProducts(t)

		require.NoError(t, db.Affect(len(l)).InsertBatch("products", 100, l).Error)
		require.NoError(t, db.Query(`select count(1) from products`).Scan(&count).Error)
		require.Equal(t, int64(len(l)), count)
	}

	{
		cleanProducts(t)
		var items []*Product
		err := db.Affect(len(l)).InsertBatch("products", 1, l, batis.Returning("*")).Scan(&items).Error
		require.NoError(t, err)
	}
}

func TestInsertBatchExecutorConflict(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()
	l := getProductsList()
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).Insert("products", &Product{}).Error)
	cleanProducts(t)
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).Update("products", nil, batis.Where("")).Error)
	cleanProducts(t)
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).Query(``).Error)
	cleanProducts(t)
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).Exec(``).Error)
	cleanProducts(t)
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).InsertBatch(``, 2, nil).Error)
	cleanProducts(t)
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).ParallelQuery(batis.ParallelQuery{}).Error)
	cleanProducts(t)
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).PagingQuery(batis.PagingQuery{}).Error)
	cleanProducts(t)
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).FetchQuery(batis.FetchQuery{}).Error)
	cleanProducts(t)
	expectExecutorConflictError(t, db.InsertBatch("products", 1, l).AssociateQuery(batis.AssociateQuery{}).Error)
}

func TestInsertBatchRowsAffected(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()
	l := getProductsList()
	{
		rowsAffected, err := db.InsertBatch("products", 1, l).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(len(l)), rowsAffected)
	}
	{
		cleanProducts(t)
		var items []*Product
		rowsAffected, err := db.InsertBatch("products", 1, l, batis.Returning("*")).Scan(&items).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(len(l)), rowsAffected)
	}
}

func TestInsertBatchContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	l := getProductsList()
	expectContextCanceled(t, db.WithContext(ctx).InsertBatch("products", 1, l).Error)
}

func TestInsertBatchTraceId(t *testing.T) {
	l := getProductsList()
	{
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithTraceId("id").InsertBatch("products2", 1, l)
		w.expectTraceId(t, "id")
	}
	{
		ctx := batis.WithTraceId(context.Background(), "ctx")
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithContext(ctx).InsertBatch("products2", 1, l)
		w.expectTraceId(t, "ctx")
	}
}

func TestInsertBatchColumnTag(t *testing.T) {

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

	require.NoError(t, s.InsertBatch("products", 1, []*ProductJ{i}).Error)

	var n *ProductJ

	require.NoError(t, s.Query(`select * from products where product_name = #{ name }`,
		batis.Param("name", i.JProductName)).Scan(&n).Error)

	require.True(t, *n.JId > 0)
	n.JId = nil

	compareProductJ(t, i, n)
}
