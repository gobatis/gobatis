package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/driver/mysql"
	"github.com/gobatis/gobatis/driver/postgresql"
	"github.com/koyeo/_log"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
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
	slaves          []*DB
	logger          Logger
	bundle          http.FileSystem
	fragmentManager *fragmentManager
}

func (p *Engine) Master() *DB {
	return p.master
}

func (p *Engine) SetTag(tag string) {
	reflect_tag = tag
}

func (p *Engine) UseJsonTag() {
	reflect_tag = "json"
}

func (p *Engine) SetLogLevel(level Level) {
	p.logger.SetLevel(level)
}

func (p *Engine) SetLogger(logger Logger) {
	p.logger = logger
	if p.master != nil {
		p.master.logger = logger
	}
	for _, v := range p.slaves {
		v.logger = logger
	}
}

func (p *Engine) Init(bundle Bundle) (err error) {
	
	if p.logger == nil {
		p.logger = _log.NewStdLogger()
		p.logger.SetLevel(InfoLevel)
	}
	
	p.bundle = bundle
	err = p.parseBundle()
	if err != nil {
		return
	}
	err = p.master.initDB()
	p.master.logger = p.logger
	if err != nil {
		err = fmt.Errorf("init master db error: %s", err)
		return
	}
	return
}

func (p *Engine) Close() {
	if p.fragmentManager != nil {
		for _, v := range p.fragmentManager.all() {
			if v._stmt != nil {
				err := v._stmt.Close()
				if err != nil {
					p.logger.Errorf("[gobatis] close stmt error: %s", err)
				}
			}
		}
	}
	for _, v := range p.slaves {
		err := v.Close()
		if err != nil {
			p.logger.Errorf("[gobatis] close slave db error: %s", err)
		}
	}
	if p.master != nil {
		err := p.master.Close()
		if err != nil {
			p.logger.Errorf("[gobatis] close master db error: %s", err)
		}
	}
}

func (p *Engine) SQL(name string, args ...interface{}) {

}

func (p *Engine) Call(name string, args ...interface{}) {
	//f, ok := p.fragmentManager.get(name)
	//if !ok {
	//	panic(fmt.Errorf("method '%s' not exist", name))
	//}
	//return &caller{fragment: f, args: args, logger: p.logger}
}

//func (p *Engine) Call(name string, args ...reflect.Value) *caller {
//	f, ok := p.fragmentManager.get(name)
//	if !ok {
//		panic(fmt.Errorf("method '%s' not exist", name))
//	}
//	return &caller{fragment: f, args: args, logger: p.logger}
//}

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
		m, ok := p.fragmentManager.get(id)
		if !ok {
			if must {
				return fmt.Errorf("%s.(Must)%s statement not defined", rt.Name(), id)
			}
			return fmt.Errorf("%s.%s statement not defined", rt.Name(), id)
		}
		m.must = must
		m.stmt = stmt
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
	
	f, err := p.bundle.Open(config_xml)
	if err != nil {
		err = nil
		return
	}
	_ = f.Close()
	
	bs, err := p.readBundleFile(config_xml)
	if err != nil {
		return
	}
	p.logger.Infof("[gobatis] register fragment: gobatis.xml")
	err = parseConfig(p, config_xml, string(bs))
	if err != nil {
		return
	}
	return
}

func (p *Engine) readBundleFile(path string) (bs []byte, err error) {
	file, err := p.bundle.Open(path)
	if err != nil {
		err = fmt.Errorf("open %s error: %s", path, err)
		return
	}
	defer func() {
		_ = file.Close()
	}()
	
	bs, err = ioutil.ReadAll(file)
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
		var bs []byte
		bs, err = p.readBundleFile(v)
		if err != nil {
			return
		}
		p.logger.Infof("[gobatis] register fragment: %s.xml", v)
		err = parseMapper(p, v, string(bs))
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
