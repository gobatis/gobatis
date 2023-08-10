package batis

import (
	"context"
	"database/sql"
	"strings"
	"sync"
	
	"github.com/gobatis/gobatis/builder"
	"github.com/gobatis/gobatis/executor"
	"golang.org/x/sync/errgroup"
)

func Open(d Dialector, options ...Option) (db *DB, err error) {
	db = &DB{
		db:     nil,
		logger: newLogger(),
		tx:     nil,
		mu:     sync.RWMutex{},
		err:    nil,
		namer:  d.WrapName,
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
	logger Logger
	tx     *sql.Tx
	mu     sync.RWMutex
	ctx    context.Context
	debug  bool
	must   bool
	loose  bool
	err    error
	namer  func(name string) string
}

func (d *DB) fork() *DB {
	return &DB{
		db:     d.db,
		logger: d.logger,
		tx:     d.tx,
		mu:     sync.RWMutex{},
		ctx:    d.ctx,
		debug:  d.debug,
		must:   d.must,
		err:    d.err,
	}
}

func (d *DB) WithContext(ctx context.Context) *DB {
	v, ok := ctx.Value(contextTxKey).(*DB)
	if ok {
		f := v.fork()
		f.ctx = ctx
		return f
	}
	f := d.fork()
	f.ctx = ctx
	return f
}

func (d *DB) Debug() *DB {
	f := d.fork()
	f.debug = true
	return f
}

func (d *DB) Must() *DB {
	f := d.fork()
	f.debug = true
	return f
}

func (d *DB) SetLogLevel(level Level) {
	d.logger.SetLevel(level)
}

func (d *DB) SetLogger(logger Logger) {
	d.logger = logger
}

func (d *DB) useLogger() Logger {
	if d.logger == nil {
		d.logger = newLogger()
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

func (d *DB) Loose() *DB {
	f := d.fork()
	f.loose = true
	return f
}

//func (d *DB) Prepare(sql string, params ...executor.NameValue) *Stmt {
//	return &Stmt{}
//}
const space = " "

func (d *DB) exec2(et int, elems ...Element) executor.Scanner {
	var sqls []string
	params := map[string]executor.NameValue{}
	for _, v := range elems {
		s, err := v.SQL(d.namer)
		if err != nil {
			return executor.Scanner{
				Error: err,
			}
		}
		sqls = append(sqls, s)
		for _, vv := range v.Params() {
			params[vv.Name] = vv
		}
	}
	s := &executor.Scanner{}
	e := &executor.Executor{
		Type: et,
		SQL:  strings.Join(sqls, space),
		Err:  d.err,
		Conn: d.db,
	}
	for _, v := range params {
		e.Params = append(e.Params, v)
	}
	e.Exec(s)
	
	return *s
}

func (d *DB) exec(_type int, sql string, params []executor.NameValue) executor.Scanner {
	s := &executor.Scanner{}
	e := &executor.Executor{
		Type:   _type,
		SQL:    sql,
		Params: params,
		Err:    d.err,
		Conn:   d.db,
	}
	e.Exec(s)
	return *s
}

func (d *DB) Query(sql string, params ...executor.NameValue) executor.Scanner {
	return d.exec(executor.Query, sql, params)
}

func (d *DB) Build(b builder.Builder) executor.Scanner {
	
	es, err := b.Build()
	if err != nil {
		return executor.Scanner{Error: err}
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
		s.Error = err
		return *s
	}
	
	return *s
}

func (d *DB) Execute(sql string, params ...executor.NameValue) executor.Scanner {
	return d.exec(executor.Exec, sql, params)
}

func (d *DB) Delete(table string, where Element) executor.Scanner {
	//s := fmt.Sprintf("delete from %s %s", table, where.SQL())
	//return d.exec(executor.Exec, s, where.Params())
	e := &_delete{
		table: table,
		where: where,
	}
	return d.exec2(executor.Exec, e)
}

func (d *DB) Update(table string, data map[string]any, where Element) executor.Scanner {
	u := &update{
		table: "",
		data:  nil,
		where: where,
	}
	return d.exec2(executor.Exec, u)
}

func (d *DB) Insert(table string, data any, onConflict ...Element) executor.Scanner {
	//s := fmt.Sprintf("insert into %s", table)
	//return d.exec(executor.Exec, s, nil)
	i := &insert{
		table: table,
		data:  data,
		elems: onConflict,
	}
	return d.exec2(executor.Query, i)
	
}

func (d *DB) InsertBatch(table string, data any, batch int, onConflict ...Element) executor.Scanner {
	//s := fmt.Sprintf("insert into %s", table)
	//return d.exec(executor.Exec, s, nil)
	ib := &insertBatch{
		table: table,
		data:  data,
		batch: batch,
		elems: onConflict,
	}
	return d.exec2(executor.Exec, ib)
}

func (d *DB) Fetch(sql string, params ...executor.NameValue) <-chan executor.Scanner {
	
	ch := make(chan executor.Scanner)
	
	f := &fetch{}
	d.exec2(executor.Query, f)
	
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
