---
id: R-REL-011
applies_to: [".github/workflows/*.yml", ".github/workflows/*.yaml", "VERSION", "CHANGELOG.md"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: release branches, release/vX.Y, backport, forward port, trunk
first, needs-forward-port, drift, lazy cut, maintenance branch -->

MUST develop on `main` (trunk-first) and land fixes there first, then backport down.
MUST create `release/vX.Y` branches lazily from the last `vX.Y.*` tag — never preemptively.
MUST route all changes to `release/*` by PR with required CI; direct pushes are blocked.
MUST label backports so provenance is tracked, and never merge into `release/v*` without completing the forward-port assessment.
MUST NOT hand-edit `VERSION` or `CHANGELOG.md` — release-please owns them on every branch.
MUST NOT treat a maintenance branch as a development trunk.

## Verify

```bash
# Passes (zero) while release-branch policy tooling references release/v refs.
grep -RInE 'release/v' .github/
```

## Rationale

The v2 launch drifted `main` and `release/v2` until cross-major merges became rewrites; lazily-cut, downstream-only branches with tracked forward-ports turn drift into visible debt instead of silent divergence (see ADR 011).
