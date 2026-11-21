---
title: 'ADR 023: Backports are manual, and forward-port assessments live in the PR'
description: >-
  Backports are manual cherry-pick PRs, and a forward-port assessment in the
  PR description replaces the automatic needs-forward-port issue.
---

# ADR 023: Backports are manual, and forward-port assessments live in the PR

## Status

Accepted (2026-11-21). Supersedes the label-triggered backport action in
rule 3 and the automatic `needs-forward-port` issue in rule 5 of
[ADR 011](011-release-branches-and-cross-major-flow.md). The other ADR 011
rules stand.

## Context

ADR 011 relied on three workflows: `backport.yml` opened cherry-pick PRs
from a `backport release/vX.Y` label, `forward-port-tracker.yml` opened a
`needs-forward-port` issue for each merge into `release/v*` that lacked
backport provenance, and `maintenance-label.yml` created the labels.

Maintenance branches see few changes. The three workflows carry
`pull_request_target` permissions and a third-party action, and they add
maintenance cost that the change volume doesn't justify.

Alternatives considered:

1. **Keep the workflows.** Rejected: the security surface and upkeep outweigh
   the few backports per year.
2. **Keep only the forward-port tracker.** Rejected: a tracking issue adds
   no information that the PR description can't hold, and it creates issues
   that nobody owns.

## Decision

1. A maintainer backports a fix by cherry-picking the `main` commit onto a
   branch off `release/vX.Y` and opening a PR. The PR description links the
   original PR.
2. Any PR into `release/v*` that doesn't come from `main` includes a
   forward-port assessment in its description. The assessment links the
   matching `main` PR or records a “won't port” rationale.
3. Reviewers don't approve a PR into `release/v*` without the assessment.
4. Everything else in ADR 011 stands: lazily cut branches, PR-only changes
   with required CI, and release-please ownership of `VERSION` and
   `CHANGELOG.md`.

## Consequences

- **Pros:** fewer privileged workflows, no label upkeep, and the port
  decision sits next to the change it describes.
- **Cons:** reviewer discipline replaces automation, and nothing flags a
  missing assessment. A merged PR without one leaves silent drift.
- **Rule for future release questions:** don't reintroduce a backport or
  tracker workflow without a new ADR that explains why the volume of
  maintenance changes now justifies it.
