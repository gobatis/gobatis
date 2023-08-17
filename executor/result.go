package executor

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
	
	"github.com/gobatis/gobatis/cast"
	"github.com/shopspring/decimal"
)

var (
	reflectTag = "db"
)

type queryResult struct {
	rows  *sql.Rows
	loose bool
}

func (p *queryResult) Tag() string {
	return reflectTag
}

func (p *queryResult) scan(ptr any) (err error) {
	
	pv := reflect.ValueOf(ptr)
	if pv.Kind() != reflect.Pointer || pv.IsNil() {
		return &InvalidUnmarshalError{pv.Type()}
	}
	pv = indirect(pv, false)
	
	columns, err := p.rows.Columns()
	if err != nil {
		return
	}
	l := len(columns)
	c := 0
	first := false
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
		if !first {
			first = true
		}
		var end bool
		end, err = ReflectRow(columns, row, pv, first)
		if err != nil {
			return
		}
		if end {
			break
		}
	}
	if c == 0 {
		err = sql.ErrNoRows
		return
	}
	
	return
}

//func (p *queryResult) reflectRow(columns []string, row []interface{}, ptr reflect.Value, first bool) (bool, error) {
//	if ptr.Elem().Kind() == reflect.Slice {
//		return false, p.reflectStructs(newRowMap(columns, row), ptr)
//	} else {
//		if ptr.Elem().Kind() == reflect.Struct {
//			return true, p.reflectStruct(newRowMap(columns, row), ptr, first)
//		} else {
//			return true, p.reflectValue("anonymity", ptr.Elem(), row[0])
//		}
//	}
//}

//func (p *queryResult) reflectStruct(r rowMap, ptr reflect.Value, first bool) error {
//	dv := ptr
//	if dv.Kind() == reflect.Ptr {
//		dv = dv.Elem()
//	}
//	_type := dv.Type()
//	
//	var tags map[string]struct{}
//	if first {
//		tags = map[string]struct{}{}
//	}
//	for i := 0; i < _type.NumField(); i++ {
//		n := p.prepareFieldName(_type.Field(i))
//		if first {
//			if _, ok := tags[n]; ok {
//				return fmt.Errorf("field tag: '%s' is duplicated in struct: '%s'", n, _type)
//			}
//			tags[n] = struct{}{}
//		}
//		v, ok := r[n]
//		if !ok {
//			if !p.loose {
//				return fmt.Errorf("no data for struct: '%s' field: '%s'", _type, _type.Field(i).Name)
//			}
//		} else if v != nil {
//			if dv.Field(i).Kind() == reflect.Ptr {
//				dv.Field(i).Set(reflect.New(dv.Field(i).Type().Elem()))
//			}
//			err := p.reflectValue(n, dv.Field(i), v)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}

func (p *queryResult) prepareFieldName(f reflect.StructField) string {
	
	field := f.Tag.Get(p.Tag())
	if field == "" {
		field = toSnakeCase(f.Name)
	}
	return p.trimComma(field)
}

func (p *queryResult) trimComma(field string) string {
	if strings.Contains(field, ",") {
		return strings.TrimSpace(strings.Split(field, ",")[0])
	}
	return field
}

func (p *queryResult) reflectStructs(r rowMap, ptr reflect.Value) error {
	var _type reflect.Type
	if ptr.Type().Elem().Elem().Kind() != reflect.Ptr {
		// var test []Test => Test
		_type = ptr.Type().Elem().Elem()
	} else {
		// var test []*Test => Test
		_type = ptr.Type().Elem().Elem().Elem()
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
	if ptr.Type().Elem().Elem().Kind() != reflect.Ptr {
		ptr.Elem().Set(reflect.Append(ptr.Elem(), elem.Elem()))
	} else {
		ptr.Elem().Set(reflect.Append(ptr.Elem(), elem))
	}
	return nil
}

func (p *queryResult) reflectValue(column string, dest reflect.Value, value interface{}) (err error) {
	
	var dv reflect.Value
	if dest.IsNil() {
		// TODO
	} else {
		dv = reflectValueElem(dest)
	}
	
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
		return fmt.Errorf("scan field '%s' error: %s", column, err.Error())
	}
	
	return fmt.Errorf("can't scan field '%s' type '%s' to '%s'", column, reflect.TypeOf(value), dest.Type())
}

func (p *queryResult) set(dest reflect.Value, r interface{}) error {
	// TODO test
	if dest.Kind() == reflect.Slice {
		if dest.Type().Elem().Kind() == reflect.Ptr {
			v := reflect.New(dest.Type().Elem().Elem())
			v.Elem().Set(reflect.ValueOf(r))
			dest.Set(reflect.Append(dest, v))
		} else {
			dest.Set(reflect.Append(dest, reflect.ValueOf(r)))
		}
	} else if dest.Kind() == reflect.Ptr {
		v := reflect.New(dest.Type().Elem())
		v.Set(reflect.ValueOf(r))
		dest.Set(v)
	} else {
		dest.Set(reflect.ValueOf(r))
	}
	return nil
}

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
