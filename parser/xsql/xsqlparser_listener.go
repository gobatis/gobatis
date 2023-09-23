// Code generated from XSQLParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package xsql // XSQLParser
import "github.com/antlr4-go/antlr/v4"

// XSQLParserListener is a complete listener for a parse tree produced by XSQLParser.
type XSQLParserListener interface {
	antlr.ParseTreeListener

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterChardata is called when entering the chardata production.
	EnterChardata(c *ChardataContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitChardata is called when exiting the chardata production.
	ExitChardata(c *ChardataContext)
}
