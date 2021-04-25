// Code generated from ExprParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package expr // ExprParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 178,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 45, 10, 2,
	12, 2, 14, 2, 48, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 58, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 69, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 88, 10, 7, 12, 7,
	14, 7, 91, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	101, 10, 8, 7, 8, 103, 10, 8, 12, 8, 14, 8, 106, 11, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 5, 9, 114, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 129, 10,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 7, 18, 144, 10, 18, 12, 18, 14, 18, 147, 11, 18, 3,
	19, 3, 19, 3, 19, 5, 19, 152, 10, 19, 5, 19, 154, 10, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 5, 20, 160, 10, 20, 3, 20, 3, 20, 5, 20, 164, 10, 20, 3,
	20, 5, 20, 167, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 174,
	10, 20, 3, 20, 3, 20, 3, 20, 2, 4, 12, 14, 21, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 8, 3, 2, 31, 34, 4,
	2, 26, 30, 35, 36, 4, 2, 25, 25, 32, 34, 3, 2, 19, 24, 4, 2, 37, 40, 44,
	45, 3, 2, 51, 52, 2, 185, 2, 40, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2, 6, 57,
	3, 2, 2, 2, 8, 59, 3, 2, 2, 2, 10, 61, 3, 2, 2, 2, 12, 68, 3, 2, 2, 2,
	14, 92, 3, 2, 2, 2, 16, 113, 3, 2, 2, 2, 18, 115, 3, 2, 2, 2, 20, 117,
	3, 2, 2, 2, 22, 120, 3, 2, 2, 2, 24, 128, 3, 2, 2, 2, 26, 130, 3, 2, 2,
	2, 28, 132, 3, 2, 2, 2, 30, 134, 3, 2, 2, 2, 32, 136, 3, 2, 2, 2, 34, 140,
	3, 2, 2, 2, 36, 148, 3, 2, 2, 2, 38, 157, 3, 2, 2, 2, 40, 46, 5, 6, 4,
	2, 41, 42, 5, 4, 3, 2, 42, 43, 5, 6, 4, 2, 43, 45, 3, 2, 2, 2, 44, 41,
	3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2,
	47, 49, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 2, 2, 3, 50, 3, 3, 2,
	2, 2, 51, 52, 7, 10, 2, 2, 52, 5, 3, 2, 2, 2, 53, 58, 7, 4, 2, 2, 54, 55,
	7, 4, 2, 2, 55, 56, 7, 12, 2, 2, 56, 58, 5, 8, 5, 2, 57, 53, 3, 2, 2, 2,
	57, 54, 3, 2, 2, 2, 58, 7, 3, 2, 2, 2, 59, 60, 7, 4, 2, 2, 60, 9, 3, 2,
	2, 2, 61, 62, 5, 12, 7, 2, 62, 63, 7, 2, 2, 3, 63, 11, 3, 2, 2, 2, 64,
	65, 8, 7, 1, 2, 65, 69, 5, 14, 8, 2, 66, 67, 9, 2, 2, 2, 67, 69, 5, 12,
	7, 8, 68, 64, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 89, 3, 2, 2, 2, 70, 71,
	12, 7, 2, 2, 71, 72, 9, 3, 2, 2, 72, 88, 5, 12, 7, 8, 73, 74, 12, 6, 2,
	2, 74, 75, 9, 4, 2, 2, 75, 88, 5, 12, 7, 7, 76, 77, 12, 5, 2, 2, 77, 78,
	9, 5, 2, 2, 78, 88, 5, 12, 7, 6, 79, 80, 12, 4, 2, 2, 80, 81, 7, 18, 2,
	2, 81, 88, 5, 12, 7, 5, 82, 83, 12, 3, 2, 2, 83, 84, 7, 17, 2, 2, 84, 88,
	5, 12, 7, 4, 85, 86, 12, 9, 2, 2, 86, 88, 5, 2, 2, 2, 87, 70, 3, 2, 2,
	2, 87, 73, 3, 2, 2, 2, 87, 76, 3, 2, 2, 2, 87, 79, 3, 2, 2, 2, 87, 82,
	3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2,
	89, 90, 3, 2, 2, 2, 90, 13, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 8,
	8, 1, 2, 93, 94, 5, 16, 9, 2, 94, 104, 3, 2, 2, 2, 95, 100, 12, 3, 2, 2,
	96, 101, 5, 20, 11, 2, 97, 101, 5, 32, 17, 2, 98, 101, 5, 38, 20, 2, 99,
	101, 5, 36, 19, 2, 100, 96, 3, 2, 2, 2, 100, 97, 3, 2, 2, 2, 100, 98, 3,
	2, 2, 2, 100, 99, 3, 2, 2, 2, 101, 103, 3, 2, 2, 2, 102, 95, 3, 2, 2, 2,
	103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105,
	15, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 114, 5, 22, 12, 2, 108, 114,
	5, 18, 10, 2, 109, 110, 7, 5, 2, 2, 110, 111, 5, 12, 7, 2, 111, 112, 7,
	6, 2, 2, 112, 114, 3, 2, 2, 2, 113, 107, 3, 2, 2, 2, 113, 108, 3, 2, 2,
	2, 113, 109, 3, 2, 2, 2, 114, 17, 3, 2, 2, 2, 115, 116, 7, 4, 2, 2, 116,
	19, 3, 2, 2, 2, 117, 118, 7, 13, 2, 2, 118, 119, 7, 4, 2, 2, 119, 21, 3,
	2, 2, 2, 120, 121, 5, 24, 13, 2, 121, 23, 3, 2, 2, 2, 122, 129, 7, 3, 2,
	2, 123, 129, 5, 26, 14, 2, 124, 129, 5, 28, 15, 2, 125, 129, 5, 30, 16,
	2, 126, 129, 7, 44, 2, 2, 127, 129, 7, 45, 2, 2, 128, 122, 3, 2, 2, 2,
	128, 123, 3, 2, 2, 2, 128, 124, 3, 2, 2, 2, 128, 125, 3, 2, 2, 2, 128,
	126, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 25, 3, 2, 2, 2, 130, 131, 9,
	6, 2, 2, 131, 27, 3, 2, 2, 2, 132, 133, 9, 7, 2, 2, 133, 29, 3, 2, 2, 2,
	134, 135, 7, 41, 2, 2, 135, 31, 3, 2, 2, 2, 136, 137, 7, 7, 2, 2, 137,
	138, 5, 12, 7, 2, 138, 139, 7, 8, 2, 2, 139, 33, 3, 2, 2, 2, 140, 145,
	5, 12, 7, 2, 141, 142, 7, 10, 2, 2, 142, 144, 5, 12, 7, 2, 143, 141, 3,
	2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2,
	2, 146, 35, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 153, 7, 5, 2, 2, 149,
	151, 5, 34, 18, 2, 150, 152, 7, 16, 2, 2, 151, 150, 3, 2, 2, 2, 151, 152,
	3, 2, 2, 2, 152, 154, 3, 2, 2, 2, 153, 149, 3, 2, 2, 2, 153, 154, 3, 2,
	2, 2, 154, 155, 3, 2, 2, 2, 155, 156, 7, 6, 2, 2, 156, 37, 3, 2, 2, 2,
	157, 173, 7, 7, 2, 2, 158, 160, 5, 12, 7, 2, 159, 158, 3, 2, 2, 2, 159,
	160, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 163, 7, 12, 2, 2, 162, 164,
	5, 12, 7, 2, 163, 162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 174, 3, 2,
	2, 2, 165, 167, 5, 12, 7, 2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2,
	167, 168, 3, 2, 2, 2, 168, 169, 7, 12, 2, 2, 169, 170, 5, 12, 7, 2, 170,
	171, 7, 12, 2, 2, 171, 172, 5, 12, 7, 2, 172, 174, 3, 2, 2, 2, 173, 159,
	3, 2, 2, 2, 173, 166, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 7, 8,
	2, 2, 176, 39, 3, 2, 2, 2, 18, 46, 57, 68, 87, 89, 100, 104, 113, 128,
	145, 151, 153, 159, 163, 166, 173,
}
var literalNames = []string{
	"", "'nil'", "", "'('", "')'", "'['", "']'", "'='", "','", "';'", "':'",
	"'.'", "'++'", "'--'", "'...'", "'||'", "'&&'", "'=='", "'!='", "'<'",
	"'<='", "'>'", "'>='", "'|'", "'/'", "'%'", "'<<'", "'>>'", "'&^'", "'!'",
	"'+'", "'-'", "'^'", "'*'", "'&'",
}
var symbolicNames = []string{
	"", "NIL_LIT", "IDENTIFIER", "L_PAREN", "R_PAREN", "L_BRACKET", "R_BRACKET",
	"ASSIGN", "COMMA", "SEMI", "COLON", "DOT", "PLUS_PLUS", "MINUS_MINUS",
	"ELLIPSIS", "LOGICAL_OR", "LOGICAL_AND", "EQUALS", "NOT_EQUALS", "LESS",
	"LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "OR", "DIV", "MOD", "LSHIFT",
	"RSHIFT", "BIT_CLEAR", "EXCLAMATION", "PLUS", "MINUS", "CARET", "STAR",
	"AMPERSAND", "DECIMAL_LIT", "BINARY_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT",
	"DECIMAL_FLOAT_LIT", "HEX_FLOAT_LIT", "IMAGINARY_LIT", "RUNE_LIT", "BYTE_VALUE",
	"OCTAL_BYTE_VALUE", "HEX_BYTE_VALUE", "LITTLE_U_VALUE", "BIG_U_VALUE",
	"RAW_STRING_LIT", "INTERPRETED_STRING_LIT", "WS", "TERMINATOR",
}

