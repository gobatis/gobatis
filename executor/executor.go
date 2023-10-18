package executor

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gobatis/gobatis/logger"
	"github.com/gobatis/gobatis/parser/commons"
	"github.com/gobatis/gobatis/parser/xsql"
)

type Executor interface {
	Execute(logger logger.Logger, pos string, trace, debug bool, affect any, scan func(s Scanner) error) (err error)
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
	//fragment *fragment
	//exprs    []string
	//vars     []any
	//dynamic  bool
	//sql      string
	rows   *sql.Rows
	result sql.Result
	conn   Conn
	raw    *Raw
	clean  bool
}

func (d *Default) Query() bool {
	return d.raw.Query
}

func (d *Default) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

	beginAt := time.Now()

	vars := map[string]any{}
	for _, v := range d.raw.Params {
		vars[v.Name] = v.Value
	}

	r, err := xsql.Parse(d.raw.SQL, vars)
	if err != nil {
		return
	}

	var s *scanner

	defer func() {
		if d.clean {
			if d.rows != nil {
				err = commons.AddError(err, d.rows.Close())
			}
			if !d.conn.IsTx() {
				err = commons.AddError(err, d.conn.Close())
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
			RawSQL:       r.SQL(),
			PlainSQL:     plainSQL,
			RowsAffected: s.rowsAffected,
		})
	}()

	if d.raw.Query {
		d.rows, err = d.conn.QueryContext(d.raw.Ctx, r.SQL(), r.Vars()...)
		if err != nil {
			return
		}
	} else {
		d.result, err = d.conn.ExecContext(d.raw.Ctx, r.SQL(), r.Vars()...)
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

func (d *Default) scan(s *scanner, f func(Scanner) error) error {
	if f == nil {
		return nil
	}
	return f(s)
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

func (i *InsertBatch) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

	conn := i.conn
	var tx *sql.Tx
	defer func() {
		// Indicate that the outside is a regular DB object,
		// not a transaction object.
		if !i.conn.IsTx() {
			err = commons.AddError(err, i.conn.Close())
		}
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = commons.AddError(err, tx.Rollback())
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

func (i *InsertBatch) execute(conn Conn, raw *Raw, logger logger.Logger, pos string, trace, debug bool, ibs *insertBatchScanner, scan func(Scanner) error) (err error) {
	d := NewDefault(conn, raw)
	d.clean = false
	err = d.Execute(logger, pos, trace, debug, nil, nil)
	if err != nil {
		return
	}
	defer func() {
		if d.rows != nil {
			err = commons.AddError(err, d.rows.Close())
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

type ParallelQuery struct {
	Conn Conn
	Raw  *Raw
	Dest any
}

func (p *ParallelQuery) Query() bool {
	return p.Raw.Query
}

func (p *ParallelQuery) Execute(logger logger.Logger, pos string, trace, debug bool, affect any, _ func(Scanner) error) error {
	d := NewDefault(p.Conn, p.Raw)
	return d.Execute(logger, pos, trace, debug, affect, func(s Scanner) error {
		return s.Scan(p.Dest)
	})
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

func (f *FetchQuery) Execute(log logger.Logger, pos string, trace, debug bool, affect any, scan func(Scanner) error) (err error) {

	conn := f.conn

	var tx *sql.Tx
	defer func() {
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = commons.AddError(err, tx.Rollback())
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
		err = commons.AddError(err, d.Execute(log, pos, trace, debug, affect, nil))

		now := time.Now()
		err = commons.AddError(err, tx.Commit())

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
