alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -o ./ -package xml -no-visitor -no-listener -Werror  XML*.g4
