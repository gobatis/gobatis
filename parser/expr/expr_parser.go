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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 199,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 2, 7,
	2, 47, 10, 2, 12, 2, 14, 2, 50, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 62, 10, 4, 3, 5, 3, 5, 5, 5, 66, 10, 5,
	3, 5, 3, 5, 3, 6, 7, 6, 71, 10, 6, 12, 6, 14, 6, 74, 11, 6, 3, 6, 3, 6,
	7, 6, 78, 10, 6, 12, 6, 14, 6, 81, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 92, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 109,
	10, 8, 12, 8, 14, 8, 112, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 122, 10, 9, 7, 9, 124, 10, 9, 12, 9, 14, 9, 127, 11, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 135, 10, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 150, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 7, 19, 165, 10, 19, 12, 19,
	14, 19, 168, 11, 19, 3, 20, 3, 20, 3, 20, 5, 20, 173, 10, 20, 5, 20, 175,
	10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 181, 10, 21, 3, 21, 3, 21, 5,
	21, 185, 10, 21, 3, 21, 5, 21, 188, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 5, 21, 195, 10, 21, 3, 21, 3, 21, 3, 21, 2, 4, 14, 16, 22, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 2,
	8, 3, 2, 31, 34, 4, 2, 26, 30, 35, 36, 4, 2, 25, 25, 32, 34, 3, 2, 19,
	24, 4, 2, 37, 40, 44, 45, 3, 2, 51, 52, 2, 208, 2, 42, 3, 2, 2, 2, 4, 53,
	3, 2, 2, 2, 6, 61, 3, 2, 2, 2, 8, 65, 3, 2, 2, 2, 10, 72, 3, 2, 2, 2, 12,
	84, 3, 2, 2, 2, 14, 91, 3, 2, 2, 2, 16, 113, 3, 2, 2, 2, 18, 134, 3, 2,
	2, 2, 20, 136, 3, 2, 2, 2, 22, 138, 3, 2, 2, 2, 24, 141, 3, 2, 2, 2, 26,
	149, 3, 2, 2, 2, 28, 151, 3, 2, 2, 2, 30, 153, 3, 2, 2, 2, 32, 155, 3,
	2, 2, 2, 34, 157, 3, 2, 2, 2, 36, 161, 3, 2, 2, 2, 38, 169, 3, 2, 2, 2,
	40, 178, 3, 2, 2, 2, 42, 48, 5, 6, 4, 2, 43, 44, 5, 4, 3, 2, 44, 45, 5,
	6, 4, 2, 45, 47, 3, 2, 2, 2, 46, 43, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48,
	46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 3, 2, 2, 2, 50, 48, 3, 2, 2,
	2, 51, 52, 7, 2, 2, 3, 52, 3, 3, 2, 2, 2, 53, 54, 7, 10, 2, 2, 54, 5, 3,
	2, 2, 2, 55, 62, 7, 4, 2, 2, 56, 57, 7, 12, 2, 2, 57, 62, 5, 8, 5, 2, 58,
	59, 7, 4, 2, 2, 59, 60, 7, 12, 2, 2, 60, 62, 5, 8, 5, 2, 61, 55, 3, 2,
	2, 2, 61, 56, 3, 2, 2, 2, 61, 58, 3, 2, 2, 2, 62, 7, 3, 2, 2, 2, 63, 64,
	7, 7, 2, 2, 64, 66, 7, 8, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2,
	66, 67, 3, 2, 2, 2, 67, 68, 7, 4, 2, 2, 68, 9, 3, 2, 2, 2, 69, 71, 5, 12,
	7, 2, 70, 69, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73,
	3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 79, 5, 14, 8, 2,
	76, 78, 5, 12, 7, 2, 77, 76, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3,
	2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82,
	83, 7, 2, 2, 3, 83, 11, 3, 2, 2, 2, 84, 85, 7, 53, 2, 2, 85, 86, 7, 54,
	2, 2, 86, 13, 3, 2, 2, 2, 87, 88, 8, 8, 1, 2, 88, 92, 5, 16, 9, 2, 89,
	90, 9, 2, 2, 2, 90, 92, 5, 14, 8, 8, 91, 87, 3, 2, 2, 2, 91, 89, 3, 2,
	2, 2, 92, 110, 3, 2, 2, 2, 93, 94, 12, 7, 2, 2, 94, 95, 9, 3, 2, 2, 95,
	109, 5, 14, 8, 8, 96, 97, 12, 6, 2, 2, 97, 98, 9, 4, 2, 2, 98, 109, 5,
	14, 8, 7, 99, 100, 12, 5, 2, 2, 100, 101, 9, 5, 2, 2, 101, 109, 5, 14,
	8, 6, 102, 103, 12, 4, 2, 2, 103, 104, 7, 18, 2, 2, 104, 109, 5, 14, 8,
	5, 105, 106, 12, 3, 2, 2, 106, 107, 7, 17, 2, 2, 107, 109, 5, 14, 8, 4,
	108, 93, 3, 2, 2, 2, 108, 96, 3, 2, 2, 2, 108, 99, 3, 2, 2, 2, 108, 102,
	3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2,
	2, 2, 110, 111, 3, 2, 2, 2, 111, 15, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2,
	113, 114, 8, 9, 1, 2, 114, 115, 5, 18, 10, 2, 115, 125, 3, 2, 2, 2, 116,
	121, 12, 3, 2, 2, 117, 122, 5, 22, 12, 2, 118, 122, 5, 34, 18, 2, 119,
	122, 5, 40, 21, 2, 120, 122, 5, 38, 20, 2, 121, 117, 3, 2, 2, 2, 121, 118,
	3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 124, 3, 2,
	2, 2, 123, 116, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2,
	125, 126, 3, 2, 2, 2, 126, 17, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128, 135,
	5, 24, 13, 2, 129, 135, 5, 20, 11, 2, 130, 131, 7, 5, 2, 2, 131, 132, 5,
	14, 8, 2, 132, 133, 7, 6, 2, 2, 133, 135, 3, 2, 2, 2, 134, 128, 3, 2, 2,
	2, 134, 129, 3, 2, 2, 2, 134, 130, 3, 2, 2, 2, 135, 19, 3, 2, 2, 2, 136,
	137, 7, 4, 2, 2, 137, 21, 3, 2, 2, 2, 138, 139, 7, 13, 2, 2, 139, 140,
	7, 4, 2, 2, 140, 23, 3, 2, 2, 2, 141, 142, 5, 26, 14, 2, 142, 25, 3, 2,
	2, 2, 143, 150, 7, 3, 2, 2, 144, 150, 5, 28, 15, 2, 145, 150, 5, 30, 16,
	2, 146, 150, 5, 32, 17, 2, 147, 150, 7, 44, 2, 2, 148, 150, 7, 45, 2, 2,
	149, 143, 3, 2, 2, 2, 149, 144, 3, 2, 2, 2, 149, 145, 3, 2, 2, 2, 149,
	146, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2, 2, 2, 150, 27, 3,
	2, 2, 2, 151, 152, 9, 6, 2, 2, 152, 29, 3, 2, 2, 2, 153, 154, 9, 7, 2,
	2, 154, 31, 3, 2, 2, 2, 155, 156, 7, 41, 2, 2, 156, 33, 3, 2, 2, 2, 157,
	158, 7, 7, 2, 2, 158, 159, 5, 14, 8, 2, 159, 160, 7, 8, 2, 2, 160, 35,
	3, 2, 2, 2, 161, 166, 5, 14, 8, 2, 162, 163, 7, 10, 2, 2, 163, 165, 5,
	14, 8, 2, 164, 162, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2,
	2, 166, 167, 3, 2, 2, 2, 167, 37, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169,
	174, 7, 5, 2, 2, 170, 172, 5, 36, 19, 2, 171, 173, 7, 16, 2, 2, 172, 171,
	3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 175, 3, 2, 2, 2, 174, 170, 3, 2,
	2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 177, 7, 6, 2, 2,
	177, 39, 3, 2, 2, 2, 178, 194, 7, 7, 2, 2, 179, 181, 5, 14, 8, 2, 180,
	179, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 184,
	7, 12, 2, 2, 183, 185, 5, 14, 8, 2, 184, 183, 3, 2, 2, 2, 184, 185, 3,
	2, 2, 2, 185, 195, 3, 2, 2, 2, 186, 188, 5, 14, 8, 2, 187, 186, 3, 2, 2,
	2, 187, 188, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 7, 12, 2, 2, 190,
	191, 5, 14, 8, 2, 191, 192, 7, 12, 2, 2, 192, 193, 5, 14, 8, 2, 193, 195,
	3, 2, 2, 2, 194, 180, 3, 2, 2, 2, 194, 187, 3, 2, 2, 2, 195, 196, 3, 2,
	2, 2, 196, 197, 7, 8, 2, 2, 197, 41, 3, 2, 2, 2, 21, 48, 61, 65, 72, 79,
	91, 108, 110, 121, 125, 134, 149, 166, 172, 174, 180, 184, 187, 194,
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
	"parameters", "paramComma", "paramDecl", "paramType", "expressions", "misc",
	"expression", "primaryExpr", "operand", "var_", "member", "literal", "basicLit",
	"integer", "string_", "float_", "index", "expressionList", "call", "slice_",
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
	ExprParserRULE_misc           = 5
	ExprParserRULE_expression     = 6
	ExprParserRULE_primaryExpr    = 7
	ExprParserRULE_operand        = 8
	ExprParserRULE_var_           = 9
	ExprParserRULE_member         = 10
	ExprParserRULE_literal        = 11
	ExprParserRULE_basicLit       = 12
	ExprParserRULE_integer        = 13
	ExprParserRULE_string_        = 14
	ExprParserRULE_float_         = 15
	ExprParserRULE_index          = 16
	ExprParserRULE_expressionList = 17
	ExprParserRULE_call           = 18
	ExprParserRULE_slice_         = 19
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
		p.SetState(40)
		p.ParamDecl()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(41)
				p.ParamComma()
			}
			{
				p.SetState(42)
				p.ParamDecl()
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(49)
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
		p.SetState(51)
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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(ExprParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(55)
			p.ParamType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Match(ExprParserIDENTIFIER)
		}
		{
			p.SetState(57)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(58)
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

func (s *ParamTypeContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(ExprParserL_BRACKET, 0)
}

func (s *ParamTypeContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(ExprParserR_BRACKET, 0)
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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(61)
			p.Match(ExprParserL_BRACKET)
		}
		{
			p.SetState(62)
			p.Match(ExprParserR_BRACKET)
		}

	}
	{
		p.SetState(65)
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

func (s *ExpressionsContext) AllMisc() []IMiscContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMiscContext)(nil)).Elem())
	var tst = make([]IMiscContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMiscContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Misc(i int) IMiscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMiscContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMiscContext)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(67)
				p.Misc()
			}

		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(73)
		p.expression(0)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(74)
				p.Misc()
			}

		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	{
		p.SetState(80)
		p.Match(ExprParserEOF)
	}

	return localctx
}

