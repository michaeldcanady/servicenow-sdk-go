# Consolidate the three error-sentinel locations before the API freezes

- **Priority:** P1 — breaking-change window closes at 2.0
- **Raised by:** Senior Principal Engineer
- **Area:** SDK design / error handling

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

## Recommendation

1. Inventory: `grep -rn 'errors.New(' --include='*.go' . | grep -v _test` and diff the
   messages against `errors/errors.go`.
2. Make `errors/errors.go` the single home for cross-package sentinels; deprecate (or
   delete, since v2 may break) duplicates in `tableapi/errors.go` and friends — keep
   package-local sentinels only where the condition is genuinely package-specific.
3. Replace inline `errors.New` duplicates with the shared sentinel by identity.
4. Add a unit test that asserts, for each public verb-method failure mode, that
   `errors.Is(err, snerrors.ErrX)` holds — locking the contract for consumers.

## Cross-references

- Issue 004 adds `ErrNilRequestBuilder`; do these together to avoid two sweeps.
