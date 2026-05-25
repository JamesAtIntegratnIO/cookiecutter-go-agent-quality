# Agent Operating Contract

You are working in a Go repository with a strict quality gate. Your job is not complete until the repository proves it.

## Prime directive

Do not say "done", "complete", "fixed", or equivalent based on reasoning, todo state, or partial command output.

The only valid local completion path is:

```bash
make agent-finalize
```

If `make agent-finalize` fails, do not summarize success. Report the failing command and the relevant output.

If `make agent-finalize` passes, include the contents of `.agent/final-report.md` in your final response.

Do not edit files after `make agent-finalize` succeeds. If you edit anything after finalization, run `make agent-finalize` again.

## Required quality gates

For every Go code change, the finalizer runs:

```bash
make quality-fast
```

For changes under `internal/domain` or `internal/service`, the finalizer automatically runs:

```bash
RUN_MUTATION_TESTS=true make quality-full
```

If mutation testing is too slow or unavailable, say that explicitly and provide the exact command the user must run. Do not pretend it passed.

## Acceptance checks

When adding user-facing behavior, update `scripts/check-acceptance.sh` with smoke tests for the requested behavior.

Examples:

```bash
go run ./cmd/example --help
go test ./test/acceptance/...
curl -fsS http://localhost:8080/healthz
```

Unit tests, linting, and coverage are not a substitute for acceptance checks when the task changes CLI flags, HTTP routes, request/response shapes, config behavior, or external integration behavior.

## Editing discipline

- Prefer small patches.
- Do not rewrite an entire file unless the file is short or the user explicitly asks.
- After a failed patch, read the file before trying again.
- Do not keep retrying the same failed patch shape.
- Do not leave duplicated code, malformed blocks, or partial rewrites behind.
- Do not hand-roll CLI parsing unless explicitly required. Use Go's standard `flag` package or a mature CLI library such as Cobra.

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
4. Only report success after `make agent-finalize` passes.

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
make agent-finalize
```
