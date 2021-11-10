package gobatis

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/dtd"
	"github.com/gobatis/gobatis/parser/xml"
	"github.com/stretchr/testify/require"
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
	//lexer.AddErrorListener(newLexerErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := xml.NewXMLParser(stream)
	parser.BuildParseTrees = true
	//parser.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	parser.SetErrorHandler(newXmlErrorStrategy(listener.file))
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Document())
}
