#alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr -Dlanguage=Go -o ./ -package xsql -no-listener -no-visitor -Werror -Xforce-atn XSQL*.g4
antlr XSQL*.g4
#javac *.java
#rm  *.java
#rm  *.interp
#rm  *.tokens
#sed -i "" 's/p.lineTerminatorAhead/this.lineTerminatorAhead/g' JsonPath.java
#sed -i "" 's/p.checkPreviousTokenText/this.checkPreviousTokenText/g' JsonPath.java
#javac JsonPath*.java
