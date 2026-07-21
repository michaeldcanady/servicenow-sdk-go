# ADR 009: Keep `RequestBuilder`/`RequestInformation` naming (Kiota parity)

## Status

Accepted

## Context

Backlog grooming for v2.0 (#496) flagged that `RequestBuilder` and
`RequestInformation` are, strictly speaking, misnamed: the type most callers
chain against (`core.RequestBuilder` / `core.BaseRequestBuilder` and the
per-module `*RequestBuilder` types) is really a fluent **facade** over the
API surface, while `RequestInformation` — built up inside each `ToXRequestInformation`
method and passed to `RequestAdapter.Send*` — is the object that actually
accumulates URL, headers, query parameters, and body, i.e. the real
**builder** in the classic sense.

This naming is inherited directly from Kiota's runtime abstractions
(`kiota-abstractions-go`) and from Kiota-*generated* SDKs like
`msgraph-sdk-go`, per [[003-hand-written-on-kiota]] and [[008-package-symbol-url-naming-independence]]'s underlying premise that
matching Kiota/msgraph conventions is a stated design goal so developers
arriving from other Kiota SDKs find familiar idioms.

Renaming these types is only possible as a breaking change, and v2.0 is the
last point before v3 where a rename would be low-friction. The decision
therefore has to be made deliberately now, not left implicit.

Alternatives considered:

1. **Rename to reflect actual roles** (e.g. `RequestBuilder` → something
   like `Client`/`Facade`, `RequestInformation` → `RequestBuilder`) —
   rejected. This breaks the naming correspondence with
   `kiota-abstractions-go` and every Kiota-generated SDK (msgraph-sdk-go and
   friends), which is exactly the familiarity ADR-003 says this SDK is
   trying to preserve. It would also touch nearly every exported type across
   every `*api` package, for a purely internal-precision gain with zero
   consumer-facing benefit.
2. **Keep the existing names** — matches the ecosystem convention this SDK
   deliberately mirrors; the imprecision is real but well-precedented (Kiota
   itself uses these names the same way) and costs nothing beyond an
   occasional need to explain it.

## Decision

Keep `RequestBuilder`/`BaseRequestBuilder` and `RequestInformation` named as
they are. No rename is scoped for v2.0 or planned for v3.

## Consequences

- **Pros:** naming stays consistent with `kiota-abstractions-go` and
  Kiota-generated SDKs (msgraph-sdk-go and peers); no churn across every
  `*api` package's exported surface; nothing to migrate for consumers.
- **Cons:** the names remain technically imprecise — `RequestBuilder` is a
  facade, `RequestInformation` is the actual builder — which can be mildly
  confusing on first read until a contributor connects it to the Kiota
  convention it mirrors.
- **Rule for future naming questions on these types:** this is now settled
  through v3 — do not revisit without a new major version boundary and a
  concrete consumer-facing reason, not just internal-precision cleanup.
