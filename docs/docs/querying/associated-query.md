---
sidebar_position: 10
---

# Associated Query



```go
type User struct {
		Id     int64
		Name   string
		Posts  []string
		Orders []string
	}

	type Post struct {
		Id   int64
		Tags []string
	}

	var users []User

    db.Must().Query(`select * from users`).LooseScan(&users, "$..Posts");

	db.Must().Exec(`select`)

	userIds := mapping.Merge(users, func())

	db.Query().LooseScan(batis.LooseDest(&users, "$..Posts","$..Orders")).Error

    db.Query(`select * from posts where user_id in #{userIds}`,
		batis.Param("userIds", userIds)).Link(&users, "user_id => $..Id", "$..Posts").Error

	postIds := mapping.Map(users, func())

	db.Query(`select * from tags where post_id in #{postIds}`,
		batis.Param("postIds", postIds)).
		Link(&users, "user_id => $..Posts[*].Id", "$..Post[*].Tags").Error

    db.Query(`select * from orders where user_id in #{userIds}`,
		batis.Param("userIds", userIds)).
		Link(&users, "user_id => $..Id", "$..Orders").Error
```