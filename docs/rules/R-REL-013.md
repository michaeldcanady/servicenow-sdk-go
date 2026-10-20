---
id: R-REL-013
applies_to: [".github/workflows/*.yml", ".github/workflows/*.yaml"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Event-driven provenance is phased in beside in-band calls

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST NOT: Make an event bus the only delivery mechanism for SBOM generation or signing.
MUST: Keep SBOM generation and signing in-band (R-REL-012) until both ordering and at-least-once delivery are guaranteed.
SHOULD: Add a reconciliation sweep before treating any dispatch consumer as reliable.
MAY: Phase in `repository_dispatch` emitters alongside the in-band calls.

## Verify

```bash
# Both releasing workflows must still call SBOM in-band. Fails closed if either file is missing.
grep -qE 'uses:\s*\./\.github/workflows/sbom\.yml' .github/workflows/stable-release.yml \
  && grep -qE 'uses:\s*\./\.github/workflows/sbom\.yml' .github/workflows/weekly-release.yml
```

## Rationale

Event delivery is fire-and-forget and unordered, so an event bus alone can silently drop order-sensitive steps such as SBOM-before-signing. The phased design keeps in-band calls until reconciliation proves at-least-once delivery.
