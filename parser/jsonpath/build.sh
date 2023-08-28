alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -o ./ -package jsonpath -Werror  -Xforce-atn   jsonpath.g4
antlr4 jsonpath.g4
sed -i "" 's/p.lineTerminatorAhead/this.lineTerminatorAhead/g' ExprParser.java
sed -i "" 's/p.checkPreviousTokenText/this.checkPreviousTokenText/g' ExprParser.java
javac jsonpath*.java
