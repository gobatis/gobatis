package postgres

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open(dsn string) *Dialector {
	d := &Dialector{}
	d.db, d.err = sql.Open("pgx", dsn)
	return d
}

type Dialector struct {
	db  *sql.DB
	err error
}

func (d Dialector) DB() (*sql.DB, error) {
	return d.db, d.err
}
