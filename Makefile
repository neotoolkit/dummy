.PHONY: all tidy fmt imports lint test

.PHONY: all
all: tidy fmt imports lint test

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: imports
imports:
	goimports -local github.com/go-dummy/dummy/ -w -l ./

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...
