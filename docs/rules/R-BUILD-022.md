---
id: R-BUILD-022
applies_to: ["core/**/*.go", "*api/**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Request-builder receiver checks go through `core.ValidateBuilder`

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Provide one shared `core.ValidateBuilder(builder any) error` that checks a request builder's receiver and embedded `RequestBuilder`.
MUST: Return `snerrors.ErrNilRequestBuilder` by identity for every invalid input, and never panic.
MUST NOT: Fold the request-adapter check into `core.ValidateBuilder`; adapter checks stay at their call-site position.
MUST NOT: Migrate `core.BaseRequestBuilder` to the helper; it has no embedded `RequestBuilder` field.
SHOULD: Call `core.ValidateBuilder` in new request-builder methods instead of hand-writing the receiver guard.
MAY: Migrate existing hand-written guards package by package.

## Verify

```bash
grep -qE '^func ValidateBuilder\(builder any\) error' core/request_builder_guard.go
```

## Rationale

The receiver and embedded-builder guard is the most repeated idiom in the SDK, and each hand-typed copy can return the wrong sentinel. One tested helper replaces them; see [ADR 022](../../website/docs/contributing/adrs/022-request-builder-guard-helper.md).
