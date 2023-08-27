package batis

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gobatis/gobatis/dialector"
	"github.com/gozelle/spew"
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
	db = &DB{
		Error: nil,
		tx:    nil,
	}

	return
}

type DB struct {
	*Config
	*executor
	Error        error
	rowsAffected *int64
	lastInsertId *int64
	tx           *tx
	ctx          context.Context
	debug        bool
	must         bool
	traceId      string
}

func (d *DB) addError(err error) {
	if d.Error == nil {
		d.Error = err
	} else if err != nil {
		d.Error = fmt.Errorf("%v; %w", d.Error, err)
	}
}

func (d *DB) clone() *DB {
	return &DB{Config: d.Config, Error: d.Error}
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

//func (d *DB) SetLogger(logger Logger) {
//	d.logger = logger
//}
//
//func (d *DB) useLogger() Logger {
//	if d.logger == nil {
//		d.logger = DefaultLogger()
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
	e := executor{query: query}
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
	d.executor = &e
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

	var err error
	d.rowsAffected, d.lastInsertId, err = d.executor.exec(dest)
	if err != nil {
		d.addError(err)
		return
	}

	spew.Json("exec done")

	return
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
	if !q {
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
	if !q {
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
	if !q {
		c.exec(nil)
	}
	return c
}

func (d *DB) InsertBatch(table string, batch int, data any, onConflict Element) *DB {
	c := d.clone()
	i := &insertBatch{table: table, batch: batch, data: data, elems: []Element{onConflict}}
	c.prepare(false, i)
	q := i.returning != nil
	c.prepare(q, i)
	if !q {
		c.exec(nil)
	}
	return c
}

func (d *DB) ParallelQuery(queryer ...Queryer) *DB {
	c := d.clone()
	defer func() {
		if c.Error != nil {
			c.log()
		}
	}()
	if len(queryer) == 0 {
		c.Error = fmt.Errorf("no querer")
		return c
	}
	return c
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

func (d *DB) Begin() (*sql.Tx, error) {
	return d.db.Begin()
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

func (d *DB) log() {

}

func LooseDest(dest any, fields ...string) Dest {
	return Dest{loose: true, dest: dest, fields: fields}
}

type Dest struct {
	loose  bool
	dest   any
	fields []string
}
