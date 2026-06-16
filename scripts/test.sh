#!/bin/bash

# Unified test runner for ServiceNow SDK for Go

REPORT=false

while [[ "$#" -gt 0 ]]; do
    case $1 in
        --report) REPORT=true ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

echo "Running unit and integration tests..."
go test -coverprofile=coverage.out ./...

if [ "$REPORT" = true ]; then
    echo "Generating HTML coverage report..."
    go tool cover -html=coverage.out -o coverage.html
    echo "Report generated at coverage.html"
fi
