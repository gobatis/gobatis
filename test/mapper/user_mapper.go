package mapper

import (
	"database/sql"
	"github.com/gobatis/gobatis/engine"
	"github.com/gobatis/gobatis/test/entity"
)

type UserMapper struct {
	Engine     *engine.Engine
	GetUser    func(tx *sql.Tx, session *engine.Session, id int64) (user *entity.User, err error)
	CreateUser func(tx *sql.Tx, name int64) (res sql.Result, err error)
	InsertUser func(tx *sql.Tx, name int64) (rowsAffected int64, err error)
	UpdateUser func(tx *sql.Tx, name int64) (rowsAffected int64, err error)
}

//func (p *UserMapper) GetUser(session *engine.Session, id int64) {
//	p.Engine.Call("UserMapper", session, id, )
//}
