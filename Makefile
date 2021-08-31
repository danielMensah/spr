.PHONY: build clean all

all: clean build

go_sources := $(shell find . -name "*.go")

bin/cli: $(go_sources)
	go build -o ./bin/cli ./cmd/main.go

build: bin/cli

clean:
	go clean ./...
	rm -rf ./bin