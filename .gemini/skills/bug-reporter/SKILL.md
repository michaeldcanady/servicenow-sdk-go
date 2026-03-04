---
name: bug-reporter
description: A skill to help users report bugs in the servicenow-sdk-go repository. It gathers system information, reproduction steps, formats the report into a GitHub issue template, and offers to submit it via the GitHub CLI.
---

# 🐞 Bug Reporter

This skill assists in creating high-quality bug reports for the `servicenow-sdk-go` SDK.

## Workflow

1.  **Gather Context**: When a user reports a bug, identify the core issue, expected vs. actual behavior, and steps to reproduce.
2.  **Collect System Info**: Run the `scripts/gather_info.sh` script to collect OS, Go version, and SDK version information.
3.  **Identify Reproduction Steps**: If the user hasn't provided clear steps, ask for them or attempt to derive them from the current session's history.
4.  **Format Report**: Use the `references/bug_report_template.md` as a template to structure the final bug report.
5.  **Submit to GitHub (Optional)**: If the GitHub CLI (`gh`) is available and authenticated, offer to submit the report automatically. Use the `scripts/submit_report.sh` script to handle the submission process.
6.  **Present to User**: If automatic submission is not possible or desired, provide the completed bug report in a Markdown code block, ready for the user to copy and paste into a new GitHub issue.

## Bundled Resources

### Scripts

- `scripts/gather_info.sh`: Gathers basic system information (OS, Go, SDK version).
- `scripts/submit_report.sh`: Submits a bug report to GitHub using the `gh` CLI. Requires a title and a path to a report file.

### References

- `references/bug_report_template.md`: The official GitHub issue template for bug reports in this repository.

## Triggering Examples

- "I found a bug in the attachment API."
- "Create a bug report for this error."
- "Report this issue on GitHub."
