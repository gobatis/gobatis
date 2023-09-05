package executor

import "context"

type Param struct {
	Name  string
	Value any
}

type Raw struct {
	Ctx    context.Context
	Query  bool
	SQL    string
	Params []Param
}
