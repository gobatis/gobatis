package executor

import (
	"context"
	"database/sql"
)

type Conn interface {
	IsTx() bool
	TraceId() string
	Close() error
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}
