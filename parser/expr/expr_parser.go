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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 213,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 61, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 70, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 87, 10, 6, 12, 6,
	14, 6, 90, 11, 6, 3, 7, 3, 7, 3, 7, 5, 7, 95, 10, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 100, 10, 7, 7, 7, 102, 10, 7, 12, 7, 14, 7, 105, 11, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 111, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 121, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 126, 10, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 136, 10, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 5, 16, 151, 10, 16, 3, 17, 3, 17, 5, 17, 155, 10, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 20, 7, 20, 169, 10, 20, 12, 20, 14, 20, 172, 11, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 5, 21, 179, 10, 21, 5, 21, 181, 10, 21, 3, 21, 5,
	21, 184, 10, 21, 3, 21, 5, 21, 187, 10, 21, 5, 21, 189, 10, 21, 3, 21,
	3, 21, 3, 22, 3, 22, 5, 22, 195, 10, 22, 3, 22, 3, 22, 5, 22, 199, 10,
	22, 3, 22, 5, 22, 202, 10, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22,
	209, 10, 22, 3, 22, 3, 22, 3, 22, 2, 4, 10, 12, 23, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 8, 3, 2,
	31, 36, 4, 2, 26, 30, 35, 36, 4, 2, 25, 25, 32, 34, 3, 2, 19, 24, 4, 2,
	37, 40, 44, 45, 3, 2, 51, 52, 2, 223, 2, 44, 3, 2, 2, 2, 4, 54, 3, 2, 2,
	2, 6, 60, 3, 2, 2, 2, 8, 62, 3, 2, 2, 2, 10, 69, 3, 2, 2, 2, 12, 94, 3,
	2, 2, 2, 14, 106, 3, 2, 2, 2, 16, 120, 3, 2, 2, 2, 18, 122, 3, 2, 2, 2,
	20, 127, 3, 2, 2, 2, 22, 135, 3, 2, 2, 2, 24, 137, 3, 2, 2, 2, 26, 139,
	3, 2, 2, 2, 28, 141, 3, 2, 2, 2, 30, 150, 3, 2, 2, 2, 32, 154, 3, 2, 2,
	2, 34, 156, 3, 2, 2, 2, 36, 160, 3, 2, 2, 2, 38, 165, 3, 2, 2, 2, 40, 173,
	3, 2, 2, 2, 42, 192, 3, 2, 2, 2, 44, 49, 5, 6, 4, 2, 45, 46, 7, 10, 2,
	2, 46, 48, 5, 6, 4, 2, 47, 45, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47,
	3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2,
	52, 53, 7, 2, 2, 3, 53, 3, 3, 2, 2, 2, 54, 55, 7, 4, 2, 2, 55, 5, 3, 2,
	2, 2, 56, 61, 7, 4, 2, 2, 57, 58, 7, 4, 2, 2, 58, 59, 7, 12, 2, 2, 59,
	61, 5, 4, 3, 2, 60, 56, 3, 2, 2, 2, 60, 57, 3, 2, 2, 2, 61, 7, 3, 2, 2,
	2, 62, 63, 5, 10, 6, 2, 63, 64, 7, 2, 2, 3, 64, 9, 3, 2, 2, 2, 65, 66,
	8, 6, 1, 2, 66, 70, 5, 12, 7, 2, 67, 68, 9, 2, 2, 2, 68, 70, 5, 10, 6,
	8, 69, 65, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 88, 3, 2, 2, 2, 71, 72,
	12, 7, 2, 2, 72, 73, 9, 3, 2, 2, 73, 87, 5, 10, 6, 8, 74, 75, 12, 6, 2,
	2, 75, 76, 9, 4, 2, 2, 76, 87, 5, 10, 6, 7, 77, 78, 12, 5, 2, 2, 78, 79,
	9, 5, 2, 2, 79, 87, 5, 10, 6, 6, 80, 81, 12, 4, 2, 2, 81, 82, 7, 18, 2,
	2, 82, 87, 5, 10, 6, 5, 83, 84, 12, 3, 2, 2, 84, 85, 7, 17, 2, 2, 85, 87,
	5, 10, 6, 4, 86, 71, 3, 2, 2, 2, 86, 74, 3, 2, 2, 2, 86, 77, 3, 2, 2, 2,
	86, 80, 3, 2, 2, 2, 86, 83, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3,
	2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 11, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91,
	92, 8, 7, 1, 2, 92, 95, 5, 16, 9, 2, 93, 95, 5, 14, 8, 2, 94, 91, 3, 2,
	2, 2, 94, 93, 3, 2, 2, 2, 95, 103, 3, 2, 2, 2, 96, 99, 12, 3, 2, 2, 97,
	100, 5, 28, 15, 2, 98, 100, 5, 42, 22, 2, 99, 97, 3, 2, 2, 2, 99, 98, 3,
	2, 2, 2, 100, 102, 3, 2, 2, 2, 101, 96, 3, 2, 2, 2, 102, 105, 3, 2, 2,
	2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 13, 3, 2, 2, 2, 105,
	103, 3, 2, 2, 2, 106, 107, 5, 30, 16, 2, 107, 108, 7, 5, 2, 2, 108, 110,
	5, 10, 6, 2, 109, 111, 7, 10, 2, 2, 110, 109, 3, 2, 2, 2, 110, 111, 3,
	2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 7, 6, 2, 2, 113, 15, 3, 2, 2,
	2, 114, 121, 5, 20, 11, 2, 115, 121, 5, 18, 10, 2, 116, 117, 7, 5, 2, 2,
	117, 118, 5, 10, 6, 2, 118, 119, 7, 6, 2, 2, 119, 121, 3, 2, 2, 2, 120,
	114, 3, 2, 2, 2, 120, 115, 3, 2, 2, 2, 120, 116, 3, 2, 2, 2, 121, 17, 3,
	2, 2, 2, 122, 125, 7, 4, 2, 2, 123, 124, 7, 13, 2, 2, 124, 126, 7, 4, 2,
	2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 19, 3, 2, 2, 2, 127,
	128, 5, 22, 12, 2, 128, 21, 3, 2, 2, 2, 129, 136, 7, 3, 2, 2, 130, 136,
	5, 24, 13, 2, 131, 136, 5, 26, 14, 2, 132, 136, 7, 41, 2, 2, 133, 136,
	7, 44, 2, 2, 134, 136, 7, 45, 2, 2, 135, 129, 3, 2, 2, 2, 135, 130, 3,
	2, 2, 2, 135, 131, 3, 2, 2, 2, 135, 132, 3, 2, 2, 2, 135, 133, 3, 2, 2,
	2, 135, 134, 3, 2, 2, 2, 136, 23, 3, 2, 2, 2, 137, 138, 9, 6, 2, 2, 138,
	25, 3, 2, 2, 2, 139, 140, 9, 7, 2, 2, 140, 27, 3, 2, 2, 2, 141, 142, 7,
	7, 2, 2, 142, 143, 5, 10, 6, 2, 143, 144, 7, 8, 2, 2, 144, 29, 3, 2, 2,
	2, 145, 151, 5, 32, 17, 2, 146, 147, 7, 5, 2, 2, 147, 148, 5, 30, 16, 2,
	148, 149, 7, 6, 2, 2, 149, 151, 3, 2, 2, 2, 150, 145, 3, 2, 2, 2, 150,
	146, 3, 2, 2, 2, 151, 31, 3, 2, 2, 2, 152, 155, 5, 34, 18, 2, 153, 155,
	7, 4, 2, 2, 154, 152, 3, 2, 2, 2, 154, 153, 3, 2, 2, 2, 155, 33, 3, 2,
	2, 2, 156, 157, 7, 4, 2, 2, 157, 158, 7, 13, 2, 2, 158, 159, 7, 4, 2, 2,
	159, 35, 3, 2, 2, 2, 160, 161, 7, 13, 2, 2, 161, 162, 7, 5, 2, 2, 162,
	163, 5, 30, 16, 2, 163, 164, 7, 6, 2, 2, 164, 37, 3, 2, 2, 2, 165, 170,
	5, 10, 6, 2, 166, 167, 7, 10, 2, 2, 167, 169, 5, 10, 6, 2, 168, 166, 3,
	2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2,
	2, 171, 39, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 188, 7, 5, 2, 2, 174,
	181, 5, 38, 20, 2, 175, 178, 5, 30, 16, 2, 176, 177, 7, 10, 2, 2, 177,
	179, 5, 38, 20, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181,
	3, 2, 2, 2, 180, 174, 3, 2, 2, 2, 180, 175, 3, 2, 2, 2, 181, 183, 3, 2,
	2, 2, 182, 184, 7, 16, 2, 2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2,
	184, 186, 3, 2, 2, 2, 185, 187, 7, 10, 2, 2, 186, 185, 3, 2, 2, 2, 186,
	187, 3, 2, 2, 2, 187, 189, 3, 2, 2, 2, 188, 180, 3, 2, 2, 2, 188, 189,
	3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 7, 6, 2, 2, 191, 41, 3, 2,
	2, 2, 192, 208, 7, 7, 2, 2, 193, 195, 5, 10, 6, 2, 194, 193, 3, 2, 2, 2,
	194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 198, 7, 12, 2, 2, 197,
	199, 5, 10, 6, 2, 198, 197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 209,
	3, 2, 2, 2, 200, 202, 5, 10, 6, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2,
	2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 7, 12, 2, 2, 204, 205, 5, 10, 6,
	2, 205, 206, 7, 12, 2, 2, 206, 207, 5, 10, 6, 2, 207, 209, 3, 2, 2, 2,
	208, 194, 3, 2, 2, 2, 208, 201, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210,
	211, 7, 8, 2, 2, 211, 43, 3, 2, 2, 2, 26, 49, 60, 69, 86, 88, 94, 99, 103,
	110, 120, 125, 135, 150, 154, 170, 178, 180, 183, 186, 188, 194, 198, 201,
	208,
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
	"parameters", "paramType", "paramDecl", "expressions", "expression", "primaryExpr",
	"conversion", "operand", "operandName", "literal", "basicLit", "integer",
	"string_", "index", "type_", "typeName", "qualifiedIdent", "typeAssertion",
	"expressionList", "arguments", "slice",
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
	ExprParserRULE_paramType      = 1
	ExprParserRULE_paramDecl      = 2
	ExprParserRULE_expressions    = 3
	ExprParserRULE_expression     = 4
	ExprParserRULE_primaryExpr    = 5
	ExprParserRULE_conversion     = 6
	ExprParserRULE_operand        = 7
	ExprParserRULE_operandName    = 8
	ExprParserRULE_literal        = 9
	ExprParserRULE_basicLit       = 10
	ExprParserRULE_integer        = 11
	ExprParserRULE_string_        = 12
	ExprParserRULE_index          = 13
	ExprParserRULE_type_          = 14
	ExprParserRULE_typeName       = 15
	ExprParserRULE_qualifiedIdent = 16
	ExprParserRULE_typeAssertion  = 17
	ExprParserRULE_expressionList = 18
	ExprParserRULE_arguments      = 19
	ExprParserRULE_slice          = 20
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

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExprParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, i)
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
		p.SetState(42)
		p.ParamDecl()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(43)
				p.Match(ExprParserCOMMA)
			}
			{
				p.SetState(44)
				p.ParamDecl()
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(50)
		p.Match(ExprParserEOF)
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
	p.EnterRule(localctx, 2, ExprParserRULE_paramType)

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
		p.SetState(52)
		p.Match(ExprParserIDENTIFIER)
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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Match(ExprParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Match(ExprParserIDENTIFIER)
		}
		{
			p.SetState(56)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(57)
			p.ParamType()
		}

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
	p.EnterRule(localctx, 6, ExprParserRULE_expressions)

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
		p.SetState(60)
		p.expression(0)
	}
	{
		p.SetState(61)
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

func (s *ExpressionContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(ExprParserAMPERSAND, 0)
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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, ExprParserRULE_expression, _p)
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(64)
			p.primaryExpr(0)
		}

	case 2:
		{
			p.SetState(65)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).unary_op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(ExprParserEXCLAMATION-29))|(1<<(ExprParserPLUS-29))|(1<<(ExprParserMINUS-29))|(1<<(ExprParserCARET-29))|(1<<(ExprParserSTAR-29))|(1<<(ExprParserAMPERSAND-29)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).unary_op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(66)
			p.expression(6)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(70)

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
					p.SetState(71)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(73)

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
					p.SetState(74)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(76)

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
					p.SetState(77)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(79)
					p.Match(ExprParserLOGICAL_AND)
				}
				{
					p.SetState(80)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(82)
					p.Match(ExprParserLOGICAL_OR)
				}
				{
					p.SetState(83)
					p.expression(2)
				}

			}

		}
		p.SetState(88)
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

