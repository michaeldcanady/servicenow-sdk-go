---
id: R-TBL-019
applies_to: ["tableapi/**/*.go", "tables/**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Typed tables are uniform, and user tables are first-class

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- TODO(owner): clarify "consumers cannot modify built-in tables" (immutable definitions? no struct embedding?). Split-out decisions: R-TBL-022 (descriptors), R-TBL-023 (references), R-TBL-024 (exported foundation). -->

MUST: Ship built-in typed tables (Incident and Case first) that consumers can use but not modify, reached by literal chaining off the collection builder.
MUST: Let a consumer define a typed table by wrapping the exported collection builder in a named wrapper type instantiated as a variable (`tables := mycool.New(client.Now().Table2())`).
MUST: Keep typed tables uniform, exposing record CRUD and query only.
MUST NOT: Put domain actions on a typed table type.
MUST: Put domain actions in scoped API packages.
MUST NOT: Make a user-defined table second-class; any capability of a built-in table is reachable from a user table through the exported foundation (R-TBL-024).

## Verify

```bash
# Requires an Example test that wraps Table2() in a user-defined typed table. Fails closed if absent.
go test ./... -run 'ExampleUserDefinedTable' -v -count=1 | grep -q -- '--- PASS: ExampleUserDefinedTable'
```

## Rationale

Go cannot define methods on a type from another package, so a user-defined typed table must wrap the exported collection builder in its own named type. Keeping tables uniform moves domain actions to scoped packages, so built-in and user tables stay interchangeable.
