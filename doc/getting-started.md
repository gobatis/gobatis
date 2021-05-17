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
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 迁移 schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1) // 根据整形主键查找
  db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

  // Update - 将 product 的 price 更新为 200
  db.Model(&product).Update("Price", 200)
  // Update - 更新多个字段
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - 删除 product
  db.Delete(&product, 1)
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
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "gobatis.co/dtd/config.dtd">

<configuration>
    <module base="github.com/gobatis/gobatis"/>
</configuration>
```

> sql/user.xml

``` xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "../../../dtd/mapper.dtd">

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

