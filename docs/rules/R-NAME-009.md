---
id: R-NAME-009
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: RequestBuilder, RequestInformation, kiota naming, facade,
rename, parity, settled through v3 -->

MUST keep `RequestBuilder`/`BaseRequestBuilder` and `RequestInformation` named as they are.
MUST NOT rename these types for internal precision.
This is settled through v3 — revisit only at a new major version boundary with a concrete consumer-facing reason, not internal-precision cleanup.

## Verify

```bash
# Fails (non-zero) if a rename is attempted (e.g. a new `*Facade` type).
! grep -RInE 'type [A-Za-z]*Facade\b' --include='*.go' .
```

## Rationale

The names mirror `kiota-abstractions-go` and msgraph-sdk-go, the familiarity ADR-003 deliberately preserves; the imprecision (`RequestBuilder` is a façade, `RequestInformation` is the real builder) is well-precedented and costs nothing (see ADR 009).
