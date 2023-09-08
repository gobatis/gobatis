package executor

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type Executor interface {
	Execute(logger Logger, id string, trace, debug bool, affecting any, scan func(s *Scanner) error) (err error)
	Result() sql.Result
}

var (
	_ Executor = (*Default)(nil)
	_ Executor = (*InsertBatch)(nil)
	_ Executor = (*ParallelQuery)(nil)
	_ Executor = (*FetchQuery)(nil)
)

func NewDefault(conn Conn, raw *Raw) *Default {
	return &Default{conn: conn, raw: raw}
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
}

func (d *Default) Result() sql.Result {
	//TODO implement me
	panic("implement me")
}

func (d *Default) Execute(logger Logger, id string, trace, debug bool, affecting any, scan func(s *Scanner) error) (err error) {

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

	var s *Scanner

	defer func() {
		if d.rows != nil {
			err = AddError(err, d.rows.Close())
		}
		err = AddError(err, d.conn.Close())

		if s == nil {
			s = &Scanner{}
		}
		logger.Trace(trace, debug, beginAt, id, d.conn.IsTx(), d.sql, s.RowsAffected, err)
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
			s = &Scanner{
				rows:         d.rows,
				RowsAffected: 0,
				LastInsertId: 0,
			}
			err = scan(s)
		} else {
			rowsAffected, _ := d.result.RowsAffected()
			lastInsertId, _ := d.result.LastInsertId()
			s = &Scanner{
				rows:         nil,
				RowsAffected: rowsAffected,
				LastInsertId: lastInsertId,
			}
			err = scan(s)
		}
		if err != nil {
			return
		}
	}

	return
}

func NewInsertBatch(conn Conn, raw *Raw) *InsertBatch {
	return &InsertBatch{conn: conn, raw: raw}
}

type InsertBatch struct {
	conn Conn
	raw  *Raw
}

func (i *InsertBatch) Result() sql.Result {
	//TODO implement me
	panic("implement me")
}

func (i *InsertBatch) Execute(logger Logger, id string, trace, debug bool, affecting any, scan func(s *Scanner) error) (err error) {

	return
}

type ParallelQuery struct {
	conn Conn
}

func (p *ParallelQuery) Result() sql.Result {
	//TODO implement me
	panic("implement me")
}

func (p *ParallelQuery) Execute(logger Logger, id string, trace, debug bool, affecting any, scan func(s *Scanner) error) (err error) {
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

func (f *FetchQuery) Result() sql.Result {
	//TODO implement me
	panic("implement me")
}

func (f *FetchQuery) Execute(logger Logger, id string, trace, debug bool, affecting any, scan func(s *Scanner) error) (err error) {
	return
}
