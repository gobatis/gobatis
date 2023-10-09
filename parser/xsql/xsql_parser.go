// Code generated from XSQLParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "", "", "", "'<'", "", "", "", "", "'}'", "", "", "", "'>'", "'/>'",
		"'/'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "EntityRef", "SEA_WS", "OPEN", "EXPR_OPEN", "TEXT", "DOLLAR_NOT_LBRACE",
		"HASH_NOT_LBRACE", "EXPR_CLOSE", "EXPR_VAL", "S1", "Name", "CLOSE",
		"SLASH_CLOSE", "SLASH", "EQUALS", "STRING", "S2",
	}
	staticData.RuleNames = []string{
		"content", "element", "placeholder", "reference", "attribute", "chardata",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 71, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 3, 0, 14, 8, 0, 1, 0, 1, 0, 1, 0, 3, 0, 19, 8, 0,
		1, 0, 3, 0, 22, 8, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 1, 1,
		1, 1, 1, 5, 1, 32, 8, 1, 10, 1, 12, 1, 35, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9,
		1, 1, 1, 3, 1, 53, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 69, 8, 5, 1, 5, 0, 0, 6, 0, 2,
		4, 6, 8, 10, 0, 0, 74, 0, 13, 1, 0, 0, 0, 2, 52, 1, 0, 0, 0, 4, 54, 1,
		0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 60, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12,
		14, 3, 10, 5, 0, 13, 12, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 25, 1, 0,
		0, 0, 15, 19, 3, 4, 2, 0, 16, 19, 3, 2, 1, 0, 17, 19, 3, 6, 3, 0, 18, 15,
		1, 0, 0, 0, 18, 16, 1, 0, 0, 0, 18, 17, 1, 0, 0, 0, 19, 21, 1, 0, 0, 0,
		20, 22, 3, 10, 5, 0, 21, 20, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 24, 1,
		0, 0, 0, 23, 18, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25,
		26, 1, 0, 0, 0, 26, 1, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 5, 4, 0,
		0, 29, 33, 5, 12, 0, 0, 30, 32, 3, 8, 4, 0, 31, 30, 1, 0, 0, 0, 32, 35,
		1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0,
		35, 33, 1, 0, 0, 0, 36, 37, 5, 13, 0, 0, 37, 38, 3, 0, 0, 0, 38, 39, 5,
		4, 0, 0, 39, 40, 5, 15, 0, 0, 40, 41, 5, 12, 0, 0, 41, 42, 5, 13, 0, 0,
		42, 53, 1, 0, 0, 0, 43, 44, 5, 4, 0, 0, 44, 48, 5, 12, 0, 0, 45, 47, 3,
		8, 4, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48,
		49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 53, 5, 14,
		0, 0, 52, 28, 1, 0, 0, 0, 52, 43, 1, 0, 0, 0, 53, 3, 1, 0, 0, 0, 54, 55,
		5, 5, 0, 0, 55, 56, 5, 10, 0, 0, 56, 57, 5, 9, 0, 0, 57, 5, 1, 0, 0, 0,
		58, 59, 5, 2, 0, 0, 59, 7, 1, 0, 0, 0, 60, 61, 5, 12, 0, 0, 61, 62, 5,
		16, 0, 0, 62, 63, 5, 17, 0, 0, 63, 9, 1, 0, 0, 0, 64, 69, 5, 3, 0, 0, 65,
		66, 4, 5, 0, 0, 66, 69, 5, 7, 0, 0, 67, 69, 5, 6, 0, 0, 68, 64, 1, 0, 0,
		0, 68, 65, 1, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 11, 1, 0, 0, 0, 8, 13, 18,
		21, 25, 33, 48, 52, 68,
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
	XSQLParserEOF               = antlr.TokenEOF
	XSQLParserCOMMENT           = 1
	XSQLParserEntityRef         = 2
	XSQLParserSEA_WS            = 3
	XSQLParserOPEN              = 4
	XSQLParserEXPR_OPEN         = 5
	XSQLParserTEXT              = 6
	XSQLParserDOLLAR_NOT_LBRACE = 7
	XSQLParserHASH_NOT_LBRACE   = 8
	XSQLParserEXPR_CLOSE        = 9
	XSQLParserEXPR_VAL          = 10
	XSQLParserS1                = 11
	XSQLParserName              = 12
	XSQLParserCLOSE             = 13
	XSQLParserSLASH_CLOSE       = 14
	XSQLParserSLASH             = 15
	XSQLParserEQUALS            = 16
	XSQLParserSTRING            = 17
	XSQLParserS2                = 18
)

