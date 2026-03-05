#!/bin/bash

# submit_report.sh - Submits a bug report to GitHub using 'gh' CLI.
# Usage: ./submit_report.sh "Title" "PathToReportFile"

TITLE=$1
REPORT_PATH=$2

if ! command -v gh &> /dev/null; then
    echo "Error: 'gh' CLI is not installed. Please install it to submit reports automatically."
    exit 1
fi

if ! gh auth status &> /dev/null; then
    echo "Error: 'gh' CLI is not authenticated. Please run 'gh auth login' to submit reports automatically."
    exit 1
fi

if [ -z "$TITLE" ] || [ -z "$REPORT_PATH" ]; then
    echo "Error: Missing title or report path."
    echo "Usage: ./submit_report.sh \"Title\" \"PathToReportFile\""
    exit 1
fi

if [ ! -f "$REPORT_PATH" ]; then
    echo "Error: Report file not found at $REPORT_PATH"
    exit 1
fi

# Create the issue
ISSUE_URL=$(gh issue create --title "$TITLE" --body-file "$REPORT_PATH" --label "type: bug")

if [ $? -eq 0 ]; then
    echo "Success: Bug report submitted successfully! View it here: $ISSUE_URL"
else
    echo "Error: Failed to submit the bug report to GitHub."
    exit 1
fi