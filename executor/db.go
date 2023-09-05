package executor

import (
	"database/sql"
)

func NewDB(db *sql.DB) *DB {
	return &DB{DB: db}
}

var _ Conn = (*DB)(nil)

type DB struct {
	*sql.DB
}

func (D *DB) IsTx() bool {
	return false
}
