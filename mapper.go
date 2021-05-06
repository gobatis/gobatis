package gobatis

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"sync"
)

//type Params = map[string]interface{}

type fragmentManager struct {
	mu        sync.RWMutex
	fragments map[string]*fragment
}

func newMethodManager() *fragmentManager {
	return &fragmentManager{}
}

func (p *fragmentManager) add(m *fragment) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.fragments == nil {
		p.fragments = map[string]*fragment{}
	}
	_, ok := p.fragments[m.id]
	if ok {
		return fmt.Errorf("duplicated fragment '%s'", m.id)
	}
	p.fragments[m.id] = m
	return nil
}

func (p *fragmentManager) replace(m *fragment) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.fragments == nil {
		p.fragments = map[string]*fragment{}
	}
	_, ok := p.fragments[m.id]
	if !ok {
		return fmt.Errorf("fragment '%s' not exist", m.id)
	}
	p.fragments[m.id] = m
	return nil
}

func (p *fragmentManager) get(id string) (m *fragment, ok bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.fragments == nil {
		return
	}
	m, ok = p.fragments[id]
	return
}

type param struct {
	name string
	kind reflect.Kind
}

type execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

func newFragment(id string, in, out []param, statement *xmlNode) *fragment {
	return &fragment{id: id, in: in, out: out, statement: statement}
}

type fragment struct {
	db        *DB
	id        string
	statement *xmlNode
	cacheable bool
	sql       string
	in        []param
	out       []param
}

func (p *fragment) exec(in ...reflect.Value) (out []reflect.Value) {

	return
}

func (p *fragment) query(in ...reflect.Value) (out []reflect.Value) {
	return
}

func (p *fragment) removeElem(a []reflect.Value, i int) []reflect.Value {
	return append(a[:i], a[i+1:]...)
}

func (p *fragment) execer(in ...reflect.Value) (execer, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(execer)).Elem()
		for i, v := range in {
			if v.Type().Implements(t) {
				return v.Interface().(execer), i
			}
		}
	}
	return nil, -1
}

func (p *fragment) queryer(in ...reflect.Value) (queryer, int) {
	if len(in) > 0 {
		t := reflect.TypeOf(new(queryer)).Elem()
		for i, v := range in {
			if v.Type().Implements(t) {
				return v.Interface().(queryer), i
			}
		}
	}
	return nil, -1
}
