package ast

func NewLexer(file, content string) *Lexer {
	return &Lexer{
		pos: Position{
			index:   -1,
			row:     0,
			col:     0,
			file:    file,
			content: content,
		}, // 读取字符串位置从-1开始计数
	}
}

type Lexer struct {
	pos         Position
	currentChar string // 当前处理字符

}

// 预读: 读取下一个字符
func (p *Lexer) Next() {

}
