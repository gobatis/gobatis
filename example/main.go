package main

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/example/entity"
	"github.com/gobatis/gobatis/example/mapper"
	"log"
)

func main() {
	engine := gobatis.NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/gobatis?connect_timeout=10&sslmode=disable")
	engine.BindSQL(gobatis.NewBundle("./sql"))
	
	err := engine.BindMapper(
		mapper.User,
	)
	if err != nil {
		log.Panicln("init error:", err)
	}
	
	if err = engine.Init(); err != nil {
		log.Panicln("init error:", err)
	}
	
	if err = engine.Master().Ping(); err != nil {
		log.Panicln("ping error:", err)
	}
	defer func() {
		engine.Close()
	}()
	
	rows, err := mapper.User.AddUser(&entity.User{
		Name: "Tom",
		Age:  18,
		From: "venus",
		Vip:  true,
	})
	if err != nil || rows != 1 {
		log.Panicf("add user error: <%s,%d>", err, rows)
	}
}
