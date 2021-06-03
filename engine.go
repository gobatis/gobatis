package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/driver/mysql"
	"github.com/gobatis/gobatis/driver/postgresql"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
)

func NewPostgresql(dsn string) *Engine {
	return NewEngine(NewDB(postgresql.PGX, dsn))
}

func NewMySQL(dsn string) *Engine {
	return NewEngine(NewDB(mysql.MySQL, dsn))
}

func NewEngine(db *DB) *Engine {
	engine := &Engine{master: db, fragmentManager: newMethodManager()}
	return engine
}

type Engine struct {
	master          *DB
	bundle          http.FileSystem
	slaves          []*DB
	fragmentManager *fragmentManager
}

func (p *Engine) Master() *DB {
	return p.master
}

func (p *Engine) BindSQL(bundle http.FileSystem) {
	p.bundle = bundle
}

func (p *Engine) SetTag(tag string) {
	reflect_tag = tag
}

func (p *Engine) SetLogLevel(level LogLevel) {
	_level = level
}

func (p *Engine) SetLogger(logger Logger) {
	_logger = logger
}

func (p *Engine) Init() (err error) {
	err = p.parseBundle()
	if err != nil {
		return
	}
	err = p.master.initDB()
	if err != nil {
		err = fmt.Errorf("init master db error: %s", err)
		return
	}
	return
}

func (p *Engine) Close() {
	if p.master != nil {
		_ = p.master.Close()
	}
	for _, v := range p.slaves {
		_ = v.Close()
	}
}

func (p *Engine) Call(name string, args ...reflect.Value) *caller {
	f, ok := p.fragmentManager.get(name)
	if !ok {
		panic(fmt.Errorf("method '%s' not exist", name))
	}
	return &caller{fragment: f, params: args}
}

func (p *Engine) parseBundle() (err error) {
	err = p.parseConfig()
	if err != nil {
		return
	}
	
	err = p.parseMappers()
	if err != nil {
		return
	}
	return
}

func (p *Engine) BindMapper(ptr ...interface{}) (err error) {
	for _, v := range ptr {
		err = p.bindMapper(v)
		if err != nil {
			return
		}
	}
	return
}

func (p *Engine) bindMapper(mapper interface{}) (err error) {
	defer func() {
		e := recover()
		err = castRecoverError("", e)
	}()
	
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
		id := rt.Field(i).Name
		m, ok := p.fragmentManager.get(id)
		if !ok {
			return fmt.Errorf("%s.%s not defined", rt.Name(), id)
		}
		ft := rv.Field(i).Type()
		m.checkParameter(ft, rt.Name(), rv.Type().Field(i).Name)
		m.checkResult(ft, rt.Name(), rv.Type().Field(i).Name)
		m.proxy(rv.Field(i))
	}
	return
}

func (p *Engine) parseConfig() (err error) {
	if p.bundle == nil {
		err = fmt.Errorf("no bundle")
		return
	}
	d, err := p.readBundleFile(config_xml)
	if err != nil {
		return
	}
	Infof("[gobatis] register fragment: gobatis.xml")
	err = parseConfig(p, config_xml, string(d))
	if err != nil {
		return
	}
	return
}

func (p *Engine) readBundleFile(path string) (d []byte, err error) {
	c, err := p.bundle.Open(path)
	if err != nil {
		err = fmt.Errorf("open %s error: %s", path, err)
		return
	}
	defer func() {
		_ = c.Close()
	}()
	d, err = ioutil.ReadAll(c)
	if err != nil {
		err = fmt.Errorf("read %s content error: %s", path, err)
		return
	}
	return
}

func (p *Engine) parseMappers() (err error) {
	files, err := p.walkMappers("/")
	if err != nil {
		return
	}
	for _, v := range files {
		var d []byte
		d, err = p.readBundleFile(v)
		if err != nil {
			return
		}
		Infof("[gobatis] register fragment: %s.xml", v)
		err = parseMapper(p, v, string(d))
		if err != nil {
			return
		}
	}
	return
}

func (p *Engine) walkMappers(root string) (files []string, err error) {
	handle, err := p.bundle.Open(root)
	if err != nil {
		return
	}
	defer func() {
		_ = handle.Close()
	}()
	res, err := handle.Readdir(-1)
	for _, v := range res {
		path := filepath.Join(root, v.Name())
		if v.IsDir() {
			var _files []string
			_files, err = p.walkMappers(path)
			if err != nil {
				return
			}
			files = append(files, _files...)
		} else {
			if path != filepath.Join("/", config_xml) {
				files = append(files, path)
			}
		}
	}
	return
}

func (p *Engine) addFragment(file string, ctx antlr.ParserRuleContext, id string, node *xmlNode) {
	
	m, err := parseFragment(p.master, file, id, node)
	if err != nil {
		return
	}
	err = p.fragmentManager.add(m)
	if err != nil {
		throw(file, ctx, parseMapperErr).with(err)
	}
	return
}
