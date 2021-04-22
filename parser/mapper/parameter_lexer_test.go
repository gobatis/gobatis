package xml

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestNewParamLexer(t *testing.T) {
	input := antlr.NewInputStream("id:int,name:string")
	lexer := NewParameterLexer(input)
	for {
		tok := lexer.NextToken()
		fmt.Println(tok.GetText(), tok.GetTokenType())
		if tok.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
}
