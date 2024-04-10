.PHONY: all build ship help coverage lint test vet build build-static install directories debug

SHELL=/bin/bash

### Env Variables
#!make
include .env
-include ${ENV_LOCAL_FILE}

### Variables
PROJECT_NAME := telemetry-dice-roller
VENDORFLAGS ?= -mod=vendor

GOFLAGS  = -v

GOPATH   = $(shell go env GOPATH)
GOJUNIT  = $(GOPATH)/bin/go-junit-report
GOCILINT = $(GOPATH)/bin/golangci-lint

MKDIR_P = mkdir -p
INSTALL = install

LDFLAGS += -s -w
LDFLAGS += -X "main.Version=${CI_COMMIT_TAG}"
LDFLAGS += -X "main.BuildTime=$(shell date -u '+%Y-%m-%dT%H:%M:%S.%3N%z')"
LDFLAGS += -X "main.CommitVersion=${CI_COMMIT_SHA}"
LDFLAGS += -X "main.PipelineNumber=${CI_PIPELINE_ID}"

SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

## Latest version of linter.  https://github.com/golangci/golangci-lint/releases
GOLANGCI_LINT_VERSION = v1.59.1

#### DO NOT CHANGE THOSE FOLDERS.
## Folder with all reports [tests,coverate,lint]
FOLDER_REPORT = $(CURDIR)/reports
## Folder with binary [compiled file]
FOLDER_BIN = $(CURDIR)/bin

### COMMANDS
all: build

ship: clean mod build-static install

## Display this help screen
help:
	@awk 'BEGIN {FS = ":.*?##"; printf "Usage: make <target>\n"} /^[a-zA-Z_-]+:.*?##/ {printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

## Directories
directories: $(FOLDER_REPORT)

$(FOLDER_REPORT):
	$(MKDIR_P) $@

local.env: ## Creates if missing local.env
	cp .env local.env

coverage: directories ## Generate code coverage report
	go clean -testcache
	go test $(GOFLAGS) $(VENDORFLAGS) -covermode=count -coverprofile="$(FOLDER_REPORT)/test_coverage.out" ./internal/...
	GOFLAGS=$(VENDORFLAGS) go tool cover -func="$(FOLDER_REPORT)/test_coverage.out"
	GOFLAGS=$(VENDORFLAGS) go tool cover -html="$(FOLDER_REPORT)/test_coverage.out" -o "$(FOLDER_REPORT)/test_coverage.html"

lint: directories ## Generate lint report
	$(GOCILINT) --config ./.golangci.yml run
	$(GOCILINT) --config ./.golangci.yml run --out-format=checkstyle ./... > $(FOLDER_REPORT)/linter.xml

test: directories $(GOJUNIT) ## Run unittests
	go clean -testcache
	go test $(GOFLAGS) $(VENDORFLAGS) -race ./internal/... 2>&1 | $(GOJUNIT) -iocopy -set-exit-code -out $(FOLDER_REPORT)/test.xml

vet: ## Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string
	go vet ./...

clean: ## Remove previous build
	rm -rf "$(FOLDER_BIN)"
	rm -rf "$(FOLDER_REPORT)"
	@find . -name ".DS_Store" -print0 | xargs -0 rm -f
	go clean -i ./...

## Build/install the binary file
build: $(FOLDER_BIN)/$(PROJECT_NAME)
build-static: export CGO_ENABLED=0
build-static: LDFLAGS += -extldflags "-static"
build-static: $(FOLDER_BIN)/$(PROJECT_NAME)

$(FOLDER_BIN)/$(PROJECT_NAME): $(SRC)
	go build -o $(FOLDER_BIN)/$(PROJECT_NAME) $(GOFLAGS) -ldflags '$(LDFLAGS)' .

debug: build
	dlv debug --headless --listen=:53599 --api-version=2 --accept-multiclient
	go build -o $(FOLDER_BIN)/$(PROJECT_NAME) $(GOFLAGS) -gcflags \"all=-N -l\"  -ldflags '$(LDFLAGS)' .

build-debug: $(FOLDER_BIN)/$(PROJECT_NAME)
	go build -o $(FOLDER_BIN)/$(PROJECT_NAME) $(GOFLAGS) -gcflags='all=-N -l' .

## Copy binary to bin folder
install: build-static
	$(INSTALL) -Dm755 $(FOLDER_BIN)/$(PROJECT_NAME) $(GOPATH)/bin/

api: ## Run the API cmd
	reflex -s -r '\.go$$' go run main.go

docker-api: ## Run the API inside a docker container
	cd dockers && docker-compose up api

docker-stop: ## Stops all the running docker containers
	cd dockers && docker-compose stop

docker-clean: ## Cleans all the docker images, volumes created
	cd dockers && docker-compose down --rmi local --volumes --remove-orphans

## Get the dependencies
$(GOCILINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin $(GOLANGCI_LINT_VERSION)

$(GOJUNIT):
	go install github.com/jstemmer/go-junit-report/v2@latest


mod: $(GOCILINT) $(GOJUNIT)
	go mod tidy
	go mod vendor


report: test coverage lint ## Generate all reports
