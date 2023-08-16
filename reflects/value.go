package reflects

import "github.com/gobatis/gobatis/dialector"

type Rows interface {
	Reflect(namer dialector.Namer, tag string) (rows []Row, err error)
}

var _ Rows = (*SelectColumns)(nil)
var _ Rows = (*ExceptColumns)(nil)

func NewSelectColumns(data any, columns string) *SelectColumns {
	return &SelectColumns{data: data, columns: columns}
}

type SelectColumns struct {
	data    any
	columns string
}

func (s SelectColumns) Reflect(namer dialector.Namer, tag string) (rows []Row, err error) {
	rows, err = ReflectRows(s.data)
	if err != nil {
		return
	}
	
	return
}

func NewExceptColumns(data any, columns string) *ExceptColumns {
	return &ExceptColumns{data: data, columns: columns}
}

type ExceptColumns struct {
	data    any
	columns string
}

func (e ExceptColumns) Reflect(namer dialector.Namer, tag string) (rows []Row, err error) {
	//TODO implement me
	panic("implement me")
}
