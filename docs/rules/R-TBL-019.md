---
id: R-TBL-019
applies_to: ["**/*.go", "internal/store/**", "internal/serialization/**"]
status: proposed
supersedes: []
conflicts_checked: true
---

<!-- Keywords: typed table, field descriptor, wrapper variable, user defined
table, parity, domain action, typed reference, no auto navigate, Incident,
Case, MyCoolTable, codegen -->

MUST ship built-in typed tables (Incident, Case first) that consumers use but cannot modify, reachable by literal chaining off the collection builder.
MUST let a consumer define a typed table by wrapping the exported collection builder in a named wrapper type instantiated as a variable (`tables := mycool.New(client.Now().Table2())`) — Go cannot attach methods to an anonymous chain.
MUST drive typed fields and typed query from field descriptors carrying the storage key and a typed accessor.
MUST export the store/serialization foundation (or a convenience base type over it) so a user-defined table has full parity with a built-in table.
MUST keep typed tables uniform: expose record CRUD and query, not domain actions.
MUST place domain actions in scoped API packages, not on typed table types.
MUST represent a typed reference as a wrapper carrying sys_id, display value, and link.
MUST NOT auto-navigate a typed reference to the referenced record.
MUST NOT make user-defined tables second-class — any capability a built-in table has, a user table must reach through the exported foundation.
SHOULD generate field descriptors with optional codegen, treating the exported helpers as the contract and codegen as convenience.
MAY offer both query styles (fluent builder and typed domain fields) coexisting.

## Verify

```bash
# Passes (zero) once the store/serialization foundation is reachable outside
# internal/ (the parity prerequisite); fails while it is still internal-only.
test -d store -o -d serialization
```

## Rationale

Go can't attach methods to an anonymous chain, so a user-defined typed table wraps the exported collection builder in a named variable; exporting the store/serialization foundation gives user tables full parity with built-ins, and keeping tables uniform moves domain actions to scoped packages (see ADRs 005, 007, 019).
