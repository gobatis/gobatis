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
		"content", "start", "end", "attribute", "expr", "reference", "chardata",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 108, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 20, 8, 0,
		10, 0, 12, 0, 23, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 28, 8, 1, 10, 1, 12, 1,
		31, 9, 1, 1, 1, 5, 1, 34, 8, 1, 10, 1, 12, 1, 37, 9, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 43, 8, 1, 10, 1, 12, 1, 46, 9, 1, 1, 1, 5, 1, 49, 8, 1,
		10, 1, 12, 1, 52, 9, 1, 1, 1, 1, 1, 3, 1, 56, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 67, 8, 3, 10, 3, 12, 3, 70, 9, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 76, 8, 4, 1, 4, 5, 4, 79, 8, 4, 10, 4, 12,
		4, 82, 9, 4, 1, 4, 5, 4, 85, 8, 4, 10, 4, 12, 4, 88, 9, 4, 1, 4, 5, 4,
		91, 8, 4, 10, 4, 12, 4, 94, 9, 4, 1, 4, 5, 4, 97, 8, 4, 10, 4, 12, 4, 100,
		9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8,
		10, 12, 0, 1, 1, 0, 4, 15, 116, 0, 21, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0, 4,
		57, 1, 0, 0, 0, 6, 62, 1, 0, 0, 0, 8, 75, 1, 0, 0, 0, 10, 103, 1, 0, 0,
		0, 12, 105, 1, 0, 0, 0, 14, 20, 3, 2, 1, 0, 15, 20, 3, 4, 2, 0, 16, 20,
		3, 8, 4, 0, 17, 20, 3, 10, 5, 0, 18, 20, 3, 12, 6, 0, 19, 14, 1, 0, 0,
		0, 19, 15, 1, 0, 0, 0, 19, 16, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 18,
		1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0,
		22, 1, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 24, 25, 5, 10, 0, 0, 25, 29, 5,
		5, 0, 0, 26, 28, 5, 4, 0, 0, 27, 26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29,
		27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 35, 1, 0, 0, 0, 31, 29, 1, 0, 0,
		0, 32, 34, 3, 6, 3, 0, 33, 32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33,
		1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0,
		38, 56, 5, 11, 0, 0, 39, 40, 5, 10, 0, 0, 40, 44, 5, 5, 0, 0, 41, 43, 5,
		4, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44,
		45, 1, 0, 0, 0, 45, 50, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 49, 3, 6, 3,
		0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51,
		1, 0, 0, 0, 51, 53, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 54, 5, 12, 0, 0,
		54, 56, 5, 11, 0, 0, 55, 24, 1, 0, 0, 0, 55, 39, 1, 0, 0, 0, 56, 3, 1,
		0, 0, 0, 57, 58, 5, 10, 0, 0, 58, 59, 5, 12, 0, 0, 59, 60, 5, 5, 0, 0,
		60, 61, 5, 11, 0, 0, 61, 5, 1, 0, 0, 0, 62, 63, 5, 5, 0, 0, 63, 64, 5,
		13, 0, 0, 64, 68, 5, 14, 0, 0, 65, 67, 5, 4, 0, 0, 66, 65, 1, 0, 0, 0,
		67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 7, 1, 0,
		0, 0, 70, 68, 1, 0, 0, 0, 71, 72, 5, 6, 0, 0, 72, 76, 5, 8, 0, 0, 73, 74,
		5, 7, 0, 0, 74, 76, 5, 8, 0, 0, 75, 71, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0,
		76, 80, 1, 0, 0, 0, 77, 79, 5, 4, 0, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1,
		0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 86, 1, 0, 0, 0, 82,
		80, 1, 0, 0, 0, 83, 85, 5, 5, 0, 0, 84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0,
		0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 92, 1, 0, 0, 0, 88, 86,
		1, 0, 0, 0, 89, 91, 5, 15, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0,
		92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 98, 1, 0, 0, 0, 94, 92, 1,
		0, 0, 0, 95, 97, 5, 4, 0, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98,
		96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98, 1, 0,
		0, 0, 101, 102, 5, 9, 0, 0, 102, 9, 1, 0, 0, 0, 103, 104, 5, 3, 0, 0, 104,
		11, 1, 0, 0, 0, 105, 106, 7, 0, 0, 0, 106, 13, 1, 0, 0, 0, 13, 19, 21,
		29, 35, 44, 50, 55, 68, 75, 80, 86, 92, 98,
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
	XSQLParserRULE_start     = 1
	XSQLParserRULE_end       = 2
	XSQLParserRULE_attribute = 3
	XSQLParserRULE_expr      = 4
	XSQLParserRULE_reference = 5
	XSQLParserRULE_chardata  = 6
)

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStart_() []IStartContext
	Start_(i int) IStartContext
	AllEnd() []IEndContext
	End(i int) IEndContext
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

