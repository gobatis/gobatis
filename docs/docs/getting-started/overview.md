---
sidebar_position: 2
---


# Overview

## Install 

```
go get github.com/gobatis/gobaits
```

## QuickStart

```go
package main

import (
	batis "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/builder/paging"
	"github.com/gobatis/gobatis/dirver/postgres"
)

type User struct {
	Id   *int64
	Name string
}

func main() {
	db, err := baits.Open(postgres.Open("test.db"))
	if err != nil {
		panic("failed to connect database")
	}
	
	db.Migrate(&User{})
	
	// Create
	user := &User{Name: "tom"}
	db.Insert("users", user, batis.OnConflict("id", "do noting"), batis.Returning(`id`)).Scan(&user.Id)
	
	// Read
	db.Query(`select * from users where name = #{name}`, baits.Param("name", "tom")).Scan(user)
	
	// Update - update user's name to jack
	db.Update("users", map[string]any{
		"name": "jack",
	}, batis.Where(`name = #{name}`, batis.Param("name", "tom")))
	
	// Delete - delete product
	db.Delete("users", batis.Where(`name = #{name}`, batis.Param("name", "jack")))
	
	// Pageing
	var users []*User
	var total
	db.Run(batis.Paging{
		Select: "*",
		Count:  "*",
		Common: `users where name age > #{age}`,
		Page:   0,
		Limit:  0,
		Params: []executor.Param{
			{Name: "age", Value: 18},
		},
	}).Scan(&users,total)
}

```