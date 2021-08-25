// Code generated from ExprParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package expr // ExprParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExprParserListener is a complete listener for a parse tree produced by ExprParser.
type ExprParserListener interface {
	antlr.ParseTreeListener

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParamComma is called when entering the paramComma production.
	EnterParamComma(c *ParamCommaContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterParamType is called when entering the paramType production.
	EnterParamType(c *ParamTypeContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterMisc is called when entering the misc production.
	EnterMisc(c *MiscContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterLogical is called when entering the logical production.
	EnterLogical(c *LogicalContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterVar_ is called when entering the var_ production.
	EnterVar_(c *Var_Context)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBasicLit is called when entering the basicLit production.
	EnterBasicLit(c *BasicLitContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterNil_ is called when entering the nil_ production.
	EnterNil_(c *Nil_Context)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterFloat_ is called when entering the float_ production.
	EnterFloat_(c *Float_Context)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterSlice_ is called when entering the slice_ production.
	EnterSlice_(c *Slice_Context)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParamComma is called when exiting the paramComma production.
	ExitParamComma(c *ParamCommaContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitParamType is called when exiting the paramType production.
	ExitParamType(c *ParamTypeContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitMisc is called when exiting the misc production.
	ExitMisc(c *MiscContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitLogical is called when exiting the logical production.
	ExitLogical(c *LogicalContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitVar_ is called when exiting the var_ production.
	ExitVar_(c *Var_Context)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBasicLit is called when exiting the basicLit production.
	ExitBasicLit(c *BasicLitContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitNil_ is called when exiting the nil_ production.
	ExitNil_(c *Nil_Context)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitFloat_ is called when exiting the float_ production.
	ExitFloat_(c *Float_Context)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitSlice_ is called when exiting the slice_ production.
	ExitSlice_(c *Slice_Context)
}
