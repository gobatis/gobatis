---
sidebar_position: 0
---

# DB Chain

## Clone

new clone db with every chain call

```go 
fmt.Printf("%p", db)          // #1
fmt.Printf("%p", db.Debug())  // #2
fmt.Printf("%p", db.Must())   // #3
```

## Must Mode

will check the RowsAffected > 0  or Scan rows > 0

```go
db.Must()
```

## Debug Mode

```go
db.Debug()
```