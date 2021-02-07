package compiler

//func NewLexer(content string) *Lexer {
//	lexer := &Lexer{
//		pos: Position{
//			index:  -1,
//			line:   1,
//			column: 1,
//		}, // index start with -1
//		chars: []rune(content),
//	}
//	lexer.length = len(lexer.chars)
//	lexer.next()
//	return lexer
//}
//
//type Lexer struct {
//	pos    Position
//	chars  []rune
//	length int
//	char   string // current handle character
//	root   *Token
//	depth  int
//}
//
//// read next character
//func (p *Lexer) next() {
//	p.pos.next(p.char)
//	if p.pos.index < p.length {
//		p.char = string(p.chars[p.pos.index])
//	} else {
//		p.char = EOF
//	}
//}
//
//func (p *Lexer) peek(length int) (char string) {
//	index := p.pos.peek(length)
//	if index < p.length {
//		return string(p.chars[index])
//	}
//	return EOF
//}
//
//func (p *Lexer) PrintTokens() {
//	d, _ := json.MarshalIndent(p.root, "", "\t")
//	fmt.Println(string(d))
//}
//
//func (p *Lexer) Tokenize() {
//	root := NewToken()
//	root.Type = TT_ROOT
//	p.root = root
//	p.tokenize(root)
//}
//
//func (p *Lexer) tokenize(root *Token) {
//
//	for p.char != EOF {
//
//		if isBlank(p.char) {
//			// ignore blank character
//			p.ignoreBlank()
//			continue
//		} else if p.char == "<" {
//			// tokenize tag node
//			p.parseNode(root)
//		} else {
//			// tokenize text node
//			token := new(Token)
//			token.Type = TT_TEXT
//			token.Start = p.pos.fork()
//			for not(p.char, EOF, "<") {
//				token.Value += p.char
//				p.next()
//			}
//			token.End = p.pos.fork()
//			root.addToken(token)
//			continue
//		}
//
//		p.next()
//	}
//}
//
//func (p *Lexer) parseNode(root *Token) () {
//
//	// <select>
//	// <select ok>
//	// <select id="2">
//	// <select/>
//	// <select id="2"/>
//	//statParseTag := true
//	//statParseTagName := true
//	p.next()
//
//	// closed tag
//	if p.char == "/" {
//		// just pass
//		for not(p.char, EOF, ">") {
//			p.next()
//		}
//		return
//	}
//
//	// tokenize tag name
//	token := NewToken()
//	token.Type = TT_TAG
//	token.Start = p.pos.fork()
//	root.addToken(token)
//	for not(p.char, " ", ">") {
//		token.Value += p.char
//		p.next()
//	}
//
//	// if tag not closed, tokenize attributes
//	if p.char != ">" {
//		for p.char != ">" {
//			p.parseAttribute(token)
//		}
//	}
//
//	// now, p.char == ">"
//	p.next()
//	p.tokenize(token)
//
//	return
//}
//
//func (p *Lexer) parseAttribute(token *Token) {
//
//	p.ignoreBlank()
//
//	// just ignore self closed tag
//	if p.char == "/" && p.peek(1) == ">" {
//		p.next()
//		return
//	}
//
//	// tokenize attribute name
//	attribute := NewAttribute()
//	token.addAttribute(attribute)
//	for not(p.char, "=", ">") {
//		attribute.Name += p.char
//		p.next()
//	}
//
//	// if tag closed
//	if p.char == ">" {
//		return
//	}
//
//	// pass character "="
//	p.next()
//
//	p.ignoreBlank()
//
//	if p.char != "\"" {
//		err := fmt.Errorf("非法字符 %s 行：%d,列: %d", p.char, p.pos.line, p.pos.column)
//		fmt.Println(err)
//		return
//	}
//
//	// pass the first "
//	p.next()
//
//	// 解析属性值
//	for not(p.char, "\"", ">") {
//		attribute.Value += p.char
//		p.next()
//	}
//
//	if p.char == ">" {
//		return
//	}
//
//	// pass the last "
//	p.next()
//
//	return
//}
//
//func (p *Lexer) ignoreBlank() {
//	for isBlank(p.char) {
//		p.next()
//	}
//}
//
//func is(char string, set ...string) bool {
//	for _, v := range set {
//		if v == char {
//			return true
//		}
//	}
//	return false
//}
//
//func not(char string, set ...string) bool {
//	for _, v := range set {
//		if v == char {
//			return false
//		}
//	}
//	return true
//}
//
//func isBlank(char string) bool {
//	return is(char, " ", "\t", "\n")
//}