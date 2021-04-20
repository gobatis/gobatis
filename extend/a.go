package extend

import "fmt"

type A struct {
}

func (p A) Hello() {
	fmt.Println("This is a hello")
}

func (p A) Hi() {
	fmt.Println("This is a hi")
}

type B struct {
	A
}

func (p B) Hello() {
	p.Hi()
	fmt.Println("this is b Hello")
}
