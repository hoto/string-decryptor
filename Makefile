all: clean build test run install goreleaser-release goreleaser-dry-run goreleaser-dry-run

REPO_NAME = github.com/hoto/string-decryptor
GO_RELEASER_VERSION = v1.13.1

clean:
	go clean
	rm -rf bin/ dist/

dependencies:
	go mod download
	go mod tidy
	go mod verify

build: dependencies
	go build -ldflags="-X '${REPO_NAME}/config.Version=0.0.0' -X '${REPO_NAME}/config.ShortCommit=HASH' -X '${REPO_NAME}/config.BuildDate=DATE'" -o bin/string-decryptor

test:
	go test -v ./...

run: clean build
	./bin/string-decryptor $(arg)

install: clean build
	go install -v ./...

smoke-test: clean build
	./bin/string-decryptor -a aes256 -p password -s salt -b thebody

goreleaser-release: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=${GO_RELEASER_VERSION} bash

goreleaser-dry-run: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=${GO_RELEASER_VERSION} bash -s -- --skip-publish --snapshot --rm-dist

goreleaser-dry-run-local: dependencies
	goreleaser release --skip-publish --snapshot --rm-dist

