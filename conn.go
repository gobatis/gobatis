package batis

import (
	"context"
	"database/sql"
)

type conn interface {
	IsTx() bool
	TraceId() string
	Close() error
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*connTx, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func newDBConn(db *sql.Conn, traceId string) conn {
	return &connDB{Conn: db, traceId: traceId}
}

var (
	_ conn = (*connDB)(nil)
	_ conn = (*connTx)(nil)
)

type connDB struct {
	*sql.Conn
	traceId string
}

func (d *connDB) Close() error {
	return d.Conn.Close()
}

func (d *connDB) TraceId() string {
	return d.traceId
}

func (d *connDB) IsTx() bool {
	return false
}

func (d *connDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*connTx, error) {
	t, err := d.Conn.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return newTx(t, d.traceId), nil
}

func newTx(tx *sql.Tx, traceID string) *connTx {
	return &connTx{Tx: tx, traceId: traceID}
}

type connTx struct {
	traceId string
	*sql.Tx
}

func (t *connTx) TraceId() string {
	return t.traceId
}

func (t *connTx) BeginTx(ctx context.Context, opts *sql.TxOptions) (*connTx, error) {
	return t, nil
}

func (t *connTx) IsTx() bool {
	return true
}

func (t *connTx) Close() error {
	return nil
}
