# {{ cookiecutter.project_name }}

This repository was generated from `cookiecutter-go-agent-quality`.

It is intentionally opinionated: agents are not allowed to call work complete until the quality gate passes.

## What this gives you

- Go service scaffold with clean-ish internal package boundaries
- Make targets for formatting, linting, coverage, tests, and mutation testing
- `golangci-lint` v2 config with complexity and maintainability checks
- Architecture tests that prevent domain/service layers from importing transport and infra concerns
- GitHub Actions quality gate
- Agent instruction files for Copilot, Cursor, Claude, Windsurf, Hermes, and generic `AGENTS.md`
- Optional mutation testing through Gremlins

## First run

```bash
go mod tidy
make install-tools
make quality-fast
```

## Common commands

```bash
make fmt             # gofmt + go mod tidy
make lint            # golangci-lint
make test            # go test -race ./...
make coverage        # coverage gate against ./internal/...
make mutation        # gremlins against domain/service
make quality-fast    # fmt + lint + coverage
make quality-full    # fmt + lint + coverage + mutation
```

## Quality contract

The agent must run `make quality-fast` before saying done.

If domain or service logic changes, the agent must also run:

```bash
RUN_MUTATION_TESTS=true make quality-full
```

## Architecture rules

The project enforces these boundaries:

- `internal/domain` must not import HTTP frameworks, SQL drivers, repositories, transports, or platform packages.
- `internal/service` must not import HTTP frameworks or concrete transport packages.
- `internal/transport` can depend on service interfaces.
- `cmd/` wires dependencies only.

## Generated package layout

```text
cmd/{{ cookiecutter.binary_name }}/
internal/domain/
internal/service/
internal/repository/
internal/transport/http/
internal/platform/
internal/architecture/
scripts/
.github/workflows/
.github/copilot-instructions.md
.cursor/rules/
.hermes/SOUL.md
AGENTS.md
```
