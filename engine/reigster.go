package engine

import (
	"fmt"
	"github.com/koyeo/gobatis/schema"
	"io/ioutil"
)

func (p *Engine) registerMapper(filePath string) (err error) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("read sql file error: %s", err)
		return
	}

	_schema, err := schema.UnmarshalMapper(data)
	if err != nil {
		return
	}

	mapper := NewMapper(_schema.Namespace, filePath)
	fmt.Println("注册Mapper：", mapper.SqlFile, mapper.Namespace)
	err = p.registerMapperMethod(mapper, _schema)
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
			fmt.Println("注册方法：", method)
		}
	}

	return
}
