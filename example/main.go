package main

import (
	"encoding/json"
	"fmt"
	"log"

	gbt "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/example/entity"
	"github.com/gobatis/gobatis/example/mapper"
)

func main() {
	engine := gbt.NewPostgresql("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
	err := engine.Init(batis.NewBundle("./sql"))
	if err != nil {
		log.Panicln("Init error:", err)
	}
	err = engine.BindMapper(
		mapper.Migration,
		mapper.User,
	)
	if err != nil {
		log.Panicln("Bind mapper error:", err)
	}
	if err = engine.Master().Ping(); err != nil {
		log.Panicln("Ping error:", err)
	}
	defer func() {
		engine.Close()
	}()

	err = engine.Master().Migrate(mapper.Migration)
	if err != nil {
		log.Panicf("exec migration error: %v", err)
	}

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
		Vip:  false,
	})
	if err != nil {
		log.Panicf("Call AddUserReturnId error: %v", err)
	}
	fmt.Printf("id:%d, createdAt:%v", id, createdAt)

	rows, err = mapper.User.UpdateUser(id, true)
	if err != nil {
		log.Panicf("Call UpdateUser error: %v", err)
		return
	}
	if rows != 1 {
		log.Panicf("Call UpdateUser expect: rows=1, got:%d", rows)
	}

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

	rows, err = mapper.User.DeleteUsers(id)
	if err != nil {
		log.Panicf("Call DeleteUser error: %v", err)
		return
	}
	if rows <= 1 {
		log.Panicf("Call DeleteUser expect: rows>1, got:%d", rows)
	}
}

func printJson(val interface{}) {
	d, _ := json.MarshalIndent(val, "", "\t")
	fmt.Println(string(d))
}
