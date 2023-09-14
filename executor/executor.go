package executor

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type Executor interface {
	Execute(logger Logger, trace, debug bool, affecting any, scan func(s Scanner) error) (err error)
	Query() bool
}

var (
	_ Executor = (*Default)(nil)
	_ Executor = (*InsertBatch)(nil)
	_ Executor = (*ParallelQuery)(nil)
	_ Executor = (*FetchQuery)(nil)
)

func NewDefault(conn Conn, raw *Raw) *Default {
	return &Default{conn: conn, raw: raw, clean: true}
}

type Default struct {
	fragment *fragment
	exprs    []string
	vars     []any
	dynamic  bool
	sql      string
	rows     *sql.Rows
	result   sql.Result
	conn     Conn
	raw      *Raw
	clean    bool
}

func (d *Default) Query() bool {
	return d.raw.Query
}

func (d *Default) Execute(logger Logger, trace, debug bool, affect any, scan func(Scanner) error) (err error) {
	
	beginAt := time.Now()
	
	var params []*param
	var vars []reflect.Value
	for _, v := range d.raw.Params {
		params = append(params, &param{
			name: v.Name,
			rt:   reflect.TypeOf(v.Value).Name(),
		})
		vars = append(vars, reflect.ValueOf(v.Value))
	}
	
	var node *xmlNode
	node, err = parseSQL("test.file", fmt.Sprintf("<sql>%s</sql>", d.raw.SQL))
	if err != nil {
		return
	}
	
	d.fragment = &fragment{node: node, in: params}
	d.sql, d.exprs, d.vars, d.dynamic, err = d.fragment.parseStatement(vars...)
	if err != nil {
		return
	}
	
	var s *scanner
	
	defer func() {
		if d.clean {
			if d.rows != nil {
				err = AddError(err, d.rows.Close())
			}
			if !d.conn.IsTx() {
				err = AddError(err, d.conn.Close())
			}
		}
		if s == nil {
			s = &scanner{}
		}
		logger.Trace(d.conn.TraceId(), d.conn.IsTx(), err, &SQLTrace{
			Trace:        trace,
			Debug:        debug,
			BeginAt:      beginAt,
			RawSQL:       d.sql,
			PlainSQL:     "",
			RowsAffected: s.rowsAffected,
		})
	}()
	
	if d.raw.Query {
		d.rows, err = d.conn.QueryContext(d.raw.Ctx, d.sql, d.vars...)
		if err != nil {
			return
		}
	} else {
		d.result, err = d.conn.ExecContext(d.raw.Ctx, d.sql, d.vars...)
		if err != nil {
			return
		}
	}
	
	if d.raw.Query {
		s = &scanner{
			rows:         d.rows,
			rowsAffected: 0,
			lastInsertId: 0,
		}
		err = scan(s)
	} else {
		rowsAffected, _ := d.result.RowsAffected()
		lastInsertId, _ := d.result.LastInsertId()
		s = &scanner{
			rows:         nil,
			rowsAffected: rowsAffected,
			lastInsertId: lastInsertId,
		}
		err = scan(s)
	}
	if err != nil {
		return
	}
	
	if affect != nil {
		var ac *affectConstraint
		ac, err = newAffectConstraint(affect)
		if err != nil {
			return
		}
		err = ac.Check(int(s.rowsAffected))
		if err != nil {
			return
		}
	}
	
	return
}

func NewInsertBatch(ctx context.Context, conn Conn, raws []*Raw) *InsertBatch {
	return &InsertBatch{ctx: ctx, conn: conn, raws: raws}
}

type InsertBatch struct {
	ctx  context.Context
	conn Conn
	raws []*Raw
}

func (i *InsertBatch) Query() bool {
	for _, v := range i.raws {
		if v.Query {
			return v.Query
		}
	}
	return false
}

