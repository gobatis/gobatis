package batis

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/gobatis/gobatis/dialector"
	"github.com/gobatis/gobatis/executor"
	"go.uber.org/atomic"
)

const (
	dbKey      = "GOBATIS_DB"
	traceIdKey = "GOBATIS_TRACE_ID"
	space      = " "
)

func WithTx(parent context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(parent, dbKey, tx)
}

func WithTraceId(parent context.Context, traceId string) context.Context {
	return context.WithValue(parent, traceIdKey, traceId)
}

func WithDebug(parent context.Context, debug bool) context.Context {
	return context.WithValue(parent, "debug", debug)
}

func UseDB(ctx context.Context) *DB {
	v := ctx.Value(dbKey)
	if v != nil {
		if vv, ok := v.(*DB); ok {
			c := vv.clone()
			c.ctx = ctx
			return c
		}
	}
	return nil
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
	Error        error
	tx           *executor.Tx
	ctx          context.Context
	trace        bool
	debug        bool
	traceId      string
	affect       any
	executor     executor.Executor
	executed     atomic.Bool
	rowsAffected int64
	lastInsertId int64
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
		traceId:  d.traceId,
		affect:   d.affect,
		executed: d.executed,
		executor: nil,
	}
}

func (d *DB) WithContext(ctx context.Context) *DB {
	v := UseDB(ctx)
	if v != nil {
		return v
	}
	c := d.clone()
	c.ctx = ctx
	return c
}

func (d *DB) WithTraceId(traceId string) *DB {
	c := d.clone()
	if c.traceId != "" {
		d.addError(fmt.Errorf("set traceId  repeatedly"))
		return c
	}
	c.traceId = traceId
	return c
}

func (d *DB) Affect(v any) *DB {
	c := d.clone()
	c.affect = v
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
	if d.executed.Swap(true) {
		d.addError(fmt.Errorf("repeat execution"))
		return
	}
	d.addError(d.executor.Execute(d.Logger, d.trace, d.debug, d.affect, func(s executor.Scanner) error {
		d.rowsAffected = s.RowsAffected()
		d.lastInsertId = s.LastInsertId()
		return s.Scan(dest)
	}))
}

func (d *DB) conn() executor.Conn {
	if d.tx != nil {
		return d.tx
	} else {
		conn, err := d.db.Conn(d.context())
		if err != nil {
			d.addError(err)
		}
		return executor.NewDB(conn, d.traceId)
	}
}

func (d *DB) raw(element Element) (raw *executor.Raw, err error) {
	raw, err = element.Raw(d.Dialector.Namer(), "db")
	if err != nil {
		return
	}
	raw.Ctx = d.context()
	return
}

// 执行查询语句
func (d *DB) Query(sql string, params ...executor.Param) *DB {
	c := d.clone()
	raw, err := c.raw(query{sql: sql, params: params})
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(c.conn(), raw)
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
	raw, err := c.raw(exec{sql: sql, params: params})
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(c.conn(), raw)
	c.execute(nil)
	return c
}

// 执行删除操作
func (d *DB) Delete(table string, where Element) *DB {
	c := d.clone()
	raw, err := c.raw(del{table: table, elems: []Element{where}})
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(c.conn(), raw)
	c.execute(nil)
	return c
}

// 执行更新操作
func (d *DB) Update(table string, data map[string]any, where Element, elems ...Element) *DB {
	c := d.clone()
	u := update{table: table, data: data, elems: append([]Element{where}, elems...)}
	raw, err := c.raw(u)
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(c.conn(), raw)
	if !raw.Query {
		c.execute(nil)
	}
	return c
}

// 插入数据
func (d *DB) Insert(table string, data any, elems ...Element) *DB {
	c := d.clone()
	i := &insert{table: table, data: data, elems: elems}
	raw, err := c.raw(i)
	if err != nil {
		c.addError(err)
		return c
	}
	c.executor = executor.NewDefault(c.conn(), raw)
	if !raw.Query {
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

func (d *DB) InsertBatch(table string, batch int, data any, elems ...Element) *DB {
	c := d.clone()

	if batch <= 0 {
		c.addError(fmt.Errorf("expect batch > 0, got %d", batch))
		return c
	}

	if data == nil {
		c.addError(fmt.Errorf("data is nil"))
		return c
	}

	chunks, err := executor.SplitStructSlice(data, batch)
	if err != nil {
		c.addError(err)
		return c
	}

	var raws []*executor.Raw
	var q bool
	for _, v := range chunks {
		i := &insertBatch{
			table: table,
			batch: batch,
			data:  v,
			elems: elems,
		}
		var raw *executor.Raw
		raw, err = c.raw(i)
		if err != nil {
			c.addError(err)
			return c
		}
		if raw.Query {
			q = true
		}
		raws = append(raws, raw)
	}

	conn := c.conn()
	if !conn.IsTx() {
		c = c.Begin()
		conn = c.conn()
	}
	c.setExecutor(executor.NewInsertBatch(c.context(), conn, raws))
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
		c.Error = fmt.Errorf("db executor confilct")
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
			err := v.Execute(c.Logger, d.trace, d.debug, nil, nil)
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
	return d.lastInsertId, d.Error
}

func (d *DB) RowsAffected() (int64, error) {
	return d.rowsAffected, d.Error
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
	c.setExecutor(executor.NewFetchQuery(c.conn(), raw, query.Limit))
	return c.executor.Execute(c.Logger, c.trace, c.debug, nil, func(s *executor.scanner) error {
		return query.Scan(s)
	})
}

func (d *DB) Begin() *DB {
	c := d.clone()
	if c.tx != nil {
		c.addError(fmt.Errorf("tx conflict"))
		return c
	}

	defer func() {
		d.Logger.Trace(c.traceId, true, d.Error, &executor.SQLTrace{
			Trace:        d.trace,
			Debug:        d.debug,
			BeginAt:      time.Now(),
			RawSQL:       "begin",
			PlainSQL:     "begin",
			RowsAffected: 0,
		})
	}()

	tx, err := c.db.Begin()
	if err != nil {
		c.addError(err)
		return c
	}

	if c.traceId == "" {
		c.traceId = fmt.Sprintf("%p", d)
	}

	c.tx = executor.NewTx(tx, c.traceId)

	return c
}

func (d *DB) Commit() *DB {
	if d.tx == nil {
		d.addError(executor.ErrInvalidTransaction)
		return d
	}

	defer func() {
		d.addError(d.tx.Close())
		d.Logger.Trace(d.traceId, true, d.Error, &executor.SQLTrace{
			Trace:        d.trace,
			Debug:        d.debug,
			BeginAt:      time.Now(),
			RawSQL:       "commit",
			PlainSQL:     "commit",
			RowsAffected: 0,
		})
	}()

	d.addError(d.tx.Commit())
	return d
}

func (d *DB) Rollback() *DB {
	if d.tx == nil {
		d.addError(executor.ErrInvalidTransaction)
		return d
	}
	defer func() {
		d.addError(d.tx.Close())
		d.Logger.Trace(d.traceId, true, d.Error, &executor.SQLTrace{
			Trace:        d.trace,
			Debug:        d.debug,
			BeginAt:      time.Now(),
			RawSQL:       "rollback",
			PlainSQL:     "rollback",
			RowsAffected: 0,
		})
	}()
	d.addError(d.tx.Rollback())
	return d
}
