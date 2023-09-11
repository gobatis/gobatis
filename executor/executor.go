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

func (d *Default) Result() sql.Result {
	//TODO implement me
	panic("implement me")
}

func (d *Default) Execute(logger Logger, trace, debug bool, affecting any, scan func(Scanner) error) (err error) {

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

	if scan != nil {
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

func (i *InsertBatch) Execute(logger Logger, trace, debug bool, affecting any, scan func(Scanner) error) (err error) {

	ibs := &insertBatchScanner{}

	//defer func() {
	//	for _, v := range ibs.rows {
	//		err = AddError(err, v.Close())
	//	}
	//}()

	conn := i.conn

	var tx *sql.Tx
	defer func() {
		if tx != nil {
			if err != nil {
				now := time.Now()
				err = AddError(err, tx.Rollback())
				logger.Trace(conn.TraceId(), trace, err, &SQLTrace{
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
		tx, err = conn.BeginTx(i.ctx, nil)
		if err != nil {
			return
		}
		conn = NewTx(tx, conn.TraceId())
	}

	for _, raw := range i.raws {
		d := NewDefault(conn, raw)
		d.clean = false
		err = d.Execute(logger, trace, debug, affecting, nil)
		if err != nil {
			return
		}
		ibs.rows = d.rows
		ibs.result = d.result
		err = scan(ibs)
		if err != nil {
			return
		}
		d.rows.Close()
	}

	//err = scan(ibs)
	//if err != nil {
	//	return
	//}

	now := time.Now()
	err = tx.Commit()
	if err != nil {
		return
	}

	logger.Trace(conn.TraceId(), trace, err, &SQLTrace{
		Trace:        trace,
		Debug:        debug,
		BeginAt:      now,
		RawSQL:       "commit",
		PlainSQL:     "commit",
		RowsAffected: 0,
	})

	return
}

type ParallelQuery struct {
	conn Conn
}

func (p *ParallelQuery) Execute(logger Logger, trace, debug bool, affecting any, scan func(Scanner) error) (err error) {

	return
}

func NewFetchQuery(conn Conn, raw *Raw, limit uint) *FetchQuery {
	return &FetchQuery{conn: conn, raw: raw, limit: limit}
}

type FetchQuery struct {
	conn  Conn
	raw   *Raw
	limit uint
}

func (f *FetchQuery) Execute(logger Logger, trace, debug bool, affecting any, scan func(Scanner) error) (err error) {
	return
}
