package gobatis

import (
	"database/sql"
	"fmt"
	"github.com/gozelle/spew"
	"reflect"
	"strings"
	"sync"
)

func Open(d Dialector, options ...Option) (db *DB, err error) {
	
	db = &DB{
		//bundle:          nil,
		fragmentManager: nil,
		db:              nil,
		logger:          newLogger(),
		tx:              nil,
		mu:              sync.RWMutex{},
		stmtMap:         nil,
		error:           nil,
	}
	
	db.db, err = d.DB()
	if err != nil {
		return
	}
	
	//err = db.Ping()
	//if err != nil {
	//	return
	//}
	
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
	fragmentManager *fragmentManager
	//driver          string
	//dsn             string
	db      *sql.DB
	logger  Logger
	tx      *sql.Tx
	mu      sync.RWMutex
	stmtMap map[string]*Stmt
	error   error
}

func (d *DB) SetTag(tag string) {
	reflect_tag = tag
}

func (d *DB) UseJsonTag() {
	reflect_tag = "json"
}

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
	if d.fragmentManager != nil {
		for _, v := range d.fragmentManager.all() {
			if v._stmt != nil {
				err := v._stmt.Close()
				if err != nil {
					d.logger.Errorf("[gobatis] close stmt error: %s", err)
				}
			}
		}
	}
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

func (d *DB) bindMapper(mapper interface{}) (err error) {
	
	rv := reflect.ValueOf(mapper)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("exptect *struct, got: %s", rv.Type())
	}
	rv = rv.Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		if rv.Field(i).Kind() != reflect.Func {
			continue
		}
		must := false
		stmt := false
		id := rt.Field(i).Name
		if strings.HasPrefix(id, must_prefix) {
			id = strings.TrimPrefix(id, must_prefix)
			must = true
		}
		if strings.HasSuffix(id, stmt_suffix) {
			id = strings.TrimSuffix(id, stmt_suffix)
			stmt = true
		}
		if strings.HasSuffix(id, tx_suffix) {
			id = strings.TrimSuffix(id, tx_suffix)
		}
		m, ok := d.fragmentManager.get(id)
		if !ok {
			if must {
				return fmt.Errorf("%s.(Must)%s statement not defined", rt.Name(), id)
			}
			return fmt.Errorf("%s.%s statement not defined", rt.Name(), id)
		}
		m = m.fork()
		m.must = must
		m.stmt = stmt
		m.id = rt.Field(i).Name
		ft := rv.Field(i).Type()
		m.checkParameter(ft, rt.Name(), rv.Type().Field(i).Name)
		m.checkResult(ft, rt.Name(), rv.Type().Field(i).Name)
		m.proxy(rv.Field(i))
	}
	return
}

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

func (d *DB) Prepare(sql string, params ...NameValue) *Stmt {
	return &Stmt{}
}

func (d *DB) Query(ctx Context, sql string, params ...NameValue) (scanner Scanner) {
	
	var err error
	defer func() {
		if err != nil {
			scanner.err = err
		}
	}()
	
	node, err := parseSQL("test.file", fmt.Sprintf("<sql>%s</sql>", sql))
	if err != nil {
		return
	}
	
	frag := &fragment{
		db:   nil,
		id:   "id",
		node: node,
		in: []*param{{
			name:  "age",
			_type: "int",
			slice: false,
		}},
		out:             nil,
		resultAttribute: 0,
		must:            false,
		stmt:            false,
		_stmt:           nil,
	}
	
	s, exprs, vars, dynamic, err := frag.parseStatement(reflect.ValueOf(10))
	if err != nil {
		return
	}
	spew.Json(s, exprs, vars, dynamic)
	//c := &caller{fragment: frag, args: nil, logger: newLogger()}
	//
	//c.call()
	
	return
}

func (d *DB) Build(ctx Context, b *Builder) (s Scanner) {
	
	//var sqls []string
	//var params []NameValue
	//for _, v := range elements {
	//	sqls = append(sqls, v.SQL())
	//	params = append(params, v.Params()...)
	//}
	
	//walkXMLNodes()
	
	//var f *fragment
	//f, d.error = parseFragment(d, none, anonymous, &xmlNode{
	//	File:       "",
	//	Name:       "",
	//	Text:       "",
	//	Attributes: nil,
	//	Nodes:      nil,
	//	nodesCount: nil,
	//	start:      nil,
	//	ctx:        nil,
	//	textOnly:   false,
	//})
	//if d.error != nil {
	//	return
	//}
	////f.call()
	//_ 
	return
}

func (d *DB) Execute(ctx Context, sql string, params ...NameValue) (scanner Scanner) {
	return
}

func (d *DB) Delete(ctx Context, table string, where Element) (scanner Scanner) {
	
	return
}

func (d *DB) Update(ctx Context, table string, data map[string]any, where Element) (scanner Scanner) {
	
	return
}

func (d *DB) Insert(ctx Context, table string, data any, onConflict ...Element) (scanner Scanner) {
	
	return
}

func (d *DB) InsertBatch(ctx Context, table string, data any, batch int, onConflict ...Element) (scanner Scanner) {
	
	return
}

func (d *DB) Begin() *DB {
	if d.error != nil {
		return d
	}
	d.tx, d.error = d.db.Begin()
	return d
}
