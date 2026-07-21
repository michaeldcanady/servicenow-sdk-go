#!/bin/bash
#
# Prints (or runs) the set of Go packages affected by the current diff, including
# transitive dependents. E.g. a change under core/ pulls in every *api package that
# imports core, not just core itself.
#
# Usage:
#   scripts/affected-tests.sh [base-ref] [--run]
#
#   base-ref   Ref to diff against. Resolved in this order if omitted:
#                1. $GITHUB_BASE_REF (set by GitHub Actions on pull_request
#                   events to the actual PR base branch, e.g. "main" or a
#                   release branch) - CI must fetch that ref too, since
#                   actions/checkout defaults to fetch-depth: 1.
#                2. origin/HEAD, if resolvable (local dev clones).
#                3. origin/main as a last-resort fallback.
#              Use HEAD~1 for the last commit.
#   --run      Instead of just printing the affected package list, run `go test` on it.
#
# How it works:
#   1. `git diff --name-only <base-ref>...HEAD -- '*.go'` finds changed files, mapped to
#      their containing package import paths via `go list`.
#   2. `go list -f '{{.ImportPath}} {{join .Deps ","}}' ./...` is used to get, for every
#      package in the module, its FULL TRANSITIVE dependency closure (the `Deps` field is
#      already transitive - no manual graph walk needed). A package is "affected" if it was
#      changed directly, or if any package in its dependency closure was changed.
#
# This intentionally does not try to be smarter than that (no test-to-source heuristics,
# no caching) - see CLAUDE.md's guidance against speculative abstraction. If this needs to
# get fancier later (e.g. skipping doc-only diffs, integration test tags), extend here.

set -euo pipefail

BASE_REF=""
DO_RUN=false

for arg in "$@"; do
  case "$arg" in
    --run) DO_RUN=true ;;
    *) BASE_REF="$arg" ;;
  esac
done

if [ -z "$BASE_REF" ]; then
  if [ -n "${GITHUB_BASE_REF:-}" ]; then
    # Set by GitHub Actions on pull_request events to the PR's actual base
    # branch - this is the only correct source in CI, since the PR base
    # isn't necessarily the repo's default branch (e.g. a PR into a release
    # branch), and a shallow checkout has no usable origin/HEAD anyway.
    BASE_REF="origin/${GITHUB_BASE_REF}"
  else
    BASE_REF=$(git symbolic-ref --short refs/remotes/origin/HEAD 2>/dev/null || echo "origin/main")
  fi
fi

echo "Diffing against ${BASE_REF}..." >&2

CHANGED_GO_FILES=$(git diff --name-only "${BASE_REF}...HEAD" -- '*.go' || true)

if [ -z "$CHANGED_GO_FILES" ]; then
  echo "No changed .go files vs ${BASE_REF}; nothing to test." >&2
  exit 0
fi

CHANGED_PKGS=$(printf '%s\n' "$CHANGED_GO_FILES" \
  | xargs -n1 dirname \
  | sort -u \
  | while read -r dir; do
      [ -d "$dir" ] && (cd "$dir" && go list . 2>/dev/null) || true
    done \
  | sort -u)

if [ -z "$CHANGED_PKGS" ]; then
  echo "No importable packages among changed files; nothing to test." >&2
  exit 0
fi

echo "Directly changed packages:" >&2
printf '%s\n' "$CHANGED_PKGS" | sed 's/^/  - /' >&2

AFFECTED=""
while IFS=' ' read -r pkg deps; do
  affected=false

  if printf '%s\n' "$CHANGED_PKGS" | grep -qxF "$pkg"; then
    affected=true
  else
    IFS=',' read -ra deplist <<< "$deps"
    for d in "${deplist[@]}"; do
      if printf '%s\n' "$CHANGED_PKGS" | grep -qxF "$d"; then
        affected=true
        break
      fi
    done
  fi

  if [ "$affected" = true ]; then
    AFFECTED="${AFFECTED}${pkg}
"
  fi
done < <(go list -f '{{.ImportPath}} {{join .Deps ","}}' ./... | grep -v '/docs ')

AFFECTED=$(printf '%s' "$AFFECTED" | sort -u)

echo "Affected packages (direct + transitive dependents):" >&2
printf '%s\n' "$AFFECTED" | sed 's/^/  - /' >&2

if [ "$DO_RUN" = true ]; then
  # shellcheck disable=SC2086
  go test -coverprofile=coverage.out $AFFECTED
else
  printf '%s\n' "$AFFECTED"
fi
