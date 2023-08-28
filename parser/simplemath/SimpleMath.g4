grammar SimpleMath;

expr:   expr '+' term
    |   term
    ;

term:   term '*' factor
    |   factor
    ;

factor: NUMBER
    ;

NUMBER: [0-9]+;
WS: [ \t\r\n]+ -> skip;