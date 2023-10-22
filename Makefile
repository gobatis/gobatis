.PHONY: test
test:
	gocov test -coverpkg=./... ./... | gocov-html > test/report.html

.PHONY: xsql
xsql:
	cd g4/xsql && antlr4 -Dlanguage=Go -o ./ -package xsql -no-listener -no-visitor -Werror XSQL*.g4
	mv ./g4/xsql/*.go ./parser/xsql/



.PHONY: expr
expr:
	cd g4/expr && antlr4 -Dlanguage=Go -o ./ -package expr -no-listener -no-visitor -Werror Expr*.g4
	mv ./g4/expr/*.go ./parser/expr/