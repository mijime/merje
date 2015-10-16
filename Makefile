NAME = merje
VERSION = 0.1.2

GO_FILES = $(shell find -name "*.go" -type f)
GO_TEST_FILES = $(shell find -name "*_test.go" -type f)

BUILD_DIR = pkg/$(VERSION)
DIST_DIR = pkg/dist/$(VERSION)
DIST_TARS = \
	$(DIST_DIR)/windows_amd64.tar.gz \
	$(DIST_DIR)/darwin_amd64.tar.gz \
	$(DIST_DIR)/linux_amd64.tar.gz

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
	go tool vet -v .
	go test -cover ./...

install: test
	go get -v ./...

release: ghr tarball
	ghr $(VERSION) $(DIST_DIR)

tarball: $(DIST_TARS)

$(DIST_DIR)/%.tar.gz:
	mkdir -p $(DIST_DIR)
	tar cvfz $@ $(BUILD_DIR)/$(*)

build: gox format test
	gox \
		-ldflags="-X main.GitCommit \"$$(git describe --always)\"" \
		-os="$(XC_OS)" \
		-arch="$(XC_ARCH)" \
		-output="$(BUILD_DIR)/{{.OS}}_{{.Arch}}/{{.Dir}}" \
		./...

ghr:
	go get -v github.com/tcnksm/ghr

gox:
	go get -v github.com/mitchellh/gox
