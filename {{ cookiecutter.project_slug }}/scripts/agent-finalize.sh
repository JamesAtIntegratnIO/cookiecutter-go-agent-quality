#!/usr/bin/env bash
set -euo pipefail

# This is the only valid local completion path for agents.
# It proves the requested quality gates pass against the current workspace and
# emits a local completion artifact the agent must include in its final answer.

mkdir -p .agent
rm -f .agent/done.json .agent/final-report.md

gate="make quality-fast"
mutation_gate="RUN_MUTATION_TESTS=true make quality-full"
acceptance_gate="./scripts/check-acceptance.sh"

hash_file() {
  if command -v shasum >/dev/null 2>&1; then
    shasum -a 256 "$1"
  elif command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$1"
  else
    openssl dgst -sha256 "$1"
  fi
}

hash_stdin() {
  if command -v shasum >/dev/null 2>&1; then
    shasum -a 256
  elif command -v sha256sum >/dev/null 2>&1; then
    sha256sum
  else
    openssl dgst -sha256
  fi
}

changed_files=""
if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  changed_files="$(git status --porcelain | awk '{print $2}')"
fi

requires_mutation="false"
if printf '%s\n' "${changed_files}" | grep -Eq '^(internal/domain|internal/service)/'; then
  requires_mutation="true"
fi

if [[ "${AGENT_REQUIRE_MUTATION:-auto}" == "true" ]]; then
  requires_mutation="true"
elif [[ "${AGENT_REQUIRE_MUTATION:-auto}" == "false" ]]; then
  requires_mutation="false"
fi

echo "==> running fast quality gate"
${gate}

if [[ "${requires_mutation}" == "true" ]]; then
  echo "==> domain/service changes detected; running mutation gate"
  ${mutation_gate}
else
  echo "==> no domain/service changes detected; mutation gate not required"
fi

if [[ -x scripts/check-acceptance.sh ]]; then
  echo "==> running acceptance checks"
  ${acceptance_gate}
else
  echo "==> no executable acceptance check found; skipping"
fi

# Create a deterministic fingerprint of the current workspace, excluding git and
# local agent artifacts. This lets the final report identify exactly what passed.
fingerprint="$({
  find . \
    -path './.git' -prune -o \
    -path './.agent' -prune -o \
    -type f -print \
  | LC_ALL=C sort \
  | while IFS= read -r file; do
      hash_file "$file"
    done
} | hash_stdin | awk '{print $1}')"

timestamp="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"
commit="not-a-git-repository"
if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  commit="$(git rev-parse --short HEAD 2>/dev/null || echo uncommitted)"
fi

cat > .agent/done.json <<EOF_JSON
{
  "status": "passed",
  "timestamp": "${timestamp}",
  "commit": "${commit}",
  "fingerprint": "${fingerprint}",
  "quality_gate": "${gate}",
  "mutation_required": ${requires_mutation},
  "acceptance_gate": "${acceptance_gate}"
}
EOF_JSON

cat > .agent/final-report.md <<EOF_REPORT
# Agent Final Report

Status: PASSED

Timestamp: ${timestamp}
Commit: ${commit}
Workspace fingerprint: ${fingerprint}

Validated commands:

\`\`\`bash
${gate}
$(if [[ "${requires_mutation}" == "true" ]]; then echo "${mutation_gate}"; fi)
$(if [[ -x scripts/check-acceptance.sh ]]; then echo "${acceptance_gate}"; fi)
\`\`\`

The agent completion contract was satisfied for the workspace fingerprint above.
Do not edit files after this report is generated unless you run \`make agent-finalize\` again.
EOF_REPORT

cat .agent/final-report.md
