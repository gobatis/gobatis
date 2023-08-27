package batis

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/gozelle/color"
	"go.uber.org/atomic"
)

type conn interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Close() error
}

type tx struct {
	*sql.Tx
}

func (tx) Close() error {
	return nil
}

type executor struct {
	query    bool
	params   []NameValue
	conn     conn
	rows     *sql.Rows
	result   sql.Result
	must     bool
	executed atomic.Bool
	fragment *fragment
	raw      string
	exprs    []string
	vars     []any
	dynamic  bool
	now      time.Time
	dest     any
	tx       bool
}

func (e *executor) exec() (err error) {

	if e.executed.Swap(true) {
		err = fmt.Errorf("db execution was repeated")
		return
	}

	var _params []*param
	var _vars []reflect.Value
	for _, v := range e.params {
		_params = append(_params, &param{
			name:  v.Name,
			_type: reflect.TypeOf(v.Value).Name(),
		})
		_vars = append(_vars, reflect.ValueOf(v.Value))
	}

	var node *xmlNode
	node, err = parseSQL("test.file", fmt.Sprintf("<sql>%s</sql>", e.raw))
	if err != nil {
		return
	}
	e.fragment = &fragment{node: node, in: _params}
	e.raw, e.exprs, e.vars, e.dynamic, err = e.fragment.parseStatement(_vars...)
	if err != nil {
		return
	}

	defer func() {
		if e.rows != nil {
			if v := e.rows.Close(); v != nil {
				err = addError(err, v)
			}
		}
		if !e.tx {
			if v := e.conn.Close(); v != nil {
				err = addError(err, v)
			}
		}
	}()

	if e.query {
		e.rows, err = e.conn.QueryContext(context.Background(), e.raw, e.vars...)
		if err != nil {
			return
		}
		if e.dest != nil {
			err = e.scan(e.rows, e.dest)
			if err != nil {
				return
			}
		}
		return
	} else {
		e.result, err = e.conn.ExecContext(context.Background(), e.raw, e.vars...)
		if err != nil {
			return
		}
		return
	}
}

func (e *executor) scan(rows *sql.Rows, ptr any) (err error) {
	if ptr == nil {
		return
	}
	pv := reflect.ValueOf(ptr)
	if pv.Kind() != reflect.Pointer || pv.IsNil() {
		return &InvalidUnmarshalError{pv.Type()}
	}
	pv = indirect(pv, false)

	columns, err := rows.Columns()
	if err != nil {
		return
	}
	l := len(columns)
	c := 0
	first := false
	for rows.Next() {
		c++
		row := make([]interface{}, l)
		pointers := make([]interface{}, l)
		for i, _ := range columns {
			pointers[i] = &row[i]
		}
		err = rows.Scan(pointers...)
		if err != nil {
			return
		}
		if !first {
			first = true
		}
		var end bool
		end, err = reflectRow(columns, row, pv, first)
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

func (e *executor) log(db *DB) {
	if !db.debug && db.Error == nil {
		return
	}
	cost := time.Since(db.executor.now)
	info := &strings.Builder{}
	var status string
	var out func(format string, a ...any)
	if db.Error != nil {
		status = color.RedString("error")
		out = db.Logger.Errorf
	} else {
		status = color.GreenString("success")
		out = db.Logger.Debugf
	}
	var traceId string
	if db.traceId != "" {
		traceId = fmt.Sprintf("[%s]", color.CyanString(db.traceId))
	}
	var t string
	if db.tx != nil {
		t = fmt.Sprintf("[%s]", color.CyanString("Tx"))
	}
	info.WriteString(fmt.Sprintf("%s %s", color.MagentaString("[gobatis]"), color.RedString(e.runFuncPos(4))))
	info.WriteString(fmt.Sprintf("\n[%s][%s]%s%s %s", status, cost, traceId, t, color.YellowString(db.executor.raw)))
	if db.Error != nil {
		info.WriteString(fmt.Sprintf("\n%s", color.RedString(db.Error.Error())))
	}
	out(info.String())
}

// runFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func (e *executor) runFuncPos(skip int) string {
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
