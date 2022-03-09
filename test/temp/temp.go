package temp

type Statement struct {
	Method interface{}
	SQL    string
}

type mapper struct {
	CreateUser func(user string) (err error)
}



