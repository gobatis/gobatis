package gobatis

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

const (
	result_rows = iota + 1
	result_result
)

func newResult(typ int) *Result {
	return &Result{typ: typ}
}

type Result struct {
	typ    int
	rows   *sql.Rows
	result sql.Result
}

func (p *Result) Rows() *sql.Rows {
	return p.rows
}

func (p *Result) Result() sql.Result {
	return p.result
}

func (p *Result) Bind(pointer interface{}, column ...string) error {
	switch p.typ {
	case result_rows:
		return p.bindRows(pointer, "")
	default:
		return fmt.Errorf("no rows")
	}
}

func (p *Result) bindRows(pointer interface{}, column string) (err error) {
	times := 1
	index := 0
	columns, err := p.rows.Columns()
	if err != nil {
		return
	}
	values := newResultValues(columns)
	pointers := make([]interface{}, len(columns))
	for p.rows.Next() {
		if index >= times {
			break
		}
		index++
		for i, _ := range columns {
			pointers[i] = &values.values[i]
		}
		err = p.rows.Scan(pointers...)
		if err != nil {
			_ = p.rows.Close()
			return
		}
		fmt.Println(values.String())
	}
	return
}

func newResultValues(columns []string) *resultValues {
	r := new(resultValues)
	r.values = make([]interface{}, len(columns))
	r._map = map[string]int{}
	for i, v := range columns {
		r._map[v] = i
	}
	return r
}

type resultValues struct {
	values []interface{}
	_map   map[string]int
}

func (p *resultValues) bindPointer(pointer interface{}, column string) {

}

func (p *resultValues) String() string {
	r := map[string]interface{}{}
	for k, v := range p._map {
		r[k] = p.values[v]
	}
	d, _ := json.MarshalIndent(r, "", "\t")
	return string(d)
}
