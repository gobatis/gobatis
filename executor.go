package batis

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/gobatis/gobatis/logger"
	"github.com/gobatis/gobatis/parser"
	"github.com/gobatis/gobatis/parser/xsql"
)

type executor interface {
	Method() string
	setScan(scan func(scanner) error)
	execute() (sql.Result, error)
}

var (
	_ executor = (*defaultExecutor)(nil)
	_ executor = (*insertBatchExecutor)(nil)
	_ executor = (*parallelQueryExecutor)(nil)
	_ executor = (*pagingQueryExecutor)(nil)
	_ executor = (*fetchQueryExecutor)(nil)
	//_ executor = (*associateQueryExecutor)(nil)
)

var _ sql.Result = (*queryResult)(nil)

type queryResult struct {
	rowsAffected *int64
	lastInserted *int64
}

func (q queryResult) LastInsertId() (int64, error) {
	if q.lastInserted != nil {
		return *q.lastInserted, nil
	}
	return 0, fmt.Errorf("invalid LastInsertId")
}

func (q queryResult) RowsAffected() (int64, error) {
	if q.rowsAffected != nil {
		return *q.rowsAffected, nil
	}
	return 0, fmt.Errorf("invalid RowsAffected")
}

func checkAffect(expect any, result sql.Result) (err error) {
	if expect == nil {
		return
	}
	ac, err := newAffectConstraint(expect)
	if err != nil {
		return
	}
	if result == nil {
		err = fmt.Errorf("expect sql.Result, got nil")
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	err = ac.Check(rowsAffected)
	if err != nil {
		return
	}
	return
}

type defaultExecutor struct {
	method  string
	raw     *raw
	ctx     context.Context
	conn    conn
	logger  logger.Logger
	pos     string
	trace   bool
	debug   bool
	affect  any
	scanner scanner
	scan    func(s scanner) error
}

func (d *defaultExecutor) Method() string {
	return d.method
}

func (d *defaultExecutor) setScan(scan func(scanner) error) {
	d.scan = scan
}

func (d *defaultExecutor) Query() bool {
	return d.raw.Query
}

func (d *defaultExecutor) execute() (result sql.Result, err error) {

	beginAt := time.Now()

	r, err := xsql.Parse(d.logger.Explain, d.raw.SQL, d.raw.Vars)
	if err != nil {
		return
	}

	var rows *sql.Rows
	defer func() {
		if rows != nil {
			err = parser.AddError(err, rows.Close())
		}
		if !d.conn.IsTx() {
			err = parser.AddError(err, d.conn.Close())
		}
		plainSQL, e := xsql.Explain(d.logger.Explain, d.raw.SQL, d.raw.Vars)
		if e != nil {
			plainSQL = fmt.Sprintf("explain sql error: %s", e)
		}
		t := &logger.SQLTrace{
			Trace:    d.trace,
			Debug:    d.debug,
			BeginAt:  beginAt,
			RawSQL:   r.Statement(),
			PlainSQL: plainSQL,
		}
		if result != nil {
			t.RowsAffected, _ = result.RowsAffected()
		}
		d.logger.Trace(d.pos, d.conn.TraceId(), d.conn.IsTx(), err, t)
	}()

	if !d.raw.Query {
		result, err = d.conn.ExecContext(d.ctx, r.Statement(), r.Vars()...)
		if err != nil {
			return
		}
		err = checkAffect(d.affect, result)
		if err != nil {
			return
		}
		return
	}

	rows, err = d.conn.QueryContext(d.ctx, r.Statement(), r.Vars()...)
	if err != nil {
		return
	}
	d.scanner.setRows(rows)
	err = d.scan(d.scanner)
	if err != nil {
		return
	}

	rowsAffected := d.scanner.getRowsAffected()
	result = &queryResult{
		rowsAffected: &rowsAffected,
		lastInserted: nil,
	}
	err = checkAffect(d.affect, result)
	if err != nil {
		return
	}

	return
}

type insertBatchExecutor struct {
	raws    []*raw
	method  string
	ctx     context.Context
	conn    conn
	logger  logger.Logger
	pos     string
	trace   bool
	debug   bool
	affect  any
	scanner scanner
	scan    func(s scanner) error
}

func (i *insertBatchExecutor) Method() string {
	return i.method
}

func (i *insertBatchExecutor) setScan(scan func(scanner) error) {
	i.scan = scan
}

func (i *insertBatchExecutor) execute() (result sql.Result, err error) {
	//c := i.conn
	var tx *connTx
	defer func() {
		// Indicate that the outside is a regular DB object,
		// not a transaction object.
		// TODO TEST maybe beginTx bind tx with conn with luck
		//if !i.conn.IsTx() {
		//	err = parser.AddError(err, i.conn.Close())
		//}
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = parser.AddError(err, tx.Rollback())
				i.logger.Trace(i.pos, tx.TraceId(), true, err, &logger.SQLTrace{
					Trace:    i.trace,
					Debug:    i.debug,
					BeginAt:  now,
					RawSQL:   "rollback",
					PlainSQL: "rollback",
				})
			}
		}
	}()

	if !i.conn.IsTx() {
		now := time.Now()
		tx, err = i.conn.BeginTx(i.ctx, nil)
		if err != nil {
			return
		}
		i.logger.Trace(i.pos, tx.TraceId(), true, err, &logger.SQLTrace{
			Trace:    i.trace,
			Debug:    i.debug,
			BeginAt:  now,
			RawSQL:   "begin",
			PlainSQL: "begin",
		})
	} else {
		tx = i.conn.(*connTx)
	}

	qr := &queryResult{
		rowsAffected: nil,
		lastInserted: nil,
	}
	for _, r := range i.raws {
		d := &defaultExecutor{
			method:  i.method,
			raw:     r,
			ctx:     i.ctx,
			conn:    tx,
			logger:  i.logger,
			pos:     i.pos,
			trace:   i.trace,
			debug:   i.debug,
			scanner: i.scanner,
			scan:    i.scan,
		}
		var rr sql.Result
		rr, err = d.execute()
		if err != nil {
			return
		}
		if n, e := rr.RowsAffected(); e == nil {
			if qr.rowsAffected == nil {
				t := int64(0)
				qr.rowsAffected = &t
			}
			*qr.rowsAffected += n
		}
		if n, e := rr.LastInsertId(); e == nil {
			qr.lastInserted = &n
		}
	}

	result = qr

	err = checkAffect(i.affect, result)
	if err != nil {
		return
	}

	now := time.Now()
	err = tx.Commit()
	if err != nil {
		return
	}

	i.logger.Trace(i.pos, tx.TraceId(), true, err, &logger.SQLTrace{
		Trace:        i.trace,
		Debug:        i.debug,
		BeginAt:      now,
		RawSQL:       "commit",
		PlainSQL:     "commit",
		RowsAffected: 0,
	})

	return
}

