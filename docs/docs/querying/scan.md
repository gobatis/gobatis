---
sidebar_position: 2
---

# Scanning Rows

## Precise Match

Gobatis will check the struct which receiving query rows match the the columns exactly. 

```go
err = batis.Query(`select * from users where id = #{id}`,batis.Param("id",1))).Scan(&user).Error
```


## Loose Dest

```go
err = batis.Query(`select * from users where id = #{id}`,batis.Param("id",1))).Scan(batis.LooseDest(&user)).Error
```


