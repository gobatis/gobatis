package engine

import (
	"fmt"
	"github.com/koyeo/mybatis.go/schema"
	"io/ioutil"
)

func (p *Engine) registerMapper(filePath string) (err error) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("read sql file error: %s", err)
		return
	}

	mapperSchema, err := schema.UnmarshalMapper(data)
	if err != nil {
		return
	}

	mapper := NewMapper(filePath)

	err = p.registerMapperMethod(mapper, mapperSchema)
	if err != nil {
		return
	}

	p.Mappers().Add(mapper)

	return
}

func (p *Engine) registerMapperMethod(mapper *Mapper, _schema *schema.Mapper) (err error) {

	if _schema.Selects != nil {
		for _, v := range _schema.Selects {
			method := new(Method)
			method.Name = v.Id
			method.ParameterType = v.ParameterType
			method.ResultType = v.ResultType
			method.SQL = v.SQL
			mapper.Methods().Add(method)
			fmt.Println(method)
		}
	}

	return
}
