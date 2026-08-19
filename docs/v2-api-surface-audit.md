# v2.0 Pre-GA Public API Surface Audit — Findings

**Audit date:** 2026-08-18
**Branch:** `release/2.0`
**Issue:** [#564](https://github.com/michaeldcanady/servicenow-sdk-go/issues/564)

---

## Additional finding: Version constant

`version.go:5` declares `const Version = "1.12.0"`. This must be updated before the
v2.0.0 tag is cut. The `release-please` automation may handle this, but it should be
verified — if the module path changes to `/v2` (blocker #557), the version constant
should reflect `2.0.0`.