var ruleNames = []string{
	"parameters", "paramComma", "paramDecl", "paramType", "expressions", "expression",
	"primaryExpr", "operand", "var_", "member", "literal", "basicLit", "integer",
	"string_", "float_", "index", "expressionList", "arguments", "slice_",
}

type ExprParser struct {
	*antlr.BaseParser
}

// NewExprParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ExprParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewExprParser(input antlr.TokenStream) *ExprParser {
	this := new(ExprParser)
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
	this.GrammarFileName = "ExprParser.g4"

	return this
}

// ExprParser tokens.
const (
	ExprParserEOF                    = antlr.TokenEOF
	ExprParserNIL_LIT                = 1
	ExprParserIDENTIFIER             = 2
	ExprParserL_PAREN                = 3
	ExprParserR_PAREN                = 4
	ExprParserL_BRACKET              = 5
	ExprParserR_BRACKET              = 6
	ExprParserASSIGN                 = 7
	ExprParserCOMMA                  = 8
	ExprParserSEMI                   = 9
	ExprParserCOLON                  = 10
	ExprParserDOT                    = 11
	ExprParserPLUS_PLUS              = 12
	ExprParserMINUS_MINUS            = 13
	ExprParserELLIPSIS               = 14
	ExprParserLOGICAL_OR             = 15
	ExprParserLOGICAL_AND            = 16
	ExprParserEQUALS                 = 17
	ExprParserNOT_EQUALS             = 18
	ExprParserLESS                   = 19
	ExprParserLESS_OR_EQUALS         = 20
	ExprParserGREATER                = 21
	ExprParserGREATER_OR_EQUALS      = 22
	ExprParserOR                     = 23
	ExprParserDIV                    = 24
	ExprParserMOD                    = 25
	ExprParserLSHIFT                 = 26
	ExprParserRSHIFT                 = 27
	ExprParserBIT_CLEAR              = 28
	ExprParserEXCLAMATION            = 29
	ExprParserPLUS                   = 30
	ExprParserMINUS                  = 31
	ExprParserCARET                  = 32
	ExprParserSTAR                   = 33
	ExprParserAMPERSAND              = 34
	ExprParserDECIMAL_LIT            = 35
	ExprParserBINARY_LIT             = 36
	ExprParserOCTAL_LIT              = 37
	ExprParserHEX_LIT                = 38
	ExprParserFLOAT_LIT              = 39
	ExprParserDECIMAL_FLOAT_LIT      = 40
	ExprParserHEX_FLOAT_LIT          = 41
	ExprParserIMAGINARY_LIT          = 42
	ExprParserRUNE_LIT               = 43
	ExprParserBYTE_VALUE             = 44
	ExprParserOCTAL_BYTE_VALUE       = 45
	ExprParserHEX_BYTE_VALUE         = 46
	ExprParserLITTLE_U_VALUE         = 47
	ExprParserBIG_U_VALUE            = 48
	ExprParserRAW_STRING_LIT         = 49
	ExprParserINTERPRETED_STRING_LIT = 50
	ExprParserWS                     = 51
	ExprParserTERMINATOR             = 52
)

