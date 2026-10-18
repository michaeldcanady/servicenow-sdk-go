---
id: R-REL-012
applies_to: [".github/workflows/*.yml", ".github/workflows/*.yaml"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: release provenance, sbom, signing, signature, release verify,
in-band, workflow_call, release created, GITHUB_TOKEN suppression -->

MUST generate SBOM, signatures, and verification as reusable workflows called in-band from the releasing workflows, gated on `release_created`.
MUST order signing after SBOM attachment (`needs: attach-sbom`) so SHA256SUMS covers the SBOM.
MUST keep a `workflow_dispatch` trigger with a `tag` input on each provenance workflow for backfill.
MUST NOT trigger provenance workflows on `release: published` or `push tags v*` for releases release-please creates with the default token — GitHub suppresses those events and they never arrive.

## Verify

```bash
# Fails (non-zero) if a provenance workflow re-adds a suppressed event trigger.
! grep -RInE 'release:\s*published' .github/workflows/sbom.yml .github/workflows/sign-release.yml
```

## Rationale

GitHub suppresses new runs for events created by the default token, so `release: published` triggers are dead on arrival for release-please releases; in-band `workflow_call` fan-out is deterministic and ordered with no new credentials (see ADR 012).
