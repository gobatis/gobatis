package batis

import (
	"database/sql"
)

func NewDB2(db *sql.Conn, traceId string) *DB2 {
	return &DB2{Conn: db, traceId: traceId}
}

var _ Conn2 = (*DB2)(nil)

type DB2 struct {
	*sql.Conn
	traceId string
}

func (d *DB2) Close() error {
	return d.Conn.Close()
}

func (d *DB2) TraceId() string {
	return d.traceId
}

func (d *DB2) IsTx() bool {
	return false
}
