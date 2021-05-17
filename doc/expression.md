# 表达式

Gobatis 表达式采用 Golang 语法实现，支持特性：

* 类型安全转换；
* 访问对象方法；
* 操作数组切片；
* 函数调用；
* 支持单引号包裹字符串。

## 基础表语法

**运算**

```
a + b > 1
#{ a + 10 }
#{ Name.A / 3600 * 360 }
<if test="User != nil && user.Status == 'normal'"></if>
```

**成员访问**

```
#{ User.Name }
#{ User.Passowrd() }
#{ Members[0].Username }
```

## 内置函数

函数名 | 描述
---|---
`len()` | 获取切片、Map的长度
`int()、int8()...` | int 类型转换
`uint()、uint8()...` | uint 类型转换


