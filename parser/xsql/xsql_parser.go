// Code generated from XSQLParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package xsql // XSQLParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type XSQLParser struct {
	*antlr.BaseParser
}

var XSQLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func xsqlparserParserInit() {
	staticData := &XSQLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "'<'", "", "", "'>'", "", "'/>'", "'/'",
		"'='",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "CDATA", "DTD", "EntityRef", "CharRef", "SEA_WS", "OPEN",
		"XMLDeclOpen", "TEXT", "CLOSE", "SPECIAL_CLOSE", "SLASH_CLOSE", "SLASH",
		"EQUALS", "STRING", "Name", "S", "PI",
	}
	staticData.RuleNames = []string{
		"content", "element", "reference", "attribute", "chardata",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 63, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 3, 0, 12, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 19, 8, 0,
		1, 0, 3, 0, 22, 8, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 1, 1,
		1, 1, 1, 5, 1, 32, 8, 1, 10, 1, 12, 1, 35, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9,
		1, 1, 1, 3, 1, 53, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 2, 1, 0, 4, 5, 2, 0, 6, 6, 9, 9, 67, 0,
		11, 1, 0, 0, 0, 2, 52, 1, 0, 0, 0, 4, 54, 1, 0, 0, 0, 6, 56, 1, 0, 0, 0,
		8, 60, 1, 0, 0, 0, 10, 12, 3, 8, 4, 0, 11, 10, 1, 0, 0, 0, 11, 12, 1, 0,
		0, 0, 12, 25, 1, 0, 0, 0, 13, 19, 3, 2, 1, 0, 14, 19, 3, 4, 2, 0, 15, 19,
		5, 2, 0, 0, 16, 19, 5, 18, 0, 0, 17, 19, 5, 1, 0, 0, 18, 13, 1, 0, 0, 0,
		18, 14, 1, 0, 0, 0, 18, 15, 1, 0, 0, 0, 18, 16, 1, 0, 0, 0, 18, 17, 1,
		0, 0, 0, 19, 21, 1, 0, 0, 0, 20, 22, 3, 8, 4, 0, 21, 20, 1, 0, 0, 0, 21,
		22, 1, 0, 0, 0, 22, 24, 1, 0, 0, 0, 23, 18, 1, 0, 0, 0, 24, 27, 1, 0, 0,
		0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 1, 1, 0, 0, 0, 27, 25, 1,
		0, 0, 0, 28, 29, 5, 7, 0, 0, 29, 33, 5, 16, 0, 0, 30, 32, 3, 6, 3, 0, 31,
		30, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0,
		0, 34, 36, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 37, 5, 10, 0, 0, 37, 38,
		3, 0, 0, 0, 38, 39, 5, 7, 0, 0, 39, 40, 5, 13, 0, 0, 40, 41, 5, 16, 0,
		0, 41, 42, 5, 10, 0, 0, 42, 53, 1, 0, 0, 0, 43, 44, 5, 7, 0, 0, 44, 48,
		5, 16, 0, 0, 45, 47, 3, 6, 3, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0,
		48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48, 1,
		0, 0, 0, 51, 53, 5, 12, 0, 0, 52, 28, 1, 0, 0, 0, 52, 43, 1, 0, 0, 0, 53,
		3, 1, 0, 0, 0, 54, 55, 7, 0, 0, 0, 55, 5, 1, 0, 0, 0, 56, 57, 5, 16, 0,
		0, 57, 58, 5, 14, 0, 0, 58, 59, 5, 15, 0, 0, 59, 7, 1, 0, 0, 0, 60, 61,
		7, 1, 0, 0, 61, 9, 1, 0, 0, 0, 7, 11, 18, 21, 25, 33, 48, 52,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// XSQLParserInit initializes any static state used to implement XSQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewXSQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func XSQLParserInit() {
	staticData := &XSQLParserParserStaticData
	staticData.once.Do(xsqlparserParserInit)
}

// NewXSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewXSQLParser(input antlr.TokenStream) *XSQLParser {
	XSQLParserInit()
	this := new(XSQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &XSQLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "XSQLParser.g4"

	return this
}

// XSQLParser tokens.
const (
	XSQLParserEOF           = antlr.TokenEOF
	XSQLParserCOMMENT       = 1
	XSQLParserCDATA         = 2
	XSQLParserDTD           = 3
	XSQLParserEntityRef     = 4
	XSQLParserCharRef       = 5
	XSQLParserSEA_WS        = 6
	XSQLParserOPEN          = 7
	XSQLParserXMLDeclOpen   = 8
	XSQLParserTEXT          = 9
	XSQLParserCLOSE         = 10
	XSQLParserSPECIAL_CLOSE = 11
	XSQLParserSLASH_CLOSE   = 12
	XSQLParserSLASH         = 13
	XSQLParserEQUALS        = 14
	XSQLParserSTRING        = 15
	XSQLParserName          = 16
	XSQLParserS             = 17
	XSQLParserPI            = 18
)

// XSQLParser rules.
const (
	XSQLParserRULE_content   = 0
	XSQLParserRULE_element   = 1
	XSQLParserRULE_reference = 2
	XSQLParserRULE_attribute = 3
	XSQLParserRULE_chardata  = 4
)

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChardata() []IChardataContext
	Chardata(i int) IChardataContext
	AllElement() []IElementContext
	Element(i int) IElementContext
	AllReference() []IReferenceContext
	Reference(i int) IReferenceContext
	AllCDATA() []antlr.TerminalNode
	CDATA(i int) antlr.TerminalNode
	AllPI() []antlr.TerminalNode
	PI(i int) antlr.TerminalNode
	AllCOMMENT() []antlr.TerminalNode
	COMMENT(i int) antlr.TerminalNode

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_content
	return p
}

