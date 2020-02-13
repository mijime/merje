install: test
	go install -race -gcflags '-m -m' ./... 2>&1 | tee gcflags.out | sed -e 's/.go:[0-9]*:[0-9]*:/.go/' > gcflags-0.out

test: lint cover.out

cover.out:
	go test -race -cover -coverprofile=$@ ./...

lint:
	go vet ./...
	golangci-lint run --enable-all ./...

update:
	go mod tidy
	go get -u -v ./...
