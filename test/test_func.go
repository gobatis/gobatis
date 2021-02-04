package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := func(name string) {}

	v := reflect.TypeOf(a)
	fmt.Println(v.In(0).Name())
}
