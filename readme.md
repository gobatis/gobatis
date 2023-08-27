# Gobatis

面向工程的 ORM，采用 MyBatis 的标签语法。

## 特性

* 便捷的增、删、查、改操作；
* 便捷的变量绑定；
* 支持语句 Debug
* 支持执行 hook.
* 支持事务、Prepare操作；
* 支持快捷方法：分页查询、批量插入；
* 更好的并发支持；

## 模块

* SQL 执行；
* 事务；
* 日志，位置打印；
* Hook；
* 并发；

## Tests

```go
    type User struct {
		Id     int64
		Name   string
		Posts  []string
		Orders []string
	}

	type Post struct {
		Id   int64
		Tags []string
	}

	var users []User

    db.Must().Query(`select * from users`).LooseScan(&users, "$..Posts");

	db.Must().Exec(`select`)

	userIds := mapping.Merge(users, func())

	db.Query().LooseScan(batis.LooseDest(&users, "$..Posts","$..Orders")).Error

    db.Query(`select * from posts where user_id in #{userIds}`,batis.Param("userIds", userIds)).Link(&users, "user_id => $..Id", "$..Posts").Error

	postIds := mapping.Map(users, func())

	db.Query(`select * from tags where post_id in #{postIds}`, batis.Param("postIds", postIds)).Link(&users, "user_id => $..Posts[*].Id", "$..Post[*].Tags").Error

    db.Query(`select * from orders where user_id in #{userIds}`,batis.Param("userIds", userIds)).Link(&users, "user_id => $..Id", "$..Orders").Error
```