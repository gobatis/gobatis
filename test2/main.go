package main

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/bundle"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	engine := gobatis.NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/angel?connect_timeout=10&sslmode=disable")
	engine.SetBundle(bundle.Dir(filepath.Join(pwd, "test")))
	////engine.BindMapper(&testMapper)
	//
	err = engine.Init()
	if err != nil {
		log.Println("DB init error:", err)
		return
	}
	err = engine.Master().Ping()
	if err != nil {
		log.Println("DB ping error:", err)
		return
	}
	defer func() {
		err = engine.Master().Close()
		if err != nil {
			log.Println("DB close error:", err)
			return
		}
	}()

	engine.Call("SelectTestById", 1, "hello", 2)
}
