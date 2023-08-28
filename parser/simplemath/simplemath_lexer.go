// Code generated from SimpleMath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package simplemath

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

type SimpleMathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SimpleMathLexerLexerStaticData struct {
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

func simplemathlexerLexerInit() {
	staticData := &SimpleMathLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'+'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "NUMBER", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "NUMBER", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 25, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 4, 2, 15, 8, 2, 11, 2, 12, 2, 16, 1, 3, 4, 3,
		20, 8, 3, 11, 3, 12, 3, 21, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4,
		1, 0, 2, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 26, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0,
		0, 3, 11, 1, 0, 0, 0, 5, 14, 1, 0, 0, 0, 7, 19, 1, 0, 0, 0, 9, 10, 5, 43,
		0, 0, 10, 2, 1, 0, 0, 0, 11, 12, 5, 42, 0, 0, 12, 4, 1, 0, 0, 0, 13, 15,
		7, 0, 0, 0, 14, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 14, 1, 0, 0, 0,
		16, 17, 1, 0, 0, 0, 17, 6, 1, 0, 0, 0, 18, 20, 7, 1, 0, 0, 19, 18, 1, 0,
		0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23,
		1, 0, 0, 0, 23, 24, 6, 3, 0, 0, 24, 8, 1, 0, 0, 0, 3, 0, 16, 21, 1, 6,
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

// SimpleMathLexerInit initializes any static state used to implement SimpleMathLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSimpleMathLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleMathLexerInit() {
	staticData := &SimpleMathLexerLexerStaticData
	staticData.once.Do(simplemathlexerLexerInit)
}

// NewSimpleMathLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSimpleMathLexer(input antlr.CharStream) *SimpleMathLexer {
	SimpleMathLexerInit()
	l := new(SimpleMathLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SimpleMathLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SimpleMath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleMathLexer tokens.
const (
	SimpleMathLexerT__0   = 1
	SimpleMathLexerT__1   = 2
	SimpleMathLexerNUMBER = 3
	SimpleMathLexerWS     = 4
)
