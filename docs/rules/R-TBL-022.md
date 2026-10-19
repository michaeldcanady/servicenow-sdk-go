---
id: R-TBL-022
applies_to: ["tables/**/*.go", "store/**/*.go", "serialization/**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Typed fields and queries are driven by field descriptors

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- TODO(owner): replace the descriptor type name in Verify with the real one. -->

MUST: Drive typed fields and typed queries from field descriptors that carry the storage key and a typed accessor.
MUST NOT: Require codegen to define a typed table.
SHOULD: Generate field descriptors with optional codegen, treating the exported helpers as the contract and codegen as a convenience.
MAY: Offer the fluent query builder and typed domain-field queries side by side.

## Verify

```bash
# Fails closed (exit 1) if no generic field-descriptor type is declared.
grep -RqE 'type [A-Za-z]*Field(Descriptor)?\[' --include='*.go' .
```

## Rationale

Descriptors tie a field's storage key to its Go type once, so typed queries and typed model access cannot drift apart. Hand-written descriptors keep user tables at parity with generated ones.
