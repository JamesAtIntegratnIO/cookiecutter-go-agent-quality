#!/usr/bin/env bash
set -euo pipefail

COVERAGE_THRESHOLD="${COVERAGE_THRESHOLD:-{{ cookiecutter.coverage_threshold }}}"
RUN_MUTATION_TESTS="${RUN_MUTATION_TESTS:-{{ cookiecutter.enable_mutation_testing }}}"

echo "==> fmt"
gofmt -w .
go mod tidy

echo "==> lint"
golangci-lint run ./...

echo "==> coverage"
go test -race -covermode=atomic -coverpkg=./internal/... -coverprofile=coverage.out ./internal/...
./scripts/check-coverage.sh coverage.out "$COVERAGE_THRESHOLD"

if [[ "$RUN_MUTATION_TESTS" == "true" ]]; then
  echo "==> mutation"
  gremlins unleash ./internal/domain ./internal/service
else
  echo "==> mutation skipped; set RUN_MUTATION_TESTS=true to enable"
fi

echo "quality gate passed"
