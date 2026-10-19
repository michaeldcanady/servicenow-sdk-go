---
id: R-PAGE-005
applies_to: ["core/**/*.go", "*api/**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Pagination has one generic entry point

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Expose pagination only through `core.NewPageIterator[T]`.
MUST NOT: Add a `New<X>PageIterator` wrapper for an API module.
SHOULD: Document, per endpoint, whether it emits `Link` headers.

## Verify

```bash
# grep exit 1 = no match = pass; exit 2 fails closed.
grep -RInE 'func New[A-Za-z]+PageIterator' --include='*.go' --exclude-dir=core --exclude-dir=.git .; [ $? -eq 1 ]
```

## Rationale

Per-module wrappers duplicate surface without adding capability, and they imply pagination support the endpoint may lack, because the iterator follows `Link` headers that not every endpoint emits. msgraph-sdk-go-core likewise exposes a single generic iterator.
