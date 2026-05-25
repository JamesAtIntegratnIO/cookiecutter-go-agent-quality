#!/usr/bin/env bash
set -euo pipefail

profile="${1:-coverage.out}"
threshold="${2:-80}"

if [[ ! -f "$profile" ]]; then
  echo "coverage profile not found: $profile" >&2
  exit 1
fi

coverage="$(go tool cover -func="$profile" | awk '/total:/ { gsub("%", "", $3); print $3 }')"

awk -v coverage="$coverage" -v threshold="$threshold" '
BEGIN {
  if (coverage + 0 < threshold + 0) {
    printf "coverage %.2f%% is below threshold %.2f%%\n", coverage, threshold
    exit 1
  }

  printf "coverage %.2f%% meets threshold %.2f%%\n", coverage, threshold
}
'
