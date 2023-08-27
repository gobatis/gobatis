package batis

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"sync/atomic"
)

type conn interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Close() error
}

type tx struct {
	*sql.Tx
}

func (tx) Close() error {
	return nil
}

type executor struct {
	query    bool
	params   []NameValue
	conn     conn
	rows     *sql.Rows
	result   *sql.Result
	must     bool
	raw      string
	executed atomic.Bool
	tracer   *tracer
}

func (e *executor) exec(dest any) (rowsAffected *int64, lastedInsertId *int64, err error) {

	defer func() {
		//e.log(err)
	}()

	if e.executed.Swap(true) {
		err = fmt.Errorf("db execution was repeated")
		return
	}

	var _params []*param
	var _vars []reflect.Value
	for _, v := range e.params {
		_params = append(_params, &param{
			name:  v.Name,
			_type: reflect.TypeOf(v.Value).Name(),
		})
		_vars = append(_vars, reflect.ValueOf(v.Value))
	}

	var node *xmlNode
	node, err = parseSQL("test.file", fmt.Sprintf("<sql>%s</sql>", e.raw))
	if err != nil {
		return
	}

	frag := &fragment{node: node, in: _params}

	raw, exprs, vars, dynamic, err := frag.parseStatement(_vars...)
	if err != nil {
		return
	}
	_ = exprs
	_ = dynamic

	var rows *sql.Rows
	var result sql.Result

	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	if e.query {
		rows, err = e.conn.QueryContext(context.Background(), raw, vars...)
		if err != nil {
			return
		}
		if dest != nil {
			err = e.scan(rows, dest)
			if err != nil {
				return
			}
		}
		return
	} else {
		result, err = e.conn.ExecContext(context.Background(), raw, vars...)
		if err != nil {
			return
		}
		affected, _err := result.RowsAffected()
		if _err == nil {
			rowsAffected = &affected
		}
		id, _err := result.LastInsertId()
		if _err == nil {
			lastedInsertId = &id
		}
		return
	}
}

func (e *executor) scan(rows *sql.Rows, ptr any) (err error) {
	if ptr == nil {
		return
	}
	qr := queryResult{
		rows: rows,
	}
	err = qr.scan(ptr)
	if err != nil {
		err = fmt.Errorf("scan rows error: %w", err)
		return
	}

	return
}
