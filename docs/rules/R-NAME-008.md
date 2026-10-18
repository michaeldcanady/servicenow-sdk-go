---
id: R-NAME-008
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: naming axes, package name, symbol name, url segment, official
api name, wire format, accessor, rename, independent -->

MUST name the package after ServiceNow's official name for the API surface.
MUST name exported symbols after the actual wire format (the JSON shape).
MUST name accessor methods after the literal URL segment.
MUST NOT assume a rename must change all three axes together — evaluate each independently.
SHOULD check each axis separately when a module name diverges from ServiceNow's official naming.

## Verify

```bash
# Guards the ADR-008 rename specifically: the /api/now/stats module package
# must be `aggregationapi`, not `statsapi`. Fails (non-zero) if reverted.
! grep -RInE 'package statsapi' --include='*.go' .
```

## Rationale

Package, symbol, and URL-segment names serve different audiences (discoverability, wire fidelity, navigation) and need not agree; renaming all three for superficial uniformity breaks wire/URL correspondence for no gain (see ADR 008).
