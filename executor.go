package batis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gozelle/color"
	"reflect"
	"runtime"
	"strings"
	"time"
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
	now     time.Time
	query   bool
	logger  Logger
	params  []NameValue
	conn    conn
	rows    *sql.Rows
	result  *sql.Result
	debug   bool
	must    bool
	tx      bool
	sql     string
	raw     string
	exprs   []string
	vars    []any
	traceId string
}

func (e executor) log(err error) {
	if !e.debug && err == nil {
		return
	}
	cost := time.Since(e.now)
	info := &strings.Builder{}
	var status string
	var out func(format string, a ...any)
	if err != nil {
		status = color.RedString("Error")
		out = e.logger.Errorf
	} else {
		status = color.GreenString("Success")
		out = e.logger.Debugf
	}
	var traceId string
	if e.traceId != "" {
		traceId = fmt.Sprintf("[%s]", color.CyanString(e.traceId))
	}
	var t string
	if e.tx {
		t = fmt.Sprintf("[%s]", color.CyanString("Tx"))
	}
	info.WriteString(fmt.Sprintf("%s %s", color.MagentaString("[gobatis]"), color.RedString(e.runFuncPos(4))))
	info.WriteString(fmt.Sprintf("\n[%s][%s]%s%s %s", status, cost, traceId, t, color.YellowString(e.sql)))
	if err != nil {
		info.WriteString(fmt.Sprintf("\n%s", color.RedString(err.Error())))
	}
	out(info.String())
}

// runFuncPos returns the file name and line number of the caller of the function calling it.
// skip: 0 for the current function, 1 for the caller of the current function
func (e executor) runFuncPos(skip int) string {
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

func (e executor) Exec(dest any) (rowsAffected *int64, lastedInsertId *int64, err error) {
	
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
	node, err = parseSQL("test.file", fmt.Sprintf("<sql>%s</sql>", e.sql))
	if err != nil {
		return
	}
	
	frag := &fragment{node: node, in: _params}
	
	raw, exprs, vars, dynamic, err := frag.parseStatement(_vars...)
	if err != nil {
		return
	}
	
	_ = exprs
	_ = dynamic
	
	var rows *sql.Rows
	var result sql.Result
	
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	
	if e.query {
		rows, err = e.conn.QueryContext(context.Background(), raw, vars...)
		if err != nil {
			return
		}
		err = e.scan(rows, dest)
		if err != nil {
			return
		}
		return
	} else {
		result, err = e.conn.ExecContext(context.Background(), raw, vars...)
		if err != nil {
			return
		}
		affected, _err := result.RowsAffected()
		if _err == nil {
			rowsAffected = &affected
		}
		id, _err := result.LastInsertId()
		if _err == nil {
			lastedInsertId = &id
		}
		return
	}
}

func (e executor) scan(rows *sql.Rows, ptr any) (err error) {
	if ptr == nil {
		return
	}
	qr := queryResult{
		rows: rows,
	}
	err = qr.scan(ptr)
	if err != nil {
		err = fmt.Errorf("scan rows error: %w", err)
		return
	}
	
	return
}
