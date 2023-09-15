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

func (q ParallelQuery) executor(namer dialector.Namer, tag string) (*executor.ParallelQuery, error) {
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
	raw := &executor.Raw{
		Ctx:    nil,
		Query:  true,
		SQL:    q.SQL,
		Params: nil,
	}
	for k, v := range q.Params {
		raw.Params = append(raw.Params, Param(k, v))
	}
	return &executor.ParallelQuery{Raw: raw, Dest: q.Scan}, nil
}

func PagingScan(items any, count any) []any {
	return []any{items, count}
}

type PagingQuery struct {
	Select string
	Count  string
	From   string
	Where  string
	Order  string
	Page   int64
	Limit  int64
	Params map[string]any
	Scan   []any
	elems  map[int][]Element
}

func (p PagingQuery) executors(namer dialector.Namer, tag string) ([]ParallelQuery, error) {
	
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
	w := ""
	if p.Where != "" {
		w = fmt.Sprintf(" where %s", p.Where)
	}
	o := ""
	if p.Order != "" {
		o = fmt.Sprintf(" order by %s", p.Order)
	}
	q := ParallelQuery{
		SQL:    fmt.Sprintf("select %s from %s%s%s limit %d offset %d", p.Select, p.From, w, o, p.Limit, p.Limit*p.Page),
		Params: p.Params,
		Scan:   p.Scan[0],
	}
	
	c := ParallelQuery{
		SQL:    fmt.Sprintf("select count(%s) from %s%s", p.Count, p.From, w),
		Params: p.Params,
		Scan:   p.Scan[1],
	}
	
	return []ParallelQuery{q, c}, nil
}

type FetchQuery struct {
	SQL    string
	Params map[string]any
	Batch  uint
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
