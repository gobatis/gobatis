package batis

type Paging struct {
	Select string
	Count  string
	Common string
	Page   int64
	Limit  int64
	Params []NameValue
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
		elems: nil,
	}
	_ = a
}

func (b *Paging) Build() (executors []executor, err error) {
	
	return
}
