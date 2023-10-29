package tests

import (
	"context"
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

func TestExec(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()
	err := db.Exec(`insert into products(product_name,price) values(#{name},#{price})`,
		batis.Param("name", TV),
		batis.Param("price", 1),
	).Error
	require.NoError(t, err)
}

func TestExecAffect(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()
	expectAffectConstrictError(t, db.Affect(0).Exec(`insert into products(product_name,price) values(#{name},#{price})`,
		batis.Param("name", TV),
		batis.Param("price", 1),
	).Error)

	require.NoError(t, db.Affect(1).Exec(`insert into products(product_name,price) values(#{name},#{price})`,
		batis.Param("name", TV),
		batis.Param("price", 1),
	).Error)
}

func TestExecExecutorConflict(t *testing.T) {
	expectExecutorConflictError(t, db.Exec(`select 1`).Insert("products", &Product{}).Error)
	expectExecutorConflictError(t, db.Exec(`select 1`).Update("products", nil, batis.Where("")).Error)
	expectExecutorConflictError(t, db.Exec(`select 1`).Query(``).Error)
	expectExecutorConflictError(t, db.Exec(`select 1`).Exec(``).Error)
	expectExecutorConflictError(t, db.Exec(`select 1`).InsertBatch(``, 2, nil).Error)
	expectExecutorConflictError(t, db.Exec(`select 1`).ParallelQuery(batis.ParallelQuery{}).Error)
	expectExecutorConflictError(t, db.Exec(`select 1`).PagingQuery(batis.PagingQuery{}).Error)
	expectExecutorConflictError(t, db.Exec(`select 1`).FetchQuery(batis.FetchQuery{}).Error)
	expectExecutorConflictError(t, db.Exec(`select 1`).AssociateQuery(batis.AssociateQuery{}).Error)
}

func TestExecRowsAffected(t *testing.T) {
	defer func() {
		cleanProducts(t)
	}()
	{
		rowsAffected, err := db.Exec(`insert into products(product_name,price) values(#{name1},#{price1}),(#{name2},#{price2})`,
			batis.Param("name1", TV),
			batis.Param("price1", 1),
			batis.Param("name2", Smartwatch),
			batis.Param("price2", 2),
		).RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(2), rowsAffected)
	}
}

func TestExecContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	expectContextCanceled(t, db.WithContext(ctx).Exec(`select 1`).Error)
}

func TestExecTraceId(t *testing.T) {
	{
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithTraceId("id").Exec("")
		w.expectTraceId(t, "id")
	}
	{
		ctx := batis.WithTraceId(context.Background(), "ctx")
		w := newTraceWriter()
		db.Session(&batis.Session{Logger: traceLogger(w)}).WithContext(ctx).Exec("")
		w.expectTraceId(t, "ctx")
	}
}
