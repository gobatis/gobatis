---
sidebar_position: 3
---

# Insert

**postgres:**

```go
type User struct {
	Id *int64
	Username string
}

err = batis.Insert("users", user, batis.Returning("id")).Scan(&user.Id).Error
```


**mysql:**

```go
r = batis.Insert("users", user, batis.Returning("id"))
err = r.Error
id = r.LastInsertId
```


## InsertBatch

```go
err = batis.InsertBatch("users", 10, users).Error
```

## Nested Struct

```go
type Id struct {
	Id *int64
}

type User struct {
	Id
	Username string
}

err = batis.Insert("users", user, batis.Returning("id")).Scan(&user.Id).Error
```