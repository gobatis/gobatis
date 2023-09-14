GO=go
GOCOVER=$(GO) tool cover
.PHONY: test
test:
	go test -coverprofile=test/coverage.out ./...
	go tool cover -html=test/coverage.out -o test/coverage.html
	go test -cover -json  ./... | go-test-html-report -o test
