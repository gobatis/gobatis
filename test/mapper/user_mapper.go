package mapper

import (
	"github.com/gobatis/gobatis/engine"
	"github.com/gobatis/gobatis/test/entity"
)

type UserMapper struct {
	Engine  *engine.Engine
	GetUser func(session *engine.Session, id int64) (user *entity.User, err error)
}

//func (p *UserMapper) GetUser(session *engine.Session, id int64) {
//	p.Engine.Call("UserMapper", session, id, )
//}