// ExprParser rules.
const (
	ExprParserRULE_parameters     = 0
	ExprParserRULE_paramComma     = 1
	ExprParserRULE_paramDecl      = 2
	ExprParserRULE_paramType      = 3
	ExprParserRULE_expressions    = 4
	ExprParserRULE_expression     = 5
	ExprParserRULE_primaryExpr    = 6
	ExprParserRULE_operand        = 7
	ExprParserRULE_var_           = 8
	ExprParserRULE_member         = 9
	ExprParserRULE_literal        = 10
	ExprParserRULE_basicLit       = 11
	ExprParserRULE_integer        = 12
	ExprParserRULE_string_        = 13
	ExprParserRULE_float_         = 14
	ExprParserRULE_index          = 15
	ExprParserRULE_expressionList = 16
	ExprParserRULE_arguments      = 17
	ExprParserRULE_slice_         = 18
)

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllParamDecl() []IParamDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamDeclContext)(nil)).Elem())
	var tst = make([]IParamDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamDeclContext)
		}
	}

	return tst
}

func (s *ParametersContext) ParamDecl(i int) IParamDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *ParametersContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExprParserEOF, 0)
}

func (s *ParametersContext) AllParamComma() []IParamCommaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamCommaContext)(nil)).Elem())
	var tst = make([]IParamCommaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamCommaContext)
		}
	}

	return tst
}

