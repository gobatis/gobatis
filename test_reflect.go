package main

import (
	"fmt"
	"github.com/koyeo/mybatis.go/core"
	"reflect"
)

func main() {
	model := new(core.Model)
	t := reflect.TypeOf(model)
	//v := reflect.ValueOf(model)
	fmt.Println(t.PkgPath())
	fmt.Println(t.Name())
	fmt.Println(t.String())
	fmt.Println(t.Elem().PkgPath())
}
