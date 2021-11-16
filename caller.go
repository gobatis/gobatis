package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/dtd"
	"reflect"
	"time"
)

type conn interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}

type caller struct {
	mt       reflect.Type
	fragment *fragment
	logger   Logger
	result   []reflect.Value
}

func (p *caller) call(in []reflect.Value) *caller {
	
	start := time.Now()
	defer func() {
		p.logger.Debugf("[gobatis] [%s] cost: %s", p.fragment.id, time.Since(start))
	}()
	
	var err error
	defer func() {
		p.export(err)
	}()
	
	switch p.fragment.node.Name {
	case dtd.SELECT, dtd.INSERT, dtd.DELETE, dtd.UPDATE:
		err = p.exec(in)
	case dtd.SAVE:
		err = p.save(in)
	case dtd.QUERY:
		err = p.query(in)
	default:
		err = fmt.Errorf("unsupported fragment %s<%s>", p.fragment.node.Name, p.fragment.id)
	}
	return p
}

func (p *caller) exec(in []reflect.Value) error {
	s, err := p.fragment.buildStmt(in)
	if err != nil {
		return err
	}
	if p.fragment.stmt {
		p.result[0] = reflect.ValueOf(s)
		return nil
	}
	defer func() {
		s.close()
	}()
	err = p.run(s)
	if err != nil {
		return err
	}
	if s.query && len(p.result) > 0 {
		return p.scanRows(p.fragment.rt, p.fragment.out, s.rows, p.result...)
	}
	if len(p.result) > 0 {
		return p.scanResult(s.result, p.result)
	}
	return nil
}

func (p *caller) query(in []reflect.Value) (err error) {
	
	//TODO 检查 Out 类型
	ss, err := p.fragment.buildQuery(in)
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
		err = p.scanRows(result_result, []*param{{name: "count"}}, ss[0].rows, p.result[0])
		if err != nil {
			return
		}
	}
	
	if ss[1] != nil {
		err = p.run(ss[1])
		if err != nil {
			return
		}
		err = p.scanRows(result_none, nil, ss[1].rows, p.result[1])
		if err != nil {
			return
		}
	}
	
	return
}

func (p *caller) save(in []reflect.Value) (err error) {
	
	s, err := p.fragment.buildSave(in)
	if err != nil {
		return
	}
	
	err = p.run(s)
	if err != nil {
		return
	}
	
	return p.scanResult(s.result, p.result)
}

func (p caller) run(s *Stmt) (err error) {
	
	if s.conn == nil {
		s.conn, err = p.fragment.db.Conn(s.ctx)
		if err != nil {
			return
		}
	}
	
	defer func() {
		if err != nil {
			p.logger.Errorf("[gobatis] [%s] exec error\n[sql]: %s\n[vars]: %v\n[detail]: %s",
				p.fragment.id, s.sql, s.vars, err)
		} else {
			if p.logger.Level() == DebugLevel {
				p.logger.Debugf("[gobatis] [%s] exec '%s'", p.fragment.id, s.sql)
			}
		}
	}()
	
	if s.query {
		s.rows, err = s.conn.QueryContext(s.ctx, s.sql, s.vars...)
		if err != nil {
			return
		}
		return
	}
	
	s.result, err = s.conn.ExecContext(s.ctx, s.sql, s.vars...)
	if err != nil {
		return
	}
	return
}

func (p *caller) export(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			if p.fragment.must {
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

func (p caller) scanRows(rt int, params []*param, rows *sql.Rows, values ...reflect.Value) (err error) {
	defer func() {
		if _err := rows.Close(); _err != nil {
			p.logger.Errorf("[gobatis] [%s] close rows error: %s", p.fragment.id, _err)
		}
	}()
	scanner := queryScanner{
		rows:    rows,
		tag:     p.fragment.scanTag(),
		scanner: p.fragment.engine.scannerFactory(),
	}
	err = scanner.setSelected(rt, params, values)
	if err != nil {
		return err
	}
	return scanner.scan()
}

func (p *caller) scanResult(res sql.Result, values []reflect.Value) error {
	// ignore RowsAffected to support database that not support
	affected, _ := res.RowsAffected()
	if p.fragment.must && affected != 1 {
		return fmt.Errorf("expect affect 1 row, got %d", affected)
	}
	scanner := execScanner{affected: affected, values: values}
	return scanner.scan()
}