func InitEmptyContentContext(p *ContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_content
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) AllChardata() []IChardataContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChardataContext); ok {
			len++
		}
	}

	tst := make([]IChardataContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChardataContext); ok {
			tst[i] = t.(IChardataContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Chardata(i int) IChardataContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChardataContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChardataContext)
}

func (s *ContentContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ContentContext) AllReference() []IReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReferenceContext); ok {
			len++
		}
	}

	tst := make([]IReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReferenceContext); ok {
			tst[i] = t.(IReferenceContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Reference(i int) IReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *ContentContext) AllCDATA() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserCDATA)
}

func (s *ContentContext) CDATA(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserCDATA, i)
}

func (s *ContentContext) AllPI() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserPI)
}

func (s *ContentContext) PI(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserPI, i)
}

func (s *ContentContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserCOMMENT)
}

func (s *ContentContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserCOMMENT, i)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.ExitContent(s)
	}
}

func (s *ContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XSQLParserVisitor:
		return t.VisitContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XSQLParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, XSQLParserRULE_content)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(10)
			p.Chardata()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(18)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(13)
					p.Element()
				}

			case 2:
				{
					p.SetState(14)
					p.Reference()
				}

			case 3:
				{
					p.SetState(15)
					p.Match(XSQLParserCDATA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 4:
				{
					p.SetState(16)
					p.Match(XSQLParserPI)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 5:
				{
					p.SetState(17)
					p.Match(XSQLParserCOMMENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}
			p.SetState(21)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(20)
					p.Chardata()
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOPEN() []antlr.TerminalNode
	OPEN(i int) antlr.TerminalNode
	AllName() []antlr.TerminalNode
	Name(i int) antlr.TerminalNode
	AllCLOSE() []antlr.TerminalNode
	CLOSE(i int) antlr.TerminalNode
	Content() IContentContext
	SLASH() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
	SLASH_CLOSE() antlr.TerminalNode

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_element
	return p
}

func InitEmptyElementContext(p *ElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_element
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) AllOPEN() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserOPEN)
}

func (s *ElementContext) OPEN(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN, i)
}

func (s *ElementContext) AllName() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserName)
}

func (s *ElementContext) Name(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserName, i)
}

func (s *ElementContext) AllCLOSE() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserCLOSE)
}

func (s *ElementContext) CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE, i)
}

func (s *ElementContext) Content() IContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *ElementContext) SLASH() antlr.TerminalNode {
	return s.GetToken(XSQLParserSLASH, 0)
}

func (s *ElementContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *ElementContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ElementContext) SLASH_CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserSLASH_CLOSE, 0)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XSQLParserVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XSQLParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, XSQLParserRULE_element)
	var _alt int

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Match(XSQLParserName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(30)
					p.Attribute()
				}

			}
			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Content()
		}
		{
			p.SetState(38)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(XSQLParserSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(XSQLParserName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(XSQLParserName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(45)
					p.Attribute()
				}

			}
			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(XSQLParserSLASH_CLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EntityRef() antlr.TerminalNode
	CharRef() antlr.TerminalNode

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) EntityRef() antlr.TerminalNode {
	return s.GetToken(XSQLParserEntityRef, 0)
}

func (s *ReferenceContext) CharRef() antlr.TerminalNode {
	return s.GetToken(XSQLParserCharRef, 0)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.ExitReference(s)
	}
}

func (s *ReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XSQLParserVisitor:
		return t.VisitReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XSQLParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, XSQLParserRULE_reference)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		_la = p.GetTokenStream().LA(1)

		if !(_la == XSQLParserEntityRef || _la == XSQLParserCharRef) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_attribute
	return p
}

func InitEmptyAttributeContext(p *AttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_attribute
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) Name() antlr.TerminalNode {
	return s.GetToken(XSQLParserName, 0)
}

func (s *AttributeContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(XSQLParserEQUALS, 0)
}

func (s *AttributeContext) STRING() antlr.TerminalNode {
	return s.GetToken(XSQLParserSTRING, 0)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (s *AttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XSQLParserVisitor:
		return t.VisitAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XSQLParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, XSQLParserRULE_attribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(XSQLParserName)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(XSQLParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(XSQLParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChardataContext is an interface to support dynamic dispatch.
type IChardataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEA_WS() antlr.TerminalNode
	TEXT() antlr.TerminalNode

	// IsChardataContext differentiates from other interfaces.
	IsChardataContext()
}

type ChardataContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChardataContext() *ChardataContext {
	var p = new(ChardataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_chardata
	return p
}

func InitEmptyChardataContext(p *ChardataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_chardata
}

func (*ChardataContext) IsChardataContext() {}

func NewChardataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChardataContext {
	var p = new(ChardataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_chardata

	return p
}

func (s *ChardataContext) GetParser() antlr.Parser { return s.parser }

func (s *ChardataContext) SEA_WS() antlr.TerminalNode {
	return s.GetToken(XSQLParserSEA_WS, 0)
}

func (s *ChardataContext) TEXT() antlr.TerminalNode {
	return s.GetToken(XSQLParserTEXT, 0)
}

func (s *ChardataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChardataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChardataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.EnterChardata(s)
	}
}

func (s *ChardataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XSQLParserListener); ok {
		listenerT.ExitChardata(s)
	}
}

func (s *ChardataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XSQLParserVisitor:
		return t.VisitChardata(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XSQLParser) Chardata() (localctx IChardataContext) {
	localctx = NewChardataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, XSQLParserRULE_chardata)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		_la = p.GetTokenStream().LA(1)

		if !(_la == XSQLParserSEA_WS || _la == XSQLParserTEXT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
