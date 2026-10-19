---
id: R-SURF-017
applies_to: ["**/*now_request_builder.go", "website/docs/**"]
status: proposed
supersedes: []
conflicts_checked: true
---

# The public surface is domain-first with a raw escape hatch

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- Scope: URL-mirroring builders (R-SURF-019) are the raw layer. Domain affordances sit above them and may use semantic names. -->

MUST: Provide a domain-oriented API for each core workflow (CRUD, query, pagination, attachments) so callers express intent without knowing ServiceNow's REST paths.
MUST: Provide a documented escape hatch from every domain affordance to the raw ServiceNow surface (raw table access, raw query parameters, raw response envelopes).
MUST NOT: Require a caller to know ServiceNow's URL structure for a core workflow.
MUST NOT: Close off the raw surface.
SHOULD: Ship domain APIs for the core workflows first and grow outward.
SHOULD NOT: Add a speculative domain API ahead of a concrete workflow (R-BUILD-007).

## Verify

```bash
# The raw-table escape hatch (R-SURF-018) must exist. Fails closed (exit 1) if absent.
grep -RqE 'func \([a-zA-Z]+ \*?[A-Za-z]+\) ByName\(' --include='*.go' .
```

## Rationale

A thin wrapper fails the premise because callers still need ServiceNow paths, and a fully opinionated SDK cannot cover ServiceNow's breadth. A hybrid gives domain intent for common cases with an always-reachable raw path. Consumer-defined typed tables are covered by R-TBL-019 and R-TBL-024.
