package main

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/bundle"
	"log"
)

type UserMapper struct {
}

func main() {
	engine := gobatis.NewPostgresql("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
	engine.SetBundle(bundle.Dir("./"))
	err := engine.Init()
	if err != nil {
		log.Println("Gobatis init error", err)
		return
	}
}
