lexer grammar ExprLexer;

NIL_LIT                : 'nil';

// Logical
LOGICAL_OR             : '||';
LOGICAL_AND            : '&&';

// Identifier
IDENTIFIER             : LETTER (LETTER | UNICODE_DIGIT)*;

// Punctuation
L_PAREN                : '(';
R_PAREN                : ')';
L_BRACKET              : '[';
R_BRACKET              : ']';
ASSIGN                 : '=';
COMMA                  : ',';
SEMI                   : ';';
COLON                  : ':';
QUESTION               : '?';
DOT                    : '.';
//PLUS_PLUS              : '++';
//MINUS_MINUS            : '--';
//ELLIPSIS               : '...';

// Relation operators
EQUALS                 : '==';
NOT_EQUALS             : '!=';
LESS                   : '<';
LESS_OR_EQUALS         : '<=';
GREATER                : '>';
GREATER_OR_EQUALS      : '>=';

// Arithmetic operators
OR                     : '|';
DIV                    : '/';
MOD                    : '%';
LSHIFT                 : '<<';
RSHIFT                 : '>>';
BIT_CLEAR              : '&^';

// Unary operators
EXCLAMATION            : '!';

// Mixed operators
PLUS                   : '+';
MINUS                  : '-';
CARET                  : '^';
STAR                   : '*';
AMPERSAND              : '&';

// Number literals
DECIMAL_LIT            : '0' | [1-9] ('_'? [0-9])*;
BINARY_LIT             : '0' [bB] ('_'? BIN_DIGIT)+;
OCTAL_LIT              : '0' [oO]? ('_'? OCTAL_DIGIT)+;
HEX_LIT                : '0' [xX]  ('_'? HEX_DIGIT)+;


FLOAT_LIT : DECIMAL_FLOAT_LIT | HEX_FLOAT_LIT;

DECIMAL_FLOAT_LIT      : DECIMALS ('.' DECIMALS? EXPONENT? | EXPONENT)
                       | '.' DECIMALS EXPONENT?
                       ;

HEX_FLOAT_LIT          : '0' [xX] HEX_MANTISSA HEX_EXPONENT
                       ;

fragment HEX_MANTISSA  : ('_'? HEX_DIGIT)+ ('.' ( '_'? HEX_DIGIT )*)?
                       | '.' HEX_DIGIT ('_'? HEX_DIGIT)*;

fragment HEX_EXPONENT  : [pP] [+-] DECIMALS;


IMAGINARY_LIT          : (DECIMAL_LIT | BINARY_LIT |  OCTAL_LIT | HEX_LIT | FLOAT_LIT) 'i';

// Rune literals
RUNE_LIT               : '\'' (UNICODE_VALUE | BYTE_VALUE) '\'';//: '\'' (~[\n\\] | ESCAPED_VALUE) '\'';

BYTE_VALUE : OCTAL_BYTE_VALUE | HEX_BYTE_VALUE;

OCTAL_BYTE_VALUE: '\\' OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT;

HEX_BYTE_VALUE: '\\' 'x'  HEX_DIGIT HEX_DIGIT;

LITTLE_U_VALUE: '\\' 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT;

BIG_U_VALUE: '\\' 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT;

// String literals
RAW_STRING_LIT         : '`' ~'`'*                      '`';
INTERPRETED_STRING_LIT : '"' (~["\\] | ESCAPED_VALUE)*  '"' | '\'' (~['\\] | ESCAPED_VALUE)*  '\'';

// Skip tokens
WS                     : [ \t]+             -> skip;
TERMINATOR             : [\r\n]+            -> skip;

fragment UNICODE_VALUE: ~[\r\n'] | LITTLE_U_VALUE | BIG_U_VALUE | ESCAPED_VALUE;

// Fragments
fragment ESCAPED_VALUE
    : '\\' ('u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | [abfnrtv\\'"]
           | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
           | 'x' HEX_DIGIT HEX_DIGIT)
    ;
fragment DECIMALS
    : [0-9] ('_'? [0-9])*
    ;
fragment OCTAL_DIGIT
    : [0-7]
    ;
fragment HEX_DIGIT
    : [0-9a-fA-F]
    ;
fragment BIN_DIGIT
    : [01]
    ;
fragment EXPONENT
    : [eE] [+-]? DECIMALS
    ;
fragment LETTER
    : UNICODE_LETTER
    | '_'
    ;
fragment UNICODE_DIGIT
    : [\p{Nd}]
    ;
fragment UNICODE_LETTER
    : [\p{L}]
    ;