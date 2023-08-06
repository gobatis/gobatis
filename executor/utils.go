package executor

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"runtime"
	
	"github.com/ttacon/chalk"
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

const lt = "&lt;"

func replaceIsolatedLessThanWithEntity(s string) string {
	// 将字符串转换为 rune 切片，以支持多字节字符
	runes := []rune(s)
	lastLeftBracket := -1
	pos := map[int]struct{}{}
	for i, r := range runes {
		switch r {
		case '<':
			// 如果之前已经有标记的 '<'，替换它
			if lastLeftBracket != -1 {
				pos[lastLeftBracket] = struct{}{}
			}
			lastLeftBracket = i
		case '>':
			// 清除之前标记的 '<'
			lastLeftBracket = -1
		}
	}
	
	// 检查是否在字符串的结尾有一个标记的 '<'
	if lastLeftBracket != -1 {
		pos[lastLeftBracket] = struct{}{}
	}
	
	var r []rune
	for i := range runes {
		if _, ok := pos[i]; ok {
			r = append(r, []rune(lt)...)
		} else {
			r = append(r, runes[i])
		}
	}
	
	return string(r)
}

func debug(raw string) {
	fmt.Printf("%s\n%s\n", runFuncPos(5), raw)
}

func runFuncPos(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}
