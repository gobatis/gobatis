---
sidebar_position: 7
---

# Exec

```go
batis.Exec(`update users set age = #{user.Age} where id = #{user.Id}`, batis.Param("user",user))
```