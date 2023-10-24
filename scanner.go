package batis

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gobatis/gobatis/reflects"
)

type Scanner interface {
	Scan(ptr any, ignore ...string) error
	RowsAffected() int64
	LastInsertId() int64
}

type PagingScanner interface {
	Scan(listPtr, countPtr any, ignore ...string) error
}

type AssociateScanner interface {
	Scan(ptr any, bindingPath, mappingPath string, ignore ...string) error
}

var _ Scanner = (*scanner)(nil)
var _ Scanner = (*insertBatchScanner)(nil)
var _ AssociateScanner = (*associateScanner)(nil)

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

func (s *scanner) Scan(ptr any, ignore ...string) (err error) {
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

func (i *insertBatchScanner) Scan(ptr any, ignore ...string) error {
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

type associateScanner struct {
	rows         *sql.Rows
	rowsAffected int64
	lastInsertId int64
	bindingPaths []associateBindingPath
	mappingPath  string
}

var bindingPathReg = regexp.MustCompile(`^(([a-zA-Z]\w*)+\s*=>\s*(\$(\.[a-zA-Z]\w*)+))(\s*,\s*([a-zA-Z]\w*)+\s*=>\s*(\$(\.[a-zA-Z]\w*)+))*$`)

func (a *associateScanner) Scan(ptr any, bindingPath, mappingPath string, ignore ...string) (err error) {

	err = a.parseBindingPath(bindingPath)
	if err != nil {
		return
	}
	err = a.parseMappingPath(mappingPath)
	if err != nil {
		return
	}

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

func (a *associateScanner) parseBindingPath(bindingPath string) (err error) {
	r := bindingPathReg.FindStringSubmatch(bindingPath)
	if len(r) == 0 {
		err = fmt.Errorf("invaild binding path foramt: %s, expect format like: a => $.A, b => $.B", bindingPath)
		return
	}

	for i := 1; i < len(r); i += 4 {
		if r[i+1] == "" {
			continue
		}
		a.bindingPaths = append(a.bindingPaths, associateBindingPath{
			column: r[i+1],
			path:   strings.TrimPrefix(r[i+2], "$."),
		})
	}
	return
}

func (a *associateScanner) parseMappingPath(mappingPath string) (err error) {
	if !mappingPathReg.MatchString(mappingPath) {
		err = fmt.Errorf("invalid mapping path format: %s", mappingPath)
		return
	}
	a.mappingPath = strings.TrimPrefix(mappingPath, "$.")
	return
}

func (a *associateScanner) reflectRow(columns []string, row []interface{}, pv reflect.Value, first bool) (end bool, err error) {
	m := map[string]any{}
	for i, v := range columns {
		m[v] = row[i]
	}
	err = a.matchBindingValue(m, columns, row, pv, first)
	if err != nil {
		return
	}
	return
}

func (a *associateScanner) matchBindingValue(m map[string]any, columns []string, row []interface{}, pv reflect.Value, first bool) (err error) {

	for _, v := range a.bindingPaths {
		if _, ok := m[v.column]; !ok {
			err = fmt.Errorf("binding column: %s not found in query result columns", v.column)
			return
		}
	}

	pv = reflects.ValueElem(pv)
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

func (a *associateScanner) equal(rv reflect.Value, v any) bool {
	return fmt.Sprintf("%v", rv.Interface()) == fmt.Sprintf("%v", v)
}

func (a *associateScanner) fetchPathValue(pv reflect.Value, path string) (rv reflect.Value) {
	path = strings.TrimPrefix(path, "$.")
	//fmt.Println("path:", path)
	//spew.Json(pv.Type().String(), pv.Interface())
	//spew.Json(pv.FieldByName(path).Interface())
	return pv.FieldByName(path)
}

func (a *associateScanner) RowsAffected() int64 {
	return a.rowsAffected
}

func (a *associateScanner) LastInsertId() int64 {
	return a.lastInsertId
}
