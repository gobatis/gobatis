package batis

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

type executor struct {
	query  bool
	sql    string
	logger Logger
	Params []KeyValue
	err    error
	conn   conn
	rows   *sql.Rows
	result *sql.Result
	Debug  bool
	must   bool
}

func (e *executor) Merge(s executor) {
	e.sql = fmt.Sprintf("%s %s", e.sql, s.sql)
	e.Params = append(e.Params, s.Params...)
}

func (e *executor) Exec(s *Scanner) {
	if e.err != nil {
		s.Error = e.err
		return
	}
	
	var err error
	var raw string
	now := time.Now()
	defer func() {
		if err != nil || e.Debug {
			log(e.logger, raw, time.Since(now), err)
		}
		s.Error = err
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
	
	node, err := parseSQL("test.file", fmt.Sprintf("<sql>%s</sql>", e.sql))
	if err != nil {
		return
	}
	
	frag := &fragment{node: node, in: _params}
	
	raw, exprs, vars, dynamic, err := frag.parseStatement(_vars...)
	if err != nil {
		return
	}
	
	//spew.Json(raw, exprs, vars, dynamic)
	_ = exprs
	_ = dynamic
	
	if e.query {
		var result sql.Result
		result, err = e.conn.ExecContext(context.Background(), raw, vars...)
		if err != nil {
			return
		}
		s.result = append(s.result, &result)
		return
	} else {
		var rows *sql.Rows
		rows, err = e.conn.QueryContext(context.Background(), raw, vars...)
		if err != nil {
			return
		}
		s.rows = append(s.rows, rows)
		return
	}
}