// IMiscContext is an interface to support dynamic dispatch.
type IMiscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMiscContext differentiates from other interfaces.
	IsMiscContext()
}

type MiscContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMiscContext() *MiscContext {
	var p = new(MiscContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_misc
	return p
}

func (*MiscContext) IsMiscContext() {}

func NewMiscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MiscContext {
	var p = new(MiscContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_misc

	return p
}

func (s *MiscContext) GetParser() antlr.Parser { return s.parser }

func (s *MiscContext) WS() antlr.TerminalNode {
	return s.GetToken(ExprParserWS, 0)
}

func (s *MiscContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(ExprParserTERMINATOR, 0)
}

func (s *MiscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MiscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MiscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterMisc(s)
	}
}

func (s *MiscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitMisc(s)
	}
}

func (p *ExprParser) Misc() (localctx IMiscContext) {
	localctx = NewMiscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExprParserRULE_misc)

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
		p.SetState(82)
		p.Match(ExprParserWS)
	}
	{
		p.SetState(83)
		p.Match(ExprParserTERMINATOR)
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
	_startState := 12
	p.EnterRecursionRule(localctx, 12, ExprParserRULE_expression, _p)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(86)
			p.primaryExpr(0)
		}

	case 2:
		{
			p.SetState(87)

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
			p.SetState(88)
			p.expression(6)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(92)

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
					p.SetState(93)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(94)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(95)

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
					p.SetState(96)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(98)

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
					p.SetState(99)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(101)
					p.Match(ExprParserLOGICAL_AND)
				}
				{
					p.SetState(102)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(104)
					p.Match(ExprParserLOGICAL_OR)
				}
				{
					p.SetState(105)
					p.expression(2)
				}

			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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

