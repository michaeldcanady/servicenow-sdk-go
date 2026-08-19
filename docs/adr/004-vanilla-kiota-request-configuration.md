# ADR 004: Vanilla Kiota request-configuration pattern for query parameters

## Status

Accepted

## Context

Every API module exposes a `*QueryParameters` struct per verb (e.g.
`TableRequestBuilderGetQueryParameters`) that gets applied to the outgoing
request. Historically these used plain value fields (`Limit int`,
`Query string`, `DisplayValue DisplayValue`) and were encoded onto the
request via a repo-specific `internal.ConfigureRequestInformation` /
`internal.KiotaRequestInformation` wrapper built on the `go-querystring`
library.

This diverged from how Kiota-generated SDKs (e.g. msgraph-sdk-go) shape query
parameters, and the divergence had a real cost: value fields can't represent
"the caller didn't set this" separately from "the caller explicitly set the
zero value," so `go-querystring` either always emitted a field or needed
extra `omitempty`-style conventions layered on top to approximate what
`abstractions.RequestInformation.AddQueryParameters` gives you for free with
nil-checkable pointers.

Alternatives considered:

1. **Keep the value-typed structs and the go-querystring wrapper** — rejected
   because it's the thing ADR-003 already argues against: reinventing a
   Kiota runtime capability (`AddQueryParameters`) that ships for free, and
   it keeps the SDK's query-parameter shape looking different from every
   Kiota-generated SDK a caller might already know.
2. **Pointer fields wired through `abstractions.ConfigureRequestInformation`
   directly** — matches msgraph-sdk-go's convention exactly and removes a
   dependency.

A timeboxed spike (#494) was run specifically to classify this as breaking
vs. additive before the v2.0 cut, since it could not land after the tag
without becoming a v3 change.

## Decision

Every exported `*QueryParameters` struct across every API module uses pointer
fields with `uriparametername` tags (`Limit *int32`, `Query *string`,
`DisplayValue *DisplayValue`, etc.), and request builders call
`abstractions.ConfigureRequestInformation` directly instead of a bespoke
wrapper. The `internal.ConfigureRequestInformation` / `internal.KiotaRequestInformation`
wrapper and the `go-querystring` dependency are removed entirely (#511).

A non-obvious constraint this decision carries: the native encoder only
recognizes `*int32` explicitly — a bare `*int` (even as a pointer) is
silently dropped from the request. Any new integer query parameter must be
declared as `*int32`, not `*int`, or it will vanish without error.

## Consequences

- **Pros:** query-parameter structs now look and behave like every other
  Kiota-generated SDK's `*QueryParameters`; "unset" and "zero" are
  distinguishable the same way ADR-002 already made models' properties
  distinguishable; one fewer third-party dependency.
- **Cons:** breaking for every caller constructing these structs with value
  literals — they must switch to pointers (see `internal.ToPointer`).
  Landed as a breaking v2.0 change specifically because it could not wait
  for v3.
- **Sharp edge to guard in review:** a new `*int` query-parameter field is a
  silent bug, not a compile error — it must be `*int32`.
