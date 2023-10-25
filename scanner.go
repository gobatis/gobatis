package batis

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/gobatis/gobatis/logger"
	"github.com/gobatis/gobatis/parser"
	"github.com/gobatis/gobatis/reflects"
)

type scanner interface {
	setRows(rows *sql.Rows)
	setDest(dest any, ignore ...string)
	scan() error
}

type Scanner interface {
	scanner
	Scan(ptr any, ignore ...string) error
}

type PagingScanner interface {
	scanner
	Scan(countPtr, listPtr any, ignore ...string) error
}

type AssociateScanner interface {
	scanner
	Scan(ptr any, bindingPath, mappingPath string, ignore ...string) error
}

var _ Scanner = (*defaultScanner)(nil)
var _ Scanner = (*insertBatchScanner)(nil)
var _ PagingScanner = (*pagingScanner)(nil)
var _ AssociateScanner = (*associateScanner)(nil)

type defaultScanner struct {
	rows         *sql.Rows
	reflectRow   func(columns []string, row []interface{}, pv reflect.Value, first bool) (bool, error)
	dest         any
	ignore       []string
	rowsAffected int64
}

func (d *defaultScanner) setDest(dest any, ignore ...string) {
	d.dest = dest
	d.ignore = ignore
}

// 与结构体里的 dest, ignore... 作用一样
// 在不走 Scan 接口时，通过注入 scanner 到回调函数中，从回调函数中获取 dest, ignore...
func (d *defaultScanner) Scan(dest any, ignore ...string) error {
	d.dest = dest
	d.ignore = ignore
	return d.scan()
}

func (d *defaultScanner) setRows(rows *sql.Rows) {
	d.rows = rows
}

func (d *defaultScanner) scan() (err error) {
	if d.rows == nil {
		err = fmt.Errorf("scan rows error: rows is nil")
		return
	}
	var pv reflect.Value
	var or reflect.Value
	var ov interface{}
	defer func() {
		if d.dest != nil && d.rowsAffected == 0 {
			if or.Kind() == reflect.Pointer {
				or.Set(reflect.ValueOf(ov))
			}
		}
	}()

	if d.dest != nil {
		pv = reflect.ValueOf(d.dest)
		if pv.Kind() != reflect.Pointer || pv.IsNil() {
			return &InvalidUnmarshalError{pv.Type()}
		}
		or = pv.Elem()
		ov = or.Interface()
		pv = indirect(pv, false)
	}
	columns, err := d.rows.Columns()
	if err != nil {
		return
	}
	l := len(columns)
	first := false
	d.rowsAffected = 0
	for d.rows.Next() {
		d.rowsAffected++
		if d.dest != nil {
			row := make([]interface{}, l)
			pointers := make([]interface{}, l)
			for i, _ := range columns {
				pointers[i] = &row[i]
			}
			err = d.rows.Scan(pointers...)
			if err != nil {
				return
			}
			if !first {
				first = true
			}
			var end bool
			if d.reflectRow != nil {
				end, err = d.reflectRow(columns, row, pv, first)
			} else {
				end, err = d.defaultReflectRow(columns, row, pv, first)
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

func (d *defaultScanner) defaultReflectRow(columns []string, row []interface{}, pv reflect.Value, first bool) (bool, error) {
	return reflectRow(columns, row, pv, first)
}

type insertBatchScanner struct {
	rows         *sql.Rows
	rowsAffected int64
	lastInsertId int64
}

func (i *insertBatchScanner) setDest(dest any, ignore ...string) {
	//TODO implement me
	panic("implement me")
}

func (i *insertBatchScanner) scan() error {
	return nil
}

func (i *insertBatchScanner) setRows(rows *sql.Rows) {
	i.rows = rows
}

func (i *insertBatchScanner) Scan(ptr any, ignore ...string) error {
	rv := reflect.ValueOf(ptr)
	if rv.Elem().Type().Kind() != reflect.Slice {
		return fmt.Errorf("expect slice, got %s", rv.Elem().Type())
	}
	s := &defaultScanner{rows: i.rows}
	err := s.Scan(ptr)
	if err != nil {
		return err
	}
	i.rowsAffected += s.rowsAffected
	//i.lastInsertId = s.lastInsertId
	return nil
}

type pagingScanner struct {
	query        *raw
	count        *raw
	method       string
	ctx          context.Context
	conn         func() conn
	logger       logger.Logger
	pos          string
	trace        bool
	debug        bool
	rowsAffected int64
}

func (p *pagingScanner) setDest(dest any, ignore ...string) {
	panic("implement me")
}

func (p *pagingScanner) setRows(rows *sql.Rows) {
	panic("implement me")
}

func (p *pagingScanner) scan() error {
	panic("implement me")
}

func (p *pagingScanner) prepareDefaultExecutor() *defaultExecutor {
	return &defaultExecutor{
		method: p.method,
		ctx:    p.ctx,
		conn:   p.conn(),
		logger: p.logger,
		pos:    p.pos,
		trace:  p.trace,
		debug:  p.debug,
	}
}

func (p *pagingScanner) Scan(countPtr, listPtr any, ignore ...string) (err error) {
	fns := []func() error{
		func() error {
			d := p.prepareDefaultExecutor()
			d.raw = p.count
			d.scanner = &defaultScanner{
				dest: countPtr,
			}
			d.scan = func(s scanner) error {
				return s.scan()
			}
			_, e := d.execute()
			return e
		},
		func() error {
			d := p.prepareDefaultExecutor()
			d.raw = p.query
			d.scanner = &defaultScanner{
				dest:   listPtr,
				ignore: ignore,
			}
			d.scan = func(s scanner) error {
				e := s.scan()
				if e != nil {
					return e
				}
				//p.rowsAffected = s.RowsAffected()
				return nil
			}
			_, e := d.execute()
			return e
		},
	}

	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, fn := range fns {
		wg.Add(1)
		go func(fn func() error) {
			defer func() {
				wg.Done()
			}()
			e := fn()
			if e != nil {
				lock.Lock()
				err = parser.AddError(err, e)
				lock.Unlock()
			}
		}(fn)
	}
	wg.Wait()

	return
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

func (a *associateScanner) setDest(dest any, ignore ...string) {
	panic("implement me")
}

func (a *associateScanner) scan() error {
	panic("implement me")
}

func (a *associateScanner) setRows(rows *sql.Rows) {
	a.rows = rows
}

func (a *associateScanner) Scan(ptr any, bindingPath, mappingPath string, ignore ...string) (err error) {

	err = a.parseBindingPath(bindingPath)
	if err != nil {
		return
	}
	err = a.parseMappingPath(mappingPath)
	if err != nil {
		return
	}

	s := &defaultScanner{
		rows:       a.rows,
		reflectRow: a.reflectRow,
	}
	err = s.Scan(ptr)
	if err != nil {
		return
	}
	a.rowsAffected = s.rowsAffected
	//a.lastInsertId = s.lastInsertId
	return
}

var bindingPathReg = regexp.MustCompile(`^(([a-zA-Z]\w*)+\s*=>\s*(\$(\.[a-zA-Z]\w*)+))(\s*,\s*([a-zA-Z]\w*)+\s*=>\s*(\$(\.[a-zA-Z]\w*)+))*$`)

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
