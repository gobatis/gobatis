package gobatis

import (
	"database/sql"
	"reflect"
)

var errorType reflect.Type
var scannerType reflect.Type

func init() {
	errorType = reflect.TypeOf((*error)(nil)).Elem()
	scannerType = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
}

func isContext(v reflect.Type) bool {
	if v.Name() == "Context" && v.PkgPath() == "context" {
		return true
	}
	return false
}

func isTx(v reflect.Type) bool {
	if v.Kind() == reflect.Ptr && v.Elem().Name() == "Tx" && v.Elem().PkgPath() == "database/sql" {
		return true
	}
	return false
}

func isErrorType(_type reflect.Type) bool {
	return _type.Implements(reflect.TypeOf((*error)(nil)).Elem())
}

func reflectValueElem(vt reflect.Value) reflect.Value {
	for {
		if vt.Kind() != reflect.Ptr {
			break
		}
		vt = vt.Elem()
	}
	return vt
}

func reflectTypeElem(vt reflect.Type) reflect.Type {
	for {
		if vt.Kind() != reflect.Ptr {
			break
		}
		vt = vt.Elem()
	}
	return vt
}
