#!/bin/bash
# Run tests and capture JSON output
go test -json -v ./... > test-output.json
TEST_EXIT_CODE=$?

# Generate the report
{
  echo "### 🧪 Go Test Report"
  echo ""
  # Tparse summary is cleaner and automatically includes pass/fail counts
  tparse -file test-output.json -format markdown --summary
  echo ""
  echo "### ❌ Failed Tests"
  echo "<details>"
  echo "<summary>Click to expand failed tests</summary>"
  echo ""
  # Tparse's default output without --summary shows failed tests
  tparse -file test-output.json -format markdown
  echo ""
  echo "</details>"
} > test-summary.md

# Output to GitHub Actions Job Summary
cat test-summary.md >> $GITHUB_STEP_SUMMARY

# Exit with original test exit code
exit $TEST_EXIT_CODE
