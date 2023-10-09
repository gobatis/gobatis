package xsql

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gobatis/gobatis/parser/commons"
	"github.com/gozelle/spew"
)

type XSQL struct {
	raw     strings.Builder
	dynamic bool
	vars    []any
	sql     string
}

func (X XSQL) Raw() string {
	return X.raw.String()
}

func (X XSQL) SQL() string {
	return X.sql
}

func (X XSQL) Dynamic() bool {
	return X.dynamic
}

func (X XSQL) Vars() []any {
	return X.vars
}

//const lt = "&lt;"

//func replaceIsolatedLessThanWithEntity(s string) string {
//
//	runes := []rune(s)
//	lastLeftBracket := -1
//	pos := map[int]struct{}{}
//	for i, r := range runes {
//		switch r {
//		case '<':
//			// if a '<' is previously marked, replace it."
//			if lastLeftBracket != -1 {
//				pos[lastLeftBracket] = struct{}{}
//			}
//			lastLeftBracket = i
//		case '>':
//			// clear the previously marked '<'.
//			lastLeftBracket = -1
//		}
//	}
//	// check if there is a marked '<' at the end of the string.
//	if lastLeftBracket != -1 {
//		pos[lastLeftBracket] = struct{}{}
//	}
//
//	var r []rune
//	for i := range runes {
//		if _, ok := pos[i]; ok {
//			r = append(r, []rune(lt)...)
//		} else {
//			r = append(r, runes[i])
//		}
//	}
//
//	return string(r)
//}

func Parse(source string) (*XSQL, error) {

	errs := &commons.CustomErrorListener{}

	//source = replaceIsolatedLessThanWithEntity(source)
	lexer := NewXSQLLexer(antlr.NewInputStream(source))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewXSQLParser(stream)
	p.BuildParseTrees = true
	//p.RemoveErrorListeners()
	//p.AddErrorListener(errs)
	spew.Json(p.GetInterpreter().GetPredictionMode())
	//p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL) // 设置为SLL模式
	//p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL) // 设置为SLL模式

	tree := p.Content()

	if errs.Error() != nil {
		return nil, errs.Error()
	}

	v := &Visitor{
		errs: errs,
		xsql: &XSQL{},
	}
	_ = v.VisitContent(tree.(*ContentContext))

	return v.xsql, errs.Error()
}

type Visitor struct {
	errs *commons.CustomErrorListener
	xsql *XSQL
}

func (v Visitor) VisitContent(ctx *ContentContext) interface{} {
	fmt.Println("content:", ctx.GetText())
	for _, c := range ctx.GetChildren() {
		if v.errs.Error() != nil {
			return nil
		}
		switch t := c.(type) {
		case *TagStartContext:
			v.visitTagStart(t)
		case *TagEndContext:
			v.visitTagEnd(t)
		case *CloseTagContext:
			v.visitCloseTag(t)
		case *ContentContext:
			v.VisitContent(t)
		case *ExprContext:
			v.visitExpr(t)
		case *ReferenceContext:
			v.visitReference(t)
		case *ChardataContext:
			v.visitCharData(t)
		default:
			v.errs.AddError(fmt.Errorf("unsupport rule: %v", c.GetPayload()))
		}
	}

	return "a"
}

func (v Visitor) visitTagStart(ctx *TagStartContext) {
	fmt.Println("tag start::", ctx.GetText(), "attributes", len(ctx.AllAttribute()))
}

func (v Visitor) visitTagEnd(ctx *TagEndContext) {
	fmt.Println("tag end::", ctx.NAME())
}

func (v Visitor) visitCloseTag(ctx *CloseTagContext) {
	fmt.Println("close tag::", ctx.NAME())
}

func (v Visitor) visitExpr(ctx *ExprContext) {
	if v.errs.Error() != nil {
		return
	}
	if ctx.HASH() != nil {
		fmt.Println("##:", ctx.GetText())
		v.xsql.raw.WriteString(fmt.Sprintf("##{%s}", ctx.Chardata().GetText()))
	} else {
		v.xsql.raw.WriteString(fmt.Sprintf("$${%s}", ctx.Chardata().GetText()))
	}
}

func (v Visitor) visitAttribute(ctx *AttributeContext) {
	if v.errs.Error() != nil {
		return
	}
	spew.Json(ctx.GetText())
}

func (v Visitor) visitCharData(ctx *ChardataContext) {
	if v.errs.Error() != nil {
		return
	}

	if ctx.WS() != nil {
		v.xsql.raw.WriteString(ctx.WS().GetText())
	} else {
		fmt.Println("chardata:", ctx.GetText())
		v.xsql.raw.WriteString(ctx.GetText())
	}
}

func (v Visitor) visitReference(ctx *ReferenceContext) {
	fmt.Println("reference:", ctx.GetText())
}

func isNameStartChar(c int) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
