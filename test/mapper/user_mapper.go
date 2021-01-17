package mapper

import (
	"github.com/koyeo/mybatis.go/test/entity"
)

type UserMapper struct {
	GetUser func(id int64) (user *entity.User, err error)
}
