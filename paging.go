package batis

type Query struct {
	SQL    string
	Params []NameValue
	Scan   any
}

func (q *Query) Queries() ([]executor, error) {
	return nil, nil
}

type Paging struct {
	Select string
	Count  string
	Common string
	Page   int64
	Limit  int64
	Params []NameValue
	Scan   []any
	elems  map[int][]Element
}

func init() {
	a := &Paging{
		Select: "*",
		Count:  "*",
		Common: `users where name age > #{age}`,
		Page:   0,
		Limit:  0,
		Params: []NameValue{
			{Name: "age", Value: 18},
		},
		Scan:  []any{nil, nil},
		elems: nil,
	}
	_ = a
}

func (p *Paging) Queries() ([]executor, error) {
	return nil, nil
}

func (p *Paging) Build() (executors []executor, err error) {
	
	return
}
