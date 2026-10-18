---
id: R-PAGE-005
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: page iterator, pagination, NewPageIterator, link headers,
per-module wrapper, generic iterator -->

MUST use `core.NewPageIterator[T]` as the single pagination entry point.
MUST NOT add a `New<X>PageIterator` wrapper for an API module.
SHOULD document per-endpoint whether it actually emits `Link` headers instead of adding a wrapper.

## Verify

```bash
# Fails (non-zero) if a per-module page-iterator wrapper is reintroduced.
! grep -RInE 'func New[A-Za-z]+PageIterator' --include='*.go' --exclude-dir=core .
```

## Rationale

`core.PageIterator` follows response `Link` headers that not every endpoint emits, so a per-module wrapper advertises pagination the endpoint may not support; msgraph-sdk-go-core also exposes only one generic iterator (see ADR 005).
