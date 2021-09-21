package gobatis

import (
	"context"
	"strings"
)

type sentence struct {
	sql     string
	exprs   []string
	vars    []interface{}
	dynamic bool
	ctx     context.Context
	conn    conn
}

func (p *sentence) merge(s ...*sentence) {
	for _, v := range s {
		p.sql += " " + strings.TrimSpace(v.sql)
		if !p.dynamic && v.dynamic {
			p.dynamic = v.dynamic
		}
	}
}

func (p *sentence) printLog() {
	//p.logger.Errorf("[gobatis] [%s] exec statement: %s", p.fragment.id, s)
	//p.logger.Errorf("[gobatis] [%s] exec parameter: %s", p.fragment.id, printVars(vars))
	//p.logger.Errorf("[gobatis] [%s] prepare error: %v", p.fragment.id, err)
}
