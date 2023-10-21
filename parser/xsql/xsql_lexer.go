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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "BLOCK_COMMENT",
		"EntityRef", "WS", "NAME", "DOLLAR", "HASH", "OPEN_CURLY_BRAXE", "CLOSE_CURLY_BRAXE",
		"OPEN", "CLOSE", "SLASH", "EQUALS", "STRING", "TEXT", "DIGIT", "Char",
		"NameChar",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 273, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 193, 8, 16, 10, 16, 12, 16, 196, 9,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 3, 18, 210, 8, 18, 1, 18, 4, 18, 213, 8, 18, 11, 18, 12,
		18, 214, 1, 19, 1, 19, 5, 19, 219, 8, 19, 10, 19, 12, 19, 222, 9, 19, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 5, 28, 242, 8, 28, 10,
		28, 12, 28, 245, 9, 28, 1, 28, 1, 28, 1, 28, 5, 28, 250, 8, 28, 10, 28,
		12, 28, 253, 9, 28, 1, 28, 3, 28, 256, 8, 28, 1, 29, 4, 29, 259, 8, 29,
		11, 29, 12, 29, 260, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 269,
		8, 31, 1, 32, 3, 32, 272, 8, 32, 2, 194, 260, 0, 33, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 0,
		63, 0, 65, 0, 1, 0, 7, 2, 0, 9, 9, 32, 32, 1, 0, 34, 34, 1, 0, 39, 39,
		1, 0, 48, 57, 2, 0, 45, 46, 95, 95, 3, 0, 183, 183, 768, 879, 8255, 8256,
		8, 0, 58, 58, 65, 90, 97, 122, 8304, 8591, 11264, 12271, 12289, 55295,
		63744, 64975, 65008, 65533, 281, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 1, 67, 1, 0, 0, 0, 3, 70, 1, 0, 0, 0, 5, 76, 1, 0, 0, 0,
		7, 83, 1, 0, 0, 0, 9, 93, 1, 0, 0, 0, 11, 98, 1, 0, 0, 0, 13, 106, 1, 0,
		0, 0, 15, 116, 1, 0, 0, 0, 17, 129, 1, 0, 0, 0, 19, 134, 1, 0, 0, 0, 21,
		141, 1, 0, 0, 0, 23, 147, 1, 0, 0, 0, 25, 156, 1, 0, 0, 0, 27, 160, 1,
		0, 0, 0, 29, 167, 1, 0, 0, 0, 31, 175, 1, 0, 0, 0, 33, 186, 1, 0, 0, 0,
		35, 203, 1, 0, 0, 0, 37, 212, 1, 0, 0, 0, 39, 216, 1, 0, 0, 0, 41, 223,
		1, 0, 0, 0, 43, 225, 1, 0, 0, 0, 45, 227, 1, 0, 0, 0, 47, 229, 1, 0, 0,
		0, 49, 231, 1, 0, 0, 0, 51, 233, 1, 0, 0, 0, 53, 235, 1, 0, 0, 0, 55, 237,
		1, 0, 0, 0, 57, 255, 1, 0, 0, 0, 59, 258, 1, 0, 0, 0, 61, 262, 1, 0, 0,
		0, 63, 268, 1, 0, 0, 0, 65, 271, 1, 0, 0, 0, 67, 68, 5, 105, 0, 0, 68,
		69, 5, 102, 0, 0, 69, 2, 1, 0, 0, 0, 70, 71, 5, 60, 0, 0, 71, 72, 5, 47,
		0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 102, 0, 0, 74, 75, 5, 62, 0, 0,
		75, 4, 1, 0, 0, 0, 76, 77, 5, 99, 0, 0, 77, 78, 5, 104, 0, 0, 78, 79, 5,
		111, 0, 0, 79, 80, 5, 111, 0, 0, 80, 81, 5, 115, 0, 0, 81, 82, 5, 101,
		0, 0, 82, 6, 1, 0, 0, 0, 83, 84, 5, 60, 0, 0, 84, 85, 5, 47, 0, 0, 85,
		86, 5, 99, 0, 0, 86, 87, 5, 104, 0, 0, 87, 88, 5, 111, 0, 0, 88, 89, 5,
		111, 0, 0, 89, 90, 5, 115, 0, 0, 90, 91, 5, 101, 0, 0, 91, 92, 5, 62, 0,
		0, 92, 8, 1, 0, 0, 0, 93, 94, 5, 119, 0, 0, 94, 95, 5, 104, 0, 0, 95, 96,
		5, 101, 0, 0, 96, 97, 5, 110, 0, 0, 97, 10, 1, 0, 0, 0, 98, 99, 5, 60,
		0, 0, 99, 100, 5, 47, 0, 0, 100, 101, 5, 119, 0, 0, 101, 102, 5, 104, 0,
		0, 102, 103, 5, 101, 0, 0, 103, 104, 5, 110, 0, 0, 104, 105, 5, 62, 0,
		0, 105, 12, 1, 0, 0, 0, 106, 107, 5, 111, 0, 0, 107, 108, 5, 116, 0, 0,
		108, 109, 5, 104, 0, 0, 109, 110, 5, 101, 0, 0, 110, 111, 5, 114, 0, 0,
		111, 112, 5, 119, 0, 0, 112, 113, 5, 105, 0, 0, 113, 114, 5, 115, 0, 0,
		114, 115, 5, 101, 0, 0, 115, 14, 1, 0, 0, 0, 116, 117, 5, 60, 0, 0, 117,
		118, 5, 47, 0, 0, 118, 119, 5, 111, 0, 0, 119, 120, 5, 116, 0, 0, 120,
		121, 5, 104, 0, 0, 121, 122, 5, 101, 0, 0, 122, 123, 5, 114, 0, 0, 123,
		124, 5, 119, 0, 0, 124, 125, 5, 105, 0, 0, 125, 126, 5, 115, 0, 0, 126,
		127, 5, 101, 0, 0, 127, 128, 5, 62, 0, 0, 128, 16, 1, 0, 0, 0, 129, 130,
		5, 116, 0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 105, 0, 0, 132, 133,
		5, 109, 0, 0, 133, 18, 1, 0, 0, 0, 134, 135, 5, 60, 0, 0, 135, 136, 5,
		47, 0, 0, 136, 137, 5, 116, 0, 0, 137, 138, 5, 114, 0, 0, 138, 139, 5,
		105, 0, 0, 139, 140, 5, 109, 0, 0, 140, 20, 1, 0, 0, 0, 141, 142, 5, 119,
		0, 0, 142, 143, 5, 104, 0, 0, 143, 144, 5, 101, 0, 0, 144, 145, 5, 114,
		0, 0, 145, 146, 5, 101, 0, 0, 146, 22, 1, 0, 0, 0, 147, 148, 5, 60, 0,
		0, 148, 149, 5, 47, 0, 0, 149, 150, 5, 119, 0, 0, 150, 151, 5, 104, 0,
		0, 151, 152, 5, 101, 0, 0, 152, 153, 5, 114, 0, 0, 153, 154, 5, 101, 0,
		0, 154, 155, 5, 62, 0, 0, 155, 24, 1, 0, 0, 0, 156, 157, 5, 115, 0, 0,
		157, 158, 5, 101, 0, 0, 158, 159, 5, 116, 0, 0, 159, 26, 1, 0, 0, 0, 160,
		161, 5, 60, 0, 0, 161, 162, 5, 47, 0, 0, 162, 163, 5, 115, 0, 0, 163, 164,
		5, 101, 0, 0, 164, 165, 5, 116, 0, 0, 165, 166, 5, 62, 0, 0, 166, 28, 1,
		0, 0, 0, 167, 168, 5, 102, 0, 0, 168, 169, 5, 111, 0, 0, 169, 170, 5, 114,
		0, 0, 170, 171, 5, 101, 0, 0, 171, 172, 5, 97, 0, 0, 172, 173, 5, 99, 0,
		0, 173, 174, 5, 104, 0, 0, 174, 30, 1, 0, 0, 0, 175, 176, 5, 60, 0, 0,
		176, 177, 5, 47, 0, 0, 177, 178, 5, 102, 0, 0, 178, 179, 5, 111, 0, 0,
		179, 180, 5, 114, 0, 0, 180, 181, 5, 101, 0, 0, 181, 182, 5, 97, 0, 0,
		182, 183, 5, 99, 0, 0, 183, 184, 5, 104, 0, 0, 184, 185, 5, 62, 0, 0, 185,
		32, 1, 0, 0, 0, 186, 187, 5, 60, 0, 0, 187, 188, 5, 33, 0, 0, 188, 189,
		5, 45, 0, 0, 189, 190, 5, 45, 0, 0, 190, 194, 1, 0, 0, 0, 191, 193, 9,
		0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 195, 1, 0, 0,
		0, 194, 192, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197,
		198, 5, 45, 0, 0, 198, 199, 5, 45, 0, 0, 199, 200, 5, 62, 0, 0, 200, 201,
		1, 0, 0, 0, 201, 202, 6, 16, 0, 0, 202, 34, 1, 0, 0, 0, 203, 204, 5, 38,
		0, 0, 204, 205, 3, 39, 19, 0, 205, 206, 5, 59, 0, 0, 206, 36, 1, 0, 0,
		0, 207, 213, 7, 0, 0, 0, 208, 210, 5, 13, 0, 0, 209, 208, 1, 0, 0, 0, 209,
		210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213, 5, 10, 0, 0, 212, 207,
		1, 0, 0, 0, 212, 209, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 212, 1, 0,
		0, 0, 214, 215, 1, 0, 0, 0, 215, 38, 1, 0, 0, 0, 216, 220, 3, 65, 32, 0,
		217, 219, 3, 63, 31, 0, 218, 217, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220,
		218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 40, 1, 0, 0, 0, 222, 220, 1,
		0, 0, 0, 223, 224, 5, 36, 0, 0, 224, 42, 1, 0, 0, 0, 225, 226, 5, 35, 0,
		0, 226, 44, 1, 0, 0, 0, 227, 228, 5, 123, 0, 0, 228, 46, 1, 0, 0, 0, 229,
		230, 5, 125, 0, 0, 230, 48, 1, 0, 0, 0, 231, 232, 5, 60, 0, 0, 232, 50,
		1, 0, 0, 0, 233, 234, 5, 62, 0, 0, 234, 52, 1, 0, 0, 0, 235, 236, 5, 47,
		0, 0, 236, 54, 1, 0, 0, 0, 237, 238, 5, 61, 0, 0, 238, 56, 1, 0, 0, 0,
		239, 243, 5, 34, 0, 0, 240, 242, 8, 1, 0, 0, 241, 240, 1, 0, 0, 0, 242,
		245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 246,
		1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 256, 5, 34, 0, 0, 247, 251, 5, 39,
		0, 0, 248, 250, 8, 2, 0, 0, 249, 248, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0,
		251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 254, 1, 0, 0, 0, 253,
		251, 1, 0, 0, 0, 254, 256, 5, 39, 0, 0, 255, 239, 1, 0, 0, 0, 255, 247,
		1, 0, 0, 0, 256, 58, 1, 0, 0, 0, 257, 259, 9, 0, 0, 0, 258, 257, 1, 0,
		0, 0, 259, 260, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0,
		261, 60, 1, 0, 0, 0, 262, 263, 7, 3, 0, 0, 263, 62, 1, 0, 0, 0, 264, 269,
		3, 65, 32, 0, 265, 269, 7, 4, 0, 0, 266, 269, 3, 61, 30, 0, 267, 269, 7,
		5, 0, 0, 268, 264, 1, 0, 0, 0, 268, 265, 1, 0, 0, 0, 268, 266, 1, 0, 0,
		0, 268, 267, 1, 0, 0, 0, 269, 64, 1, 0, 0, 0, 270, 272, 7, 6, 0, 0, 271,
		270, 1, 0, 0, 0, 272, 66, 1, 0, 0, 0, 12, 0, 194, 209, 212, 214, 220, 243,
		251, 255, 260, 268, 271, 1, 6, 0, 0,
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
	XSQLLexerT__0              = 1
	XSQLLexerT__1              = 2
	XSQLLexerT__2              = 3
	XSQLLexerT__3              = 4
	XSQLLexerT__4              = 5
	XSQLLexerT__5              = 6
	XSQLLexerT__6              = 7
	XSQLLexerT__7              = 8
	XSQLLexerT__8              = 9
	XSQLLexerT__9              = 10
	XSQLLexerT__10             = 11
	XSQLLexerT__11             = 12
	XSQLLexerT__12             = 13
	XSQLLexerT__13             = 14
	XSQLLexerT__14             = 15
	XSQLLexerT__15             = 16
	XSQLLexerBLOCK_COMMENT     = 17
	XSQLLexerEntityRef         = 18
	XSQLLexerWS                = 19
	XSQLLexerNAME              = 20
	XSQLLexerDOLLAR            = 21
	XSQLLexerHASH              = 22
	XSQLLexerOPEN_CURLY_BRAXE  = 23
	XSQLLexerCLOSE_CURLY_BRAXE = 24
	XSQLLexerOPEN              = 25
	XSQLLexerCLOSE             = 26
	XSQLLexerSLASH             = 27
	XSQLLexerEQUALS            = 28
	XSQLLexerSTRING            = 29
	XSQLLexerTEXT              = 30
)
