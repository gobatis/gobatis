grammar XSQL;

BLOCK_COMMENT     :   '<!--' .*? '-->'-> skip;
LINE_COMMENT      :   '//' ~[\r\n]* -> skip;
EntityRef         :   '&' NAME ';' ;
WS                :   (' '|'\t'|'\r'? '\n')+ ;
NAME              :   NameChar Char* ;
DOLLAR            :   '$';
HASH              :   '#';
OPEN_CURLY_BRAXE  :   '{';
CLOSE_CURLY_BRAXE :   '}';
OPEN              :   '<';
CLOSE             :   '>';
//SLASH_CLOSE       :   '/>';
SLASH             :   '/';
EQUALS            :   '=';
STRING            :   '"' ~["]* '"'
                  |   '\'' ~[']* '\''
                  ;
TEXT              :    .+?;
//TEXT: ~[<>&]+?;
//TEXT              :    ~[<]+?;



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

content      :  (element | expr | reference | chardata)*;

element      :   OPEN NAME WS* attribute* CLOSE content OPEN SLASH NAME CLOSE
             |   OPEN NAME WS* attribute* SLASH CLOSE
             ;

attribute    : NAME '=' STRING WS* ;

expr         : ((DOLLAR OPEN_CURLY_BRAXE) | (HASH OPEN_CURLY_BRAXE)) val = .*? CLOSE_CURLY_BRAXE;

reference    : EntityRef;

chardata     : WS
             | NAME
             | OPEN
             | CLOSE
             | SLASH ~'>'
             | EQUALS
             | DOLLAR
             | HASH
             | OPEN_CURLY_BRAXE
             | CLOSE_CURLY_BRAXE
             | STRING
             | TEXT
             ;

//data:    SLASH ~'>'
//              | '<' ~NameChar
//              | ~[<>/]+;