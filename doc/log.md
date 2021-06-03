# 日志

## Logger Interface

```go
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
```

实现 Logger interface 的对象即可通过 `engine.SetLogger()` 方法替换引擎默认的 logger。

## 日志方法

可在项目中使用，将项目运行信息随 SQL 日志打印。

```go
gobatis.Debugf(format string, args ...interface{})
gobatis.Infof(format string, args ...interface{})
gobatis.Warnf(format string, args ...interface{})
gobatis.Errorf(format string, args ...interface{})
```

## 日志等级

等级 | 说明
:---|---
`gobatis.StackLevel` | 打印全部日志，报错时显示出错位置调用栈信息。
`gobatis.DebugLevel` | 打印全部日志，不显示调用栈信息。
`gobatis.InfoLevel` |  打印 Info，Warn，Error 等级的日志
`gobatis.WarnLevel` | 打印 Warn，Error 等级的日志
`gobatis.ErrorLevel` | 只打印 Error 等级的日志

通过 `engine.SetLogLevel()` 改变日志输出等级：

```go
engine.SetLogLevel(gobatis.ErrorLevel)
```