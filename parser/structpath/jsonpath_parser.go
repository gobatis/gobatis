// Code generated from JsonPath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package structpath // JsonPath
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

type JsonPathParser struct {
	*antlr.BaseParser
}

var JsonPathParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonpathParserInit() {
	staticData := &JsonPathParserStaticData
	staticData.LiteralNames = []string{
		"", "'$.'", "'.'", "'[*]'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "INDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"jsonpath", "dotnotation", "dotnotation_expr", "identifierWithQualifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 27, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 15, 8, 1, 10, 1, 12, 1, 18, 9, 1, 1, 2, 1,
		2, 3, 2, 22, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6, 0, 0, 24,
		0, 8, 1, 0, 0, 0, 2, 10, 1, 0, 0, 0, 4, 21, 1, 0, 0, 0, 6, 23, 1, 0, 0,
		0, 8, 9, 3, 2, 1, 0, 9, 1, 1, 0, 0, 0, 10, 11, 5, 1, 0, 0, 11, 16, 3, 4,
		2, 0, 12, 13, 5, 2, 0, 0, 13, 15, 3, 4, 2, 0, 14, 12, 1, 0, 0, 0, 15, 18,
		1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 3, 1, 0, 0, 0,
		18, 16, 1, 0, 0, 0, 19, 22, 3, 6, 3, 0, 20, 22, 5, 4, 0, 0, 21, 19, 1,
		0, 0, 0, 21, 20, 1, 0, 0, 0, 22, 5, 1, 0, 0, 0, 23, 24, 5, 4, 0, 0, 24,
		25, 5, 3, 0, 0, 25, 7, 1, 0, 0, 0, 2, 16, 21,
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

// JsonPathParserInit initializes any static state used to implement JsonPathParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonPathParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonPathParserInit() {
	staticData := &JsonPathParserStaticData
	staticData.once.Do(jsonpathParserInit)
}

// NewJsonPathParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonPathParser(input antlr.TokenStream) *JsonPathParser {
	JsonPathParserInit()
	this := new(JsonPathParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JsonPathParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JsonPath.g4"

	return this
}

// JsonPathParser tokens.
const (
	JsonPathParserEOF         = antlr.TokenEOF
	JsonPathParserT__0        = 1
	JsonPathParserT__1        = 2
	JsonPathParserT__2        = 3
	JsonPathParserINDENTIFIER = 4
	JsonPathParserWS          = 5
)

// JsonPathParser rules.
const (
	JsonPathParserRULE_jsonpath                = 0
	JsonPathParserRULE_dotnotation             = 1
	JsonPathParserRULE_dotnotation_expr        = 2
	JsonPathParserRULE_identifierWithQualifier = 3
)

// IJsonpathContext is an interface to support dynamic dispatch.
type IJsonpathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dotnotation() IDotnotationContext

	// IsJsonpathContext differentiates from other interfaces.
	IsJsonpathContext()
}

type JsonpathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonpathContext() *JsonpathContext {
	var p = new(JsonpathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonPathParserRULE_jsonpath
	return p
}

func InitEmptyJsonpathContext(p *JsonpathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonPathParserRULE_jsonpath
}

func (*JsonpathContext) IsJsonpathContext() {}

func NewJsonpathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonpathContext {
	var p = new(JsonpathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_jsonpath

	return p
}

func (s *JsonpathContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonpathContext) Dotnotation() IDotnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDotnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDotnotationContext)
}

func (s *JsonpathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonpathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonpathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterJsonpath(s)
	}
}

func (s *JsonpathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitJsonpath(s)
	}
}

func (s *JsonpathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonPathVisitor:
		return t.VisitJsonpath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonPathParser) Jsonpath() (localctx IJsonpathContext) {
	localctx = NewJsonpathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonPathParserRULE_jsonpath)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.Dotnotation()
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

// IDotnotationContext is an interface to support dynamic dispatch.
type IDotnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDotnotation_expr() []IDotnotation_exprContext
	Dotnotation_expr(i int) IDotnotation_exprContext

	// IsDotnotationContext differentiates from other interfaces.
	IsDotnotationContext()
}

type DotnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotnotationContext() *DotnotationContext {
	var p = new(DotnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonPathParserRULE_dotnotation
	return p
}

func InitEmptyDotnotationContext(p *DotnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonPathParserRULE_dotnotation
}

func (*DotnotationContext) IsDotnotationContext() {}

func NewDotnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotnotationContext {
	var p = new(DotnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_dotnotation

	return p
}

func (s *DotnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *DotnotationContext) AllDotnotation_expr() []IDotnotation_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDotnotation_exprContext); ok {
			len++
		}
	}

	tst := make([]IDotnotation_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDotnotation_exprContext); ok {
			tst[i] = t.(IDotnotation_exprContext)
			i++
		}
	}

	return tst
}

func (s *DotnotationContext) Dotnotation_expr(i int) IDotnotation_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDotnotation_exprContext); ok {
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

	return t.(IDotnotation_exprContext)
}

func (s *DotnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DotnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterDotnotation(s)
	}
}

func (s *DotnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitDotnotation(s)
	}
}

func (s *DotnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonPathVisitor:
		return t.VisitDotnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonPathParser) Dotnotation() (localctx IDotnotationContext) {
	localctx = NewDotnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonPathParserRULE_dotnotation)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Match(JsonPathParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(11)
		p.Dotnotation_expr()
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(12)
				p.Match(JsonPathParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(13)
				p.Dotnotation_expr()
			}

		}
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
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

// IDotnotation_exprContext is an interface to support dynamic dispatch.
type IDotnotation_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierWithQualifier() IIdentifierWithQualifierContext
	INDENTIFIER() antlr.TerminalNode

	// IsDotnotation_exprContext differentiates from other interfaces.
	IsDotnotation_exprContext()
}

type Dotnotation_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotnotation_exprContext() *Dotnotation_exprContext {
	var p = new(Dotnotation_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonPathParserRULE_dotnotation_expr
	return p
}

func InitEmptyDotnotation_exprContext(p *Dotnotation_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonPathParserRULE_dotnotation_expr
}

func (*Dotnotation_exprContext) IsDotnotation_exprContext() {}

func NewDotnotation_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dotnotation_exprContext {
	var p = new(Dotnotation_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_dotnotation_expr

	return p
}

func (s *Dotnotation_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Dotnotation_exprContext) IdentifierWithQualifier() IIdentifierWithQualifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierWithQualifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierWithQualifierContext)
}

func (s *Dotnotation_exprContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JsonPathParserINDENTIFIER, 0)
}

func (s *Dotnotation_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dotnotation_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dotnotation_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterDotnotation_expr(s)
	}
}

func (s *Dotnotation_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitDotnotation_expr(s)
	}
}

func (s *Dotnotation_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonPathVisitor:
		return t.VisitDotnotation_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonPathParser) Dotnotation_expr() (localctx IDotnotation_exprContext) {
	localctx = NewDotnotation_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonPathParserRULE_dotnotation_expr)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.IdentifierWithQualifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.Match(JsonPathParserINDENTIFIER)
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

// IIdentifierWithQualifierContext is an interface to support dynamic dispatch.
type IIdentifierWithQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDENTIFIER() antlr.TerminalNode

	// IsIdentifierWithQualifierContext differentiates from other interfaces.
	IsIdentifierWithQualifierContext()
}

type IdentifierWithQualifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierWithQualifierContext() *IdentifierWithQualifierContext {
	var p = new(IdentifierWithQualifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonPathParserRULE_identifierWithQualifier
	return p
}

func InitEmptyIdentifierWithQualifierContext(p *IdentifierWithQualifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonPathParserRULE_identifierWithQualifier
}

func (*IdentifierWithQualifierContext) IsIdentifierWithQualifierContext() {}

func NewIdentifierWithQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierWithQualifierContext {
	var p = new(IdentifierWithQualifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_identifierWithQualifier

	return p
}

func (s *IdentifierWithQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierWithQualifierContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JsonPathParserINDENTIFIER, 0)
}

func (s *IdentifierWithQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierWithQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierWithQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterIdentifierWithQualifier(s)
	}
}

func (s *IdentifierWithQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitIdentifierWithQualifier(s)
	}
}

func (s *IdentifierWithQualifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonPathVisitor:
		return t.VisitIdentifierWithQualifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonPathParser) IdentifierWithQualifier() (localctx IIdentifierWithQualifierContext) {
	localctx = NewIdentifierWithQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonPathParserRULE_identifierWithQualifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Match(JsonPathParserINDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Match(JsonPathParserT__2)
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
