grammar JsonPath;

/*

 Limited set of JsonPath notation

 TODO
 - bracketnotation
 - Check against spec description (http://goessner.net/articles/JsonPath/), e.g. '..' for recursive descent

*/

jsonpath: dotnotation;

dotnotation : '$.' dotnotation_expr ('.' dotnotation_expr)*;

dotnotation_expr : identifierWithQualifier
                 | INDENTIFIER
                 ;

identifierWithQualifier : INDENTIFIER '[]'
                        | INDENTIFIER '[' INT ']'
                        | INDENTIFIER '[?(' query_expr ')]'
                        ;

query_expr : query_expr ('&&' query_expr)+
           | query_expr ('||' query_expr)+
           | '*'
           | '@.' INDENTIFIER
           | '@.' INDENTIFIER '>' INT
           | '@.' INDENTIFIER '<' INT
           | '@.length-' INT
           | '@.' INDENTIFIER '==' INT
           | '@.' INDENTIFIER '==\'' INDENTIFIER '\''
           ;

INDENTIFIER : [a-zA-Z][a-zA-Z0-9]* ;
INT         : '0' | [1-9][0-9]* ;
WS  :   [ \t\n\r]+ -> skip ;