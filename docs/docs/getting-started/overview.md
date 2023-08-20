---
sidebar_position: 2
title: Overview
---

<img src="/img/logo.png" style={{width: 300, display: 'block', margin:'0 auto'}} />

Most database operation scenarios involve simple CRUD (Create, Read, Update, Delete) operations. For these cases,
we need straightforward methods. When conducting complex queries, excessive parameter bindings become a headache; at
these times, a convenient template syntax is required. In some specific scenarios, such as paginated queries, we need
some quick methods.

Taking into account the above requirements, Gobatis was designed. It adheres to the traditional usage habits of Go ORMs,
and also draws from MyBatis's Dynamic SQL syntax. Additionally, it offers numerous other features, making system
development simpler and more efficient.

## Feature

* Simple, An engineering-oriented ORM
* Intuitive and convenient API design
* Targeted at users who prefer using native SQL
* Transaction tracing.
* Mybatis parameter syntax and Dynamic SQL syntax
* Hooks (Before/After, Insert/Update/Delete/Query/Exec)
* More rigorous query result matching mechanism
* Context, Prepared Statement Mode, Debug Mode, DryRun Mode, Loose Mode
* Logger
* Every feature comes with tests
* Developer Friendly

## Install

```
go get github.com/gobatis/gobaits
```

## Quick Start

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
	}).Scan(&users, total)
}

```