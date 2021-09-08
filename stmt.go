package gobatis

import (
	"context"
	"database/sql"
	"reflect"
)

type Stmt struct {
	stmt   *sql.Stmt
	exprs  []string
	sql    string
	conn   *sql.Conn
	caller *caller
}

func (p *Stmt) Close() error {
	if p.conn != nil {
		err := p.conn.Close()
		if err != nil {
			return err
		}
	}
	_ = p.stmt.Close()
	return nil
}

func (p *Stmt) exec(tx bool, ctx context.Context, in []reflect.Value) (err error) {
	
	parser := newExprParser(in...)
	for i, v := range p.caller.fragment.in {
		err = parser.paramsStack.list.Front().Next().Value.(*exprParams).bind(v, i)
		if err != nil {
			throw(p.caller.fragment.node.File, p.caller.fragment.node.ctx, parasFragmentErr).with(err)
		}
	}
	
	vars := make([]interface{}, 0)
	for _, v := range p.exprs {
		var _var interface{}
		_var, _, err = parser.parseExpression(p.caller.fragment.node.ctx, v)
		if err != nil {
			return err
		}
		vars = append(vars, _var)
	}
	
	tf := ""
	if tx {
		tf = "[tx]"
	}
	
	p.caller.logger.Debugf("[gobatis] [%s]%s[stmt] exec statement: %s", p.caller.fragment.id, tf, p.sql)
	p.caller.logger.Debugf("[gobatis] [%s]%s[stmt] exec parameter: %s", p.caller.fragment.id, tf, printVars(vars))
	
	res, err := p.stmt.ExecContext(ctx, vars...)
	if err != nil {
		p.caller.logger.Errorf("[gobatis][%s]%s[stmt] exec statement: %s", p.caller.fragment.id, tf, p.sql)
		p.caller.logger.Errorf("[gobatis][%s]%s[stmt] exec parameter: %s", p.caller.fragment.id, tf, printVars(vars))
		p.caller.logger.Errorf("[gobatis][%s]%s[stmt] exec error: %v", p.caller.fragment.id, tf, err)
		return
	}
	
	return p.caller.parseExecResult(res, p.caller.values)
}

func (p *Stmt) query(tx bool, ctx context.Context, in []reflect.Value, values []reflect.Value) (err error) {
	
	parser := newExprParser(in...)
	for i, v := range p.caller.fragment.in {
		err = parser.paramsStack.list.Front().Next().Value.(*exprParams).bind(v, i)
		if err != nil {
			throw(p.caller.fragment.node.File, p.caller.fragment.node.ctx, parasFragmentErr).with(err)
		}
	}
	
	vars := make([]interface{}, 0)
	for _, v := range p.exprs {
		var _var interface{}
		_var, _, err = parser.parseExpression(p.caller.fragment.node.ctx, v)
		if err != nil {
			return err
		}
		vars = append(vars, _var)
	}
	
	tf := ""
	if tx {
		tf = "[tx]"
	}
	
	p.caller.logger.Debugf("[gobatis] [%s]%s[stmt] exec statement: %s", p.caller.fragment.id, tf, p.sql)
	p.caller.logger.Debugf("[gobatis] [%s]%s[stmt] exec parameter: %s", p.caller.fragment.id, tf, printVars(vars))
	
	rows, err := p.stmt.QueryContext(ctx, vars...)
	if err != nil {
		p.caller.logger.Errorf("[gobatis][%s]%s[stmt] query statement: %s", p.caller.fragment.id, tf, p.sql)
		p.caller.logger.Errorf("[gobatis][%s]%s[stmt] query parameter: %s", p.caller.fragment.id, tf, printVars(vars))
		p.caller.logger.Errorf("[gobatis][%s]%s[stmt] query error: %v", p.caller.fragment.id, tf, err)
		return
	}
	
	err = p.caller.parseQueryResult(rows, values)
	if err != nil {
		return
	}
	
	return
}
