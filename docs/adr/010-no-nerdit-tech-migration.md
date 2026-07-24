# ADR 010: Stay at `github.com/michaeldcanady/servicenow-sdk-go` — no org migration at v2

## Status

Accepted

## Context

Backlog grooming for v2.0 (#493) flagged that a GitHub org migration — moving
the repo from `michaeldcanady/servicenow-sdk-go` to a `NerdIT-Tech` org path —
is inherently breaking, since it changes the module import path. v2 already
forces every consumer to update their import paths for the `/v2` semantic-import-versioning
suffix (see the release-day runbook, `release-2.0-issues/003-v2-module-path-runbook.md`),
so folding an org migration into that same bump would have been free from a
consumer-churn standpoint. Migrating the org *after* v2 ships would mean
either a v3 major bump just for the path change, or a permanent
redirect/fork story — "TBD" was not viable, since v2.0.0 is the only point
where this decision is free.

Alternatives considered:

1. **Migrate to `NerdIT-Tech` at v2** — transfer the GitHub repo before
   release day, fold the new org path into the same commit as the `/v2`
   module bump. Rejected: no concrete driver for the org move materialized
   during grooming beyond a speculative future-org-structure benefit: the
   maintainer decided against it (2026-07-24).
2. **Stay at `michaeldcanady`** — commit to the current path for the entire
   v2 lifecycle, revisit only at a future major version if a real reason
   emerges.

## Decision

Keep the module path at `github.com/michaeldcanady/servicenow-sdk-go`. The
`/v2` module-path bump (release-day runbook item, `release-2.0-issues/003`)
changes only the semantic-import-versioning suffix — it does **not** change
the org/owner segment. No GitHub repo transfer is planned for v2.0.0 or v3.

## Consequences

- **Pros:** no repo transfer to coordinate before release day (Actions,
  apps, branch protection rules, CODEOWNERS handles all stay as-is); no risk
  to the `go get` redirect story; one less moving part in the release
  runbook.
- **Cons:** none identified — this was a speculative option, not a
  committed direction, so declining it costs nothing.
- **Rule for future naming/org questions:** this is settled through v3 —
  do not revisit without a new major version boundary and a concrete,
  non-speculative reason to move.
