package batis

type ParallelQuery struct {
	SQL    string
	Params map[string]any
	Scan   func(s Scanner) error
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
