package gobatis

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/koyeo/gobatis/dtd"
	"github.com/koyeo/gobatis/parser/xml"
	"strings"
	"sync"
)

func parseConfig(db *DB, file string, content []byte) (err error) {
	listener := &xmlParser{
		file:          file,
		stack:         newXMLStack(),
		rootElement:   dtd.Configuration,
		elementGetter: dtd.ConfigElement,
	}
	err = parseNode(listener, content)
	if err != nil {
		return
	}
	
	d, _ := json.MarshalIndent(listener.rootNode, "", "\t")
	fmt.Println(string(d))
	return
}

func parseMapper(db *DB, fileName string, content []byte) {

}

func parseNode(listener *xmlParser, content []byte) (err error) {
	lexer := xml.NewXMLLexer(antlr.NewInputStream(string(content)))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := xml.NewXMLParser(stream)
	parser.BuildParseTrees = true
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//parser.SetErrorHandler()
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Document())
	if listener.error != nil {
		err = listener.error
		return
	}
	return
}

func newXMLNode(file, name string, token antlr.Token) *xmlNode {
	return &xmlNode{
		File:  file,
		Name:  name,
		start: token,
	}
}

type xmlNode struct {
	File       string `json:"-"`
	Name       string
	Text       string
	Attributes map[string]*xmlNodeAttribute
	Nodes      []*xmlNode
	nodesCount map[string]int
	start      antlr.Token
}

type xmlNodeAttribute struct {
	File  string      `json:"-"`
	Start antlr.Token `json:"-"`
	Value string
}

func (p *xmlNode) AddAttribute(name string, value *xmlNodeAttribute) {
	if p.Attributes == nil {
		p.Attributes = map[string]*xmlNodeAttribute{}
	}
	p.Attributes[name] = value
}

func (p *xmlNode) HasAttribute(name string) bool {
	if p.Attributes == nil {
		return false
	}
	_, ok := p.Attributes[name]
	return ok
}

func (p *xmlNode) AddNode(node *xmlNode) {
	if p.nodesCount == nil {
		p.nodesCount = map[string]int{}
	}
	p.nodesCount[node.Name]++
	p.Nodes = append(p.Nodes, node)
}

func (p *xmlNode) countNode(name string) int {
	if p.nodesCount == nil {
		return 0
	}
	return p.nodesCount[name]
}

type xmlNodeStack struct {
	list *list.List
	lock *sync.RWMutex
}

func newXMLStack() *xmlNodeStack {
	l := list.New()
	lock := &sync.RWMutex{}
	return &xmlNodeStack{l, lock}
}

func (stack *xmlNodeStack) Push(value *xmlNode) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}

func (stack *xmlNodeStack) Pop() *xmlNode {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value.(*xmlNode)
	}
	return nil
}

func (stack *xmlNodeStack) Peak() *xmlNode {
	e := stack.list.Back()
	if e != nil {
		return e.Value.(*xmlNode)
	}
	return nil
}

func (stack *xmlNodeStack) Len() int {
	return stack.list.Len()
}

func (stack *xmlNodeStack) Empty() bool {
	return stack.list.Len() == 0
}

type xmlParser struct {
	file          string
	content       []byte
	depth         int
	error         error
	stack         *xmlNodeStack
	rootNode      *xmlNode
	rootElement   *dtd.Element
	elementGetter func(name string) (elem *dtd.Element, err error)
}

func (p *xmlParser) trimAttributeValueQuote(val string) string {
	if strings.HasPrefix(val, "'") || strings.HasSuffix(val, "'") {
		val = strings.TrimPrefix(val, "'")
		return strings.TrimSuffix(val, "'")
	}
	if strings.HasPrefix(val, "\"") || strings.HasSuffix(val, "\"") {
		val = strings.TrimPrefix(val, "\"")
		return strings.TrimSuffix(val, "\"")
	}
	return val
}

