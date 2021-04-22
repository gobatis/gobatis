parser grammar ParameterParser;

options {
    tokenVocab=ParameterLexer;
}

expression: varSpec (COMMA varSpec)* EOF;
varSpec: IDENTIFIER | (IDENTIFIER COLON IDENTIFIER);
