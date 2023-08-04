package main

import (
	gbt "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/gobatis/gobatis/example/entity"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	db, err := gbt.Open(postgres.Open("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable"))
	if err != nil {
		return
	}
	
	user := entity.User{}
	err = db.Insert(gbt.Background(), "users", user, gbt.OnConflict("do update set a = columnd.a")).Error()
	if err != nil {
		return
	}
}
