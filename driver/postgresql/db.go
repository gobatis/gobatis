package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const PGX = "pgx"

func InitDB(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open(PGX, dsn)
	if err != nil {
		err = fmt.Errorf("postgresql connnet error: %s", err)
		return
	}
	return
}
