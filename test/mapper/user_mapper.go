package mapper

import (
	"github.com/koyeo/gobatis/engine"
)

type UserMapper struct {
	Engine *engine.Engine
	//GetUser func(session *engine.Session, id int64) (user *entity.User, err error)
}

func (p *UserMapper) GetUser(session *engine.Session, id int64) {
	p.Engine.Call("UserMapper", session, id, )
}
