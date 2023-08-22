package executor

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type conn interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

const (
	Query = iota + 1
	Exec
)

type Executor struct {
	Type   int
	SQL    string
	Logger Logger
	Params []Param
	Err    error
	Conn   conn
	rows   *sql.Rows
	result *sql.Result
	Debug  bool
	must   bool
}

func (e *Executor) Merge(s Executor) {
	e.SQL = fmt.Sprintf("%s %s", e.SQL, s.SQL)
	e.Params = append(e.Params, s.Params...)
}

func (e *Executor) Exec(s *Scanner) {
	if e.Err != nil {
		s.err = e.Err
		return
	}
	
	var err error
	var raw string
	now := time.Now()
	defer func() {
		if err != nil || e.Debug {
			log(e.Logger, raw, time.Since(now), err)
		}
		s.err = err
	}()
	
	var _params []*param
	var _vars []reflect.Value
	for _, v := range e.Params {
		_params = append(_params, &param{
			name:  v.Name,
			_type: reflect.TypeOf(v.Value).Name(),
		})
		_vars = append(_vars, reflect.ValueOf(v.Value))
	}
	
	node, err := parseSQL("test.file", fmt.Sprintf("<sql>%s</sql>", e.SQL))
	if err != nil {
		return
	}
	
	frag := &fragment{node: node, in: _params}
	
	raw, exprs, vars, dynamic, err := frag.parseStatement(_vars...)
	if err != nil {
		return
	}
	//raw = strings.ReplaceAll(raw, "\\u003e", ">")
	//raw = strings.ReplaceAll(raw, "\\u003c", "<")
	
	//spew.Json(raw, exprs, vars, dynamic)
	_ = exprs
	_ = dynamic
	
	switch e.Type {
	case Exec:
		var result sql.Result
		result, err = e.Conn.ExecContext(context.Background(), raw, vars...)
		if err != nil {
			return
		}
		s.result = append(s.result, &result)
		return
	case Query:
		var rows *sql.Rows
		rows, err = e.Conn.QueryContext(context.Background(), raw, vars...)
		if err != nil {
			return
		}
		s.rows = append(s.rows, rows)
		return
	default:
		s.err = fmt.Errorf("unexpect executor type: %d", e.Type)
		return
	}
}
