# Nil-receiver guards return `nil, nil`, silently swallowing bugs

- **Priority:** P1 — API design; 2.0 is the only window to fix it
- **Raised by:** Senior Principal Engineer
- **Area:** SDK design / public API contract

## Problem

The standard verb-method preamble across every `*api` package is:

```go
if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
    return nil, nil
}
```

(e.g. `tableapi/table_request_builder.go:76`, `:120`, `:164`; the pattern is
codified in CLAUDE.md and repeated in dozens of builders.)

A caller who accidentally holds a nil builder gets `(nil, nil)` — no result **and no
error**. The inevitable outcome is a confusing downstream nil-pointer dereference far
from the actual mistake, or worse, code that treats "no error" as success. This
violates the Go convention that a nil error means the operation succeeded and the
result is usable.

The adapter guard right next to it already does the right thing
(`return nil, snerrors.ErrNilRequestAdapter`), so the codebase is internally
inconsistent about how nil-guards behave.

## Recommendation

Since 2.0 is a breaking release, change the contract now:

1. Add a shared sentinel to `errors/errors.go` (e.g. `ErrNilRequestBuilder` — check
   whether one already exists before adding, per the sentinel-duplication issue 005).
2. Replace every `return nil, nil` nil-receiver guard with
   `return nil, snerrors.ErrNilRequestBuilder`.
3. Update the corresponding tests (many currently assert `nil, nil`).
4. Update CLAUDE.md and the `new-api-module` skill / `api-module-consistency-reviewer`
   agent so new modules follow the corrected pattern.

This is mechanical but wide; a scripted sweep plus the consistency-reviewer agent can
verify completeness.
