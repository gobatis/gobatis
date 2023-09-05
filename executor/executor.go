package executor

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Executor interface {
	Execute(scan func(s *Scanner) error) (err error)
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

func (d *Default) Execute(scan func(s *Scanner) error) (err error) {

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

	defer func() {
		if d.rows != nil {
			err = AddError(err, d.rows.Close())
		}
		err = AddError(err, d.conn.Close())
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
		err = scan(&Scanner{
			rows:   d.rows,
			result: d.result,
		})
		if err != nil {
			return
		}
	}

	return
}

type InsertBatch struct {
}

func (i InsertBatch) Result() sql.Result {
	//TODO implement me
	panic("implement me")
}

func (i InsertBatch) Execute(scan func(s *Scanner) error) (err error) {

	return
}

type ParallelQuery struct {
}

func (p ParallelQuery) Result() sql.Result {
	//TODO implement me
	panic("implement me")
}

func (p ParallelQuery) Execute(scan func(s *Scanner) error) (err error) {
	return
}

type FetchQuery struct {
}

func (f FetchQuery) Result() sql.Result {
	//TODO implement me
	panic("implement me")
}

func (f FetchQuery) Execute(scan func(s *Scanner) error) (err error) {
	return
}
