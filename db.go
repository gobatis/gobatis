package batis

import (
	"context"
	"database/sql"

	"github.com/gobatis/gobatis/builder"
	"github.com/gobatis/gobatis/dialector"
	"github.com/gobatis/gobatis/executor"
	"github.com/gozelle/logger"
	"golang.org/x/sync/errgroup"
)

func Open(d dialector.Dialector, options ...Option) (db *DB, err error) {
	db = &DB{
		db:     nil,
		logger: logger.NewLogger("[gobatis]"),
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

const contextTxKey = "GOBATIS_TX"

func NewTxContext(parent context.Context, tx *DB) context.Context {
	return context.WithValue(parent, contextTxKey, tx)
}

type DB struct {
	db     *sql.DB
	logger logger.EventLogger
	tx     *sql.Tx
	ctx    context.Context
	debug  bool
	must   bool
	loose  bool
	err    error
	namer  dialector.Namer
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
	v, ok := ctx.Value(contextTxKey).(*DB)
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

//func (d *DB) SetLogLevel(level Level) {
//	d.logger.SetLevel(level)
//}
//
//func (d *DB) SetLogger(logger Logger) {
//	d.logger = logger
//}
//
//func (d *DB) useLogger() Logger {
//	if d.logger == nil {
//		d.logger = newLogger()
//	}
//	return d.logger
//}

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

func (d *DB) Loose() *DB {
	f := d.clone()
	f.loose = true
	return f
}

//func (d *DB) Prepare(sql string, params ...executor.NameValue) *Stmt {
//	return &Stmt{}
//}

const space = " "

func (d *DB) execute(et int, elem Element) executor.Scanner {
	e := &executor.Executor{Type: et, Conn: d.db, Debug: d.debug, Logger: d.logger}
	e.SQL, e.Params, e.Err = elem.SQL(d.namer, "db")
	if e.Err != nil {
		return executor.NewErrorScanner(e.Err)
	}
	s := &executor.Scanner{}
	e.Exec(s)
	return *s
}

//func (d *DB) exec(typ int, sql string, params []executor.NameValue) executor.Scanner {
//	s := &executor.Scanner{}
//	e := &executor.Executor{Type: typ, SQL: sql, Params: params, Err: d.err, Conn: d.db}
//	e.Exec(s)
//	return *s
//}

func (d *DB) Query(sql string, params ...executor.NameValue) executor.Scanner {
	return d.execute(executor.Query, query{sql: sql, params: params})
}

func (d *DB) Build(b builder.Builder) executor.Scanner {

	es, err := b.Build()
	if err != nil {
		return executor.NewErrorScanner(err)
	}

	s := &executor.Scanner{}
	g := errgroup.Group{}
	for _, v := range es {
		e := v
		e.Conn = d.db
		e.Err = d.err
		g.Go(func() error {
			// todo auto cancel
			e.Exec(s)
			return e.Err
		})
	}
	err = g.Wait()
	if err != nil {
		return executor.NewErrorScanner(err)
	}

	return *s
}

func (d *DB) Exec(sql string, params ...executor.NameValue) executor.Scanner {
	return d.execute(executor.Exec, exec{sql: sql, params: params})
}

func (d *DB) Delete(table string, where Element) executor.Scanner {
	e := &del{table: table, elems: []Element{where}}
	return d.execute(executor.Exec, e)
}

func (d *DB) Update(table string, data map[string]any, where Element) executor.Scanner {
	return d.execute(executor.Exec, update{table: table, data: data, elems: []Element{where}})
}

func (d *DB) Insert(table string, data any, elems ...Element) executor.Scanner {
	i := &insert{table: table, data: data, elems: elems}
	return d.execute(executor.Query, i)
}

func (d *DB) InsertBatch(table string, batch int, data any, onConflict Element) executor.Scanner {
	i := &insertBatch{table: table, batch: batch, data: data, elems: []Element{onConflict}}
	return d.execute(executor.Query, i)
}

func (d *DB) Fetch(sql string, params ...executor.NameValue) <-chan executor.Scanner {

	ch := make(chan executor.Scanner)

	f := &fetch{}
	d.execute(executor.Query, f)

	return ch
}

func (d *DB) Begin() *DB {
	if d.err != nil {
		return d
	}
	d.tx, d.err = d.db.Begin()
	return d
}

func (d *DB) Prepare(ctx context.Context, sql string, params ...executor.NameValue) (*sql.Stmt, error) {
	//return d.tx.PrepareContext(ctx, query)
	panic("todo")
}

func (d *DB) PrepareContext(sql string, params ...executor.NameValue) (*sql.Stmt, error) {
	panic("todo")
}

func (d *DB) Commit() error {
	return d.tx.Commit()
}

func (d *DB) Rollback() error {
	return d.tx.Rollback()
}
