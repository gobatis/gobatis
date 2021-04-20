package xml

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewVisitor() *Visitor {
	return &Visitor{}
}

type Visitor struct {
}

func (p *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return nil
}

func (p *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	panic("implement me")
}

func (p *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("implement me")
}

func (p *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("implement me")
}

func (p *Visitor) VisitDocument(ctx *DocumentContext) interface{} {
	fmt.Println(ctx.Element())
	return ctx.Element().Accept(p)
}

func (p *Visitor) VisitProlog(ctx *PrologContext) interface{} {
	fmt.Println("visit Prolog")
	//fmt.Println(ctx.getC)
	fmt.Println(ctx.GetText())
	panic("implement me")
}

func (p *Visitor) VisitContent(ctx *ContentContext) interface{} {
	panic("implement me")
}

func (p *Visitor) VisitElement(ctx *ElementContext) interface{} {
	fmt.Println("call VisitElement: ", ctx.GetText())
	fmt.Println(ctx.OPEN(0))
	fmt.Println(ctx.CLOSE(0))
	fmt.Println(ctx.Name(0))
	return nil
}

func (p *Visitor) VisitReference(ctx *ReferenceContext) interface{} {
	panic("implement me")
}

func (p *Visitor) VisitAttribute(ctx *AttributeContext) interface{} {
	panic("implement me")
}

func (p *Visitor) VisitChardata(ctx *ChardataContext) interface{} {
	panic("implement me")
}

func (p *Visitor) VisitMisc(ctx *MiscContext) interface{} {
	panic("implement me")
}
