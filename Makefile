COMMIT_HASH=$(shell git rev-parse --verify HEAD | cut -c 1-8)
BUILD_DATE=$(shell date +%Y-%m-%d_%H:%M:%S%z)
GIT_TAG=$(shell git describe --tags)
GIT_AUTHOR=$(shell git show -s --format=%an)
SHELL:=/bin/bash
BIN_NAME="mock"
path=$(shell pwd)


all: build test # golint

.PHONY: build
build: mod
	go build -ldflags "-X main.GitTag=$(GIT_TAG) -X main.BuildTime=$(BUILD_DATE) -X main.GitCommit=$(COMMIT_HASH) -X main.GitAuthor=$(GIT_AUTHOR)"  -o ${BIN_NAME} ./main/main.go
	mkdir -p output/bin
	cp -r ${BIN_NAME} output/bin

.PHONY: cover
cover: mod
	@echo "build cover test"
	go test -c -covermode=count -coverpkg=gitlab.mobvista.com/voyager/pioneer/internal/... -o ${BIN_NAME}_cover  ./cmd/main_test.go

mod: golang
	go mod download && go mod tidy

.PHONY: test
test:
	@echo "Run unit tests"
	go test -test.short -cover -gcflags=-l ./...

.PHONY: golang
golang:
	@hash go 2>/dev/null || (go version | grep "$NEED_GO_VERSION" > /dev/null) || { \
		echo "install go1.14.13" && \
		mkdir -p ${THIRD_DIR} && cd ${THIRD_DIR} && \
		wget https://dl.google.com/go/go1.14.13.linux-amd64.tar.gz && \
		tar -xzvf go1.14.13.linux-amd64.tar.gz && \
		cd .. && \
		export PATH="${THIRD_DIR}/go/bin/" && \
		go version; \
	}


golint: golang
	@echo "Run golangci-lint linters"
	golangci-lint run cmd/... internal/... -v

clean:
	rm -rf output
