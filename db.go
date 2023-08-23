package batis

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	
	"github.com/gobatis/gobatis/dialector"
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
	Error   error
	tracer  *tracer
	dirty   bool
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
		tracer: d.tracer,
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

func (d *DB) initTracer() *tracer {
	t := &tracer{
		//err:     d.err,
		now:     time.Now(),
		debug:   d.debug,
		logger:  d.useLogger(),
		traceId: d.traceId,
	}
	return t
}

func (d *DB) context() context.Context {
	if d.ctx == nil {
		return context.Background()
	}
	return d.ctx
}

func (d *DB) execute(query bool, elem Element) {
	
	d.tracer = d.initTracer()
	
	var err error
	defer func() {
		if err != nil {
			d.Error = err
		}
	}()
	
	if d.Error != nil {
		err = d.Error
		return
	}
	if d.tracer != nil {
		d.Error = fmt.Errorf("db is dirty")
		return
	}
	
	//if t.err != nil {
	//	t.log()
	//	return
	//}
	//var c *sql.Conn
	//defer func() {
	//	if c != nil {
	//		_ = c.Close()
	//	}
	//}()
	//e := &executor{
	//	query:  query,
	//	tracer: t,
	//}
	//if d.tx != nil {
	//	e.conn = d.tx
	//} else {
	//	c, t.err = d.db.Conn(d.context())
	//	if t.err != nil {
	//		t.log()
	//		return
	//	}
	//	e.conn = c
	//}
	//e.sql, e.params, e.tracer.err = elem.SQL(d.namer, "db")
	//if e.tracer.err != nil {
	//	t.log()
	//	return
	//}
	//s := &Scanner{tracer: t}
	//e.Exec(s)
	//for _, v := range s.rows {
	//	_ = v.Close()
	//}
	return
}

func (d *DB) Query(sql string, params ...NameValue) *DB {
	d.execute(true, query{sql: sql, params: params})
	return d
}

func (d *DB) Build(b Builder) Scanner {
	//t := d.initTracer()
	//es, err := b.Build()
	//if err != nil {
	//	t.log()
	//	return Scanner{Error: err}
	//}
	//
	s := &Scanner{}
	//g := errgroup.Group{}
	//for _, v := range es {
	//	e := v
	//	e.conn = d.db
	//	t.err = d.err
	//	g.Go(func() error {
	//		// todo auto cancel
	//		e.Exec(s)
	//		return t.err
	//	})
	//}
	//err = g.Wait()
	//if err != nil {
	//	return Scanner{Error: err}
	//}
	
	return *s
}

func (d *DB) Exec(sql string, params ...NameValue) *DB {
	d.execute(false, exec{sql: sql, params: params})
	return d
}

func (d *DB) Delete(table string, where Element) *DB {
	e := &del{table: table, elems: []Element{where}}
	d.execute(false, e)
	return d
}

func (d *DB) Update(table string, data map[string]any, where Element) *DB {
	d.execute(false, update{table: table, data: data, elems: []Element{where}})
	return d
}

func (d *DB) Insert(table string, data any, elems ...Element) *DB {
	i := &insert{table: table, data: data, elems: elems}
	d.execute(true, i)
	return d
}

func (d *DB) InsertBatch(table string, batch int, data any, onConflict Element) *DB {
	i := &insertBatch{table: table, batch: batch, data: data, elems: []Element{onConflict}}
	d.execute(false, i)
	return d
}

//func (d *DB) Fetch(sql string, params ...NameValue) <-chan Scanner {
//	ch := make(chan Scanner)
//	f := &fetch{}
//	d.execute(true, f)
//	return ch
//}

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

// Scan 扫描结果集
func (d *DB) Scan(dest ...any) *DB {
	
	return d
}

// LooseScan 宽松扫描模式
func (d *DB) LooseScan(dest ...LDest) *DB {
	return d
}

func LooseDest(dest any, fields ...string) LDest {
	return LDest{dest: dest, fields: fields}
}

type LDest struct {
	dest   any
	fields []string
}

/*
    type User struct {
		Id     int64
		Name   string
		Posts  []string
		Orders []string
	}

	type Post struct {
		Id   int64
		Tags []string
	}

	var users []User
	db.Query(`select * from users`).Scan(&users);

	userIds := mapping.Merge(users, func())

	db.Query().LooseScan(batis.LooseDest(&users, "$..Posts","$..Orders")).Error

    db.Query(`select * from posts where user_id in #{userIds}`,batis.Param("userIds", userIds)).Link(&users, "user_id => $..Id", "$..Posts").Error

	postIds := mapping.Map(users, func())

	db.Query(`select * from tags where post_id in #{postIds}`, batis.Param("postIds", postIds)).Link(&users, "user_id => $..Posts[*].Id", "$..Post[*].Tags").Error

    db.Query(`select * from orders where user_id in #{userIds}`,batis.Param("userIds", userIds)).Link(&users, "user_id => $..Id", "$..Orders").Error

*/
