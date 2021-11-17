package postgresql

import "github.com/gobatis/gobatis"

type Mapper struct {
	*MakeMapper
	Migrate    func(db *gobatis.DB) error
	ResetTable func() error
	//InsertNullType func(a interface{}) (row error)
	//SelectNullInt  func(id int) error
	//InsertReturnId func(sid string) (id int, err error)
	//MustGetSidById func(id int) (sid string, err error)
}
