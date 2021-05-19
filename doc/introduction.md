# 简介

## 什么是 GoBatis？

[MyBatis](https://mybatis.org/mybatis-3) 是一款优秀的持久层框架，致敬！

Gobatis 借鉴了 MyBatis 的 XML 语法，结合 Golang 的语法特点实现，旨在简化 Gopher 日程的数据库操作。除开 XML 部分外，本质上 Gobatis 和 MyBatis 完全是两个不同的产品。

所以我们将 MyBatis 中相同的 XML 使用部分迁移至本文档中，对于不同的或不支持的特性给予说明，在使用时请以本文档为准。

## 特性

* 用户体验优先
* 全面的单元测试
* 基于 database/sql 实现与数据库无关；
* 支持 xml 组装动态 SQL , 适用 golang 语法，支持多参数传递和多参数返回；
* 支持 xml 文件打包成 bin 文件；
* 支持 null 数据扫描；
* 采用 decimal.Decimal 进行浮点数运算；
* 数据类型安全转换，自动检查整型溢出；

## 未来发展

Gobatis 有着自己的发展路线，不会完全遵循 MyBatis 的特性去研发。

