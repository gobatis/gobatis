package batis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/builder"
	"github.com/gobatis/gobatis/executor"
	"golang.org/x/sync/errgroup"
	"sync"
)

func Open(d Dialector, options ...Option) (db *DB, err error) {
	
	db = &DB{
		//bundle:          nil,
		//fragmentManager: nil,
		db:     nil,
		logger: newLogger(),
		tx:     nil,
		mu:     sync.RWMutex{},
		//stmtMap:         nil,
		err: nil,
	}
	
	db.db, err = d.DB()
	if err != nil {
		return
	}
	
	err = db.Ping()
	if err != nil {
		return
	}
	
	return
}

//func (d *DB) initDB() (err error) {
//	switch d.driver {
//	case postgresql.PGX:
//		d.db, err = postgresql.InitDB(d.dsn)
//	case mysql.MySQL:
//		d.db, err = mysql.InitDB(d.dsn)
//	default:
//		d.db, err = sql.Open(d.driver, d.dsn)
//		if err != nil {
//			err = fmt.Errorf("%s connnet error: %s", d.driver, err)
//			return
//		}
//	}
//	d.dsn = ""
//	return
//}

type DB struct {
	//bundle          Bundle
	//fragmentManager *fragmentManager
	//driver          string
	//dsn             string
	db     *sql.DB
	logger Logger
	tx     *sql.Tx
	mu     sync.RWMutex
	//stmtMap map[string]*Stmt
	ctx   context.Context
	debug bool
	must  bool
	loose bool
	err   error
}

