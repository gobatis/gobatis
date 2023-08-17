package executor

import (
	"fmt"
	"reflect"
	"strings"
	
	"github.com/gobatis/gobatis/dialector"
)

type Row []*Column

type Column struct {
	column string
	value  any
}

func ReflectRows(v any, namer dialector.Namer, tag string) (rows []Row, err error) {
	
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	
	rt := rv.Type()
	multiple := false
	
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		multiple = true
		rt = rv.Type().Elem()
	}
	
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	
	if rt.Kind() != reflect.Struct {
		err = fmt.Errorf("only accept struct, got: %s", rt.Kind())
		return
	}
	
	if multiple {
		for i := 0; i < rv.Len(); i++ {
			rows = append(rows, reflectStruct(rt, rv.Index(i), namer, tag))
		}
	} else {
		rows = append(rows, reflectStruct(rt, rv, namer, tag))
	}
	
	return
}

func reflectStruct(rt reflect.Type, rv reflect.Value, namer dialector.Namer, tag string) Row {
	
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	
	var r Row
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		var n string
		if t := f.Tag.Get(tag); t != "" {
			n = ExtractTag(t)
		} else {
			n = namer.ColumnName(f.Name)
		}
		v := rv.Field(i)
		if !IsNil(v) {
			r = append(r, &Column{
				column: n,
				value:  v.Interface(),
			})
		}
	}
	
	return r
}

func IsNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		if v.IsNil() {
			return true
		}
	}
	return false
}

func RowColumns(row Row, namer dialector.Namer) (columns []string) {
	for _, v := range row {
		columns = append(columns, namer.ReservedName(v.column))
	}
	return
}

func RowVars(row Row) (vars []string) {
	for _, v := range row {
		vars = append(vars, fmt.Sprintf("#{%s}", v.column))
	}
	return
}

func RowParams(row Row) (params []NameValue) {
	for _, v := range row {
		params = append(params, NameValue{
			Name:  v.column,
			Value: v.value,
		})
	}
	return
}

func RowsVars(rows []Row) (vars []string) {
	for i, v := range rows {
		var s []string
		for _, vv := range v {
			s = append(s, fmt.Sprintf("#{%s%d}", vv.column, i))
		}
		vars = append(vars, fmt.Sprintf("(%s)", strings.Join(s, ",")))
	}
	return
}

func RowsParams(rows []Row) (params []NameValue) {
	for i, v := range rows {
		for _, vv := range v {
			params = append(params, NameValue{
				Name:  fmt.Sprintf("%s%d", vv.column, i),
				Value: vv.value,
			})
		}
		
	}
	return
}
