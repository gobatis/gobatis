package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/parser/xml"
	"strconv"
	"strings"
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
	expr_syntax_err
	xml_syntax_err
	result_attribute_conflict_err
	cast_bool_err
	parse_inserter_err
	parse_query_err
	tag_not_match_err
)

type ErrorMessage struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Context string `json:"context,omitempty"`
}

func (p ErrorMessage) Error() string {
	return fmt.Sprintf("%s line %d:%d [%d] %s at '%s'", p.File, p.Line, p.Column, p.Code, p.Message, p.Context)
}

func ParseErrorMessage(msg string) ErrorMessage {
	r := ErrorMessage{}
	sc := 0
	chars := ""
	for i := 0; i < len(msg); i++ {
		v := msg[i]
		if sc < 7 && v == 32 {
			sc++
		} else {
			chars += fmt.Sprintf("%c", v)
		}
		switch sc {
		case 1:
			r.File = chars
			chars = ""
			sc++
		case 3:
			chars = ""
			sc++
		case 5:
			items := strings.Split(chars, ":")
			if len(items) == 2 {
				r.Line, _ = strconv.Atoi(items[0])
				r.Column, _ = strconv.Atoi(items[1])
			}
			chars = ""
			sc++
		case 7:
			r.Code, _ = strconv.Atoi(strings.TrimLeft(strings.TrimRight(chars, "]"), "["))
			chars = ""
			sc++
		default:
			if sc > 7 && v == 32 && i+5 <= len(msg) && msg[i:i+5] == " at '" {
				r.Message = strings.TrimSuffix(chars, " ")
				r.Context = strings.TrimRight(msg[i+5:], "'")
				chars = ""
				
				break
			}
		}
	}
	if r.Message == "" {
		r.Message = chars
	}
	return r
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
		column += 1
	} else {
		ctx = p.line
		column += 1
	}
	return ErrorMessage{
		File:    p.file,
		Line:    line,
		Column:  column,
		Code:    p.code,
		Message: p.message,
		Context: ctx,
	}.Error()
}

func throw(file string, ctx antlr.ParserRuleContext, code int) *_error {
	return &_error{
		file: file,
		ctx:  ctx,
		code: code,
	}
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
	
	return strings.TrimSpace(s)
}

func newExprErrorStrategy() *exprErrorStrategy {
	return &exprErrorStrategy{BailErrorStrategy: antlr.NewBailErrorStrategy()}
}

type exprErrorStrategy struct {
	*antlr.BailErrorStrategy
}

func (p *exprErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	throw("", nil, expr_syntax_err).
		setLine(getLine(e.GetInputStream().(antlr.CharStream))).
		format("express syntax error: %s", e.GetMessage())
}

func (p *exprErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	p.Recover(recognizer, antlr.NewInputMisMatchException(recognizer))
	return nil
}

func (p *exprErrorStrategy) Sync(recognizer antlr.Parser) {
	// pass
}

func (p *exprErrorStrategy) ReportError(parser antlr.Parser, e antlr.RecognitionException) {
	// pass
}

func (p *exprErrorStrategy) ReportMatch(parser antlr.Parser) {
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
	throw(p.file, nil, xml_syntax_err).
		setLine(getLine(cs)).
		format("xml syntax error")
	
}

func (p *xmlErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	p.Recover(recognizer, antlr.NewInputMisMatchException(recognizer))
	return nil
}

func (p *xmlErrorStrategy) Sync(recognizer antlr.Parser) {
	// pass
}

func (p *xmlErrorStrategy) ReportError(parser antlr.Parser, err antlr.RecognitionException) {
	// pass
}

func (p *xmlErrorStrategy) ReportMatch(parser antlr.Parser) {
	// pass
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
	return strings.TrimSpace(s.GetText(start, end))
}
