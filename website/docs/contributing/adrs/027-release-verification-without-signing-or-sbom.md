---
title: 'ADR 027: Releases ship without signing or an SBOM, and the merge gate replaces post-release verification'
description: >-
  Supersedes ADR 012 and withdraws ADR 013. Releases carry no signatures and
  no SBOM. release-verify.yml goes away once the gate is required on release
  PRs and a tag ruleset blocks hand-pushed tags.
---

# ADR 027: Releases ship without signing or an SBOM, and the merge gate replaces post-release verification

## Status

Proposed (2026-11-21). When accepted, this record supersedes
[ADR 012](012-release-provenance-in-band.md) and withdraws
[ADR 013](013-event-driven-release-dispatch.md). The plan is in
`specs/010-holistic-delivery-pipeline/spec.md`.

## Context

[ADR 012](012-release-provenance-in-band.md) made SBOM generation, artifact
signing, and release verification run as reusable workflows that the
releasing workflow calls. GitHub doesn't start new runs from events that
`GITHUB_TOKEN` causes, so a release-event trigger never fired.
[ADR 013](013-event-driven-release-dispatch.md) proposed moving those tasks to
`repository_dispatch` events, possibly ending at a GitHub App token.

Since then:

- #983 removed the SBOM workflow, and #985 removed the sign-release workflow.
- Only `release-verify.yml` remains. `stable-release.yml` calls it after
  release-please creates a tag. It also runs on hand-pushed `v*` tags.
  It reruns the module, build, lint, and unit-test checks on the tagged ref.
  It can't undo a published tag, so a failure means a human yanks or
  patches the release.
- [ADR 024](024-no-long-lived-ci-secrets.md) rejects a stored GitHub App
  key, which was the end state of ADR 013.

The reason for dropping signing and SBOM, **pending review by the
maintainer**: the upkeep cost more than the value, for a library that people
get through the Go module proxy, where `sum.golang.org` already gives checksum
integrity, and that has no binaries to sign.

Alternatives considered:

1. **Restore signing and SBOM.** Rejected for the reason above.
2. **Keep `release-verify.yml` as a backup check after the gate is required.**
   Rejected: once release PRs pass `PR Gate` with strict status checks, the
   tagged commit is the tested commit, and the backup reruns the same checks
   after the release can no longer be stopped.
3. **Remove `release-verify.yml` now.** Rejected: until the gate is required
   on release PRs, it is the only check that runs on release-please's code
   changes (`version.go`).

## Decision

1. **Releases carry no artifact signatures and no SBOM.** Don't add either
   back without a new ADR.
2. **Remove `release-verify.yml` and its call in `stable-release.yml` once
   both of these are true:**
   - `PR Gate` is a required check on release PRs, and strict status checks
     are on ([ADR 025](025-delivery-system-boundary-and-merge-gate.md)).
   - A tag ruleset blocks hand-pushed `v*` tags, which is the other case that
     `release-verify.yml` covers.
3. **Mark ADR 012 superseded** by this record, and set `R-REL-012` to
   `status: superseded`.
4. **Mark ADR 013 withdrawn.** With signing and SBOM gone and bot PRs using
   GitHub's built-in approval, there are no release tasks left to dispatch.
   Set `R-REL-013` to `status: superseded`.

## Consequences

- **Pros:**
  - Two release workflows and a planned dispatch design are gone.
  - The release path has one point of verification: the gate on the release
    PR.
- **Cons:**
  - Consumers can't verify release signatures, and no SBOM is published.
  - The OpenSSF Scorecard "Signed-Releases" score drops.
  - Until the two conditions in decision 2 hold, `release-verify.yml` stays
    as the only check on release-please code changes.
- **Rule for future release work:** a new release-time check goes into the
  gate on the release PR, not into a workflow that runs after the tag
  exists.
