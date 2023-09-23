package xsql

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gobatis/gobatis/parser/commons"
	"github.com/gozelle/spew"
)

type XSQL struct {
	placeholder string
	dynamic     bool
	vars        []any
	sql         string
}

func (X XSQL) Placeholder() string {
	return X.placeholder
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

const lt = "&lt;"

func replaceIsolatedLessThanWithEntity(s string) string {

	runes := []rune(s)
	lastLeftBracket := -1
	pos := map[int]struct{}{}
	for i, r := range runes {
		switch r {
		case '<':
			// if a '<' is previously marked, replace it."
			if lastLeftBracket != -1 {
				pos[lastLeftBracket] = struct{}{}
			}
			lastLeftBracket = i
		case '>':
			// clear the previously marked '<'.
			lastLeftBracket = -1
		}
	}
	// check if there is a marked '<' at the end of the string.
	if lastLeftBracket != -1 {
		pos[lastLeftBracket] = struct{}{}
	}

	var r []rune
	for i := range runes {
		if _, ok := pos[i]; ok {
			r = append(r, []rune(lt)...)
		} else {
			r = append(r, runes[i])
		}
	}

	return string(r)
}

func Parse(source string) (*XSQL, error) {

	errs := &commons.CustomErrorListener{}

	source = replaceIsolatedLessThanWithEntity(source)
	lexer := NewXSQLLexer(antlr.NewInputStream(source))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewXSQLParser(stream)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)

	tree := p.Content()

	if errs.Error() != nil {
		return nil, errs.Error()
	}

	v := &Visitor{
		errs: errs,
		xsql: &XSQL{},
	}
	_ = v.Visit(tree)

	return v.xsql, errs.Error()
}

type Visitor struct {
	errs *commons.CustomErrorListener
	xsql *XSQL
}

func (v Visitor) Visit(tree antlr.ParseTree) interface{} {

	for _, c := range tree.GetChildren() {
		if v.errs.Error() != nil {
			return nil
		}
		switch t := c.(type) {
		case *ChardataContext:
			v.visitCharData(t)
		case *ElementContext:
			v.visitElement(t)
		case *ReferenceContext:
		}
	}

	spew.Json("result")

	return "a"
}

func (v Visitor) visitContent(ctx *ContentContext) {
	spew.Json(ctx.GetText())
}

func (v Visitor) visitElement(ctx *ElementContext) {
	spew.Json(ctx.GetText())
}

func (v Visitor) visitAttribute(ctx *AttributeContext) {
	spew.Json(ctx.GetText())
}

func (v Visitor) visitCharData(ctx *ChardataContext) {

	if ctx.TEXT() != nil {
		spew.Json(ctx.TEXT().GetText())
		fmt.Println(strings.TrimSpace(ctx.TEXT().GetText()))
	}

	if ctx.SEA_WS() != nil {
		spew.Json(ctx.SEA_WS().GetText())
	}
}
