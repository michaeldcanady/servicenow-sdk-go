---
title: 'ADR 012: Release provenance is generated in-band, not event-driven'
description: >-
  SBOM generation, artifact signing, and release verification run as reusable
  workflows called by the releasing workflow, because GitHub suppresses new
  workflow runs from events created by the default token.
---

# ADR 012: Release provenance is generated in-band, not event-driven

## Status

Accepted (2026-09-02). This is the root record for the release-provenance
pipeline. The rules apply to any workflow that must react to a release-please-
created release.

## Context

The v2.0.3 release (2026-09-02) shipped with zero provenance assets: no SBOM,
no checksums, and no signature. v2.0.2's SBOM only existed because a maintainer
attached it by manual dispatch. The releases themselves were created correctly
by release-please; the follow-on work never ran.

Root cause: GitHub suppresses new workflow runs for repository events created
by the default `GITHUB_TOKEN`, with two explicit exceptions,
`workflow_dispatch` and `repository_dispatch`. `sbom.yml` and `sign-release.yml`
triggered on `release: published`, and `release-verify.yml` triggered on
`push tags v*`. release-please-action creates both the tag and the release as
a GitHub Action running with the default token, so neither event ever arrived.
`sign-release.yml` had never run once; `sbom.yml` had only manual runs.

Alternatives considered:

1. **Repository-event triggers (`release: published`, `push tags v*`)** — this
   is what shipped and what broke. Dead on arrival for any release created by
   automation running on the default token. Rejected as the operating model.
2. **A real principal for release-please (PAT or GitHub App installation
   token)** — events created by an app or PAT do start new workflow runs, so
   this restores the event-driven design. It adds a secret or key lifecycle,
   a security review of a write-scoped credential, and no delivery guarantee:
   a missed event is still silent. Rejected for now; revisited in ADR 013.
3. **`repository_dispatch` event bridge** — dispatch events are explicitly
   exempt from suppression, so an in-band POST reliably starts consumers.
   The POST is fire-and-forget: ordering across consumers is not guaranteed,
   and a lost dispatch fails silently. Rejected as the primary path; the
   payload versioning and reconciliation needed to make it reliable are
   deferred to ADR 013.
4. **In-band `workflow_call` fan-out** (chosen) — the releasing workflow calls
   each dependent directly as a reusable workflow, gated on the release-please
   job's `release_created` output, passing the created tag as an input.
   Deterministic, ordered, and auditable in one run; no new credentials.
5. **State-file polling** — a scheduled workflow diffs the release manifest
   (`.release-please-manifest*.json`) against created tags. Fully decoupled
   but laggy and failure-agnostic; not suitable as the only mechanism on the
   release-critical path.

## Decision

Couple release-please's follow-on work to the releasing workflow in-band.
The releasing workflows (`stable-release.yml`, `weekly-release.yml`) publish
`release_created` and `tag_name` outputs from their release-please jobs, then
call each dependent as a reusable workflow via
`uses: ./.github/workflows/<name>.yml` with a `tag` input, gated on
`release_created == 'true'`. The pattern covers:

- **SBOM generation and attachment** (`sbom.yml`).
- **Artifact signing and optional tag signing** (`sign-release.yml`), ordered
  by `needs: attach-sbom` so the SHA256SUMS always cover the SBOM.
- **Release verification** (`release-verify.yml`), which keeps its
  `push tags v*` trigger only for tags that a human pushes directly, because
  those do not pass through the releasing workflows.

Each reusable workflow also keeps a `workflow_dispatch` trigger with the same
`tag` input so a maintainer can backfill older releases, as happened for
v2.0.3 and v2.0.2.

## Consequences

- **Pros:** deterministic — the follow-on work fires in the same run that
  created the release, with no reliance on event delivery; ordered, so signing
  always sees the SBOM; one workflow run to inspect after a release; manual
  backfill stays available; no new secrets or credentials.
- **Cons:** the releasing workflows must name every dependent, so adding a
  consumer is a change to the production release path; the gate expression is
  hand-maintained and a wrong `if` silently skips coverage; concerns that are
  logically separate share one run's history and retry semantics.
- **Rule for future release questions:** do not add `release: published` or
  `push tags v*` triggers for releases that release-please creates with the
  default token — the trigger is dead on arrival. Add the work as a reusable
  `workflow_call` workflow and invoke it from `stable-release.yml` and
  `weekly-release.yml`. Revisit this rule only under the conditions in
  [ADR 013](013-event-driven-release-dispatch.md), which may supersede this
  record.