# 快速上手

## 安装

``` sh
$ go get -v github.com/gobatis/gobatis
```

## 目录结构

```
|- sql          // 存放 xml 映射文件
|- mapper       // 存放 mapper 映射文件
main.go         // 定义 goabtis engine
gobatis.xml     // xml 配置文件
```

## 项目启动

> main.go

``` go
package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
    engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/gobatis?connect_timeout=10&sslmode=disable")
	engine.SetBundle(bundle.Dir("test"))
	err := engine.Init()
	if err != nil {
	    log.Println("gobatis engine init error:", err)
	    return
	}
	
	err = engine.master.Ping()
	if err != nil {
	    log.Println("gobatis engine master ping error:", err)
	    return
	}
	
	defer func() {
		err = engine.Close()
		require.NoError(t, err)
	}()
	
	_testMapper := new(mapper.TestMapper)
	err = engine.BindMapper(_testMapper)
	require.NoError(t, err)
}
```

## Mapper

> mapper/user_mapper.go

``` go 
type UserMapper struct{
    InsertUser func(username string, password string) error
    GetUser func(id int64)(username,password string, err error)
}
```

## XML文件

> gobatis.xml

``` xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
        PUBLIC "-//gobatis.co//DTD Config 1.0//EN"
        "gobatis.co/dtd/config.dtd">

<configuration>
    <module base="github.com/gobatis/gobatis"/>
</configuration>
```

> sql/user.xml

``` xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//gobatis.co//DTD Mapper 1.0//EN" "gobatis.co/dtd/mapper.dtd">

<mapper>
    <insert id="CreateProduct" parameter="product:struct">
        insert into products(name,width,height,price)
        values(#{product.Name},#{product.Width},#{product.Height},#{product.Price})
    </insert>

    <select id="GetProductById" parameter="id:int64">
        select * from products where id = #{ id };
    </select>

    <select id="GetProductsById" parameter="id:int64" resultType="[]struct">
        select * from products where id >= #{ id } order by id limit 1;
    </select>
</mapper>
```

