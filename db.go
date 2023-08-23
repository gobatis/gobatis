package batis

import (
	"context"
	"database/sql"
	"fmt"
	"sync/atomic"
	
	"github.com/gobatis/gobatis/dialector"
)

const (
	txKey = "GOBATIS_TX"
	space = " "
)

func WrapTx(parent context.Context, tx *DB) context.Context {
	return context.WithValue(parent, txKey, tx)
}

func Open(d dialector.Dialector, options ...Option) (db *DB, err error) {
	db = &DB{
		Error:  nil,
		db:     nil,
		logger: nil,
		tx:     nil,
		namer:  d.Namer(),
	}
	db.db, err = d.DB()
	if err != nil {
		return
	}
	return
}

type DB struct {
	Error        error
	RowsAffected *int64
	LastInsertId *int64
	db           *sql.DB
	logger       Logger
	tx           *tx
	ctx          context.Context
	debug        bool
	must         bool
	namer        dialector.Namer
	traceId      string
	executors    []executor
	executed     atomic.Bool
}

func (d *DB) clone() *DB {
	return &DB{
		db:      d.db,
		logger:  d.logger,
		tx:      d.tx,
		ctx:     d.ctx,
		debug:   d.debug,
		must:    d.must,
		Error:   d.Error,
		namer:   d.namer,
		traceId: "",
	}
}

func (d *DB) WithTraceId(traceId string) *DB {
	c := d.clone()
	c.traceId = traceId
	
	return c
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

func (d *DB) context() context.Context {
	if d.ctx == nil {
		return context.Background()
	}
	return d.ctx
}

func (d *DB) execute(dest ...any) {
	
	defer func() {
		if d.Error != nil {
			// log
		}
	}()
	
	if d.Error != nil {
		return
	}
	
	if d.executed.Swap(true) {
		d.Error = fmt.Errorf("db has executed")
		return
	}
	
	l1 := len(dest)
	l2 := len(d.executors)
	
	if l2 == 0 {
		d.Error = fmt.Errorf("no exector")
		return
	}
	
	if l1 > l2 {
		d.Error = fmt.Errorf("expect %d dest, got %d exector", l1, l2)
		return
	}
	
	for i, v := range d.executors {
		var vv any
		if i < l1 {
			vv = dest[i]
		}
		d.RowsAffected, d.LastInsertId, d.Error = v.Exec(vv)
		if d.Error != nil {
			return
		}
	}
	
	return
}

func (d *DB) prepare(query bool, elem Element) {
	
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
	
	e := executor{query: query}
	if d.tx != nil {
		e.conn = d.tx
	} else {
		e.conn, err = d.db.Conn(d.context())
		if err != nil {
			return
		}
	}
	e.sql, e.params, err = elem.SQL(d.namer, "db")
	if err != nil {
		return
	}
	
	return
}

// Query 执行查询语句
func (d *DB) Query(sql string, params ...NameValue) *DB {
	//d.prepare(true, query{sql: sql, params: params})
	//queryer.Queries()
	return d
}

func test2() {
	db := &DB{}
	var users []string
	err := db.Query(`select * from public.users where active = #{args.Paging} and age < #{args.Node}`,
		Param("args", 18)).Scan(&users).Error
	_ = err
}

// Scan 扫描结果集
func (d *DB) Scan(dest ...any) *DB {
	d.execute(dest...)
	return d
}

//func (d *DB) Build(b Builder) Scanner {
//	//t := d.initTracer()
//	//es, err := b.Build()
//	//if err != nil {
//	//	t.log()
//	//	return Scanner{Error: err}
//	//}
//	//
//	s := &Scanner{}
//	//g := errgroup.Group{}
//	//for _, v := range es {
//	//	e := v
//	//	e.conn = d.db
//	//	t.err = d.err
//	//	g.Go(func() error {
//	//		// todo auto cancel
//	//		e.Exec(s)
//	//		return t.err
//	//	})
//	//}
//	//err = g.Wait()
//	//if err != nil {
//	//	return Scanner{Error: err}
//	//}
//	
//	return *s
//}

func (d *DB) Exec(sql string, params ...NameValue) *DB {
	d.prepare(false, exec{sql: sql, params: params})
	d.execute()
	return d
}

func (d *DB) Delete(table string, where Element) *DB {
	e := &del{table: table, elems: []Element{where}}
	q := e.returning != nil
	d.prepare(q, e)
	if !q {
		d.execute()
	}
	return d
}

func (d *DB) Update(table string, data map[string]any, where Element) *DB {
	u := update{table: table, data: data, elems: []Element{where}}
	q := u.returning != nil
	d.prepare(q, u)
	if !q {
		d.execute()
	}
	return d
}

func (d *DB) Insert(table string, data any, elems ...Element) *DB {
	i := &insert{table: table, data: data, elems: elems}
	q := i.returning != nil
	d.prepare(q, i)
	if !q {
		d.execute()
	}
	return d
}

func (d *DB) InsertBatch(table string, batch int, data any, onConflict Element) *DB {
	i := &insertBatch{table: table, batch: batch, data: data, elems: []Element{onConflict}}
	d.prepare(false, i)
	q := i.returning != nil
	d.prepare(q, i)
	if !q {
		d.execute()
	}
	return d
}

func (d *DB) ParallelQuery(queryer ...Queryer) *DB {
	
	return d
}

func test() {
	d := &DB{}
	var items []string
	d.ParallelQuery(&Paging{
		Select: "",
		Count:  "",
		Common: "",
		Page:   0,
		Limit:  0,
		Params: nil,
		Scan:   []any{&items},
	})
	
	d.ParallelQuery(
		&Query{
			SQL:    "",
			Params: nil,
			Scan:   nil,
		},
		&Query{
			SQL:    "",
			Params: nil,
			Scan:   nil,
		},
	)
}

type Queryer interface {
	Queries() ([]executor, error)
}

//func (d *DB) Fetch(sql string, params ...NameValue) <-chan Scanner {
//	ch := make(chan Scanner)
//	f := &fetch{}
//	d.prepareExecutor(true, f)
//	return ch
//}

func (d *DB) Begin() *DB {
	if d.Error != nil {
		return d
	}
	var t *sql.Tx
	t, d.Error = d.db.Begin()
	if d.Error != nil {
		return d
	}
	d.tx = &tx{Tx: t}
	return d
}

func (d *DB) Commit() error {
	return d.tx.Commit()
}

func (d *DB) Rollback() error {
	return d.tx.Rollback()
}

func LooseDest(dest any, fields ...string) Dest {
	return Dest{loose: true, dest: dest, fields: fields}
}

type Dest struct {
	loose  bool
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

    db.Must().Query(`select * from users`).LooseScan(&users, "$..Posts");

	db.Must().Exec(`select`)

	userIds := mapping.Merge(users, func())

	db.Query().LooseScan(batis.LooseDest(&users, "$..Posts","$..Orders")).Error

    db.Query(`select * from posts where user_id in #{userIds}`,batis.Param("userIds", userIds)).Link(&users, "user_id => $..Id", "$..Posts").Error

	postIds := mapping.Map(users, func())

	db.Query(`select * from tags where post_id in #{postIds}`, batis.Param("postIds", postIds)).Link(&users, "user_id => $..Posts[*].Id", "$..Post[*].Tags").Error

    db.Query(`select * from orders where user_id in #{userIds}`,batis.Param("userIds", userIds)).Link(&users, "user_id => $..Id", "$..Orders").Error

*/
