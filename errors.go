package batis

import (
	"errors"
)

var (
	ErrRecordNotFound     = errors.New("record not found")
	ErrInvalidTransaction = errors.New("invalid transaction")
	NoScanDestErr         = errors.New("expect 1 scan dest, got nil")
)

var (
	InvalidLimitErr               = errors.New("invalid limit")
	InvalidPagingScanDestErr      = errors.New("paging query scan expect 2 dest")
	NoParallelQueryerErr          = errors.New("no queryer")
	InvalidInsertBatchBatchErr    = errors.New("expect InsertBatch batch > 0")
	InvalidInsertBatchDataErr     = errors.New("invalid InsertBatch data")
	InvalidInsertBatchDataTypeErr = errors.New("invalid InsertBatch data type")
	PrepareSQLRawErr              = errors.New("prepare sql error")
	//NoScanDestErr                 = NoScanDestErr
)
