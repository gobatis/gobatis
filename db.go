package batis

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/gobatis/gobatis/dialector"
	"github.com/gobatis/gobatis/executor"
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
		Logger:    executor.DefaultLogger(),
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
	Error    error
	tx       *executor.Tx
	ctx      context.Context
	trace    bool
	debug    bool
	must     bool
	traceId  string
	executor executor.Executor
}

func (d *DB) addError(err error) {
	d.Error = executor.AddError(d.Error, err)
}

func (d *DB) clone() *DB {
	return &DB{
		Config:   d.Config,
		Error:    d.Error,
		tx:       d.tx,
		ctx:      d.ctx,
		trace:    d.trace,
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

func (d *DB) Trace() *DB {
	c := d.clone()
	c.trace = true
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

func (d *DB) execute(dest any) {
	if d.Error != nil {
		return
	}
	if d.executor == nil {
		d.addError(fmt.Errorf("no executor"))
		return
	}
	//d.dest = dest
	//d.addError(d.executor.execute())
	//if d.tx == nil {
	//	d.addError(d.executor.conn.Close())
	//}
	//d.executor.log(d)
}

func (d *DB) conn() executor.Conn {
	if d.tx != nil {
		return d.tx
	} else {
		return executor.NewDB(d.db)
	}
}

func (d *DB) prepare(query bool, element Element) (conn executor.Conn, raw *executor.Raw, err error) {

	conn = d.conn()

	raw, err = element.Raw(d.Dialector.Namer(), "db")
	if err != nil {
		return
	}
	raw.Query = query
	raw.Ctx = d.context()

	return
}

// 执行查询语句
func (d *DB) Query(sql string, params ...executor.Param) *DB {
	c := d.clone()
	conn, raw, err := c.prepare(true, query{sql: sql, params: params})
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(conn, raw)
	return c
}

// 扫描结果集
func (d *DB) Scan(dest any) *DB {
	d.execute(dest)
	return d
}

// 执行 SQL 语句
func (d *DB) Exec(sql string, params ...executor.Param) *DB {
	c := d.clone()
	conn, raw, err := c.prepare(false, exec{sql: sql, params: params})
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(conn, raw)
	c.execute(nil)
	return c
}

// 执行删除操作
func (d *DB) Delete(table string, where Element) *DB {
	c := d.clone()
	conn, raw, err := c.prepare(false, del{table: table, elems: []Element{where}})
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(conn, raw)
	c.execute(nil)
	return c
}

// 执行更新操作
func (d *DB) Update(table string, data map[string]any, where Element) *DB {
	c := d.clone()
	u := update{table: table, data: data, elems: []Element{where}}
	q := u.returning != nil
	conn, raw, err := c.prepare(false, u)
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(conn, raw)
	if q {
		c.execute(nil)
	}
	return c
}

// 插入数据
func (d *DB) Insert(table string, data any, elems ...Element) *DB {
	c := d.clone()
	i := &insert{table: table, data: data, elems: elems}
	q := i.returning != nil
	conn, raw, err := c.prepare(false, i)
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(conn, raw)
	if q {
		c.execute(nil)
	}
	return c
}

func (d *DB) setExecutor(e executor.Executor) {
	if d.executor != nil {
		d.addError(fmt.Errorf("executor duplicated"))
		return
	}
	d.executor = e
}

func (d *DB) InsertBatch(table string, batch int, data any, onConflict, returning Element) *DB {
	c := d.clone()
	i := &insertBatch{
		table: table,
		batch: batch,
		data:  data,
		elems: []Element{
			onConflict,
			returning,
		}}
	q := i.returning != nil
	conn, raw, err := c.prepare(q, i)
	if err != nil {
		c.addError(err)
		return c
	}
	c.setExecutor(executor.NewInsertBatch(conn, raw))
	if !q {
		c.execute(nil)
	}
	return c
}

func (d *DB) ParallelQuery(queryer ...ParallelQueryer) *DB {
	c := d.clone()
	if len(queryer) == 0 {
		c.Error = fmt.Errorf("no querer")
		return c
	}
	if d.executor != nil {
		c.Error = fmt.Errorf("db executor is not empty")
		return c
	}

	var executors []executor.Executor
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
		wg.Add(1)
		go func(v executor.Executor) {
			defer func() {
				wg.Done()
			}()
			err := v.Execute(nil)
			if err != nil {
				lock.Lock()
				c.addError(err)
				lock.Unlock()
			}
		}(v)
	}
	wg.Wait()

	return c
}

func (d *DB) LastInsertId() (int64, error) {
	return d.executor.Result().LastInsertId()
}

func (d *DB) RowsAffected() (int64, error) {
	return d.executor.Result().RowsAffected()
}

func (d *DB) Result() (r sql.Result, err error) {
	//if d.executor == nil || d.executor.query {
	//	err = fmt.Errorf("no execute result")
	//	return
	//}
	//r = d.executor.result
	return
}

func (d *DB) FetchQuery(query FetchQuery) error {

	c := d.clone()

	raw := &executor.Raw{
		Ctx:    d.context(),
		Query:  true,
		SQL:    query.SQL,
		Params: nil,
	}

	for k, v := range query.Params {
		raw.Params = append(raw.Params, executor.Param{
			Name:  k,
			Value: v,
		})
	}

	d.setExecutor(executor.NewFetchQuery(c.conn(), raw, query.Limit))

	return d.executor.Execute(func(s *executor.Scanner) error {
		return query.Scan(s)
	})
}

func (d *DB) Begin() {
	c := d.clone()
	if c.tx != nil {
		c.addError(fmt.Errorf("tx conflict"))
	} else {
		tx, err := d.db.Begin()
		if err != nil {
			c.addError(err)
			return
		}
		c.tx = executor.NewTx(tx)
	}
	return
}

func (d *DB) Commit() *DB {
	if d.tx == nil {
		d.addError(executor.ErrInvalidTransaction)
	} else {
		d.addError(d.tx.Commit())
	}
	return d
}

func (d *DB) Rollback() *DB {
	if d.tx == nil {
		d.addError(executor.ErrInvalidTransaction)
	} else {
		d.addError(d.tx.Rollback())
	}
	return d
}
