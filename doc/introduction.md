# 简介

## 什么是 GoBatis？

Java 的 [MyBatis](https://mybatis.org/mybatis-3) 是一款优秀的持久层框架，它支持自定义 SQL、存储过程以及高级映射。MyBatis 免除了几乎所有的 JDBC
代码以及设置参数和获取结果集的工作。MyBatis 可以通过简单的 XML 或注解来配置和映射原始类型、接口和 Java POJO（Plain Old Java Objects，普通老式 Java 对象）为数据库中的记录。

Gobatis 借鉴了 MyBatis 的 XML 语法，结合 Golang 的语法特点实现，旨在简化 Gopher 日程的数据库操作。除开 XML 部分外，本质上 Gobatis 和 MyBatis 完全是两个不同的产品。

所以我们将 MyBatis 中相同的 XML 使用部分迁移至本文档中，同时对于不同的，或不支持的特性给予说明，在使用时请以本文档为准。

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

