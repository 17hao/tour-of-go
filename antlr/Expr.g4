//  antlr4 Expr.g4 -Dlanguage=Go -o ../antlr/expr/parser
grammar Expr;

// the start rule
prog: stat+ ;

stat: expr NEWLINE
    | ID '=' expr NEWLINE
    | NEWLINE
    ;

expr: expr ('*'|'/') expr
    | expr ('+'|'-') expr
    | INT
    | ID
    | '(' expr ')'
    ;

ID : [a-zA-A]+ ;
INT : [0-9]+ ;
NEWLINE : '\r'? '\n' ; // return newlines to parser (is end-statement signal)
WS : [ \t]+ -> skip ; // toss out whitespace