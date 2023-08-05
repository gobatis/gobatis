package executor

import batis "github.com/gobatis/gobatis"

type Element struct {
	Name   string
	SQL    string
	Params []batis.NameValue
}
