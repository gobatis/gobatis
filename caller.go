package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dtd"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type conn interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}

type stmt struct {
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
}

func (p stmt) close() {
	if p.conn != nil {
		_ = p.conn.Close()
	}
}

func (p stmt) realSQL() string {
	s := p.sql
	for i, v := range p.vars {
		s = strings.Replace(s, fmt.Sprintf("$%d", i+1), p.realValue(v), 1)
		//s = strings.Replace(s, fmt.Sprintf("$%d", i+1), fmt.Sprintf("$%d[%v]", i+1, v), 1)
	}
	return s
}

func (p stmt) realValue(v interface{}) string {
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

func (p *stmt) concatSQL(s string) {
	if p.sql == "" {
		p.sql = s
	} else {
		p.sql += " " + s
	}
}

type caller struct {
	mt     reflect.Type
	method *fragment
	logger Logger
	result []reflect.Value
}

func (p *caller) call(in []reflect.Value) *caller {
	
	start := time.Now()
	defer func() {
		p.logger.Debugf("[gobatis] [%s] cost: %s", p.method.id, time.Since(start))
	}()
	
	var err error
	defer func() {
		p.convertResult(err)
	}()
	
	switch p.method.node.Name {
	case dtd.SELECT, dtd.INSERT, dtd.DELETE, dtd.UPDATE:
		err = p.exec(in)
	case dtd.SAVE:
		err = p.save(in)
	case dtd.QUERY:
		err = p.query(in)
	default:
		throw(p.method.node.File, p.method.node.ctx, callerErr).
			format("unsupported call fragment '%s'", p.method.node.Name)
	}
	return p
}

func (p *caller) exec(in []reflect.Value) error {
	s, err := p.method.buildSegment(in)
	if err != nil {
		return err
	}
	defer func() {
		s.close()
	}()
	err = p.run(s)
	if err != nil {
		return err
	}
	if s.query {
		return p.parseQueryResult(p.method.rt, p.method.out, s.rows, p.result...)
	}
	return p.parseExecResult(s.result, p.result)
}

func (p *caller) query(in []reflect.Value) (err error) {
	
	//TODO 检查 Out 类型
	
	ss, err := p.method.buildQuery(in)
	if err != nil {
		return
	}
	
	defer func() {
		if ss[0] != nil {
			ss[0].close()
		}
		if ss[1] != nil {
			ss[1].close()
		}
	}()
	
	if ss[0] != nil {
		err = p.run(ss[0])
		if err != nil {
			return
		}
		err = p.parseQueryResult(result_result, []*param{{name: "count"}}, ss[0].rows, p.result[0])
		if err != nil {
			return
		}
	}
	
	if ss[1] != nil {
		err = p.run(ss[1])
		if err != nil {
			return
		}
		err = p.parseQueryResult(result_none, nil, ss[1].rows, p.result[1])
		if err != nil {
			return
		}
	}
	
	return
}

func (p *caller) save(in []reflect.Value) (err error) {
	
	return
}

func (p caller) run(s *stmt) (err error) {
	
	if s.conn == nil {
		s.conn, err = p.method.db.Conn(s.ctx)
		if err != nil {
			return
		}
	}
	
	defer func() {
		if err != nil {
			p.logger.Errorf("[gobatis] [%s] exec error\n[sql]: %s\n[vars]: %v\n[detail]: %s",
				p.method.id, s.sql, s.vars, err)
		} else {
			if p.logger.Level() == DebugLevel {
				p.logger.Debugf("[gobatis] [%s] exec '%s'", p.method.id, s.sql)
			}
		}
	}()
	
	if s.query {
		var rows *sql.Rows
		rows, err = s.conn.QueryContext(s.ctx, s.sql, s.vars...)
		if err != nil {
			return
		}
		s.rows = rows
		return
	}
	
	var r sql.Result
	r, err = s.conn.ExecContext(s.ctx, s.sql, s.vars...)
	if err != nil {
		return
	}
	s.result = r
	return
}

func (p *caller) convertResult(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			if p.method.must {
				p.result = append(p.result, reflect.ValueOf(err))
			} else {
				p.result = append(p.result, reflect.Zero(errorType))
			}
		} else {
			p.result = append(p.result, reflect.ValueOf(err))
		}
	} else {
		p.result = append(p.result, reflect.Zero(errorType))
	}
	if p.mt != nil {
		for i := 0; i < p.mt.NumOut()-1; i++ {
			if p.mt.Out(i).Kind() == reflect.Ptr {
				if err == sql.ErrNoRows {
					p.result[i] = reflect.Zero(p.result[i].Type())
				}
			} else {
				p.result[i] = p.result[i].Elem()
			}
		}
	}
}

func (p *caller) parseExecResult(res sql.Result, values []reflect.Value) error {
	// ignore RowsAffected to support database that not support
	affected, _ := res.RowsAffected()
	if p.method.must && affected != 1 {
		return fmt.Errorf("expect affect 1 row, got %d", affected)
	}
	return (&execResult{affected: affected, values: values}).scan()
}

func (p caller) parseQueryResult(rt int, params []*param, rows *sql.Rows, values ...reflect.Value) (err error) {
	defer func() {
		if _err := rows.Close(); _err != nil {
			p.logger.Errorf("[gobatis] [%s] close rows error: %s", p.method.id, _err)
		}
	}()
	res := queryResult{rows: rows, tag: p.method.tag()}
	err = res.setSelected(rt, params, values)
	if err != nil {
		return err
	}
	return res.scan()
}
