---
sidebar_position: 5
---

# Query

```go
var user *User

err = batis.Query(`select * from users where id = #{id}`,batis.Param("id",1))).Scan(&user).Error

var users []*User

err = batis.Query(`select * from users`).Scan(&users).Error
```