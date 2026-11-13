---
title: 'ADR 022: A shared request-builder guard in core replaces hand-rolled receiver checks'
description: >-
  One total, panic-free ValidateBuilder helper owns the receiver/embedded-builder
  precondition; adapter checks stay at their call-site position.
---

# ADR 022: A shared request-builder guard in `core` replaces hand-rolled receiver checks

## Status

Proposed. Accepted on merge; this record and its rule file
(`docs/rules/R-BUILD-022.md`) ship together with the helper.

## Context

Every request-builder method opens with the same two-line precondition
(CLAUDE.md's constructor-triad section):

```go
if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
    return nil, snerrors.ErrNilRequestBuilder // ADR-006
}
```

That pair is repeated verbatim across every `*api` package's verb,
navigation, and `To*RequestInformation` methods — dozens of copies of one
idea (Epic #568). The repetition is not free: each copy is a place to get the
sentinel wrong (ADR-006 exists because some historically returned `(nil, nil)`),
and the embedding contract (`struct { core.RequestBuilder }`) is enforced only
by convention and the consistency reviewer.

Alternatives considered:

1. **Keep per-method inline guards** — rejected; the duplication is the bug
   class. Every new module re-types the idiom by hand.
2. **An all-in-one guard that also checks the request adapter** — rejected;
   several verbs interleave `requestConfiguration`/query-parameter checks
   between the builder guard and the adapter guard. Folding them together
   silently reorders which sentinel a caller holding multiple nils receives,
   which would make the #633 migration a behaviour change instead of a
   refactor.
3. **A generics/closure-based helper taking lazy accessors** — rejected; the
   closure boilerplate per call site exceeds the lines it saves.
4. **A method on `core.BaseRequestBuilder`** — impossible; a typed-nil outer
   receiver cannot be detected from inside an embedded value, which is exactly
   the failure this guard exists to catch.
5. **One `core.ValidateBuilder(builder any) error` owning only the
   always-paired receiver + embedded check** — accepted.

## Decision

`core.ValidateBuilder` (in `core/request_builder_guard.go`) enforces the
receiver + embedded-builder precondition for any input:

- accepts pointers to structs and struct values; everything else — untyped
  nil, typed-nil pointers (including typed nils inside interfaces), pointers
  to non-structs, multi-level pointers — reports
  `snerrors.ErrNilRequestBuilder`;
- never panics. Reflection traversal uses `IsNil`/`Elem`, never
  `Value.Interface` on fields reflect may consider unexported, and follows
  anonymous struct / non-nil anonymous-pointer promotions to at most two
  levels. Promotion through a nil intermediate pointer reports invalid rather
  than followed — the legacy idiom panics on exactly those shapes;
- the embedded field must be named `RequestBuilder` and hold a non-nil
  interface with a non-nil concrete value; impostor kinds (map, chan, func,
  concrete struct) do not satisfy the contract;
- returns the shared sentinel **by identity** so `errors.Is` works uniformly
  across modules (ADR-001/ADR-006 rules).

The request-adapter precondition deliberately remains at call sites, keeping
its existing position between other checks.

## Consequences

- **Pros:** one tested implementation of the most-repeated idiom in the SDK;
  the totality and differential-oracle properties are proven by an adversarial
  suite (`request_builder_guard_test.go`,
  `request_builder_guard_adversarial_test.go`) rather than asserted; future
  modules stop hand-typing the guard (#632 exit criteria feeds the
  `new-api-module` skill).
- **Cons:** the embedding contract ("field literally named `RequestBuilder`,
  interface kind") becomes load-bearing in runtime code, not just convention;
  reflection costs a small constant per call, irrelevant next to any HTTP
  round-trip.
- **Migration:** #633 swaps call sites package-by-package. For sites using the
  canonical two-line idiom, the helper agrees exactly with the legacy code on
  every shape where the legacy idiom terminates (differential-oracle tested),
  so migration there is semantics-preserving by construction. Two site
  families are NOT covered by that guarantee and must be decided explicitly
  during #633: (a) sites using `IsNil(rB)` alone — today they panic later at
  first field use when only the embedded builder is nil; swapping to the guard
  converts that panic into the sentinel (a behaviour change needing a changelog
  note); (b) `core.BaseRequestBuilder` itself — it has no embedded
  `RequestBuilder` field and must be EXCLUDED from migration entirely.
  `To*RequestInformation` methods without any guard today gain a first-ever
  one under the same panic-to-error caveat.
- **Enforced by:** `docs/rules/R-BUILD-022.md`; an inline re-implementation of
  the guard in a request-builder method is a review finding once this ADR is
  accepted.

See also: ADR 001 (error standardization), ADR 006 (nil-receiver sentinel).
