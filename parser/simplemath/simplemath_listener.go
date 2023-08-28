// Code generated from SimpleMath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package simplemath // SimpleMath
import "github.com/antlr4-go/antlr/v4"

// SimpleMathListener is a complete listener for a parse tree produced by SimpleMathParser.
type SimpleMathListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)
}
