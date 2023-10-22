parser grammar ExprParser;

options {
    tokenVocab=ExprLexer;
}

expressions: expression EOF;

expression:
	primary
	| unary = (
        PLUS
        | MINUS
        | EXCLAMATION
        | CARET
    ) expression
	| expression rel = (
	    EQUALS
        | NOT_EQUALS
        | LESS
        | LESS_OR_EQUALS
        | GREATER
        | GREATER_OR_EQUALS
	)  expression
	| expression logical expression
	| expression tertiary = QUESTION expression COLON expression;

primary:
    operand
	| primary (
		member
		| index
		| slice
		| call
	);

logical: LOGICAL_AND | LOGICAL_OR;

operand: literal | var | L_PAREN expression R_PAREN;

var: IDENTIFIER;

member: DOT IDENTIFIER;

literal:
	nil
	| integer
	| string
	| float
//	| IMAGINARY_LIT
//	| RUNE_LIT
	;

integer:
	DECIMAL_LIT
//	| BINARY_LIT
//	| OCTAL_LIT
//	| HEX_LIT
//	| IMAGINARY_LIT
//	| RUNE_LIT
    ;

nil: NIL_LIT;

string: RAW_STRING_LIT | INTERPRETED_STRING_LIT;

float: FLOAT_LIT;

index: L_BRACKET expression R_BRACKET;

expressionList: expression (COMMA expression)*;

call:
	 L_PAREN (
        expressionList ELLIPSIS?
	)? R_PAREN;

slice:
	L_BRACKET (
		sea=expression? COLON seb=expression?
		| sea=expression? COLON seb=expression COLON sec=expression
	) R_BRACKET;