func (s *ParametersContext) ParamComma(i int) IParamCommaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamCommaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamCommaContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *ExprParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExprParserRULE_parameters)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.ParamDecl()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(39)
				p.ParamComma()
			}
			{
				p.SetState(40)
				p.ParamDecl()
			}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(47)
		p.Match(ExprParserEOF)
	}

	return localctx
}

// IParamCommaContext is an interface to support dynamic dispatch.
type IParamCommaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamCommaContext differentiates from other interfaces.
	IsParamCommaContext()
}

type ParamCommaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamCommaContext() *ParamCommaContext {
	var p = new(ParamCommaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_paramComma
	return p
}

func (*ParamCommaContext) IsParamCommaContext() {}

func NewParamCommaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamCommaContext {
	var p = new(ParamCommaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_paramComma

	return p
}

func (s *ParamCommaContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamCommaContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, 0)
}

func (s *ParamCommaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamCommaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamCommaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterParamComma(s)
	}
}

func (s *ParamCommaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitParamComma(s)
	}
}

func (p *ExprParser) ParamComma() (localctx IParamCommaContext) {
	localctx = NewParamCommaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExprParserRULE_paramComma)

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
		p.SetState(49)
		p.Match(ExprParserCOMMA)
	}

	return localctx
}

// IParamDeclContext is an interface to support dynamic dispatch.
type IParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_paramDecl
	return p
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *ParamDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(ExprParserCOLON, 0)
}

func (s *ParamDeclContext) ParamType() IParamTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamTypeContext)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterParamDecl(s)
	}
}

func (s *ParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitParamDecl(s)
	}
}

func (p *ExprParser) ParamDecl() (localctx IParamDeclContext) {
	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExprParserRULE_paramDecl)

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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Match(ExprParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Match(ExprParserIDENTIFIER)
		}
		{
			p.SetState(53)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(54)
			p.ParamType()
		}

	}

	return localctx
}

// IParamTypeContext is an interface to support dynamic dispatch.
type IParamTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamTypeContext differentiates from other interfaces.
	IsParamTypeContext()
}

type ParamTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamTypeContext() *ParamTypeContext {
	var p = new(ParamTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_paramType
	return p
}

func (*ParamTypeContext) IsParamTypeContext() {}

func NewParamTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamTypeContext {
	var p = new(ParamTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_paramType

	return p
}

func (s *ParamTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *ParamTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterParamType(s)
	}
}

func (s *ParamTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitParamType(s)
	}
}

