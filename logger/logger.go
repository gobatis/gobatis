package logger

import (
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gozelle/color"
)

type Logger interface {
	Debugf(format string, a ...any)
	Infof(format string, a ...any)
	Errorf(format string, a ...any)
	Warnf(format string, a ...any)
	Trace(pos, id string, tx bool, err error, st *SQLTrace)
	Explain(rv reflect.Value, escaper string) (s string, err error)
}

type Writer interface {
	Printf(string, ...interface{})
}

func NewtLogger(w Writer) Logger {
	return &logger{
		Writer: w,
	}
}

var _ Logger = (*logger)(nil)

var Default = NewtLogger(log.New(os.Stdout, "\r\n", log.LstdFlags))

type logger struct {
	Writer
}

func (l logger) Trace(pos, traceId string, tx bool, err error, tr *SQLTrace) {
	if !tr.Trace && !tr.Debug && err == nil {
		return
	}
	info := &strings.Builder{}
	var f func(format string, a ...any)
	if err != nil {
		f = l.Errorf
	} else {
		f = l.Debugf
	}
	if traceId != "" {
		traceId = fmt.Sprintf("[%s]", color.CyanString(traceId))
	}
	var t string
	if tx {
		t = fmt.Sprintf("[%s]", color.MagentaString("tx"))
	}
	if pos == "" {
		pos = CallFuncPos(7)
	}
	info.WriteString(fmt.Sprintf("%s%s %s", traceId, t, color.RedString(pos)))
	if err != nil {
		info.WriteString(color.RedString(fmt.Sprintf(" error: %s", err.Error())))
	}

	if tr != nil {
		cost := time.Since(tr.BeginAt)
		info.WriteString(fmt.Sprintf("\n%s %s %s",
			color.YellowString(fmt.Sprintf("[%s]", cost)),
			color.BlueString(fmt.Sprintf("[rows:%d]", tr.RowsAffected)),
			tr.PlainSQL,
		))
	}
	f(info.String())
}

func (l logger) Debugf(format string, a ...any) {
	l.Printf(format, a...)
}

func (l logger) Infof(format string, a ...any) {
	l.Printf(format, a...)
}

func (l logger) Errorf(format string, a ...any) {
	l.Printf(format, a...)
}

func (l logger) Warnf(format string, a ...any) {
	l.Printf(format, a...)
}

const (
	null = "null"
	tsf  = "2006-01-02 15:04:05.999"
	tsz  = "0000-00-00 00:00:00"
)

func elemOf(rv reflect.Value) reflect.Value {
	for {
		if rv.Kind() != reflect.Pointer {
			return rv
		}
		rv = rv.Elem()
	}
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func (l logger) Explain(rv reflect.Value, escaper string) (s string, err error) {

	escape := func(v string) string {
		return fmt.Sprintf("%s%s%s", escaper, v, escaper)
	}
	if rv.Kind() == reflect.Invalid || rv.Kind() == reflect.Pointer && rv.IsNil() {
		s = null
		return
	}
	rv = elemOf(rv)
	switch rv.Kind() {
	case reflect.String:
		s = escape(strings.ReplaceAll(fmt.Sprintf("%s", rv.Interface()), escaper, "\\"+escaper))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s = fmt.Sprintf("%d", rv.Interface())
	case reflect.Float32, reflect.Float64:
		s = fmt.Sprintf("%.6f", rv.Interface())
	case reflect.Bool:
		s = strconv.FormatBool(rv.Interface().(bool))
	default:
		switch t := rv.Interface().(type) {
		case time.Time:
			if rv.IsZero() {
				s = escape(tsz)
			} else {
				s = escape(t.Format(tsf))
			}
		case []byte:
			if vv := string(t); isPrintable(vv) {
				s = escape(strings.ReplaceAll(s, escaper, "\\"+escaper))
			} else {
				s = escape("<binary>")
			}
		case fmt.Stringer:
			s = escape(t.String())
		case driver.Valuer:
			var vv driver.Value
			vv, err = t.Value()
			if err != nil {
				return
			}
			s, err = l.Explain(reflect.ValueOf(vv), escaper)
		default:
			err = fmt.Errorf("unsupported explain type: %s", rv.Type())
		}
	}
	return
}

type SQLTrace struct {
	Trace        bool
	Debug        bool
	BeginAt      time.Time
	RawSQL       string
	PlainSQL     string
	RowsAffected int64
}

type TraceSQL struct {
}

type TraceLog struct {
	Duration     time.Duration
	Location     string
	RawSQL       string
	RowsAffected int64
	Error        string
}

// CallFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func CallFuncPos(skip int) string {
	i := skip
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok || i > 20 {
			break
		}
		if !strings.Contains(file, "/gobatis/gobatis") || strings.HasSuffix(file, "_test.go") {
			return fmt.Sprintf("%s:%d", file, line)
		}
		i++
	}
	return ""
}
