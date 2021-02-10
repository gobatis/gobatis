package compiler

// 记录代码位置
type Position struct {
	line   int
	column int
}

func (p *Position) next(char rune) {
	if char == LINE_FEED {
		p.column = 1
		p.line++
	} else {
		p.column++
	}
}

func (p *Position) fork() *Position {
	return &Position{
		line:   p.line,
		column: p.column,
	}
}
