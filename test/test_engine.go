package main

import (
	"encoding/json"
	"fmt"
	"github.com/koyeo/gobatis/engine"
	"github.com/koyeo/gobatis/test/mapper"
)

func main() {
	instance, err := engine.NewEngine("./gobatis.xml")
	if err != nil {
		panic(err)
	}
	err = instance.DB.Ping()
	if err != nil {
		fmt.Println("ping error:", err)
		return
	}

	defer func() {
		err = instance.DB.Close()
		if err != nil {
			fmt.Println("数据库关闭失败:", err)
			return
		}
	}()

	//instance.DB.Conn()
	fmt.Println("数据库连接成功!")

	//userMapper := new(mapper.UserMapper)
	userMapper := new(mapper.UserMapper)
	err = instance.BindMapper(userMapper)
	if err != nil {
		fmt.Println("mapper 绑定错误:", err)
		return
	}

	user, err := userMapper.GetUser(1)
	if err != nil {
		fmt.Println("查询错误:", err)
		return
	}
	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println("解析错误:", err)
		return
	}
	fmt.Println(string(res))
}
