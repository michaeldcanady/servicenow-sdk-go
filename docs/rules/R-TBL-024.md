---
id: R-TBL-024
applies_to: ["store/**/*.go", "serialization/**/*.go", "internal/**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# The store and serialization foundation is exported

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- Paths follow the original ADR 019 (`store`, `serialization`); adjust if the foundation is exported elsewhere. -->

MUST: Export the store and serialization foundation, or a convenience base type over it, outside `internal/`.
MUST NOT: Keep any capability that built-in typed tables rely on reachable only from `internal/`.

## Verify

```bash
# Passes only once the foundation has left internal/. Expected to fail while this rule is proposed.
[ -d store ] && [ -d serialization ] && [ ! -d internal/store ] && [ ! -d internal/serialization ]
```

## Rationale

User-defined tables reach parity with built-in tables only if they can use the same backing-store and serialization machinery. Leaving it in `internal/` would make built-in tables permanently privileged (see R-TBL-019).
