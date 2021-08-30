package gobatis

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
	"time"
)

var (
	reflect_tag = "sql"
)

type queryResult struct {
	rows     *sql.Rows
	first    bool
	reflect  bool
	all      bool
	selected map[string]int
	values   []reflect.Value
}

func (p *queryResult) Tag() string {
	return reflect_tag
}

func (p *queryResult) Rows() *sql.Rows {
	return p.rows
}

func (p *queryResult) setSelected(typeAttribute int, params []*param, values []reflect.Value) error {
	
	p.values = values
	p.reflect = len(params) == 0
	
	if typeAttribute != result_result {
		return nil
	}
	
	var el int
	if p.reflect {
		el = 1
	} else {
		el = len(params)
	}
	
	if el != len(values) {
		return fmt.Errorf("expected to receive %d result filed(s), got %d (except error)", el, len(values))
	}
	
	if p.reflect {
		return nil
	}
	
	p.first = true
	for i, v := range params {
		err := p.addSelected(i, v.name)
		if err != nil {
			return err
		}
		if v.slice {
			p.first = false
		}
	}
	
	return nil
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
	c := 0
	for p.rows.Next() {
		c++
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
	if len(p.values) > 0 && c == 0 {
		err = sql.ErrNoRows
		return
	}
	
	return
}

func (p *queryResult) reflectRow(columns []string, row []interface{}) error {
	
	if p.reflect {
		if p.values[0].Elem().Kind() == reflect.Slice {
			return p.reflectStructs(newRowMap(columns, row))
		} else {
			p.first = true
			return p.reflectStruct(newRowMap(columns, row))
		}
	} else {
		for i, column := range columns {
			if p.isSelected(column) && row[i] != nil {
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
		field = p.trimComma(field)
		v, ok := r[field]
		if ok && v != nil {
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

func (p *queryResult) trimComma(field string) string {
	// TODO 也许可以更加优化
	if strings.Contains(field, ",") {
		return strings.TrimSpace(strings.Split(field, ",")[0])
	}
	return field
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
		field = p.trimComma(field)
		v, ok := r[field]
		if ok && v != nil {
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
	var err error
	dv := reflectValueElem(dest)
	dt := reflectTypeElem(dest.Type())
	if dt.Kind() == reflect.Slice {
		dt = reflectTypeElem(dt.Elem())
	}
	dtv := reflect.New(dt)
	
	if dtv.Type().Implements(scannerType) {
		errs := dtv.MethodByName("Scan").Call([]reflect.Value{reflect.ValueOf(value)})
		if len(errs) > 0 && errs[0].Interface() != nil {
			err = errs[0].Interface().(error)
			return err
		}
		return p.set(dv, dtv.Elem().Interface())
	} else {
		switch dt.Name() {
		case "int8":
			var v int8
			v, err = cast.ToInt8E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "int16":
			var v int16
			v, err = cast.ToInt16E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "int32":
			var v int32
			v, err = cast.ToInt32E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "int64":
			var v int64
			v, err = cast.ToInt64E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "int":
			var v int
			v, err = cast.ToIntE(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "uint8":
			var v uint8
			v, err = cast.ToUint8E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "uint16":
			var v uint16
			v, err = cast.ToUint16E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "uint32":
			var v uint32
			v, err = cast.ToUint32E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "uint64":
			var v uint64
			v, err = cast.ToUint64E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "uint":
			var v uint
			v, err = cast.ToUintE(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "string":
			var v string
			v, err = cast.ToStringE(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "float32":
			var v float32
			v, err = cast.ToFloat32E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "float64":
			var v float64
			v, err = cast.ToFloat64E(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "bool":
			var v bool
			v, err = cast.ToBoolE(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "Time":
			var v time.Time
			v, err = cast.ToTimeE(value)
			if err == nil {
				return p.set(dv, v)
			}
		case "Duration":
			var t time.Time
			t, err = cast.ToTimeE(value)
			if err == nil {
				return p.set(dv, time.Duration(t.Second())*time.Second)
			}
		case "Decimal":
			var v decimal.Decimal
			v, err = cast.ToDecimalE(value)
			if err == nil {
				return p.set(dv, v)
			}
		}
	}
	if err != nil {
		return fmt.Errorf("scan field %s error: %s", column, err.Error())
	}
	
	return fmt.Errorf("can't scan field '%s' type '%s' to '%s'", column, reflect.TypeOf(value), dest.Type())
}

func (p *queryResult) set(dest reflect.Value, r interface{}) error {
	if dest.Kind() == reflect.Slice {
		if dest.Type().Elem().Kind() == reflect.Ptr {
			v := reflect.New(dest.Type().Elem().Elem())
			v.Elem().Set(reflect.ValueOf(r))
			dest.Set(reflect.Append(dest, v))
		} else {
			dest.Set(reflect.Append(dest, reflect.ValueOf(r)))
		}
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

type execResult struct {
	affected int64
	values   []reflect.Value
}

func (p *execResult) scan() error {
	if len(p.values) > 0 {
		switch p.values[0].Elem().Kind() {
		case reflect.Int:
			r, e := cast.ToIntE(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetInt(int64(r))
		case reflect.Int8:
			r, e := cast.ToInt8E(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetInt(int64(r))
		case reflect.Int16:
			r, e := cast.ToInt16E(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetInt(int64(r))
		case reflect.Int32:
			r, e := cast.ToInt32E(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetInt(int64(r))
		case reflect.Int64:
			p.values[0].Elem().SetInt(p.affected)
		case reflect.Uint:
			r, e := cast.ToUintE(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetUint(uint64(r))
		case reflect.Uint8:
			r, e := cast.ToUint8E(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetUint(uint64(r))
		case reflect.Uint16:
			r, e := cast.ToUint16E(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetUint(uint64(r))
		case reflect.Uint32:
			r, e := cast.ToUint32E(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetUint(uint64(r))
		case reflect.Uint64:
			r, e := cast.ToUint64E(p.affected)
			if e != nil {
				return e
			}
			p.values[0].Elem().SetUint(r)
		}
	}
	return nil
}