func (s *PrimaryExprContext) Conversion() IConversionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConversionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConversionContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExprContext) Slice() ISliceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISliceContext)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ExprParserRULE_primaryExpr, _p)

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
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(90)
			p.Operand()
		}

	case 2:
		{
			p.SetState(91)
			p.Conversion()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_primaryExpr)
			p.SetState(94)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(95)
					p.Index()
				}

			case 2:
				{
					p.SetState(96)
					p.Slice()
				}

			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IConversionContext is an interface to support dynamic dispatch.
type IConversionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConversionContext differentiates from other interfaces.
	IsConversionContext()
}

type ConversionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConversionContext() *ConversionContext {
	var p = new(ConversionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_conversion
	return p
}

func (*ConversionContext) IsConversionContext() {}

func NewConversionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConversionContext {
	var p = new(ConversionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_conversion

	return p
}

func (s *ConversionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConversionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ConversionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserL_PAREN, 0)
}

func (s *ConversionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConversionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserR_PAREN, 0)
}

func (s *ConversionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, 0)
}

func (s *ConversionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConversionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConversionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterConversion(s)
	}
}

func (s *ConversionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitConversion(s)
	}
}

func (p *ExprParser) Conversion() (localctx IConversionContext) {
	localctx = NewConversionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExprParserRULE_conversion)

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
		p.SetState(104)
		p.Type_()
	}
	{
		p.SetState(105)
		p.Match(ExprParserL_PAREN)
	}
	{
		p.SetState(106)
		p.expression(0)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(107)
			p.Match(ExprParserCOMMA)
		}

	}
	{
		p.SetState(110)
		p.Match(ExprParserR_PAREN)
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

func (s *OperandContext) OperandName() IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.OperandName()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(114)
			p.Match(ExprParserL_PAREN)
		}
		{
			p.SetState(115)
			p.expression(0)
		}
		{
			p.SetState(116)
			p.Match(ExprParserR_PAREN)
		}

	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ExprParserIDENTIFIER)
}

