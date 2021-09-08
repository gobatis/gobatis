package gobatis

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/ttacon/chalk"
	"reflect"
)

var errorType reflect.Type
var scannerType reflect.Type
var valuerType reflect.Type

type Valuer interface {
	Value() (driver.Value, error)
}

func init() {
	errorType = reflect.TypeOf((*error)(nil)).Elem()
	scannerType = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
	valuerType = reflect.TypeOf((*Valuer)(nil)).Elem()
}

func isContext(v reflect.Type) bool {
	if v.Name() == "Context" && v.PkgPath() == "context" {
		return true
	}
	return false
}

func isTx(v reflect.Type) bool {
	if v.Kind() == reflect.Ptr && v.Elem().Name() == "Tx" &&
		(v.Elem().PkgPath() == "database/sql" ||
			v.Elem().PkgPath() == "github.com/gobatis/gobatis") {
		return true
	}
	return false
}

func isDB(v reflect.Type) bool {
	if v.Kind() == reflect.Ptr && v.Elem().Name() == "DB" && v.Elem().PkgPath() == "github.com/gobatis/gobatis" {
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

func printVars(vars []interface{}) string {
	if len(vars) == 0 {
		return ""
	}
	r := "\n"
	for i, v := range vars {
		_t := ""
		if v != nil {
			_t = reflect.TypeOf(v).String()
		}
		r += fmt.Sprintf("   $%d %s (%s) %+v\n",
			i+1, chalk.Green.Color("=>"), chalk.Yellow.Color(_t), v)
	}
	return r
}
