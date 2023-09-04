package batis

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gobatis/gobatis/dialector"
)

const (
	txKey      = "GOBATIS_TX"
	traceIdKey = "GOBATIS_TRACE_ID"
	space      = " "
)

func WithTx(parent context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(parent, txKey, tx)
}

func WithTraceId(parent context.Context, traceId string) context.Context {
	return context.WithValue(parent, traceIdKey, traceId)
}

func WithDebug(parent context.Context, debug bool) context.Context {
	return context.WithValue(parent, "debug", debug)
}

func Open(d dialector.Dialector, options ...Option) (db *DB, err error) {

	config := &Config{
		CreateBatchSize: 10,
		Plugins:         nil,
		NowFunc: func() time.Time {
			return time.Now()
		},
		Dialector: d,
		Logger:    &logger{},
		db:        nil,
	}

	config.db, err = d.DB()
	if err != nil {
		return
	}

	db = &DB{Config: config, Error: nil}

	return
}

type DB struct {
	*Config
	//*executor
	Error   error
	tx      *tx
	ctx     context.Context
	debug   bool
	must    bool
	traceId string
	closure *closure
}

func (d *DB) addError(err error) {
	d.Error = addError(d.Error, err)
}

func (d *DB) clone() *DB {
	return &DB{
		Config:  d.Config,
		Error:   d.Error,
		tx:      d.tx,
		ctx:     d.ctx,
		debug:   d.debug,
		must:    d.must,
		traceId: d.traceId,
		//executor: nil,
	}
}

func (d *DB) WithTraceId(traceId string) *DB {
	c := d.clone()
	c.traceId = traceId
	return c
}

