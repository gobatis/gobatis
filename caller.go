package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dtd"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type conn interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}

//type execer interface {
//	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
//	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
//}
//
//type queryer interface {
//	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
//	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
//}

type segment struct {
	query   bool
	sql     string
	in      []reflect.Value
	vars    []interface{}
	dynamic bool
	ctx     context.Context
	conn    conn
}

//func (p segment) fork() *segment {
//	return &segment{
//		ctx:  p.ctx,
//		conn: p.conn,
//		in:   p.in,
//	}
//}

func (p *segment) merge(s ...*segment) {
	for _, v := range s {
		p.sql += " " + strings.TrimSpace(v.sql)
		if !p.dynamic && v.dynamic {
			p.dynamic = v.dynamic
		}
	}
}

func (p segment) printLog() {
	//p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.method.id, s)
	//p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.method.id, printVars(vars))
	//p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.method.id, err)
}

func (p segment) realSql() string {
	s := p.sql
	for i, v := range p.vars {
		s = strings.Replace(s, fmt.Sprintf("$%d", i+1), p.realValue(v), 1)
	}
	return s
}

func (p segment) realValue(v interface{}) string {
	vv := cast.IndirectToStringerOrError(v)
	switch s := vv.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	case bool:
		return fmt.Sprintf("%v", v)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(uint64(s), 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case []byte:
		return fmt.Sprintf("'%s'", v)
	case nil:
		return ""
	case fmt.Stringer:
		return fmt.Sprintf("'%s'", v)
	case error:
		return fmt.Sprintf("'%s'", v)
	default:
		return ""
	}
}

func (p *segment) concatSQL(s string) {
	if p.sql == "" {
		p.sql = s
	} else {
		p.sql += " " + s
	}
}

type caller struct {
	mt     reflect.Type
	method *method
	logger Logger
	result []reflect.Value
}

func (p *caller) call(in []reflect.Value) *caller {
	start := time.Now()
	defer func() {
		p.logger.Debugf("[gobatis] [%s] cost: %s", p.method.id, time.Since(start))
	}()
	
	var err error
	defer func() {
		if err != nil {
			p.injectError(err)
		}
	}()
	
	switch p.method.node.Name {
	case dtd.SELECT, dtd.INSERT, dtd.DELETE, dtd.UPDATE:
		err = p.exec(in)
	case dtd.SAVE:
		err = p.save(in)
	case dtd.QUERY:
		err = p.query(in)
	default:
		throw(p.method.node.File, p.method.node.ctx, callerErr).
			format("unsupported call method '%s'", p.method.node.Name)
	}
	return p
}

func (p *caller) exec(in []reflect.Value) error {
	s, err := p.method.buildSegment(in)
	if err != nil {
		return err
	}
	return p.run(s)
}

func (p *caller) run(s *segment) (err error) {
	
	if s.conn == nil {
		s.conn, err = p.method.db.Conn(s.ctx)
		if err != nil {
			return
		}
	}
	
	defer func() {
		if s.conn != nil {
			if _err := s.conn.Close(); _err != nil {
				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.method.id, err)
			}
		}
	}()
	
	if s.query {
		var rows *sql.Rows
		rows, err = s.conn.QueryContext(s.ctx, s.sql, s.vars...)
		if err != nil {
			return
		}
		return p.parseQueryResult(rows, p.result)
	}
	
	var r sql.Result
	r, err = s.conn.ExecContext(s.ctx, s.sql, s.vars...)
	if err != nil {
		return
	}
	
	return p.parseExecResult(r, p.result)
}

func (p *caller) injectError(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			if p.method.must {
				p.result = append(p.result, reflect.ValueOf(err))
			} else {
				p.result = append(p.result, reflect.Zero(errorType))
			}
		} else {
			p.result = append(p.result, reflect.ValueOf(err))
		}
	} else {
		p.result = append(p.result, reflect.Zero(errorType))
	}
	if p.mt != nil {
		for i := 0; i < p.mt.NumOut()-1; i++ {
			if p.mt.Out(i).Kind() == reflect.Ptr {
				if err == sql.ErrNoRows {
					p.result[i] = reflect.Zero(p.result[i].Type())
				}
			} else {
				p.result[i] = p.result[i].Elem()
			}
		}
	}
}

func (p caller) query(in []reflect.Value) (err error) {
	
	return
}

func (p caller) save(in []reflect.Value) (err error) {
	
	return
}

