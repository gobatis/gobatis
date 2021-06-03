package mapper

import (
	"github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/example/entity"
	"time"
)

var (
	User      = &userMapper{}
	Migration = &migrationMapper{}
)

type migrationMapper struct {
	CreateTable func(db *gobatis.DB) error
}

type userMapper struct {
	AddUser         func(user *entity.User) (rows int, err error)
	AddUserReturnId func(user *entity.User) (id int, createdAt time.Time, err error)
	UpdateUser      func(id int, vip bool) (rows int, err error)
	GetUserById     func(id int64) (name string, age int, err error)
	GetUserByName   func(name string) (user *entity.User, err error)
	GetUserByFrom   func(places []string) ([]*entity.User, error)
	QueryUsers      func(m map[string]interface{}) ([]*entity.User, error)
	DeleteUsers     func(id int) (rows int, err error)
}
