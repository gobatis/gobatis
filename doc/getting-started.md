# 快速上手

## 安装

``` sh
# 引入库
go get -u github.com/gobatis/gobatis

# 安装 XML 语法映射文件
sh -c "$(curl -fsSL https://gobatis.co/dtd/dtd.sh)"
```

示例地址：[https://github.com/gobatis/gobatis/tree/master/example](https://github.com/gobatis/gobatis/tree/master/example)


## 目录结构

```
project
├── entity
│   └── user.go        // 存放实体类
├── mapper
│   ├── mapper.go      // 定义 mapper
└── sql
│   └── migration.xml  // 迁移语句    
│   └── user.xml       // users 表操作语句    
├── main.go            // 程序入口    
├── table.sql          // 示例数据表    
```

## Main

::: tip 文件路径
project/main.go
:::

``` go
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
	engine := gobatis.NewPostgresql("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
	err := engine.Init(gobatis.NewBundle("./sql"))
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
	
    // ...
}

```

## Entity

::: tip 文件路径
project/entity/user.go
:::

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

::: tip 文件路径
project/mapper/mapper.go
:::

``` go 
package mapper

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/example/entity"
	"time"
)

var (
	User      = &userMapper{}
	Migration = &migrationMapper{}
)

type migrationMapper struct {
	CreateTable func(db *gobatis.DB) error
}

type userMapper struct {
	AddUser         func(user *entity.User) (rows int, err error)
	AddUserReturnId func(user *entity.User) (id int, createdAt time.Time, err error)
	UpdateUser      func(id int, vip bool) (rows int, err error)
	GetUserById     func(id int64) (name string, age int, err error)
	GetUserByName   func(name string) (user *entity.User, err error)
	GetUserByFrom   func(places []string) ([]*entity.User, error)
	QueryUsers      func(m map[string]interface{}) ([]*entity.User, error)
	DeleteUsers     func(id int) (rows int, err error)
}
```

## SQL

::: tip 文件路径
project/sql/migration.xml
:::

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//gobatis.co//DTD Mapper 1.0//EN" "gobatis.co/dtd/mapper.dtd">

<mapper>
    <insert id="CreateTable">
        create schema if not exists public;

        create table if not exists users(
        id serial constraint users_pk primary key,
        name varchar,
        age int,
        "from" varchar,
        vip bool,
        created_at timestamp default current_timestamp
        );
    </insert>
</mapper>
```


::: tip 文件路径
project/sql/user.xml
:::

``` xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//gobatis.co//DTD Mapper 1.0//EN" "gobatis.co/dtd/mapper.dtd">

<mapper>
    <insert id="AddUser" parameter="user:struct">
        insert into users( name, age, "from", vip ) values( #{user.Name}, #{user.Age}, #{user.From}, ${user.Vip})
    </insert>

    <update id="UpdateUser" parameter="id,vip">
        update users set vip = #{vip} where id = #{ id };
    </update>

    <!--  &lt; 转义小于符号  -->
    <delete id="DeleteUsers" parameter="id">
        delete from users where id &lt;= #{id};
    </delete>

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

