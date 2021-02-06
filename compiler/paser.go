package compiler

type Parser struct {
	tokens []Token
	token  Token
	index  int
}

func (p *Parser) next() Token {
	p.index += 1
	if p.index < len(p.tokens) {
		p.token = p.tokens[p.index]
	}
	return p.token
}

func (p *Parser) Parse(root *Token) {
	
}

func (p *Parser) expr() {

}

func (p *Parser) term() {

}

func (p *Parser) factor() {

}

type ParseResult struct {
	err  error
	node string
}

func (p *ParseResult) success() {

}

func (p *ParseResult) failed(err error) {
	p.err = err
}

func (p *ParseResult) register() {

}
