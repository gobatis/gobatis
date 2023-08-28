package structpath

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestNewParamLexer(t *testing.T) {
	input := antlr.NewInputStream("$.User.Roles[*].ID")
	lexer := NewJsonPathLexer(input)
	for {
		tok := lexer.NextToken()
		fmt.Println(tok.GetText(), tok.GetTokenType())
		if tok.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
}
