package batis

import "context"

type NameValue struct {
	Name  string
	Value any
}

type Raw struct {
	Ctx    context.Context
	Query  bool
	SQL    string
	Params []NameValue
}
