# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOBIN=$(shell pwd)/build

PROJECT_NAME=okex
BINARY_NAME=$(PROJECT_NAME)

.PHONY: all okex clean help

all: okex

okex: ## Build the binary file
	@go mod download
	@go mod tidy
	@go build -o $(GOBIN)/$(BINARY_NAME)  ./cmd/okex/*.go
	@echo "Done building."
	@echo "Go to build folder and run \"$(GOBIN)/$(BINARY_NAME)\" to launch."
.PHONY: okex

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) -v   ./cmd/okex/main.go
.PHONY: build_linux


install: ## Install binary to system
	sudo install -C $(GOBIN)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

clean: ## Remove previous build
	@go clean
	@rm -rf $(GOBIN)
	@echo "Done cleaning."
.PHONY: clean

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: clean
