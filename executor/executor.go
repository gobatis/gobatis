package executor

import (
	"fmt"
	batis "github.com/gobatis/gobatis"
)

type Executor struct {
	SQL    string
	Params []batis.NameValue
}

func (e Executor) Merge(s Executor) {
	e.SQL = fmt.Sprintf(" %s", s.SQL)
	e.Params = append(e.Params, s.Params...)
}
