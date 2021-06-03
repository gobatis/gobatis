# 快速上手

## 安装

``` sh
# 引入库
go get -v github.com/gobatis/gobatis

# 安装 XML 语法映射文件
sh -c "$(curl -fsSL https://gobatis.co/dtd/dtd.sh)"
```

## 项目示例

::: tip 示例地址
[http]()
:::

### 目录结构

```
project
├── entity
│   └── user.go   // 存放实体类
├── mapper
│   ├── mapper.go // 定义 mapper
└── sql
│   └── user.xml  // sql 文件    
├── main.go       // 程序入口    
```

### Main

> project/main.go

``` go
package main

import (
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
	err := engine.BindMapper(
		mapper.User,
	)
	if err != nil {
		log.Panicln("init error:", err)
	}
	defer func() {
		engine.Close()
	}()
	
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
	
	// ...
}
```

## Entity

> project/entity/user.go

```go
package entity

import "time"

type User struct {
	Id        int        `sql:"id"`
	Name      string     `sql:"name"`
	Age       int        `sql:"age"`
	From      string     `sql:"from"`
	Vip       bool       `sql:"vip"`
	CreatedAt *time.Time `sql:"created_at"`
}
```

## Mapper

> project/mapper/mapper.go

``` go 
package mapper

import (
	"github.com/gobatis/gobatis/example/entity"
	"time"
)

var (
	User = &userMapper{}
)

type userMapper struct {
	AddUser         func(user *entity.User) (rows int, err error)
	AddUserReturnId func(user *entity.User) (id int, createdAt time.Time, err error)
	GetUserById     func(id int64) (name string, age int, err error)
	GetUserByName   func(name string) (user *entity.User, err error)
	GetUserByFrom   func(places []string) ([]*entity.User, error)
	QueryUsers      func(m map[string]interface{}) ([]*entity.User, error)
}
```

## SQL

> project/sql/user.xml

``` xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//gobatis.co//DTD Mapper 1.0//EN" "gobatis.co/dtd/mapper.dtd">

<mapper>
    <insert id="AddUser" parameter="user:struct">
        insert into users( name, age, "from", vip ) values( #{user.Name}, #{user.Age}, #{user.From}, ${user.Vip})
    </insert>

    <select id="AddUserReturnId" parameter="user" result="id,created_at">
        insert into users(name,age,"from",vip ) values(#{user.Name},#{user.Age},#{user.From},${user.Vip} )
        returning id,created_at;
    </select>

    <select id="GetUserById" parameter="id:int64" result="name,age">
        select * from users where id=#{id};
    </select>

    <select id="GetUserByName" parameter="name">
        select * from users where name=#{name};
    </select>

    <select id="GetUserByFrom" parameter="places:[]string">
        select * from users where "from" in
        <foreach index="index" item="item" collection="places" open="(" separator="," close=")">
            #{item}
        </foreach>
    </select>

    <select id="QueryUsers" parameter="m" result="*">
        select * from users
        <where>
            <if test='m["name"] != nil'>
                name = #{m["name"]}
            </if>
            <if test="m['from'] != nil">
                and "from" = #{m['from']}
            </if>
            <if test="m['vip'] != nil">
                and vip = #{m['vip']}
            </if>
        </where>
    </select>
</mapper>
```

