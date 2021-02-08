package gobatis

import (
	"fmt"
	"github.com/koyeo/gobatis/compiler"
	"io/ioutil"
	"net/http"
)

const CONFIG_FILE = "gobatis.xml"

func NewEngine(resource http.FileSystem) *Engine {
	engine := &Engine{
		configFile: CONFIG_FILE,
		resource:   resource,
	}
	return engine
}

type Engine struct {
	configFile string
	resource   http.FileSystem
}

func (p *Engine) ConfigFile() string {
	return p.configFile
}

func (p *Engine) SetConfigFile(configFile string) {
	p.configFile = configFile
}

func (p *Engine) Init() (err error) {

	return
}

func (p *Engine) parseConfigXML() (err error) {
	if p.resource == nil {
		err = fmt.Errorf("no resource bundle")
		return
	}
	c, err := p.resource.Open(p.configFile)
	if err != nil {
		err = fmt.Errorf("open config.xml error: %s", err)
		return
	}
	defer func() {
		_ = c.Close()
	}()
	d, err := ioutil.ReadAll(c)
	if err != nil {
		err = fmt.Errorf("read config.xml content error: %s", err)
		return
	}

	tokenizer := compiler.NewXMLTokenizer(d)

	parser := compiler.NewXMLParser(tokenizer.Parse())

	return p.registerConfiguration(parser.Parse())
}

func (p *Engine) registerConfiguration(nodes []*compiler.XMLNode) (err error) {

	return
}
