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

func (p *Result) scan() (err error) {
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
		if len(p.selectedList) > 0 && !p.selectedList[0].isArray {
			break
		}
	}
	return
}

func (p *Result) reflectValue(index int, value interface{}) error {

	switch p.values[index].Elem().Kind() {
	case reflect.Int8:
		v, err := cast.ToInt8E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetInt(int64(v))
		return nil
	case reflect.Int16:
		v, err := cast.ToInt16E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetInt(int64(v))
		return nil
	case reflect.Int32:
		v, err := cast.ToInt32E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetInt(int64(v))
		return nil
	case reflect.Int64:
		v, err := cast.ToInt64E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetInt(v)
		return nil
	case reflect.Uint8:
		v, err := cast.ToUint8E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetUint(uint64(v))
		return nil
	case reflect.Uint16:
		v, err := cast.ToUint16E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetUint(uint64(v))
		return nil
	case reflect.Uint32:
		v, err := cast.ToUint32E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetUint(uint64(v))
		return nil
	case reflect.Uint64:
		v, err := cast.ToUint64E(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetUint(v)
		return nil
	case reflect.String:
		v, err := cast.ToStringE(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetString(v)
		return nil
	case reflect.Float32:
		v, err := cast.ToStringE(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetString(v)
		return nil
	case reflect.Float64:
		v, err := cast.ToStringE(value)
		if err != nil {
			return err
		}
		p.values[index].Elem().SetString(v)
		return nil
	case reflect.Slice, reflect.Array:
		err := p.append(index, p.values[index].Elem(), value)
		if err != nil {
			return err
		}
		return nil
	case reflect.Map:

	case reflect.Struct:
		if p.values[index].Elem().Type().Name() == "decimal.Decimal" {
			v, err := cast.ToDecimalE(value)
			if err != nil {
				return err
			}
			p.values[index].Elem().Set(reflect.ValueOf(v))
			return nil
		}
	}

	return fmt.Errorf(
		"unsupport convert field '%s' type '%s' to '%s'",
		p.selectedList[index].name, reflect.TypeOf(value).Kind(), p.values[index].Elem().Kind(),
	)
}

func (p *Result) append(index int, elem reflect.Value, value interface{}) error {
	var rv reflect.Value
	switch elem.Type().Elem().Kind() {
	case reflect.Int8:
		v, err := cast.ToInt8E(value)
		if err != nil {
			return err
		}
		rv = reflect.ValueOf(v)
	case reflect.Int16:
		v, err := cast.ToInt16E(value)
		if err != nil {
			return err
		}
		rv = reflect.ValueOf(v)
	case reflect.Int32:
		v, err := cast.ToInt32E(value)
		if err != nil {
			return err
		}
		rv = reflect.ValueOf(v)
	case reflect.Int64:
		v, err := cast.ToInt64E(value)
		if err != nil {
			return err
		}
		rv = reflect.ValueOf(v)
	case reflect.String:
		v, err := cast.ToStringE(value)
		if err != nil {
			return err
		}
		rv = reflect.ValueOf(v)
	default:
		return fmt.Errorf("unsupport convert slice field '%s' type '%s'",
			p.selectedList[index].name, p.values[index].Elem().Kind())
	}

	p.values[index].Elem().Set(reflect.Append(p.values[index].Elem(), rv))

	return nil
}
