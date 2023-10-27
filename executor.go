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
	method() string
	setScan(scan func(scanner) error)
	execute() (sql.Result, error)
}

var (
	_ executor = (*defaultExecutor)(nil)
	_ executor = (*insertBatchExecutor)(nil)
	_ executor = (*parallelQueryExecutor)(nil)
	_ executor = (*pagingQueryExecutor)(nil)
	_ executor = (*fetchQueryExecutor)(nil)
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

func withTx(log logger.Logger, pos string, trace, debug bool, ctx context.Context, c conn, f func(tx *connTx) error) (err error) {
	var tx *connTx
	defer func() {
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = parser.AddError(err, tx.Rollback())
				log.Trace(pos, tx.TraceId(), true, err, &logger.SQLTrace{
					Trace:    trace,
					Debug:    debug,
					BeginAt:  now,
					RawSQL:   "rollback",
					PlainSQL: "rollback",
				})
			}
		}
		// Indicate that the outside is a regular DB object,
		if !c.IsTx() {
			err = parser.AddError(err, c.Close())
		}
	}()

	if !c.IsTx() {
		now := time.Now()
		tx, err = c.BeginTx(ctx, nil)
		if err != nil {
			return
		}
		log.Trace(pos, tx.TraceId(), true, err, &logger.SQLTrace{
			Trace:    trace,
			Debug:    debug,
			BeginAt:  now,
			RawSQL:   "begin",
			PlainSQL: "begin",
		})
	} else {
		tx = c.(*connTx)
	}

	err = f(tx)
	if err != nil {
		return
	}

	now := time.Now()
	err = tx.Commit()
	if err != nil {
		return
	}
	log.Trace(pos, tx.TraceId(), true, err, &logger.SQLTrace{
		Trace:        trace,
		Debug:        debug,
		BeginAt:      now,
		RawSQL:       "commit",
		PlainSQL:     "commit",
		RowsAffected: 0,
	})
	return
}

type defaultExecutor struct {
	name    string
	raw     *raw
	ctx     context.Context
	conn    conn
	logger  logger.Logger
	pos     string
	trace   bool
	debug   bool
	tx      bool
	affect  any
	scanner scanner
	scan    func(s scanner) error
}

func (d *defaultExecutor) method() string {
	return d.name
}

func (d *defaultExecutor) setScan(scan func(scanner) error) {
	d.scan = scan
}

func (d *defaultExecutor) Query() bool {
	return d.raw.Query
}

func (d *defaultExecutor) execute() (sql.Result, error) {
	if d.tx {
		var result sql.Result
		err := withTx(d.logger, d.pos, d.trace, d.debug, d.ctx, d.conn, func(tx *connTx) error {
			var err error
			result, err = d.f(tx)
			return err
		})
		return result, err
	}
	return d.f(d.conn)
}

func (d *defaultExecutor) f(c conn) (result sql.Result, err error) {

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
		if !c.IsTx() && !d.tx {
			err = parser.AddError(err, c.Close())
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
		d.logger.Trace(d.pos, c.TraceId(), c.IsTx(), err, t)
	}()

	if !d.raw.Query {
		result, err = c.ExecContext(d.ctx, r.Statement(), r.Vars()...)
		if err != nil {
			return
		}
		err = checkAffect(d.affect, result)
		if err != nil {
			return
		}
		return
	}

	rows, err = c.QueryContext(d.ctx, r.Statement(), r.Vars()...)
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
	name    string
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

func (i *insertBatchExecutor) method() string {
	return i.name
}

func (i *insertBatchExecutor) setScan(scan func(scanner) error) {
	i.scan = scan
}

func (i *insertBatchExecutor) execute() (sql.Result, error) {
	var result sql.Result
	err := withTx(i.logger, i.pos, i.trace, i.debug, i.ctx, i.conn, func(tx *connTx) (err error) {

		qr := &queryResult{
			rowsAffected: nil,
			lastInserted: nil,
		}
		for _, r := range i.raws {
			d := &defaultExecutor{
				name:    i.name,
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
		return
	})

	return result, err
}

type parallelQueryExecutor struct {
	queries []ParallelQuery
	name    string
	ctx     context.Context
	conn    func() conn
	logger  logger.Logger
	pos     string
	trace   bool
	debug   bool
}

func (p parallelQueryExecutor) method() string {
	return p.name
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
				name:    p.name,
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
	name   string
	ctx    context.Context
	conn   func() conn
	logger logger.Logger
	pos    string
	trace  bool
	debug  bool
}

func (p pagingQueryExecutor) method() string {
	return p.name
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
		method: p.name,
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
		name:    f.name,
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

func (f *fetchQueryExecutor) execute() (sql.Result, error) {
	var result sql.Result
	err := withTx(f.logger, f.pos, f.trace, f.debug, f.ctx, f.conn, func(tx *connTx) (err error) {
		// TODO Complete cursor ID
		cursor := fmt.Sprintf("curosr_%s", tx.TraceId())
		r := newRaw(false, fmt.Sprintf("declare %s cursor for %s", cursor, f.raw.SQL), nil)
		r.mergeVars(f.raw.Vars)
		err = f.exec(tx, r, nil, nil)
		if err != nil {
			return
		}
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
		err = parser.AddError(err, f.exec(tx, newRaw(false, fmt.Sprintf("close %s", cursor), nil), nil, nil))
		if err != nil {
			return
		}
		return
	})
	return result, err
}
