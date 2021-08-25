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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 58, 194,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 52, 10, 2, 12, 2, 14, 2, 55, 11, 2,
	5, 2, 57, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	67, 10, 4, 3, 5, 3, 5, 5, 5, 71, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 85, 10, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 100,
	10, 8, 12, 8, 14, 8, 103, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 113, 10, 9, 7, 9, 115, 10, 9, 12, 9, 14, 9, 118, 11, 9,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 128, 10,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 143, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	7, 21, 160, 10, 21, 12, 21, 14, 21, 163, 11, 21, 3, 22, 3, 22, 3, 22, 5,
	22, 168, 10, 22, 5, 22, 170, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 5, 23,
	176, 10, 23, 3, 23, 3, 23, 5, 23, 180, 10, 23, 3, 23, 5, 23, 183, 10, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 190, 10, 23, 3, 23, 3, 23, 3,
	23, 2, 4, 14, 16, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 2, 9, 3, 2, 35, 38, 4, 2, 30, 34, 39, 40,
	4, 2, 29, 29, 36, 38, 3, 2, 23, 28, 3, 2, 4, 9, 4, 2, 41, 44, 48, 49, 3,
	2, 55, 56, 2, 198, 2, 56, 3, 2, 2, 2, 4, 60, 3, 2, 2, 2, 6, 66, 3, 2, 2,
	2, 8, 70, 3, 2, 2, 2, 10, 74, 3, 2, 2, 2, 12, 77, 3, 2, 2, 2, 14, 84, 3,
	2, 2, 2, 16, 104, 3, 2, 2, 2, 18, 119, 3, 2, 2, 2, 20, 127, 3, 2, 2, 2,
	22, 129, 3, 2, 2, 2, 24, 131, 3, 2, 2, 2, 26, 134, 3, 2, 2, 2, 28, 142,
	3, 2, 2, 2, 30, 144, 3, 2, 2, 2, 32, 146, 3, 2, 2, 2, 34, 148, 3, 2, 2,
	2, 36, 150, 3, 2, 2, 2, 38, 152, 3, 2, 2, 2, 40, 156, 3, 2, 2, 2, 42, 164,
	3, 2, 2, 2, 44, 173, 3, 2, 2, 2, 46, 57, 7, 39, 2, 2, 47, 53, 5, 6, 4,
	2, 48, 49, 5, 4, 3, 2, 49, 50, 5, 6, 4, 2, 50, 52, 3, 2, 2, 2, 51, 48,
	3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2,
	54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 46, 3, 2, 2, 2, 56, 47, 3,
	2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 7, 2, 2, 3, 59, 3, 3, 2, 2, 2, 60,
	61, 7, 16, 2, 2, 61, 5, 3, 2, 2, 2, 62, 67, 7, 10, 2, 2, 63, 64, 7, 10,
	2, 2, 64, 65, 7, 18, 2, 2, 65, 67, 5, 8, 5, 2, 66, 62, 3, 2, 2, 2, 66,
	63, 3, 2, 2, 2, 67, 7, 3, 2, 2, 2, 68, 69, 7, 13, 2, 2, 69, 71, 7, 14,
	2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73,
	7, 10, 2, 2, 73, 9, 3, 2, 2, 2, 74, 75, 5, 14, 8, 2, 75, 76, 7, 2, 2, 3,
	76, 11, 3, 2, 2, 2, 77, 78, 7, 57, 2, 2, 78, 79, 7, 58, 2, 2, 79, 13, 3,
	2, 2, 2, 80, 81, 8, 8, 1, 2, 81, 85, 5, 16, 9, 2, 82, 83, 9, 2, 2, 2, 83,
	85, 5, 14, 8, 7, 84, 80, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 101, 3, 2,
	2, 2, 86, 87, 12, 6, 2, 2, 87, 88, 9, 3, 2, 2, 88, 100, 5, 14, 8, 7, 89,
	90, 12, 5, 2, 2, 90, 91, 9, 4, 2, 2, 91, 100, 5, 14, 8, 6, 92, 93, 12,
	4, 2, 2, 93, 94, 9, 5, 2, 2, 94, 100, 5, 14, 8, 5, 95, 96, 12, 3, 2, 2,
	96, 97, 5, 18, 10, 2, 97, 98, 5, 14, 8, 4, 98, 100, 3, 2, 2, 2, 99, 86,
	3, 2, 2, 2, 99, 89, 3, 2, 2, 2, 99, 92, 3, 2, 2, 2, 99, 95, 3, 2, 2, 2,
	100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 15,
	3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 105, 8, 9, 1, 2, 105, 106, 5, 20,
	11, 2, 106, 116, 3, 2, 2, 2, 107, 112, 12, 3, 2, 2, 108, 113, 5, 24, 13,
	2, 109, 113, 5, 38, 20, 2, 110, 113, 5, 44, 23, 2, 111, 113, 5, 42, 22,
	2, 112, 108, 3, 2, 2, 2, 112, 109, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112,
	111, 3, 2, 2, 2, 113, 115, 3, 2, 2, 2, 114, 107, 3, 2, 2, 2, 115, 118,
	3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 17, 3, 2,
	2, 2, 118, 116, 3, 2, 2, 2, 119, 120, 9, 6, 2, 2, 120, 19, 3, 2, 2, 2,
	121, 128, 5, 26, 14, 2, 122, 128, 5, 22, 12, 2, 123, 124, 7, 11, 2, 2,
	124, 125, 5, 14, 8, 2, 125, 126, 7, 12, 2, 2, 126, 128, 3, 2, 2, 2, 127,
	121, 3, 2, 2, 2, 127, 122, 3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 128, 21, 3,
	2, 2, 2, 129, 130, 7, 10, 2, 2, 130, 23, 3, 2, 2, 2, 131, 132, 7, 19, 2,
	2, 132, 133, 7, 10, 2, 2, 133, 25, 3, 2, 2, 2, 134, 135, 5, 28, 15, 2,
	135, 27, 3, 2, 2, 2, 136, 143, 5, 32, 17, 2, 137, 143, 5, 30, 16, 2, 138,
	143, 5, 34, 18, 2, 139, 143, 5, 36, 19, 2, 140, 143, 7, 48, 2, 2, 141,
	143, 7, 49, 2, 2, 142, 136, 3, 2, 2, 2, 142, 137, 3, 2, 2, 2, 142, 138,
	3, 2, 2, 2, 142, 139, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 141, 3, 2,
	2, 2, 143, 29, 3, 2, 2, 2, 144, 145, 9, 7, 2, 2, 145, 31, 3, 2, 2, 2, 146,
	147, 7, 3, 2, 2, 147, 33, 3, 2, 2, 2, 148, 149, 9, 8, 2, 2, 149, 35, 3,
	2, 2, 2, 150, 151, 7, 45, 2, 2, 151, 37, 3, 2, 2, 2, 152, 153, 7, 13, 2,
	2, 153, 154, 5, 14, 8, 2, 154, 155, 7, 14, 2, 2, 155, 39, 3, 2, 2, 2, 156,
	161, 5, 14, 8, 2, 157, 158, 7, 16, 2, 2, 158, 160, 5, 14, 8, 2, 159, 157,
	3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2,
	2, 2, 162, 41, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 164, 169, 7, 11, 2, 2,
	165, 167, 5, 40, 21, 2, 166, 168, 7, 22, 2, 2, 167, 166, 3, 2, 2, 2, 167,
	168, 3, 2, 2, 2, 168, 170, 3, 2, 2, 2, 169, 165, 3, 2, 2, 2, 169, 170,
	3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 172, 7, 12, 2, 2, 172, 43, 3, 2,
	2, 2, 173, 189, 7, 13, 2, 2, 174, 176, 5, 14, 8, 2, 175, 174, 3, 2, 2,
	2, 175, 176, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 179, 7, 18, 2, 2, 178,
	180, 5, 14, 8, 2, 179, 178, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 190,
	3, 2, 2, 2, 181, 183, 5, 14, 8, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2,
	2, 2, 183, 184, 3, 2, 2, 2, 184, 185, 7, 18, 2, 2, 185, 186, 5, 14, 8,
	2, 186, 187, 7, 18, 2, 2, 187, 188, 5, 14, 8, 2, 188, 190, 3, 2, 2, 2,
	189, 175, 3, 2, 2, 2, 189, 182, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	192, 7, 14, 2, 2, 192, 45, 3, 2, 2, 2, 20, 53, 56, 66, 70, 84, 99, 101,
	112, 116, 127, 142, 161, 167, 169, 175, 179, 182, 189,
}
var literalNames = []string{
	"", "'nil'", "'||'", "'&&'", "'and'", "'or'", "'AND'", "'OR'", "", "'('",
	"')'", "'['", "']'", "'='", "','", "';'", "':'", "'.'", "'++'", "'--'",
	"'...'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'/'", "'%'",
	"'<<'", "'>>'", "'&^'", "'!'", "'+'", "'-'", "'^'", "'*'", "'&'",
}
var symbolicNames = []string{
	"", "NIL_LIT", "LOGICAL_OR", "LOGICAL_AND", "LOGICAL_AND_LOWER", "LOGICAL_OR_LOWER",
	"LOGICAL_AND_UPPER", "LOGICAL_OR_UPPER", "IDENTIFIER", "L_PAREN", "R_PAREN",
	"L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "SEMI", "COLON", "DOT", "PLUS_PLUS",
	"MINUS_MINUS", "ELLIPSIS", "EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS",
	"GREATER", "GREATER_OR_EQUALS", "OR", "DIV", "MOD", "LSHIFT", "RSHIFT",
	"BIT_CLEAR", "EXCLAMATION", "PLUS", "MINUS", "CARET", "STAR", "AMPERSAND",
	"DECIMAL_LIT", "BINARY_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT",
	"HEX_FLOAT_LIT", "IMAGINARY_LIT", "RUNE_LIT", "BYTE_VALUE", "OCTAL_BYTE_VALUE",
	"HEX_BYTE_VALUE", "LITTLE_U_VALUE", "BIG_U_VALUE", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
	"WS", "TERMINATOR",
}

