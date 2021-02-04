package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string
}

func main() {

	a := new(Test)
	a.Name = "a"
	test(a)
	fmt.Println(a)

	//x := new(Test) // type of x *float64
	//x.Value = "你好"
	//
	//oldXPointer := x
	//
	//fmt.Println("oldXPointer:", oldXPointer) // oldXPointer 34
	//// 改变原先的值
	//assign(x)
	//
	//fmt.Println("x", *x) // x 33
}

func assign(x interface{}) {
	v := reflect.ValueOf(&x).Elem()        // 这里传递的是指针的指针
	fmt.Println("v.CanSet():", v.CanSet()) // v.CanSet(): true

	fmt.Println(v.Type().Kind())
	c := reflect.New(v.Type().Elem())   // 创建原始的类型
	(c.Interface().(*Test)).Name = "不好" // 转换为 Interface 并且赋予新的值
	v.Set(c)
}

func test(a interface{}) {
	v := reflect.ValueOf(a)
	//fmt.Println(v.Kind())
	//fmt.Println(v.CanSet())
	//fmt.Println(v.Elem().CanSet())
	//name := v.Elem().FieldByName("Value")
	//name.Set(reflect.ValueOf("b"))
	b := reflect.New(v.Type().Elem())
	b.Elem().FieldByName("Value").Set(reflect.ValueOf("b"))
	v.Elem().Set(b.Elem())
}
