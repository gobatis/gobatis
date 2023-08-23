---
sidebar_position: 6
---

# Update

```go
batis.Must().Update("users", map[string]any{
      "updated_at": time.Now(),
   }, batis.Where("id = #{user.Id}", batis.Param("user",user)))
```