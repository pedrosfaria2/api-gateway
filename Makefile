SHELL := /bin/bash
GO := go
SERVICE_NAME := $(shell git remote get-url origin | sed 's/.*\/\([^\/]*\)\.git/\1/')

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  install-deps     Install system dependencies (requires sudo)"
	@echo "  setup            Install Go dependencies"
	@echo "  setup-hooks      Install git hooks and commit message validation"
	@echo "  build            Build the application"
	@echo "  run              Run the application"
	@echo "  run-dev          Run the application with hot-reload using air"
	@echo "  run-docker       Run the application using Docker"
	@echo "  run-compose      Run the application with Docker Compose"
	@echo "  test             Run tests"
	@echo "  test-coverage    Run tests with coverage"
	@echo "  lint             Run linters (golangci-lint)"
	@echo "  format           Format code using gofmt"
	@echo "  clean            Clean build artifacts"
	@echo "  mock-generate    Generate mocks using mockgen"
	@echo "  proto-generate   Generate protobuf files"

.PHONY: install-deps
install-deps:
	@echo "This command requires sudo access to install system dependencies"
	sudo apt-get update
	sudo apt-get install -y pipx
	pipx ensurepath
	pipx install pre-commit
	pipx install commitizen

.PHONY: setup
setup:
	$(GO) mod download
	$(GO) install github.com/air-verse/air@latest
	$(GO) install github.com/golang/mock/mockgen@latest
	# Install golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(HOME)/go/bin v1.55.2

.PHONY: setup-hooks
setup-hooks:
	pre-commit install
	pre-commit install --hook-type commit-msg
	git config commit.template .gitmessage

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux $(GO) build -o bin/main cmd/api/main.go

.PHONY: run
run: build
	./bin/main

.PHONY: run-dev
run-dev:
	air -c .air.toml

.PHONY: run-docker
run-docker:
	docker build --pull --rm -f Dockerfile -t $(SERVICE_NAME):latest .
	docker run -d --name $(SERVICE_NAME) -p 8080:8080 $(SERVICE_NAME):latest

.PHONY: run-compose
run-compose:
	docker-compose up --build

.PHONY: test
test:
	$(GO) test ./... -v

.PHONY: test-coverage
test-coverage:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) tool cover -html=coverage.out

.PHONY: lint
lint:
	$(HOME)/go/bin/golangci-lint run ./...

.PHONY: format
format:
	$(GO) fmt ./...

.PHONY: clean
clean:
	rm -rf bin/
	rm -f coverage.out

.PHONY: mock-generate
mock-generate:
	go generate ./...

.PHONY: proto-generate
proto-generate:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/*.proto
