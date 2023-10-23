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
		return nil, NoScanDestErr
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

func (p PagingQuery) GetCount() int64 {
	return 0
}

func (p PagingQuery) executors(namer dialector.Namer, tag string) ([]ParallelQuery, error) {

	if p.Limit <= 0 {
		return nil, InvalidLimitErr
	}

	if l := len(p.Scan); l != 2 {
		return nil, fmt.Errorf("%w; got: %d", InvalidPagingScanDestErr, l)
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
	//Associate associateDest
	Associate func() (any, string, string)
	Scan      func() (any, string, string, []string)
}

func Scan(dest any, options ...func(c *scanOptions)) func() (any, string, string, []string) {
	return func() (any, string, string, []string) {
		return dest, "", "", nil
	}
}

type scanOptions struct {
}

func BindingPath(v string) func(c *scanOptions) {
	return func(c *scanOptions) {

	}
}

func MappingPath(v string) func(c *scanOptions) {
	return func(c *scanOptions) {

	}
}

func Ignore(fields ...string) func(c *scanOptions) {
	return func(c *scanOptions) {

	}
}

//type associateDest struct {
//	dest        any
//	bindingPath string
//	mappingPath string
//}

//func Associate(dest any, bindingPath, mappingPath string) func() (any, string, string) {
//	return func() (any, string, string) {
//		return dest, bindingPath, mappingPath
//	}
//}

func Associate(bindingPath, mappingPath string) func(c *scanOptions) {
	return func(c *scanOptions) {

	}
}
