package batis

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gobatis/gobatis/dialector"
	"github.com/gobatis/gobatis/logger"
	"github.com/gobatis/gobatis/parser"
	"go.uber.org/atomic"
)

const (
	dbKey      = "GOBATIS_DB"
	traceIdKey = "GOBATIS_TRACE_ID"
	debugKey   = "GOBATIS_DEBUG"
	space      = " "
)

const (
	methodQuery          = "Query"
	methodExec           = "Exec"
	methodDelete         = "Delete"
	methodUpdate         = "Update"
	methodInsert         = "Insert"
	methodInsertBatch    = "InsertBatch"
	methodParallelQuery  = "ParallelQuery"
	methodPagingQuery    = "PagingQuery"
	methodFetchQuery     = "FetchQuery"
	methodAssociateQuery = "AssociateQuery"
)

func WithTx(parent context.Context, tx *DB) context.Context {
	if tx.tx != nil {
		return context.WithValue(parent, dbKey, tx)
	}
	return context.WithValue(parent, dbKey, tx.Begin())
}

func WithTraceId(parent context.Context, traceId string) context.Context {
	return context.WithValue(parent, traceIdKey, traceId)
}

func WithDebug(parent context.Context, debug bool) context.Context {
	return context.WithValue(parent, debugKey, debug)
}

func Open(d dialector.Dialector, options ...Option) (db *DB, err error) {
	config := Config{
		//CreateBatchSize: 10,
		Plugins: nil,
		NowFunc: func() time.Time {
			return time.Now()
		},
		Dialector: d,
		Logger:    logger.Default,
		ColumnTag: "db",
		db:        nil,
	}
	config.db, err = d.DB()
	if err != nil {
		return
	}
	db = &DB{Config: &config, Error: nil}
	return
}

type DB struct {
	*Config
	Error    error
	tx       *connTx
	ctx      context.Context
	trace    bool
	debug    bool
	traceId  string
	affect   any
	executor executor
	executed atomic.Bool
	result   sql.Result
}

func (d *DB) addError(err error) {
	d.Error = parser.AddError(d.Error, err)
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
		executor: d.executor,
	}
}

func (d *DB) WithContext(ctx context.Context) *DB {
	c := d.clone()
	c.ctx = ctx
	if vv := c.ctx.Value(dbKey); vv != nil {
		db, ok := vv.(*DB)
		if ok {
			c = db.clone()
			c.ctx = ctx
		}
	}

	if c.executed.Load() {
		c.addError(fmt.Errorf("cannot apply the WithContext method %w", ErrApplyMethodOnExecutedDBChain))
		return c
	}

	if vv := c.ctx.Value(traceIdKey); vv != nil {
		traceId, ok := vv.(string)
		if ok {
			c.traceId = traceId
		}
	}
	if vv := c.ctx.Value(debugKey); vv != nil {
		debug, ok := vv.(bool)
		if ok {
			c.debug = debug
		}
	}
	return c
}

func (d *DB) WithTraceId(traceId string) *DB {
	c := d.clone()
	if c.executed.Load() {
		c.addError(fmt.Errorf("cannot apply the WithTraceId method %w", ErrApplyMethodOnExecutedDBChain))
		return c
	}
	if c.traceId != "" {
		d.addError(fmt.Errorf("set traceId repeatedly"))
		return c
	}
	c.traceId = traceId
	return c
}

type Session struct {
	Logger    logger.Logger
	ColumnTag string
}

func (d *DB) Session(s *Session) *DB {
	c := d.clone()
	if c.executed.Load() {
		c.addError(fmt.Errorf("cannot apply the Session method %w", ErrApplyMethodOnExecutedDBChain))
		return c
	}
	c.Config = c.Config.clone()
	if s.Logger != nil {
		c.Logger = s.Logger
	}
	if s.ColumnTag != "" {
		c.ColumnTag = s.ColumnTag
	}
	return c
}

func (d *DB) Affect(v any) *DB {
	c := d.clone()
	if c.executed.Load() {
		c.addError(fmt.Errorf("cannot apply the Affect method %w", ErrApplyMethodOnExecutedDBChain))
		return c
	}
	c.affect = v
	return c
}

