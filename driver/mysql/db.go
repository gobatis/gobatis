package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Open(dsn string) *Dialector {
	d := &Dialector{}
	d.db, d.err = sql.Open("mysql", dsn)
	return d
}

type Dialector struct {
	db  *sql.DB
	err error
}

func (d Dialector) DB() (*sql.DB, error) {
	return d.db, d.err
}
