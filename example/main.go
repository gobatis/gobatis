package main

import (
	"encoding/json"
	"fmt"
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/example/entity"
	"github.com/gobatis/gobatis/example/mapper"
	"log"
)

func main() {
	engine := gobatis.NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/gobatis?connect_timeout=10&sslmode=disable")
	engine.BindSQL(gobatis.NewBundle("./sql"))
	if err := engine.Init(); err != nil {
		log.Panicln("init error:", err)
	}
	
	if err := engine.Master().Ping(); err != nil {
		log.Panicln("ping error:", err)
	}
	err := engine.BindMapper(
		mapper.User,
	)
	if err != nil {
		log.Panicln("init error:", err)
	}
	defer func() {
		engine.Close()
	}()
	
	// AddUser
	rows, err := mapper.User.AddUser(&entity.User{
		Name: "Tom",
		Age:  18,
		From: "venus",
		Vip:  true,
	})
	if err != nil || rows != 1 {
		log.Panicf("Call AddUser error: <%v, %d>", err, rows)
	}
	
	// AddUserReturnId
	id, createdAt, err := mapper.User.AddUserReturnId(&entity.User{
		Name: "Tom",
		Age:  18,
		From: "venus",
		Vip:  true,
	})
	if err != nil {
		log.Panicf("Call AddUserReturnId error: %v", err)
	}
	fmt.Printf("id:%d, createdAt:%v", id, createdAt)
	
	// GetUserById
	name, age, err := mapper.User.GetUserById(int64(id))
	if err != nil {
		log.Panicf("Call GetUserById error: %v", err)
	}
	if name != "Tom" {
		log.Panicf("GetUserById expect name=Tom: got:%s", name)
	}
	if age != 18 {
		log.Panicf("GetUserById expect age=18: got:%d", age)
	}
	
	user, err := mapper.User.GetUserByName("Tom")
	if err != nil {
		log.Panicf("Call GetUserByName error: %v", err)
	}
	fmt.Println("Call GetUserByName:")
	printJson(user)
	
	users, err := mapper.User.GetUserByFrom([]string{"venus"})
	if err != nil {
		log.Panicf("Call GetUserByFrom error: %v", err)
	}
	fmt.Println("Call GetUserByFrom:")
	printJson(users)
	
	users, err = mapper.User.QueryUsers(map[string]interface{}{
		"name": "Tom",
	})
	if err != nil {
		log.Panicf("Call QueryUsers error: %v", err)
	}
	fmt.Println("Call QueryUsers:")
	printJson(users)
}

func printJson(val interface{}) {
	d, _ := json.MarshalIndent(val, "", "\t")
	fmt.Println(string(d))
}
