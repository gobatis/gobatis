// Code generated from JsonPath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package gpath

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JsonPathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var JsonPathLexerLexerStaticData struct {
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

func jsonpathlexerLexerInit() {
	staticData := &JsonPathLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'$.'", "'.'", "'[*]'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "INDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "INDENTIFIER", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 34, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 5, 3, 23, 8, 3, 10, 3, 12, 3, 26, 9, 3, 1, 4, 4, 4, 29, 8, 4, 11, 4,
		12, 4, 30, 1, 4, 1, 4, 0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0, 3,
		2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 3, 0, 9, 10, 13,
		13, 32, 32, 35, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 14, 1, 0, 0, 0,
		5, 16, 1, 0, 0, 0, 7, 20, 1, 0, 0, 0, 9, 28, 1, 0, 0, 0, 11, 12, 5, 36,
		0, 0, 12, 13, 5, 46, 0, 0, 13, 2, 1, 0, 0, 0, 14, 15, 5, 46, 0, 0, 15,
		4, 1, 0, 0, 0, 16, 17, 5, 91, 0, 0, 17, 18, 5, 42, 0, 0, 18, 19, 5, 93,
		0, 0, 19, 6, 1, 0, 0, 0, 20, 24, 7, 0, 0, 0, 21, 23, 7, 1, 0, 0, 22, 21,
		1, 0, 0, 0, 23, 26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0,
		25, 8, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 27, 29, 7, 2, 0, 0, 28, 27, 1, 0,
		0, 0, 29, 30, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32,
		1, 0, 0, 0, 32, 33, 6, 4, 0, 0, 33, 10, 1, 0, 0, 0, 3, 0, 24, 30, 1, 6,
		0, 0,
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

// JsonPathLexerInit initializes any static state used to implement JsonPathLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonPathLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonPathLexerInit() {
	staticData := &JsonPathLexerLexerStaticData
	staticData.once.Do(jsonpathlexerLexerInit)
}

// NewJsonPathLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonPathLexer(input antlr.CharStream) *JsonPathLexer {
	JsonPathLexerInit()
	l := new(JsonPathLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &JsonPathLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "JsonPath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonPathLexer tokens.
const (
	JsonPathLexerT__0        = 1
	JsonPathLexerT__1        = 2
	JsonPathLexerT__2        = 3
	JsonPathLexerINDENTIFIER = 4
	JsonPathLexerWS          = 5
)
