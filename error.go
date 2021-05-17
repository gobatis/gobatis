package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func parseError(file string, ctx antlr.ParserRuleContext, msg string) error {
	return fmt.Errorf("%s line %d:%d:\n%s\nparse error: %s",
		file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText(), msg)
}
