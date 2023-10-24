package batis

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gobatis/gobatis/logger"
	"github.com/gobatis/gobatis/parser"
	"github.com/gobatis/gobatis/parser/xsql"
)

type executor interface {
	Execute(logger logger.Logger, pos string, trace, debug bool, affect any, scan func(s Scanner) error) (err error)
	Query() bool
}

var (
	_ executor = (*defaultExecutor)(nil)
	_ executor = (*insertBatchExecutor)(nil)
	_ executor = (*parallelQueryExecutor)(nil)
	_ executor = (*fetchQueryExecutor)(nil)
	_ any      = (*associateQueryExecutor)(nil)
)

func newDefaultExecutor(conn conn, raw *Raw) *defaultExecutor {
	return &defaultExecutor{conn: conn, raw: raw, clean: true}
}

type defaultExecutor struct {
	rows   *sql.Rows
	result sql.Result
	conn   conn
	raw    *Raw
	clean  bool
}

func (d *defaultExecutor) Query() bool {
	return d.raw.Query
}

func (d *defaultExecutor) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

	beginAt := time.Now()

	vars := map[string]any{}
	for _, v := range d.raw.Params {
		vars[v.Name] = v.Value
	}

	r, err := xsql.Parse(log.Explain, d.raw.SQL, vars)
	if err != nil {
		return
	}

	var s *scanner

	defer func() {
		if d.clean {
			if d.rows != nil {
				err = parser.AddError(err, d.rows.Close())
			}
			if !d.conn.IsTx() {
				err = parser.AddError(err, d.conn.Close())
			}
		}
		if s == nil {
			s = &scanner{}
		}
		plainSQL, e := xsql.Explain(log.Explain, d.raw.SQL, vars)
		if e != nil {
			plainSQL = fmt.Sprintf("explain sql error: %s", e)
		}
		log.Trace(pos, d.conn.TraceId(), d.conn.IsTx(), err, &logger.SQLTrace{
			Trace:        trace,
			Debug:        debug,
			BeginAt:      beginAt,
			RawSQL:       r.Statement(),
			PlainSQL:     plainSQL,
			RowsAffected: s.rowsAffected,
		})
	}()

	if d.raw.Query {
		d.rows, err = d.conn.QueryContext(d.raw.Ctx, r.Statement(), r.Vars()...)
		if err != nil {
			return
		}
	} else {
		d.result, err = d.conn.ExecContext(d.raw.Ctx, r.Statement(), r.Vars()...)
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
		err = d.scan(s, scan)
	} else {
		rowsAffected, _ := d.result.RowsAffected()
		lastInsertId, _ := d.result.LastInsertId()
		s = &scanner{
			rows:         nil,
			rowsAffected: rowsAffected,
			lastInsertId: lastInsertId,
		}
		err = d.scan(s, scan)
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

func (d *defaultExecutor) scan(s *scanner, f func(Scanner) error) error {
	if f == nil {
		return nil
	}
	return f(s)
}

func newInsertBatch(ctx context.Context, conn conn, raws []*Raw) *insertBatchExecutor {
	return &insertBatchExecutor{ctx: ctx, conn: conn, raws: raws}
}

type insertBatchExecutor struct {
	ctx  context.Context
	conn conn
	raws []*Raw
}

func (i *insertBatchExecutor) Query() bool {
	for _, v := range i.raws {
		if v.Query {
			return v.Query
		}
	}
	return false
}

func (i *insertBatchExecutor) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

	c := i.conn
	var tx *sql.Tx
	defer func() {
		// Indicate that the outside is a regular DB object,
		// not a transaction object.
		if !i.conn.IsTx() {
			err = parser.AddError(err, i.conn.Close())
		}
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = parser.AddError(err, tx.Rollback())
				log.Trace(pos, c.TraceId(), true, err, &logger.SQLTrace{
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

	if !c.IsTx() {
		now := time.Now()
		tx, err = c.BeginTx(i.ctx, nil)
		if err != nil {
			return
		}
		c = NewTx(tx, c.TraceId())
		log.Trace(pos, c.TraceId(), true, err, &logger.SQLTrace{
			Trace:        trace,
			Debug:        debug,
			BeginAt:      now,
			RawSQL:       "begin",
			PlainSQL:     "begin",
			RowsAffected: 0,
		})
	}

	ibs := &insertBatchScanner{}
	for _, raw := range i.raws {
		err = i.execute(c, raw, log, pos, trace, debug, ibs, func(s Scanner) error {
			return scan(s)
		})
		if err != nil {
			return
		}
	}

	if affect != nil {
		var ac *affectConstraint
		ac, err = newAffectConstraint(affect)
		if err != nil {
			return
		}
		err = ac.Check(int(ibs.rowsAffected))
		if err != nil {
			return
		}
	}

	now := time.Now()
	err = tx.Commit()
	if err != nil {
		return
	}

	log.Trace(pos, c.TraceId(), true, err, &logger.SQLTrace{
		Trace:        trace,
		Debug:        debug,
		BeginAt:      now,
		RawSQL:       "commit",
		PlainSQL:     "commit",
		RowsAffected: 0,
	})

	return
}

func (i *insertBatchExecutor) execute(conn conn, raw *Raw, logger logger.Logger, pos string, trace, debug bool, ibs *insertBatchScanner, scan func(Scanner) error) (err error) {
	d := newDefaultExecutor(conn, raw)
	d.clean = false
	err = d.Execute(logger, pos, trace, debug, nil, nil)
	if err != nil {
		return
	}
	defer func() {
		if d.rows != nil {
			err = parser.AddError(err, d.rows.Close())
		}
	}()
	if d.result != nil {
		lastInsertId, _ := d.result.LastInsertId()
		rowsAffected, _ := d.result.RowsAffected()
		ibs.lastInsertId = lastInsertId
		ibs.rowsAffected += rowsAffected
	}
	ibs.rows = d.rows
	err = scan(ibs)
	if err != nil {
		return
	}
	return
}

type parallelQueryExecutor struct {
	Conn    conn
	Raw     *Raw
	scanner func(s Scanner) error
}

func (p *parallelQueryExecutor) Query() bool {
	return p.Raw.Query
}

func (p *parallelQueryExecutor) Execute(logger logger.Logger, pos string, trace, debug bool, affect any, f func(Scanner) error) error {
	d := newDefaultExecutor(p.Conn, p.Raw)
	return d.Execute(logger, pos, trace, debug, affect, func(s Scanner) error {
		if f == nil {
			return nil
		}
		return f(s)
	})
}

func newFetchQuery(ctx context.Context, conn conn, raw *Raw, limit uint) *fetchQueryExecutor {
	return &fetchQueryExecutor{ctx: ctx, conn: conn, raw: raw, limit: limit}
}

type fetchQueryExecutor struct {
	ctx   context.Context
	conn  conn
	raw   *Raw
	limit uint
}

func (f *fetchQueryExecutor) Query() bool {
	return f.raw.Query
}

func (f *fetchQueryExecutor) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

	c := f.conn

	var tx *sql.Tx
	defer func() {
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = parser.AddError(err, tx.Rollback())
				log.Trace(pos, c.TraceId(), true, err, &logger.SQLTrace{
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

	if !c.IsTx() {
		tx, err = c.BeginTx(f.ctx, nil)
		if err != nil {
			return
		}
		c = NewTx(tx, c.TraceId())
	}

	cursor := fmt.Sprintf("curosr_%s", c.TraceId())

	d := newDefaultExecutor(c, &Raw{
		Ctx:    f.ctx,
		Query:  false,
		SQL:    fmt.Sprintf("declare %s cursor for %s", cursor, f.raw.SQL),
		Params: f.raw.Params,
	})
	err = d.Execute(log, pos, trace, debug, affect, nil)
	if err != nil {
		return
	}

	defer func() {
		d = newDefaultExecutor(c, &Raw{
			Ctx:    f.ctx,
			Query:  false,
			SQL:    fmt.Sprintf("close %s", cursor),
			Params: nil,
		})
		err = parser.AddError(err, d.Execute(log, pos, trace, debug, affect, nil))

		now := time.Now()
		err = parser.AddError(err, tx.Commit())

		log.Trace(pos, c.TraceId(), true, err, &logger.SQLTrace{
			Trace:        trace,
			Debug:        debug,
			BeginAt:      now,
			RawSQL:       "commit",
			PlainSQL:     "commit",
			RowsAffected: 0,
		})
	}()

	for {
		d = newDefaultExecutor(c, &Raw{
			Ctx:    f.ctx,
			Query:  true,
			SQL:    fmt.Sprintf("fetch forward %d from %s", f.limit, cursor),
			Params: nil,
		})
		var rowsAffected int64
		err = d.Execute(log, pos, trace, debug, affect, func(s Scanner) error {
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

func newAssociateQueryExecutor(conn conn, raw *Raw) *associateQueryExecutor {
	return &associateQueryExecutor{conn: conn, raw: raw}
}

type associateQueryExecutor struct {
	conn conn
	raw  *Raw
}

func (a associateQueryExecutor) Execute(logger logger.Logger, pos string, trace, debug bool, affect any, scan func(s AssociateScanner) error) (err error) {
	d := newDefaultExecutor(a.conn, a.raw)
	d.clean = false

	err = d.Execute(logger, pos, trace, debug, affect, nil)
	if err != nil {
		return
	}
	defer func() {
		_ = d.rows.Close()
	}()

	if scan == nil {
		return
	}

	err = scan(&associateScanner{
		rows: d.rows,
	})
	if err != nil {
		return
	}

	return
}

func (a associateQueryExecutor) Query() bool {
	return a.raw.Query
}
