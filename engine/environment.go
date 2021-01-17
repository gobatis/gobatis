package engine

import (
	"fmt"
	"github.com/koyeo/mybatis.go/schema"
)

func (p *Engine) defaultEnvironment() (env *schema.Environment, err error) {

	if p.configuration.Environments == nil {
		err = fmt.Errorf("mybatis-config.xml not set <environments>")
		return
	}
	if p.configuration.Environments.Default == "" {
		err = fmt.Errorf("mybatis-config.xml <environments> atrribute defualt is empty")
		return
	}

	for _, v := range p.configuration.Environments.Children {
		if v.Id == p.configuration.Environments.Default {
			env = v
			break
		}
	}

	if env == nil {
		err = fmt.Errorf("mybatis-config.xml not find defulat environment: %s",
			p.configuration.Environments.Default)
		return
	}

	if env.DataSource == nil {
		err = fmt.Errorf("mybatis-config.xml environment: %s not set data source",
			p.configuration.Environments.Default)
		return
	}

	if !env.DataSource.PropertyMap().Has(schema.DRIVER) {
		err = fmt.Errorf("mybatis-config.xml environment: %s data source not set driver property",
			p.configuration.Environments.Default)
		return
	}

	if !env.DataSource.PropertyMap().Has(schema.URL) {
		err = fmt.Errorf("mybatis-config.xml environment: %s data source not set url property",
			p.configuration.Environments.Default)
		return
	}

	if !env.DataSource.PropertyMap().Has(schema.USERNAME) {
		err = fmt.Errorf("mybatis-config.xml environment: %s data source not set username property",
			p.configuration.Environments.Default)
		return
	}

	if !env.DataSource.PropertyMap().Has(schema.PASSWORD) {
		err = fmt.Errorf("mybatis-config.xml environment: %s data source not set password property",
			p.configuration.Environments.Default)
		return
	}

	return
}
