.PHONY: all tidy fmt vet lint sec test check build clean

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "v0.0.0-dev")
COMMIT := $(shell git rev-parse HEAD 2>/dev/null || echo "unknown")
BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown")

LDFLAGS := -ldflags "-X gotcode.org/sourcevault/internal/system.AppVersion=$(VERSION) -X gotcode.org/sourcevault/internal/system.Commit=$(COMMIT) -X gotcode.org/sourcevault/internal/system.Branch=$(BRANCH)"

all: check build

tidy:
	go mod tidy
	go mod download

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...

sec:
	go run github.com/securego/gosec/v2/cmd/gosec@latest ./...

test:
	go test -v -race ./...

check: tidy fmt vet lint sec test

build:
	go build $(LDFLAGS) -o bin/sourcevault ./cmd/sourcevault
	go build $(LDFLAGS) -o bin/sourcevaultd ./cmd/sourcevaultd

clean:
	rm -rf bin/
