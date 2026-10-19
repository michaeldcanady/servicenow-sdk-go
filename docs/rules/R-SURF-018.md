---
id: R-SURF-018
applies_to: ["now_request_builder.go", "tableapi/**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Table2 and ByName bridge the v2 table entry point

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- TODO(owner): Kiota convention for index builders is `ByTableName`; confirm `ByName` is intentional (R-BUILD-003 makes msgraph-sdk-go the tie-breaker). -->

MUST: Expose the raw-table escape hatch as `.ByName(name)` on the collection builder returned by `Table2()`.
MUST: Name the v2 collection-builder method `Table2()`.
MUST: Rename `Table2()` to `Table()` at the next major, once the legacy `Table(name string)` is removed.
MUST: Keep the legacy `Now().Table(name string) *TableRecord` working for all of v2.
MUST: Mark the legacy `Table(name string)` with a `// Deprecated:` paragraph naming `Table2().ByName` as the replacement (R-REL-025).
MUST NOT: Remove the legacy `Table(name string)` before the next major version.
SHOULD NOT: Use number-suffixed builder names outside this Table2 transition.
SHOULD: Document the URL mapping for scoped APIs that keep semantic (non-path) names.

## Verify

```bash
# Fails closed (exit 1) if Table2 is missing or the legacy Table(name) lacks a Deprecated: paragraph.
grep -RqE 'func \([a-zA-Z]+ \*?[A-Za-z]+\) Table2\(' --include='*.go' . \
  && grep -RIn -B6 -E 'func \([a-zA-Z]+ \*?[A-Za-z]+\) Table\(name string\)' --include='*.go' . | grep -q 'Deprecated:'
```

## Rationale

Go cannot overload methods, so the collection-builder entry and the legacy `Table(name)` share v2 under a number-suffixed `Table2()`, renamed to `Table()` at the next major. `.ByName` preserves the raw escape hatch, and the name still mirrors its URL segment.
