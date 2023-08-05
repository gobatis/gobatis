package executor

type Element struct {
	Name   int
	SQL    string
	Params []NameValue
}

type NameValue struct {
	Name  string
	Value any
}
