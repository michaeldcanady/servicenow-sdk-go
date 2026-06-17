#!/bin/bash
# Run tests and capture JSON output
go test -json -v ./... > test-output.json
TEST_EXIT_CODE=$?

# Generate the report
{
  echo "### 🧪 Go Test Report"
  echo ""
  
  # 1. Full Package Table (tparse defaults to this without --summary)
  tparse -file test-output.json -format markdown
  echo ""
  
  # 2. Collapsible Failed Tests
  echo "### ❌ Failed Tests"
  echo "<details>"
  echo "<summary>Click to expand failed tests</summary>"
  echo ""
  
  # Extract failures and format specifically
  # This relies on tparse structure to output failures
  tparse -file test-output.json -format markdown --all | sed -n '/FAIL/,$p'
  
  echo ""
  echo "</details>"
  echo "<!-- Sticky Pull Request Commenttest-failure-summary -->"
} > test-summary.md

# Output to GitHub Actions Job Summary
cat test-summary.md >> $GITHUB_STEP_SUMMARY

# Exit with original test exit code
exit $TEST_EXIT_CODE
