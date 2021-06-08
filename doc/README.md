---
home: true
heroImage: /assets/img/logo.png
heroText: Gobatis
tagline: 适用于 golang 的持久层框架
actionText: 快速上手 →
actionLink: /getting-started
features:
- title: 简洁高效
  details: 不做多余的事情，本质上只是做 SQL 语句的拼装，避免因侵入用户代码而出错。
- title: 体验至上
  details: 以人为本，从命名到工具链的打磨，为了愉悦编程而不断探索。
- title: 传承创新
  details: 以 MyBatis 语法为参照，结合 Golang 语言特点，自成一套体系。
footer: MIT Licensed | Copyright © 2020-present gobatis.co
---


## 安装

```
go get -v github.com/gobatis/gobatis
```

## 用法概览

**user.go:**
```go
type User struct{
	Id   int    `sql:"id"`
	Name string `sql:"name"`
	Age  string `sql:"age"`
}
```

**mapper.go:**
```go
var UserMapper = &usermapper{}

type userMapper struct{
    CreateUser func (user *User) error                         // 创建用户
    GetUserById func (id int) (name string,age int, err error) // 查询单个用户
    GetUsers func()([]*User, error)                            // 查询多个用户
}
```

**user.xml:**

```xml
<insert id="CreateUser" parameter="user">
    insert into users(name,age) values( #{user.Name}, #{user.Age} );
</insert>

<select id="GetUserById" parameter="id:int" result="name,age"> 
    select name,age from users where id=${id};
</select>

<select id="GetUsers" resultType="[]struct"> 
    select * from users
</select>
```