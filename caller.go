package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dtd"
	"reflect"
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

func (p segment) fork() *segment {
	return &segment{
		ctx:  p.ctx,
		conn: p.conn,
	}
}

func (p *segment) merge(s ...*segment) {
	for _, v := range s {
		p.sql += " " + strings.TrimSpace(v.sql)
		if !p.dynamic && v.dynamic {
			p.dynamic = v.dynamic
		}
	}
}

func (p segment) printLog() {
	//p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
	//p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
	//p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.fragment.id, err)
}

func (p *segment) realSql() string {
	s := p.sql
	for i, v := range p.vars {
		s = strings.Replace(s, fmt.Sprintf("$%d", i+1), fmt.Sprintf("%v", v), 1)
	}
	return s
}

type caller struct {
	mt       reflect.Type
	fragment *method
	logger   Logger
	result   []reflect.Value
}

type blocks struct {
	items map[string]*xmlNode
}

func (p blocks) get(name string) *xmlNode {
	if p.items == nil {
		return nil
	}
	return p.items[name]
}

func (p blocks) len() int {
	return len(p.items)
}

func (p *caller) call(in []reflect.Value) *caller {
	
	start := time.Now()
	defer func() {
		p.logger.Debugf("[gobatis] [%s] cost: %s", p.fragment.id, time.Since(start))
	}()
	
	var err error
	defer func() {
		e := recover()
		if err == nil {
			err = catch(p.fragment.node.File, e)
		}
		if err != nil {
			p.injectError(err)
		}
	}()
	
	switch p.fragment.node.Name {
	case dtd.SELECT, dtd.INSERT, dtd.DELETE, dtd.UPDATE:
		s := p.prepareSegment(in)
		var parser *exprParser
		parser, err = p.fragment.prepareParser(s.in)
		err = p.fragment.buildSegment(parser, s, p.fragment.node)
		if err != nil {
			return p
		}
		err = p.exec(s)
	case dtd.SAVE:
		err = p.save(in)
	case dtd.QUERY:
		err = p.query(in)
	default:
		throw(p.fragment.node.File, p.fragment.node.ctx, callerErr).
			format("unsupported call method '%s'", p.fragment.node.Name)
	}
	return p
}

func (p *caller) prepareSegment(in []reflect.Value) (s *segment) {
	var index int
	s = &segment{
		query: p.fragment.node.Name == dtd.SELECT,
		in:    in,
	}
	s.ctx, index = p.context(in)
	if index > -1 {
		s.in = p.removeParam(s.in, index)
	}
	s.conn, index = p.conn(in)
	if index > -1 {
		s.in = p.removeParam(s.in, index)
	}
	return
}

