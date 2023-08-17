package main

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgres"
)

type User struct {
	Id   *int64
	Name string
	Age  int64
}

func main() {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	db, err := batis.Open(postgres.Open("postgresql://root:123456@192.168.1.189:5432/example?connect_timeout=10&sslmode=disable"))
	if err != nil {
		return
	}
	
	err = db.Ping()
	if err != nil {
		return
	}
	
	//user := &User{
	//	Id:   nil,
	//	Name: "tom",
	//	Age:  18,
	//}
	////id := new(int64)
	//err = db.Debug().Insert("users", user, batis.Returning("id")).Scan(&user.Id)
	//if err != nil {
	//	return
	//}
	//spew.Json(user)
	//
	//var users []****User
	//err = db.Query(`select * from users where id = #{id}`, batis.Param("id", 18)).Scan(&users)
	//if err != nil {
	//	return
	//}
	//
	//spew.Json(user)
	
	err = db.Update("users", map[string]any{
		"age": 99,
	}, batis.Where("id = #{id}", batis.Param("id", 18))).Error()
	if err != nil {
		return
	}
	
	//db.Query(
	//	`select * from users where age = #{age}`,
	//	batis.Param("age", 1),
	//)
	
	//db.Update("users", map[string]any{}, batis.Where(""))
	//
	//var users []*entity.User
	//var total int64
	//err = db.Build(paging.Select("id,username").
	//	Count("*").
	//	From("public.users").
	//	Where("age > 18"),
	//).Scan(&users, &total)
	//if err != nil {
	//	return
	//}
	//
	//db.Insert("users", users, batis.OnConflict("a,b", "do noting"), batis.Returning("*"))
	//db.InsertBatch("users", 10, batis.Except(user, "id"))
	//db.Delete("users", batis.Where("id = ?"))
	//db.Update("users", map[string]any{}, batis.Where("id = ?"))
	//db.Query(``, batis.Param("", ""))
	//db.Exec(``, batis.Param("", ""))
	//db.Build(paging.Select("").
	//	Count("").
	//	From("").
	//	Where(""))
	//
	//ch := db.Fetch("select * from users")
	//for a := range ch {
	//	_ = a.Error
}

//db.Exec(``, batis.Param("user", ""))
//}
