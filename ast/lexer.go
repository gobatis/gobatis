package ast

import (
	"encoding/json"
	"fmt"
)

const EOF = "EOF"
const (
	TT_ROOT          = "ROOT"
	TT_TAG           = "TAG"
	TT_NODE_START    = "NODE_START"
	TT_NODE_END      = "NODE_END"
	TT_SELF_END_NODE = "SELF_NODE"
	TT_TEXT          = "TEXT"
	TT_ID            = "ID"
)

func NewLexer(file, content string) *Lexer {
	lexer := &Lexer{
		pos: Position{
			index:   -1,
			row:     0,
			col:     0,
			file:    file,
			content: content,
		}, // 读取字符串位置从-1开始计数
		chars: []rune(content),
	}
	lexer.length = len(lexer.chars)
	lexer.next()
	return lexer
}

type Lexer struct {
	pos    Position
	chars  []rune
	length int
	char   string // 当前处理字符
	root   *Token
	depth  int
}

// 预读: 读取下一个字符
func (p *Lexer) next() {
	p.pos.next(p.char)
	if p.pos.index < p.length {
		p.char = string(p.chars[p.pos.index])
	} else {
		p.char = EOF
	}
}

func (p *Lexer) peek(length int) (char string) {
	index := p.pos.peek(length)
	if index < p.length {
		return string(p.chars[index])
	}
	return EOF
}

func (p *Lexer) PrintTokens() {
	d, _ := json.MarshalIndent(p.root, "", "\t")
	fmt.Println(string(d))
}

func (p *Lexer) Parse() {
	root := NewToken()
	root.Type = TT_ROOT
	p.root = root
	p.parse(root)
}

func (p *Lexer) parse(root *Token) {

	for p.char != EOF {

		if isBlank(p.char) {
			// 过滤空白字符
			p.ignoreBlank()
			continue
		} else if p.char == "<" {
			p.parseNode(root)
		} else {
			// 解析文本节点
			token := new(Token)
			token.Type = TT_TEXT
			for not(p.char, EOF, "<") {
				token.Value += p.char
				p.next()
			}
			root.addToken(token)
			continue
		}

		p.next()
	}
}

func (p *Lexer) parseNode(root *Token) () {

	// <select>
	// <select ok>
	// <select id="2">
	// <select/>
	// <select id="2"/>
	//statParseTag := true
	//statParseTagName := true
	p.next()

	// 闭合标签
	if p.char == "/" {
		// 不处理
		for not(p.char, EOF, ">") {
			p.next()
		}
		return
	}

	// 解析标签名称
	token := NewToken()
	token.Type = TT_TAG
	root.addToken(token)
	for not(p.char, " ", ">") {
		token.Value += p.char
		p.next()
	}

	// 标签未结束
	if p.char != ">" {
		for p.char != ">" {
			p.parseAttribute(token)
		}
	}

	// 此时最后一个字符 >
	p.next()
	p.parse(token)

	return
}

func (p *Lexer) parseAttribute(token *Token) {

	// 过滤空白字符
	p.ignoreBlank()

	// 自闭和标签
	if p.char == "/" && p.peek(1) == ">" {
		p.next()
		return
	}

	// 解析属性名称
	attribute := NewAttribute()
	token.addAttribute(attribute)
	for not(p.char, "=", ">") {
		attribute.Name += p.char
		p.next()
	}

	// 标签已闭合
	if p.char == ">" {
		return
	}

	// 跳过等号
	p.next()

	// 过滤空白字符
	p.ignoreBlank()

	if p.char != "\"" {
		err := fmt.Errorf("非法字符 %s 行：%d,列: %d", p.char, p.pos.row, p.pos.col)
		fmt.Println(err)
		return
	}

	// 跳过第一个双引号
	p.next()

	// 解析属性值
	for not(p.char, "\"", ">") {
		attribute.Value += p.char
		p.next()
	}

	if p.char == ">" {
		return
	}

	// 跳过最后一个双引号
	p.next()

	return
}

// 过滤空白字符
func (p *Lexer) ignoreBlank() {
	for isBlank(p.char) {
		p.next()
	}
}

func is(char string, set ...string) bool {
	for _, v := range set {
		if v == char {
			return true
		}
	}
	return false
}

func not(char string, set ...string) bool {
	for _, v := range set {
		if v == char {
			return false
		}
	}
	return true
}

func isBlank(char string) bool {
	return is(char, " ", "\t", "\n")
}
