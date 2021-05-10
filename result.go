package gobatis

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"reflect"
)

func newQueryResult(rows *sql.Rows) *queryResult {
	return &queryResult{rows: rows}
}

type queryResult struct {
	rows     *sql.Rows
	first    bool
	tag      string
	dest     *dest
	selected map[string]int
	values   []reflect.Value
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

func (p *queryResult) setSelected(_dest *dest, params []param, values []reflect.Value) error {

	p.dest = _dest
	p.values = values

	var el int
	if p.dest != nil {
		return nil
	}

	if p.dest != nil {
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

	if p.dest != nil {
		return nil
	}

	p.first = true
	for i, v := range params {
		err = p.addSelected(i, v.name)
		if err != nil {
			return err
		}
		if v.isArray {
			p.first = false
		}
	}

	return nil
}

func (p *queryResult) checkResultType() (err error) {
	if p.dest != nil && p.values[0].Elem().Kind() != reflect.Struct {
		return fmt.Errorf("bind value is not struct")
	}
	return
}

func (p *queryResult) reflectElem(v reflect.Value) reflect.Value {
	for {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		} else {
			return v
		}
	}
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
	if rc == 0 {
		err = sql.ErrNoRows
		return
	}
	return
}

func (p *queryResult) reflectRow(columns []string, row []interface{}) error {
	if p.dest != nil {
		//fmt.Println(p.values[0].Elem().Kind())
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
	elem := p.values[0]
	//fmt.Println(elem.Kind(), elem.Elem().Kind())
	//if elem.Kind() == reflect.Ptr && elem.Elem().Kind() == reflect.Invalid {
	//	// var test *Test
	//	p.values[0] = reflect.New(p.values[0].Type().Elem().Elem())
	//	elem = p.values[0].Elem()
	//} else {
	//	// test := new(Test)
	//	//elem = elem.Elem()
	//	fmt.Println(elem.Type())
	//	p.values[0] = reflect.New(elem.Type())
	//}
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	_type := elem.Type()
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i).Tag.Get(p.Tag())
		v, ok := r[field]
		if ok {
			err := p.reflectValue(field, elem.Field(i), v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *queryResult) reflectStructs(r rowMap) error {
	// var test []*Test => Test
	var _type reflect.Type
	if p.values[0].Type().Elem().Elem().Kind() != reflect.Ptr {
		_type = p.values[0].Elem().Type().Elem()
	} else {
		_type = p.values[0].Elem().Type().Elem().Elem()
	}
	elem := reflect.New(_type)
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i).Tag.Get(p.Tag())
		v, ok := r[field]
		if ok {
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

func (p *queryResult) reflectValue(column string, rv reflect.Value, value interface{}) error {

	var kind reflect.Kind
	var r reflect.Value

	switch rv.Kind() {
	case reflect.Slice:
		kind = rv.Type().Elem().Kind()
	default:
		kind = rv.Kind()
	}

	switch kind {
	case reflect.Int8:
		v, err := cast.ToInt8E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Int16:
		v, err := cast.ToInt16E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Int32:
		v, err := cast.ToInt32E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Int64:
		v, err := cast.ToInt64E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Int:
		v, err := cast.ToIntE(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Uint8:
		v, err := cast.ToUint8E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Uint16:
		v, err := cast.ToUint16E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Uint32:
		v, err := cast.ToUint32E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Uint64:
		v, err := cast.ToUint64E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Uint:
		v, err := cast.ToUintE(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.String:
		v, err := cast.ToStringE(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Float32:
		v, err := cast.ToFloat32E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	case reflect.Float64:
		v, err := cast.ToFloat64E(value)
		if err != nil {
			return err
		}
		r = reflect.ValueOf(v)
	default:
		if rv.Type().Name() == "Decimal" && rv.Type().PkgPath() == "github.com/shopspring/decimal" {
			v, err := cast.ToDecimalE(value)
			if err != nil {
				return err
			}
			r = reflect.ValueOf(v)
		} else if rv.Kind() != reflect.Slice {
			return fmt.Errorf(
				"unsupport convert field '%s' type '%s' to '%s'",
				column, reflect.TypeOf(value), rv.Type(),
			)
		}
	}
	if rv.Kind() == reflect.Slice {
		rv.Set(reflect.Append(rv, r))
	} else {
		rv.Set(r)
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
