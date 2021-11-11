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
	syntaxError  string
	runtimeError string
}

func testError(t *testing.T, item errorCase) {
	fs, err := parseMapper("test_error.xml", wrapMapperSchema(item.definition))
	if item.syntaxError != "" {
		require.Error(t, err, item)
		//require.Equal(t, item.syntaxError, err.Error(), item)
		t.Log(err.Error())
	} else {
		require.NoError(t, err, item)
		require.Equal(t, 1, len(fs), item.definition)
		f := fs[0]
		
		c := f.newCaller(nil)
		_, err = c.fragment.buildStmt(item.in)
		t.Log(err)
	}
	
}

func TestXMLErrors(t *testing.T) {
	errors := []errorCase{
		{
			definition: `
				<insert></insert>
			`,
			syntaxError: "[ERROR 32]",
		},
		{
			definition: `
				<insert id=""></insert>
			`,
			syntaxError: "[ERROR 32]",
		},
		{
			definition: `
				<insert id=''></insert>
			`,
			syntaxError: "[ERROR 32]",
		},
		{
			definition: `
				<insert id=''>/insert>
			`,
			syntaxError: "[ERROR 32]",
		},
		{
			definition: `
				<insert id=''></ok>
			`,
			syntaxError: "[ERROR 32]",
		},
	}
	for _, v := range errors {
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
			runtimeError: "[ERROR 32]",
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