func (p *xmlParser) validateNode(node *xmlNode, elem *dtd.Element) {
	
	// check required attributes
	if elem.Attributes != nil {
		for k, v := range elem.Attributes {
			if v == dtd.REQUIRED && !node.HasAttribute(k) {
				p.setError(fmt.Sprintf("element: %s miss required attribute: %s", node.Name, k), node.start)
				return
			}
		}
	}
	
	for _, childNode := range node.Nodes {
		
		// check not supported node
		if !elem.HasNode(childNode.Name) {
			p.setError(
				fmt.Sprintf("element: %s not support child element: %s", node.Name, childNode.Name),
				childNode.start,
			)
			return
		}
		
		// check at most once node
		if elem.GetNodeCount(childNode.Name) == dtd.AT_MOST_ONCE && node.countNode(childNode.Name) > 1 {
			p.setError(
				fmt.Sprintf("element: %s not support duplicate element: %s", node.Name, childNode.Name),
				childNode.start,
			)
			return
		}
		
		childElem, err := p.elementGetter(childNode.Name)
		if err != nil {
			p.setError(err.Error(), childNode.start)
			return
		}
		
		p.validateNode(childNode, childElem)
		if p.error != nil {
			return
		}
	}
	
	// check at least once node
	if elem.Nodes != nil {
		for k, v := range elem.Nodes {
			if v == dtd.AT_LEAST_ONCE && node.countNode(k) == 0 {
				p.setError(fmt.Sprintf("element %s miss required element %s", node.Name, k), node.start)
				return
			}
		}
	}
}

func (p *xmlParser) setError(msg string, token antlr.Token) {
	p.error = fmt.Errorf("%s line %d:%d, parse error: %s", p.file, token.GetLine(), token.GetColumn(), msg)
}

func (p *xmlParser) EnterElement(c *xml.ElementContext) {
	if p.error != nil {
		return
	}
	name := c.Name(0)
	if p.depth == 0 {
		if name.GetText() != p.rootElement.Name {
			p.setError(fmt.Sprintf("first level tag %s unsupported", name.GetText()), name.GetSymbol())
			return
		}
	}
	p.stack.Push(newXMLNode(p.file, name.GetText(), name.GetSymbol()))
	p.depth++
}

func (p *xmlParser) ExitElement(_ *xml.ElementContext) {
	if p.error != nil || p.stack.Peak() == nil {
		return
	}
	p.depth--
	child := p.stack.Pop()
	
	if p.stack.Len() > 0 {
		p.stack.Peak().AddNode(child)
	} else {
		p.rootNode = child
		p.validateNode(child, p.rootElement)
	}
}

func (p *xmlParser) EnterAttribute(c *xml.AttributeContext) {
	// <?xml version="1.0" encoding="UTF-8" ?>
	if p.error != nil || p.stack.Peak() == nil {
		return
	}
	p.stack.Peak().AddAttribute(c.Name().GetText(), &xmlNodeAttribute{
		File:  p.file,
		Start: c.STRING().GetSymbol(),
		Value: p.trimAttributeValueQuote(c.STRING().GetText()),
	})
}

func (p *xmlParser) EnterChardata(c *xml.ChardataContext) {
	if p.error != nil || p.stack.Peak() == nil {
		return
	}
	text := strings.TrimSpace(c.GetText())
	if text != "" {
		p.stack.Peak().AddNode(&xmlNode{File: p.file, Text: text, start: c.GetStart()})
	}
}

func (p *xmlParser) VisitTerminal(_ antlr.TerminalNode) {
	// pass
}

func (p *xmlParser) VisitErrorNode(_ antlr.ErrorNode) {
	// pass
}

func (p *xmlParser) EnterEveryRule(_ antlr.ParserRuleContext) {
	// pass
}

func (p *xmlParser) ExitEveryRule(_ antlr.ParserRuleContext) {
	// pass
}

func (p *xmlParser) EnterDocument(_ *xml.DocumentContext) {
	// pass
}

func (p *xmlParser) EnterProlog(_ *xml.PrologContext) {
	// pass
}

func (p *xmlParser) EnterContent(_ *xml.ContentContext) {
	// pass
}

func (p *xmlParser) EnterReference(_ *xml.ReferenceContext) {
	// pass
}

func (p *xmlParser) ExitReference(_ *xml.ReferenceContext) {
	// pass
}

func (p *xmlParser) ExitAttribute(_ *xml.AttributeContext) {
	// pass
}

func (p *xmlParser) ExitChardata(_ *xml.ChardataContext) {
	// pass
}

func (p *xmlParser) ExitMisc(_ *xml.MiscContext) {
	// pass
}

func (p *xmlParser) EnterMisc(_ *xml.MiscContext) {
	// pass
}

func (p *xmlParser) ExitDocument(_ *xml.DocumentContext) {
	// pass
}

func (p *xmlParser) ExitProlog(_ *xml.PrologContext) {
	// pass
}

func (p *xmlParser) ExitContent(_ *xml.ContentContext) {
	// pass
}
