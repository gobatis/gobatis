package simplemath

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

type EvalVisitor struct {
	*BaseSimpleMathVisitor
}

func (v *EvalVisitor) VisitExpr(ctx *ExprContext) interface{} {
	if ctx.Term() != nil {
		return v.Visit(ctx.Term())
	}
	left := v.Visit(ctx.Expr()).(int)
	right := v.Visit(ctx.Term()).(int)
	fmt.Println(left, right)
	return left + right
}

func (v *EvalVisitor) VisitTerm(ctx *TermContext) interface{} {
	if ctx.Factor() != nil {
		return v.Visit(ctx.Factor())
	}
	left := v.Visit(ctx.Term()).(int)
	right := v.Visit(ctx.Factor()).(int)
	fmt.Println(left, right)
	return left * right
}

func (v *EvalVisitor) VisitFactor(ctx *FactorContext) interface{} {
	if ctx.NUMBER() != nil {
		val, err := strconv.Atoi(ctx.NUMBER().GetText())
		if err != nil {
			panic(fmt.Sprintf("Invalid number: %s", ctx.NUMBER().GetText()))
		}
		return val
	}
	panic("todo")
	return nil
}

func TestVisitor(t *testing.T) {
	input := "3+5"
	is := antlr.NewInputStream(input)
	lexer := NewSimpleMathLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSimpleMathParser(ts)

	// Error strategy
	p.RemoveErrorListeners()
	errorListener := &CustomErrorListener{}
	p.AddErrorListener(errorListener)

	tree := p.Expr()

	if errorListener.HasError {
		fmt.Println("Failed to parse the expression due to errors.")
		return
	}

	visitor := &EvalVisitor{
		BaseSimpleMathVisitor: &BaseSimpleMathVisitor{},
	}
	result := tree.Accept(visitor)
	fmt.Println(result)
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	HasError bool
}

func (d *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.HasError = true
	fmt.Printf("Syntax Error: %s at line %d:%d\n", msg, line, column)
}
