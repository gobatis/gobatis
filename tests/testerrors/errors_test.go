package testerrors

import (
	"errors"
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/stretchr/testify/require"
)

var db *batis.DB

func initDB(t *testing.T) {
	var err error
	db, err = batis.Open(postgres.Open("postgresql://test:test@127.0.0.1:8432/gobatis-test-db?connect_timeout=10&sslmode=disable"))
	require.NoError(t, err)
}

func TestErrors(t *testing.T) {
	initDB(t)
	testParallelQueryErr(t)
	testPagingQueryErr(t)
	testInsertBatchErr(t)
	testUpdateErr(t)
}

func testParallelQueryErr(t *testing.T) {
	err := db.ParallelQuery(batis.ParallelQuery{
		SQL:    "",
		Params: nil,
		Scan:   nil,
	}).Error
	require.True(t, errors.Is(err, batis.ErrNoScanDest))

	err = db.ParallelQuery().Error
	require.True(t, errors.Is(err, batis.NoParallelQueryerErr))
}

func testPagingQueryErr(t *testing.T) {
	err := db.PagingQuery(batis.PagingQuery{
		Limit: 0,
	}).Error
	require.True(t, errors.Is(err, batis.InvalidLimitErr))

	err = db.PagingQuery(batis.PagingQuery{
		Limit: -1,
	}).Error
	require.True(t, errors.Is(err, batis.InvalidLimitErr))

	err = db.PagingQuery(batis.PagingQuery{
		Limit: 1,
	}).Error
	require.True(t, errors.Is(err, batis.InvalidPagingScanDestErr))
}

func testInsertBatchErr(t *testing.T) {
	err := db.InsertBatch("products", 0, nil).Error
	require.True(t, errors.Is(err, batis.InvalidInsertBatchBatchErr))

	err = db.InsertBatch("products", 1, nil).Error
	require.True(t, errors.Is(err, batis.InvalidInsertBatchDataErr))

	err = db.InsertBatch("products", 1, "hello").Error
	require.True(t, errors.Is(err, batis.InvalidInsertBatchDataTypeErr))
}

func testUpdateErr(t *testing.T) {
	err := db.Update("products", nil, batis.Where("id = 0"), batis.Where("id = 1")).Error
	require.True(t, errors.Is(err, batis.PrepareSQLRawErr))

	err = db.Update("products", nil, batis.Returning("*"), batis.Returning("*")).Error
	require.True(t, errors.Is(err, batis.PrepareSQLRawErr))

	err = db.Update("products", nil, batis.Where("*"), batis.OnConflict("", "")).Error
	require.True(t, errors.Is(err, batis.PrepareSQLRawErr))
}