var ruleNames = []string{
	"parameters", "paramComma", "paramDecl", "paramType", "expressions", "misc",
	"expression", "primaryExpr", "logical", "operand", "var_", "member", "literal",
	"basicLit", "integer", "nil_", "string_", "float_", "index", "expressionList",
	"call", "slice_",
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
	ExprParserLOGICAL_OR             = 2
	ExprParserLOGICAL_AND            = 3
	ExprParserLOGICAL_AND_LOWER      = 4
	ExprParserLOGICAL_OR_LOWER       = 5
	ExprParserLOGICAL_AND_UPPER      = 6
	ExprParserLOGICAL_OR_UPPER       = 7
	ExprParserIDENTIFIER             = 8
	ExprParserL_PAREN                = 9
	ExprParserR_PAREN                = 10
	ExprParserL_BRACKET              = 11
	ExprParserR_BRACKET              = 12
	ExprParserASSIGN                 = 13
	ExprParserCOMMA                  = 14
	ExprParserSEMI                   = 15
	ExprParserCOLON                  = 16
	ExprParserDOT                    = 17
	ExprParserPLUS_PLUS              = 18
	ExprParserMINUS_MINUS            = 19
	ExprParserELLIPSIS               = 20
	ExprParserEQUALS                 = 21
	ExprParserNOT_EQUALS             = 22
	ExprParserLESS                   = 23
	ExprParserLESS_OR_EQUALS         = 24
	ExprParserGREATER                = 25
	ExprParserGREATER_OR_EQUALS      = 26
	ExprParserOR                     = 27
	ExprParserDIV                    = 28
	ExprParserMOD                    = 29
	ExprParserLSHIFT                 = 30
	ExprParserRSHIFT                 = 31
	ExprParserBIT_CLEAR              = 32
	ExprParserEXCLAMATION            = 33
	ExprParserPLUS                   = 34
	ExprParserMINUS                  = 35
	ExprParserCARET                  = 36
	ExprParserSTAR                   = 37
	ExprParserAMPERSAND              = 38
	ExprParserDECIMAL_LIT            = 39
	ExprParserBINARY_LIT             = 40
	ExprParserOCTAL_LIT              = 41
	ExprParserHEX_LIT                = 42
	ExprParserFLOAT_LIT              = 43
	ExprParserDECIMAL_FLOAT_LIT      = 44
	ExprParserHEX_FLOAT_LIT          = 45
	ExprParserIMAGINARY_LIT          = 46
	ExprParserRUNE_LIT               = 47
	ExprParserBYTE_VALUE             = 48
	ExprParserOCTAL_BYTE_VALUE       = 49
	ExprParserHEX_BYTE_VALUE         = 50
	ExprParserLITTLE_U_VALUE         = 51
	ExprParserBIG_U_VALUE            = 52
	ExprParserRAW_STRING_LIT         = 53
	ExprParserINTERPRETED_STRING_LIT = 54
	ExprParserWS                     = 55
	ExprParserTERMINATOR             = 56
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
	ExprParserRULE_logical        = 8
	ExprParserRULE_operand        = 9
	ExprParserRULE_var_           = 10
	ExprParserRULE_member         = 11
	ExprParserRULE_literal        = 12
	ExprParserRULE_basicLit       = 13
	ExprParserRULE_integer        = 14
	ExprParserRULE_nil_           = 15
	ExprParserRULE_string_        = 16
	ExprParserRULE_float_         = 17
	ExprParserRULE_index          = 18
	ExprParserRULE_expressionList = 19
	ExprParserRULE_call           = 20
	ExprParserRULE_slice_         = 21
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

func (s *ParametersContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExprParserEOF, 0)
}

