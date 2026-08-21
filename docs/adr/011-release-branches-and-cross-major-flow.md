# ADR 011: Release branches are lazily cut, downstream-only, and never silently diverge

## Status

Accepted

## Context

The v2 launch left `main` (v1) and `release/v2` significantly drifted, because
new development landed on `release/v2` while `main` stayed on v1. Go's
semantic-import-versioning module paths make this especially expensive to
recover from: every touched file carries `/v1` vs `/v2` import suffixes, so
cross-line merges conflict in nearly every file and effectively become
rewrites rather than merges. Drift between major lines therefore compounds
fast and does not heal on its own.

Tooling reinforces the problem today: release-please fires only on `main`
(`.github/workflows/stable-release.yml`), so there is no supported path for
patching — or extending — a shipped major once `main` has moved on. That gap
becomes acute whenever a new ServiceNow API surface ships mid-cycle that
consumers still pinned to the previous major want *now*, not at the next
major release.

Alternatives considered:

1. **GitFlow-style per-major development** — develop v2 work directly on
   `release/v2`, as happened before. Rejected: this *is* the drift incident;
   the failure mode is structural, not incidental.
2. **Tags-only, no maintenance branches** — cut an ephemeral branch per
   patch, tag, delete. Rejected: once more than one patch accumulates,
   the line needs a durable target for PRs, required CI, and a
   release-please config; repeated ad-hoc cuts invite inconsistency.
3. **Preemptive branch per minor (`release/vX.Y.0` at every minor)** —
   rejected: most minors never receive a patch, and idle long-lived
   branches are exactly how silent drift happens again.

## Decision

Adopt trunk-first development with lazily-cut, downstream-only maintenance
branches (maintainer decision, 2026-08-21):

1. **`main` is always the tip of the next major's development.**
2. **Maintenance branches are named `release/vX.Y`** (no patch segment) and
   are **created lazily** from the last `vX.Y.*` tag at first need —
   `git branch release/v2.x v2.4.5 && git push origin release/v2.x`. Tags
   are immutable, so when the branch is cut is irrelevant; never cut one
   preemptively.
3. **Downstream-only within a major:** fixes land on `main` first, then move
   down via a label-triggered backport action (`backport release/v2.x`
   opens the cherry-pick PR). Direct pushes to `release/*` are blocked;
   all changes arrive by PR with required CI.
4. **Cross-major features** (a new ServiceNow capability arriving while
   `main` develops the next major):
   - Useful for both majors and small/portable (typical new API module) —
     build on `main`, ship in the next major, then open a matching PR
     against `release/vX.Y`, rewriting the import suffix (`/vN+1` → `/vN`).
   - Useful for both but touching diverged core — build on `main`; backport
     only if consumer demand justifies the port cost, otherwise document
     the change as next-major-only in its PR.
   - Old-major-only by design (depends on prior-major model shapes) — PR
     targets `release/vX.Y` directly and ships as a **minor** bump
     (`feat:` → `vX.(Y+1).0`), but triggers rule 5 below.
5. **Drift is tracked debt, never silent:** any merge into `release/v*`
   without backport provenance automatically opens a `needs-forward-port`
   issue ("assess porting #N to `main`"). It closes only by porting the
   change up or by recording an explicit "won't port" rationale.
6. **Releases off maintenance lines are batched and demand-driven**, driven
   by Conventional Commits through a second release-please configuration
   scoped to `release/v*` refs. Weekly preview releases remain `main`-only.
   `VERSION` and `CHANGELOG.md` stay release-please-owned on every branch —
   never hand-edited.
7. **EOL policy:** the current major's latest minor receives fixes and
   features; the prior major receives critical/security fixes only; its
   `release/*` branch is deleted at EOL with an announcement in the release
   notes. Forgotten long-lived branches are how the v2 incident happens
   twice.

Follow-up tooling work implied by this decision: a `stable-release` job
keyed on `release/v*` refs with an alternate release-please config, the
backport label action, the forward-port tracker workflow, and
`branch-policy.yml` updates covering `release/*`.

## Consequences

- **Pros:** drift becomes visible, tracked debt instead of silent
  divergence; consumers on shipped majors get a predictable path to fixes
  and selected features; maintenance releases gain real CI, changelogs, and
  tagging via release-please instead of hand-run tags.
- **Cons:** cross-major ports stay semi-manual — import-suffix differences
  conflict in nearly every touched file, so ports must be kept small;
  two release-please configs and several new workflows add maintenance
  surface; rules 4–5 depend on reviewer discipline that automation only
  partially enforces.
- **Rule for future release questions:** do not create preemptive
  `release/*` branches, never land feature work on a maintenance branch
  without completing the forward-port assessment, and do not treat a
  maintenance branch as a development trunk — if v2-era work starts feeling
  like active development again, that is the signal to re-evaluate the EOL
  clock, not to grow the branch.
