package engine

import (
	"fmt"
	"github.com/koyeo/gobatis/test/entity"
	"reflect"
)

// 绑定Mapper实例
func (p *Engine) BindMapper(instance interface{}) (err error) {

	v := reflect.ValueOf(instance)
	t := v.Type()
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
	//	method,ok := t.Elem().MethodByName(v.Value)
	//	if !ok{
	//		err = fmt.Errorf("")
	//	}
	//}
	fmt.Println("尝试修改值")
	//fmt.Println(v.Elem().CanSet())

	d := reflect.New(elem)
	fmt.Println("初始化值 Mapper:")
	d.Elem().FieldByName("GetUser").Set(reflect.ValueOf(func(id int64) (user *entity.User, err error) {
		fmt.Println("赋值成功!")
		return
	}))
	v.Elem().Set(d.Elem())

	//for i := 0; i < elem.NumField(); i++ {
	//	fmt.Println(elem.Field(i).Value)
	//	fmt.Println(elem.Field(i).Type)
	//	fmt.Println(elem.Field(i).Type.NumIn())
	//	fmt.Println(elem.Field(i).Type.NumOut())
	//	fieldType := elem.Field(i).Type
	//	for j := 0; j < fieldType.NumOut(); j++ {
	//		if fieldType.Out(j).Kind() == reflect.Ptr {
	//			// github.com/koyeo/gobatis/test/entity User
	//			fmt.Println(fieldType.Out(j))
	//			a := reflect.New(fieldType.Out(j).Elem())
	//			for k := 0; k < a.Elem().NumField(); k++ {
	//				a.Elem().Field(k).Set(reflect.ValueOf("abc"))
	//			}
	//			fmt.Println(a)
	//			fmt.Println(fieldType.Out(j).Elem().PkgPath(), fieldType.Out(j).Elem().Value())
	//
	//		} else {
	//			fmt.Println(fieldType.Out(j).Value())
	//		}
	//	}
	//}

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
