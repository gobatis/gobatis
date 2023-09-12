package executor

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Scanner interface {
	Scan(ptr any) error
	RowsAffected() int64
	LastInsertId() int64
}

var _ Scanner = (*scanner)(nil)
var _ Scanner = (*insertBatchScanner)(nil)

type scanner struct {
	rows         *sql.Rows
	rowsAffected int64
	lastInsertId int64
}

func (s *scanner) RowsAffected() int64 {
	return s.rowsAffected
}

func (s *scanner) LastInsertId() int64 {
	return s.lastInsertId
}

func (s *scanner) Scan(ptr any) (err error) {
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
	s.rowsAffected = 0
	for s.rows.Next() {
		s.rowsAffected++
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
	//if s.rowsAffected == 0 {
	//	err = sql.ErrNoRows
	//	return
	//}
	return
}

type insertBatchScanner struct {
	rows         *sql.Rows
	result       sql.Result
	rowsAffected int64
	lastInsertId int64
}

func (i insertBatchScanner) Scan(ptr any) error {
	
	rv := reflect.ValueOf(ptr)
	if rv.Elem().Type().Kind() != reflect.Slice {
		return fmt.Errorf("expect slice, got %s", rv.Elem().Type())
	}
	
	s := &scanner{rows: i.rows}
	err := s.Scan(ptr)
	if err != nil {
		return err
	}
	
	return nil
}

func (i insertBatchScanner) RowsAffected() int64 {
	return i.rowsAffected
}

func (i insertBatchScanner) LastInsertId() int64 {
	return i.lastInsertId
}
