# disable default rules
.SUFFIXES:
MAKEFLAGS+=-r -R
GOBIN = $(shell go env GOPATH)/bin
DATE  = $(shell date +%Y%m%d%H%M%S)

default: test

.PHONY: test
test:
	go test -race -shuffle=on -v ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: ci-tidy
ci-tidy:
	go mod tidy
	git status --porcelain go.mod go.sum || { echo "Please run 'go mod tidy'."; exit 1; }

.PHONY: staticcheck
staticcheck: $(GOBIN)/staticcheck
	$(GOBIN)/staticcheck ./...

bin/noiota.so: noiota/noiota.go
	CGO_ENABLED=1 go build -o $@ -buildmode=plugin github.com/jacoelho/noiota/plugin