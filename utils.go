package gobatis

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/ttacon/chalk"
	"reflect"
	"strings"
)

const (
	json_tag    = "json"
	default_tag = "sql"
)

var errorType reflect.Type
var scannerType reflect.Type

type Valuer interface {
	Value() (driver.Value, error)
}

func init() {
	errorType = reflect.TypeOf((*error)(nil)).Elem()
	scannerType = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
}

func isCtx(v reflect.Type) bool {
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

func isError(t reflect.Type) bool {
	return t.Implements(reflect.TypeOf((*error)(nil)).Elem())
}

func toReflectValueElem(s interface{}) reflect.Value {
	v := reflect.ValueOf(s)
	for {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		} else {
			return v
		}
	}
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

func innerVar(id, name string) string {
	return fmt.Sprintf("%s_%s", id, name)
}

func innerExpr(name string) string {
	return fmt.Sprintf("#{%s}", name)
}

func innerMethod(id, method string) string {
	return fmt.Sprintf("%s@%s", id, method)
}

func snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
