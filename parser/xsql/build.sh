#alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -o ./ -package xsql -no-listener -no-visitor -Werror XSQL*.g4
#antlr4 XSQL*.g4
#javac *.java
#rm  *.java
rm  *.interp
#rm  *.tokens
#sed -i "" 's/p.lineTerminatorAhead/this.lineTerminatorAhead/g' JsonPath.java
#sed -i "" 's/p.checkPreviousTokenText/this.checkPreviousTokenText/g' JsonPath.java
#javac JsonPath*.java