func (s *OperandNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, i)
}

func (s *OperandNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(ExprParserDOT, 0)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterOperandName(s)
	}
}

func (s *OperandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitOperandName(s)
	}
}

func (p *ExprParser) OperandName() (localctx IOperandNameContext) {
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExprParserRULE_operandName)

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
		p.SetState(120)
		p.Match(ExprParserIDENTIFIER)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(121)
			p.Match(ExprParserDOT)
		}
		{
			p.SetState(122)
			p.Match(ExprParserIDENTIFIER)
		}

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
	p.EnterRule(localctx, 18, ExprParserRULE_literal)

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
		p.SetState(125)
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

func (s *BasicLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(ExprParserFLOAT_LIT, 0)
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
	p.EnterRule(localctx, 20, ExprParserRULE_basicLit)

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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(ExprParserNIL_LIT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Integer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.Match(ExprParserFLOAT_LIT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.Match(ExprParserIMAGINARY_LIT)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
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
	p.EnterRule(localctx, 22, ExprParserRULE_integer)
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
		p.SetState(135)
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
	p.EnterRule(localctx, 24, ExprParserRULE_string_)
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
		p.SetState(137)
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
	p.EnterRule(localctx, 26, ExprParserRULE_index)

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
		p.Match(ExprParserL_BRACKET)
	}
	{
		p.SetState(140)
		p.expression(0)
	}
	{
		p.SetState(141)
		p.Match(ExprParserR_BRACKET)
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *Type_Context) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserL_PAREN, 0)
}

