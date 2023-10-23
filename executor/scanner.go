package executor

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/gobatis/gobatis/parser"
)

type Scanner interface {
	Scan(ptr any) error
	RowsAffected() int64
	LastInsertId() int64
}

var _ Scanner = (*scanner)(nil)
var _ Scanner = (*insertBatchScanner)(nil)
var _ Scanner = (*associateQueryScan)(nil)

type scanner struct {
	rows         *sql.Rows
	rowsAffected int64
	lastInsertId int64
	reflectRow   func(columns []string, row []interface{}, pv reflect.Value, first bool) (bool, error)
}

func (s *scanner) RowsAffected() int64 {
	return s.rowsAffected
}

func (s *scanner) LastInsertId() int64 {
	return s.lastInsertId
}

func (s *scanner) Scan(ptr any) (err error) {
	if s.rows == nil {
		err = fmt.Errorf("scan rows error: rows is nil")
		return
	}
	var pv reflect.Value
	var or reflect.Value
	var ov interface{}
	defer func() {
		if ptr != nil && s.rowsAffected == 0 {
			if or.Kind() == reflect.Pointer {
				or.Set(reflect.ValueOf(ov))
			}
		}
	}()

	if ptr != nil {
		pv = reflect.ValueOf(ptr)
		if pv.Kind() != reflect.Pointer || pv.IsNil() {
			return &InvalidUnmarshalError{pv.Type()}
		}
		or = pv.Elem()
		ov = or.Interface()
		pv = indirect(pv, false)
	}
	columns, err := s.rows.Columns()
	if err != nil {
		return
	}
	l := len(columns)
	first := false
	s.rowsAffected = 0
	for s.rows.Next() {
		s.rowsAffected++
		if ptr != nil {
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
			if s.reflectRow != nil {
				end, err = s.reflectRow(columns, row, pv, first)
			} else {
				end, err = s.defaultReflectRow(columns, row, pv, first)
			}
			if err != nil {
				return
			}
			if end {
				break
			}
		}
	}
	return
}

func (s *scanner) defaultReflectRow(columns []string, row []interface{}, pv reflect.Value, first bool) (bool, error) {
	return reflectRow(columns, row, pv, first)
}

type insertBatchScanner struct {
	rows         *sql.Rows
	rowsAffected int64
	lastInsertId int64
}

func (i *insertBatchScanner) Scan(ptr any) error {
	rv := reflect.ValueOf(ptr)
	if rv.Elem().Type().Kind() != reflect.Slice {
		return fmt.Errorf("expect slice, got %s", rv.Elem().Type())
	}
	s := &scanner{rows: i.rows}
	err := s.Scan(ptr)
	if err != nil {
		return err
	}
	i.rowsAffected += s.rowsAffected
	i.lastInsertId = s.lastInsertId
	return nil
}

func (i *insertBatchScanner) RowsAffected() int64 {
	return i.rowsAffected
}

func (i *insertBatchScanner) LastInsertId() int64 {
	return i.lastInsertId
}

type associateBindingPath struct {
	column string
	path   string
}

type associateQueryScan struct {
	rows         *sql.Rows
	rowsAffected int64
	lastInsertId int64
	bindingPaths []*associateBindingPath
	mappingPath  string
}

func (a *associateQueryScan) Scan(ptr any) (err error) {
	s := &scanner{
		rows:       a.rows,
		reflectRow: a.reflectRow,
	}
	err = s.Scan(ptr)
	if err != nil {
		return
	}
	a.rowsAffected = s.rowsAffected
	a.lastInsertId = s.lastInsertId
	return
}

func (a *associateQueryScan) reflectRow(columns []string, row []interface{}, pv reflect.Value, first bool) (end bool, err error) {

	m := map[string]any{}
	for i, v := range columns {
		m[v] = row[i]
	}

	err = a.matchBindingValue(m, columns, row, pv, first)

	return
}

func (a *associateQueryScan) matchBindingValue(m map[string]any, columns []string, row []interface{}, pv reflect.Value, first bool) (err error) {

	for _, v := range a.bindingPaths {
		if _, ok := m[v.column]; !ok {
			err = fmt.Errorf("binding column: %s not found in query result columns", v.column)
			return
		}
	}

	pv = parser.ValueElem(pv)
	if pv.Kind() == reflect.Slice || pv.Kind() == reflect.Array {
		for i := 0; i < pv.Len(); i++ {
			err = a.matchBindingValue(m, columns, row, pv.Index(i), first)
			if err != nil {
				return
			}
		}
		return
	}

	if pv.Kind() != reflect.Struct {
		err = fmt.Errorf("binding value only accept starut as elem, got: %s", pv.Type())
		return
	}

	equal := true
	for _, v := range a.bindingPaths {
		if !a.equal(a.fetchPathValue(pv, v.path), m[v.column]) {
			equal = false
		}
	}
	if !equal {
		return
	}

	//fmt.Println("OK111")

	mv := indirect(a.fetchPathValue(pv, a.mappingPath), false)
	//mv := a.fetchPathValue(pv, a.mappingPath)

	//spew.Json(a.mappingPath, mv.Kind().String())
	_, err = reflectRow(columns, row, mv, first)

	return
}

func (a *associateQueryScan) equal(rv reflect.Value, v any) bool {
	return fmt.Sprintf("%v", rv.Interface()) == fmt.Sprintf("%v", v)
}

func (a *associateQueryScan) fetchPathValue(pv reflect.Value, path string) (rv reflect.Value) {
	path = strings.TrimPrefix(path, "$.")
	//fmt.Println("path:", path)
	//spew.Json(pv.Type().String(), pv.Interface())
	//spew.Json(pv.FieldByName(path).Interface())
	return pv.FieldByName(path)
}

func (a *associateQueryScan) RowsAffected() int64 {
	return a.rowsAffected
}

func (a *associateQueryScan) LastInsertId() int64 {
	return a.lastInsertId
}
