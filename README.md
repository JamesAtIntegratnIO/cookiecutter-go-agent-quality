# Cookiecutter Go Agent Quality Template

A Cookiecutter template for creating Go services with enforceable quality gates designed for AI-assisted development.

This template is opinionated around one core rule:

> The agent is not done until the repo proves it is done.

Generated projects include Go scaffolding, linting, coverage enforcement, architecture boundary tests, mutation testing support, GitHub Actions, and agent instruction files for tools like GitHub Copilot, Claude, Cursor, Windsurf, and Hermes.

## What this template creates

Generated repositories include:

```text
.
├── cmd/
│   └── <binary_name>/
│       └── main.go
├── internal/
│   ├── architecture/
│   ├── domain/
│   ├── repository/
│   ├── service/
│   └── transport/
├── scripts/
│   ├── check-coverage.sh
│   └── quality.sh
├── .github/
│   ├── workflows/
│   │   └── quality.yml
│   └── copilot-instructions.md
├── .cursor/
│   └── rules/
│       └── go-quality.mdc
├── .hermes/
│   └── SOUL.md
├── AGENTS.md
├── CLAUDE.md
├── .windsurfrules
├── .golangci.yml
├── Makefile
├── go.mod
└── README.md
```

## Quality gates included

The generated project includes:

- `gofmt`
- `go mod tidy`
- `go test ./...`
- race-enabled tests
- coverage threshold enforcement
- `golangci-lint`
- architecture boundary tests
- optional mutation testing with `gremlins`
- GitHub Actions CI
- agent instruction files that tell coding agents exactly how to work in the repo

## Prerequisites

Install Cookiecutter:

```bash
brew install cookiecutter
```

Or with Python:

```bash
python3 -m pip install --user cookiecutter
```

For generated Go projects, you will also need:

```bash
go version
golangci-lint version
```

The generated project includes a `make install-tools` target to install Go-based tooling such as `gremlins`.

## Using this template locally

From a local clone of this template:

```bash
cookiecutter ./cookiecutter-go-agent-quality
```

Then answer the prompts:

```text
project_name [Go Agent Quality Service]:
project_slug [go-agent-quality-service]:
module_path [github.com/example/go-agent-quality-service]:
package_name [goagentqualityservice]:
binary_name [go-agent-quality-service]:
coverage_threshold [80]:
```

Then enter the generated project:

```bash
cd go-agent-quality-service
go mod tidy
make install-tools
make quality-fast
```

## Using this template after pushing it to GitHub

Once this Cookiecutter template is pushed to GitHub, users can generate a new project directly from the GitHub repo URL.

Example:

```bash
cookiecutter gh:YOUR_GITHUB_ORG/cookiecutter-go-agent-quality
```

Or using the full URL:

```bash
cookiecutter https://github.com/YOUR_GITHUB_ORG/cookiecutter-go-agent-quality
```

For a private repository, make sure your local Git authentication is configured first:

```bash
gh auth login
```

Then run:

```bash
cookiecutter git@github.com:YOUR_GITHUB_ORG/cookiecutter-go-agent-quality.git
```

## Recommended flow for creating a new project

```bash
cookiecutter gh:YOUR_GITHUB_ORG/cookiecutter-go-agent-quality
cd <generated-project>
go mod tidy
make install-tools
make quality-fast
git init
git add .
git commit -m "Initial Go service scaffold"
gh repo create <generated-project> --private --source=. --remote=origin --push
```

For a public repo:

```bash
gh repo create <generated-project> --public --source=. --remote=origin --push
```

## Recommended flow for using with an AI agent

After generating the project, tell the agent:

```text
Read AGENTS.md, .github/copilot-instructions.md, and the README before making changes.

Before saying a task is complete, run:

make quality-fast

If domain or service logic changed, also run:

RUN_MUTATION_TESTS=true make quality-full

Do not weaken lint rules, coverage thresholds, or architecture tests without explicit approval.
```

The generated repo includes agent instruction files for:

- `AGENTS.md`
- GitHub Copilot
- Claude
- Cursor
- Windsurf
- Hermes

These files are intentionally repetitive. The goal is to make sure whichever agent or editor is active receives the same constraints.

## Generated project commands

Inside a generated project:

```bash
make test
```

Runs normal Go tests.

```bash
make coverage
```

Runs tests with coverage and enforces the configured threshold.

```bash
make lint
```

Runs `golangci-lint`.

```bash
make mutation
```

Runs mutation tests with `gremlins`.

```bash
make quality-fast
```

Runs the normal local quality gate.

```bash
make quality-full
```

Runs the full quality gate, including mutation testing.

## Suggested local development loop

Use this loop while working manually or with an agent:

```bash
make quality-fast
```

For significant domain/service changes:

```bash
RUN_MUTATION_TESTS=true make quality-full
```

Use mutation testing selectively. It is valuable, but it can be slow on larger packages.

## CI behavior

The generated project includes a GitHub Actions workflow at:

```text
.github/workflows/quality.yml
```

The default intent is:

- Pull requests run the fast quality gate.
- Pushes to `main` can run the fuller gate.
- Mutation testing is available but should be enabled intentionally.

If mutation testing is too slow for regular CI, move it to a scheduled workflow:

```yaml
on:
  schedule:
    - cron: "0 6 * * *"
```

## Updating this template

After modifying the template, test generation locally:

```bash
cookiecutter . --no-input
```

Then enter the generated repo and run:

```bash
go mod tidy
make install-tools
make quality-fast
```

If the generated project fails quality checks, fix the template before pushing.

## Template repository maintenance checklist

Before pushing changes to this template repo:

```bash
cookiecutter . --no-input
cd <generated-project>
go mod tidy
make quality-fast
```

Also verify:

- generated project names render correctly
- module path is valid
- scripts are executable
- GitHub Actions YAML is valid
- agent instruction files are present
- coverage threshold is achievable
- architecture tests pass
- lint config matches the installed `golangci-lint` major version

## Why this exists

AI coding agents are fast, but they are also very good at producing code that looks correct while quietly degrading maintainability.

This template makes the repo enforce the standard:

- tests prove behavior
- coverage prevents fake completeness
- linting catches maintainability problems
- architecture tests protect package boundaries
- mutation testing checks whether tests actually fail when behavior changes
- agent instructions make the workflow explicit

The goal is not to slow down development. The goal is to prevent “vibe-coded” software from reaching `main` without objective proof that it works.