func (d *DB) Trace() *DB {
	c := d.clone()
	if c.executed.Load() {
		c.addError(fmt.Errorf("cannot apply the Trace method %w", ErrApplyMethodOnExecutedDBChain))
		return c
	}
	c.trace = true
	return c
}

func (d *DB) Debug() *DB {
	c := d.clone()
	if c.executed.Load() {
		c.addError(fmt.Errorf("cannot apply the Debug method %w", ErrApplyMethodOnExecutedDBChain))
		return c
	}
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

func (d *DB) prepareContext() context.Context {
	if d.ctx == nil {
		return context.Background()
	}
	return d.ctx
}

func (d *DB) duplicatedExecutor() bool {
	if d.executor != nil {
		d.addError(fmt.Errorf("%w with method: %s", ErrExecutorConflict, d.executor.method()))
		return false
	}
	return true
}

func (d *DB) setExecutor(e executor) {
	if d.duplicatedExecutor() {
		d.executor = e
	}
}

func (d *DB) execute() {
	if d.Error != nil {
		return
	}
	if d.executor == nil {
		d.addError(fmt.Errorf("no executor in db chain"))
		return
	}

	if d.executed.Swap(true) {
		d.addError(fmt.Errorf("repeat execution"))
		return
	}
	r, err := d.executor.execute()
	if err != nil {
		d.addError(err)
		return
	}
	d.result = r
}

func (d *DB) prepareDefaultExecutor(method string, r *raw) *defaultExecutor {
	return &defaultExecutor{
		name:   method,
		raw:    r,
		ctx:    d.prepareContext(),
		conn:   d.prepareConn(),
		logger: d.Logger,
		pos:    "",
		trace:  d.trace,
		debug:  d.debug,
		affect: d.affect,
		scanner: &defaultScanner{
			columnTag: d.ColumnTag,
		},
		scan: nil,
	}
}

func (d *DB) prepareConn() conn {
	if d.tx != nil {
		return d.tx
	} else {
		cn, err := d.db.Conn(d.prepareContext())
		if err != nil {
			d.addError(err)
		}
		return newDBConn(cn, d.traceId)
	}
}

func (d *DB) prepareRaw(elem Elem) (r *raw, err error) {
	r, err = elem.Raw(d.Dialector.Namer(), d.ColumnTag)
	if err != nil {
		err = fmt.Errorf("%w, %s", PrepareSQLRawErr, err)
		return
	}
	return
}

// 扫描结果集
func (d *DB) Scan(dest any, ignore ...string) *DB {
	if d.executor == nil {
		d.addError(fmt.Errorf("no executor found in db chain"))
		return d
	}
	switch d.executor.method() {
	case methodQuery, methodExec, methodUpdate, methodDelete, methodInsert, methodInsertBatch:
		d.executor.setScan(func(s scanner) error {
			s.setDest(dest, ignore...)
			e := s.scan()
			if e != nil {
				return e
			}
			return nil
		})
		d.execute()
	default:
		d.addError(fmt.Errorf("moethod: %s unsuppoted Scan method", d.executor.method()))
	}
	return d
}

func (d *DB) LastInsertId() (int64, error) {
	if d.Error != nil {
		return 0, d.Error
	}
	if d.result == nil {
		return 0, ErrNoSQLResultExists
	}
	return d.result.LastInsertId()
}

func (d *DB) RowsAffected() (int64, error) {
	if d.Error != nil {
		return 0, d.Error
	}
	if d.result == nil {
		return 0, ErrNoSQLResultExists
	}
	return d.result.RowsAffected()
}

// 执行查询语句
func (d *DB) Query(sql string, params ...NameValue) *DB {
	c := d.clone()
	r, err := c.prepareRaw(query{sql: sql, params: params})
	if err != nil {
		c.addError(err)
		return c
	}
	c.setExecutor(c.prepareDefaultExecutor(methodQuery, r))
	return c
}

// 执行 SQL 语句
func (d *DB) Exec(sql string, params ...NameValue) *DB {
	c := d.clone()
	r, err := c.prepareRaw(exec{sql: sql, params: params})
	if err != nil {
		c.addError(err)
		return c
	}
	e := c.prepareDefaultExecutor(methodExec, r)
	e.withTx = true
	c.setExecutor(e)
	c.execute()
	return c
}

// 执行删除操作
func (d *DB) Delete(table string, where Elem, elems ...Elem) *DB {
	c := d.clone()
	r, err := c.prepareRaw(del{table: table, elems: append([]Elem{where}, elems...)})
	if err != nil {
		c.addError(err)
		return c
	}
	e := c.prepareDefaultExecutor(methodDelete, r)
	e.withTx = true
	c.setExecutor(e)
	if !r.Query {
		c.execute()
	}
	return c
}

// 执行更新操作
func (d *DB) Update(table string, data map[string]any, where Elem, elems ...Elem) *DB {
	c := d.clone()
	u := update{table: table, data: data, elems: append([]Elem{where}, elems...)}
	r, err := c.prepareRaw(u)
	if err != nil {
		c.addError(err)
		return c
	}
	e := c.prepareDefaultExecutor(methodUpdate, r)
	e.withTx = true
	c.setExecutor(e)
	if !r.Query {
		c.execute()
	}
	return c
}

// 插入数据
func (d *DB) Insert(table string, data any, elems ...Elem) *DB {
	c := d.clone()
	i := &insert{table: table, data: data, elems: elems}
	r, err := c.prepareRaw(i)
	if err != nil {
		c.addError(err)
		return c
	}
	e := c.prepareDefaultExecutor(methodInsert, r)
	e.withTx = true
	c.setExecutor(e)
	if !r.Query {
		c.execute()
	}
	return c
}

func (d *DB) InsertBatch(table string, batch int, data any, elems ...Elem) *DB {
	c := d.clone()

	if !c.duplicatedExecutor() {
		return c
	}

	if batch <= 0 {
		c.addError(fmt.Errorf("%w, got %d", InvalidInsertBatchBatchErr, batch))
		return c
	}

	if data == nil {
		c.addError(InvalidInsertBatchDataErr)
		return c
	}

	chunks, err := SplitStructSlice(data, batch)
	if err != nil {
		c.addError(fmt.Errorf("%w, %s", InvalidInsertBatchDataTypeErr, err))
		return c
	}

	var raws []*raw
	var q bool
	for _, v := range chunks {
		i := &insertBatch{
			table: table,
			batch: batch,
			data:  v,
			elems: elems,
		}
		var r *raw
		r, err = c.prepareRaw(i)
		if err != nil {
			c.addError(err)
			return c
		}
		if r.Query {
			q = true
		}
		raws = append(raws, r)
	}

	c.setExecutor(&insertBatchExecutor{
		raws:   raws,
		name:   methodInsertBatch,
		ctx:    c.prepareContext(),
		conn:   c.prepareConn(),
		logger: c.Logger,
		trace:  c.trace,
		debug:  c.debug,
		affect: c.affect,
		scanner: &insertBatchScanner{
			columnTag: c.ColumnTag,
		},
	})
	if !q {
		c.execute()
	}
	return c
}

func (d *DB) insertBatch() {

}

func (d *DB) ParallelQuery(queryer ...ParallelQuery) *DB {
	c := d.clone()
	if len(queryer) == 0 {
		c.addError(NoParallelQueryerErr)
		return c
	}
	if c.tx != nil {
		c.addError(fmt.Errorf("method %s %w", methodParallelQuery, ErrNotCompatibleWithTransactionMode))
		return c
	}
	if c.affect != nil {
		c.addError(fmt.Errorf("mehod %s %w", methodParallelQuery, ErrNotSupportAffectConstraint))
		return c
	}

	pos := logger.CallFuncPos(0)
	c.setExecutor(&parallelQueryExecutor{
		queries:   queryer,
		name:      methodParallelQuery,
		ctx:       c.prepareContext(),
		conn:      c.prepareConn,
		logger:    c.Logger,
		pos:       pos,
		trace:     c.trace,
		debug:     c.debug,
		columnTag: c.ColumnTag,
	})
	c.execute()

	return c
}

func (d *DB) PagingQuery(query PagingQuery) *DB {
	c := d.clone()

	if c.tx != nil {
		c.addError(fmt.Errorf("method %s %w", methodPagingQuery, ErrNotCompatibleWithTransactionMode))
		return c
	}
	if c.affect != nil {
		c.addError(fmt.Errorf("mehod %s %w", methodPagingQuery, ErrNotSupportAffectConstraint))
		return c
	}

	pos := logger.CallFuncPos(0)

	c.setExecutor(&pagingQueryExecutor{
		query: query,
		name:  methodPagingQuery,
		ctx:   c.prepareContext(),
		conn: func() (conn, error) {
			return c.prepareConn(), c.Error
		},
		logger:    c.Logger,
		pos:       pos,
		trace:     c.trace,
		debug:     c.debug,
		columnTag: c.ColumnTag,
	})
	c.execute()

	return c
}

func (d *DB) AssociateQuery(query AssociateQuery) *DB {

	c := d.clone()

	if c.affect != nil {
		c.addError(fmt.Errorf("mehod %s %w", methodAssociateQuery, ErrNotSupportAffectConstraint))
		return c
	}

	r := &raw{
		Query: true,
		SQL:   query.SQL,
		Vars:  query.Params,
	}

	e := c.prepareDefaultExecutor(methodAssociateQuery, r)
	e.scanner = &associateScanner{
		columnTag: c.ColumnTag,
	}
	e.scan = func(s scanner) error {
		return query.Scan(s.(AssociateScanner))
	}

	c.setExecutor(e)
	c.execute()

	return c
}

func (d *DB) FetchQuery(query FetchQuery) *DB {

	c := d.clone()

	if c.affect != nil {
		c.addError(fmt.Errorf("mehod %s %w", methodFetchQuery, ErrNotSupportAffectConstraint))
		return c
	}

	if c.traceId == "" {
		c.traceId = fmt.Sprintf("TID%d", time.Now().UnixNano())
	}

	e := c.prepareDefaultExecutor(methodFetchQuery, &raw{
		Query: true,
		SQL:   query.SQL,
		Vars:  query.Params,
	})
	e.scan = func(s scanner) error {
		return query.Scan(s.(Scanner))
	}
	c.setExecutor(&fetchQueryExecutor{columnTag: c.ColumnTag, limit: query.Batch, defaultExecutor: e})
	c.execute()

	return c
}

func (d *DB) Begin() *DB {
	return d.BeginWithOption(nil)
}

func (d *DB) BeginWithOption(opts *sql.TxOptions) *DB {
	c := d.clone()
	if c.tx != nil {
		c.addError(fmt.Errorf("there is already a transaction"))
		return c
	}
	defer func() {
		d.Logger.Trace("", c.traceId, true, d.Error, &logger.SQLTrace{
			Trace:    d.trace,
			Debug:    d.debug,
			BeginAt:  time.Now(),
			RawSQL:   "begin",
			PlainSQL: "begin",
		})
	}()
	tx, err := c.db.BeginTx(d.prepareContext(), opts)
	if err != nil {
		c.addError(err)
		return c
	}
	if c.traceId == "" {
		c.traceId = fmt.Sprintf("%p", d)
	}
	c.tx = newTx(tx, c.traceId)
	return c
}

func (d *DB) Commit() *DB {
	if d.tx == nil {
		d.addError(ErrInvalidTransaction)
		return d
	}

	defer func() {
		d.addError(d.tx.Close())
		d.Logger.Trace("", d.traceId, true, d.Error, &logger.SQLTrace{
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
		d.addError(ErrInvalidTransaction)
		return d
	}
	defer func() {
		d.addError(d.tx.Close())
		d.Logger.Trace("", d.traceId, true, d.Error, &logger.SQLTrace{
			Trace:    d.trace,
			Debug:    d.debug,
			BeginAt:  time.Now(),
			RawSQL:   "rollback",
			PlainSQL: "rollback",
		})
	}()
	d.addError(d.tx.Rollback())
	return d
}
