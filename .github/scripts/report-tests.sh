#!/bin/bash
# Run tests and capture JSON output
go test -json -v ./... > test-output.json
TEST_EXIT_CODE=$?

# Initialize report files
echo "### 🧪 Go Test Report" > test-summary.md
echo "" >> test-summary.md

# 1. Summary Table
echo "#### Summary" >> test-summary.md
tparse -file test-output.json -format markdown --summary >> test-summary.md
echo "" >> test-summary.md

# 2. Collapsible Failed Tests
echo "<details><summary>❌ Failed Tests</summary>" >> test-summary.md
echo "" >> test-summary.md
# We use tparse's ability to show just the failures
tparse -file test-output.json -format markdown >> test-summary.md
echo "" >> test-summary.md
echo "</details>" >> test-summary.md

# Output to GitHub Actions Job Summary (Native visualization)
echo "### 🧪 Test Results" >> $GITHUB_STEP_SUMMARY
cat test-summary.md >> $GITHUB_STEP_SUMMARY

# Exit with the original test exit code
exit $TEST_EXIT_CODE