type parallelQueryExecutor struct {
	queries []ParallelQuery
	method  string
	ctx     context.Context
	conn    func() conn
	logger  logger.Logger
	pos     string
	trace   bool
	debug   bool
}

func (p parallelQueryExecutor) Method() string {
	return p.method
}

func (p parallelQueryExecutor) setScan(scan func(scanner) error) {
	panic("unimplemented methods were accessed")
}

func (p parallelQueryExecutor) execute() (result sql.Result, err error) {
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, v := range p.queries {
		wg.Add(1)
		go func(v ParallelQuery) {
			defer func() {
				wg.Done()
			}()
			r := newRaw(true, v.SQL, nil)
			r.mergeVars(v.Params)
			d := &defaultExecutor{
				method:  p.method,
				raw:     r,
				ctx:     p.ctx,
				conn:    p.conn(),
				logger:  p.logger,
				pos:     p.pos,
				trace:   p.trace,
				debug:   p.debug,
				scanner: &defaultScanner{},
				scan: func(s scanner) error {
					return v.Scan(s.(Scanner))
				},
			}
			_, e := d.execute()
			if e != nil {
				lock.Lock()
				err = parser.AddError(err, e)
				lock.Unlock()
			}
		}(v)
	}
	wg.Wait()
	return
}

type pagingQueryExecutor struct {
	query  PagingQuery
	method string
	ctx    context.Context
	conn   func() conn
	logger logger.Logger
	pos    string
	trace  bool
	debug  bool
}

func (p pagingQueryExecutor) Method() string {
	return p.method
}

func (p pagingQueryExecutor) setScan(scan func(scanner) error) {
	//TODO implement me
	panic("implement me")
}

