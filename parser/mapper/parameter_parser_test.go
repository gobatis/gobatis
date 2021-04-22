package xml

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

type testParameterListener struct {
}

func (t testParameterListener) VisitTerminal(node antlr.TerminalNode) {
	//fmt.Println("visit terminal", node.GetText())
}

func (t testParameterListener) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Println("visit error node", node.GetText())
}

func (t testParameterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetRuleIndex(), ctx.GetText())
}

func (t testParameterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {

}

func TestNewParameterParser(t *testing.T) {
	lexer := NewParameterLexer(antlr.NewInputStream("id:int64, mix:a234 name:any,age:string"))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewParameterParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	antlr.ParseTreeWalkerDefault.Walk(&testParameterListener{}, p.Expression())
}
