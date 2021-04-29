package engine

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/driver/mysql"
	"github.com/gobatis/gobatis/schema"
	"io/ioutil"
)

func NewEngine(configPath string) (instance *Engine, err error) {
	instance = &Engine{configPath: configPath}
	err = instance.init()
	if err != nil {
		return
	}
	return
}

type Engine struct {
	DB            *sql.DB
	module        string // go.mod module
	configPath    string
	configuration *schema.Configuration
	mappers       *Mappers
}

func (p *Engine) Mappers() *Mappers {
	if p.mappers == nil {
		p.mappers = NewMappers()
	}
	return p.mappers
}

// 加载配置文件
func (p *Engine) loadConfig() (err error) {
	data, err := ioutil.ReadFile(p.configPath)
	if err != nil {
		err = fmt.Errorf("read gobatis.xml error: %s", err)
		return
	}
	p.configuration, err = schema.UnmarshalConfiguration(data)
	if err != nil {
		err = fmt.Errorf("parse gobatis.xml error: %s", err)
		return
	}
	
	if p.configuration.Properties == nil || !p.configuration.Properties.PropertyMap().Has(schema.MODULE) {
		err = fmt.Errorf("gobatis.xml not set module property")
		return
	}
	
	p.module = p.configuration.Module()
	
	return
}

func (p *Engine) init() (err error) {
	
	err = p.loadConfig()
	if err != nil {
		return
	}
	
	err = p.registerMappers()
	if err != nil {
		return
	}
	
	err = p.initDB()
	if err != nil {
		return
	}
	
	return
}

// 注册 Mapper
func (p *Engine) registerMappers() (err error) {
	if p.configuration.Mappers == nil {
		return
	}
	
	for _, v := range p.configuration.Mappers.Children {
		err = p.registerMapper(p.join(v.Resource))
		if err != nil {
			return
		}
	}
	
	return
}

// 初始化引擎
func (p *Engine) initDB() (err error) {
	
	env, err := p.defaultEnvironment()
	if err != nil {
		return
	}
	
	// p.defaultEnvironment 中会限制 env.DataSource != nil
	switch env.DataSource.Driver() {
	case MYSQL:
		p.DB, err = mysql.InitDB(env.DataSource)
		if err != nil {
			return
		}
	default:
		err = fmt.Errorf("gobatis.xml environment:%s data source driver not support", env.Id)
		return
	}
	
	return
}
