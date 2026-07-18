# Stable and weekly-preview release-please channels can fight each other

- **Priority:** P2 — decide before GA
- **Raised by:** Senior DevOps
- **Area:** Release automation

## Problem

Two release-please workflows operate on the same branch (`main`), same package, same
`CHANGELOG.md`, and same tag namespace:

- `stable-release.yml` — every push to `main`, `release-please-config.json`.
- `weekly-release.yml` — Mondays 3 AM, `weekly-release-please-config.json` with
  `"prerelease": true`.

Because both configs share `changelog-path` and version files, they generate competing
release PRs proposing different versions from the same commit range; whichever merges
second rebases the other's changelog edits, and prerelease tags interleave with stable
tags in one sequence. Additionally, `weekly-release.yml` grants `contents: write` at
the workflow level to *both* jobs (detect-changes only needs read — it re-declares
read, which is fine, but the release job inherits broad top-level grants).

For a Go module this weekly prerelease channel has limited value anyway: Go consumers
can already get `main` via `go get ...@main` pseudo-versions, and prerelease tags
(`v2.1.0-pre.1`) are opt-in-only for `go get`.

## Recommendation

Pick one:

- **Retire the weekly preview at GA** (simplest; recommend this). Pseudo-versions
  cover the "try latest" use case.
- Or keep it, but move preview releases to a dedicated branch/tag prefix so the two
  channels never propose PRs against the same changelog, and document the channel in
  the Readme.

Either way, do it before v2.0.0 so launch week doesn't start with a Monday 3 AM
prerelease PR colliding with the GA release PR.
