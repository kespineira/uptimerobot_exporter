.PHONY: all build clean default help init test format docker-build docker-run
default: help

GO_TEST_PATH ?= ./...
GO_TEST_EXTRA ?=
GO_TEST_COVER_PROFILE ?= cover.out
GO_TEST_CODECOV ?=

BUILD ?= $(shell date +%FT%T%z)
GOVERSION ?= $(shell go version | cut -d " " -f3)
COMPONENT_VERSION ?= $(shell git describe --abbrev=0 --always)
COMPONENT_BRANCH ?= $(shell git describe --always --contains --all)
PMM_RELEASE_FULLCOMMIT ?= $(shell git rev-parse HEAD)
GO_BUILD_LDFLAGS = -X main.version=${COMPONENT_VERSION} -X main.buildDate=${BUILD} -X main.commit=${PMM_RELEASE_FULLCOMMIT} -X main.Branch=${COMPONENT_BRANCH} -X main.GoVersion=${GOVERSION} -s -w
NAME ?= uptimerobot_exporter
REPO ?= kespineira/$(NAME)
GORELEASER_FLAGS ?=
UID ?= $(shell id -u)

build: ## Build the binary
	@echo "Building binary..."
	@CGO_ENABLED=0 go build -ldflags "$(GO_BUILD_LDFLAGS)" -o $(NAME) ./cmd/uptimerobot_exporter

all: init format test build ## Run all the targets

test: ## Run tests
	@echo "Running tests..."
	@go test -v $(GO_TEST_EXTRA) $(GO_TEST_PATH)

format: ## Format the code
	@echo "Formatting code..."
	@go fmt ./...

init: ## Install dependencies
	@echo "Installing dependencies..."
	@go mod tidy

clean: ## Clean the binary
	@echo "Cleaning..."
	@rm -f $(NAME)

help: ## Show this help
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

docker-build: ## Build the docker image
	@echo "Building docker image..."
	@docker build --build-arg UID=$(UID) -t $(REPO):latest .

docker-run: ## Run the docker image
	@echo "Running docker image..."
	@docker run --rm -p 8080:8080 $(REPO):latest