func (s *Type_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Type_Context) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserR_PAREN, 0)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *ExprParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ExprParserRULE_type_)

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

	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.TypeName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Match(ExprParserL_PAREN)
		}
		{
			p.SetState(145)
			p.Type_()
		}
		{
			p.SetState(146)
			p.Match(ExprParserR_PAREN)
		}

	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) QualifiedIdent() IQualifiedIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentContext)
}

func (s *TypeNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *ExprParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ExprParserRULE_typeName)

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

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.QualifiedIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Match(ExprParserIDENTIFIER)
		}

	}

	return localctx
}

// IQualifiedIdentContext is an interface to support dynamic dispatch.
type IQualifiedIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedIdentContext differentiates from other interfaces.
	IsQualifiedIdentContext()
}

type QualifiedIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentContext() *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_qualifiedIdent
	return p
}

func (*QualifiedIdentContext) IsQualifiedIdentContext() {}

func NewQualifiedIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_qualifiedIdent

	return p
}

func (s *QualifiedIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ExprParserIDENTIFIER)
}

func (s *QualifiedIdentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, i)
}

func (s *QualifiedIdentContext) DOT() antlr.TerminalNode {
	return s.GetToken(ExprParserDOT, 0)
}

func (s *QualifiedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitQualifiedIdent(s)
	}
}

