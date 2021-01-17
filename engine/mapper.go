package engine

import (
	"fmt"
	"reflect"
)

// 绑定Mapper实例
func (p *Engine) BindMapper(instance interface{}) (err error) {

	t := reflect.TypeOf(instance)
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("engine.BindMapper method paramenter should be a point")
		return
	}

	elem := t.Elem()
	path := fmt.Sprintf("%s@%s", p.resolve(elem.PkgPath()), elem.Name())

	m := p.Mappers().Get(path)
	if m == nil {
		err = fmt.Errorf("engine.BindMapper mapper binded with %s not found", path)
		return
	}

	//for _, v := range m.Methods().All() {
	//	method,ok := t.Elem().MethodByName(v.Name)
	//	if !ok{
	//		err = fmt.Errorf("")
	//	}
	//}

	for i := 0; i < elem.NumField(); i++ {
		fmt.Println(elem.Field(i).Name)
		fmt.Println(elem.Field(i).Type)
		fmt.Println(elem.Field(i).Type.NumIn())
		fmt.Println(elem.Field(i).Type.NumOut())
		fieldType := elem.Field(i).Type
		for j := 0; j < fieldType.NumOut(); j++ {
			if fieldType.Out(j).Kind() == reflect.Ptr {
				// github.com/koyeo/mybatis.go/test/entity User
				fmt.Println(fieldType.Out(j))
				a := reflect.New(fieldType.Out(j).Elem())
				for k := 0; k < a.Elem().NumField(); k++ {
					a.Elem().Field(k).Set(reflect.ValueOf("abc"))
				}
				fmt.Println(a)
				fmt.Println(fieldType.Out(j).Elem().PkgPath(), fieldType.Out(j).Elem().Name())

			} else {
				fmt.Println(fieldType.Out(j).Name())
			}
		}
	}

	return
}

func NewMapper(namespace, sqlFile string) *Mapper {
	return &Mapper{SqlFile: sqlFile, Namespace: namespace}
}

type Mapper struct {
	SqlFile   string
	Namespace string
	methods   *Methods
}

func (p *Mapper) Methods() *Methods {
	if p.methods == nil {
		p.methods = NewMethods()
	}
	return p.methods
}
