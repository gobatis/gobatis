package executor

import (
	"database/sql"
)

func NewDB(db *sql.Conn, traceId string) *DB {
	return &DB{Conn: db, traceId: traceId}
}

var _ Conn = (*DB)(nil)

type DB struct {
	*sql.Conn
	traceId string
}

func (d *DB) Close() error {
	return d.Conn.Close()
}

func (d *DB) TraceId() string {
	return d.traceId
}

func (d *DB) IsTx() bool {
	return false
}
