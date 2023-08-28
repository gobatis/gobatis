package batis

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
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
	*executor
	Error   error
	tx      *tx
	ctx     context.Context
	debug   bool
	must    bool
	traceId string
}

func (d *DB) addError(err error) {
	d.Error = addError(d.Error, err)
}

func (d *DB) clone() *DB {
	return &DB{
		Config:   d.Config,
		Error:    d.Error,
		tx:       d.tx,
		ctx:      d.ctx,
		debug:    d.debug,
		must:     d.must,
		traceId:  d.traceId,
		executor: nil,
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

func (d *DB) prepare(query bool, elem Element) {
	if d.executor != nil {
		d.addError(fmt.Errorf("executor overridden"))
		return
	}

	var err error
	e := &executor{query: query, now: d.NowFunc()}
	if d.tx != nil {
		e.conn = d.tx
	} else {
		e.conn, err = d.db.Conn(d.context())
		if err != nil {
			d.addError(err)
			return
		}
	}
	e.raw, e.params, err = elem.SQL(d.Dialector.Namer(), "db")
	if err != nil {
		d.addError(err)
		return
	}
	d.executor = e
	return
}

func (d *DB) exec(dest any) {
	if d.Error != nil {
		return
	}
	if d.executor == nil {
		d.addError(fmt.Errorf("no executor"))
		return
	}
	d.dest = dest
	d.addError(d.executor.execute())
	if d.tx == nil {
		d.addError(d.executor.conn.Close())
	}
	d.executor.log(d)
}

// 执行查询语句
func (d *DB) Query(sql string, params ...NameValue) *DB {
	c := d.clone()
	c.prepare(true, query{sql: sql, params: params})
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
	c.prepare(false, exec{sql: sql, params: params})
	c.exec(nil)
	return c
}

// 执行删除操作
func (d *DB) Delete(table string, where Element) *DB {
	c := d.clone()
	e := &del{table: table, elems: []Element{where}}
	q := e.returning != nil
	c.prepare(q, e)
	if q {
		c.exec(nil)
	}
	return c
}

// 执行更新操作
func (d *DB) Update(table string, data map[string]any, where Element) *DB {
	c := d.clone()
	u := update{table: table, data: data, elems: []Element{where}}
	q := u.returning != nil
	c.prepare(q, u)
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
	c.prepare(q, i)
	if q {
		c.exec(nil)
	}
	return c
}

func (d *DB) InsertBatch(table string, batch int, data any, onConflict Element) *DB {
	c := d.clone()
	var t bool
	if c.tx == nil {
		c.Begin()
		t = true
	}
	if c.Error != nil {
		return c
	}
	defer func() {
		if t && c.Error != nil {
			c.Rollback()
		}
		if c.executor.conn != nil {
			c.addError(c.executor.conn.Close())
		}
	}()
	i := &insertBatch{table: table, batch: batch, data: data, elems: []Element{onConflict}}
	c.prepare(false, i)
	c.addError(c.executor.insertBatch(batch))
	if t && c.Error == nil {
		c.Commit()
	}
	return c
}

func (d *DB) ParallelQuery(queryer ...Queryer) *DB {
	c := d.clone()
	if len(queryer) == 0 {
		c.Error = fmt.Errorf("no querer")
		return c
	}
	if d.executor != nil {
		c.Error = fmt.Errorf("db has origin executor")
		return c
	}

	var executors []*executor
	for _, v := range queryer {
		items, err := v.executors(d.Dialector.Namer(), "db")
		if err != nil {
			c.addError(err)
			return c
		}
		executors = append(executors, items...)
	}

	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for _, v := range executors {
		v.query = true
		if d.tx != nil {
			v.conn = d.tx
		} else {
			var err error
			v.conn, err = d.db.Conn(d.context())
			if err != nil {
				d.addError(err)
				return c
			}
		}
		wg.Add(1)
		go func(v *executor) {
			defer func() {
				wg.Done()
			}()
			err := v.execute()
			if err != nil {
				lock.Lock()
				c.addError(err)
				lock.Unlock()
			}
		}(v)
	}
	wg.Wait()
	//d.executor.log(d)

	return c
}

func (d *DB) Result() (r sql.Result, err error) {
	if d.executor == nil || d.executor.query {
		err = fmt.Errorf("no execute result")
		return
	}
	r = d.executor.result
	return
}

//func (d *DB) Fetch(sql string, params ...NameValue) <-chan Scanner {
//	ch := make(chan Scanner)
//	f := &fetch{}
//	d.prepareExecutor(true, f)
//	return ch
//}

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
