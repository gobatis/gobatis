// Code generated from XSQLLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package xsql

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type XSQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var XSQLLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func xsqllexerLexerInit() {
	staticData := &XSQLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "EXPR", "INSIDE",
	}
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
		"COMMENT", "EntityRef", "SEA_WS", "OPEN", "EXPR_OPEN", "TEXT", "DOLLAR_NOT_LBRACE",
		"HASH_NOT_LBRACE", "EXPR_CLOSE", "EXPR_VAL", "S1", "Name", "CLOSE",
		"SLASH_CLOSE", "SLASH", "EQUALS", "STRING", "S2", "DIGIT", "NameChar",
		"NameStartChar",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 167, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2,
		2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8,
		2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2,
		14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19,
		7, 19, 2, 20, 7, 20, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 52, 8, 0,
		10, 0, 12, 0, 55, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 69, 8, 2, 1, 2, 4, 2, 72, 8, 2, 11, 2, 12,
		2, 73, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 84, 8, 4,
		1, 4, 1, 4, 1, 5, 4, 5, 89, 8, 5, 11, 5, 12, 5, 90, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 106, 8,
		9, 10, 9, 12, 9, 109, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 5,
		11, 117, 8, 11, 10, 11, 12, 11, 120, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 5, 16, 137, 8, 16, 10, 16, 12, 16, 140, 9, 16, 1, 16, 1, 16, 1, 16,
		5, 16, 145, 8, 16, 10, 16, 12, 16, 148, 9, 16, 1, 16, 3, 16, 151, 8, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3,
		19, 163, 8, 19, 1, 20, 3, 20, 166, 8, 20, 1, 53, 0, 21, 3, 1, 5, 2, 7,
		3, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 27,
		13, 29, 14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 0, 41, 0, 43, 0, 3, 0,
		1, 2, 11, 2, 0, 9, 9, 32, 32, 3, 0, 35, 36, 38, 38, 60, 60, 1, 0, 123,
		123, 3, 0, 46, 46, 91, 91, 93, 93, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 34,
		34, 60, 60, 2, 0, 39, 39, 60, 60, 1, 0, 48, 57, 2, 0, 45, 46, 95, 95, 3,
		0, 183, 183, 768, 879, 8255, 8256, 8, 0, 58, 58, 65, 90, 97, 122, 8304,
		8591, 11264, 12271, 12289, 55295, 63744, 64975, 65008, 65533, 176, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1,
		19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 1, 23, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0,
		2, 27, 1, 0, 0, 0, 2, 29, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 2, 33, 1, 0, 0,
		0, 2, 35, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5, 62, 1, 0,
		0, 0, 7, 71, 1, 0, 0, 0, 9, 75, 1, 0, 0, 0, 11, 83, 1, 0, 0, 0, 13, 88,
		1, 0, 0, 0, 15, 92, 1, 0, 0, 0, 17, 95, 1, 0, 0, 0, 19, 98, 1, 0, 0, 0,
		21, 102, 1, 0, 0, 0, 23, 110, 1, 0, 0, 0, 25, 114, 1, 0, 0, 0, 27, 121,
		1, 0, 0, 0, 29, 125, 1, 0, 0, 0, 31, 130, 1, 0, 0, 0, 33, 132, 1, 0, 0,
		0, 35, 150, 1, 0, 0, 0, 37, 152, 1, 0, 0, 0, 39, 156, 1, 0, 0, 0, 41, 162,
		1, 0, 0, 0, 43, 165, 1, 0, 0, 0, 45, 46, 5, 60, 0, 0, 46, 47, 5, 33, 0,
		0, 47, 48, 5, 45, 0, 0, 48, 49, 5, 45, 0, 0, 49, 53, 1, 0, 0, 0, 50, 52,
		9, 0, 0, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		53, 51, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5,
		45, 0, 0, 57, 58, 5, 45, 0, 0, 58, 59, 5, 62, 0, 0, 59, 60, 1, 0, 0, 0,
		60, 61, 6, 0, 0, 0, 61, 4, 1, 0, 0, 0, 62, 63, 5, 38, 0, 0, 63, 64, 3,
		25, 11, 0, 64, 65, 5, 59, 0, 0, 65, 6, 1, 0, 0, 0, 66, 72, 7, 0, 0, 0,
		67, 69, 5, 13, 0, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 1,
		0, 0, 0, 70, 72, 5, 10, 0, 0, 71, 66, 1, 0, 0, 0, 71, 68, 1, 0, 0, 0, 72,
		73, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 8, 1, 0, 0,
		0, 75, 76, 5, 60, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 6, 3, 1, 0, 78, 10,
		1, 0, 0, 0, 79, 80, 5, 35, 0, 0, 80, 84, 5, 123, 0, 0, 81, 82, 5, 36, 0,
		0, 82, 84, 5, 123, 0, 0, 83, 79, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 85,
		1, 0, 0, 0, 85, 86, 6, 4, 2, 0, 86, 12, 1, 0, 0, 0, 87, 89, 8, 1, 0, 0,
		88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1,
		0, 0, 0, 91, 14, 1, 0, 0, 0, 92, 93, 5, 36, 0, 0, 93, 94, 8, 2, 0, 0, 94,
		16, 1, 0, 0, 0, 95, 96, 5, 35, 0, 0, 96, 97, 8, 2, 0, 0, 97, 18, 1, 0,
		0, 0, 98, 99, 5, 125, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 6, 8, 3, 0,
		101, 20, 1, 0, 0, 0, 102, 107, 3, 43, 20, 0, 103, 106, 3, 41, 19, 0, 104,
		106, 7, 3, 0, 0, 105, 103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 109,
		1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 22, 1, 0,
		0, 0, 109, 107, 1, 0, 0, 0, 110, 111, 7, 4, 0, 0, 111, 112, 1, 0, 0, 0,
		112, 113, 6, 10, 0, 0, 113, 24, 1, 0, 0, 0, 114, 118, 3, 43, 20, 0, 115,
		117, 3, 41, 19, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116,
		1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 26, 1, 0, 0, 0, 120, 118, 1, 0,
		0, 0, 121, 122, 5, 62, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 6, 12, 3,
		0, 124, 28, 1, 0, 0, 0, 125, 126, 5, 47, 0, 0, 126, 127, 5, 62, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 129, 6, 13, 3, 0, 129, 30, 1, 0, 0, 0, 130, 131,
		5, 47, 0, 0, 131, 32, 1, 0, 0, 0, 132, 133, 5, 61, 0, 0, 133, 34, 1, 0,
		0, 0, 134, 138, 5, 34, 0, 0, 135, 137, 8, 5, 0, 0, 136, 135, 1, 0, 0, 0,
		137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 151, 5, 34, 0, 0, 142, 146,
		5, 39, 0, 0, 143, 145, 8, 6, 0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0,
		0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0,
		148, 146, 1, 0, 0, 0, 149, 151, 5, 39, 0, 0, 150, 134, 1, 0, 0, 0, 150,
		142, 1, 0, 0, 0, 151, 36, 1, 0, 0, 0, 152, 153, 7, 4, 0, 0, 153, 154, 1,
		0, 0, 0, 154, 155, 6, 17, 0, 0, 155, 38, 1, 0, 0, 0, 156, 157, 7, 7, 0,
		0, 157, 40, 1, 0, 0, 0, 158, 163, 3, 43, 20, 0, 159, 163, 7, 8, 0, 0, 160,
		163, 3, 39, 18, 0, 161, 163, 7, 9, 0, 0, 162, 158, 1, 0, 0, 0, 162, 159,
		1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 42, 1, 0,
		0, 0, 164, 166, 7, 10, 0, 0, 165, 164, 1, 0, 0, 0, 166, 44, 1, 0, 0, 0,
		17, 0, 1, 2, 53, 68, 71, 73, 83, 90, 105, 107, 118, 138, 146, 150, 162,
		165, 4, 6, 0, 0, 5, 2, 0, 5, 1, 0, 4, 0, 0,
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

