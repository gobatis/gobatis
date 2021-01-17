package main

import (
	"fmt"
	"github.com/koyeo/mybatis.go/engine"
	"github.com/koyeo/mybatis.go/test/mapper"
)

func main() {
	instance, err := engine.NewEngine("./mybatis-config.xml")
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

	fmt.Println("数据库连接成功!")

	userMapper := new(mapper.UserMapper)
	err = instance.BindMapper(userMapper)
	if err != nil {
		fmt.Println("mapper 绑定错误:", err)
		return
	}

}
