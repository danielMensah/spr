.PHONY: build clean all

all: clean test build

go_sources := $(shell find . -name "*.go")

bin/cli: $(go_sources)
	go build -o ./bin/cli ./cmd/main.go

build: bin/cli

clean:
	go clean ./...
	rm -rf ./bin

test.fmt:
	go fmt ./...;

test.vet:
	go vet ./...;

test.lint:
	golangci-lint run ./...;

test.testfmt:
	test -z $(gofmt -s -l -w .);

test.tests:
	go test -cover -race -coverprofile=c.out ./...;

test.coverage: test.tests
	go tool cover -html=c.out -o coverage.html;

test: test.fmt test.vet test.lint test.testfmt test.tests test.coverage