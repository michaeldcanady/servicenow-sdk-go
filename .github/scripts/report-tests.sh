#!/bin/bash
# Run tests and capture JSON output
go test -json -v ./... > test-output.json
TEST_EXIT_CODE=$?

# Generate markdown summary using tparse, ensuring it's cleaner
# We use tparse to generate the tabular summary
echo "### 🧪 Go Test Report" > test-summary.md
echo "" >> test-summary.md
tparse -file test-output.json -format markdown >> test-summary.md

# Append a summary table if possible (tparse includes this by default in markdown format)

# Output to GitHub Actions Job Summary (Native visualization)
echo "### 🧪 Test Results" >> $GITHUB_STEP_SUMMARY
cat test-summary.md >> $GITHUB_STEP_SUMMARY

# Exit with original test exit code
exit $TEST_EXIT_CODE