func (s *PrimaryExprContext) Call() ICallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallContext)
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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ExprParserRULE_primaryExpr, _p)

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
		p.SetState(112)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_primaryExpr)
			p.SetState(114)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(115)
					p.Member()
				}

			case 2:
				{
					p.SetState(116)
					p.Index()
				}

			case 3:
				{
					p.SetState(117)
					p.Slice_()
				}

			case 4:
				{
					p.SetState(118)
					p.Call()
				}

			}

		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 16, ExprParserRULE_operand)

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Var_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Match(ExprParserL_PAREN)
		}
		{
			p.SetState(129)
			p.expression(0)
		}
		{
			p.SetState(130)
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
	p.EnterRule(localctx, 18, ExprParserRULE_var_)

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
	p.EnterRule(localctx, 20, ExprParserRULE_member)

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
		p.SetState(136)
		p.Match(ExprParserDOT)
	}
	{
		p.SetState(137)
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
	p.EnterRule(localctx, 22, ExprParserRULE_literal)

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
		p.SetState(139)
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
	p.EnterRule(localctx, 24, ExprParserRULE_basicLit)

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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(ExprParserNIL_LIT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Integer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(143)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(144)
			p.Float_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(145)
			p.Match(ExprParserIMAGINARY_LIT)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(146)
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
	p.EnterRule(localctx, 26, ExprParserRULE_integer)
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
		p.SetState(149)
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
	p.EnterRule(localctx, 28, ExprParserRULE_string_)
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
		p.SetState(151)
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
	p.EnterRule(localctx, 30, ExprParserRULE_float_)

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
		p.SetState(153)
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
	p.EnterRule(localctx, 32, ExprParserRULE_index)

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
	{
		p.SetState(156)
		p.expression(0)
	}
	{
		p.SetState(157)
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
	p.EnterRule(localctx, 34, ExprParserRULE_expressionList)

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
		p.SetState(159)
		p.expression(0)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(160)
				p.Match(ExprParserCOMMA)
			}
			{
				p.SetState(161)
				p.expression(0)
			}

		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserL_PAREN, 0)
}

func (s *CallContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserR_PAREN, 0)
}

func (s *CallContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *CallContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(ExprParserELLIPSIS, 0)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *ExprParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ExprParserRULE_call)

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
		p.SetState(167)
		p.Match(ExprParserL_PAREN)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(168)
			p.ExpressionList()
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(169)
				p.Match(ExprParserELLIPSIS)
			}

		}

	}
	{
		p.SetState(174)
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
	p.EnterRule(localctx, 38, ExprParserRULE_slice_)

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
		p.SetState(176)
		p.Match(ExprParserL_BRACKET)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.SetState(178)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(177)
				p.expression(0)
			}

		}
		{
			p.SetState(180)
			p.Match(ExprParserCOLON)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(181)
				p.expression(0)
			}

		}

	case 2:
		p.SetState(185)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(184)
				p.expression(0)
			}

		}
		{
			p.SetState(187)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(188)
			p.expression(0)
		}
		{
			p.SetState(189)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(190)
			p.expression(0)
		}

	}
	{
		p.SetState(194)
		p.Match(ExprParserR_BRACKET)
	}

	return localctx
}

func (p *ExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 7:
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

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ExprParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
