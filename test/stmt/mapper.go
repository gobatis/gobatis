package stmt

import "github.com/gobatis/gobatis"

func NewMapper() *Mapper {
	return &Mapper{
		&insertMapper{},
		&updateMapper{},
	}
}

type Mapper struct {
	*insertMapper
	*updateMapper
}

type insertMapper struct {
	InsertS001Stmt func() (stmt *gobatis.Stmt, err error)
	InsertS002Stmt func() (stmt *gobatis.Stmt, err error)
	InsertS003Stmt func() (stmt *gobatis.Stmt, err error)
	InsertS004Stmt func() (stmt *gobatis.Stmt, err error)
	InsertS005Stmt func(items []Item) (stmt *gobatis.Stmt, err error)
}

type updateMapper struct {
	UpdateS001Stmt func() (stmt *gobatis.Stmt, err error)
	UpdateS002Stmt func() (stmt *gobatis.Stmt, err error)
	UpdateS003Stmt func() (stmt *gobatis.Stmt, err error)
	UpdateS004Stmt func() (stmt *gobatis.Stmt, err error)
	UpdateS005Stmt func(items []Item) (stmt *gobatis.Stmt, err error)
}
