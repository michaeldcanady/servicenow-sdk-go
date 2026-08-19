#!/usr/bin/env bash
# Verifies the doc-snippet contract between website/docs and website/snippets:
#   1. every region referenced from a page exists in a snippet file,
#   2. every region defined in a snippet file is referenced by some page
#      (orphans rot silently otherwise),
#   3. every [START x] has a matching [END x] in the same file.
set -euo pipefail

cd "$(dirname "$0")/.."
snippets=website/snippets
docs=website/docs
fail=0

defined=$(grep -ho '\[START [A-Za-z0-9_]*\]' "$snippets"/*.go | sed 's/\[START \(.*\)\]/\1/' | sort -u)
referenced=$(
  {
    grep -rho 'region="[A-Za-z0-9_]*"' "$docs" | sed 's/region="\(.*\)"/\1/'
    grep -rho '{{[A-Za-z0-9_]*:[A-Za-z0-9_]*}}' "$docs" | sed 's/{{.*:\(.*\)}}/\1/'
  } | sort -u
)

missing=$(comm -13 <(echo "$defined") <(echo "$referenced"))
if [ -n "$missing" ]; then
  echo "ERROR: regions referenced in $docs but not defined in $snippets:" >&2
  echo "$missing" | sed 's/^/  /' >&2
  fail=1
fi

orphans=$(comm -23 <(echo "$defined") <(echo "$referenced"))
if [ -n "$orphans" ]; then
  echo "ERROR: regions defined in $snippets but never referenced by a page (delete the markers or wire them in):" >&2
  echo "$orphans" | sed 's/^/  /' >&2
  fail=1
fi

for f in "$snippets"/*.go; do
  starts=$(grep -o '\[START [A-Za-z0-9_]*\]' "$f" | sed 's/\[START \(.*\)\]/\1/' | sort)
  ends=$(grep -o '\[END [A-Za-z0-9_]*\]' "$f" | sed 's/\[END \(.*\)\]/\1/' | sort)
  if [ "$starts" != "$ends" ]; then
    echo "ERROR: unbalanced START/END markers in $f:" >&2
    diff <(echo "$starts") <(echo "$ends") | sed 's/^/  /' >&2 || true
    fail=1
  fi
done

if [ "$fail" -eq 0 ]; then
  echo "snippet regions OK: $(echo "$defined" | wc -l) regions, all referenced and balanced"
fi
exit "$fail"
