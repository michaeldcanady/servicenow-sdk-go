---
id: R-REL-011
applies_to: [".github/workflows/*.yml", ".github/workflows/*.yaml", "VERSION", "CHANGELOG.md"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Development is trunk-first with lazily cut maintenance branches

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- Assumed label names: `backport`, `needs-forward-port`. Enforcement of PR-only and required CI is a GitHub ruleset, not repo content. -->

MUST: Develop on `main` and land every fix there first.
MUST: Backport fixes from `main` to `release/vX.Y` after they land.
MUST: Create `release/vX.Y` only when a backport is needed, from the last `vX.Y.*` tag.
MUST NOT: Create a `release/vX.Y` branch ahead of need.
MUST: Land every change on `release/*` through a PR with required CI.
MUST NOT: Push directly to `release/*`.
MUST: Label every backport PR `backport`.
MUST: Complete the forward-port assessment in the PR before merging into `release/v*`, stating whether the change already exists on `main` and applying `needs-forward-port` if it does not.
MUST NOT: Edit `VERSION` or `CHANGELOG.md` by hand; release-please owns them on every branch.
MUST NOT: Treat a `release/*` branch as a development trunk.
MAY: Fix an older major by cutting `release/vX.Y` from that major's last tag, because each major has its own module path (see R-NAME-010).

## Verify

```bash
# Requires an authenticated `gh`. Fails closed if CI is not wired to release branches or no active branch ruleset exists.
grep -RqE 'release/(\*\*|v)' .github/workflows/ \
  && gh api "repos/{owner}/{repo}/rulesets" --jq 'any(.[]; .target == "branch" and .enforcement == "active")' | grep -qx true
```

## Rationale

The v2 launch drifted `main` and `release/v2` until cross-major merges became rewrites. Lazily cut, downstream-only branches with tracked forward-ports turn drift into visible debt instead of silent divergence.
