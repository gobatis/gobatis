package gobatis

import (
	"encoding/json"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/dtd"
	"github.com/gobatis/gobatis/parser/xml"
	"github.com/stretchr/testify/require"
	"reflect"
	"strings"
	"testing"
)

func TestXMLParseError(t *testing.T) {
	listener := &xmlParser{
		file:          "text.xml",
		stack:         newXMLStack(),
		coverage:      newCoverage(),
		rootElement:   dtd.Mapper,
		elementGetter: dtd.MapperElement,
	}
	_ = listener
	tokens := wrapMapperSchema(`
<update id="updateUser">some...</update>
<insert id="ok">some...</insert>
`)
	inputStream := antlr.NewInputStream(strings.TrimSpace(tokens))
	defer func() {
		err := catch(listener.file, recover())
		t.Log(err)
		require.Error(t, err)
	}()
	lexer := xml.NewXMLLexer(inputStream)
	lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := xml.NewXMLParser(stream)
	parser.BuildParseTrees = true
	parser.SetErrorHandler(newXmlErrorStrategy(listener.file))
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Document())
}

type errorCase struct {
	definition   string
	in           []reflect.Value
	message      string
	context      string
	syntaxError  int
	runtimeError int
}

func testError(t *testing.T, item errorCase) {
	fs, err := parseMapper("test_error.xml", wrapMapperSchema(item.definition))
	if item.syntaxError > 0 {
		require.Error(t, err, item)
		em := ParseErrorMessage(err.Error())
		require.Equal(t, item.syntaxError, em.Code, item)
		require.Equal(t, item.message, em.Message, item)
		require.Equal(t, item.context, em.Context, item)
		t.Log(err.Error())
	} else {
		require.NoError(t, err, item)
		require.Equal(t, 1, len(fs), item.definition)
		f := fs[0]
		
		c := f.newCaller(nil)
		_, err = c.fragment.buildStmt(item.in)
		if item.runtimeError > 0 {
			require.Error(t, err, item)
			em := ParseErrorMessage(err.Error())
			require.Equal(t, item.runtimeError, em.Code, item)
			require.Equal(t, item.message, em.Message, item)
			require.Equal(t, item.context, em.Context, item)
			t.Log(err)
		} else {
			require.NoError(t, err, item)
		}
	}
	
}

var xmlSyntaxTestCases = []errorCase{
	{
		definition: `
				<a
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b=
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b='
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b=""
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="">c
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="">c<
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="">c</
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="">c</d
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="">c</dd
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="" c />
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     `<aa b="" c />`,
	},
	{
		definition: `
				<aa b="" ' />
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="" c " />
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     `</mapper>`,
	},
	{
		definition: `
				<aa b="" " />
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     "</mapper>",
	},
	{
		definition: `
				<aa b="" / />
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     `<aa b="" / />`,
	},
	{
		definition: `
				<aa b="" //>
			`,
		syntaxError: xml_syntax_err,
		message:     "xml syntax error",
		context:     `<aa b="" //>`,
	},
}

func TestXMLErrors(t *testing.T) {
	for _, v := range xmlSyntaxTestCases {
		testError(t, v)
	}
}

func TestExprErrors(t *testing.T) {
	errors := []errorCase{
		{
			definition: `
				<insert id="TestExpr1">
					a + #{a}
				</insert>
			`,
			runtimeError: expr_syntax_err,
		},
		{
			definition: `
				<insert id="TestExpr1" parameter="user">
					a + #{user.Sex}
				</insert>
			`,
			in: []reflect.Value{rv(struct {
				Name string
				Age  int
			}{
				Name: "tom",
				Age:  18,
			})},
			runtimeError: expr_syntax_err,
		},
		{
			definition: `
				<insert id="TestExpr1" parameter="user">
					a + #{user[0]}
				</insert>
			`,
			in: []reflect.Value{rv(struct {
				Name string
				Age  int
			}{
				Name: "tom",
				Age:  18,
			})},
			runtimeError: expr_syntax_err,
		},
		{
			definition: `
				<insert id="TestExpr1" parameter="user">
					#{1 + user.Name}
				</insert>
			`,
			in: []reflect.Value{rv(struct {
				Name string
				Age  int
			}{
				Name: "tom",
				Age:  18,
			})},
			runtimeError: expr_syntax_err,
		},
	}
	for _, v := range errors {
		testError(t, v)
	}
}

func TestParseErrorMessage(t *testing.T) {
	d, _ := json.Marshal(ParseErrorMessage("test_error.xml line 2:28 [4] var 'a' not found at 'a + #{a}'"))
	t.Log(string(d))
	
	d, _ = json.Marshal(ParseErrorMessage("test_error.xml line 2:28 [4] var 'a' not found"))
	t.Log(string(d))
}
