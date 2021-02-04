package ast

import "fmt"

func NewToken() *Token {
	return &Token{}
}

type Token struct {
	Value      string       `json:"value,omitempty"`
	Type       string       `json:"type,omitempty"`
	Attributes []*Attribute `json:"attributes,omitempty"`
	Tokens     []*Token     `json:"tokens,omitempty"`
}

func (p *Token) addToken(token ...*Token) {
	p.Tokens = append(p.Tokens, token...)
}

func (p *Token) addAttribute(attribute ...*Attribute) {
	p.Attributes = append(p.Attributes, attribute...)
}

func (p Token) String() string {
	return fmt.Sprintf("name:%s type:%s", p.Value, p.Type)
}

func NewAttribute() *Attribute {
	return &Attribute{}
}

type Attribute struct {
	Name  string
	Value string
}
