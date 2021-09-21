package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/dtd"
	"reflect"
	"time"
)

type caller struct {
	fragment *fragment
	logger   Logger
	args     []reflect.Value
	values   []reflect.Value
	t        reflect.Type
}

func (p *caller) call() *caller {
	var err error
	start := time.Now()
	defer func() {
		p.logger.Debugf("[gobatis] [%s] cost: %s", p.fragment.id, time.Since(start))
	}()
	switch p.fragment.node.Name {
	case dtd.SELECT:
		err = p.query(p.args...)
	case dtd.INSERT, dtd.DELETE, dtd.UPDATE:
		err = p.exec(p.args...)
	default:
		throw(p.fragment.node.File, p.fragment.node.ctx, callerErr).
			format("unsupported call method '%s'", p.fragment.node.Name)
		return nil
	}
	p.handleCallResult(err)
	return p
}

func (p *caller) prepare() {

}

func (p *caller) handleCallResult(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			if p.fragment.must {
				p.values = append(p.values, reflect.ValueOf(err))
			} else {
				p.values = append(p.values, reflect.Zero(errorType))
			}
		} else {
			p.values = append(p.values, reflect.ValueOf(err))
		}
	} else {
		p.values = append(p.values, reflect.Zero(errorType))
	}
	for i := 0; i < p.t.NumOut()-1; i++ {
		if p.t.Out(i).Kind() == reflect.Ptr {
			if err == sql.ErrNoRows {
				p.values[i] = reflect.Zero(p.values[i].Type())
			}
		} else {
			p.values[i] = p.values[i].Elem()
		}
	}
}

func (p *caller) exec(in ...reflect.Value) (err error) {
	
	_execer, index := p.execer(in)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	ctx, index := p.context(in)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	tx, _ := _execer.(*Tx)
	if tx != nil {
		stmt := tx.getStmt(p.fragment.id)
		if stmt != nil {
			return stmt.exec(true, ctx, in)
		}
	}
	
	if p.fragment._stmt != nil {
		return p.fragment._stmt.exec(false, ctx, in)
	}
	
	var conn *sql.Conn
	if _execer == nil {
		conn, err = p.fragment.db.Conn(ctx)
		if err != nil {
			return
		}
		_execer = conn
	}
	defer func() {
		if conn != nil && p.fragment._stmt == nil {
			if _err := conn.Close(); _err != nil {
				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.fragment.id, err)
			}
		}
	}()
	
	s, exprs, vars, dynamic, err := p.fragment.parseStatement(in...)
	if err != nil {
		return
	}
	
	p.logger.Debugf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
	p.logger.Debugf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
	
	var res sql.Result
	if p.fragment.stmt {
		var _stmt *sql.Stmt
		_stmt, err = _execer.PrepareContext(ctx, s)
		if err != nil {
			p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
			p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
			p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.fragment.id, err)
			return err
		}
		
		if !dynamic {
			stmt := &Stmt{
				stmt:   _stmt,
				exprs:  exprs,
				sql:    s,
				conn:   conn,
				caller: p,
			}
			if tx != nil {
				tx.addStmt(stmt)
			} else {
				p.fragment._stmt = stmt
			}
		}
		res, err = _stmt.ExecContext(ctx, vars...)
	} else {
		res, err = _execer.ExecContext(ctx, s, vars...)
	}
	if err != nil {
		p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
		p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
		p.logger.Errorf("[gobatis] [%s] exec error: %v", p.fragment.id, err)
		return
	}
	
	return p.parseExecResult(res, p.values)
}

func (p *caller) parseExecResult(res sql.Result, values []reflect.Value) error {
	// ignore RowsAffected to support database that not support
	affected, _ := res.RowsAffected()
	if p.fragment.must && affected != 1 {
		return fmt.Errorf("expect affect 1 row, got %d", affected)
	}
	return (&execResult{affected: affected, values: values}).scan()
}

