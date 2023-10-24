package batis

import (
	"context"
	"database/sql"
)

type conn interface {
	IsTx() bool
	TraceId() string
	Close() error
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func newDBConn(db *sql.Conn, traceId string) conn {
	return &dbConn{Conn: db, traceId: traceId}
}

var _ conn = (*dbConn)(nil)

type dbConn struct {
	*sql.Conn
	traceId string
}

func (d *dbConn) Close() error {
	return d.Conn.Close()
}

func (d *dbConn) TraceId() string {
	return d.traceId
}

func (d *dbConn) IsTx() bool {
	return false
}