func (p *ExprParser) QualifiedIdent() (localctx IQualifiedIdentContext) {
	localctx = NewQualifiedIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ExprParserRULE_qualifiedIdent)

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
		p.SetState(154)
		p.Match(ExprParserIDENTIFIER)
	}
	{
		p.SetState(155)
		p.Match(ExprParserDOT)
	}
	{
		p.SetState(156)
		p.Match(ExprParserIDENTIFIER)
	}

	return localctx
}

// ITypeAssertionContext is an interface to support dynamic dispatch.
type ITypeAssertionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeAssertionContext differentiates from other interfaces.
	IsTypeAssertionContext()
}

type TypeAssertionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAssertionContext() *TypeAssertionContext {
	var p = new(TypeAssertionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_typeAssertion
	return p
}

func (*TypeAssertionContext) IsTypeAssertionContext() {}

func NewTypeAssertionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAssertionContext {
	var p = new(TypeAssertionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_typeAssertion

	return p
}

func (s *TypeAssertionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAssertionContext) DOT() antlr.TerminalNode {
	return s.GetToken(ExprParserDOT, 0)
}

func (s *TypeAssertionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserL_PAREN, 0)
}

func (s *TypeAssertionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TypeAssertionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserR_PAREN, 0)
}

func (s *TypeAssertionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAssertionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAssertionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterTypeAssertion(s)
	}
}

func (s *TypeAssertionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitTypeAssertion(s)
	}
}

func (p *ExprParser) TypeAssertion() (localctx ITypeAssertionContext) {
	localctx = NewTypeAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ExprParserRULE_typeAssertion)

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
		p.SetState(158)
		p.Match(ExprParserDOT)
	}
	{
		p.SetState(159)
		p.Match(ExprParserL_PAREN)
	}
	{
		p.SetState(160)
		p.Type_()
	}
	{
		p.SetState(161)
		p.Match(ExprParserR_PAREN)
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
	p.EnterRule(localctx, 36, ExprParserRULE_expressionList)

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
		p.SetState(163)
		p.expression(0)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(164)
				p.Match(ExprParserCOMMA)
			}
			{
				p.SetState(165)
				p.expression(0)
			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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

func (s *ArgumentsContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ArgumentsContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(ExprParserELLIPSIS, 0)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExprParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, i)
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
	p.EnterRule(localctx, 38, ExprParserRULE_arguments)

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
		p.Match(ExprParserL_PAREN)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(172)
				p.ExpressionList()
			}

		case 2:
			{
				p.SetState(173)
				p.Type_()
			}
			p.SetState(176)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(174)
					p.Match(ExprParserCOMMA)
				}
				{
					p.SetState(175)
					p.ExpressionList()
				}

			}

		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(180)
				p.Match(ExprParserELLIPSIS)
			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(183)
				p.Match(ExprParserCOMMA)
			}

		}

	}
	{
		p.SetState(188)
		p.Match(ExprParserR_PAREN)
	}

	return localctx
}

// ISliceContext is an interface to support dynamic dispatch.
type ISliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceContext differentiates from other interfaces.
	IsSliceContext()
}

type SliceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceContext() *SliceContext {
	var p = new(SliceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_slice
	return p
}

func (*SliceContext) IsSliceContext() {}

func NewSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceContext {
	var p = new(SliceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_slice

	return p
}

func (s *SliceContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(ExprParserL_BRACKET, 0)
}

func (s *SliceContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(ExprParserR_BRACKET, 0)
}

func (s *SliceContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ExprParserCOLON)
}

func (s *SliceContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserCOLON, i)
}

func (s *SliceContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SliceContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterSlice(s)
	}
}

func (s *SliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitSlice(s)
	}
}

func (p *ExprParser) Slice() (localctx ISliceContext) {
	localctx = NewSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ExprParserRULE_slice)

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
		p.SetState(190)
		p.Match(ExprParserL_BRACKET)
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.SetState(192)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(191)
				p.expression(0)
			}

		}
		{
			p.SetState(194)
			p.Match(ExprParserCOLON)
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(195)
				p.expression(0)
			}

		}

	case 2:
		p.SetState(199)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(198)
				p.expression(0)
			}

		}
		{
			p.SetState(201)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(202)
			p.expression(0)
		}
		{
			p.SetState(203)
			p.Match(ExprParserCOLON)
		}
		{
			p.SetState(204)
			p.expression(0)
		}

	}
	{
		p.SetState(208)
		p.Match(ExprParserR_BRACKET)
	}

	return localctx
}

func (p *ExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 5:
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
