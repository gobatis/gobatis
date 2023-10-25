package batis

//type ParallelQueryer interface {
//	executors(namer dialector.Namer, tag string) ([]executor.Executor, error)
//}

type ParallelQuery struct {
	SQL    string
	Params map[string]any
	Scan   func(s Scanner) error
}

func (q ParallelQuery) raw() *raw {
	r := newRaw(true, q.SQL, nil)
	r.mergeVars(q.Params)
	return r
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
	Scan   func(scanner PagingScanner) error
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
