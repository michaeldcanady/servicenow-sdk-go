#!/usr/bin/env bash
# Copyright (c) 2021 Michael Canady
# SPDX-License-Identifier: MIT

# Fails when the changed, non-test Go lines of a change fall below a coverage floor.
#
# Usage: scripts/check-patch-coverage.sh <profile> <base-rev> <floor-percent>
#
# A changed line counts as "relevant" when a block in the coverage profile
# contains it, and as "covered" when one of those blocks ran at least once.
# Lines outside every block (comments, declarations, blank lines) are ignored.
# The script passes when no relevant line changed.
set -euo pipefail

if [ "$#" -ne 3 ]; then
  echo "usage: $0 <profile> <base-rev> <floor-percent>" >&2
  exit 2
fi
profile=$1
base=$2
floor=$3

if [ ! -f "$profile" ]; then
  echo "::error::Coverage profile '$profile' does not exist" >&2
  exit 2
fi

module=$(go list -m)

# Emit "path start end" for each added or modified line range in non-test Go files.
ranges=$(git diff --unified=0 --no-color --diff-filter=AM "$base" HEAD -- '*.go' ':(exclude)*_test.go' |
  awk '
    /^\+\+\+ b\// { file = substr($0, 7); next }
    /^@@/ {
      split($3, a, ","); start = substr(a[1], 2); n = (a[2] == "" ? 1 : a[2])
      if (n > 0) print file, start, start + n - 1
    }')

if [ -z "$ranges" ]; then
  echo "::notice::No non-test Go lines changed; skipping the patch coverage floor."
  exit 0
fi

result=$(awk -v module="$module" '
  FNR == NR { nr++; f[nr] = $1; s[nr] = $2; e[nr] = $3; next }
  FNR == 1 && /^mode:/ { next }
  {
    split($1, p, ":"); path = p[1]; sub("^" module "/", "", path)
    split(p[2], r, ","); split(r[1], a, "."); split(r[2], b, ".")
    for (l = a[1]; l <= b[1]; l++) { rel[path SUBSEP l] = 1; if ($3 > 0) cov[path SUBSEP l] = 1 }
  }
  END {
    for (i = 1; i <= nr; i++)
      for (l = s[i]; l <= e[i]; l++) {
        k = f[i] SUBSEP l
        if (k in rel) { total++; if (k in cov) hit++; else miss[f[i] ":" l] = 1 }
      }
    for (m in miss) print "MISS " m
    printf "TOTAL %d %d\n", total, hit
  }' <(printf '%s\n' "$ranges") "$profile")

read -r _ total hit < <(grep '^TOTAL' <<<"$result")
if [ "${total:-0}" -eq 0 ]; then
  echo "::notice::Changed Go lines contain no statements; skipping the patch coverage floor."
  exit 0
fi

pct=$(awk -v h="$hit" -v t="$total" 'BEGIN { printf "%.1f", 100 * h / t }')
summary="Patch coverage: ${pct}% (${hit}/${total} changed statement lines); floor: ${floor}%."
echo "$summary"
if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
  echo "$summary" >>"$GITHUB_STEP_SUMMARY"
fi

if awk -v p="$pct" -v f="$floor" 'BEGIN { exit !(p + 0 < f + 0) }'; then
  grep '^MISS' <<<"$result" | sort | sed 's/^MISS //' | while IFS=: read -r file line; do
    echo "::error file=${file},line=${line}::Changed line is not covered by tests."
  done
  echo "::error::Patch coverage ${pct}% is below the ${floor}% floor. Add tests for the changed lines listed above."
  exit 1
fi