func (p pagingQueryExecutor) execute() (sql.Result, error) {
	if p.query.Limit <= 0 {
		return nil, InvalidLimitErr
	}

	w := ""
	if p.query.Where != "" {
		w = fmt.Sprintf(" where %s", p.query.Where)
	}
	o := ""
	if p.query.Order != "" {
		o = fmt.Sprintf(" order by %s", p.query.Order)
	}

	q := newRaw(true, fmt.Sprintf("select %s from %s%s%s limit %d offset %d", p.query.Select, p.query.From, w, o, p.query.Limit, p.query.Limit*p.query.Page), nil)
	q.mergeVars(p.query.Params)

	c := newRaw(true, fmt.Sprintf("select count(%s) from %s%s", p.query.Count, p.query.From, w), nil)
	c.mergeVars(p.query.Params)

	s := &pagingScanner{
		query:  q,
		count:  c,
		method: p.method,
		ctx:    p.ctx,
		conn:   p.conn,
		logger: p.logger,
		pos:    p.pos,
		trace:  p.trace,
		debug:  p.debug,
	}
	return nil, p.query.Scan(s)
}

type fetchQueryExecutor struct {
	limit uint
	*defaultExecutor
}

func (f *fetchQueryExecutor) exec(t *connTx, r *raw, s scanner, c func(scanner) error) error {
	d := &defaultExecutor{
		method:  f.method,
		raw:     r,
		ctx:     f.ctx,
		conn:    t,
		logger:  f.logger,
		pos:     f.pos,
		trace:   f.trace,
		debug:   f.debug,
		affect:  nil,
		scanner: s,
		scan:    c,
	}
	_, err := d.execute()
	return err
}

func (f *fetchQueryExecutor) execute() (result sql.Result, err error) {
	c := f.conn
	var tx *connTx
	defer func() {
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = parser.AddError(err, tx.Rollback())
				f.logger.Trace(f.pos, c.TraceId(), true, err, &logger.SQLTrace{
					Trace:    f.trace,
					Debug:    f.debug,
					BeginAt:  now,
					RawSQL:   "rollback",
					PlainSQL: "rollback",
				})
			}
		}
	}()

	if !c.IsTx() {
		tx, err = c.BeginTx(f.ctx, nil)
		if err != nil {
			return
		}
	} else {
		tx = c.(*connTx)
	}

	// TODO Complete cursor ID
	cursor := fmt.Sprintf("curosr_%s", c.TraceId())

	r := newRaw(false, fmt.Sprintf("declare %s cursor for %s", cursor, f.raw.SQL), nil)
	r.mergeVars(f.raw.Vars)
	err = f.exec(tx, r, nil, nil)
	if err != nil {
		return
	}
	now := time.Now()

	defer func() {
		err = parser.AddError(err, f.exec(tx, newRaw(false, fmt.Sprintf("close %s", cursor), nil), nil, nil))
		err = parser.AddError(err, tx.Commit())
		f.logger.Trace(f.pos, c.TraceId(), true, err, &logger.SQLTrace{
			Trace:    f.trace,
			Debug:    f.debug,
			BeginAt:  now,
			RawSQL:   "commit",
			PlainSQL: "commit",
		})
	}()

	for {
		var rowsAffected int64
		err = f.exec(
			tx,
			newRaw(true, fmt.Sprintf("fetch forward %d from %s", f.limit, cursor), nil),
			&defaultScanner{},
			func(s scanner) error {
				e := f.scan(s)
				if e != nil {
					return e
				}
				rowsAffected += s.getRowsAffected()
				return nil
			},
		)
		if err != nil {
			return
		}
		if rowsAffected == 0 {
			break
		}
	}
	return
}

//	func newAssociateQueryExecutor(base baseExecutor) *associateQueryExecutor {
//		return &associateQueryExecutor{baseExecutor: base}
//	}
//type associateQueryExecutor struct {
//	*defaultExecutor
//}

//
//func (a associateQueryExecutor) execute(f func(s scanner) error) error {
//	//TODO implement me
//	panic("implement me")
//}
//
////func (a associateQueryExecutor) Execute(logger logger.Logger, pos string, trace, debug bool, affect any, s scanner) error {
////	d := newDefaultExecutor(a.conn, a.raw)
////	return d.Execute(logger, pos, trace, debug, affect, s)
////}
//
//func (a associateQueryExecutor) Query() bool {
//	return a.raw.Query
//}
