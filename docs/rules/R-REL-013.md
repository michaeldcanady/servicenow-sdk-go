---
id: R-REL-013
applies_to: [".github/workflows/*.yml", ".github/workflows/*.yaml"]
status: proposed
supersedes: []
conflicts_checked: true
---

<!-- Keywords: repository_dispatch, event bus, event-driven, reconciliation,
at-least-once, github app, dispatch emitter, phased -->

MUST NOT make an event bus the only mechanism for a step that must not be missed.
MUST keep SBOM and signing in-band until both ordering and at-least-once delivery are guaranteed.
SHOULD add a reconciliation sweep before treating any dispatch consumer as reliable.
MAY phase in `repository_dispatch` emitters alongside the in-band calls.

## Verify

```bash
# Passes (zero) while SBOM is still called in-band from the releasing workflows.
grep -RInE 'uses:\s*\./\.github/workflows/sbom\.yml' .github/workflows/stable-release.yml .github/workflows/weekly-release.yml
```

## Rationale

Event delivery is fire-and-forget and unordered, so an event bus alone can silently drop order-sensitive steps like SBOM-before-signing; the phased design keeps in-band calls until reconciliation proves at-least-once delivery (see ADR 013, ADR 012).
