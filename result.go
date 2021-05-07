package gobatis

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"reflect"
)

const (
	result_rows = iota + 1
	result_result
)

func newResult(typ int) *Result {
	return &Result{typ: typ}
}

type Result struct {
	typ          int
	rows         *sql.Rows
	result       sql.Result
	selectedMap  map[string]int
	selectedList []param
	values       []reflect.Value
}

func (p *Result) Rows() *sql.Rows {
	return p.rows
}

func (p *Result) Result() sql.Result {
	return p.result
}

func (p *Result) Bind(pointer interface{}, column ...string) error {
	switch p.typ {
	case result_rows:
		return nil
	default:
		return fmt.Errorf("no rows")
	}
}

func (p *Result) setSelected(fields []param) {
	for i, v := range fields {
		if p.selectedMap == nil {
			p.selectedMap = map[string]int{}
		}
		p.selectedMap[v.name] = i
		p.selectedList = append(p.selectedList, v)
	}
}

func (p *Result) isSelected(field string) bool {
	if p.selectedMap == nil {
		return false
	}
	_, ok := p.selectedMap[field]
	return ok
}

func (p *Result) setValues(values []reflect.Value) {
	p.values = values
}

func (p *Result) scanAll() (err error) {
	columns, err := p.rows.Columns()
	if err != nil {
		return
	}
	l := len(columns)
	for p.rows.Next() {
		values := make([]interface{}, l)
		pointers := make([]interface{}, l)
		for i, _ := range columns {
			pointers[i] = &values[i]
		}
		err = p.rows.Scan(pointers...)
		if err != nil {
			_ = p.rows.Close()
			return
		}
		for i, v := range columns {
			if p.isSelected(v) {
				err = p.reflectValue(p.selectedMap[v], values[i])
			}
		}

	}
	return
}

func (p *Result) reflectValue(index int, value interface{}) error {

	switch p.selectedList[index].kind {
	case reflect.Int8:
		v, err := cast.ToInt8E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetInt(int64(v))
	case reflect.Int16:
		v, err := cast.ToInt16E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetInt(int64(v))
	case reflect.Int32:
		v, err := cast.ToInt32E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetInt(int64(v))
	case reflect.Int64:
		v, err := cast.ToInt64E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetInt(v)
	case reflect.String:
		v, err := cast.ToStringE(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetString(v)
	}

	return nil
}
