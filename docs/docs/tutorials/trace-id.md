---
sidebar_position: 1
---

# Trace ID

```go
db = db.WithTraceId("1")

ctx := batis.WrapTraceId(context.Backgroun(), "1")
db.WithContext(ctx).Insert()
```