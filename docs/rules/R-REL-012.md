---
id: R-REL-012
applies_to: [".github/workflows/*.yml", ".github/workflows/*.yaml"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Release provenance runs in-band, not on release events

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Implement SBOM generation, signing, and verification as reusable `workflow_call` workflows called in-band from the releasing workflows.
MUST: Gate each in-band call on `release_created`.
MUST: Order signing after SBOM attachment (`needs: attach-sbom`) so `SHA256SUMS` covers the SBOM.
MUST: Keep a `workflow_dispatch` trigger with a `tag` input on each provenance workflow for backfill.
MUST NOT: Trigger a provenance workflow on `release: published` or tag-push events for releases release-please creates with the default token.

## Verify

```bash
# Parses YAML (handles multi-line `release:` / `types:` forms). Anything other than "false" fails closed.
for f in .github/workflows/sbom.yml .github/workflows/sign-release.yml; do
  [ "$(yq '(.on | has("release")) or ((.on.push.tags // []) | length > 0)' "$f")" = "false" ] || exit 1
done
```

## Rationale

GitHub does not start new workflow runs for events created by the default token (apart from `workflow_dispatch` and `repository_dispatch`), so `release: published` triggers never fire for release-please releases. In-band `workflow_call` fan-out is deterministic and ordered with no new credentials.