func (s *ContentContext) AllStart_() []IStartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStartContext); ok {
			len++
		}
	}

	tst := make([]IStartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStartContext); ok {
			tst[i] = t.(IStartContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Start_(i int) IStartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartContext); ok {
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

	return t.(IStartContext)
}

func (s *ContentContext) AllEnd() []IEndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndContext); ok {
			len++
		}
	}

	tst := make([]IEndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndContext); ok {
			tst[i] = t.(IEndContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) End(i int) IEndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndContext); ok {
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

	return t.(IEndContext)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65528) != 0 {
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(14)
				p.Start_()
			}

		case 2:
			{
				p.SetState(15)
				p.End()
			}

		case 3:
			{
				p.SetState(16)
				p.Expr()
			}

		case 4:
			{
				p.SetState(17)
				p.Reference()
			}

		case 5:
			{
				p.SetState(18)
				p.Chardata()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(23)
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

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
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
	SLASH() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) OPEN() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN, 0)
}

func (s *StartContext) NAME() antlr.TerminalNode {
	return s.GetToken(XSQLParserNAME, 0)
}

func (s *StartContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE, 0)
}

func (s *StartContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserWS)
}

func (s *StartContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserWS, i)
}

func (s *StartContext) AllAttribute() []IAttributeContext {
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

func (s *StartContext) Attribute(i int) IAttributeContext {
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

func (s *StartContext) SLASH() antlr.TerminalNode {
	return s.GetToken(XSQLParserSLASH, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, XSQLParserRULE_start)
	var _la int

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Match(XSQLParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(26)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(31)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(32)
				p.Attribute()
			}

			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(38)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(XSQLParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(41)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(47)
				p.Attribute()
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(53)
			p.Match(XSQLParserSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(XSQLParserCLOSE)
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

// IEndContext is an interface to support dynamic dispatch.
type IEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	NAME() antlr.TerminalNode
	CLOSE() antlr.TerminalNode

	// IsEndContext differentiates from other interfaces.
	IsEndContext()
}

type EndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndContext() *EndContext {
	var p = new(EndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_end
	return p
}

func InitEmptyEndContext(p *EndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_end
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }

func (s *EndContext) OPEN() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN, 0)
}

func (s *EndContext) SLASH() antlr.TerminalNode {
	return s.GetToken(XSQLParserSLASH, 0)
}

func (s *EndContext) NAME() antlr.TerminalNode {
	return s.GetToken(XSQLParserNAME, 0)
}

func (s *EndContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE, 0)
}

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) End() (localctx IEndContext) {
	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, XSQLParserRULE_end)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(XSQLParserOPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(XSQLParserSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(XSQLParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
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
	p.EnterRule(localctx, 6, XSQLParserRULE_attribute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(XSQLParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(XSQLParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(XSQLParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserWS {
		{
			p.SetState(65)
			p.Match(XSQLParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(70)
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

	// GetVal returns the val token.
	GetVal() antlr.Token

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// Getter signatures
	CLOSE_CURLY_BRAXE() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	DOLLAR() antlr.TerminalNode
	OPEN_CURLY_BRAXE() antlr.TerminalNode
	HASH() antlr.TerminalNode
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    antlr.Token
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

func (s *ExprContext) GetVal() antlr.Token { return s.val }

func (s *ExprContext) SetVal(v antlr.Token) { s.val = v }

func (s *ExprContext) CLOSE_CURLY_BRAXE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE_CURLY_BRAXE, 0)
}

func (s *ExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserWS)
}

func (s *ExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserWS, i)
}

func (s *ExprContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserTEXT)
}

func (s *ExprContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserTEXT, i)
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

func (s *ExprContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserNAME)
}

func (s *ExprContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserNAME, i)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, XSQLParserRULE_expr)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case XSQLParserDOLLAR:
		{
			p.SetState(71)
			p.Match(XSQLParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(XSQLParserOPEN_CURLY_BRAXE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case XSQLParserHASH:
		{
			p.SetState(73)
			p.Match(XSQLParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(77)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserNAME {
		{
			p.SetState(83)

			var _m = p.Match(XSQLParserNAME)

			localctx.(*ExprContext).val = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserTEXT {
		{
			p.SetState(89)
			p.Match(XSQLParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserWS {
		{
			p.SetState(95)
			p.Match(XSQLParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
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
	p.EnterRule(localctx, 10, XSQLParserRULE_reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
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
	WS() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	OPEN() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	DOLLAR() antlr.TerminalNode
	HASH() antlr.TerminalNode
	OPEN_CURLY_BRAXE() antlr.TerminalNode
	CLOSE_CURLY_BRAXE() antlr.TerminalNode
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

func (s *ChardataContext) WS() antlr.TerminalNode {
	return s.GetToken(XSQLParserWS, 0)
}

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
	p.EnterRule(localctx, 12, XSQLParserRULE_chardata)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
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
