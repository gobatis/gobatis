// Code generated from JsonPath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package structpath // JsonPath
import "github.com/antlr4-go/antlr/v4"

// JsonPathListener is a complete listener for a parse tree produced by JsonPathParser.
type JsonPathListener interface {
	antlr.ParseTreeListener

	// EnterJsonpath is called when entering the jsonpath production.
	EnterJsonpath(c *JsonpathContext)

	// EnterDotnotation is called when entering the dotnotation production.
	EnterDotnotation(c *DotnotationContext)

	// EnterDotnotation_expr is called when entering the dotnotation_expr production.
	EnterDotnotation_expr(c *Dotnotation_exprContext)

	// EnterIdentifierWithQualifier is called when entering the identifierWithQualifier production.
	EnterIdentifierWithQualifier(c *IdentifierWithQualifierContext)

	// ExitJsonpath is called when exiting the jsonpath production.
	ExitJsonpath(c *JsonpathContext)

	// ExitDotnotation is called when exiting the dotnotation production.
	ExitDotnotation(c *DotnotationContext)

	// ExitDotnotation_expr is called when exiting the dotnotation_expr production.
	ExitDotnotation_expr(c *Dotnotation_exprContext)

	// ExitIdentifierWithQualifier is called when exiting the identifierWithQualifier production.
	ExitIdentifierWithQualifier(c *IdentifierWithQualifierContext)
}
