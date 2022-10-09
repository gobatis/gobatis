# Gobatis

适用于 Golang，基于 MyBatis 标签语法的持久层框架。

## 特性

- 基于 database/sql 实现与数据库无关；
- 适用 golang 语法，支持多参数传递和多参数返回；
- 支持 xml 文件打包成 bin 文件；
- 支持 mybatis 的标签语法；

## 示例
1. 单条数据手动插入
2. 单条数据自动映射插入
3. 多条数据手动插入
4. 多条数据自动映射插入
5. 单结果简单查询
6. 多结果简单查询
7. 复杂条件分页查询
8. 复杂子查询
9. 数据表迁移

## 命令行工具
1. Dump 表结构
2. 数据库比较
3. 数据库迁移及语句生成
4. 数据库自动备份

## TODO

- 持续性测试；
- 优化错误返回信息，增加上下文提示；

## 支持扫描数据类型

- int,int8,int16,int32,int64
- uint,uint8,uint16,uint32,uint64
- float32,float64
- time.Time
- decimal.Decimal

## 支持顶级 SQL 标签

- insert
- select
- update
- delete

## 支持动态 SQL 标签

- if
- choose、when、otherwise
- trim、where、set
- foreach

## 安装

```
$ go get -v github.com/gobatis/gobatis
```

## 初始化

参考 engine_test.go -> TestEngine。

