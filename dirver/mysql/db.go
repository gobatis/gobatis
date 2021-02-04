package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/koyeo/gobatis/schema"
)

func InitDB(dataSource *schema.DataSource) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dataSource.Username(),
		dataSource.Password(),
		dataSource.URL(),
		dataSource.Database(),
	))
	if err != nil {
		err = fmt.Errorf("mysql connnet error: %s", err)
		return
	}
	return
}
