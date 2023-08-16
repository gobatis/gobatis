package main

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/builder/paging"
	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/gobatis/gobatis/example/entity"
	"github.com/gobatis/gobatis/reflects"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	db, err := batis.Open(postgres.Open("postgresql://postgres:postgres@127.0.0.1:5432/gobatis?connect_timeout=10&sslmode=disable"))
	if err != nil {
		return
	}
	
	user := entity.User{}
	err = db.Insert("users", user, batis.OnConflict("", "do update set a = columnd.a")).Error()
	if err != nil {
		return
	}
	
	db.Query(
		`select * from users where age = #{age}`,
		batis.Param("age", 1),
	)
	
	db.Update("users", map[string]any{}, batis.Where(""))
	
	var users []*entity.User
	var total int64
	err = db.Build(paging.Select("id,username").
		Count("*").
		From("public.users").
		Where("age > 18"),
	).Scan(&users, &total)
	if err != nil {
		return
	}
	
	db.Insert("users", users, batis.OnConflict("a,b", "do noting"), batis.Returning("*"))
	db.InsertBatch("users", 10, reflects.Except(user, "id"))
	db.Delete("users", batis.Where("id = ?"))
	db.Update("users", map[string]any{}, batis.Where("id = ?"))
	db.Query(``, batis.Param("", ""))
	db.Exec(``, batis.Param("", ""))
	db.Build(paging.Select("").
		Count("").
		From("").
		Where(""))
	
	ch := db.Fetch("select * from users")
	for a := range ch {
		_ = a.Error
	}
	
	db.Exec(``, batis.Param("user", ""))
}
