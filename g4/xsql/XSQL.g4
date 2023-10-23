grammar XSQL;

BLOCK_COMMENT     :   ('<!--' .*? '-->' |  '/*' .*? '*/') -> skip;
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

document     : content EOF;

content      :  (element | expr | reference | chardata)*;


element     :  '<' name='if'  WS* attribute* '>' content  '</if>'
            |  '<' name='choose'  WS* attribute* '>' content  '</choose>'
            |  '<' name='when'  WS* attribute* '>' content  '</when>'
            |  '<' name='otherwise'  WS* attribute* '>' content  '</otherwise>'
            |  '<' name='trim'  WS* attribute* '>' content  '</trim>'
            |  '<' name='where'  WS* attribute* '>' content  '</where>'
            |  '<' name='set'  WS* attribute* '>' content  '</set>'
            |  '<' name='foreach'  WS* attribute* '>' content  '</foreach>'
            ;

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
             | 'if'
             | 'choose'
             | 'when'
             | 'otherwise'
             | 'trim'
             | 'where'
             | 'set'
             | 'foreach'
             | NAME
             | STRING
             | TEXT
             ;
