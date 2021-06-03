# 事务支持

事务支持需要 Mapper 方法可接受 `*sql.Tx` 类型参数，如：

```go
type userMapper{
	UpdateUser(tx *sql.Tx) error
}
```

调用示例：

```go
User := &userMapper{}
engine.Bind(User)

// 初始化事务
tx,err := engine.Master().Begin()
if err != nil{
	return err
}

err = User.CreateUser(tx)
if err != nil{
	// 事务回滚
	_ = tx.Rollback()
	return err
}

// 事务提交
err = tx.Commit()
if err != nil{
	return err
}

// 事务释放
_ = tx.Colse()
```

涉及到事务的方法一般要更加细心地处理，Gobatis 并没有做太多封装，而是将事务操作暴露在外面，以提供更高的灵活性。