func (i *InsertBatch) Execute(logger Logger, trace, debug bool, affect any, scan func(Scanner) error) (err error) {
	
	conn := i.conn
	
	var tx *sql.Tx
	defer func() {
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = AddError(err, tx.Rollback())
				logger.Trace(conn.TraceId(), true, err, &SQLTrace{
					Trace:        trace,
					Debug:        debug,
					BeginAt:      now,
					RawSQL:       "rollback",
					PlainSQL:     "rollback",
					RowsAffected: 0,
				})
			}
		}
	}()
	
	if !conn.IsTx() {
		now := time.Now()
		tx, err = conn.BeginTx(i.ctx, nil)
		if err != nil {
			return
		}
		conn = NewTx(tx, conn.TraceId())
		logger.Trace(conn.TraceId(), true, err, &SQLTrace{
			Trace:        trace,
			Debug:        debug,
			BeginAt:      now,
			RawSQL:       "begin",
			PlainSQL:     "begin",
			RowsAffected: 0,
		})
	}
	
	for _, raw := range i.raws {
		err = i.execute(conn, raw, logger, trace, debug, nil, func(s Scanner) error {
			return scan(s)
		})
		if err != nil {
			return
		}
	}
	
	//if affect != nil {
	//	var ac *affectConstraint
	//	ac, err = newAffectConstraint(affect)
	//	if err != nil {
	//		return
	//	}
	//	err = ac.Check(int(s.rowsAffected))
	//	if err != nil {
	//		return
	//	}
	//}
	
	now := time.Now()
	err = tx.Commit()
	if err != nil {
		return
	}
	
	logger.Trace(conn.TraceId(), true, err, &SQLTrace{
		Trace:        trace,
		Debug:        debug,
		BeginAt:      now,
		RawSQL:       "commit",
		PlainSQL:     "commit",
		RowsAffected: 0,
	})
	
	return
}

func (i *InsertBatch) execute(conn Conn, raw *Raw, logger Logger, trace, debug bool, affect any, scan func(Scanner) error) (err error) {
	d := NewDefault(conn, raw)
	d.clean = false
	err = d.Execute(logger, trace, debug, nil, nil)
	if err != nil {
		return
	}
	defer func() {
		if d.rows != nil {
			err = AddError(err, d.rows.Close())
		}
	}()
	ibs := &insertBatchScanner{}
	ibs.rows = d.rows
	ibs.result = d.result
	err = scan(ibs)
	if err != nil {
		return
	}
	return
}

type ParallelQuery struct {
}

func (p *ParallelQuery) Query() bool {
	return true
}

func (p *ParallelQuery) Execute(logger Logger, trace, debug bool, affecting any, scan func(Scanner) error) (err error) {
	
	return
}

func NewFetchQuery(ctx context.Context, conn Conn, raw *Raw, limit uint) *FetchQuery {
	return &FetchQuery{ctx: ctx, conn: conn, raw: raw, limit: limit}
}

type FetchQuery struct {
	ctx   context.Context
	conn  Conn
	raw   *Raw
	limit uint
}

func (f *FetchQuery) Query() bool {
	return f.raw.Query
}

func (f *FetchQuery) Execute(logger Logger, trace, debug bool, affect any, scan func(Scanner) error) (err error) {
	
	conn := f.conn
	
	var tx *sql.Tx
	defer func() {
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = AddError(err, tx.Rollback())
				logger.Trace(conn.TraceId(), true, err, &SQLTrace{
					Trace:        trace,
					Debug:        debug,
					BeginAt:      now,
					RawSQL:       "rollback",
					PlainSQL:     "rollback",
					RowsAffected: 0,
				})
			}
		}
	}()
	
	if !conn.IsTx() {
		tx, err = conn.BeginTx(f.ctx, nil)
		if err != nil {
			return
		}
		conn = NewTx(tx, conn.TraceId())
	}
	
	cursor := fmt.Sprintf("curosr_%s", conn.TraceId())
	
	d := NewDefault(conn, &Raw{
		Ctx:    f.ctx,
		Query:  false,
		SQL:    fmt.Sprintf("declare %s cursor for %s", cursor, f.raw.SQL),
		Params: f.raw.Params,
	})
	err = d.Execute(logger, trace, debug, affect, nil)
	if err != nil {
		return
	}
	
	defer func() {
		d = NewDefault(conn, &Raw{
			Ctx:    f.ctx,
			Query:  false,
			SQL:    fmt.Sprintf("close %s", cursor),
			Params: nil,
		})
		err = AddError(err, d.Execute(logger, trace, debug, affect, nil))
		
		now := time.Now()
		err = AddError(err, tx.Commit())
		
		logger.Trace(conn.TraceId(), true, err, &SQLTrace{
			Trace:        trace,
			Debug:        debug,
			BeginAt:      now,
			RawSQL:       "commit",
			PlainSQL:     "commit",
			RowsAffected: 0,
		})
	}()
	
	for {
		d = NewDefault(conn, &Raw{
			Ctx:    f.ctx,
			Query:  true,
			SQL:    fmt.Sprintf("fetch forward %d from %s", f.limit, cursor),
			Params: nil,
		})
		var rowsAffected int64
		err = d.Execute(logger, trace, debug, affect, func(s Scanner) error {
			e := scan(s)
			if e != nil {
				return e
			}
			rowsAffected = s.RowsAffected()
			return nil
		})
		if err != nil {
			return
		}
		if rowsAffected == 0 {
			break
		}
	}
	
	return
}
