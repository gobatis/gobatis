package expr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

type parameterListener struct {
	*antlr.BaseParseTreeListener
}

func (p *parameterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetRuleIndex() == ExprParserRULE_paramDecl {
		_ctx := ctx.(*ParamDeclContext)
		_type := _ctx.ParamType()
		if _type != nil {
			fmt.Println(_ctx.IDENTIFIER(), _type.(*ParamTypeContext).IDENTIFIER())
		} else {
			fmt.Println(_ctx.IDENTIFIER())
		}
	}
}

func TestParseParameters(t *testing.T) {
	lexer := NewExprLexer(antlr.NewInputStream(`id:int64, mix, name:any, age:string`))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewExprParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	antlr.ParseTreeWalkerDefault.Walk(&parameterListener{}, p.Parameters())
}

type expressionListener struct {
	*antlr.BaseParseTreeListener
}

func (p *expressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetRuleIndex() == ExprParserRULE_expressions {
		fmt.Println("总TOKEN:", ctx.GetStop().GetTokenIndex())
	}
}

func (p *expressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	
	fmt.Println(ctx.GetRuleIndex(), ctx.GetText(), ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex())
}

func TestParseExpression(t *testing.T) {
	//lexer := NewExprLexer(antlr.NewInputStream("int(a) > 0"))
	//lexer := NewExprLexer(antlr.NewInputStream("a.B > 0"))
	//lexer := NewExprLexer(antlr.NewInputStream("a[0] > 0"))
	lexer := NewExprLexer(antlr.NewInputStream("(aa+300)*2"))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewExprParser(stream)
	
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	antlr.ParseTreeWalkerDefault.Walk(&expressionListener{}, p.Expressions())
}
