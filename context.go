package batis

import (
	"context"
	"github.com/gobatis/gobatis/executor"
	"time"
)

func Background() Context {
	return Context{}
}

func WithContext(ctx context.Context) Context {
	return Context{}
}

func WithDebug() Context {
	return Context{debug: true}
}

func WithMust() Context {
	return Context{must: true}
}

func WithTx(tx *DB) Context {
	return Context{tx: tx}
}

var _ context.Context = (*Context)(nil)

type Context struct {
	dirty   bool
	ctx     context.Context
	debug   bool
	must    bool
	cost    time.Duration
	sql     string
	params  []*executor.NameValue
	tx      *DB
	error   error
	exact   bool
	loose   bool
	traceId string
	analyze bool
	mute    bool
}

func (c Context) Copy() Context {
	return c
}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	deadline, ok = c.Context().Deadline()
	return
}

func (c Context) Done() <-chan struct{} {
	return c.Context().Done()
}

func (c Context) Err() error {
	if c.error != nil {
		return c.error
	}
	return c.Context().Err()
}

func (c Context) Value(key any) any {
	return c.Context().Value(key)
}

func (c Context) Cost() time.Duration {
	return c.cost
}

func (c Context) WithContext(ctx context.Context) Context {
	c.ctx = ctx
	return c
}

func (c Context) WithTraceId(id string) Context {
	c.traceId = id
	return c
}

type Trace struct {
	Caller string
	Func   string
	Cost   time.Duration
	SQL    string
}

func (c Context) Traces() []Trace {
	return nil
}

func (c Context) Context() (ctx context.Context) {
	if c.ctx != nil {
		ctx = c.ctx
	} else {
		ctx = context.Background()
	}
	return
}

func (c Context) Debug() Context {
	c.debug = true
	return Context{debug: true}
}

func (c Context) Must() Context {
	c.must = true
	return c
}

func (c Context) Strict() Context {
	c.exact = true
	return c
}

func (c Context) Mute() Context {
	c.mute = true
	return c
}

func (c Context) Loose() Context {
	c.loose = true
	return c
}

func (c Context) Analyze() Context {
	c.analyze = true
	return c
}

//func (c Context) WithTx(tx *DB) Context {
//	c.tx = tx
//	return c
//}

func (c Context) isTx() bool {
	return c.tx != nil
}

func (c Context) useTx() (tx *DB) {
	return c.tx
}