func (d *DB) WithContext(ctx context.Context) *DB {
	v := ctx.Value(txKey)
	if v != nil {
		if vv, ok := v.(*DB); ok {
			c := vv.clone()
			c.ctx = ctx
			return c
		}
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

func (d *DB) Close() error {
	return d.db.Close()
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

func (d *DB) context() context.Context {
	if d.ctx == nil {
		return context.Background()
	}
	return d.ctx
}

func (d *DB) exec(dest any) {
	if d.Error != nil {
		return
	}
	//if d.executor == nil {
	//	d.addError(fmt.Errorf("no executor"))
	//	return
	//}
	//d.dest = dest
	//d.addError(d.executor.execute())
	//if d.tx == nil {
	//	d.addError(d.executor.conn.Close())
	//}
	//d.executor.log(d)
}

// 执行查询语句
func (d *DB) Query(sql string, params ...NameValue) *DB {
	c := d.clone()
	e := newExecutor(c.Dialector.Namer(), true, query{sql: sql, params: params})
	if e.err != nil {
		c.addError(e.err)
		return c
	}
	c.closure = newClosure(c.db, c.tx, e)
	return c
}

// 扫描结果集
func (d *DB) Scan(dest any) *DB {
	d.exec(dest)
	return d
}

// 执行 SQL 语句
func (d *DB) Exec(sql string, params ...NameValue) *DB {
	c := d.clone()
	e := newExecutor(c.Dialector.Namer(), false, exec{sql: sql, params: params})
	if e.err != nil {
		c.addError(e.err)
		return c
	}
	c.closure = newClosure(c.db, c.tx, e)
	c.exec(nil)
	return c
}

// 执行删除操作
func (d *DB) Delete(table string, where Element) *DB {
	c := d.clone()
	e := newExecutor(c.Dialector.Namer(), false, del{table: table, elems: []Element{where}})
	if e.err != nil {
		c.addError(e.err)
		return c
	}
	c.closure = newClosure(c.db, c.tx, e)
	c.exec(nil)
	return c
}

// 执行更新操作
func (d *DB) Update(table string, data map[string]any, where Element) *DB {
	c := d.clone()
	u := update{table: table, data: data, elems: []Element{where}}
	q := u.returning != nil
	c.closure = newClosure(c.db, c.tx, newExecutor(d.Dialector.Namer(), q, u))
	if q {
		c.exec(nil)
	}
	return c
}

// 插入数据
func (d *DB) Insert(table string, data any, elems ...Element) *DB {
	c := d.clone()
	i := &insert{table: table, data: data, elems: elems}
	q := i.returning != nil
	c.closure = newClosure(c.db, c.tx, newExecutor(c.Dialector.Namer(), q, i))
	if q {
		c.exec(nil)
	}
	return c
}

func (d *DB) InsertBatch(table string, batch int, data any, onConflict, returning Element) *DB {
	c := d.clone()
	//var t bool
	//if c.tx == nil {
	//	c.Begin()
	//	t = true
	//}
	//if c.Error != nil {
	//	return c
	//}
	//defer func() {
	//	if t && c.Error != nil {
	//		c.Rollback()
	//	}
	//	if c.executor.conn != nil {
	//		c.addError(c.executor.conn.Close())
	//	}
	//}()
	//i := &insertBatch{
	//	table: table,
	//	batch: batch,
	//	data:  data,
	//	elems: []Element{
	//		onConflict,
	//		returning,
	//	}}
	//c.prepare(false, i)
	//if returning == nil {
	//	c.addError(c.executor.insertBatch(batch))
	//}
	//if t && c.Error == nil {
	//	c.Commit()
	//}
	return c
}

func (d *DB) ParallelQuery(queryer ...Queryer) *DB {
	c := d.clone()
	//if len(queryer) == 0 {
	//	c.Error = fmt.Errorf("no querer")
	//	return c
	//}
	//if d.executor != nil {
	//	c.Error = fmt.Errorf("db has origin executor")
	//	return c
	//}
	//
	//var executors []*executor
	//for _, v := range queryer {
	//	items, err := v.executors(d.Dialector.Namer(), "db")
	//	if err != nil {
	//		c.addError(err)
	//		return c
	//	}
	//	executors = append(executors, items...)
	//}
	//
	//wg := sync.WaitGroup{}
	//lock := sync.Mutex{}
	//
	//for _, v := range executors {
	//	v.query = true
	//	if d.tx != nil {
	//		v.conn = d.tx
	//	} else {
	//		var err error
	//		v.conn, err = d.db.Conn(d.context())
	//		if err != nil {
	//			d.addError(err)
	//			return c
	//		}
	//	}
	//	wg.Add(1)
	//	go func(v *executor) {
	//		defer func() {
	//			wg.Done()
	//		}()
	//		err := v.execute()
	//		if err != nil {
	//			lock.Lock()
	//			c.addError(err)
	//			lock.Unlock()
	//		}
	//	}(v)
	//}
	//wg.Wait()
	//d.executor.log(d)

	return c
}

func (d *DB) Result() (r sql.Result, err error) {
	//if d.executor == nil || d.executor.query {
	//	err = fmt.Errorf("no execute result")
	//	return
	//}
	//r = d.executor.result
	return
}

func (d *DB) Fetch(sql string, params ...NameValue) <-chan Scanner {

	c := d.clone()

	c.closure = newClosure(d.db, d.tx,
		newExecutor(d.Dialector.Namer(), false, newInnerSQL("begin;")),
		newExecutor(d.Dialector.Namer(), false, newInnerSQL("forward 10;")),
		newExecutor(d.Dialector.Namer(), false, newInnerSQL("commit;")),
	)

	ch := make(chan Scanner)
	//f := &fetch{}
	//d.prepareExecutor(true, f)

	return ch
}

func (d *DB) Begin() {
	c := d.clone()
	if c.tx != nil {
		c.addError(fmt.Errorf("tx conflict"))
	} else {
		t, err := d.db.Begin()
		if err != nil {
			c.addError(err)
			return
		}
		c.tx = &tx{Tx: t}
	}
	return
}

func (d *DB) Commit() *DB {
	if d.tx == nil {
		d.addError(ErrInvalidTransaction)
	} else {
		d.addError(d.tx.Commit())
	}
	return d
}

func (d *DB) Rollback() *DB {
	if d.tx == nil {
		d.addError(ErrInvalidTransaction)
	} else {
		d.addError(d.tx.Rollback())
	}
	return d
}
