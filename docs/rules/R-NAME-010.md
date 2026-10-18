---
id: R-NAME-010
applies_to: ["go.mod", "docs/**", "README.md"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: module path, org migration, michaeldcanady, nerdit, import path,
go get, semantic import versioning -->

MUST keep the module path at `github.com/michaeldcanady/servicenow-sdk-go`.
MUST change only the semantic-import-versioning suffix (`/vN`) at a major bump, never the org/owner segment.
MUST NOT migrate the repo to a different GitHub org for v2 or v3.

## Verify

```bash
# Passes (zero) only while the module path is correct; a wrong org fails.
grep -qE '^module github.com/michaeldcanady/servicenow-sdk-go(/v[0-9]+)?$' go.mod
```

## Rationale

An org migration changes the import path and is inherently breaking; v2 already forces a `/vN` bump, but no concrete driver for an org move materialized, so declining it costs nothing and keeps the release runbook simpler (see ADR 010).
