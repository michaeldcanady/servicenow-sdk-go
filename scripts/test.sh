#!/bin/bash

# Unified test runner for ServiceNow SDK for Go

REPORT=false
MD_REPORT=false

while [[ "$#" -gt 0 ]]; do
    case $1 in
        --report) REPORT=true ;;
        --md-report) MD_REPORT=true ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

if [ "$MD_REPORT" = true ]; then
    echo "Running tests and generating Markdown report..."
    go test -json ./... | go run scripts/generate_test_report.go
else
    echo "Running unit and integration tests..."
    go test -coverprofile=coverage.out ./...
fi

if [ "$REPORT" = true ]; then
    echo "Generating HTML coverage report..."
    go tool cover -html=coverage.out -o coverage.html
    echo "Report generated at coverage.html"
fi