func (p *caller) exec(s *segment) (err error) {
	
	if s.conn == nil {
		s.conn, err = p.fragment.db.Conn(s.ctx)
		if err != nil {
			return
		}
	}
	
	defer func() {
		if s.conn != nil {
			if _err := s.conn.Close(); _err != nil {
				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.fragment.id, err)
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
			if p.fragment.must {
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

func (p *caller) query(in []reflect.Value) (err error) {
	var (
		cs     *segment
		ss     *segment
		parser *exprParser
	)
	bs := p.prepareBlocks()
	cn := bs.get(dtd.BLOCK_COUNT)
	if cn != nil {
		cs = p.prepareSegment(in)
	}
	sn := bs.get(dtd.BLOCK_SELECT)
	if sn != nil {
		if cs != nil {
			ss = cs.fork()
		} else {
			ss = p.prepareSegment(in)
		}
	}
	fn := bs.get(dtd.BLOCK_FROM)
	ln := bs.get(dtd.BLOCK_LIMIT)
	if cs != nil {
		parser, err = p.fragment.prepareParser(cs.in)
		if err != nil {
			return
		}
		err = p.fragment.buildSegment(parser, cs, sn, fn, ln)
		if err != nil {
			return
		}
	}
	if ss != nil {
		if parser == nil {
			parser, err = p.fragment.prepareParser(ss.in)
			if err != nil {
				return
			}
		}
		err = p.fragment.buildSegment(parser, ss, sn, fn, ln)
		if err != nil {
			return
		}
	}
	return
}

func (p *caller) save(in []reflect.Value) (err error) {
	
	s := p.prepareSegment(in)
	parser, err := p.fragment.prepareParser(s.in)
	if err != nil {
		return
	}
	bs := p.prepareBlocks()
	n := bs.get(dtd.BLOCK_INSERT)
	update := true
	if n != nil {
		var v interface{}
		v, _, err = parser.parseExpression(n.ctx, n.GetAttribute(dtd.TEST))
		if err != nil {
			return
		}
		var ok bool
		ok, err = cast.ToBoolE(v)
		if err != nil {
			return
		}
		if ok {
			update = false
			err = p.fragment.buildSegment(parser, s, n)
			if err != nil {
				return
			}
		}
	}
	
	if update {
		n = bs.get(dtd.BLOCK_SELECT)
		if n != nil {
			err = p.fragment.buildSegment(parser, s, n)
			if err != nil {
				return
			}
		}
	}
	
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
//		stmt := tx.getStmt(p.fragment.id)
//		if stmt != nil {
//			return stmt.exec(true, ctx, in)
//		}
//	}
//
//	if p.fragment._stmt != nil {
//		return p.fragment._stmt.exec(false, ctx, in)
//	}
//
//	var conn *sql.Conn
//	if _execer == nil {
//		conn, err = p.fragment.db.Conn(ctx)
//		if err != nil {
//			return
//		}
//		_execer = conn
//	}
//	defer func() {
//		if conn != nil && p.fragment._stmt == nil {
//			if _err := conn.Close(); _err != nil {
//				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.fragment.id, err)
//			}
//		}
//	}()
//
//	s, exprs, vars, dynamic, err := p.fragment.buildSegment(in...)
//	if err != nil {
//		return
//	}
//
//	p.logger.Debugf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
//	p.logger.Debugf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
//
//	var res sql.Result
//	if p.fragment.stmt {
//		var _stmt *sql.Stmt
//		_stmt, err = _execer.PrepareContext(ctx, s)
//		if err != nil {
//			p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
//			p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
//			p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.fragment.id, err)
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
//				p.fragment._stmt = stmt
//			}
//		}
//		res, err = _stmt.ExecContext(ctx, vars...)
//	} else {
//		res, err = _execer.ExecContext(ctx, s, vars...)
//	}
//	if err != nil {
//		p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
//		p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
//		p.logger.Errorf("[gobatis] [%s] exec error: %v", p.fragment.id, err)
//		return
//	}
//
//	return p.parseExecResult(res, p.result)
//}

func (p *caller) parseExecResult(res sql.Result, values []reflect.Value) error {
	// ignore RowsAffected to support database that not support
	affected, _ := res.RowsAffected()
	if p.fragment.must && affected != 1 {
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
//		stmt := tx.getStmt(p.fragment.id)
//		if stmt != nil {
//			err = stmt.query(true, ctx, in, p.result)
//			if err != nil {
//				return
//			}
//			return
//		}
//	}
//
//	if p.fragment._stmt != nil {
//		err = p.fragment._stmt.query(false, ctx, in, p.result)
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	var conn *sql.Conn
//	if _queryer == nil {
//		conn, err = p.fragment.db.Conn(ctx)
//		if err != nil {
//			return
//		}
//		_queryer = conn
//	}
//	defer func() {
//		if conn != nil && p.fragment._stmt == nil {
//			if _err := conn.Close(); _err != nil {
//				p.logger.Errorf("[gobatis] [%s] close conn error: %s", p.fragment.id, err)
//			}
//		}
//	}()
//
//	s, exprs, vars, dynamic, err := p.fragment.buildSegment(in...)
//	if err != nil {
//		return
//	}
//
//	p.logger.Debugf("[gobatis] [%s] query statement: %s", p.fragment.id, s)
//	p.logger.Debugf("[gobatis] [%s] query parameter: [%+v]", p.fragment.id, printVars(vars))
//
//	var rows *sql.Rows
//	if p.fragment.stmt {
//		var _stmt *sql.Stmt
//		_stmt, err = _queryer.PrepareContext(ctx, s)
//		if err != nil {
//			p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
//			p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
//			p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.fragment.id, err)
//			return err
//		}
//
//		if p.fragment.stmt && !dynamic {
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
//				p.fragment._stmt = stmt
//			}
//		}
//
//		rows, err = _stmt.QueryContext(ctx, vars...)
//	} else {
//		rows, err = _queryer.QueryContext(ctx, s, vars...)
//	}
//	if err != nil {
//		p.logger.Errorf("[gobatis] [%s] query statement: %s", p.fragment.id, s)
//		p.logger.Errorf("[gobatis] [%s] query parameter: [%+v]", p.fragment.id, printVars(vars))
//		p.logger.Errorf("[gobatis] [%s] query error: %v", p.fragment.id, err)
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
			p.logger.Errorf("[gobatis] [%s] close rows error: %s", p.fragment.id, _err)
		}
	}()
	
	res := queryResult{rows: rows}
	err = res.setSelected(p.fragment.ra, p.fragment.out, values)
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
		if isCtx(v.Type()) {
			return v.Interface().(context.Context), i
		}
	}
	return context.Background(), -1
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

func (p *caller) conn(in []reflect.Value) (conn, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(conn)).Elem()
		for i, v := range in {
			if v.Type().Implements(t) {
				return v.Interface().(conn), i
			}
		}
	}
	return nil, -1
}

func (p caller) prepareBlocks() *blocks {
	bs := new(blocks)
	if len(p.fragment.node.Nodes) > 0 {
		bs.items = map[string]*xmlNode{}
	}
	for _, v := range p.fragment.node.Nodes {
		if v.Name == dtd.BLOCK {
			bs.items[v.GetAttribute(dtd.TYPE)] = v
		} else if !v.EmptyText() {
			throw(p.fragment.node.File, p.fragment.node.ctx, parseFragmentErr).
				with(fmt.Errorf("unsupported ohter element"))
		}
	}
	return bs
}
