// Code generated from XSQL.g4 by ANTLR 4.13.1. DO NOT EDIT.

package xsql // XSQL
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

var XSQLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func xsqlParserInit() {
	staticData := &XSQLParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "'$'", "'#'", "'{'", "'}'", "'<'", "'>'", "'/'",
		"'='",
	}
	staticData.SymbolicNames = []string{
		"", "BLOCK_COMMENT", "LINE_COMMENT", "EntityRef", "WS", "NAME", "DOLLAR",
		"HASH", "OPEN_CURLY_BRAXE", "CLOSE_CURLY_BRAXE", "OPEN", "CLOSE", "SLASH",
		"EQUALS", "STRING", "TEXT",
	}
	staticData.RuleNames = []string{
		"content", "tagStart", "tagEnd", "closeTag", "attribute", "expr", "reference",
		"chardata",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 5, 0, 23, 8, 0, 10, 0, 12, 0, 26, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 31,
		8, 1, 10, 1, 12, 1, 34, 9, 1, 1, 1, 5, 1, 37, 8, 1, 10, 1, 12, 1, 40, 9,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 52,
		8, 3, 10, 3, 12, 3, 55, 9, 3, 1, 3, 5, 3, 58, 8, 3, 10, 3, 12, 3, 61, 9,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 70, 8, 4, 10, 4, 12,
		4, 73, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 79, 8, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1,
		1, 0, 4, 15, 91, 0, 24, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 43, 1, 0, 0,
		0, 6, 48, 1, 0, 0, 0, 8, 65, 1, 0, 0, 0, 10, 78, 1, 0, 0, 0, 12, 83, 1,
		0, 0, 0, 14, 85, 1, 0, 0, 0, 16, 23, 3, 2, 1, 0, 17, 23, 3, 4, 2, 0, 18,
		23, 3, 6, 3, 0, 19, 23, 3, 10, 5, 0, 20, 23, 3, 12, 6, 0, 21, 23, 3, 14,
		7, 0, 22, 16, 1, 0, 0, 0, 22, 17, 1, 0, 0, 0, 22, 18, 1, 0, 0, 0, 22, 19,
		1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 21, 1, 0, 0, 0, 23, 26, 1, 0, 0, 0,
		24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 1, 1, 0, 0, 0, 26, 24, 1, 0,
		0, 0, 27, 28, 5, 10, 0, 0, 28, 32, 5, 5, 0, 0, 29, 31, 5, 4, 0, 0, 30,
		29, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0,
		0, 33, 38, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 37, 3, 8, 4, 0, 36, 35,
		1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0,
		39, 41, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 42, 5, 11, 0, 0, 42, 3, 1,
		0, 0, 0, 43, 44, 5, 10, 0, 0, 44, 45, 5, 12, 0, 0, 45, 46, 5, 5, 0, 0,
		46, 47, 5, 11, 0, 0, 47, 5, 1, 0, 0, 0, 48, 49, 5, 10, 0, 0, 49, 53, 5,
		5, 0, 0, 50, 52, 5, 4, 0, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53,
		51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 59, 1, 0, 0, 0, 55, 53, 1, 0, 0,
		0, 56, 58, 3, 8, 4, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57,
		1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0,
		62, 63, 5, 12, 0, 0, 63, 64, 5, 11, 0, 0, 64, 7, 1, 0, 0, 0, 65, 66, 5,
		5, 0, 0, 66, 67, 5, 13, 0, 0, 67, 71, 5, 14, 0, 0, 68, 70, 5, 4, 0, 0,
		69, 68, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1,
		0, 0, 0, 72, 9, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 5, 6, 0, 0, 75,
		79, 5, 8, 0, 0, 76, 77, 5, 7, 0, 0, 77, 79, 5, 8, 0, 0, 78, 74, 1, 0, 0,
		0, 78, 76, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 3, 14, 7, 0, 81, 82,
		5, 9, 0, 0, 82, 11, 1, 0, 0, 0, 83, 84, 5, 3, 0, 0, 84, 13, 1, 0, 0, 0,
		85, 86, 7, 0, 0, 0, 86, 15, 1, 0, 0, 0, 8, 22, 24, 32, 38, 53, 59, 71,
		78,
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
	staticData := &XSQLParserStaticData
	staticData.once.Do(xsqlParserInit)
}

// NewXSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewXSQLParser(input antlr.TokenStream) *XSQLParser {
	XSQLParserInit()
	this := new(XSQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &XSQLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "XSQL.g4"

	return this
}

// XSQLParser tokens.
const (
	XSQLParserEOF               = antlr.TokenEOF
	XSQLParserBLOCK_COMMENT     = 1
	XSQLParserLINE_COMMENT      = 2
	XSQLParserEntityRef         = 3
	XSQLParserWS                = 4
	XSQLParserNAME              = 5
	XSQLParserDOLLAR            = 6
	XSQLParserHASH              = 7
	XSQLParserOPEN_CURLY_BRAXE  = 8
	XSQLParserCLOSE_CURLY_BRAXE = 9
	XSQLParserOPEN              = 10
	XSQLParserCLOSE             = 11
	XSQLParserSLASH             = 12
	XSQLParserEQUALS            = 13
	XSQLParserSTRING            = 14
	XSQLParserTEXT              = 15
)

// XSQLParser rules.
const (
	XSQLParserRULE_content   = 0
	XSQLParserRULE_tagStart  = 1
	XSQLParserRULE_tagEnd    = 2
	XSQLParserRULE_closeTag  = 3
	XSQLParserRULE_attribute = 4
	XSQLParserRULE_expr      = 5
	XSQLParserRULE_reference = 6
	XSQLParserRULE_chardata  = 7
)

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTagStart() []ITagStartContext
	TagStart(i int) ITagStartContext
	AllTagEnd() []ITagEndContext
	TagEnd(i int) ITagEndContext
	AllCloseTag() []ICloseTagContext
	CloseTag(i int) ICloseTagContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllReference() []IReferenceContext
	Reference(i int) IReferenceContext
	AllChardata() []IChardataContext
	Chardata(i int) IChardataContext

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

func (s *ContentContext) AllTagStart() []ITagStartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagStartContext); ok {
			len++
		}
	}

	tst := make([]ITagStartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagStartContext); ok {
			tst[i] = t.(ITagStartContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) TagStart(i int) ITagStartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagStartContext); ok {
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

	return t.(ITagStartContext)
}

func (s *ContentContext) AllTagEnd() []ITagEndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagEndContext); ok {
			len++
		}
	}

	tst := make([]ITagEndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagEndContext); ok {
			tst[i] = t.(ITagEndContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) TagEnd(i int) ITagEndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagEndContext); ok {
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

	return t.(ITagEndContext)
}

func (s *ContentContext) AllCloseTag() []ICloseTagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICloseTagContext); ok {
			len++
		}
	}

	tst := make([]ICloseTagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICloseTagContext); ok {
			tst[i] = t.(ICloseTagContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) CloseTag(i int) ICloseTagContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICloseTagContext); ok {
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

	return t.(ICloseTagContext)
}

func (s *ContentContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
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

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, XSQLParserRULE_content)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65528) != 0 {
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(16)
				p.TagStart()
			}

		case 2:
			{
				p.SetState(17)
				p.TagEnd()
			}

		case 3:
			{
				p.SetState(18)
				p.CloseTag()
			}

		case 4:
			{
				p.SetState(19)
				p.Expr()
			}

		case 5:
			{
				p.SetState(20)
				p.Reference()
			}

		case 6:
			{
				p.SetState(21)
				p.Chardata()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ITagStartContext is an interface to support dynamic dispatch.
type ITagStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN() antlr.TerminalNode
	NAME() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsTagStartContext differentiates from other interfaces.
	IsTagStartContext()
}

type TagStartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagStartContext() *TagStartContext {
	var p = new(TagStartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_tagStart
	return p
}

func InitEmptyTagStartContext(p *TagStartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_tagStart
}

func (*TagStartContext) IsTagStartContext() {}

func NewTagStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagStartContext {
	var p = new(TagStartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_tagStart

	return p
}

func (s *TagStartContext) GetParser() antlr.Parser { return s.parser }

func (s *TagStartContext) OPEN() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN, 0)
}

func (s *TagStartContext) NAME() antlr.TerminalNode {
	return s.GetToken(XSQLParserNAME, 0)
}

func (s *TagStartContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE, 0)
}

func (s *TagStartContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserWS)
}

func (s *TagStartContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserWS, i)
}

func (s *TagStartContext) AllAttribute() []IAttributeContext {
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

func (s *TagStartContext) Attribute(i int) IAttributeContext {
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

func (s *TagStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) TagStart() (localctx ITagStartContext) {
	localctx = NewTagStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, XSQLParserRULE_tagStart)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(XSQLParserOPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(XSQLParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserWS {
		{
			p.SetState(29)
			p.Match(XSQLParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserNAME {
		{
			p.SetState(35)
			p.Attribute()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(41)
		p.Match(XSQLParserCLOSE)
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

// ITagEndContext is an interface to support dynamic dispatch.
type ITagEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	NAME() antlr.TerminalNode
	CLOSE() antlr.TerminalNode

	// IsTagEndContext differentiates from other interfaces.
	IsTagEndContext()
}

type TagEndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagEndContext() *TagEndContext {
	var p = new(TagEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_tagEnd
	return p
}

func InitEmptyTagEndContext(p *TagEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_tagEnd
}

func (*TagEndContext) IsTagEndContext() {}

func NewTagEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagEndContext {
	var p = new(TagEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_tagEnd

	return p
}

func (s *TagEndContext) GetParser() antlr.Parser { return s.parser }

func (s *TagEndContext) OPEN() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN, 0)
}

func (s *TagEndContext) SLASH() antlr.TerminalNode {
	return s.GetToken(XSQLParserSLASH, 0)
}

func (s *TagEndContext) NAME() antlr.TerminalNode {
	return s.GetToken(XSQLParserNAME, 0)
}

func (s *TagEndContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE, 0)
}

func (s *TagEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) TagEnd() (localctx ITagEndContext) {
	localctx = NewTagEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, XSQLParserRULE_tagEnd)
	p.EnterOuterAlt(localctx, 1)
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
		p.Match(XSQLParserSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(XSQLParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(XSQLParserCLOSE)
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

// ICloseTagContext is an interface to support dynamic dispatch.
type ICloseTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN() antlr.TerminalNode
	NAME() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsCloseTagContext differentiates from other interfaces.
	IsCloseTagContext()
}

type CloseTagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCloseTagContext() *CloseTagContext {
	var p = new(CloseTagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_closeTag
	return p
}

func InitEmptyCloseTagContext(p *CloseTagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_closeTag
}

func (*CloseTagContext) IsCloseTagContext() {}

func NewCloseTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CloseTagContext {
	var p = new(CloseTagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_closeTag

	return p
}

func (s *CloseTagContext) GetParser() antlr.Parser { return s.parser }

func (s *CloseTagContext) OPEN() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN, 0)
}

func (s *CloseTagContext) NAME() antlr.TerminalNode {
	return s.GetToken(XSQLParserNAME, 0)
}

func (s *CloseTagContext) SLASH() antlr.TerminalNode {
	return s.GetToken(XSQLParserSLASH, 0)
}

func (s *CloseTagContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE, 0)
}

func (s *CloseTagContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserWS)
}

func (s *CloseTagContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserWS, i)
}

func (s *CloseTagContext) AllAttribute() []IAttributeContext {
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

func (s *CloseTagContext) Attribute(i int) IAttributeContext {
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

func (s *CloseTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) CloseTag() (localctx ICloseTagContext) {
	localctx = NewCloseTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, XSQLParserRULE_closeTag)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(XSQLParserOPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(XSQLParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserWS {
		{
			p.SetState(50)
			p.Match(XSQLParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserNAME {
		{
			p.SetState(56)
			p.Attribute()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(XSQLParserSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(XSQLParserCLOSE)
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
	NAME() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	STRING() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

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

func (s *AttributeContext) NAME() antlr.TerminalNode {
	return s.GetToken(XSQLParserNAME, 0)
}

func (s *AttributeContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(XSQLParserEQUALS, 0)
}

func (s *AttributeContext) STRING() antlr.TerminalNode {
	return s.GetToken(XSQLParserSTRING, 0)
}

func (s *AttributeContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserWS)
}

func (s *AttributeContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserWS, i)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(XSQLParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(XSQLParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(XSQLParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserWS {
		{
			p.SetState(68)
			p.Match(XSQLParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Chardata() IChardataContext
	CLOSE_CURLY_BRAXE() antlr.TerminalNode
	DOLLAR() antlr.TerminalNode
	OPEN_CURLY_BRAXE() antlr.TerminalNode
	HASH() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Chardata() IChardataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChardataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChardataContext)
}

func (s *ExprContext) CLOSE_CURLY_BRAXE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE_CURLY_BRAXE, 0)
}

func (s *ExprContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(XSQLParserDOLLAR, 0)
}

func (s *ExprContext) OPEN_CURLY_BRAXE() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN_CURLY_BRAXE, 0)
}

func (s *ExprContext) HASH() antlr.TerminalNode {
	return s.GetToken(XSQLParserHASH, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, XSQLParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case XSQLParserDOLLAR:
		{
			p.SetState(74)
			p.Match(XSQLParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(XSQLParserOPEN_CURLY_BRAXE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case XSQLParserHASH:
		{
			p.SetState(76)
			p.Match(XSQLParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(XSQLParserOPEN_CURLY_BRAXE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(80)
		p.Chardata()
	}
	{
		p.SetState(81)
		p.Match(XSQLParserCLOSE_CURLY_BRAXE)
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
	p.EnterRule(localctx, 12, XSQLParserRULE_reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
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

// IChardataContext is an interface to support dynamic dispatch.
type IChardataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLOSE() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	OPEN() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	DOLLAR() antlr.TerminalNode
	HASH() antlr.TerminalNode
	OPEN_CURLY_BRAXE() antlr.TerminalNode
	CLOSE_CURLY_BRAXE() antlr.TerminalNode
	WS() antlr.TerminalNode
	NAME() antlr.TerminalNode
	STRING() antlr.TerminalNode
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

func (s *ChardataContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE, 0)
}

func (s *ChardataContext) SLASH() antlr.TerminalNode {
	return s.GetToken(XSQLParserSLASH, 0)
}

func (s *ChardataContext) OPEN() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN, 0)
}

func (s *ChardataContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(XSQLParserEQUALS, 0)
}

func (s *ChardataContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(XSQLParserDOLLAR, 0)
}

func (s *ChardataContext) HASH() antlr.TerminalNode {
	return s.GetToken(XSQLParserHASH, 0)
}

func (s *ChardataContext) OPEN_CURLY_BRAXE() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN_CURLY_BRAXE, 0)
}

func (s *ChardataContext) CLOSE_CURLY_BRAXE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE_CURLY_BRAXE, 0)
}

func (s *ChardataContext) WS() antlr.TerminalNode {
	return s.GetToken(XSQLParserWS, 0)
}

func (s *ChardataContext) NAME() antlr.TerminalNode {
	return s.GetToken(XSQLParserNAME, 0)
}

func (s *ChardataContext) STRING() antlr.TerminalNode {
	return s.GetToken(XSQLParserSTRING, 0)
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
	p.EnterRule(localctx, 14, XSQLParserRULE_chardata)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65520) != 0) {
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
