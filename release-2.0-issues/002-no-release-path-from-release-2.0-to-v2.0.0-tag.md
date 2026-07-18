# No defined release path from `release/2.0` to a v2.0.0 tag

- **Priority:** P0 — release blocker
- **Raised by:** Senior DevOps + Senior Product Manager
- **Area:** Release automation

## Problem

`stable-release.yml` runs release-please only on pushes to `main`. The 2.0 work lives on
`release/2.0`, currently **90 commits ahead of main**, and nothing in the automation
knows about that branch. There is no documented or automated way for this branch to
become the `v2.0.0` tag.

Two specific hazards when the merge to `main` happens:

1. **Merge strategy determines what release-please parses.** With a merge commit,
   release-please walks the individual commits — and the branch contains
   non-conventional messages that will parse wrong or not at all:
   - `d863ddb major` (not conventional at all)
   - `7e31c01 fix: merge errors` (parses as a patch fix)
   - `95ee5a9 refator: model accessor mutators use store (#474)` (typo — `refator` is
     not a recognized type)
2. **Nothing on the branch carries a `BREAKING CHANGE:` footer** (spot-check the 90
   commits). Without one, release-please will compute a **minor** bump (v1.13.0), not
   v2.0.0, and pkg.go.dev will index a breaking release under a v1 tag — the worst
   possible outcome for consumers.

Also note: PRs targeting `release/2.0` bypass `pr.yml` title lint (it filters
`branches: [main]`), which is how the malformed titles got in.

## Recommendation

Write a short release runbook (can live in this directory) and follow it:

1. Merge `release/2.0` into `main` via a **single squash PR** whose title is
   conventional (e.g. `feat!: v2 rework`) and whose body carries a
   `BREAKING CHANGE:` footer summarizing the v2 changes.
2. Verify the resulting release-please PR proposes **2.0.0** before merging it.
3. Sequence this with the `/v2` module-path change (see issue 003) — the tag must not
   land before the module path is correct.
4. Extend `pr.yml`'s branch filter to include `release/*` so future release branches get
   title linting.
