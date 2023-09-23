package xsql

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	input := antlr.NewInputStream("$.User.Roles[*].ID")
	lexer := NewXSQLLexer(input)
	for {
		tok := lexer.NextToken()
		fmt.Println(tok.GetText(), tok.GetTokenType())
		if tok.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
}

func TestReplaceIsolatedLessThan(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is a <test> string with a single < and another <tag> and one more <.", "This is a <test> string with a single &lt; and another <tag> and one more &lt;."},
		{"<Hello<world", "&lt;Hello&lt;world"},
		{"Hello<world>", "Hello<world>"},
		{"<<<<<", "&lt;&lt;&lt;&lt;&lt;"},
		{"<", "&lt;"},
		{"Hello<", "Hello&lt;"},
		{">Hello", ">Hello"},
		{"<10", "&lt;10"},
	}

	for _, test := range tests {
		result := replaceIsolatedLessThanWithEntity(test.input)
		require.Equal(t, test.expected, result)
	}
}

func TestParser(t *testing.T) {
	xs, err := Parse(`
		select * from products where
		weight > ${ weight }
		<if test="price > 0">
			and price < #{ price }
		</if>
	`)
	require.NoError(t, err)
	t.Log(xs.Placeholder())
	t.Log(xs.SQL())
}
