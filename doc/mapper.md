# Mapper 映射

Mapper 是 xml 中定义的 SQL 方法在 Go 代码中的映射。

## 定义

```go
type userMapper struct{
// 普通方法
CreateUser func (user *entity.User)(rows int, err error)
// 事务方法
UpdateUser func (tx *sql.Tx, user *entity.User)(rows int, err error)
// 自定义 Context 
DeleteUser func (ctx context.Context, tx *sql.Tx, id int)(rows int, err error)
}
```

可注入的非 SQL 方法参数类型：

参数 | 描述
:---|---
`*sql.Tx` |  用于事务
`*gobatis.DB` | 常用于数据库迁移
`context.Context` | 使用 Context

::: warning 注意
`*sql.Tx` 和 `*gobatis.DB` 在同一个方法的入参中，只能出现一个。
:::

## 方法类型及返回结果

方法类型 | XML 标签 | 返回结果
:---|:---|---
添加 | `<insert>` | `rows(int)` 影响行数，`(err error)` 执行错误
更新 | `<update>` | `rows(int)` 影响行数，`(err error)` 执行错误
删除 | `<delete>` | `rows(int)` 影响行数，`(err error)` 执行错误
查询 | `<select>` | result 属性选中的结果，`(err error)` 执行错误

::: tip 注意 所有的 mapper 方法都要求以 error 作为最后一个返回参数，其余参数为选填。
:::

如下所示：

```go
type userMapper struct{
CreateUser func (user *entity.User)(rows int, err error) // 正确
CreateUser func (user *entity.User)(rows int, err error) // 正确
CreateUser func (user *entity.User)(err error, rows int) // 错误
CreateUser func (user *entity.User) // 错误
}
```

## 迁移 Mapper

作为迁移数据库建表用的 Mapper，SQL 方法和普通方法一样。

sql:

```xml

<insert id="CreateUserTable">
    create table if no exist ...
</insert>
```

mapper:

```go
// Migration mapper
type migrationMapper struct{
CreateUserTable func (db *gobatis.DB) error
}
```

要求 Mapper 方法函数只能有一个入参 `*gobatis.DB` 和一个出参 `error`。可使用 `DB.Migrate()` 方法自动调用 mapper 下的所有方法。如下所示 :

```go
Migration := &migrationMapper{}
engine.BindMapper(Migration)
engnie.Master().Migrate(Migration) // 向主库中执行迁移语句
```