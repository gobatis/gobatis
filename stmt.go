package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"reflect"
	"strconv"
	"strings"
)

type Stmt struct {
	ctx     context.Context
	in      []reflect.Value
	sql     string
	vars    []interface{}
	query   bool
	dynamic bool
	strict  bool
	conn    conn
	rows    *sql.Rows
	result  sql.Result
	id      string
}

func (p Stmt) close() {
	if p.conn != nil {
		_ = p.conn.Close()
	}
}

func (p Stmt) ID() string {
	return p.id
}

func (p Stmt) RealSQL() string {
	s := p.sql
	for i, v := range p.vars {
		s = strings.Replace(s, fmt.Sprintf("$%d", i+1), p.realValue(v), 1)
		//s = strings.Replace(s, fmt.Sprintf("$%d", i+1), fmt.Sprintf("$%d[%v]", i+1, v), 1)
	}
	return s
}

func (p Stmt) realValue(v interface{}) string {
	vv := cast.IndirectToStringerOrError(v)
	switch s := vv.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	case bool:
		return fmt.Sprintf("%v", v)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(uint64(s), 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case []byte:
		return fmt.Sprintf("'%s'", v)
	case nil:
		return ""
	case fmt.Stringer:
		return fmt.Sprintf("'%s'", v)
	case error:
		return fmt.Sprintf("'%s'", v)
	default:
		return ""
	}
}

func (p *Stmt) concatSQL(s string) {
	if p.sql == "" {
		p.sql = s
	} else {
		p.sql += " " + s
	}
}
