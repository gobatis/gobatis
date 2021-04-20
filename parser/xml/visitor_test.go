package xml

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestVisitor(t *testing.T) {
	visitor := NewVisitor()
	input := antlr.NewInputStream(`<if test="123">ok</if>`)
	//input, _ := newCharStream("./web.xml")
	lexer := NewXMLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewXMLParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Document()
	tree.Accept(visitor)
}
