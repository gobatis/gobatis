package main

import (
	g "github.com/gobatis/gobatis"
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
	db, err := g.Open(postgres.Open("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable"))
	if err != nil {
		return
	}
	
	user := entity.User{}
	err = db.Insert(g.Background(), "users", user, g.OnConflict([]string{}, "do update set a = columnd.a")).Error()
	if err != nil {
		return
	}
	
	//db.Build(
	//	
	//)
	
	//db.SetTimeout()
	
	db.Query(``, g.Param("", 1), g.Param("", 1))
}