func (p *ExprParser) ParamType() (localctx IParamTypeContext) {
	localctx = NewParamTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExprParserRULE_paramType)

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
		p.SetState(57)
		p.Match(ExprParserIDENTIFIER)
	}

	return localctx
}

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionsContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExprParserEOF, 0)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *ExprParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExprParserRULE_expressions)

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
		p.SetState(59)
		p.expression(0)
	}
	{
		p.SetState(60)
		p.Match(ExprParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUnary_op returns the unary_op token.
	GetUnary_op() antlr.Token

	// GetMul_op returns the mul_op token.
	GetMul_op() antlr.Token

	// GetAdd_op returns the add_op token.
	GetAdd_op() antlr.Token

	// GetRel_op returns the rel_op token.
	GetRel_op() antlr.Token

	// SetUnary_op sets the unary_op token.
	SetUnary_op(antlr.Token)

	// SetMul_op sets the mul_op token.
	SetMul_op(antlr.Token)

	// SetAdd_op sets the add_op token.
	SetAdd_op(antlr.Token)

	// SetRel_op sets the rel_op token.
	SetRel_op(antlr.Token)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	unary_op antlr.Token
	mul_op   antlr.Token
	add_op   antlr.Token
	rel_op   antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetUnary_op() antlr.Token { return s.unary_op }

func (s *ExpressionContext) GetMul_op() antlr.Token { return s.mul_op }

func (s *ExpressionContext) GetAdd_op() antlr.Token { return s.add_op }

func (s *ExpressionContext) GetRel_op() antlr.Token { return s.rel_op }

func (s *ExpressionContext) SetUnary_op(v antlr.Token) { s.unary_op = v }

func (s *ExpressionContext) SetMul_op(v antlr.Token) { s.mul_op = v }

func (s *ExpressionContext) SetAdd_op(v antlr.Token) { s.add_op = v }

func (s *ExpressionContext) SetRel_op(v antlr.Token) { s.rel_op = v }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ExprParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ExprParserMINUS, 0)
}

func (s *ExpressionContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(ExprParserEXCLAMATION, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(ExprParserCARET, 0)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(ExprParserSTAR, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(ExprParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(ExprParserMOD, 0)
}

func (s *ExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(ExprParserLSHIFT, 0)
}

func (s *ExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(ExprParserRSHIFT, 0)
}

func (s *ExpressionContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(ExprParserAMPERSAND, 0)
}

func (s *ExpressionContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(ExprParserBIT_CLEAR, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(ExprParserOR, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ExprParserEQUALS, 0)
}

func (s *ExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(ExprParserNOT_EQUALS, 0)
}

func (s *ExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(ExprParserLESS, 0)
}

func (s *ExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(ExprParserLESS_OR_EQUALS, 0)
}

func (s *ExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(ExprParserGREATER, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(ExprParserGREATER_OR_EQUALS, 0)
}

func (s *ExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_AND, 0)
}

func (s *ExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_OR, 0)
}

func (s *ExpressionContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExprParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ExprParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ExprParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(63)
			p.primaryExpr(0)
		}

	case 2:
		{
			p.SetState(64)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).unary_op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(ExprParserEXCLAMATION-29))|(1<<(ExprParserPLUS-29))|(1<<(ExprParserMINUS-29))|(1<<(ExprParserCARET-29)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).unary_op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(65)
			p.expression(6)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(69)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).mul_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(ExprParserDIV-24))|(1<<(ExprParserMOD-24))|(1<<(ExprParserLSHIFT-24))|(1<<(ExprParserRSHIFT-24))|(1<<(ExprParserBIT_CLEAR-24))|(1<<(ExprParserSTAR-24))|(1<<(ExprParserAMPERSAND-24)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).mul_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(70)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(72)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).add_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-23)&-(0x1f+1)) == 0 && ((1<<uint((_la-23)))&((1<<(ExprParserOR-23))|(1<<(ExprParserPLUS-23))|(1<<(ExprParserMINUS-23))|(1<<(ExprParserCARET-23)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).add_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(73)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(75)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).rel_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExprParserEQUALS)|(1<<ExprParserNOT_EQUALS)|(1<<ExprParserLESS)|(1<<ExprParserLESS_OR_EQUALS)|(1<<ExprParserGREATER)|(1<<ExprParserGREATER_OR_EQUALS))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).rel_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(76)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(78)
					p.Match(ExprParserLOGICAL_AND)
				}
				{
					p.SetState(79)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(81)
					p.Match(ExprParserLOGICAL_OR)
				}
				{
					p.SetState(82)
					p.expression(2)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(84)
					p.Parameters()
				}

			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) Member() IMemberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *PrimaryExprContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExprContext) Slice_() ISlice_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISlice_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISlice_Context)
}

