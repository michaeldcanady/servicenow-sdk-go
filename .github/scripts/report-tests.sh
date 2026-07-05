#!/bin/bash
set -o pipefail

TEST_EXIT_CODE=0
go test -coverprofile=coverage.out -json -v ./... > test-output.json
echo "exit: $?"
grep '"Action":"fail"' test-output.json

cat test-output.json | go run scripts/generate_test_report.go > test-summary.md || true

echo "" >> test-summary.md
echo "<!-- Sticky Pull Request Comment: test-failure-summary -->" >> test-summary.md

cat test-summary.md >> "$GITHUB_STEP_SUMMARY"

exit $TEST_EXIT_CODE