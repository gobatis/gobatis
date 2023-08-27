package batis

import (
	"fmt"

	"github.com/gobatis/gobatis/dialector"
)

type Queryer interface {
	executors(namer dialector.Namer, tag string) ([]*executor, error)
}

type Query struct {
	SQL    string
	Params map[string]any
	Scan   any
}

func (q Query) executors(namer dialector.Namer, tag string) ([]*executor, error) {

	if q.Scan == nil {
		return nil, fmt.Errorf("expect 1 scan dest; got nil")
	}

	var params []NameValue
	for k, v := range q.Params {
		params = append(params, NameValue{
			Name:  k,
			Value: v,
		})
	}
	e := &executor{}
	e.raw = q.SQL
	e.params = params
	e.dest = q.Scan

	return []*executor{e}, nil
}

var _ Queryer = (*Paging)(nil)

type Paging struct {
	Select string
	Count  string
	From   string
	Where  string
	Page   int64
	Limit  int64
	Params map[string]any
	Scan   []any
	elems  map[int][]Element
}

func (p Paging) executors(namer dialector.Namer, tag string) ([]*executor, error) {

	if p.Limit <= 0 {
		return nil, fmt.Errorf("invalid limit")
	}

	if l := len(p.Scan); l != 2 {
		return nil, fmt.Errorf("expect 2 scan dest; got: %d", l)
	}

	var params []NameValue
	for k, v := range p.Params {
		params = append(params, NameValue{
			Name:  k,
			Value: v,
		})
	}

	q := &executor{}
	q.raw = fmt.Sprintf("select %s from %s limit %d offset %d", p.Select, p.From, p.Limit, p.Limit*p.Page)
	q.params = params
	q.dest = p.Scan[0]

	c := &executor{}
	c.raw = fmt.Sprintf("select count(%s) from %s", p.Count, p.From)
	c.params = params
	c.dest = p.Scan[1]

	return []*executor{q, c}, nil
}
