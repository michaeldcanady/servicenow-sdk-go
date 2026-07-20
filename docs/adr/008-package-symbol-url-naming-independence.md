# ADR 008: Package names, exported symbol names, and URL segments are independent naming axes

## Status

Accepted

## Context

The module behind ServiceNow's `/api/now/stats` endpoint was implemented as
package `statsapi`. ServiceNow's own documentation calls this the
**Aggregate API**, and the SDK's docs site already labeled the module
"Aggregation" — the package name (`statsapi`) was the one remaining place
still using the old name, creating a naming inconsistency between the code,
the docs, and ServiceNow's own terminology (#491, #505).

The naive fix — rename everything, package and symbols and accessor alike,
to `aggregation*` — would have been a larger and less-justified break: the
exported `Stats*` type names mirror the actual wire format (ServiceNow's
JSON response shape uses `stats`), and the `Now().Stats()` accessor mirrors
the literal URL segment (`/api/now/stats`). Neither of those has anything to
do with what the *package* should be called for discoverability.

Alternatives considered:

1. **Rename nothing, live with the inconsistency** — rejected; a
   contributor searching ServiceNow's docs for "Aggregate API" would not
   find `statsapi` by name.
2. **Rename everything (package, types, accessor) to `aggregation*`** —
   rejected; it would break the correspondence between `Stats*` types and
   the actual `stats` wire field, and between `.Stats()` and the actual URL
   segment — changes with no real justification, made only for superficial
   uniformity.
3. **Rename only the package** (`statsapi` → `aggregationapi`), keep
   `Stats*` symbol names and the `Now().Stats()` accessor as-is.

## Decision

Package names follow ServiceNow's official name for the API surface.
Exported symbol names follow the actual wire format. Accessor method names
follow the literal URL segment. These three do not have to agree with each
other, and a naming inconsistency in one axis doesn't imply the other two
are wrong.

## Consequences

- **Pros:** the package is discoverable under ServiceNow's official
  terminology; `Stats*` types and `.Stats()` still read correctly against
  the actual JSON payload and URL a reader sees on the wire.
- **Cons:** breaking for import paths
  (`github.com/.../statsapi` → `github.com/.../aggregationapi`); mildly
  surprising on first read that `aggregationapi.StatsResponse` isn't a typo.
- **Rule for future renames:** when a module's name diverges from
  ServiceNow's official naming, check independently whether the package
  name, the symbol names, and the accessor/URL segment each need to change
  — don't assume a rename must be all-or-nothing across all three.
