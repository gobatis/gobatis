// Code generated from XSQL.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"DEFAULT_MODE",
	}
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
		"BLOCK_COMMENT", "LINE_COMMENT", "EntityRef", "WS", "NAME", "DOLLAR",
		"HASH", "OPEN_CURLY_BRAXE", "CLOSE_CURLY_BRAXE", "OPEN", "CLOSE", "SLASH",
		"EQUALS", "STRING", "TEXT", "DIGIT", "Char", "NameChar",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 135, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		5, 0, 44, 8, 0, 10, 0, 12, 0, 47, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 59, 8, 1, 10, 1, 12, 1, 62, 9, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 72, 8, 3, 1, 3, 4, 3, 75,
		8, 3, 11, 3, 12, 3, 76, 1, 4, 1, 4, 5, 4, 81, 8, 4, 10, 4, 12, 4, 84, 9,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 104, 8, 13, 10, 13,
		12, 13, 107, 9, 13, 1, 13, 1, 13, 1, 13, 5, 13, 112, 8, 13, 10, 13, 12,
		13, 115, 9, 13, 1, 13, 3, 13, 118, 8, 13, 1, 14, 4, 14, 121, 8, 14, 11,
		14, 12, 14, 122, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 131,
		8, 16, 1, 17, 3, 17, 134, 8, 17, 2, 45, 122, 0, 18, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 0, 33, 0, 35, 0, 1, 0, 8, 2, 0, 10, 10, 13, 13, 2, 0, 9,
		9, 32, 32, 1, 0, 34, 34, 1, 0, 39, 39, 1, 0, 48, 57, 2, 0, 45, 46, 95,
		95, 3, 0, 183, 183, 768, 879, 8255, 8256, 8, 0, 58, 58, 65, 90, 97, 122,
		8304, 8591, 11264, 12271, 12289, 55295, 63744, 64975, 65008, 65533, 144,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 37, 1,
		0, 0, 0, 3, 54, 1, 0, 0, 0, 5, 65, 1, 0, 0, 0, 7, 74, 1, 0, 0, 0, 9, 78,
		1, 0, 0, 0, 11, 85, 1, 0, 0, 0, 13, 87, 1, 0, 0, 0, 15, 89, 1, 0, 0, 0,
		17, 91, 1, 0, 0, 0, 19, 93, 1, 0, 0, 0, 21, 95, 1, 0, 0, 0, 23, 97, 1,
		0, 0, 0, 25, 99, 1, 0, 0, 0, 27, 117, 1, 0, 0, 0, 29, 120, 1, 0, 0, 0,
		31, 124, 1, 0, 0, 0, 33, 130, 1, 0, 0, 0, 35, 133, 1, 0, 0, 0, 37, 38,
		5, 60, 0, 0, 38, 39, 5, 33, 0, 0, 39, 40, 5, 45, 0, 0, 40, 41, 5, 45, 0,
		0, 41, 45, 1, 0, 0, 0, 42, 44, 9, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 47,
		1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 48, 49, 5, 45, 0, 0, 49, 50, 5, 45, 0, 0, 50, 51, 5,
		62, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 6, 0, 0, 0, 53, 2, 1, 0, 0, 0, 54,
		55, 5, 47, 0, 0, 55, 56, 5, 47, 0, 0, 56, 60, 1, 0, 0, 0, 57, 59, 8, 0,
		0, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61,
		1, 0, 0, 0, 61, 63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 64, 6, 1, 0, 0,
		64, 4, 1, 0, 0, 0, 65, 66, 5, 38, 0, 0, 66, 67, 3, 9, 4, 0, 67, 68, 5,
		59, 0, 0, 68, 6, 1, 0, 0, 0, 69, 75, 7, 1, 0, 0, 70, 72, 5, 13, 0, 0, 71,
		70, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 75, 5, 10,
		0, 0, 74, 69, 1, 0, 0, 0, 74, 71, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 74,
		1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 8, 1, 0, 0, 0, 78, 82, 3, 35, 17, 0,
		79, 81, 3, 33, 16, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1,
		0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 10, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85,
		86, 5, 36, 0, 0, 86, 12, 1, 0, 0, 0, 87, 88, 5, 35, 0, 0, 88, 14, 1, 0,
		0, 0, 89, 90, 5, 123, 0, 0, 90, 16, 1, 0, 0, 0, 91, 92, 5, 125, 0, 0, 92,
		18, 1, 0, 0, 0, 93, 94, 5, 60, 0, 0, 94, 20, 1, 0, 0, 0, 95, 96, 5, 62,
		0, 0, 96, 22, 1, 0, 0, 0, 97, 98, 5, 47, 0, 0, 98, 24, 1, 0, 0, 0, 99,
		100, 5, 61, 0, 0, 100, 26, 1, 0, 0, 0, 101, 105, 5, 34, 0, 0, 102, 104,
		8, 2, 0, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0,
		0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0,
		108, 118, 5, 34, 0, 0, 109, 113, 5, 39, 0, 0, 110, 112, 8, 3, 0, 0, 111,
		110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114,
		1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 118, 5, 39,
		0, 0, 117, 101, 1, 0, 0, 0, 117, 109, 1, 0, 0, 0, 118, 28, 1, 0, 0, 0,
		119, 121, 9, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 30, 1, 0, 0, 0, 124, 125, 7,
		4, 0, 0, 125, 32, 1, 0, 0, 0, 126, 131, 3, 35, 17, 0, 127, 131, 7, 5, 0,
		0, 128, 131, 3, 31, 15, 0, 129, 131, 7, 6, 0, 0, 130, 126, 1, 0, 0, 0,
		130, 127, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131,
		34, 1, 0, 0, 0, 132, 134, 7, 7, 0, 0, 133, 132, 1, 0, 0, 0, 134, 36, 1,
		0, 0, 0, 13, 0, 45, 60, 71, 74, 76, 82, 105, 113, 117, 122, 130, 133, 1,
		6, 0, 0,
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
	l.GrammarFileName = "XSQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// XSQLLexer tokens.
const (
	XSQLLexerBLOCK_COMMENT     = 1
	XSQLLexerLINE_COMMENT      = 2
	XSQLLexerEntityRef         = 3
	XSQLLexerWS                = 4
	XSQLLexerNAME              = 5
	XSQLLexerDOLLAR            = 6
	XSQLLexerHASH              = 7
	XSQLLexerOPEN_CURLY_BRAXE  = 8
	XSQLLexerCLOSE_CURLY_BRAXE = 9
	XSQLLexerOPEN              = 10
	XSQLLexerCLOSE             = 11
	XSQLLexerSLASH             = 12
	XSQLLexerEQUALS            = 13
	XSQLLexerSTRING            = 14
	XSQLLexerTEXT              = 15
)
