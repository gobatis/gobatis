# Engine 引擎

`Engine` 是 Gobatis 的操作入口，一个 `Engine` 下可以挂载一个 master `DB` 和多个 slave `DB`。

## 初始化

初始化 Postgresql Engine:

```go
engine := gobatis.NewPostgresql("postgresql://user:password@127.0.0.1:5432/database")
```

初始化 MySQL Engine:

```go
engine := gobatis.NewMySQL("mysql://user:password@127.0.0.1:3306/database")
```

初始化其他数据库 Engine:

```go
db := gobatis.NewDB("clickhouse", ("clickhouse://user:password@127.0.0.1:5672/database")
engine := gobatis.NewEngine(db)
```

## 读写分离（暂未实现 salve 逻辑）

可读写数据库实例：

```go
master := engnie.Master()
```

只读数据库实例：

```go
// 添加只读数据库实例
engine.Add(slaveDB...)

// 获取只读实例列表
slaves := engnie.Slaves()

// 根据策略分配只读实例
slave = engnie.Slave()
```

## 注册 SQL 文件

通过 SQL 目录的 bundle 注册 SQL 文件，并且连接数据库 master 实例。

```go
engine.Init(gobatis.NewBundle("./sql"))
```

## 绑定 Mapper

为抽象 mapper 实例赋值数据库操作方法。

```go
engnie.BindMapper(mappers...)
```

## 日志设置

通过 `SetLogLevel` 和 `SetLogger` 方法设置日志输出等级和打印日志的 Logger 实例，详情参看[日志](log.html)。

```go
// 设置日志等级
engnie.SetLogLevel(gobatis.InfoLevel)

// 设置 Logger
engnie.SetLogger(logger)
```

## Reflect Tag

改变默认 Struct 字段反射标签。

```go
engine.SetTag("json")

// 便可使用 json 作为实体类的反射标签
type User struct{
    Name stirng `json:"name"`
}
```

## Close

释放 Engine 资源。

```go
engine.Close()
```