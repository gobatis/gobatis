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

identifierWithQualifier : INDENTIFIER '[*]';

INDENTIFIER : [a-zA-Z][a-zA-Z0-9]* ;
WS  :   [ \t\n\r]+ -> skip ;