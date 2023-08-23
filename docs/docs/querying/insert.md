---
sidebar_position: 3
---

# Insert


with returning:
```go
err = batis.Insert("users", user, batis.Returning("id")).Scan(&user.Id).Error
```


with lastInsertedId:
```go
r = batis.Insert("users", user, batis.Returning("id"))
err = r.Error
id = r.LastInsertId
```


## InsertBatch

```go
err = batis.InsertBatch("users", 10, users).Error
```