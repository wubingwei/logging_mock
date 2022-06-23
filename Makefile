COMMIT_HASH=$(shell git rev-parse --verify HEAD | cut -c 1-8)
BUILD_DATE=$(shell date +%Y-%m-%d_%H:%M:%S%z)
GIT_TAG=$(shell git describe --tags)
GIT_AUTHOR=$(shell git show -s --format=%an)
SHELL:=/bin/bash
BIN_NAME="mock_server"
path=$(shell pwd)


all: build test # golint

.PHONY: build
build: mod
	go build -ldflags "-X main.BuildTime=$(BUILD_DATE) -X main.GitCommit=$(COMMIT_HASH) -X main.GitAuthor=$(GIT_AUTHOR)"  -o ${BIN_NAME} ./main/main.go
	mkdir -p output/bin
	mv ${BIN_NAME} output/bin

.PHONY: cover
cover: mod
	@echo "build cover test"
	go test -c -covermode=count -coverpkg=gitlab.mobvista.com/voyager/pioneer/internal/... -o ${BIN_NAME}_cover  ./cmd/main_test.go

.PHONY: mod
mod:
	go mod download && go mod tidy

.PHONY: test
test:
	@echo "Run unit tests"
	go test -test.short -cover -gcflags=-l ./...

clean:
	rm -rf output
