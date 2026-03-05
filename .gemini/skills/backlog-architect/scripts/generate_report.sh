#!/bin/bash

# generate_report.sh - Generates a bug report from provided inputs.
# Usage: ./generate_report.sh "Title" "Description" "ReproductionSteps" "ExpectedBehavior" "AdditionalContext"

TITLE=$1
DESCRIPTION=$2
REPRODUCTION_STEPS=$3
EXPECTED_BEHAVIOR=$4
ADDITIONAL_CONTEXT=$5
OUTPUT_FILE=$6

TEMPLATE_PATH="/workspaces/servicenow-sdk-go/.gemini/skills/bug-reporter/references/bug_report_template.md"
GATHER_INFO_SCRIPT="/workspaces/servicenow-sdk-go/.gemini/skills/bug-reporter/scripts/gather_info.sh"

if [ -z "$TITLE" ] || [ -z "$DESCRIPTION" ] || [ -z "$OUTPUT_FILE" ]; then
    echo "Error: Missing required arguments (Title, Description, OutputFile)."
    exit 1
fi

# Gather system info
SYSTEM_INFO=$($GATHER_INFO_SCRIPT)
OS=$(echo "$SYSTEM_INFO" | grep "OS:" | sed 's/OS: //')
GO_VERSION=$(echo "$SYSTEM_INFO" | grep "Go version:" | sed 's/Go version: //')
MODULE_VERSION=$(echo "$SYSTEM_INFO" | grep "Module version:" | sed 's/Module version: //')

# Create the report from template
cat <<EOF > "$OUTPUT_FILE"
---
name: "🐞 Bug Report"
about: "Report an issue to help the project improve."
title: "$TITLE"
labels: ["type: bug"]
assignees: 

---

# **🐞 Bug Report**

## **Describe the bug**
$DESCRIPTION

---

### **To Reproduce**
$REPRODUCTION_STEPS

---

### **Expected behaviour**
$EXPECTED_BEHAVIOR

---

### **Media prove**
N/A (Generated automatically)

---

### **Your environment**

* OS: $OS
* Golang version: $GO_VERSION
* module version: $MODULE_VERSION

---

### **Additional context**
$ADDITIONAL_CONTEXT
EOF

echo "Success: Bug report generated at $OUTPUT_FILE"
