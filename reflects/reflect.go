package reflects

import (
	"fmt"
	"github.com/gobatis/gobatis/dialector"
	"reflect"
)

type Row []*Column

type Column struct {
	column string
	value  reflect.Value
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
	
	if rt.Kind() != reflect.Struct {
		err = fmt.Errorf("only accept struct")
		return
	}
	
	if multiple {
		for i := 0; i < rv.Len(); i++ {
			rows = append(rows, reflectStruct(rt, rv.Slice(i, i)))
		}
	} else {
		rows = append(rows, reflectStruct(rt, rv))
	}
	
	return
}

func reflectStruct(rt reflect.Type, rv reflect.Value) Row {
	
	var r Row
	for i := 0; i < rt.NumField(); i++ {
		r = append(r, &Column{
			column: rt.Field(i).Name,
			value:  rv.Field(i),
		})
	}
	
	return r
}

func RowsColumns(rows []Row, namer dialector.Namer) (columns []string) {
	if len(rows) == 0 {
		return
	}
	for _, v := range rows[0] {
		columns = append(columns, namer.ReservedName(v.column))
	}
	return
}

func RowsVars(rows []Row) (vars []string) {
	if len(rows) == 0 {
		return
	}
	for _, v := range rows[0] {
		vars = append(vars, fmt.Sprintf("#{%s}", v.column))
	}
	return
}
