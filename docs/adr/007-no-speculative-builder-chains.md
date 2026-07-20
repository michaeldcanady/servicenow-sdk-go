# ADR 007: Don't scaffold a builder chain ahead of an implemented operation

## Status

Accepted

## Context

`Cdm().Policies().Mappings().Inputs().Resolved()` could be constructed —
both `Inputs()` and `Resolved()` existed as request-builder accessors — but
neither had a verb method (`Get`/`Post`/etc.) or a `To*RequestInformation`
method. The resolved-inputs endpoint was advertised in the fluent chain and
in docs/snippets, but was actually unreachable: calling it did nothing,
because there was nothing to call. `Inputs()` itself existed only as a
stepping stone to reach `Resolved()` (#489).

Alternatives considered:

1. **Implement the missing verb method(s) to make the chain functional** —
   rejected for this PR specifically because the endpoint's actual behavior
   wasn't yet scoped; forcing an implementation just to match the existing
   (accidental) shape risked guessing at the wrong contract.
2. **Leave the dead-end chain in place as a documented "coming soon"** —
   rejected; an SDK builder chain that compiles and can be navigated but
   silently does nothing on every verb is worse than a compile error, and
   contradicts the nil-receiver-guard principle in ADR-006 of failing loud
   and early rather than silently.
3. **Remove the dead-end chain entirely** (`Inputs()`/`Resolved()` and their
   builder types, 177 lines) until the endpoint is implemented additively.

## Decision

Remove builder-chain segments that have no reachable verb operation, rather
than leaving them in place as a stub. A request-builder accessor should only
exist once the operation(s) it leads to are actually implemented — not
speculatively, to reserve a URL segment or match an API's documented shape
in advance.

## Consequences

- **Pros:** every reachable point in a fluent chain does something; a
  builder existing is a reliable signal that its operations work, matching
  the "fail loud" spirit of ADR-006 rather than silently advertising
  unreachable surface area.
- **Cons:** breaking removal for anyone who had already started depending on
  the chain shape (even though no verb ever worked). Re-adding the
  Inputs/Resolved endpoint later is an additive change, not a revert.
- **Rule for new modules and reviews:** don't scaffold a full URL-segment
  chain "for completeness" ahead of the operation being implemented. Add the
  builder when the verb method(s) it exposes are ready to ship together.