func (s *ParametersContext) STAR() antlr.TerminalNode {
	return s.GetToken(ExprParserSTAR, 0)
}

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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(44)
			p.Match(ExprParserSTAR)
		}

	case 2:
		{
			p.SetState(45)
			p.ParamDecl()
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(46)
					p.ParamComma()
				}
				{
					p.SetState(47)
					p.ParamDecl()
				}

			}
			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(56)
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
		p.SetState(58)
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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(ExprParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(ExprParserIDENTIFIER)
		}
		{
			p.SetState(62)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(63)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(66)
			p.Match(ExprParserL_BRACKET)
		}
		{
			p.SetState(67)
			p.Match(ExprParserR_BRACKET)
		}

	}
	{
		p.SetState(70)
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
		p.SetState(72)
		p.expression(0)
	}
	{
		p.SetState(73)
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
		p.SetState(75)
		p.Match(ExprParserWS)
	}
	{
		p.SetState(76)
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

func (s *ExpressionContext) Logical() ILogicalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalContext)
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
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(79)
			p.primaryExpr(0)
		}

	case 2:
		{
			p.SetState(80)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).unary_op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ExprParserEXCLAMATION-33))|(1<<(ExprParserPLUS-33))|(1<<(ExprParserMINUS-33))|(1<<(ExprParserCARET-33)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).unary_op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(81)
			p.expression(5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(85)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).mul_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-28)&-(0x1f+1)) == 0 && ((1<<uint((_la-28)))&((1<<(ExprParserDIV-28))|(1<<(ExprParserMOD-28))|(1<<(ExprParserLSHIFT-28))|(1<<(ExprParserRSHIFT-28))|(1<<(ExprParserBIT_CLEAR-28))|(1<<(ExprParserSTAR-28))|(1<<(ExprParserAMPERSAND-28)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).mul_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(86)
					p.expression(5)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(88)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).add_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(ExprParserOR-27))|(1<<(ExprParserPLUS-27))|(1<<(ExprParserMINUS-27))|(1<<(ExprParserCARET-27)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).add_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)
					p.expression(4)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(91)

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
					p.SetState(92)
					p.expression(3)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(94)
					p.Logical()
				}
				{
					p.SetState(95)
					p.expression(2)
				}

			}

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
		p.SetState(103)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_primaryExpr)
			p.SetState(105)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(106)
					p.Member()
				}

			case 2:
				{
					p.SetState(107)
					p.Index()
				}

			case 3:
				{
					p.SetState(108)
					p.Slice_()
				}

			case 4:
				{
					p.SetState(109)
					p.Call()
				}

			}

		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// ILogicalContext is an interface to support dynamic dispatch.
type ILogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalContext differentiates from other interfaces.
	IsLogicalContext()
}

type LogicalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalContext() *LogicalContext {
	var p = new(LogicalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_logical
	return p
}

func (*LogicalContext) IsLogicalContext() {}

func NewLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalContext {
	var p = new(LogicalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_logical

	return p
}

func (s *LogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_AND, 0)
}

func (s *LogicalContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_OR, 0)
}

func (s *LogicalContext) LOGICAL_AND_LOWER() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_AND_LOWER, 0)
}

func (s *LogicalContext) LOGICAL_OR_LOWER() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_OR_LOWER, 0)
}

func (s *LogicalContext) LOGICAL_AND_UPPER() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_AND_UPPER, 0)
}

func (s *LogicalContext) LOGICAL_OR_UPPER() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_OR_UPPER, 0)
}

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterLogical(s)
	}
}

func (s *LogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitLogical(s)
	}
}

func (p *ExprParser) Logical() (localctx ILogicalContext) {
	localctx = NewLogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExprParserRULE_logical)
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
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExprParserLOGICAL_OR)|(1<<ExprParserLOGICAL_AND)|(1<<ExprParserLOGICAL_AND_LOWER)|(1<<ExprParserLOGICAL_OR_LOWER)|(1<<ExprParserLOGICAL_AND_UPPER)|(1<<ExprParserLOGICAL_OR_UPPER))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 18, ExprParserRULE_operand)

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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Var_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.Match(ExprParserL_PAREN)
		}
		{
			p.SetState(122)
			p.expression(0)
		}
		{
			p.SetState(123)
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
	p.EnterRule(localctx, 20, ExprParserRULE_var_)

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
		p.SetState(127)
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
	p.EnterRule(localctx, 22, ExprParserRULE_member)

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
		p.SetState(129)
		p.Match(ExprParserDOT)
	}
	{
		p.SetState(130)
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
	p.EnterRule(localctx, 24, ExprParserRULE_literal)

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

func (s *BasicLitContext) Nil_() INil_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INil_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INil_Context)
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
	p.EnterRule(localctx, 26, ExprParserRULE_basicLit)

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

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Nil_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Integer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.Float_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			p.Match(ExprParserIMAGINARY_LIT)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(139)
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
	p.EnterRule(localctx, 28, ExprParserRULE_integer)
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
		p.SetState(142)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(ExprParserDECIMAL_LIT-39))|(1<<(ExprParserBINARY_LIT-39))|(1<<(ExprParserOCTAL_LIT-39))|(1<<(ExprParserHEX_LIT-39))|(1<<(ExprParserIMAGINARY_LIT-39))|(1<<(ExprParserRUNE_LIT-39)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INil_Context is an interface to support dynamic dispatch.
type INil_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNil_Context differentiates from other interfaces.
	IsNil_Context()
}

