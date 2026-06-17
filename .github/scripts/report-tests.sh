#!/bin/bash
# Run tests and capture JSON output
go test -json -v ./... > test-output.json
TEST_EXIT_CODE=$?

# Generate markdown summary
echo "### ❌ Test Failure Summary" > test-summary.md
tparse -file test-output.json -format markdown >> test-summary.md

# Output to GitHub Actions Job Summary (Native visualization)
echo "### 🧪 Test Results" >> $GITHUB_STEP_SUMMARY
cat test-summary.md >> $GITHUB_STEP_SUMMARY

# Exit with original test exit code
exit $TEST_EXIT_CODE
