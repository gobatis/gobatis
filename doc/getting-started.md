# 快速上手

## 安装

``` sh
# 引入库
go get -v github.com/gobatis/gobatis

# 安装 XML 语法映射文件
sh -c "$(curl -fsSL https://gobatis.co/dtd/dtd.sh)"
```

## 目录结构

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

## 项目启动

> project/main.go

``` go
package main

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
    engine := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable")
	engine.BindSQL(gobatis.NewBundle("./sql"))
	err := engine.Init()
	if err != nil {
	    log.Println("engine init error:", err)
	    return
	}
	
	err = engine.master.Ping()
	if err != nil {
	    log.Println("engine master ping error:", err)
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

> project/mapper/mapper.go

``` go 
type UserMapper struct{
    InsertUser func(username string, password string) error
    GetUser func(id int64)(username,password string, err error)
}
```

## SQL
> project/sql/user.xml

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

