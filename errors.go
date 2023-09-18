package batis

import (
	"errors"
	"github.com/gobatis/gobatis/executor"
)

var (
	InvalidLimitErr               = errors.New("invalid limit")
	InvalidPagingScanDestErr      = errors.New("paging query scan expect 2 dest")
	NoParallelQueryerErr          = errors.New("no queryer")
	InvalidInsertBatchBatchErr    = errors.New("expect InsertBatch batch > 0")
	InvalidInsertBatchDataErr     = errors.New("invalid InsertBatch data")
	InvalidInsertBatchDataTypeErr = errors.New("invalid InsertBatch data type")
	PrepareSQLRawErr              = errors.New("prepare sql error")
	NoScanDestErr                 = executor.NoScanDestErr
)
