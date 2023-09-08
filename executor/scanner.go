package executor

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Scanner struct {
	rows         *sql.Rows
	RowsAffected int64
	LastInsertId int64
}

func (s *Scanner) Scan(ptr any) (err error) {
	if ptr == nil {
		err = fmt.Errorf("ptr is nil")
		return
	}
	if s.rows == nil {
		err = fmt.Errorf("rows is nil")
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
	first := false
	s.RowsAffected = 0
	for s.rows.Next() {
		s.RowsAffected++
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
	if s.RowsAffected == 0 {
		err = sql.ErrNoRows
		return
	}
	return
}
