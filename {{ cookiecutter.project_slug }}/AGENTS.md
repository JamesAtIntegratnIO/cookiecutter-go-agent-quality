# Agent Operating Contract

You are working in a Go repository with a strict quality gate. Your job is not complete until the gate passes.

## Prime directive

Do not say "done", "complete", "fixed", or equivalent unless the required commands have passed in this working tree.

## Required completion commands

For every Go code change:

```bash
make quality-fast
```

For changes under `internal/domain` or `internal/service`:

```bash
RUN_MUTATION_TESTS=true make quality-full
```

If mutation testing is too slow for the current environment, say that explicitly and provide the exact command the user must run. Do not pretend it passed.

## Conventions

- Use small functions.
- Prefer explicit errors and `errors.Is`/`errors.As` compatible wrapping.
- Keep domain code free of HTTP, database, framework, and platform concerns.
- Keep services framework-free.
- Put wiring in `cmd/`.
- Put concrete adapters in `internal/repository`, `internal/transport`, or `internal/platform`.
- Tests should verify behavior, not just chase coverage numbers.
- Do not weaken linter, coverage, architecture, or mutation thresholds unless the user explicitly asks.
- Do not add `//nolint` without a short reason and the specific linter name.
- Do not add meaningless tests only to increase coverage.
- Do not ignore failing quality gates.

## When a gate fails

1. Read the failure.
2. Fix the root cause.
3. Re-run the failed command.
4. Only report success after the command passes.

## Preferred Go style

- Table-driven tests for branching behavior.
- Interfaces at consumer boundaries, not producer boundaries.
- `context.Context` as the first parameter when needed.
- No package-level mutable state except deliberate, documented constants/errors.
- No panics outside impossible programmer errors or startup wiring.
- No framework types in domain or service packages.

## Project commands

```bash
make help
make fmt
make lint
make test
make coverage
make mutation
make quality-fast
make quality-full
```
