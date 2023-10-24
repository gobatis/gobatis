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

type Executor interface {
	Execute(logger logger.Logger, pos string, trace, debug bool, affect any, scan func(s Scanner) error) (err error)
	Query() bool
}

var (
	_ Executor = (*DefaultExecutor)(nil)
	_ Executor = (*InsertBatchExecutor)(nil)
	_ Executor = (*ParallelQueryExecutor)(nil)
	_ Executor = (*FetchQueryExecutor)(nil)
	_ Executor = (*AssociateQueryExecutor)(nil)
)

func NewDefault(conn Conn2, raw *Raw) *DefaultExecutor {
	return &DefaultExecutor{conn: conn, raw: raw, clean: true}
}

type DefaultExecutor struct {
	rows   *sql.Rows
	result sql.Result
	conn   Conn2
	raw    *Raw
	clean  bool
}

func (d *DefaultExecutor) Query() bool {
	return d.raw.Query
}

func (d *DefaultExecutor) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

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

func (d *DefaultExecutor) scan(s *scanner, f func(Scanner) error) error {
	if f == nil {
		return nil
	}
	return f(s)
}

func NewInsertBatch(ctx context.Context, conn Conn2, raws []*Raw) *InsertBatchExecutor {
	return &InsertBatchExecutor{ctx: ctx, conn: conn, raws: raws}
}

type InsertBatchExecutor struct {
	ctx  context.Context
	conn Conn2
	raws []*Raw
}

func (i *InsertBatchExecutor) Query() bool {
	for _, v := range i.raws {
		if v.Query {
			return v.Query
		}
	}
	return false
}

func (i *InsertBatchExecutor) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

	conn := i.conn
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
				log.Trace(pos, conn.TraceId(), true, err, &logger.SQLTrace{
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
		log.Trace(pos, conn.TraceId(), true, err, &logger.SQLTrace{
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
		err = i.execute(conn, raw, log, pos, trace, debug, ibs, func(s Scanner) error {
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

	log.Trace(pos, conn.TraceId(), true, err, &logger.SQLTrace{
		Trace:        trace,
		Debug:        debug,
		BeginAt:      now,
		RawSQL:       "commit",
		PlainSQL:     "commit",
		RowsAffected: 0,
	})

	return
}

func (i *InsertBatchExecutor) execute(conn Conn2, raw *Raw, logger logger.Logger, pos string, trace, debug bool, ibs *insertBatchScanner, scan func(Scanner) error) (err error) {
	d := NewDefault(conn, raw)
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

type ParallelQueryExecutor struct {
	Conn Conn2
	Raw  *Raw
	Dest any
}

func (p *ParallelQueryExecutor) Query() bool {
	return p.Raw.Query
}

func (p *ParallelQueryExecutor) Execute(logger logger.Logger, pos string, trace, debug bool, affect any, _ func(Scanner) error) error {
	d := NewDefault(p.Conn, p.Raw)
	return d.Execute(logger, pos, trace, debug, affect, func(s Scanner) error {
		return s.Scan(p.Dest)
	})
}

func NewFetchQuery(ctx context.Context, conn Conn2, raw *Raw, limit uint) *FetchQueryExecutor {
	return &FetchQueryExecutor{ctx: ctx, conn: conn, raw: raw, limit: limit}
}

type FetchQueryExecutor struct {
	ctx   context.Context
	conn  Conn2
	raw   *Raw
	limit uint
}

func (f *FetchQueryExecutor) Query() bool {
	return f.raw.Query
}

func (f *FetchQueryExecutor) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

	conn := f.conn

	var tx *sql.Tx
	defer func() {
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = parser.AddError(err, tx.Rollback())
				log.Trace(pos, conn.TraceId(), true, err, &logger.SQLTrace{
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
	err = d.Execute(log, pos, trace, debug, affect, nil)
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
		err = parser.AddError(err, d.Execute(log, pos, trace, debug, affect, nil))

		now := time.Now()
		err = parser.AddError(err, tx.Commit())

		log.Trace(pos, conn.TraceId(), true, err, &logger.SQLTrace{
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

func NewAssociateQuery(conn Conn2, raw *Raw, dest any, bindingPath string, mappingPath string) *AssociateQueryExecutor {
	return &AssociateQueryExecutor{conn: conn, raw: raw, dest: dest, bindingPath: bindingPath, mappingPath: mappingPath}
}

type AssociateQueryExecutor struct {
	conn        Conn2
	raw         *Raw
	dest        any
	bindingPath string
	mappingPath string
}

func (a AssociateQueryExecutor) Execute(logger logger.Logger, pos string, trace, debug bool, affect any, _ func(s Scanner) error) (err error) {
	d := NewDefault(a.conn, a.raw)
	d.clean = false

	err = d.Execute(logger, pos, trace, debug, affect, nil)
	if err != nil {
		return
	}

	var abs = &associateScanner{
		rows:         d.rows,
		rowsAffected: 0,
		lastInsertId: 0,
		bindingPaths: []*associateBindingPath{{
			column: "product_name",
			path:   "$.Name",
		}},
		mappingPath: a.mappingPath,
	}
	err = abs.Scan(a.dest)

	return
}

func (a AssociateQueryExecutor) Query() bool {
	return a.raw.Query
}
