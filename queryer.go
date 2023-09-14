package batis

import (
	"fmt"

	"github.com/gobatis/gobatis/dialector"
	"github.com/gobatis/gobatis/executor"
)

//type ParallelQueryer interface {
//	executors(namer dialector.Namer, tag string) ([]executor.Executor, error)
//}

type ParallelQuery struct {
	SQL    string
	Params map[string]any
	Scan   any
}

func (q ParallelQuery) executors(namer dialector.Namer, tag string) ([]executor.Executor, error) {

	if q.Scan == nil {
		return nil, fmt.Errorf("expect 1 scan dest; got nil")
	}

	var params []executor.Param
	for k, v := range q.Params {
		params = append(params, executor.Param{
			Name:  k,
			Value: v,
		})
	}
	//e := &executor{}
	//e.raw = q.SQL
	//e.params = params
	//e.dest = q.Scan
	//[]*executor{e}, nil
	return nil, nil
}

type PagingQuery struct {
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

func (p PagingQuery) executors(namer dialector.Namer, tag string) ([]executor.Executor, error) {

	if p.Limit <= 0 {
		return nil, fmt.Errorf("invalid limit")
	}

	if l := len(p.Scan); l != 2 {
		return nil, fmt.Errorf("expect 2 scan dest; got: %d", l)
	}

	var params []executor.Param
	for k, v := range p.Params {
		params = append(params, executor.Param{
			Name:  k,
			Value: v,
		})
	}

	//q := &executor{}
	//q.raw = fmt.Sprintf("select %s from %s limit %d offset %d", p.Select, p.From, p.Limit, p.Limit*p.Page)
	//q.params = params
	//q.dest = p.Scan[0]
	//
	//c := &executor{}
	//c.raw = fmt.Sprintf("select count(%s) from %s", p.Count, p.From)
	//c.params = params
	//c.dest = p.Scan[1]
	//[]*executor{q, c}, nil
	return nil, nil
}

type FetchQuery struct {
	SQL    string
	Params map[string]any
	Limit  uint
	Scan   func(scanner Scanner) error
}

type AssociateQuery struct {
	SQL    string
	Params map[string]any
	Link   any
}

func AssociateLink(dest any, condition, inject string) any {
	return nil
}
