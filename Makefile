.PHONY: test
test:
	gocov test -coverpkg=./... ./... | gocov-html > test/report.html
