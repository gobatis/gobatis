parser grammar ExprParser;

options {
    tokenVocab=ExprLexer;
}

parameters: (STAR | paramDecl (paramComma paramDecl)*) EOF;

paramComma: COMMA;

paramDecl: IDENTIFIER | (IDENTIFIER COLON paramType);

paramType: (L_BRACKET R_BRACKET)? IDENTIFIER;

expressions: expression EOF;

misc: WS TERMINATOR;

expression:
	primaryExpr
	| unary_op = (
	    PLUS
        | MINUS
        | EXCLAMATION
        | CARET
	) expression
	| expression mul_op = (
         STAR
        | DIV
        | MOD
        | LSHIFT
        | RSHIFT
        | AMPERSAND
        | BIT_CLEAR
	) expression
	| expression add_op = (PLUS | MINUS | OR | CARET) expression
	| expression rel_op = (
	    EQUALS
        | NOT_EQUALS
        | LESS
        | LESS_OR_EQUALS
        | GREATER
        | GREATER_OR_EQUALS
	)  expression
	| expression logical expression
	| expression tertiary = QUESTION expression COLON expression;

primaryExpr:
    operand
	| primaryExpr (
		member
		| index
		| slice_
		| call
	);

logical: LOGICAL_AND | LOGICAL_OR | LOGICAL_AND_LOWER | LOGICAL_OR_LOWER | LOGICAL_AND_UPPER | LOGICAL_OR_UPPER;

operand: literal | var_ | L_PAREN expression R_PAREN;

var_: IDENTIFIER;

member: DOT IDENTIFIER;

literal: basicLit;

basicLit:
	nil_
	| integer
	| string_
	| float_
	| IMAGINARY_LIT
	| RUNE_LIT;

integer:
	DECIMAL_LIT
	| BINARY_LIT
	| OCTAL_LIT
	| HEX_LIT
	| IMAGINARY_LIT
	| RUNE_LIT;

nil_: NIL_LIT;

string_: RAW_STRING_LIT | INTERPRETED_STRING_LIT;

float_: FLOAT_LIT;

index: L_BRACKET expression R_BRACKET;

expressionList: expression (COMMA expression)*;

call:
	 L_PAREN (
        expressionList ELLIPSIS?
	)? R_PAREN;

slice_:
	L_BRACKET (
		expression? COLON expression?
		| expression? COLON expression COLON expression
	) R_BRACKET;
