---
title: 'ADR 028: An API-compatibility check blocks incompatible changes on release branches and advises on main'
description: >-
  PR Gate runs gorelease (or apidiff) against the latest tag on the line
  that the PR targets. An incompatible exported-API change fails on release/*
  and is reported, not blocked, on main.
---

# ADR 028: An API-compatibility check blocks incompatible changes on release branches and advises on main

## Status

Proposed (2026-11-21). The plan is in
`specs/010-holistic-delivery-pipeline/spec.md`.

## Context

Nothing in CI detects an incompatible change to the exported API, such as a
removed function or a changed signature. Today only reviewers catch it. Go's
semantic import versioning makes the cost high:

- On a maintenance line (`release/vX.Y`), an incompatible change is never
  allowed. It breaks every consumer pinned to that major.
- On `main`, an incompatible change needs a new major module path and a
  `BREAKING CHANGE:` footer, so that release-please bumps the major version.

Tools exist for this check. `golang.org/x/exp/cmd/gorelease` compares a module
against a base version and suggests the next version. `apidiff` reports
incompatible changes between two versions of a package.

Alternatives considered:

1. **Block on both `main` and `release/*`.** On `main`, the check would fail
   unless the PR title has `!` or the body has `BREAKING CHANGE:`. Rejected
   for now: the maintainer chose to start with a report on `main`, and to
   learn how often the check flags changes before it blocks any.
2. **No check.** Rejected: reviewers miss incompatible changes, and a miss on
   a maintenance line breaks consumers.
3. **Block on `release/*`, advise on `main`** (chosen).

## Decision

1. **`PR Gate` runs an API-compatibility job** whenever Go code changes. It
   uses `gorelease`, or `apidiff` if `gorelease` doesn't fit, against the
   latest tag on the line that the PR targets.
2. **On PRs into `release/*`, an incompatible change fails the job.**
3. **On PRs into `main`, the job reports incompatible changes and doesn't
   fail.** The report names each change, so the author can add `!` or a
   `BREAKING CHANGE:` footer.
4. **Pin the tool** under [ADR 026](026-tool-version-pins-and-drift-check.md).

## Consequences

- **Pros:**
  - Maintenance lines can't take an incompatible change by accident.
  - On `main`, the version bump and the API change are visible in the same
    PR.
- **Cons:**
  - On `main`, an incompatible change can still merge without a
    `BREAKING CHANGE:` footer if the author ignores the report.
  - `gorelease` lives in `golang.org/x/exp`, which makes no compatibility
    promise of its own.
- **Rule for future work:** revisit decision 3 once the report has run on
  `main` for a few months. If it flags real breaks that merge without a
  footer, make it blocking on `main` too, with a new ADR.
