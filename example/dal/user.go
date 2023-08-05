package dal

type UserDal struct {
	db *batis.DB
}

func (u UserDal) SaveUser(ctx batis.Context, user string) (err error) {
	
	err = u.db.Insert(ctx.Debug().Must(), "users", user, nil).Error()
	if err != nil {
		return
	}
	
	var age int64
	err = u.db.Execute(ctx.Must().Strict().Analyze(), "select * from users where id = ${id}", batis.Param("id", 123)).Scan(&age)
	if err != nil {
		return
	}
	
	err = u.db.Update(ctx.Must(), "users", map[string]any{
		"name": nil,
	}, nil).Error()
	if err != nil {
		return
	}
	
	return
}
