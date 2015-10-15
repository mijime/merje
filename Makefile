GO_FILES = $(wildcard *.go */*.go */*/*.go)

all: test

format:
	gofmt -d -s -w -e $(GO_FILES)

test: format
	go test

install: test
	go get -v
