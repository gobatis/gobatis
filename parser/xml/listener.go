package xml

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Listener struct {
}

func (l *Listener) VisitTerminal(c antlr.TerminalNode) {
	fmt.Println("VisitTerminal:", c.GetText())
}

func (l *Listener) VisitErrorNode(c antlr.ErrorNode) {
	fmt.Println("VisitErrorNode:", c.GetText())
}

func (l *Listener) ExitEveryRule(c antlr.ParserRuleContext) {
	fmt.Println("ExitEveryRule:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterDocument(c *DocumentContext) {
	fmt.Println("EnterDocument:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterProlog(c *PrologContext) {
	fmt.Println("EnterProlog:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterContent(c *ContentContext) {
	fmt.Println("EnterContent:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterElement(c *ElementContext) {
	fmt.Println("EnterElement:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterReference(c *ReferenceContext) {
	fmt.Println("EnterReference:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterAttribute(c *AttributeContext) {
	fmt.Println("EnterAttribute:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterChardata(c *ChardataContext) {
	fmt.Println("EnterChardata:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterMisc(c *MiscContext) {
	fmt.Println("EnterMisc:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) ExitDocument(c *DocumentContext) {
	fmt.Println("ExitDocument:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) ExitProlog(c *PrologContext) {
	fmt.Println("ExitProlog:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) ExitContent(c *ContentContext) {
	fmt.Println("ExitContent:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) ExitElement(c *ElementContext) {
	fmt.Println("ExitElement:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) ExitReference(c *ReferenceContext) {
	fmt.Println("ExitReference:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) ExitAttribute(c *AttributeContext) {
	fmt.Println("ExitAttribute:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) ExitChardata(c *ChardataContext) {
	fmt.Println("ExitChardata:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) ExitMisc(c *MiscContext) {
	fmt.Println("ExitMisc:", c.GetRuleIndex(), c.GetText())
	
}

func (l *Listener) EnterEveryRule(c antlr.ParserRuleContext) {
	fmt.Println("EnterEveryRule:", c.GetRuleIndex(), c.GetText())
}
