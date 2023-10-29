package batis

import (
	"errors"
)

var (
	ErrRecordNotFound                   = errors.New("record not found")
	ErrInvalidTransaction               = errors.New("invalid transaction")
	ErrExecutorConflict                 = errors.New("executor conflict")
	ErrAffectConstrict                  = errors.New("affect constrict error")
	ErrNotSupportAffectConstraint       = errors.New("not support Affect() method")
	ErrInvalidAffectValue               = errors.New("db.Affect() only accept int type or string like 1+")
	ErrNoScanDest                       = errors.New("expect 1 scan dest, got nil")
	ErrApplyMethodOnExecutedDBChain     = errors.New("on an already executed db chain")
	ErrNotCompatibleWithTransactionMode = errors.New("not compatible with transaction mode")
	ErrNoSQLResultExists                = errors.New("no sql.Result produced")
)

var (
	InvalidLimitErr               = errors.New("invalid limit")
	InvalidPagingScanDestErr      = errors.New("paging query scan expect 2 dest")
	NoParallelQueryerErr          = errors.New("no queryer")
	InvalidInsertBatchBatchErr    = errors.New("expect InsertBatch batch > 0")
	InvalidInsertBatchDataErr     = errors.New("invalid InsertBatch data")
	InvalidInsertBatchDataTypeErr = errors.New("invalid InsertBatch data type")
	PrepareSQLRawErr              = errors.New("prepare sql error")
)
