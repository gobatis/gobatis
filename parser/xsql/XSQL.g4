grammar XSQL;

BLOCK_COMMENT     :   '<!--' .*? '-->'-> skip;
LINE_COMMENT      :   '//' ~[\r\n\t]* -> skip;
EntityRef         :   '&' NAME ';' ;
WS                :   (' '|'\t'|'\r'? '\n')+ ;

NAME              :   NameChar Char* ;
DOLLAR            :   '$';
HASH              :   '#';
OPEN_CURLY_BRAXE  :   '{';
CLOSE_CURLY_BRAXE :   '}';
OPEN              :   '<';
CLOSE             :   '>';
SLASH             :   '/';
EQUALS            :   '=';
STRING            :   '"' ~["]* '"'
                  |   '\'' ~[']* '\''
                  ;
TEXT              :    .+?;
fragment
DIGIT        :   [0-9];

fragment
Char         :   NameChar
             |   '-' | '_' | '.' | DIGIT
             |   '\u00B7'
             |   '\u0300'..'\u036F'
             |   '\u203F'..'\u2040'
             ;

fragment
NameChar     :   [:a-zA-Z]
             |   '\u2070'..'\u218F'
             |   '\u2C00'..'\u2FEF'
             |   '\u3001'..'\uD7FF'
             |   '\uF900'..'\uFDCF'
             |   '\uFDF0'..'\uFFFD'
             ;

content      :  (start | end | expr | reference | chardata)* EOF;


start        :  '<' NAME WS* attribute* '>'
             |  '<' NAME WS* attribute* '/' '>'
             ;
end   :  '<' '/' NAME '>';

attribute    :  NAME '=' STRING WS* ;

expr         :  (('$' '{') | ('#' '{')) WS* val=NAME* TEXT*  WS* '}';

reference    :  EntityRef;

chardata     : WS
             | '>'
             | '/'
             | '<'
             | '='
             | '$'
             | '#'
             | '{'
             | '}'
             | NAME
             | STRING
             | TEXT
             ;
