---
sidebar_position: 4
---

# Delete

common:

```go
batis.Delete("users", batis.Where("id = #{id}", batis.Param("id",1)))
```

with must:
```go
batis.Must().Delete("users", batis.Where("id = #{id}", batis.Param("id",1)))
```
