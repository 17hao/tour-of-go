// antlr4 Calc.g4 -Dlanguage=Go -o ../antlr/calc/parser -no-listener -visitor
grammar Calc;

// the start rule
prog: stat+                 # calc
    ;

stat: expr NEWLINE          # printExpr
    | ID '=' expr NEWLINE   # assign
    | NEWLINE               # blank
    ;

expr: expr op=('*'|'/') expr   # MulDiv
    | expr op=('+'|'-') expr   # AddSub
    | INT                   # int
    | ID                    # id
    | '(' expr ')'          # parens
    ;

ID : [a-zA-A]+ ;
INT : [0-9]+ ;
NEWLINE : '\r'? '\n' ; // return newlines to parser (is end-statement signal)
WS : [ \t]+ -> skip ; // toss out whitespace

MUL : '*' ; // assign token name to '*' used above in grammar
DIV : '/' ;
ADD : '+' ;
SUB : '-' ;