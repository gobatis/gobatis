package batis

import (
	"context"
	"database/sql"
	"fmt"

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
	executor     *executor
	tracer       *tracer
}

func (d *DB) clone() *DB {
	return &DB{
		Error:        d.Error,
		RowsAffected: nil,
		LastInsertId: nil,
		db:           d.db,
		logger:       d.logger,
		tx:           d.tx,
		ctx:          d.ctx,
		debug:        d.debug,
		must:         d.must,
		namer:        d.namer,
		traceId:      "",
		executor:     nil,
		tracer:       nil,
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

func (d *DB) prepare(query bool, elem Element) {
	if d.Error != nil {
		return
	}
	if d.executor != nil {
		d.Error = fmt.Errorf("executor overridden")
		return
	}
	e := executor{query: query}
	if d.tx != nil {
		e.conn = d.tx
		d.tracer.tx = true
	} else {
		e.conn, d.Error = d.db.Conn(d.context())
		if d.Error != nil {
			return
		}
	}
	e.raw, e.params, d.Error = elem.SQL(d.namer, "db")
	if d.Error != nil {
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
		d.Error = fmt.Errorf("no executor")
		return
	}
	d.RowsAffected, d.LastInsertId, d.Error = d.executor.exec(dest)
	if d.Error != nil {
		return
	}
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
	if q {
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
	if q {
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

func (d *DB) Begin() *DB {
	c := d.clone()
	if c.Error != nil {
		c.log()
		return c
	}
	var t *sql.Tx
	t, c.Error = c.db.Begin()
	if c.Error != nil {
		return c
	}
	c.tx = &tx{Tx: t}
	return c
}

func (d *DB) Commit() error {
	return d.tx.Commit()
}

func (d *DB) Rollback() error {
	return d.tx.Rollback()
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
