# Consolidate the three error-sentinel locations before the API freezes

- **Priority:** P1 — breaking-change window closes at 2.0
- **Raised by:** Senior Principal Engineer
- **Area:** SDK design / error handling
- **Status (2026-07-24):** Confirmed still open by a follow-up audit against
  `release/2.0` @ `aca3942`, with new concrete duplicate-name examples added below.

## Problem

Sentinel errors live in three places with similar-but-different wording:

1. Root `errors.go` (`package servicenowsdkgo`) — client-config sentinels.
2. `errors/errors.go` (`snerrors`) — the intended shared set (`ErrNilRequestAdapter`,
   `ErrNilResponse`, `ErrNilConfig`, `ErrNilBody`, ...).
3. Package-local files such as `tableapi/errors.go` with their own variants.

On top of that, many historical call sites return a fresh `errors.New("...")` with text
matching a sentinel instead of the sentinel itself, which breaks `errors.Is` for
callers. Once v2.0.0 ships, removing or re-homing any exported sentinel is a new
breaking change, so this cleanup is now-or-wait-for-v3.

## Confirmed concrete examples (2026-07-24 audit)

Same identifier name, two different sentinel *values* (different message text),
exported from two different importable packages — this is the exact shape that
breaks `errors.Is` for a consumer who imports both:

- `ErrNilContext`:
  - Root `errors.go:7` — `errors.New("ctx cannot be nil")`
  - `errors/errors.go:10` (`snerrors.ErrNilContext`) — `errors.New("context cannot be
    nil")`
- `ErrNilResponse`:
  - `tableapi/errors.go:9` — `errors.New("response can't be nil")`
  - `errors/errors.go:9` (`snerrors.ErrNilResponse`) — `errors.New("response cannot
    be nil")`

Additionally, `attachmentapi/errors.go` is a **fourth** micro-location for sentinels
(e.g. `ErrNilParams`) that was not in the original three-location inventory above and
should fold into the same consolidation pass rather than being missed.

## Recommendation

1. Inventory: `grep -rn 'errors.New(' --include='*.go' . | grep -v _test` and diff the
   messages against `errors/errors.go`.
2. Make `errors/errors.go` the single home for cross-package sentinels; deprecate (or
   delete, since v2 may break) duplicates in `tableapi/errors.go`, `attachmentapi/errors.go`,
   and root `errors.go` — keep package-local sentinels only where the condition is
   genuinely package-specific.
3. Replace inline `errors.New` duplicates with the shared sentinel by identity.
4. Add a unit test that asserts, for each public verb-method failure mode, that
   `errors.Is(err, snerrors.ErrX)` holds — locking the contract for consumers.

## Cross-references

- Issue 004 adds `ErrNilRequestBuilder`; do these together to avoid two sweeps. (That
  addition has since shipped for most `*api` packages — see the doc-004 status note
  and [#551](https://github.com/michaeldcanady/servicenow-sdk-go/issues/551) for the
  remaining CDM-package gap, which should reuse the same `snerrors.ErrNilRequestBuilder`
  sentinel this consolidation is meant to protect.)
