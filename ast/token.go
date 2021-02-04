package ast

import "fmt"

type Token struct {
	Value     string
	Type      string
	Attribute []Attribute
}

func (p Token) String() string {
	return fmt.Sprintf("name:%s type:%s", p.Value, p.Type)
}

type Attribute struct {
	Name  string
	Value string
}