// XSQLParser rules.
const (
	XSQLParserRULE_content     = 0
	XSQLParserRULE_element     = 1
	XSQLParserRULE_placeholder = 2
	XSQLParserRULE_reference   = 3
	XSQLParserRULE_attribute   = 4
	XSQLParserRULE_chardata    = 5
)

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChardata() []IChardataContext
	Chardata(i int) IChardataContext
	AllPlaceholder() []IPlaceholderContext
	Placeholder(i int) IPlaceholderContext
	AllElement() []IElementContext
	Element(i int) IElementContext
	AllReference() []IReferenceContext
	Reference(i int) IReferenceContext

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

func (s *ContentContext) AllPlaceholder() []IPlaceholderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPlaceholderContext); ok {
			len++
		}
	}

	tst := make([]IPlaceholderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPlaceholderContext); ok {
			tst[i] = t.(IPlaceholderContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Placeholder(i int) IPlaceholderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlaceholderContext); ok {
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

	return t.(IPlaceholderContext)
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

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, XSQLParserRULE_content)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(12)
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
					p.SetState(15)
					p.Placeholder()
				}

			case 2:
				{
					p.SetState(16)
					p.Element()
				}

			case 3:
				{
					p.SetState(17)
					p.Reference()
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

// IPlaceholderContext is an interface to support dynamic dispatch.
type IPlaceholderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXPR_OPEN() antlr.TerminalNode
	EXPR_VAL() antlr.TerminalNode
	EXPR_CLOSE() antlr.TerminalNode

	// IsPlaceholderContext differentiates from other interfaces.
	IsPlaceholderContext()
}

type PlaceholderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlaceholderContext() *PlaceholderContext {
	var p = new(PlaceholderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_placeholder
	return p
}

func InitEmptyPlaceholderContext(p *PlaceholderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_placeholder
}

func (*PlaceholderContext) IsPlaceholderContext() {}

func NewPlaceholderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlaceholderContext {
	var p = new(PlaceholderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_placeholder

	return p
}

func (s *PlaceholderContext) GetParser() antlr.Parser { return s.parser }

func (s *PlaceholderContext) EXPR_OPEN() antlr.TerminalNode {
	return s.GetToken(XSQLParserEXPR_OPEN, 0)
}

func (s *PlaceholderContext) EXPR_VAL() antlr.TerminalNode {
	return s.GetToken(XSQLParserEXPR_VAL, 0)
}

func (s *PlaceholderContext) EXPR_CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserEXPR_CLOSE, 0)
}

func (s *PlaceholderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlaceholderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Placeholder() (localctx IPlaceholderContext) {
	localctx = NewPlaceholderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, XSQLParserRULE_placeholder)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(XSQLParserEXPR_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(XSQLParserEXPR_VAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(XSQLParserEXPR_CLOSE)
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

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EntityRef() antlr.TerminalNode

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

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, XSQLParserRULE_reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(XSQLParserEntityRef)
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

func (p *XSQLParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, XSQLParserRULE_attribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(XSQLParserName)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(XSQLParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
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
	DOLLAR_NOT_LBRACE() antlr.TerminalNode
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

func (s *ChardataContext) DOLLAR_NOT_LBRACE() antlr.TerminalNode {
	return s.GetToken(XSQLParserDOLLAR_NOT_LBRACE, 0)
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

func (p *XSQLParser) Chardata() (localctx IChardataContext) {
	localctx = NewChardataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, XSQLParserRULE_chardata)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(XSQLParserSEA_WS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(65)

		if !(p.GetInputStream().LA(1) == '$' && p.GetInputStream().LA(2) == '{') {
			p.SetError(antlr.NewFailedPredicateException(p, " p.GetInputStream().LA(1) == '$' && p.GetInputStream().LA(2) == '{' ", ""))
			goto errorExit
		}
		{
			p.SetState(66)
			p.Match(XSQLParserDOLLAR_NOT_LBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(XSQLParserTEXT)
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

func (p *XSQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ChardataContext = nil
		if localctx != nil {
			t = localctx.(*ChardataContext)
		}
		return p.Chardata_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *XSQLParser) Chardata_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.GetInputStream().LA(1) == '$' && p.GetInputStream().LA(2) == '{'

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
