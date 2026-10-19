---
id: R-ERR-006
applies_to: ["*api/**/*_builder.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Verb methods fail loudly on a nil receiver

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Return `snerrors.ErrNilRequestBuilder` from every verb method's nil-receiver guard, with the package zero value for any other return.
MUST NOT: Return `(nil, nil)` or a bare `nil` from a verb method's nil-receiver guard.
MUST: Leave navigation methods unchanged; they return a bare child builder.
MUST: Treat a `nil` response from the request adapter (for example, no content) as success, distinct from the guard.

## Verify

```bash
# Flags `if <non-res> == nil { return nil, nil }` in builder files; `res == nil` is the legitimate no-content path.
# grep exit 1 = no match = pass; exit 2 fails closed.
grep -RPzl --include='*_builder.go' 'if (?!res\b)\w+ == nil \{\s*return nil, nil' .; [ $? -eq 1 ]
```

## Rationale

A `(nil, nil)` guard turns a nil builder into silent success that fails later, far from the mistake. A shared sentinel makes it fail at the call site with an identity callers can check uniformly (see R-ERR-001).
