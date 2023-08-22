package batis

import (
	"context"
	"database/sql"
	"time"
	
	"github.com/gobatis/gobatis/dialector"
	"golang.org/x/sync/errgroup"
)

const txKey = "GOBATIS_TX"

func WithTx(parent context.Context, tx *DB) context.Context {
	return context.WithValue(parent, txKey, tx)
}

func Open(d dialector.Dialector, options ...Option) (db *DB, err error) {
	db = &DB{
		db:     nil,
		logger: nil,
		tx:     nil,
		err:    nil,
		namer:  d.Namer(),
	}
	db.db, err = d.DB()
	if err != nil {
		return
	}
	return
}

type DB struct {
	db      *sql.DB
	logger  Logger
	tx      *sql.Tx
	ctx     context.Context
	debug   bool
	must    bool
	loose   bool
	err     error
	namer   dialector.Namer
	traceId string
}

func (d *DB) clone() *DB {
	return &DB{
		db:     d.db,
		logger: d.logger,
		tx:     d.tx,
		ctx:    d.ctx,
		debug:  d.debug,
		must:   d.must,
		loose:  d.loose,
		err:    d.err,
		namer:  d.namer,
	}
}

func (d *DB) WithContext(ctx context.Context) *DB {
	v, ok := ctx.Value(txKey).(*DB)
	if ok {
		c := v.clone()
		c.ctx = ctx
		return c
	}
	c := d.clone()
	c.ctx = ctx
	return c
}

func (d *DB) Debug() *DB {
	c := d.clone()
	c.debug = true
	return c
}

func (d *DB) Must() *DB {
	c := d.clone()
	c.debug = true
	return c
}

func (d *DB) Loose() *DB {
	f := d.clone()
	f.loose = true
	return f
}

func (d *DB) SetLogger(logger Logger) {
	d.logger = logger
}

func (d *DB) useLogger() Logger {
	if d.logger == nil {
		d.logger = DefaultLogger()
	}
	return d.logger
}

func (d *DB) Close() {
	_ = d.db.Close()
}

func (d *DB) DB() *sql.DB {
	return d.db
}

func (d *DB) Ping() error {
	return d.db.Ping()
}

func (d *DB) Stats() sql.DBStats {
	return d.db.Stats()
}

func (d *DB) Prepare(sql string, params ...NameValue) *Stmt {
	return &Stmt{}
}

const space = " "

func (d *DB) tracer() *tracer {
	t := &tracer{
		now:     time.Now(),
		debug:   d.debug,
		logger:  d.useLogger(),
		traceId: d.traceId,
	}
	return t
}

func (d *DB) execute(query bool, elem Element) Scanner {
	t := d.tracer()
	e := &executor{
		query:  query,
		conn:   d.db,
		tracer: t,
	}
	e.sql, e.params, e.tracer.err = elem.SQL(d.namer, "db")
	if e.tracer.err != nil {
		t.log()
		return Scanner{Error: t.err}
	}
	s := &Scanner{tracer: t}
	e.Exec(s)
	return *s
}

func (d *DB) Query(sql string, params ...NameValue) Scanner {
	return d.execute(true, query{sql: sql, params: params})
}

func (d *DB) Build(b Builder) Scanner {
	t := d.tracer()
	es, err := b.Build()
	if err != nil {
		t.log()
		return Scanner{Error: err}
	}
	
	s := &Scanner{}
	g := errgroup.Group{}
	for _, v := range es {
		e := v
		e.conn = d.db
		t.err = d.err
		g.Go(func() error {
			// todo auto cancel
			e.Exec(s)
			return t.err
		})
	}
	err = g.Wait()
	if err != nil {
		return Scanner{Error: err}
	}
	
	return *s
}

func (d *DB) Exec(sql string, params ...NameValue) Scanner {
	return d.execute(false, exec{sql: sql, params: params})
}

func (d *DB) Delete(table string, where Element) Scanner {
	e := &del{table: table, elems: []Element{where}}
	return d.execute(false, e)
}

func (d *DB) Update(table string, data map[string]any, where Element) Scanner {
	return d.execute(false, update{table: table, data: data, elems: []Element{where}})
}

func (d *DB) Insert(table string, data any, elems ...Element) Scanner {
	i := &insert{table: table, data: data, elems: elems}
	return d.execute(true, i)
}

func (d *DB) InsertBatch(table string, batch int, data any, onConflict Element) Scanner {
	i := &insertBatch{table: table, batch: batch, data: data, elems: []Element{onConflict}}
	return d.execute(false, i)
}

func (d *DB) Fetch(sql string, params ...NameValue) <-chan Scanner {
	ch := make(chan Scanner)
	f := &fetch{}
	d.execute(true, f)
	return ch
}

func (d *DB) Begin() *DB {
	if d.err != nil {
		return d
	}
	d.tx, d.err = d.db.Begin()
	return d
}

func (d *DB) Commit() error {
	return d.tx.Commit()
}

func (d *DB) Rollback() error {
	return d.tx.Rollback()
}