//func (p *caller) exec(in ...reflect.Value) (err error) {
//
//	_execer, index := p.execer(in)
//	if index > -1 {
//		in = p.removeParam(in, index)
//	}
//	ctx, index := p.context(in)
//	if index > -1 {
//		in = p.removeParam(in, index)
//	}
//
//	tx, _ := _execer.(*Tx)
//	if tx != nil {
//		stmt := tx.getStmt(p.method.id)
//		if stmt != nil {
//			return stmt.exec(true, ctx, in)
//		}
//	}
//
//	if p.method._stmt != nil {
//		return p.method._stmt.exec(false, ctx, in)
//	}
//
//	var conn *sql.Conn
//	if _execer == nil {
//		conn, err = p.method.db.Conn(ctx)
//		if err != nil {
//			return
//		}
//		_execer = conn
//	}
//	defer func() {
//		if conn != nil && p.method._stmt == nil {
//			if _err := conn.Close(); _err != nil {
//				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.method.id, err)
//			}
//		}
//	}()
//
//	s, exprs, vars, dynamic, err := p.method.buildSegment(in...)
//	if err != nil {
//		return
//	}
//
//	p.logger.Debugf("[gobatis] [%s] exec statement: %s", p.method.id, s)
//	p.logger.Debugf("[gobatis] [%s] exec parameter: %s", p.method.id, printVars(vars))
//
//	var res sql.Result
//	if p.method.stmt {
//		var _stmt *sql.Stmt
//		_stmt, err = _execer.PrepareContext(ctx, s)
//		if err != nil {
//			p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.method.id, s)
//			p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.method.id, printVars(vars))
//			p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.method.id, err)
//			return err
//		}
//
//		if !dynamic {
//			stmt := &Stmt{
//				stmt:   _stmt,
//				exprs:  exprs,
//				sql:    s,
//				conn:   conn,
//				caller: p,
//			}
//			if tx != nil {
//				tx.addStmt(stmt)
//			} else {
//				p.method._stmt = stmt
//			}
//		}
//		res, err = _stmt.ExecContext(ctx, vars...)
//	} else {
//		res, err = _execer.ExecContext(ctx, s, vars...)
//	}
//	if err != nil {
//		p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.method.id, s)
//		p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.method.id, printVars(vars))
//		p.logger.Errorf("[gobatis] [%s] exec error: %v", p.method.id, err)
//		return
//	}
//
//	return p.parseExecResult(res, p.result)
//}

func (p *caller) parseExecResult(res sql.Result, values []reflect.Value) error {
	// ignore RowsAffected to support database that not support
	affected, _ := res.RowsAffected()
	if p.method.must && affected != 1 {
		return fmt.Errorf("expect affect 1 row, got %d", affected)
	}
	return (&execResult{affected: affected, values: values}).scan()
}

//func (p *caller) query(in ...reflect.Value) (err error) {
//
//	ctx, index := p.context(in)
//	if index > -1 {
//		in = p.removeParam(in, index)
//	}
//
//	_queryer, index := p.queryer(in)
//	if index > -1 {
//		in = p.removeParam(in, index)
//	}
//
//	tx, _ := _queryer.(*Tx)
//	if tx != nil {
//		stmt := tx.getStmt(p.method.id)
//		if stmt != nil {
//			err = stmt.query(true, ctx, in, p.result)
//			if err != nil {
//				return
//			}
//			return
//		}
//	}
//
//	if p.method._stmt != nil {
//		err = p.method._stmt.query(false, ctx, in, p.result)
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	var conn *sql.Conn
//	if _queryer == nil {
//		conn, err = p.method.db.Conn(ctx)
//		if err != nil {
//			return
//		}
//		_queryer = conn
//	}
//	defer func() {
//		if conn != nil && p.method._stmt == nil {
//			if _err := conn.Close(); _err != nil {
//				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.method.id, err)
//			}
//		}
//	}()
//
//	s, exprs, vars, dynamic, err := p.method.buildSegment(in...)
//	if err != nil {
//		return
//	}
//
//	p.logger.Debugf("[gobatis] [%s] query statement: %s", p.method.id, s)
//	p.logger.Debugf("[gobatis] [%s] query parameter: [%+v]", p.method.id, printVars(vars))
//
//	var rows *sql.Rows
//	if p.method.stmt {
//		var _stmt *sql.Stmt
//		_stmt, err = _queryer.PrepareContext(ctx, s)
//		if err != nil {
//			p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.method.id, s)
//			p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.method.id, printVars(vars))
//			p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.method.id, err)
//			return err
//		}
//
//		if p.method.stmt && !dynamic {
//			stmt := &Stmt{
//				stmt:   _stmt,
//				exprs:  exprs,
//				sql:    s,
//				conn:   conn,
//				caller: p,
//			}
//			if tx != nil {
//				tx.addStmt(stmt)
//			} else {
//				p.method._stmt = stmt
//			}
//		}
//
//		rows, err = _stmt.QueryContext(ctx, vars...)
//	} else {
//		rows, err = _queryer.QueryContext(ctx, s, vars...)
//	}
//	if err != nil {
//		p.logger.Errorf("[gobatis] [%s] query statement: %s", p.method.id, s)
//		p.logger.Errorf("[gobatis] [%s] query parameter: [%+v]", p.method.id, printVars(vars))
//		p.logger.Errorf("[gobatis] [%s] query error: %v", p.method.id, err)
//		return
//	}
//	err = p.parseQueryResult(rows, p.result)
//	if err != nil {
//		return
//	}
//	return
//}

func (p *caller) parseQueryResult(rows *sql.Rows, values []reflect.Value) (err error) {
	defer func() {
		if _err := rows.Close(); _err != nil {
			p.logger.Errorf("[gobatis] [%s] close rows error: %s", p.method.id, _err)
		}
	}()
	
	res := queryResult{rows: rows}
	err = res.setSelected(p.method.ra, p.method.out, values)
	if err != nil {
		return err
	}
	return res.scan()
}

//func (p *caller) execer(in []reflect.Value) (execer, int) {
//	if len(in) > 0 {
//		t := reflect.TypeOf(new(execer)).Elem()
//		for i, v := range in {
//			if v.Type().Implements(t) {
//				return v.Interface().(execer), i
//			}
//		}
//	}
//	return nil, -1
//}
//
//func (p *caller) queryer(in []reflect.Value) (queryer, int) {
//	if len(in) > 0 {
//		t := reflect.TypeOf(new(queryer)).Elem()
//		for i, v := range in {
//			if v.Type().Implements(t) {
//				return v.Interface().(queryer), i
//			}
//		}
//	}
//	return nil, -1
//}