// XSQLLexerInit initializes any static state used to implement XSQLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewXSQLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func XSQLLexerInit() {
	staticData := &XSQLLexerLexerStaticData
	staticData.once.Do(xsqllexerLexerInit)
}

// NewXSQLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewXSQLLexer(input antlr.CharStream) *XSQLLexer {
	XSQLLexerInit()
	l := new(XSQLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &XSQLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "XSQLLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// XSQLLexer tokens.
const (
	XSQLLexerCOMMENT           = 1
	XSQLLexerEntityRef         = 2
	XSQLLexerSEA_WS            = 3
	XSQLLexerOPEN              = 4
	XSQLLexerEXPR_OPEN         = 5
	XSQLLexerTEXT              = 6
	XSQLLexerDOLLAR_NOT_LBRACE = 7
	XSQLLexerHASH_NOT_LBRACE   = 8
	XSQLLexerEXPR_CLOSE        = 9
	XSQLLexerEXPR_VAL          = 10
	XSQLLexerS1                = 11
	XSQLLexerName              = 12
	XSQLLexerCLOSE             = 13
	XSQLLexerSLASH_CLOSE       = 14
	XSQLLexerSLASH             = 15
	XSQLLexerEQUALS            = 16
	XSQLLexerSTRING            = 17
	XSQLLexerS2                = 18
)

// XSQLLexer modes.
const (
	XSQLLexerEXPR = iota + 1
	XSQLLexerINSIDE
)
