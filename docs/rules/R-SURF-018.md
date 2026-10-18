---
id: R-SURF-018
applies_to: ["now_request_builder.go", "core/**/*.go", "**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

<!-- Keywords: Table2, collection builder, ByName, table entry, raw table,
method overload, number suffix, url segment, rename to Table -->

MUST expose the raw-table escape hatch as `.ByName(name)` on the collection builder returned by `Table2()`.
MUST name the v2 collection-builder method `Table2()` and rename it to `Table()` at the next major once the legacy `Table(name string)` is removed.
MUST keep the legacy `Now().Table(name string) *TableRecord` working for all of v2.
MUST map a builder method name to its URL path segment; on a name collision append a number (`Table2`).
MUST NOT remove the legacy `Table(name string)` before the next major version.
MUST NOT attempt to overload `Table` for both shapes — Go has no method overloading.
SHOULD document the URL mapping for scoped APIs that keep semantic (non-path) names.

## Verify

```bash
# Passes (zero) while the v2 collection-builder entry exists.
grep -RInE 'func .*Table2\(' --include='*.go' .
```

## Rationale

Go can't overload methods, so the collection-builder entry and the legacy `Table(name)` share v2 under a number-suffixed `Table2`, renamed to `Table` at the next major; `.ByName` preserves the raw escape hatch and the name still mirrors the URL segment (see ADRs 008, 018).
