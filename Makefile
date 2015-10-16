NAME = merje
VERSION = 0.1.1

GO_FILES = $(shell find -name "*.go" -type f)
GO_TEST_FILES = $(shell find -name "*_test.go" -type f)
BUILD_DIRS = \
	pkg/dist/$(VERSION)/windows_amd64.tar.gz \
	pkg/dist/$(VERSION)/darwin_amd64.tar.gz \
	pkg/dist/$(VERSION)/linux_amd64.tar.gz

# XC_ARCH = 386 amd64
XC_ARCH = amd64
XC_OS = windows darwin linux

all: $(GO_FILES)

%.go:
	gofmt -d -s -w -e $*.go

%_test.go:
	go test -v ./$(*D)

format: $(GO_FILES)
	gofmt -d -s -w -e $(GO_FILES)

test: format
	go test ./...

install: test
	go get -v ./...

release: ghr tarball
	ghr $(VERSION) pkg/dist/$(VERSION)

tarball: $(BUILD_DIRS)

pkg/dist/$(VERSION)/%.tar.gz:
	mkdir -p pkg/dist/$(VERSION)
	tar cvfz $@ pkg/$(VERSION)/$(*)

build: gox format test
	gox \
		-ldflags="-X main.GitCommit \"$$(git describe --always)\"" \
		-os="$(XC_OS)" \
		-arch="$(XC_ARCH)" \
		-output="pkg/$(VERSION)/{{.OS}}_{{.Arch}}/{{.Dir}}" \
		./...

ghr:
	go get -v github.com/tcnksm/ghr

gox:
	go get -v github.com/mitchellh/gox
