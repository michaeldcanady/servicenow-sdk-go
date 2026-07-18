# Pre-GA public API surface audit

- **Priority:** P1 — the surface freezes at v2.0.0
- **Raised by:** Senior Principal Engineer
- **Area:** SDK design

## Problem

Everything exported at tag time is frozen under Go 1 compatibility expectations until
v3. The v2 rework was large (90 commits) and fast; a deliberate final pass over what is
actually exported has not happened. Known smells worth a look:

- `mocking.go` in the root package exports test doubles to consumers (see issue 006).
- The constructor triad (`New<X>RequestBuilderInternal` / `NewDefault<X>RequestBuilder`
  / `New<X>RequestBuilder`) exposes `...Internal` constructors publicly in every
  package — if they're internal, the name says so but the export doesn't.
- Backing-store accessor/mutator plumbing: verify `internal/store` types don't leak
  through exported signatures (a consumer should never need to name an `internal` type;
  the compiler forbids importing them, so a leaked one makes the API uncallable).
- Deprecated v1 surfaces: confirm nothing marked `// Deprecated:` in v1 survived into
  v2 exports unintentionally.
- Response envelopes (`core.ServiceNowItemResponse[T]` etc.): check the generic
  constraints and method sets read as intended final API, since generics choices are
  especially hard to walk back.

## Recommendation

1. Generate the surface: `go doc -all ./...` or
   `golang.org/x/exp/cmd/apidiff` v1-tag → release/2.0 to see exactly what changed and
   what's new.
2. Review with the rule "unexported by default; export only what a consumer journey
   needs." Demote `...Internal` constructors if feasible (or document why they must be
   public — Kiota-style cross-package construction may genuinely need them; then say so
   in doc comments).
3. Ensure every exported identifier has a doc comment (`staticcheck`'s ST1000/ST1020
   family or `revive`'s exported rule can enforce; currently not enabled in
   `.golangci.yml`).
4. Run the `api-module-consistency-reviewer` agent across all `*api` packages as the
   final consistency gate.
