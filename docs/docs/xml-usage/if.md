---
sidebar_position: 2
title: if
---

## Tag

```xml
<if test="">
    
</if>
```

## Usage

使用动态 SQL 最常见情景是根据条件包含 where 子句的一部分。比如：

```go
db.Query(`
    select * from blog
    where state = 'active'
    <if test="len(title) > 0">
        and title like #{title}
    </if>
`, batis.Param("title", "tom's blog"))
```