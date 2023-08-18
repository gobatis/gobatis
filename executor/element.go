package executor

type Element struct {
	Name   int
	SQL    string
	Params []Param
}

type Param struct {
	Name  string
	Value any
}
