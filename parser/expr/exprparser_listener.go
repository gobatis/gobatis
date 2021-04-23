// Code generated from ExprParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package expr // ExprParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExprParserListener is a complete listener for a parse tree produced by ExprParser.
type ExprParserListener interface {
	antlr.ParseTreeListener

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterParamType is called when entering the paramType production.
	EnterParamType(c *ParamTypeContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterConversion is called when entering the conversion production.
	EnterConversion(c *ConversionContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterOperandName is called when entering the operandName production.
	EnterOperandName(c *OperandNameContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBasicLit is called when entering the basicLit production.
	EnterBasicLit(c *BasicLitContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterFloat_ is called when entering the float_ production.
	EnterFloat_(c *Float_Context)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterQualifiedIdent is called when entering the qualifiedIdent production.
	EnterQualifiedIdent(c *QualifiedIdentContext)

	// EnterSlice is called when entering the slice production.
	EnterSlice(c *SliceContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitParamType is called when exiting the paramType production.
	ExitParamType(c *ParamTypeContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitConversion is called when exiting the conversion production.
	ExitConversion(c *ConversionContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitOperandName is called when exiting the operandName production.
	ExitOperandName(c *OperandNameContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBasicLit is called when exiting the basicLit production.
	ExitBasicLit(c *BasicLitContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitFloat_ is called when exiting the float_ production.
	ExitFloat_(c *Float_Context)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitQualifiedIdent is called when exiting the qualifiedIdent production.
	ExitQualifiedIdent(c *QualifiedIdentContext)

	// ExitSlice is called when exiting the slice production.
	ExitSlice(c *SliceContext)
}
