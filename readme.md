# Gobatis

适用于 Golang，基于 MyBatis 标签语法的持久层框架，[使用文档](https://gobatis.io)。

## 特性

- 基于 database/sql 实现与数据库无关；
- 适用 golang 语法，支持多参数传递和多参数返回；
- 支持 xml 文件打包成 bin 文件；
- 支持 mybatis 的标签语法；

## 包安装

```shell
go get -u github.com/gobatis/gobatis
```

## 命令行工具

```
go get -u github.com/gobatis/gobatis/cmd/gobatis
```

## 合适使用

## 使用示例

提供一些例子说明你是如何部署运行的环境的。

表述部署的步骤是什么，并给出示例：

给出示例 最后，以一个小 demo 结束：从外部获取数据，并使用部署好的环境展示项目。

## 运行测试用例

gobatis 合并代码进入 master 分支时，自动触发 github actions 执行测试用例，历史测试用例执行记录请访问: [集成测试任务]();

测试用例以如下步骤进行：

1. 语法测试；
2. SQL 执行结果扫描测试；
3. 并发测试；

## 内置

* 依赖项管理：go mod。
* 使用日志包：zap。

## 投稿

请阅读 CONTRIBUTING.md，了解我们的行为准则以及如何向我们提交拉取请求的过程。

## 版本

我们使用 https://semver.org/lang/zh-CN/ 进行版本控制。有关可用的版本，请:book:此存储库中的标记。

gobatis 还处于设计开发阶段，暂时未发布 1.0 版本。

## 作者

- Koyeo
- Alpha

另请参阅参与此项目的贡献者列表。

## 许可证

此项目根据 MIT 许可证获得许可 - 有关详细信息，请参阅 LICENSE.md 文件。

## 致谢

感谢所有使用本项目的人

