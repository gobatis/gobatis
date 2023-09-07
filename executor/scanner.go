package executor

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Scanner struct {
	rows   *sql.Rows
	result sql.Result
}

func (s Scanner) Scan(ptr any) (err error) {
	if ptr == nil {
		err = fmt.Errorf("ptr is nil")
		return
	}
	pv := reflect.ValueOf(ptr)
	if pv.Kind() != reflect.Pointer || pv.IsNil() {
		return &InvalidUnmarshalError{pv.Type()}
	}
	pv = indirect(pv, false)

	columns, err := s.rows.Columns()
	if err != nil {
		return
	}
	l := len(columns)
	c := 0
	first := false
	for s.rows.Next() {
		c++
		row := make([]interface{}, l)
		pointers := make([]interface{}, l)
		for i, _ := range columns {
			pointers[i] = &row[i]
		}
		err = s.rows.Scan(pointers...)
		if err != nil {
			return
		}
		if !first {
			first = true
		}
		var end bool
		end, err = reflectRow(columns, row, pv, first)
		if err != nil {
			return
		}
		if end {
			break
		}
	}
	if c == 0 {
		err = sql.ErrNoRows
		return
	}
	return
}
