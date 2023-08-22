package batis

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

type conn interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type executor struct {
	query  bool
	sql    string
	logger Logger
	params []NameValue
	conn   conn
	rows   *sql.Rows
	result *sql.Result
	tracer *tracer
	debug  bool
	must   bool
}

func (e *executor) Exec(s *Scanner) {
	
	defer func() {
		if e.tracer.err != nil {
			e.tracer.log()
		}
	}()
	
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
	node, e.tracer.err = parseSQL("test.file", fmt.Sprintf("<sql>%s</sql>", e.sql))
	if e.tracer.err != nil {
		return
	}
	
	frag := &fragment{node: node, in: _params}
	
	e.tracer.raw, e.tracer.exprs, e.tracer.vars, e.tracer.dynamic, e.tracer.err = frag.parseStatement(_vars...)
	if e.tracer.err != nil {
		return
	}
	
	if e.query {
		var result sql.Result
		result, e.tracer.err = e.conn.ExecContext(context.Background(), e.tracer.raw, e.tracer.vars...)
		if e.tracer.err != nil {
			return
		}
		s.result = append(s.result, &result)
		return
	} else {
		var rows *sql.Rows
		rows, e.tracer.err = e.conn.QueryContext(context.Background(), e.tracer.raw, e.tracer.vars...)
		if e.tracer.err != nil {
			return
		}
		s.rows = append(s.rows, rows)
		return
	}
}
