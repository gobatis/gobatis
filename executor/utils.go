package executor

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
	
	"github.com/gozelle/color"
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

func log(logger Logger, raw string, cost time.Duration, err error) {
	info := &strings.Builder{}
	var status string
	var out func(format string, a ...any)
	if err != nil {
		status = color.RedString("Error")
		out = logger.Errorf
	} else {
		status = color.GreenString("Success")
		out = logger.Debugf
	}
	info.WriteString(fmt.Sprintf("%s\n[%s][%s]%s", color.RedString(runFuncPos(4)), status, color.WhiteString(cost.String()), color.CyanString("[Tx][1692287996356]")))
	info.WriteString(fmt.Sprintf("\n%s", color.YellowString(raw)))
	out(info.String())
}

// runFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func runFuncPos(skip int) string {
	i := skip
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok || i > 10 {
			break
		}
		if (!strings.Contains(file, "gobatis/executor/") &&
			!strings.Contains(file, "gobatis/db.go")) ||
			strings.HasSuffix(file, "_test.go") {
			return fmt.Sprintf("%s:%d", file, line)
		}
		i++
	}
	return ""
}

//func toSnakeCase(s string) string {
//	var re = regexp.MustCompile(`([^A-Z_])([A-Z])`)
//	snakeStr := re.ReplaceAllString(s, "${1}_${2}")
//	return strings.ToLower(snakeStr)
//}
