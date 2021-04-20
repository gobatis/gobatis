package xml

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestListener(t *testing.T) {
	input := antlr.NewInputStream(`<if test="123" name="k">ok</if>`)
	//input, _ := newCharStream("./web.xml")
	lexer := NewXMLLexer(input)
	
	//for {
	//	tok := lexer.NextToken()
	//	fmt.Println(tok)
	//	if tok.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//}
	
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewXMLParser(stream)
	p.BuildParseTrees = true
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	errorHandler := &ErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorHandler)
	tree := p.Document()
	if errorHandler.HasError() {
		for _, v := range errorHandler.Errors() {
			fmt.Println(v)
		}
	} else {
		antlr.ParseTreeWalkerDefault.Walk(&Listener{}, tree)
	}
}
