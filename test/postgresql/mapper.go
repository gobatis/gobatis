package postgresql

import (
	_decimal "github.com/gozelle/decimal"
	_time "time"
)

type decimal = _decimal.Decimal
type time = _time.Time
type Mapper struct {
	*MakeMapper
	Migrate    func() error
	ResetTable func() error
	//InsertNullType func(a interface{}) (row error)
	//SelectNullInt  func(id int) error
	//InsertReturnId func(sid string) (id int, err error)
	//MustGetSidById func(id int) (sid string, err error)
}
