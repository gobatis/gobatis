package stmt

import "github.com/gobatis/gobatis"

type InsertMapper struct {
	InsertS001Stmt func() (stmt *gobatis.Stmt, err error)
	InsertS002Stmt func() (stmt *gobatis.Stmt, err error)
	InsertS003Stmt func() (stmt *gobatis.Stmt, err error)
	InsertS004Stmt func() (stmt *gobatis.Stmt, err error)
	InsertS005Stmt func(items []Item) (stmt *gobatis.Stmt, err error)
}
