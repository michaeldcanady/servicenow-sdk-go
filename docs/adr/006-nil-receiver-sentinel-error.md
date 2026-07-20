# ADR 006: Nil-receiver guards return a sentinel error, never `(nil, nil)`

## Status

Accepted

## Context

Every request-builder verb method (`Get`/`Post`/`Patch`/`Put`/`Delete`)
starts with a nil-receiver guard (see `CLAUDE.md`'s constructor-triad
section). Historically these guards returned `return nil, nil` — no result
and no error — when the receiver or its embedded `RequestBuilder` was nil.
`Delete`/`Patch` variants that only return an error swallowed the problem
entirely with a bare `return nil`.

This meant a caller holding a nil builder (e.g. from a failed prior chain
step) saw silent success and only failed later, at a point in the code far
from the actual mistake, with no error to grep for or wrap.

Alternatives considered:

1. **Leave `(nil, nil)` / bare `nil`** — rejected; it's a silent-failure
   trap that pushes the debugging cost onto whoever hits the downstream
   symptom.
2. **A new, verb-specific error per package** — rejected; it would multiply
   the "three separate error-sentinel locations" problem CLAUDE.md already
   warns about, and give every module its own wording for the same failure.
3. **One shared sentinel in `errors/errors.go`**, returned by every guard
   across every module — matches ADR-001's centralization principle and
   lets callers use `errors.Is` uniformly regardless of which module or verb
   they called.

## Decision

Added `snerrors.ErrNilRequestBuilder` to `errors/errors.go`. Every
nil-receiver guard across every `*api` package, `core.BaseRequestBuilder`
setters, and the generic `appserviceapi` post builder now returns this
sentinel (or the package's zero value + sentinel for single-error-return
verbs) instead of `(nil, nil)` / bare `nil`.

Navigation methods that return a bare `*Builder` (no error), and
nil-response/null-element semantics elsewhere (`res == nil`,
`ElementValue.IsNil`), are unchanged — this decision is scoped to
verb-method nil-*receiver* guards specifically, not every nil check in the
codebase.

## Consequences

- **Pros:** a nil builder now fails loudly and immediately, at the call
  site, with an error identity (`errors.Is(err, snerrors.ErrNilRequestBuilder)`)
  every caller can check the same way regardless of module.
- **Cons:** breaking — any caller relying on `(nil, nil)` as a success-shaped
  no-op must be updated to handle the returned error.
- **Enforced by:** the `api-module-consistency-reviewer` agent and the
  `new-api-module` skill both check for this pattern; a new module's guard
  that returns bare `nil`/`(nil, nil)` instead of the sentinel is a
  consistency-review finding, not a style nit.
