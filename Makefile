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
	goimports -local github.com/neotoolkit/dummy/ -w -l ./

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...
cover:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out

.PHONY: local
local:
	go run ./cmd/dummy -port=8080 server examples/docker/openapi3.yml