func (s *PrimaryExprContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (p *ExprParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *ExprParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, ExprParserRULE_primaryExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_primaryExpr)
			p.SetState(93)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(94)
					p.Member()
				}

			case 2:
				{
					p.SetState(95)
					p.Index()
				}

			case 3:
				{
					p.SetState(96)
					p.Slice_()
				}

			case 4:
				{
					p.SetState(97)
					p.Arguments()
				}

			}

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) Var_() IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *OperandContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserL_PAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserR_PAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *ExprParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ExprParserRULE_operand)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Var_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)
			p.Match(ExprParserL_PAREN)
		}
		{
			p.SetState(108)
			p.expression(0)
		}
		{
			p.SetState(109)
			p.Match(ExprParserR_PAREN)
		}

	}

	return localctx
}

// IVar_Context is an interface to support dynamic dispatch.
type IVar_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_Context differentiates from other interfaces.
	IsVar_Context()
}

type Var_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_Context() *Var_Context {
	var p = new(Var_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_var_
	return p
}

func (*Var_Context) IsVar_Context() {}

func NewVar_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_Context {
	var p = new(Var_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_var_

	return p
}

func (s *Var_Context) GetParser() antlr.Parser { return s.parser }

func (s *Var_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *Var_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterVar_(s)
	}
}

func (s *Var_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitVar_(s)
	}
}

func (p *ExprParser) Var_() (localctx IVar_Context) {
	localctx = NewVar_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExprParserRULE_var_)

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
		p.SetState(113)
		p.Match(ExprParserIDENTIFIER)
	}

	return localctx
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_member
	return p
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) DOT() antlr.TerminalNode {
	return s.GetToken(ExprParserDOT, 0)
}

func (s *MemberContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterMember(s)
	}
}

func (s *MemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitMember(s)
	}
}

func (p *ExprParser) Member() (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ExprParserRULE_member)

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
		p.SetState(115)
		p.Match(ExprParserDOT)
	}
	{
		p.SetState(116)
		p.Match(ExprParserIDENTIFIER)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *ExprParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ExprParserRULE_literal)

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
		p.SetState(118)
		p.BasicLit()
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserNIL_LIT, 0)
}

func (s *BasicLitContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *BasicLitContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *BasicLitContext) Float_() IFloat_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_Context)
}

func (s *BasicLitContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserIMAGINARY_LIT, 0)
}

func (s *BasicLitContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserRUNE_LIT, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterBasicLit(s)
	}
}

func (s *BasicLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitBasicLit(s)
	}
}

func (p *ExprParser) BasicLit() (localctx IBasicLitContext) {
	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ExprParserRULE_basicLit)

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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(ExprParserNIL_LIT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Integer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			p.Float_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(124)
			p.Match(ExprParserIMAGINARY_LIT)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(125)
			p.Match(ExprParserRUNE_LIT)
		}

	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserDECIMAL_LIT, 0)
}

func (s *IntegerContext) BINARY_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserBINARY_LIT, 0)
}

func (s *IntegerContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserOCTAL_LIT, 0)
}

func (s *IntegerContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserHEX_LIT, 0)
}

func (s *IntegerContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserIMAGINARY_LIT, 0)
}

func (s *IntegerContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserRUNE_LIT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *ExprParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ExprParserRULE_integer)
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
		p.SetState(128)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(ExprParserDECIMAL_LIT-35))|(1<<(ExprParserBINARY_LIT-35))|(1<<(ExprParserOCTAL_LIT-35))|(1<<(ExprParserHEX_LIT-35))|(1<<(ExprParserIMAGINARY_LIT-35))|(1<<(ExprParserRUNE_LIT-35)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserRAW_STRING_LIT, 0)
}