func (d *DB) fork() *DB {
	return &DB{
		//fragmentManager: nil,
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
	dd := d.fork()
	dd.ctx = ctx
	return dd
}

func (d *DB) Debug() *DB {
	dd := d.fork()
	dd.debug = true
	return dd
}

func (d *DB) Must() *DB {
	dd := d.fork()
	dd.debug = true
	return dd
}

//func (d *DB) SetTag(tag string) {
//	executor.reflect_tag = tag
//}

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

//func (d *DB) Init(bundle Bundle) (err error) {
//
//	if d.logger == nil {
//		d.logger = _log.NewStdLogger()
//		d.logger.SetLevel(InfoLevel)
//	}
//
//	d.bundle = bundle
//	err = d.parseBundle()
//	if err != nil {
//		return
//	}
//	//err = d.master.initDB()
//	//d.master.logger = d.logger
//	//if err != nil {
//	//	err = fmt.Errorf("init master db error: %s", err)
//	//	return
//	//}
//	return
//}

func (d *DB) Close() {
	//if d.fragmentManager != nil {
	//	for _, v := range d.fragmentManager.all() {
	//		if v._stmt != nil {
	//			err := v._stmt.Close()
	//			if err != nil {
	//				d.logger.Errorf("[gobatis] close stmt error: %s", err)
	//			}
	//		}
	//	}
	//}
	_ = d.db.Close()
}

//func (d *DB) parseBundle() (err error) {
//	err = d.parseConfig()
//	if err != nil {
//		return
//	}
//	
//	err = d.parseMappers()
//	if err != nil {
//		return
//	}
//	return
//}
//
//func (d *DB) BindMapper(bundle Bundle, ptr ...interface{}) (err error) {
//	for _, v := range ptr {
//		err = d.bindMapper(v)
//		if err != nil {
//			return
//		}
//	}
//	return
//}

//func (d *DB) bindMapper(mapper interface{}) (err error) {
//	
//	rv := reflect.ValueOf(mapper)
//	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
//		return fmt.Errorf("exptect *struct, got: %s", rv.Type())
//	}
//	rv = rv.Elem()
//	rt := rv.Type()
//	for i := 0; i < rt.NumField(); i++ {
//		if rv.Field(i).Kind() != reflect.Func {
//			continue
//		}
//		must := false
//		stmt := false
//		id := rt.Field(i).Name
//		if strings.HasPrefix(id, must_prefix) {
//			id = strings.TrimPrefix(id, must_prefix)
//			must = true
//		}
//		if strings.HasSuffix(id, stmt_suffix) {
//			id = strings.TrimSuffix(id, stmt_suffix)
//			stmt = true
//		}
//		if strings.HasSuffix(id, tx_suffix) {
//			id = strings.TrimSuffix(id, tx_suffix)
//		}
//		m, ok := d.fragmentManager.get(id)
//		if !ok {
//			if must {
//				return fmt.Errorf("%s.(Must)%s statement not defined", rt.Name(), id)
//			}
//			return fmt.Errorf("%s.%s statement not defined", rt.Name(), id)
//		}
//		m = m.fork()
//		m.must = must
//		m.stmt = stmt
//		m.id = rt.Field(i).Name
//		ft := rv.Field(i).Type()
//		m.checkParameter(ft, rt.Name(), rv.Type().Field(i).Name)
//		m.checkResult(ft, rt.Name(), rv.Type().Field(i).Name)
//		m.proxy(rv.Field(i))
//	}
//	return
//}

//func (d *DB) parseConfig() (err error) {
//	if d.bundle == nil {
//		err = fmt.Errorf("no bundle")
//		return
//	}
//	
//	f, err := d.bundle.Open(config_xml)
//	if err != nil {
//		err = nil
//		return
//	}
//	_ = f.Close()
//	
//	bs, err := d.readBundleFile(config_xml)
//	if err != nil {
//		return
//	}
//	d.logger.Infof("[gobatis] register fragment: gobatis.xml")
//	err = parseConfig(d, config_xml, string(bs))
//	if err != nil {
//		return
//	}
//	return
//}

//func (d *DB) readBundleFile(path string) (bs []byte, err error) {
//	file, err := d.bundle.Open(path)
//	if err != nil {
//		err = fmt.Errorf("open %s error: %s", path, err)
//		return
//	}
//	defer func() {
//		_ = file.Close()
//	}()
//	
//	bs, err = io.ReadAll(file)
//	if err != nil {
//		err = fmt.Errorf("read %s content error: %s", path, err)
//		return
//	}
//	return
//}

//func (d *DB) parseMappers() (err error) {
//	files, err := d.walkMappers("/")
//	if err != nil {
//		return
//	}
//	for _, v := range files {
//		var bs []byte
//		bs, err = d.readBundleFile(v)
//		if err != nil {
//			return
//		}
//		d.logger.Infof("register fragment: %s.xml", v)
//		err = parseMapper(d, v, string(bs))
//		if err != nil {
//			return
//		}
//	}
//	return
//}

//func (d *DB) walkMappers(root string) (files []string, err error) {
//	handle, err := d.bundle.Open(root)
//	if err != nil {
//		return
//	}
//	defer func() {
//		_ = handle.Close()
//	}()
//	res, err := handle.Readdir(-1)
//	for _, v := range res {
//		path := filepath.Join(root, v.Name())
//		if v.IsDir() {
//			var _files []string
//			_files, err = d.walkMappers(path)
//			if err != nil {
//				return
//			}
//			files = append(files, _files...)
//		} else {
//			if path != filepath.Join("/", config_xml) {
//				files = append(files, path)
//			}
//		}
//	}
//	return
//}

//func (d *DB) addFragment(file string, ctx antlr.ParserRuleContext, id string, node *xmlNode) {
//	
//	f, err := parseFragment(d, file, id, node)
//	if err != nil {
//		return
//	}
//	err = d.fragmentManager.add(f)
//	if err != nil {
//		throw(file, ctx, parseMapperErr).with(err)
//	}
//	return
//}

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
	dd := d.fork()
	dd.loose = true
	return dd
}

//func (d *DB) Prepare(sql string, params ...executor.NameValue) *Stmt {
//	return &Stmt{}
//}

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
		return executor.WithErrScanner(err)
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
		return executor.WithErrScanner(err)
	}
	
	return *s
}

func (d *DB) Execute(sql string, params ...executor.NameValue) executor.Scanner {
	return d.exec(executor.Exec, sql, params)
	
}

func (d *DB) Delete(table string, where Element) executor.Scanner {
	s := fmt.Sprintf("delete from %s %s", table, where.SQL())
	return d.exec(executor.Exec, s, where.Params())
}

func (d *DB) Update(table string, data map[string]any, where Element) executor.Scanner {
	
	s := fmt.Sprintf("update %s set %s", table, where.SQL())
	return d.exec(executor.Exec, s, where.Params())
}

func (d *DB) Insert(table string, data any, onConflict ...Element) executor.Scanner {
	
	s := fmt.Sprintf("insert into %s", table)
	return d.exec(executor.Exec, s, nil)
}

func (d *DB) InsertBatch(table string, data any, batch int, onConflict ...Element) executor.Scanner {
	s := fmt.Sprintf("insert into %s", table)
	return d.exec(executor.Exec, s, nil)
}

func (d *DB) Fetch(sql string, params ...executor.NameValue) <-chan executor.Executor {
	return nil
}

func (d *DB) Begin() *DB {
	if d.err != nil {
		return d
	}
	d.tx, d.err = d.db.Begin()
	return d
}
