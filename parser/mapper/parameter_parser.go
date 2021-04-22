// Code generated from ParameterParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package xml // ParameterParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 23, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 7, 2, 10, 10, 2, 12, 2, 14, 2, 13,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 21, 10, 3, 3, 3, 2, 2,
	4, 2, 4, 2, 2, 2, 22, 2, 6, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 11, 5, 4,
	3, 2, 7, 8, 7, 4, 2, 2, 8, 10, 5, 4, 3, 2, 9, 7, 3, 2, 2, 2, 10, 13, 3,
	2, 2, 2, 11, 9, 3, 2, 2, 2, 11, 12, 3, 2, 2, 2, 12, 14, 3, 2, 2, 2, 13,
	11, 3, 2, 2, 2, 14, 15, 7, 2, 2, 3, 15, 3, 3, 2, 2, 2, 16, 21, 7, 3, 2,
	2, 17, 18, 7, 3, 2, 2, 18, 19, 7, 5, 2, 2, 19, 21, 7, 3, 2, 2, 20, 16,
	3, 2, 2, 2, 20, 17, 3, 2, 2, 2, 21, 5, 3, 2, 2, 2, 4, 11, 20,
}
var literalNames = []string{
	"", "", "','", "':'",
}
var symbolicNames = []string{
	"", "IDENTIFIER", "COMMA", "COLON", "WS",
}

var ruleNames = []string{
	"expression", "varSpec",
}

type ParameterParser struct {
	*antlr.BaseParser
}

// NewParameterParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ParameterParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewParameterParser(input antlr.TokenStream) *ParameterParser {
	this := new(ParameterParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ParameterParser.g4"

	return this
}

// ParameterParser tokens.
const (
	ParameterParserEOF        = antlr.TokenEOF
	ParameterParserIDENTIFIER = 1
	ParameterParserCOMMA      = 2
	ParameterParserCOLON      = 3
	ParameterParserWS         = 4
)

// ParameterParser rules.
const (
	ParameterParserRULE_expression = 0
	ParameterParserRULE_varSpec    = 1
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParameterParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParameterParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllVarSpec() []IVarSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarSpecContext)(nil)).Elem())
	var tst = make([]IVarSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarSpecContext)
		}
	}

	return tst
}

func (s *ExpressionContext) VarSpec(i int) IVarSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarSpecContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParameterParserEOF, 0)
}

func (s *ExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParameterParserCOMMA)
}

func (s *ExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParameterParserCOMMA, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParameterParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParameterParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.VarSpec()
	}
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ParameterParserCOMMA {
		{
			p.SetState(5)
			p.Match(ParameterParserCOMMA)
		}
		{
			p.SetState(6)
			p.VarSpec()
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(12)
		p.Match(ParameterParserEOF)
	}

	return localctx
}

// IVarSpecContext is an interface to support dynamic dispatch.
type IVarSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarSpecContext differentiates from other interfaces.
	IsVarSpecContext()
}

type VarSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarSpecContext() *VarSpecContext {
	var p = new(VarSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParameterParserRULE_varSpec
	return p
}

func (*VarSpecContext) IsVarSpecContext() {}

func NewVarSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarSpecContext {
	var p = new(VarSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParameterParserRULE_varSpec

	return p
}

func (s *VarSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *VarSpecContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ParameterParserIDENTIFIER)
}

func (s *VarSpecContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ParameterParserIDENTIFIER, i)
}

func (s *VarSpecContext) COLON() antlr.TerminalNode {
	return s.GetToken(ParameterParserCOLON, 0)
}

func (s *VarSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParameterParser) VarSpec() (localctx IVarSpecContext) {
	localctx = NewVarSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParameterParserRULE_varSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.Match(ParameterParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Match(ParameterParserIDENTIFIER)
		}
		{
			p.SetState(16)
			p.Match(ParameterParserCOLON)
		}
		{
			p.SetState(17)
			p.Match(ParameterParserIDENTIFIER)
		}

	}

	return localctx
}
