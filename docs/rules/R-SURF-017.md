---
id: R-SURF-017
applies_to: ["**/*.go", "website/docs/**"]
status: proposed
supersedes: []
conflicts_checked: true
---

<!-- Keywords: hybrid surface, domain first, north star, escape hatch, raw
access, intent, byName, user defined sdk, core workflows -->

MUST present a domain-oriented API for each core workflow (CRUD, query, pagination, attachments) so a caller can express intent without knowing ServiceNow's REST paths.
MUST keep an escape hatch from every domain affordance to the raw ServiceNow surface (raw table access, raw query parameters, raw response envelopes).
MUST let consumers build their own strongly-typed table SDKs on the exported foundation with parity to the built-in tables.
MUST NOT require a caller to know ServiceNow's URL structure for a core workflow.
MUST NOT close off the raw surface — every domain affordance needs a documented raw path.
SHOULD ship domain APIs for the core workflows first and grow outward.
SHOULD NOT add a speculative domain API ahead of a concrete workflow (see ADR 007).

## Verify

```bash
# Passes (zero) while the raw escape hatch is present; losing it regresses
# the hybrid promise.
grep -RInE 'ByName' --include='*.go' .
```

## Rationale

A thin wrapper fails the premise (callers still need ServiceNow paths) and a fully-opinionated SDK can't cover ServiceNow's breadth, so the surface is hybrid: domain intent for common cases with an always-reachable raw escape hatch (see ADR 017).
