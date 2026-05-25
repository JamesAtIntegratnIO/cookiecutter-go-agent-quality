SHELL := /usr/bin/env bash
COVERAGE_THRESHOLD ?= {{ cookiecutter.coverage_threshold }}
RUN_MUTATION_TESTS ?= {{ cookiecutter.enable_mutation_testing }}

.PHONY: help install-tools fmt lint test coverage mutation quality-fast quality-full agent-finalize clean

help:
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-18s %s\n", $$1, $$2}'

install-tools: ## Install local quality tools
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	go install github.com/go-gremlins/gremlins/cmd/gremlins@latest

fmt: ## Format and tidy
	gofmt -w .
	go mod tidy

lint: ## Run linters
	golangci-lint run ./...

test: ## Run tests with race detector
	go test -race ./...

coverage: ## Run coverage gate for internal packages
	go test -race -covermode=atomic -coverpkg=./internal/... -coverprofile=coverage.out ./internal/...
	./scripts/check-coverage.sh coverage.out $(COVERAGE_THRESHOLD)

mutation: ## Run mutation tests against behavior-heavy packages
	gremlins unleash ./internal/domain
	gremlins unleash ./internal/service

quality-fast: fmt lint coverage ## Fast gate agents must run before claiming done

quality-full: fmt lint coverage ## Full gate with optional mutation testing
	@if [[ "$(RUN_MUTATION_TESTS)" == "true" ]]; then \
		$(MAKE) mutation; \
	else \
		echo "mutation tests skipped; set RUN_MUTATION_TESTS=true to enable"; \
	fi

agent-finalize: ## Prove the agent completion contract and emit .agent/final-report.md
	./scripts/agent-finalize.sh

clean: ## Clean generated quality artifacts
	rm -f coverage.out coverage.html
	rm -rf .agent
