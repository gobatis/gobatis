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

	errs := &CustomErrorListener{}

	lexer := NewJsonPathLexer(antlr.NewInputStream("$.User[*]"))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewJsonPathParser(stream)
	//p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)

	// 提前构建解析树
	tree := p.Jsonpath()
	// 判断过程中是否有语法错误或词法错误
	t.Log(errs.errors)

	v := &visitor{}
	a := v.Visit(tree)
	// 判断结果
	t.Log(a)
}

var _ antlr.ErrorListener = (*CustomErrorListener)(nil)

type CustomErrorListener struct {
	errors []string
}

func (d *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	errorMsg := fmt.Sprintf("Ambiguity detected between tokens %d and %d. Ambiguous alternatives: %v", startIndex, stopIndex, ambigAlts)
	d.errors = append(d.errors, errorMsg)
}

func (d *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	errorMsg := fmt.Sprintf("Attempting full context mode between tokens %d and %d", startIndex, stopIndex)
	d.errors = append(d.errors, errorMsg)
}

func (d *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	errorMsg := fmt.Sprintf("Context sensitivity detected between tokens %d and %d", startIndex, stopIndex)
	d.errors = append(d.errors, errorMsg)
}

func (d *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("Syntax Error at line %d:%d - %s", line, column, msg)
	d.errors = append(d.errors, errorMsg)
}

func (d *CustomErrorListener) GetErrors() []string {
	return d.errors
}
