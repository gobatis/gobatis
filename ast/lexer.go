package ast

import "fmt"

const EOF = "EOF"
const (
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

func (p *Lexer) Parse() {
	//tokens := make([]Token, 0)

	for p.char != EOF {

		if isBlank(p.char) {
			// 过滤空白字符
			p.ignoreBlank()
			continue
		} else if p.char == "<" {
			p.next()
			p.parseNode()
		} else {
			// 解析文本节点
			//fmt.Println(p.char)
			//fmt.Println(p.pos.row, p.pos.col)
			//panic("非法字符")
		}

		p.next()
	}

}

func (p *Lexer) parseNode() (token Token) {

	// <select>
	// <select ok>
	// <select id="2">
	// <select/>
	// <select id="2"/>
	//statParseTag := true
	//statParseTagName := true

	// 解析标签名称
	nodeName := ""

	for not(p.char, " ", ">") {
		nodeName += p.char
		p.next()
	}

	fmt.Println(nodeName)

	// 标签已闭合
	if p.char == ">" {
		return
	}

	for p.char != ">" {
		p.parseAttribute()
	}

	// 最后一个字符 >

	return
}

func (p *Lexer) parseAttribute() (attr Attribute) {

	// 过滤空白字符
	p.ignoreBlank()

	// 解析属性名称
	attrName := ""
	defer fmt.Printf("attrName: '%s'\n", attrName)

	for not(p.char, "=", ">") {
		attrName += p.char
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
	attrValue := ""
	defer func() {
		fmt.Printf("attrValue: '%s'\n", attrValue)
	}()
	for not(p.char, "\"", ">") {
		attrValue += p.char
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
