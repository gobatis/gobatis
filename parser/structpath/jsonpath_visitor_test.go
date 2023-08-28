package structpath

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

//var _ antlr.ParseTreeVisitor = (*visitor)(nil)

type visitor struct {
	//*BaseJsonPathVisitor
}

func (v visitor) Visit(tree antlr.ParseTree) interface{} {

	switch t := tree.(type) {
	case *JsonpathContext:
		fmt.Println(0, t.GetText())
		return v.VisitJsonpathContext(t)
	case *DotnotationContext:
		fmt.Println(1, t.GetText())
	case *Dotnotation_exprContext:
		fmt.Println(2, t.GetText())
	default:
		panic("unknown")
	}

	return nil
}

func (v visitor) VisitJsonpathContext(ctx *JsonpathContext) interface{} {

	fmt.Println("ok")
	for _, child := range ctx.GetChildren() {
		switch t := child.(type) {
		case *DotnotationContext:
			return v.VisitDotnotationContext(t)
		default:
			fmt.Println(reflect.TypeOf(t).String())
		}
	}

	return nil
}

func (v visitor) VisitDotnotationContext(ctx *DotnotationContext) interface{} {

	for _, c := range ctx.GetChildren() {
		switch t := c.(type) {
		case *antlr.TerminalNodeImpl:
			fmt.Println("terminal", t.String())
		case *Dotnotation_exprContext:
			return v.VisitDotnotationExpr(t)
		}
		fmt.Println("kkk", reflect.TypeOf(c).String())
	}

	return nil
}

func (v visitor) VisitDotnotationExpr(ctx *Dotnotation_exprContext) interface{} {

	for _, c := range ctx.GetChildren() {
		fmt.Println("ddd", reflect.TypeOf(c).String())
	}
	return nil
}

//func (v visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
//	//TODO implement me
//	panic("implement me 2")
//}
//
//func (v visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
//	//TODO implement me
//	panic("implement me 3")
//}

func TestVisitor(t *testing.T) {
	lexer := NewJsonPathLexer(antlr.NewInputStream("$.User[*]"))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJsonPathParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	v := &visitor{}
	a := v.Visit(p.Jsonpath())
	t.Log(a)
}
