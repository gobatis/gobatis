package main

import "fmt"

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func main() {
	tokens := []string{"cad"}
	node := new(Node)
	for _, token := range tokens {
		node, err := S(token)
		if err != nil {
			fmt.Println("注册错误")
			return
		}

	}
}

const EOF = "EOF"

func NewParser(tokens []string) *Parser {
	return &Parser{tokens: tokens}
}

type Parser struct {
	index  int
	tokens []string
	token  string
}

func (p *Parser) next() {
	p.index++
	if p.index < len(p.tokens) {
		p.token = p.tokens[p.index]
	} else {
		p.token = EOF
	}
}

func (p *Parser) S() (s *Node) {
	p.next()
	if p.token == "c" {
		left := "c"
		p.next()
		op := p.A()
		p.next()
		if p.token == "d" {
			right := "d"
		}
	}
}

func (p *Parser) A() (node *Node) {

	if p.token == "a" {
		left := "a"
		p.next()
		if p.token == "b" {
			right := "b"
		}
	}

	return
}