type Nil_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNil_Context() *Nil_Context {
	var p = new(Nil_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_nil_
	return p
}

func (*Nil_Context) IsNil_Context() {}

func NewNil_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nil_Context {
	var p = new(Nil_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_nil_

	return p
}

func (s *Nil_Context) GetParser() antlr.Parser { return s.parser }

func (s *Nil_Context) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserNIL_LIT, 0)
}

func (s *Nil_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nil_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nil_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterNil_(s)
	}
}

func (s *Nil_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitNil_(s)
	}
}

func (p *ExprParser) Nil_() (localctx INil_Context) {
	localctx = NewNil_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ExprParserRULE_nil_)

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
		p.SetState(144)
		p.Match(ExprParserNIL_LIT)
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
	p.EnterRule(localctx, 32, ExprParserRULE_string_)
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
		p.SetState(146)
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
	p.EnterRule(localctx, 34, ExprParserRULE_float_)

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
		p.SetState(148)
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
	p.EnterRule(localctx, 36, ExprParserRULE_index)

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
		p.SetState(150)
		p.Match(ExprParserL_BRACKET)
	}
	{
		p.SetState(151)
		p.expression(0)
	}
	{
		p.SetState(152)
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
	p.EnterRule(localctx, 38, ExprParserRULE_expressionList)

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
		p.SetState(154)
		p.expression(0)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(155)
				p.Match(ExprParserCOMMA)
			}
			{
				p.SetState(156)
				p.expression(0)
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 40, ExprParserRULE_call)

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
		p.SetState(162)
		p.Match(ExprParserL_PAREN)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(163)
			p.ExpressionList()
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(164)
				p.Match(ExprParserELLIPSIS)
			}

		}

	}
	{
		p.SetState(169)
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
	p.EnterRule(localctx, 42, ExprParserRULE_slice_)

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
		p.SetState(171)
		p.Match(ExprParserL_BRACKET)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.SetState(173)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(172)
				p.expression(0)
			}

		}
		{
			p.SetState(175)
			p.Match(ExprParserCOLON)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(176)
				p.expression(0)
			}

		}

	case 2:
		p.SetState(180)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(179)
				p.expression(0)
			}

		}
		{
			p.SetState(182)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(183)
			p.expression(0)
		}
		{
			p.SetState(184)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(185)
			p.expression(0)
		}

	}
	{
		p.SetState(189)
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
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ExprParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
