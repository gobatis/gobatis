package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobatis/gobatis"
)

const MySQL = "mysql"

func InitDB(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open(MySQL, dsn)
	if err != nil {
		err = fmt.Errorf("mysql connnet error: %s", err)
		return
	}
	return
}

func NewEngine(dsn string) *gobatis.Engine {
	db, err := InitDB(dsn)
	if err != nil {
		panic(err)
	}
	return gobatis.NewEngine(gobatis.NewDB(db))
}
