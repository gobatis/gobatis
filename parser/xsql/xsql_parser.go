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
		"", "'if'", "'</if>'", "'choose'", "'</choose>'", "'when'", "'</when>'",
		"'otherwise'", "'</otherwise>'", "'trim'", "'</trim'", "'where'", "'</where>'",
		"'set'", "'</set>'", "'foreach'", "'</foreach>'", "", "", "", "", "'$'",
		"'#'", "'{'", "'}'", "'<'", "'>'", "'/'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"BLOCK_COMMENT", "EntityRef", "WS", "NAME", "DOLLAR", "HASH", "OPEN_CURLY_BRAXE",
		"CLOSE_CURLY_BRAXE", "OPEN", "CLOSE", "SLASH", "EQUALS", "STRING", "TEXT",
	}
	staticData.RuleNames = []string{
		"document", "content", "element", "attribute", "expr", "reference",
		"chardata",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 218, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 22, 8, 1, 10, 1, 12, 1, 25, 9, 1, 1, 2, 1, 2, 1, 2, 5, 2, 30, 8, 2,
		10, 2, 12, 2, 33, 9, 2, 1, 2, 5, 2, 36, 8, 2, 10, 2, 12, 2, 39, 9, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 48, 8, 2, 10, 2, 12, 2, 51,
		9, 2, 1, 2, 5, 2, 54, 8, 2, 10, 2, 12, 2, 57, 9, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 66, 8, 2, 10, 2, 12, 2, 69, 9, 2, 1, 2, 5, 2,
		72, 8, 2, 10, 2, 12, 2, 75, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 5, 2, 84, 8, 2, 10, 2, 12, 2, 87, 9, 2, 1, 2, 5, 2, 90, 8, 2, 10, 2,
		12, 2, 93, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 102, 8,
		2, 10, 2, 12, 2, 105, 9, 2, 1, 2, 5, 2, 108, 8, 2, 10, 2, 12, 2, 111, 9,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 120, 8, 2, 10, 2, 12,
		2, 123, 9, 2, 1, 2, 5, 2, 126, 8, 2, 10, 2, 12, 2, 129, 9, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 138, 8, 2, 10, 2, 12, 2, 141, 9,
		2, 1, 2, 5, 2, 144, 8, 2, 10, 2, 12, 2, 147, 9, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 156, 8, 2, 10, 2, 12, 2, 159, 9, 2, 1, 2, 5,
		2, 162, 8, 2, 10, 2, 12, 2, 165, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 171,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 177, 8, 3, 10, 3, 12, 3, 180, 9, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 186, 8, 4, 1, 4, 5, 4, 189, 8, 4, 10, 4,
		12, 4, 192, 9, 4, 1, 4, 5, 4, 195, 8, 4, 10, 4, 12, 4, 198, 9, 4, 1, 4,
		5, 4, 201, 8, 4, 10, 4, 12, 4, 204, 9, 4, 1, 4, 5, 4, 207, 8, 4, 10, 4,
		12, 4, 210, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 0, 0, 7, 0,
		2, 4, 6, 8, 10, 12, 0, 1, 1, 0, 19, 30, 243, 0, 14, 1, 0, 0, 0, 2, 23,
		1, 0, 0, 0, 4, 170, 1, 0, 0, 0, 6, 172, 1, 0, 0, 0, 8, 185, 1, 0, 0, 0,
		10, 213, 1, 0, 0, 0, 12, 215, 1, 0, 0, 0, 14, 15, 3, 2, 1, 0, 15, 16, 5,
		0, 0, 1, 16, 1, 1, 0, 0, 0, 17, 22, 3, 4, 2, 0, 18, 22, 3, 8, 4, 0, 19,
		22, 3, 10, 5, 0, 20, 22, 3, 12, 6, 0, 21, 17, 1, 0, 0, 0, 21, 18, 1, 0,
		0, 0, 21, 19, 1, 0, 0, 0, 21, 20, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21,
		1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 3, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0,
		26, 27, 5, 25, 0, 0, 27, 31, 5, 1, 0, 0, 28, 30, 5, 19, 0, 0, 29, 28, 1,
		0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32,
		37, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 36, 3, 6, 3, 0, 35, 34, 1, 0, 0,
		0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40,
		1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 26, 0, 0, 41, 42, 3, 2, 1, 0,
		42, 43, 5, 2, 0, 0, 43, 171, 1, 0, 0, 0, 44, 45, 5, 25, 0, 0, 45, 49, 5,
		3, 0, 0, 46, 48, 5, 19, 0, 0, 47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49,
		47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 55, 1, 0, 0, 0, 51, 49, 1, 0, 0,
		0, 52, 54, 3, 6, 3, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53,
		1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 58, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0,
		58, 59, 5, 26, 0, 0, 59, 60, 3, 2, 1, 0, 60, 61, 5, 4, 0, 0, 61, 171, 1,
		0, 0, 0, 62, 63, 5, 25, 0, 0, 63, 67, 5, 5, 0, 0, 64, 66, 5, 19, 0, 0,
		65, 64, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1,
		0, 0, 0, 68, 73, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 72, 3, 6, 3, 0, 71,
		70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0,
		0, 74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 26, 0, 0, 77, 78,
		3, 2, 1, 0, 78, 79, 5, 6, 0, 0, 79, 171, 1, 0, 0, 0, 80, 81, 5, 25, 0,
		0, 81, 85, 5, 7, 0, 0, 82, 84, 5, 19, 0, 0, 83, 82, 1, 0, 0, 0, 84, 87,
		1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 91, 1, 0, 0, 0,
		87, 85, 1, 0, 0, 0, 88, 90, 3, 6, 3, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1,
		0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93,
		91, 1, 0, 0, 0, 94, 95, 5, 26, 0, 0, 95, 96, 3, 2, 1, 0, 96, 97, 5, 8,
		0, 0, 97, 171, 1, 0, 0, 0, 98, 99, 5, 25, 0, 0, 99, 103, 5, 9, 0, 0, 100,
		102, 5, 19, 0, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 109, 1, 0, 0, 0, 105, 103, 1, 0,
		0, 0, 106, 108, 3, 6, 3, 0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0,
		109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111,
		109, 1, 0, 0, 0, 112, 113, 5, 26, 0, 0, 113, 114, 3, 2, 1, 0, 114, 115,
		5, 10, 0, 0, 115, 171, 1, 0, 0, 0, 116, 117, 5, 25, 0, 0, 117, 121, 5,
		11, 0, 0, 118, 120, 5, 19, 0, 0, 119, 118, 1, 0, 0, 0, 120, 123, 1, 0,
		0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 127, 1, 0, 0, 0,
		123, 121, 1, 0, 0, 0, 124, 126, 3, 6, 3, 0, 125, 124, 1, 0, 0, 0, 126,
		129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 130,
		1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 131, 5, 26, 0, 0, 131, 132, 3, 2,
		1, 0, 132, 133, 5, 12, 0, 0, 133, 171, 1, 0, 0, 0, 134, 135, 5, 25, 0,
		0, 135, 139, 5, 13, 0, 0, 136, 138, 5, 19, 0, 0, 137, 136, 1, 0, 0, 0,
		138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		145, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 144, 3, 6, 3, 0, 143, 142,
		1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0,
		0, 0, 146, 148, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 26, 0, 0,
		149, 150, 3, 2, 1, 0, 150, 151, 5, 14, 0, 0, 151, 171, 1, 0, 0, 0, 152,
		153, 5, 25, 0, 0, 153, 157, 5, 15, 0, 0, 154, 156, 5, 19, 0, 0, 155, 154,
		1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0,
		0, 0, 158, 163, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 162, 3, 6, 3, 0,
		161, 160, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163,
		164, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167,
		5, 26, 0, 0, 167, 168, 3, 2, 1, 0, 168, 169, 5, 16, 0, 0, 169, 171, 1,
		0, 0, 0, 170, 26, 1, 0, 0, 0, 170, 44, 1, 0, 0, 0, 170, 62, 1, 0, 0, 0,
		170, 80, 1, 0, 0, 0, 170, 98, 1, 0, 0, 0, 170, 116, 1, 0, 0, 0, 170, 134,
		1, 0, 0, 0, 170, 152, 1, 0, 0, 0, 171, 5, 1, 0, 0, 0, 172, 173, 5, 20,
		0, 0, 173, 174, 5, 28, 0, 0, 174, 178, 5, 29, 0, 0, 175, 177, 5, 19, 0,
		0, 176, 175, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178,
		179, 1, 0, 0, 0, 179, 7, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 181, 182, 5,
		21, 0, 0, 182, 186, 5, 23, 0, 0, 183, 184, 5, 22, 0, 0, 184, 186, 5, 23,
		0, 0, 185, 181, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 190, 1, 0, 0, 0,
		187, 189, 5, 19, 0, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190,
		188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 196, 1, 0, 0, 0, 192, 190,
		1, 0, 0, 0, 193, 195, 5, 20, 0, 0, 194, 193, 1, 0, 0, 0, 195, 198, 1, 0,
		0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 202, 1, 0, 0, 0,
		198, 196, 1, 0, 0, 0, 199, 201, 5, 30, 0, 0, 200, 199, 1, 0, 0, 0, 201,
		204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 208,
		1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 207, 5, 19, 0, 0, 206, 205, 1, 0,
		0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0,
		209, 211, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 212, 5, 24, 0, 0, 212,
		9, 1, 0, 0, 0, 213, 214, 5, 18, 0, 0, 214, 11, 1, 0, 0, 0, 215, 216, 7,
		0, 0, 0, 216, 13, 1, 0, 0, 0, 25, 21, 23, 31, 37, 49, 55, 67, 73, 85, 91,
		103, 109, 121, 127, 139, 145, 157, 163, 170, 178, 185, 190, 196, 202, 208,
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
	XSQLParserT__0              = 1
	XSQLParserT__1              = 2
	XSQLParserT__2              = 3
	XSQLParserT__3              = 4
	XSQLParserT__4              = 5
	XSQLParserT__5              = 6
	XSQLParserT__6              = 7
	XSQLParserT__7              = 8
	XSQLParserT__8              = 9
	XSQLParserT__9              = 10
	XSQLParserT__10             = 11
	XSQLParserT__11             = 12
	XSQLParserT__12             = 13
	XSQLParserT__13             = 14
	XSQLParserT__14             = 15
	XSQLParserT__15             = 16
	XSQLParserBLOCK_COMMENT     = 17
	XSQLParserEntityRef         = 18
	XSQLParserWS                = 19
	XSQLParserNAME              = 20
	XSQLParserDOLLAR            = 21
	XSQLParserHASH              = 22
	XSQLParserOPEN_CURLY_BRAXE  = 23
	XSQLParserCLOSE_CURLY_BRAXE = 24
	XSQLParserOPEN              = 25
	XSQLParserCLOSE             = 26
	XSQLParserSLASH             = 27
	XSQLParserEQUALS            = 28
	XSQLParserSTRING            = 29
	XSQLParserTEXT              = 30
)

