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
			fmt.Println(ctx.GetRuleIndex(), _ctx.IDENTIFIER(), _type.(*ParamTypeContext).IDENTIFIER())
		} else {
			fmt.Println(ctx.GetRuleIndex(), _ctx.IDENTIFIER())
		}
	}
}

func TestParseParameters(t *testing.T) {
	//lexer := NewExprLexer(antlr.NewInputStream(`id:int64, mix, name:any, age:string`))
	lexer := NewExprLexer(antlr.NewInputStream(`a:struct, b:array`))
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
	//fmt.Println(ctx.GetRuleIndex(), ctx.GetText())
}

func TestParseExpression(t *testing.T) {
	// a a.b a.b.c a.b() a.b(1) a.b(b1) a.b.c()
	// a.b(c, d...) a.b().c
	// a.b.c(c1, c2...).d
	// a(int(1))
	//lexer := NewExprLexer(antlr.NewInputStream("int(a) > 0"))
	//lexer := NewExprLexer(antlr.NewInputStream("a.B > 0"))
	//lexer := NewExprLexer(antlr.NewInputStream("a[0] > 0"))
	//lexer := NewExprLexer(antlr.NewInputStream("(aa+300)*2"))
	//lexer := NewExprLexer(antlr.NewInputStream(`a + (-1)`))
	//lexer := NewExprLexer(antlr.NewInputStream(`a[0:]`))
	//lexer := NewExprLexer(antlr.NewInputStream(`a.Person(int(a),b,c) + 1`))
	//lexer := NewExprLexer(antlr.NewInputStream(`int(a...) + 1`))
	//lexer := NewExprLexer(antlr.NewInputStream(`int(a...) + 1`))
	//lexer := NewExprLexer(antlr.NewInputStream(`a.b.c(1) + 1`))
	//lexer := NewExprLexer(antlr.NewInputStream(`a.b.c(int(d))+3`))
	//lexer := NewExprLexer(antlr.NewInputStream(`test(a,b,c) + 1`))
	//lexer := NewExprLexer(antlr.NewInputStream(`a.b(1+1).c + 1`))
	//lexer := NewExprLexer(antlr.NewInputStream(`a(a...)`))
	//lexer := NewExprLexer(antlr.NewInputStream(`a.Age + b["2"]`))
	lexer := NewExprLexer(antlr.NewInputStream(`a[2:3:4]`))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewExprParser(stream)
	
	//p.BuildParseTrees = true
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	antlr.ParseTreeWalkerDefault.Walk(&expressionListener{}, p.Expressions())
}
