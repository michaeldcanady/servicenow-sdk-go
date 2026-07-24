# Nil-receiver guards return `nil, nil`, silently swallowing bugs

- **Priority:** P1 — API design; 2.0 is the only window to fix it
- **Raised by:** Senior Principal Engineer
- **Area:** SDK design / public API contract
- **Status (2026-07-24):** Original defect fixed (ADR-006 / PR #487). This doc's
  citations were stale as of a follow-up audit and have been corrected below — the
  remaining gap is now tracked as a separate GitHub issue, not in this doc.

## Problem

The standard verb-method preamble across every `*api` package used to be:

```go
if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
    return nil, nil
}
```

A caller who accidentally holds a nil builder got `(nil, nil)` — no result **and no
error**. The inevitable outcome is a confusing downstream nil-pointer dereference far
from the actual mistake, or worse, code that treats "no error" as success. This
violates the Go convention that a nil error means the operation succeeded and the
result is usable.

The adapter guard right next to it already did the right thing
(`return nil, snerrors.ErrNilRequestAdapter`), so the codebase was internally
inconsistent about how nil-guards behave.

## Recommendation (as originally written — now implemented)

Since 2.0 is a breaking release, change the contract now:

1. Add a shared sentinel to `errors/errors.go` (e.g. `ErrNilRequestBuilder` — check
   whether one already exists before adding, per the sentinel-duplication issue 005).
2. Replace every `return nil, nil` nil-receiver guard with
   `return nil, snerrors.ErrNilRequestBuilder`.
3. Update the corresponding tests (many currently assert `nil, nil`).
4. Update CLAUDE.md and the `new-api-module` skill / `api-module-consistency-reviewer`
   agent so new modules follow the corrected pattern.

## Verification update (2026-07-24)

This doc originally cited `tableapi/table_request_builder.go:76,120,164` as still
returning bare `nil, nil`. **That citation is now stale/wrong** — those exact lines
were fixed by ADR-006 / PR #487 and today correctly return
`nil, snerrors.ErrNilRequestBuilder`. `tableapi`, `caseapi`, `appserviceapi`,
`actsubapi`, `cmdbinstanceapi`, `policyapi`, and `appointmentbookingapi` were sampled
and confirmed to have the guard consistently, matching ADR-006's stated contract.

The underlying defect class (nil receiver → wrong/missing error) is **not** fully
closed, though: `cdmapplicationsapi`, `cdmchangesetapi`, and `cdmeditorapi` have
**zero** nil-receiver guards on any verb method — these three packages don't even
have the old `nil, nil` fallback, they nil-pointer-panic instead. ADR-006's claim
that the guard was "applied across every `*api` package" is therefore inaccurate for
these three packages today; the ADR text itself doesn't need editing, just the code
catching up to it.

This remaining gap is a materially different (and worse — panic, not a wrong return
value) bug than what this doc originally tracked, so it's filed separately as
[#551](https://github.com/michaeldcanady/servicenow-sdk-go/issues/551) rather than
reopening this doc's scope. Recommend treating this doc as closed/historical once
#551 lands.
