---
id: R-ERR-006
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: nil receiver, nil guard, nil request builder, ErrNilRequestBuilder,
return nil nil, sentinel error, fail loud -->

MUST return `snerrors.ErrNilRequestBuilder` (or the package zero value plus that sentinel) from every verb-method nil-receiver guard.
MUST NOT return `(nil, nil)` or a bare `nil` from a nil-receiver guard.
MUST limit this to verb-method nil-*receiver* guards — navigation methods returning a bare `*Builder`, and `res == nil` / `ElementValue.IsNil` semantics, are unchanged.

## Verify

```bash
# Fails (non-zero) if a request-builder verb method still returns bare `nil, nil`.
# Scoped to request-builder files: `nil, nil` in deserialization/model files is
# the legitimate null-element behavior ADR-006 explicitly leaves unchanged.
! grep -RInE 'return nil, nil' --include='*_request_builder.go' --include='*_item_request_builder.go' --include='*_builder.go' .
```

## Rationale

A `(nil, nil)` guard turns a nil builder into silent success that fails later, far from the mistake; a shared sentinel makes it fail loud at the call site with an identity callers check uniformly (see ADR 006, ADR 001).
