---
id: R-NAME-010
applies_to: ["go.mod", "RELEASE.md"]
status: accepted
supersedes: [R-REL-014]
conflicts_checked: true
---

# The module path carries the major version and a fixed owner

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Keep the module path rooted at `github.com/michaeldcanady/servicenow-sdk-go`.
MUST: Encode the major version only as a `/vN` suffix, for N of 2 or higher.
MUST NOT: Add a `/v0` or `/v1` suffix.
MUST: Bump the module path to `/v{N+1}` and tag `v{N+1}.0.0` when a breaking change lands.
MUST: Release minor and patch versions as Git tags (`v2.1.0`, `v2.1.1`, ...).
MUST NOT: Change the owner or org segment of the module path at a major bump.
MUST NOT: Migrate the repository to a different GitHub org for v2 or v3.

## Verify

```bash
# Run at release time with tags fetched (fetch-depth: 0). Fails closed if tags are missing.
mod=$(awk '/^module /{print $2}' go.mod)
printf '%s\n' "$mod" | grep -qE '^github\.com/michaeldcanady/servicenow-sdk-go(/v([2-9]|[1-9][0-9]+))?$' || exit 1
mod_major=$(printf '%s' "$mod" | sed -nE 's#.*/v([0-9]+)$#\1#p'); mod_major=${mod_major:-1}
tag_major=$(git tag -l 'v[0-9]*' | sort -V | tail -1 | sed -E 's/^v([0-9]+)\..*/\1/')
[ -n "$tag_major" ] && [ "$mod_major" = "$tag_major" ]
```

## Rationale

A major bump changes the import path, so consumers must opt in deliberately and two majors can coexist in one build. An org migration is inherently breaking, and no concrete driver for one has materialized, so declining it keeps the release runbook simple.