func (p *caller) query(in ...reflect.Value) (err error) {
	
	ctx, index := p.context(in)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	_queryer, index := p.queryer(in)
	if index > -1 {
		in = p.removeParam(in, index)
	}
	
	tx, _ := _queryer.(*Tx)
	if tx != nil {
		stmt := tx.getStmt(p.fragment.id)
		if stmt != nil {
			err = stmt.query(true, ctx, in, p.values)
			if err != nil {
				return
			}
			return
		}
	}
	
	if p.fragment._stmt != nil {
		err = p.fragment._stmt.query(false, ctx, in, p.values)
		if err != nil {
			return
		}
		return
	}
	
	var conn *sql.Conn
	if _queryer == nil {
		conn, err = p.fragment.db.Conn(ctx)
		if err != nil {
			return
		}
		_queryer = conn
	}
	defer func() {
		if conn != nil && p.fragment._stmt == nil {
			if _err := conn.Close(); _err != nil {
				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.fragment.id, err)
			}
		}
	}()
	
	s, exprs, vars, dynamic, err := p.fragment.parseStatement(in...)
	if err != nil {
		return
	}
	
	p.logger.Debugf("[gobatis] [%s] query statement: %s", p.fragment.id, s)
	p.logger.Debugf("[gobatis] [%s] query parameter: [%+v]", p.fragment.id, printVars(vars))
	
	var rows *sql.Rows
	if p.fragment.stmt {
		var _stmt *sql.Stmt
		_stmt, err = _queryer.PrepareContext(ctx, s)
		if err != nil {
			p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
			p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
			p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.fragment.id, err)
			return err
		}
		
		if p.fragment.stmt && !dynamic {
			stmt := &Stmt{
				stmt:   _stmt,
				exprs:  exprs,
				sql:    s,
				conn:   conn,
				caller: p,
			}
			if tx != nil {
				tx.addStmt(stmt)
			} else {
				p.fragment._stmt = stmt
			}
		}
		
		rows, err = _stmt.QueryContext(ctx, vars...)
	} else {
		rows, err = _queryer.QueryContext(ctx, s, vars...)
	}
	if err != nil {
		p.logger.Errorf("[gobatis] [%s] query statement: %s", p.fragment.id, s)
		p.logger.Errorf("[gobatis] [%s] query parameter: [%+v]", p.fragment.id, printVars(vars))
		p.logger.Errorf("[gobatis] [%s] query error: %v", p.fragment.id, err)
		return
	}
	err = p.parseQueryResult(rows, p.values)
	if err != nil {
		return
	}
	return
}

func (p *caller) parseQueryResult(rows *sql.Rows, values []reflect.Value) (err error) {
	defer func() {
		if _err := rows.Close(); _err != nil {
			p.logger.Errorf("[gobatis] [%s] close rows error: %s", p.fragment.id, _err)
		}
	}()
	
	res := queryResult{rows: rows}
	err = res.setSelected(p.fragment.resultAttribute, p.fragment.out, values)
	if err != nil {
		return err
	}
	return res.scan()
}

func (p *caller) removeParam(a []reflect.Value, i int) []reflect.Value {
	return append(a[:i], a[i+1:]...)
}

func (p *caller) context(in []reflect.Value) (context.Context, int) {
	for i, v := range in {
		if isContext(v.Type()) {
			return v.Interface().(context.Context), i
		}
	}
	return context.Background(), -1
}

func (p *caller) execer(in []reflect.Value) (execer, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(execer)).Elem()
		for i, v := range in {
			if v.Type().Implements(t) {
				return v.Interface().(execer), i
			}
		}
	}
	return nil, -1
}

func (p *caller) queryer(in []reflect.Value) (queryer, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(queryer)).Elem()
		for i, v := range in {
			if v.Type().Implements(t) {
				return v.Interface().(queryer), i
			}
		}
	}
	return nil, -1
}
