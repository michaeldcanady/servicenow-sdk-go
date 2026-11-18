#!/bin/bash
# Copyright (c) 2026 Michael Canady
# SPDX-License-Identifier: MIT

#
# Applies (or verifies) MIT license headers to this repo's in-scope source
# files using google/addlicense, pinned via `go run ...@v1.2.0` so the tool
# stays out of go.mod/go.sum (mirrors how tparse is pinned in ci.yml).
#
# Scope (locked decision):
#   - **/*.go            incl. website/snippets
#   - **/*.sh            repo-owned shell, incl. .github/scripts and .devcontainer
#   - root *.yml/*.yaml  config files only (workflow yaml under .github/ is out)
# The file list is derived from `git ls-files`, so .git/, untracked
# node_modules/build/vendor dirs can never be swept into the header pass.
#
# Year policy:
#   - Existing files keep their static creation year: the one-time backfill
#     stamped 2021, and addlicense skips any file that already carries a
#     header, so re-runs never rewrite it.
#   - New files get the current year: the default is -y $(date +%Y).
#
# Usage:
#   scripts/add-license.sh [--check-only] [--staged] [-y <year>]
#
#   --check-only   addlicense -check: verify every in-scope file carries a
#                  header; exit non-zero if any does not. No modifications.
#   --staged       limit the run to in-scope files staged for commit, and
#                  stage the headers it adds. Used by the pre-commit hook.
#                  Fails if a file that needs a header also has unstaged
#                  changes, so the hook never stages work you left out.
#   -y <year>      copyright year to stamp (default: current year).
set -euo pipefail

CHECK_ONLY=false
STAGED=false
YEAR="$(date +%Y)"

usage() {
    awk 'NR > 1 && /^#/ { sub(/^# ?/, ""); print } NR > 1 && !/^#/ { exit }' "${BASH_SOURCE[0]}"
}

while [[ $# -gt 0 ]]; do
    case $1 in
        --check-only) CHECK_ONLY=true ;;
        --staged) STAGED=true ;;
        -y) YEAR="$2"; shift ;;
        -h|--help) usage; exit 0 ;;
        *) echo "Unknown parameter passed: $1" >&2; exit 1 ;;
    esac
    shift
done

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT"

# In-scope files: every tracked *.go / *.sh, plus root-level *.yml/*.yaml
# config files (a root file has no '/' in its path).
FILES=()
while IFS= read -r f; do
    case "$f" in
        *.go|*.sh) FILES+=("$f") ;;
        *.yml|*.yaml)
            if [[ "$f" != */* ]]; then
                FILES+=("$f")
            fi
            ;;
    esac
done < <(
    if [ "$STAGED" = true ]; then
        git diff --cached --name-only --diff-filter=AM -- '*.go' '*.sh' '*.yml' '*.yaml'
    else
        git ls-files -- '*.go' '*.sh' '*.yml' '*.yaml'
    fi
)

if [ "${#FILES[@]}" -eq 0 ]; then
    echo "No in-scope files found; nothing to do." >&2
    exit 0
fi

COMMON_ARGS=(
    -c "Michael Canady"
    -l mit
    -f "$REPO_ROOT/scripts/license-header.tmpl"
    -y "$YEAR"
)

if [ "$CHECK_ONLY" = true ]; then
    echo "Checking license headers on ${#FILES[@]} in-scope files (year looks like: $YEAR)..."
    go run github.com/google/addlicense@v1.2.0 -check "${COMMON_ARGS[@]}" -- "${FILES[@]}"
    echo "OK: all in-scope files carry a license header."
else
    echo "Applying license headers (year $YEAR) to ${#FILES[@]} in-scope files..."
    if [ "$STAGED" = true ]; then
        # Record which files have unstaged changes before addlicense edits them.
        UNSTAGED="$(git diff --name-only)"
    fi
    go run github.com/google/addlicense@v1.2.0 "${COMMON_ARGS[@]}" -- "${FILES[@]}"
    if [ "$STAGED" = true ]; then
        for f in "${FILES[@]}"; do
            if git diff --quiet -- "$f"; then
                continue
            fi
            if grep -qxF -- "$f" <<<"$UNSTAGED"; then
                echo "ERROR: $f needs a license header but has unstaged changes." >&2
                echo "Stage or stash those changes, then commit again." >&2
                exit 1
            fi
            git add -- "$f"
            echo "Added a license header to $f and staged it."
        done
    fi
    echo "Done. Review with: git diff --stat"
fi