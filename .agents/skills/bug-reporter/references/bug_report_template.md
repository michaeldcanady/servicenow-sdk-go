# Bug Report

Use this as a reference when generating bug reports. The canonical form is the
YAML issue template at `.github/ISSUE_TEMPLATE/bug-template.yaml`.

## Fields

**SDK Version** — `go list -m github.com/michaeldcanady/servicenow-sdk-go` output or go.mod pin.

**Go Version** — output of `go version`.

**Module** — closest matching package (table-api, attachment-api, core / request pipeline, etc.).

**What happened?** — clear description of the bug.

**What did you expect?** — expected behavior.

**Reproduction** — minimal Go code that triggers the bug. Redact instance URLs and credentials.

**Error output** — stack trace or error messages, if any.

**Additional context** — related issues, ServiceNow version, plugin config, etc.
