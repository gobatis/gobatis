package batis

import "github.com/gobatis/gobatis/dialector"

func Select(data any, columns string) Rows {
	return &selectColumns{
		columns: columns,
		data:    data,
	}
}

func Except(data any, columns string) Rows {
	return &exceptColumns{
		data:    data,
		columns: columns,
	}
}

type Rows interface {
	Reflect(namer dialector.Namer, tag string) (rows []Row, err error)
}

var _ Rows = (*selectColumns)(nil)
var _ Rows = (*exceptColumns)(nil)

type selectColumns struct {
	data    any
	columns string
}

func (s selectColumns) Reflect(namer dialector.Namer, tag string) (rows []Row, err error) {
	rows, err = reflectRows(s.data, namer, tag)
	if err != nil {
		return
	}

	return
}

type exceptColumns struct {
	data    any
	columns string
}

func (e exceptColumns) Reflect(namer dialector.Namer, tag string) (rows []Row, err error) {
	//TODO implement me
	panic("implement me")
}
