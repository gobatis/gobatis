# SQL测试用例

## 用例1

**测试点：**

* 基础 select 表达式
* 任意匹配符
* 包裹字段
* 等于比较
* 大于等于比较
* 小于等于比较
* 字符串值
*

**正确用例：**

```sql
select * from users where `role`='admin' age >=18 and sex= 1 and `status` = 'abc' limit 2,1; 
```

**异常用例：**

```sql

```




