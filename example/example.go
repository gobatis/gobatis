package main

import (
	"fmt"
	"time"

	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/driver/postgres"
	"github.com/gozelle/spew"
	"github.com/shopspring/decimal"
)

type User struct {
	Id     *int64
	Name   string
	Age    int64
	Tstz   time.Time
	Ts     time.Time
	Wealth decimal.Decimal
}

func main() {
	var err error
	defer func() {
		if err != nil {
			panic(fmt.Errorf("\nexmaple error: %s", err))
		}
	}()
	//db, err := batis.Open(postgres.Open("postgresql://root:123456@192.168.1.189:5432/example?connect_timeout=10&sslmode=disable"))
	db, err := batis.Open(postgres.Open("postgresql://root:123456@127.0.0.1:5432/example?connect_timeout=10&sslmode=disable"))
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	//user := &User{
	//	Id:     nil,
	//	Name:   "tom",
	//	Age:    18,
	//	Tstz:   time.Now(),
	//	Ts:     time.Now(),
	//	Wealth: decimal.NewFromFloat(3.14),
	//}
	////id := new(int64)
	//err = db.Debug().Insert("users", user, batis.Returning("id")).Scan(&user.Id).Error
	//if err != nil {
	//	return
	//}
	//spew.Json(user)

	db = db.WithTraceId("123")

	var count int64
	err = db.Debug().Query("select count(1) from users").Scan(&count).Error
	if err != nil {
		return
	}

	tx := db.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Commit()
		}
	}()

	var user *User
	err = tx.Query(`select * from users where id = #{id}`, batis.Param("id", 110)).Scan(&user).Error
	if err != nil {
		return
	}
	
	err = tx.Commit().Error
	if err != nil {
		return
	}

	spew.Json(user)
	//
	//err = db.Update("users", map[string]any{
	//	"age": 99,
	//}, batis.Where("id = #{id}", batis.Param("id", 19))).Error
	//if err != nil {
	//	return
	//}
	//
	//err = db.Delete("users", batis.Where("id = #{id}", batis.Param("id", 20))).Error
	//if err != nil {
	//	return
	//}
	//
	//var total int64
	//err = db.Debug().Query(`select count(1) from users`).Scan(&total).Error
	//if err != nil {
	//	return
	//}
	//spew.Json(total)
	//
	//var items []*User
	//err = db.ParallelQuery(batis.Paging{
	//	Select: "*",
	//	Count:  "*",
	//	From:   "users where id > 10",
	//	Page:   0,
	//	Limit:  10,
	//	Params: nil,
	//	Scan:   []any{&items, &total},
	//}).Error
	//if err != nil {
	//	return
	//}
	//spew.Json(items, total)
	//
	//db.Query(`
	//		select * from users
	//		<where>
	//		     <choose>
	//		     	<when test="groupId == null">
	//		     			group_id is null
	//		     	</when>
	//		     	<otherwise>
	//		     			group_id = #{groupId}
	//		      	</otherwise>
	//		     </choose>
	//	     </where>
	//	     order by id desc`,
	//	batis.Param("group_id", 1),
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
