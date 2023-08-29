---
sidebar_position: 1
---

# Logger


## Export 

```go
db := open(config{
	LogExporter: timescaledb.LogExporter{
		Level: "warn",
		Table: "logs",
		Interval: ""
    }
})
```