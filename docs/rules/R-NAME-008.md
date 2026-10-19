---
id: R-NAME-008
applies_to: ["*api/**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Package, symbol, and accessor names are chosen independently

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Name each API package after ServiceNow's official name for that API surface.
MUST: Name exported symbols after the wire format (the JSON shape).
MUST: Name accessor methods after the literal URL segment they traverse.
MUST NOT: Rename one axis only to match another; evaluate package, symbol, and accessor names independently.
MUST: Name the `/api/now/stats` package `aggregateapi`, after ServiceNow's "Aggregate API".
SHOULD: Record the reason in the package doc when a package name diverges from ServiceNow's official name.

## Verify

```bash
# The /api/now/stats package must be `aggregateapi`; neither earlier name may reappear.
# grep exit 1 = no match = pass; exit 2 fails closed.
{ grep -RInE '^package (statsapi|aggregationapi)\b' --include='*.go' --exclude-dir=.git .; [ $? -eq 1 ]; } \
  && grep -RqE '^package aggregateapi\b' --include='*.go' .
```

## Rationale

Package, symbol, and URL-segment names serve different audiences (discoverability, wire fidelity, navigation) and need not agree. Renaming all three for superficial uniformity breaks wire and URL correspondence for no gain.
