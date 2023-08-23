---
sidebar_position: 1
---

# Transactions


## Visual call

```go
tx := db.Begin()
tx.Insert()

tx.Commit()

tx.Rollback()
```


## Hidden Call
```go
tx := db.Begin()
defer func() {
    if err != nil {
        tx.Rollback()
    }
}

ctx := batis.WrapTx(context.Backgroun(), tx)

db.WithContext(ctx).Insert()
	
tx.Commit()	
```

## Debug
Debug Tx with traceID

```go
[Success][30ms][Tx][#1]
[Success][30ms][Tx][#1]
[Success][30ms][Tx][#1]
```