// Code generated from ExprParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package expr // ExprParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExprParserListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprParserListener struct{}

var _ ExprParserListener = &BaseExprParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseExprParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseExprParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterParamComma is called when production paramComma is entered.
func (s *BaseExprParserListener) EnterParamComma(ctx *ParamCommaContext) {}

// ExitParamComma is called when production paramComma is exited.
func (s *BaseExprParserListener) ExitParamComma(ctx *ParamCommaContext) {}

// EnterParamDecl is called when production paramDecl is entered.
func (s *BaseExprParserListener) EnterParamDecl(ctx *ParamDeclContext) {}

// ExitParamDecl is called when production paramDecl is exited.
func (s *BaseExprParserListener) ExitParamDecl(ctx *ParamDeclContext) {}

// EnterParamType is called when production paramType is entered.
func (s *BaseExprParserListener) EnterParamType(ctx *ParamTypeContext) {}

// ExitParamType is called when production paramType is exited.
func (s *BaseExprParserListener) ExitParamType(ctx *ParamTypeContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseExprParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseExprParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterMisc is called when production misc is entered.
func (s *BaseExprParserListener) EnterMisc(ctx *MiscContext) {}

// ExitMisc is called when production misc is exited.
func (s *BaseExprParserListener) ExitMisc(ctx *MiscContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExprParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExprParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseExprParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseExprParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterLogical is called when production logical is entered.
func (s *BaseExprParserListener) EnterLogical(ctx *LogicalContext) {}

// ExitLogical is called when production logical is exited.
func (s *BaseExprParserListener) ExitLogical(ctx *LogicalContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseExprParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseExprParserListener) ExitOperand(ctx *OperandContext) {}

// EnterVar_ is called when production var_ is entered.
func (s *BaseExprParserListener) EnterVar_(ctx *Var_Context) {}

// ExitVar_ is called when production var_ is exited.
func (s *BaseExprParserListener) ExitVar_(ctx *Var_Context) {}

// EnterMember is called when production member is entered.
func (s *BaseExprParserListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseExprParserListener) ExitMember(ctx *MemberContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseExprParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseExprParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BaseExprParserListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BaseExprParserListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseExprParserListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseExprParserListener) ExitInteger(ctx *IntegerContext) {}

// EnterNil_ is called when production nil_ is entered.
func (s *BaseExprParserListener) EnterNil_(ctx *Nil_Context) {}

// ExitNil_ is called when production nil_ is exited.
func (s *BaseExprParserListener) ExitNil_(ctx *Nil_Context) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseExprParserListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseExprParserListener) ExitString_(ctx *String_Context) {}

// EnterFloat_ is called when production float_ is entered.
func (s *BaseExprParserListener) EnterFloat_(ctx *Float_Context) {}

// ExitFloat_ is called when production float_ is exited.
func (s *BaseExprParserListener) ExitFloat_(ctx *Float_Context) {}

// EnterIndex is called when production index is entered.
func (s *BaseExprParserListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseExprParserListener) ExitIndex(ctx *IndexContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseExprParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseExprParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterCall is called when production call is entered.
func (s *BaseExprParserListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseExprParserListener) ExitCall(ctx *CallContext) {}

// EnterSlice_ is called when production slice_ is entered.
func (s *BaseExprParserListener) EnterSlice_(ctx *Slice_Context) {}

// ExitSlice_ is called when production slice_ is exited.
func (s *BaseExprParserListener) ExitSlice_(ctx *Slice_Context) {}
