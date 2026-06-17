#!/bin/bash
# Run tests and capture JSON output
# We use '|| true' to prevent the script from exiting immediately on test failure
go test -json -v ./... > test-output.json
TEST_EXIT_CODE=$?

# Generate markdown summary using tparse
echo "### 🧪 Go Test Report" > test-summary.md
echo "" >> test-summary.md
tparse -file test-output.json -format markdown >> test-summary.md

# Output to GitHub Actions Job Summary (Native visualization)
echo "### 🧪 Test Results" >> $GITHUB_STEP_SUMMARY
cat test-summary.md >> $GITHUB_STEP_SUMMARY

# Exit with the original test exit code so CI correctly marks the job as failed if tests failed
exit $TEST_EXIT_CODE
