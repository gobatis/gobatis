package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	UnknownErr = iota
	ParameterTypeErr
)

func NewError(code int, file string, ctx antlr.ParserRuleContext, detail error) *Error {
	return &Error{code: code, file: file, ctx: ctx, detail: detail}
}

type Error struct {
	code   int
	file   string
	ctx    antlr.ParserRuleContext
	detail error
}

func (p Error) Error() string {
	msg := fmt.Sprintf("ERROR %d: %s", p.code, p.detail)
	if p.ctx != nil {
		msg += fmt.Sprintf("\n%s line %d column %d:\n%s",
			p.file, p.ctx.GetStart().GetLine(), p.ctx.GetStart().GetColumn(), p.ctx.GetText())
	}
	return msg
}

func parseError(file string, ctx antlr.ParserRuleContext, msg string) error {
	
	if ctx == nil {
		return fmt.Errorf("%s: parse error: %s", file, msg)
	}
	
	return fmt.Errorf("%s line %d:%d parse error:\n%s\n%s",
		file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText(), msg)
}
