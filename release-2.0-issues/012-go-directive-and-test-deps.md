# Consumer-facing go.mod: minimum Go version and test-only dependencies

- **Priority:** P2
- **Raised by:** Senior Principal Engineer
- **Area:** SDK design / dependency hygiene

## Problem

1. **`go 1.25.0` minimum.** The go.mod directive forces every consumer onto Go ≥ 1.25.
   Enterprise ServiceNow shops (the target audience) often trail several releases.
   Unless the code needs a 1.25 language feature or stdlib API, a lower directive
   (matching CI's `oldstable`, and ideally one more back) widens adoption at zero cost.
   Note the CI matrix already tests `oldstable`, so the floor should be an explicit,
   tested decision, not whatever the dev machine had.

2. **Test-only dependencies in the main module's require block**: `cucumber/godog`,
   `jarcoal/httpmock`, `joho/godotenv` exist solely for `tests/` (build-tag-gated) and
   unit-test mocks. Since Go 1.17 module-graph pruning, consumers don't *build* these,
   but they still appear in the SDK's go.mod/go.sum, inflate `go mod graph` output,
   trigger consumers' dependency scanners (a godog CVE would page every consumer's
   security team), and show on pkg.go.dev's imports list.

## Status update (2026-07-18)

Item 1 was attempted and is **not actionable**: `kiota-http-go` (and its otel
dependencies) declare `go 1.25.0`, so `go mod tidy` forces the directive back to
1.25.0. The floor can only drop by downgrading kiota-http-go, which isn't worth it.
Item 2 (nested test module) remains open.

## Recommendation

1. ~~Choose and document a support policy — e.g. "the two most recent Go releases" — and
   set the go directive to the older of the two (drop the patch suffix: `go 1.24`).~~
   Blocked by kiota-http-go's own `go 1.25.0` requirement; document "requires Go ≥ 1.25"
   in the Readme instead.
2. Move `tests/integration` and `tests/e2e` into their own nested module
   (`tests/go.mod`) with a `replace` or workspace (`go.work`) for local dev. That
   removes godog/godotenv/httpmock from the SDK's module. `testify` stays (used by
   co-located unit tests — acceptable and conventional).
3. Add the nested module to CI's tidy-check and test jobs (see issue 009).
