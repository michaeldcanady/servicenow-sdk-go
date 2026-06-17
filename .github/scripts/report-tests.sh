#!/bin/bash
# Run tests and capture JSON output and generate coverage profile
go test -coverprofile=coverage.out -json -v ./... > test-output.json
TEST_EXIT_CODE=$?

# Generate the report using the new custom tool
cat test-output.json | go run scripts/generate_test_report.go > test-summary.md

# Append the sticky comment footer
echo "" >> test-summary.md
echo "<!-- Sticky Pull Request Commenttest-failure-summary -->" >> test-summary.md

# Output to GitHub Actions Job Summary
cat test-summary.md >> $GITHUB_STEP_SUMMARY

# Exit with original test exit code
exit $TEST_EXIT_CODE
