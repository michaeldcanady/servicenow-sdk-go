---
title: 'ADR 013: Release tasks may move to event-driven dispatch'
description: >-
  Proposed direction: decouple release follow-on work from the releasing
  workflows using repository_dispatch events and a reconciliation sweep,
  possibly ending at a GitHub App token for native event triggers.
---

# ADR 013: Release tasks may move to event-driven dispatch

## Status

Proposed (2026-09-02). Not adopted — this records a direction for moving
release follow-on tasks from in-band coupling (ADR 012) to event-driven
dispatch. If adopted, it supersedes [ADR 012](012-release-provenance-in-band.md),
whose Status then reads "Superseded by ADR-013."

## Context

[ADR 012](012-release-provenance-in-band.md) couples follow-on work to the
releasing workflows because GitHub doesn't start new runs from events created
by `GITHUB_TOKEN`. In-band coupling is correct today, but it has real costs:

- The releasing workflows name every consumer (`attach-sbom`, `attach-sign`,
  `verify-release`), so a new consumer is a change to the production release
  path, and the fan-out list grows with each concern.
- A skipped dependency or a subtly wrong `if` gate drops coverage silently.
- Logically separate concerns share one run's history, retry, and concurrency
  semantics.

Event-driven dispatch is attractive because each concern becomes its own
workflow with an explicit contract: consumers subscribe and receive a durable
event rather than being invoked by the producer. GitHub offers one event type
that's exempt from the `GITHUB_TOKEN` suppression and can carry arbitrary
data: `repository_dispatch`. `workflow_dispatch` is the other exemption but
can't carry a custom payload.

Options for true event delivery:

1. **`repository_dispatch` event bus** — the releasing workflow POSTs a
   `release-created` dispatch with `client_payload` (`tag`, `sha`) after the
   release-please step; every consumer subscribes independently. Fire-and-
   forget: a lost dispatch is silent, so the bus needs a reconciliation sweep
   (see Decision) to become at-least-once. Payload must be versioned.
2. **GitHub App installation token** — events created with an app token start
   new runs, restoring native triggers such as `release: published` and
   `push tags v*`. Adds an org-installed GitHub App, key storage and rotation,
   permitting scope, and a write-credential security review. This is the
   "real principal" end state, but it's a heavy operational lift.
3. **`workflow_run` trigger** — keys off a dedicated "emit release events"
   workflow run rather than a repository event. The consumer receives no
   inputs and must resolve state from run metadata or artifacts, and the
   written security guidance prefers `workflow_call` chains over
   `workflow_run`. Kept out of scope.
4. **State-file polling** — a scheduled sweep diffs the release manifest
   against created tags and attached assets. Fully decoupled, slow to react,
   and
   failure-agnostic on its own; valuable as the reconciliation layer for the
   dispatch bus rather than as the primary mechanism.

### Ordering constraint

SDK releases must ship SBOM before checksums: SHA256SUMS must cover the SBOM.
Event delivery is unordered across consumers, so any event-driven design must
either keep order-sensitive steps in-band (call `sbom.yml` before
`sign-release.yml` as today), encode ordering in the payload (a `phase` that
consumers honor), or have signing wait until the SBOM asset exists and retry.

## Decision

Don't adopt yet. If adopted, pursue this phased path:

1. **Phase 1 — add a dispatch emitter alongside the in-band calls.** The
   releasing workflows POST `repository_dispatch` events of type
   `release-created` with a versioned payload (`v1`, `tag`, `sha`), with a
   small retry-and-jitter loop. In-band attachment continues unchanged.
2. **Phase 2 — move stateless consumers to the bus.** `release-verify.yml`
   (and any future consumer that doesn't affect artifact ordering) subscribes
   to `release-created` instead of being called in-band. SBOM and signing stay
   in-band because their ordering matters.
3. **Phase 3 — reconciliation sweep.** A scheduled workflow diffs tags against
   attached assets and re-dispatches anything missing, converting the bus from
   fire-and-forget to at-least-once with observable drift. Only when this sweep
   has proven itself in production would the in-band calls become redundant.
4. **Phase 4 (optional) — GitHub App for native triggers.** If a maintainer
   stands up a GitHub App, the emitter switches from `repository_dispatch` to
   the app token, native triggers return, and the event bus disappears. This
   substitutes an operational credential for pipeline complexity; don't do it
   without an explicit cost decision.

## Consequences

- **Pros:** consumers decouple from the release path — adding one doesn't
  touch the production release workflow; each concern reverts to its own
  workflow, history, and retry semantics; the design matches the mental model
  of subscriptions to durable events; reconciliation makes the pipeline fail
  loud instead of silent.
- **Cons:** the event bus and the reconciliation sweep are new pipeline to
  build and maintain; ordering between SBOM and signing must survive
  nondeterministic delivery, so the order-sensitive steps stay in-band for
  now; the phased approach intentionally keeps two mechanisms (in-band calls
  and dispatches) alive in parallel during the transition; the GitHub App
  option trades an operational credential for the least pipeline surface, and
  only a maintainer decision can settle that.
- **Rule for future work:** until both ordering and at-least-once delivery are
  guaranteed, keep SBOM and signing in-band (`attach-sbom` before
  `attach-sign`) and treat any event-driven consumer as best-effort overlays
  with their own reconciliation. Never make the event bus the only mechanism
  for a step that must not be missed.