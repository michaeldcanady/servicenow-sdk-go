#!/bin/bash
# Run tests and capture JSON output
go test -json -v ./... > test-output.json
TEST_EXIT_CODE=$?

# Extract summary data using tparse to get totals
PASS=$(tparse -file test-output.json -format markdown --summary | grep "Pass" | awk '{print $4}')
FAIL=$(tparse -file test-output.json -format markdown --summary | grep "Fail" | awk '{print $4}')
SKIP=$(tparse -file test-output.json -format markdown --summary | grep "Skip" | awk '{print $4}')
TOTAL=$(($PASS + $FAIL + $SKIP))

# Construct the custom Markdown report
{
  echo "### Go Test Report"
  echo "| Metric | Value |"
  echo "| - | - |"
  echo "| Total Tests | $TOTAL |"
  echo "| Passed | $PASS |"
  echo "| Failed | $FAIL |"
  echo "| Skipped | $SKIP |"
  echo ""
  echo "### Failed Tests"
  echo "<details>"
  echo "<summary>Click to expand failed tests</summary>"
  echo ""
  # Use tparse to show just the failed test details
  tparse -file test-output.json -format markdown
  echo ""
  echo "</details>"
} > test-summary.md

# Output to GitHub Actions Job Summary (Native visualization)
echo "### 🧪 Test Results" >> $GITHUB_STEP_SUMMARY
cat test-summary.md >> $GITHUB_STEP_SUMMARY

# Exit with the original test exit code
exit $TEST_EXIT_CODE
