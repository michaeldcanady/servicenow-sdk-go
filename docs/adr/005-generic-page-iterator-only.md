# ADR 005: One generic `core.PageIterator`, no per-module wrappers

## Status

Accepted

## Context

`tableapi` and `attachmentapi` each shipped their own one-line pagination
entry point (`tableapi.NewTablePageIterator`, `tableapi.NewDefaultTablePageIterator`,
`attachmentapi.NewAttachmentPageIterator`) that did nothing but pre-bind the
module's `Create...FromDiscriminatorValue` factory to the generic
`core.NewPageIterator[T]`. No other collection-bearing module had one, so the
SDK had three ways to spell the same operation depending on which module you
were in (#490).

Alternatives considered:

1. **Give every collection-bearing module its own wrapper** — rejected.
   `core.PageIterator` follows response `Link` headers, and not every
   ServiceNow collection endpoint emits them; a `NewDefault<X>PageIterator`
   for every module would advertise pagination an endpoint may not actually
   support. It also means auditing and maintaining N near-identical
   constructors forever, and the wrappers only help the default record type
   anyway — a caller using a custom type already passes a factory, at which
   point the wrapper saves nothing.
2. **Generic-only** — msgraph-sdk-go-core exposes exactly one generic
   `PageIterator`; no generated per-module wrappers. Callers arriving from
   other Kiota SDKs already know this shape (ADR-003's tie-breaker rule).

## Decision

Remove the three module-specific wrappers. `core.NewPageIterator[T]` is the
single, only documented pagination pattern:

```go
iterator, err := core.NewPageIterator(response, client.GetRequestAdapter(),
    tableapi.CreateTableRecordFromDiscriminatorValue)
```

## Consequences

- **Pros:** one pattern to teach, document, and review, regardless of
  module; no risk of a wrapper implying pagination support an endpoint
  doesn't have.
- **Cons:** breaking removal of the three wrapper constructors — callers
  must migrate to the generic call. Landed pre-v2.0 for that reason.
- **Rule for new modules:** do not add a `New<X>PageIterator` wrapper for a
  new API module, even for symmetry with `tableapi`'s old shape — that
  shape was removed deliberately, not left behind by omission. Document
  per-endpoint whether it actually emits `Link` headers instead.
