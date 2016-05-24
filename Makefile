VERSION = dev
NAME = $(shell pwd | xargs basename)

SRC_FILES = $(shell find . -name "*.go" -type f)
RELEASE_DIR = _obj/$(VERSION)

TARGET = \
				 release-windows-amd64 \
				 release-linux-amd64 \
				 release-darwin-amd64 \
				 release-windows-386 \
				 release-linux-386 \
				 release-darwin-386

all: $(TARGET)

ghr: all
	ghr --replace $(VERSION) _obj/$(VERSION)

test: $(SRC_FILES)
	go fmt ./...
	go vet ./...
	go test -v -race -cover ./...

release-windows-%:
		@$(MAKE) release GOOS=windows GOARCH=$* SUFFIX=.exe

release-%-amd64:
		@$(MAKE) release GOOS=$* GOARCH=amd64

release-%-386:
		@$(MAKE) release GOOS=$* GOARCH=386

release: $(RELEASE_DIR)/$(NAME)_$(GOOS)_$(GOARCH).tar.gz

$(RELEASE_DIR)/$(NAME)_$(GOOS)_$(GOARCH).tar.gz: $(RELEASE_DIR)/$(NAME)_$(GOOS)_$(GOARCH)/$(NAME)$(SUFFIX)
	tar cfz $@ -C $(RELEASE_DIR)/$(NAME)_$(GOOS)_$(GOARCH) $(NAME)$(SUFFIX)

build-windows-%: $(SRC_FILES)
		@$(MAKE) build GOOS=windows GOARCH=$* SUFFIX=.exe

build-%-amd64: $(SRC_FILES)
		@$(MAKE) build GOOS=$* GOARCH=amd64

build-%-386: $(SRC_FILES)
		@$(MAKE) build GOOS=$* GOARCH=386

build: $(RELEASE_DIR)/$(NAME)_$(GOOS)_$(GOARCH)/$(NAME)$(SUFFIX)

$(RELEASE_DIR)/$(NAME)_$(GOOS)_$(GOARCH)/$(NAME)$(SUFFIX):
	go build -ldflags "-X main.Version=$(VERSION)" -o $@
