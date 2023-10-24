package batis

import (
	"fmt"

	"github.com/gobatis/gobatis/dialector"
)

//type ParallelQueryer interface {
//	executors(namer dialector.Namer, tag string) ([]executor.Executor, error)
//}

type ParallelQuery struct {
	SQL    string
	Params map[string]any
	Scan   func(s Scanner) error
}

func (q ParallelQuery) executor(namer dialector.Namer, tag string) (*parallelQueryExecutor, error) {
	if q.Scan == nil {
		return nil, NoScanDestErr
	}
	var params []NameValue
	for k, v := range q.Params {
		params = append(params, NameValue{
			Name:  k,
			Value: v,
		})
	}
	raw := &Raw{
		Ctx:    nil,
		Query:  true,
		SQL:    q.SQL,
		Params: nil,
	}
	for k, v := range q.Params {
		raw.Params = append(raw.Params, Param(k, v))
	}
	return &parallelQueryExecutor{Raw: raw}, nil
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
	elems  map[int][]Element
	Scan   func(scanner PagingScanner) error
}

func (p PagingQuery) GetCount() int64 {
	return 0
}

func (p PagingQuery) executors(namer dialector.Namer, tag string) ([]ParallelQuery, *pagingScanner, error) {

	if p.Limit <= 0 {
		return nil, nil, InvalidLimitErr
	}

	//if l := len(p.Scan); l != 2 {
	//	return nil, fmt.Errorf("%w; got: %d", InvalidPagingScanDestErr, l)
	//}

	var params []NameValue
	for k, v := range p.Params {
		params = append(params, NameValue{
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

	s := &pagingScanner{}

	q := ParallelQuery{
		SQL:    fmt.Sprintf("select %s from %s%s%s limit %d offset %d", p.Select, p.From, w, o, p.Limit, p.Limit*p.Page),
		Params: p.Params,
		Scan: func(scanner Scanner) error {
			s.listScanner = scanner
			return nil
		},
	}

	c := ParallelQuery{
		SQL:    fmt.Sprintf("select count(%s) from %s%s", p.Count, p.From, w),
		Params: p.Params,
		Scan: func(scanner Scanner) error {
			s.countScanner = scanner
			return nil
		},
	}

	return []ParallelQuery{q, c}, s, nil
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
	Scan   func(scanner AssociateScanner) error
}
