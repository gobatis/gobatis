# 错误处理


## 查询错误

在 Gobatis 中，执行查询语句（`<select>`）时，无论是单条数据（`sql.Row`）还是多条数据（`sql.Rows`）查询，都会检查结果集是否为空，为空时返回 `sql.ErrNoRows` 错误。

在检查数据不存在则执行添加操作的场景下，有如下两种错误处理方式：

**方式一：**

```go
user, err := mapper.User.GetUserByUsername("tony")
if err != nil{
    if err != sql.NoErrRows{
        return nil, fmt.Errorf("query error: %s", err)
    }
} else {
    // 用户名已存在
    return user, nil
}

// 执行创建操作
```

**方式二：**
```go
user, err := mapper.User.GetUserByUsername("tony")
if err == nil{
    // 用户名已存在
    return user, nil
} else if err != nil {
	return nil, fmt.Errorf("query error: %s", err)
}

// 执行创建操作
```