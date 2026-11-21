---
id: R-REL-027
applies_to: [".github/workflows/*.yml", ".github/workflows/*.yaml"]
status: proposed
supersedes: [R-REL-012, R-REL-013]
conflicts_checked: true
---

# Releases ship without signing or an SBOM, and the gate verifies release PRs

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST NOT: Add artifact signing or SBOM generation to the release flow without a new ADR.
MUST NOT: Remove `release-verify.yml` before `PR Gate` is required on release PRs with strict status checks, and a tag ruleset blocks hand-pushed `v*` tags.
SHOULD: Put a new release-time check in the gate on the release PR, not in a workflow that runs after the tag exists.
MUST NOT: Add `repository_dispatch` fan-out for release tasks; ADR 013 is withdrawn.

## Verify

```bash
# Fails when a workflow reintroduces signing or SBOM tooling.
! grep -RqiE 'cosign|sigstore|syft|anchore/sbom|spdx|cyclonedx' .github/workflows .github/actions
```

## Rationale

The Go module proxy and `sum.golang.org` already give checksum integrity, and the repo ships no binaries. See [ADR 027](../../website/docs/contributing/adrs/027-release-verification-without-signing-or-sbom.md).
