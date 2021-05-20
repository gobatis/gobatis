package gobatis

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"github.com/shopspring/decimal"
	"reflect"
	"time"
)

type queryResult struct {
	rows       *sql.Rows
	first      bool
	tag        string
	resultType *resultType
	selected   map[string]int
	values     []reflect.Value
}

func (p *queryResult) Tag() string {
	if p.tag != "" {
		return p.tag
	}
	return "sql"
}

func (p *queryResult) Rows() *sql.Rows {
	return p.rows
}

func (p *queryResult) setSelected(_dest *resultType, params []*param, values []reflect.Value) error {
	
	p.resultType = _dest
	p.values = values
	
	var el int
	if p.resultType != nil {
		return nil
	}
	
	if p.resultType != nil {
		el = 1
	} else {
		el = len(params)
	}
	
	if el != len(values) {
		return fmt.Errorf("expected to receive %d result filed(s), got %d (except error)", el, len(values))
	}
	
	err := p.checkResultType()
	if err != nil {
		return err
	}
	
	if p.resultType != nil {
		return nil
	}
	
	p.first = true
	for i, v := range params {
		err = p.addSelected(i, v.name)
		if err != nil {
			return err
		}
		if v.slice {
			p.first = false
		}
	}
	
	return nil
}

func (p *queryResult) checkResultType() (err error) {
	if p.resultType != nil && p.values[0].Elem().Kind() != reflect.Struct {
		return fmt.Errorf("bind value is not struct or map")
	}
	return
}

func (p *queryResult) addSelected(index int, name string) error {
	if p.selected == nil {
		p.selected = map[string]int{}
	}
	if _, ok := p.selected[name]; ok {
		return fmt.Errorf("duplicated result filed '%s'", name)
	}
	p.selected[name] = index
	return nil
}
func (p *queryResult) isSelected(field string) bool {
	if p.selected == nil {
		return false
	}
	_, ok := p.selected[field]
	return ok
}

func (p *queryResult) scan() (err error) {
	columns, err := p.rows.Columns()
	if err != nil {
		return
	}
	l := len(columns)
	rc := 0
	for p.rows.Next() {
		rc++
		row := make([]interface{}, l)
		pointers := make([]interface{}, l)
		for i, _ := range columns {
			pointers[i] = &row[i]
		}
		err = p.rows.Scan(pointers...)
		if err != nil {
			_ = p.rows.Close()
			return
		}
		err = p.reflectRow(columns, row)
		if err != nil {
			return
		}
		if p.first {
			break
		}
	}
	
	// TODO Debug p.values
	if len(p.values) > 0 && rc == 0 {
		err = sql.ErrNoRows
		return
	}
	
	return
}

func (p *queryResult) reflectRow(columns []string, row []interface{}) error {
	if p.resultType != nil {
		if p.values[0].Elem().Kind() == reflect.Slice {
			return p.reflectStructs(newRowMap(columns, row))
		} else {
			p.first = true
			return p.reflectStruct(newRowMap(columns, row))
		}
	} else {
		for i, column := range columns {
			if p.isSelected(column) {
				err := p.reflectValue(column, p.values[p.selected[column]].Elem(), row[i])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *queryResult) reflectStruct(r rowMap) error {
	dv := p.values[0]
	if dv.Kind() == reflect.Ptr {
		dv = dv.Elem()
	}
	_type := dv.Type()
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i).Tag.Get(p.Tag())
		v, ok := r[field]
		if ok {
			if dv.Field(i).Kind() == reflect.Ptr {
				dv.Field(i).Set(reflect.New(dv.Field(i).Type().Elem()))
			}
			err := p.reflectValue(field, dv.Field(i), v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *queryResult) reflectStructs(r rowMap) error {
	var _type reflect.Type
	if p.values[0].Type().Elem().Elem().Kind() != reflect.Ptr {
		// var test []Test => Test
		_type = p.values[0].Type().Elem().Elem()
	} else {
		// var test []*Test => Test
		_type = p.values[0].Type().Elem().Elem().Elem()
	}
	elem := reflect.New(_type)
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i).Tag.Get(p.Tag())
		v, ok := r[field]
		if ok {
			if elem.Elem().Field(i).Kind() == reflect.Ptr {
				elem.Elem().Field(i).Set(reflect.New(elem.Elem().Field(i).Type().Elem()))
			}
			err := p.reflectValue(field, elem.Elem().Field(i), v)
			if err != nil {
				return err
			}
		}
	}
	if p.values[0].Type().Elem().Elem().Kind() != reflect.Ptr {
		p.values[0].Elem().Set(reflect.Append(p.values[0].Elem(), elem.Elem()))
	} else {
		p.values[0].Elem().Set(reflect.Append(p.values[0].Elem(), elem))
	}
	return nil
}

func (p *queryResult) reflectValue(column string, dest reflect.Value, value interface{}) error {
	
	var (
		isPtr = dest.Kind() == reflect.Ptr
		dv    = dest
	)
	
	if isPtr {
		dv = dest.Elem()
	}
	switch tv := dv.Interface().(type) {
	case sql.Scanner:
		// TODO debug scanner
		err := tv.Scan(value)
		if err != nil {
			return err
		}
	case int8:
		v, err := cast.ToInt8E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case int16:
		v, err := cast.ToInt16E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case int32:
		v, err := cast.ToInt32E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case int64:
		v, err := cast.ToInt64E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case int:
		v, err := cast.ToIntE(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case uint8:
		v, err := cast.ToUint8E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case uint16:
		v, err := cast.ToUint16E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case uint32:
		v, err := cast.ToUint32E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case uint64:
		v, err := cast.ToUint64E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case uint:
		v, err := cast.ToUintE(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case string:
		v, err := cast.ToStringE(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case float32:
		v, err := cast.ToFloat32E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case float64:
		v, err := cast.ToFloat64E(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case time.Time:
		v, err := cast.ToTimeE(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	case decimal.Decimal:
		v, err := cast.ToDecimalE(value)
		if err != nil {
			return err
		}
		return p.set(dv, v)
	}
	
	return fmt.Errorf("can't scan field '%s' type '%s' to '%s'", column, reflect.TypeOf(value), dest.Type())
}

func (p *queryResult) set(dest reflect.Value, r interface{}) error {
	if dest.Kind() == reflect.Slice {
		dest.Set(reflect.Append(dest, reflect.ValueOf(r)))
	} else {
		dest.Set(reflect.ValueOf(r))
	}
	return nil
}

func newRowMap(columns []string, values []interface{}) rowMap {
	m := rowMap{}
	for i, v := range columns {
		m[v] = values[i]
	}
	return m
}

type rowMap map[string]interface{}

func newExecResult(res sql.Result, values []reflect.Value) *execResult {
	return &execResult{res: res, values: values}
}

type execResult struct {
	res    sql.Result
	values []reflect.Value
}

func (p *execResult) scan() error {
	ra, err := p.res.RowsAffected()
	if err != nil {
		return err
	}
	if len(p.values) > 0 {
		p.values[0].Elem().SetInt(ra)
	}
	return nil
}
