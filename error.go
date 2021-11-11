package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/parser/xml"
	"reflect"
)

const (
	unknown_err = iota + 1
	parameter_type_err
	parameter_conflict_with_builtin_err
	parameter_not_found_err
	var_to_reflect_kind_err
	var_bind_err
	pop_value_err
	pop_params_err
	pop_binary_operands_err
	pop_tertiary_operands_err
	pop_result_err
	unsupported_relation_calc_err
	unsupported_unary_calc
	unsupported_numeric_calc
	numeric_calc_err
	unary_calc_err
	relation_calc_error
	logic_calc_err
	visit_member_err
	visit_map_err
	visit_array_err
	index_err
	call_err
	parse_integer_err
	parse_decimal_err
	parse_covered_err
	check_parameter_err
	check_result_err
	parse_mapper_err
	register_fragment_err
	parse_fragment_err
	validate_xml_node_err
	paras_fragment_err
	parser_bind_err
	caller_err
	syntax_err
	result_attribute_conflict_err
	cast_bool_err
	parse_inserter_err
	parse_query_err
	tag_not_match_err
)

func throw(file string, ctx antlr.ParserRuleContext, code int) *_error {
	return &_error{file: file, ctx: ctx, code: code}
}

type _error struct {
	code    int
	file    string
	parent  antlr.ParserRuleContext
	ctx     antlr.ParserRuleContext
	line    string
	message string
}

func (p *_error) setParent(ctx antlr.ParserRuleContext) *_error {
	p.parent = ctx
	return p
}

func (p *_error) format(format string, args ...interface{}) {
	p.message = fmt.Sprintf(format, args...)
	panic(p)
}

func (p *_error) with(err error) {
	p.message = err.Error()
	panic(p)
}

func (p *_error) setLine(line string) *_error {
	p.line = line
	return p
}

func (p *_error) Error() string {
	msg := fmt.Sprintf("[ERROR %d]: %s", p.code, p.message)
	line := 0
	column := 0
	ctx := ""
	if p.parent != nil {
		line = p.parent.GetStart().GetLine()
		column = p.parent.GetStart().GetColumn()
		ctx = getText(p.parent)
	} else if p.ctx != nil {
		line = p.ctx.GetStart().GetLine()
		column = p.ctx.GetStart().GetColumn()
		ctx = getText(p.ctx)
	}
	if p.ctx != nil {
		msg += fmt.Sprintf("\n[file]: %s near line %d column %d:\n[context]: %s", p.file, line, column+1, ctx)
	} else {
		msg += fmt.Sprintf("\n[file]: %s near line %d column %d:\n[location]: %s", p.file, line, column+1, p.line)
	}
	
	return msg
}

func catch(file string, e interface{}) error {
	if e != nil {
		//debug.PrintStack()
		_e, ok := e.(*_error)
		if ok {
			if _e.file == "" {
				_e.file = file
			}
			return _e
		}
		return &_error{
			code:    unknown_err,
			file:    file,
			message: fmt.Sprintf("%v", e),
		}
	}
	return nil
}

func getText(ctx antlr.ParserRuleContext) string {
	
	if ctx.GetChildCount() == 0 {
		return ""
	}
	
	var s string
	for _, child := range ctx.GetChildren() {
		_, ok := child.(*xml.AttributeContext)
		if ok {
			s += " " + child.(antlr.ParseTree).GetText()
		} else {
			s += child.(antlr.ParseTree).GetText()
		}
	}
	
	return s
}

func newParserErrorStrategy() *parserErrorStrategy {
	return &parserErrorStrategy{BailErrorStrategy: antlr.NewBailErrorStrategy()}
}

type parserErrorStrategy struct {
	*antlr.BailErrorStrategy
}

// Recover 直接抛出语法异常
func (p *parserErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	// TODO handle syntax error detail
	//context := recognizer.GetParserRuleContext()
	throw("", nil, syntax_err).
		setLine(getLine(e.GetInputStream().(antlr.CharStream))).
		format("express syntax error: %s", e.GetMessage())
}

// RecoverInline 确保不会试图执行行内恢复
func (p *parserErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	//p.Recover(recognizer, antlr.NewBaseRecognitionException("", recognizer, recognizer.GetInputStream(), recognizer.GetParserRuleContext()))
	p.BailErrorStrategy.RecoverInline(recognizer)
	return nil
}

// Sync 确保不会试图从子规则中恢复
func (p *parserErrorStrategy) Sync(recognizer antlr.Parser) {
	// pass
}

func (p *parserErrorStrategy) ReportError(parser antlr.Parser, e antlr.RecognitionException) {
	// pass
	p.BailErrorStrategy.ReportError(parser, e)
}

func (p *parserErrorStrategy) ReportMatch(parser antlr.Parser) {
	// pass
	p.BailErrorStrategy.ReportMatch(parser)
}

func newXmlErrorStrategy(file string) *xmlErrorStrategy {
	return &xmlErrorStrategy{
		file:              file,
		BailErrorStrategy: antlr.NewBailErrorStrategy(),
	}
}

type xmlErrorStrategy struct {
	file string
	*antlr.BailErrorStrategy
}

func (p *xmlErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	var cs antlr.CharStream
	switch e.GetInputStream().(type) {
	case antlr.TokenStream:
		cs = e.GetInputStream().(antlr.TokenStream).GetTokenSource().GetInputStream()
	case antlr.CharStream:
		cs = e.GetInputStream().(antlr.CharStream)
	default:
		// TODO 处理 unknown stream
	}
	fmt.Println(reflect.TypeOf(e).String())
	//fmt.Println(recognizer.GetCurrentToken().GetLine(), recognizer.GetCurrentToken().GetColumn())
	throw(p.file, nil, syntax_err).
		setLine(getLine(cs)).
		format("xml syntax error")
	
}

func (p *xmlErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	p.Recover(recognizer, antlr.NewInputMisMatchException(recognizer))
	return nil
}

func (p *xmlErrorStrategy) Sync(recognizer antlr.Parser) {
	//fmt.Println("func (p *xmlErrorStrategy) Sync:", recognizer.GetParserRuleContext().GetText())
	// pass
}

func (p *xmlErrorStrategy) ReportError(parser antlr.Parser, err antlr.RecognitionException) {
	fmt.Println("*xmlErrorStrategy.ReportError", parser.GetParserRuleContext().GetText(), err)
	// pass
}

func (p *xmlErrorStrategy) ReportMatch(parser antlr.Parser) {
	// pass
	//fmt.Println("func (p *xmlErrorStrategy) ReportMatch", parser.GetParserRuleContext().GetText())
}

func getLine(s antlr.CharStream) string {
	start := s.Index()
	for i := s.Index() - 1; i >= 0; i-- {
		if s.GetText(i, i) == "\n" {
			break
		}
		start = i
	}
	end := s.Index()
	for i := s.Index(); i < s.Size(); i++ {
		if s.GetText(i, i) == "\n" {
			break
		}
		end = i
	}
	return s.GetText(start, end)
}

//func newLexerErrorListener() *lexerErrorListener {
//	return new(lexerErrorListener)
//}
//
//type lexerErrorListener struct {
//}
//
//func (p *lexerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//	throw("", nil, syntax_err).
//		setLine(getLine(e.GetInputStream().(antlr.CharStream))).
//		format("lexer syntax error")
//}
//

//
//func (p *lexerErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//	// pass
//}
//
//func (p *lexerErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//	// pass
//}
//
//func (p *lexerErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//	// pass
//}