func (s *String_Context) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserINTERPRETED_STRING_LIT, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitString_(s)
	}
}

func (p *ExprParser) String_() (localctx IString_Context) {
	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ExprParserRULE_string_)
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
		p.SetState(130)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ExprParserRAW_STRING_LIT || _la == ExprParserINTERPRETED_STRING_LIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFloat_Context is an interface to support dynamic dispatch.
type IFloat_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloat_Context differentiates from other interfaces.
	IsFloat_Context()
}

type Float_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloat_Context() *Float_Context {
	var p = new(Float_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_float_
	return p
}

func (*Float_Context) IsFloat_Context() {}

func NewFloat_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_Context {
	var p = new(Float_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_float_

	return p
}

func (s *Float_Context) GetParser() antlr.Parser { return s.parser }

func (s *Float_Context) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserFLOAT_LIT, 0)
}

func (s *Float_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterFloat_(s)
	}
}

func (s *Float_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitFloat_(s)
	}
}

func (p *ExprParser) Float_() (localctx IFloat_Context) {
	localctx = NewFloat_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ExprParserRULE_float_)

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
		p.SetState(132)
		p.Match(ExprParserFLOAT_LIT)
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(ExprParserL_BRACKET, 0)
}

func (s *IndexContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(ExprParserR_BRACKET, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (p *ExprParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ExprParserRULE_index)

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
		p.SetState(134)
		p.Match(ExprParserL_BRACKET)
	}
	{
		p.SetState(135)
		p.expression(0)
	}
	{
		p.SetState(136)
		p.Match(ExprParserR_BRACKET)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExprParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *ExprParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ExprParserRULE_expressionList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.expression(0)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(139)
				p.Match(ExprParserCOMMA)
			}
			{
				p.SetState(140)
				p.expression(0)
			}

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserL_PAREN, 0)
}

func (s *ArgumentsContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserR_PAREN, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(ExprParserELLIPSIS, 0)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *ExprParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ExprParserRULE_arguments)

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
		p.SetState(146)
		p.Match(ExprParserL_PAREN)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(147)
			p.ExpressionList()
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(148)
				p.Match(ExprParserELLIPSIS)
			}

		}

	}
	{
		p.SetState(153)
		p.Match(ExprParserR_PAREN)
	}

	return localctx
}

// ISlice_Context is an interface to support dynamic dispatch.
type ISlice_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSlice_Context differentiates from other interfaces.
	IsSlice_Context()
}

type Slice_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlice_Context() *Slice_Context {
	var p = new(Slice_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_slice_
	return p
}

func (*Slice_Context) IsSlice_Context() {}

func NewSlice_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Slice_Context {
	var p = new(Slice_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_slice_

	return p
}

func (s *Slice_Context) GetParser() antlr.Parser { return s.parser }

func (s *Slice_Context) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(ExprParserL_BRACKET, 0)
}

func (s *Slice_Context) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(ExprParserR_BRACKET, 0)
}

func (s *Slice_Context) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ExprParserCOLON)
}

func (s *Slice_Context) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserCOLON, i)
}

func (s *Slice_Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Slice_Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Slice_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Slice_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Slice_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterSlice_(s)
	}
}

func (s *Slice_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitSlice_(s)
	}
}

func (p *ExprParser) Slice_() (localctx ISlice_Context) {
	localctx = NewSlice_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ExprParserRULE_slice_)

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
		p.SetState(155)
		p.Match(ExprParserL_BRACKET)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.SetState(157)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(156)
				p.expression(0)
			}

		}
		{
			p.SetState(159)
			p.Match(ExprParserCOLON)
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(160)
				p.expression(0)
			}

		}

	case 2:
		p.SetState(164)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(163)
				p.expression(0)
			}

		}
		{
			p.SetState(166)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(167)
			p.expression(0)
		}
		{
			p.SetState(168)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(169)
			p.expression(0)
		}

	}
	{
		p.SetState(173)
		p.Match(ExprParserR_BRACKET)
	}

	return localctx
}

func (p *ExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 6:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExprParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ExprParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
