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

//func TestReplaceIsolatedLessThan(t *testing.T) {
//	tests := []struct {
//		input    string
//		expected string
//	}{
//		{"This is a <test> string with a single < and another <tag> and one more <.", "This is a <test> string with a single &lt; and another <tag> and one more &lt;."},
//		{"<Hello<world", "&lt;Hello&lt;world"},
//		{"Hello<world>", "Hello<world>"},
//		{"<<<<<", "&lt;&lt;&lt;&lt;&lt;"},
//		{"<", "&lt;"},
//		{"Hello<", "Hello&lt;"},
//		{">Hello", ">Hello"},
//		{"<10", "&lt;10"},
//	}
//
//	for _, test := range tests {
//		result := replaceIsolatedLessThanWithEntity(test.input)
//		require.Equal(t, test.expected, result)
//	}
//}

func TestParser(t *testing.T) {
	//xs, err := Parse(`< ok <if test="ok">abc</if> <a  < = >`)
	//xs, err := Parse(`><`)
	//xs, err := Parse(`<a test="ok" /> / > < !=`)
	//xs, err := Parse(`ok <a/> b > c`)
	//xs, err := Parse(`</>/></a</a><b/><c></c>d<!>`)
	//xs, err := Parse(`<a></a>*`)
	//xs, err := Parse(`*~<<a test="ok">b</a>>kk`)
	//xs, err := Parse(`<a`)
	//xs, err := Parse(`
	//	select * from products where
	//	${a} weight == ${weight}
	//	<if test="price > 0">
	//		and price != #{ price[0].Age }
	//	</if>
	//`)
	xs, err := Parse(`
	select $a weight < / > =,
	'a', "b"
	${weight}
	// hi
	<if test="price <= 0">
	
	 and price &lt; #{ price[0].Age }
	 <where>
	         where2
	 </where>
	</if>
	
	kkk
		`)
	require.NoError(t, err)
	t.Log(xs.Raw())
	t.Log(xs.SQL())
}
