---
id: R-BUILD-004
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: query parameters, pointer fields, *int32, go-querystring,
ConfigureRequestInformation, RequestInformation, uri param -->

MUST declare exported `*QueryParameters` fields as pointers.
MUST declare integer query parameters as `*int32` — a bare `*int` is silently dropped by the Kiota encoder.
MUST wire query parameters through `abstractions.ConfigureRequestInformation` directly.
MUST NOT reintroduce a `go-querystring` dependency or a bespoke `ConfigureRequestInformation` wrapper.
MUST NOT use value-typed (non-pointer) query-parameter fields.

## Verify

```bash
# Fails (non-zero) if a bare `*int` (not `*int32`) appears — a silent bug.
! grep -RInE '\*int\b' --include='*.go' .
```

## Rationale

Pointer fields wired through `ConfigureRequestInformation` match msgraph-sdk-go and let "unset" differ from "zero," the same distinction ADR-002 gave models; the `*int32` requirement is a non-obvious encoder constraint (see ADR 004).