// XSQLParser rules.
const (
	XSQLParserRULE_document  = 0
	XSQLParserRULE_content   = 1
	XSQLParserRULE_element   = 2
	XSQLParserRULE_attribute = 3
	XSQLParserRULE_expr      = 4
	XSQLParserRULE_reference = 5
	XSQLParserRULE_chardata  = 6
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Content() IContentContext
	EOF() antlr.TerminalNode

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XSQLParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XSQLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) Content() IContentContext {
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

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(XSQLParserEOF, 0)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, XSQLParserRULE_document)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Content()
	}
	{
		p.SetState(15)
		p.Match(XSQLParserEOF)
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

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllElement() []IElementContext
	Element(i int) IElementContext
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
	p.EnterRule(localctx, 2, XSQLParserRULE_content)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147221504) != 0 {
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(17)
				p.Element()
			}

		case 2:
			{
				p.SetState(18)
				p.Expr()
			}

		case 3:
			{
				p.SetState(19)
				p.Reference()
			}

		case 4:
			{
				p.SetState(20)
				p.Chardata()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(25)
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

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	OPEN() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	Content() IContentContext
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
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

func (s *ElementContext) GetName() antlr.Token { return s.name }

func (s *ElementContext) SetName(v antlr.Token) { s.name = v }

func (s *ElementContext) OPEN() antlr.TerminalNode {
	return s.GetToken(XSQLParserOPEN, 0)
}

func (s *ElementContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(XSQLParserCLOSE, 0)
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

func (s *ElementContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(XSQLParserWS)
}

func (s *ElementContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(XSQLParserWS, i)
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

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *XSQLParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, XSQLParserRULE_element)
	var _la int

	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)

			var _m = p.Match(XSQLParserT__0)

			localctx.(*ElementContext).name = _m
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

		for _la == XSQLParserWS {
			{
				p.SetState(28)
				p.Match(XSQLParserWS)
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
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(34)
				p.Attribute()
			}

			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(40)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Content()
		}
		{
			p.SetState(42)
			p.Match(XSQLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)

			var _m = p.Match(XSQLParserT__2)

			localctx.(*ElementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(46)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(52)
				p.Attribute()
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(58)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Content()
		}
		{
			p.SetState(60)
			p.Match(XSQLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)

			var _m = p.Match(XSQLParserT__4)

			localctx.(*ElementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(64)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(70)
				p.Attribute()
			}

			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(76)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Content()
		}
		{
			p.SetState(78)
			p.Match(XSQLParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)

			var _m = p.Match(XSQLParserT__6)

			localctx.(*ElementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(82)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(88)
				p.Attribute()
			}

			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(94)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Content()
		}
		{
			p.SetState(96)
			p.Match(XSQLParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)

			var _m = p.Match(XSQLParserT__8)

			localctx.(*ElementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(100)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(106)
				p.Attribute()
			}

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Content()
		}
		{
			p.SetState(114)
			p.Match(XSQLParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)

			var _m = p.Match(XSQLParserT__10)

			localctx.(*ElementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(118)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(124)
				p.Attribute()
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(130)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Content()
		}
		{
			p.SetState(132)
			p.Match(XSQLParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(134)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)

			var _m = p.Match(XSQLParserT__12)

			localctx.(*ElementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(136)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(142)
				p.Attribute()
			}

			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(148)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Content()
		}
		{
			p.SetState(150)
			p.Match(XSQLParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(152)
			p.Match(XSQLParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)

			var _m = p.Match(XSQLParserT__14)

			localctx.(*ElementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserWS {
			{
				p.SetState(154)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == XSQLParserNAME {
			{
				p.SetState(160)
				p.Attribute()
			}

			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(166)
			p.Match(XSQLParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Content()
		}
		{
			p.SetState(168)
			p.Match(XSQLParserT__15)
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
		p.SetState(172)
		p.Match(XSQLParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(XSQLParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(XSQLParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserWS {
		{
			p.SetState(175)
			p.Match(XSQLParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(180)
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
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case XSQLParserDOLLAR:
		{
			p.SetState(181)
			p.Match(XSQLParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(XSQLParserOPEN_CURLY_BRAXE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case XSQLParserHASH:
		{
			p.SetState(183)
			p.Match(XSQLParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
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
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(187)
				p.Match(XSQLParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserNAME {
		{
			p.SetState(193)

			var _m = p.Match(XSQLParserNAME)

			localctx.(*ExprContext).val = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserTEXT {
		{
			p.SetState(199)
			p.Match(XSQLParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XSQLParserWS {
		{
			p.SetState(205)
			p.Match(XSQLParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(211)
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
		p.SetState(213)
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
		p.SetState(215)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2146959360) != 0) {
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
