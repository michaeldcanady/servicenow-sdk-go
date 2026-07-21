# Branch Protection Audit: `main`

This is a point-in-time audit of `main`'s branch-protection configuration
against the job names actually produced by the workflows in
`.github/workflows/`. It was performed for issue #527 and should be re-run
(or at least sanity-checked) whenever a workflow's job `name:`/`id` changes,
or a job is added/removed/renamed.

## How this was checked

- `gh api repos/michaeldcanady/servicenow-sdk-go/branches/main`—checks
  whether `main` is a protected branch at all.
- `gh api repos/michaeldcanady/servicenow-sdk-go/branches/main/protection`—the
  classic branch-protection configuration, including
  `required_status_checks.contexts`.
- `gh api repos/michaeldcanady/servicenow-sdk-go/rulesets`—repository
  rulesets (the newer replacement/complement to classic branch protection),
  which can also carry a `required_status_checks` rule.

## Finding

As of this audit, **`main` has no branch protection configured at all**:

- `GET /branches/main` reports `"protected": false`.
- `GET /branches/main/protection` returns `404 Branch not protected`.
- `GET /rulesets` returns an empty list—no repository rulesets exist.

No configured list of required status checks exists to compare against
current job names, so there is no drift to fix. This is worth calling out
explicitly rather than silently closing the issue as "N/A," since the
absence of any required checks means PRs can currently be merged into `main`
regardless of whether `go.yml`, `pages.yml`, or `pull-request-title-lint.yml`
passed.

## Current job names (for when branch protection is configured)

If/when required status checks are configured for `main`, they should
reference the **check names** below (job `name:` where set, otherwise the
job id), not raw job ids, since jobs run from a matrix report one status
check per
matrix combination.

### `.github/workflows/go.yml` ("Go CI")

Gated by the `changes` job on `**.go` (excluding `docs/snippets/**`) and
`.github/workflows/go.yml`; skipped (not failed) on PRs that don't touch
those paths.

| Job id | Check names as reported to the Checks API |
| --- | --- |
| `changes` | `changes` |
| `check-go` | `Check Go Modules (stable)`, `Check Go Modules (oldstable)` |
| `build-go` | `Build Go Code (ubuntu-latest, stable)`, `Build Go Code (ubuntu-latest, oldstable)`, `Build Go Code (macos-latest, stable)`, `Build Go Code (macos-latest, oldstable)`, `Build Go Code (windows-latest, stable)`, `Build Go Code (windows-latest, oldstable)` |
| `lint-go` | `Lint Go Code (stable)`, `Lint Go Code (oldstable)` |
| `test-go` | `Test Go Code (ubuntu-latest, stable)`, `Test Go Code (ubuntu-latest, oldstable)`, `Test Go Code (macos-latest, stable)`, `Test Go Code (macos-latest, oldstable)`, `Test Go Code (windows-latest, stable)`, `Test Go Code (windows-latest, oldstable)` |

### `.github/workflows/pages.yml` ("Documentation CI/CD")

Gated by the `changes-docs` job on `docs/**`, `mkdocs.yml`, and
`.github/workflows/pages.yml`; skipped (not failed) on PRs that don't touch
those paths.

| Job id | Check name |
| --- | --- |
| `changes-docs` | `changes-docs` |
| `build` | `build-md` |
| `lint` | `lint-md` |
| `deploy` | `deploy-docs` (push-to-`main` only, not applicable to PR checks) |

### `.github/workflows/pull-request-title-lint.yml` ("Lint PR Title")

Not path-gated; runs on every PR.

| Job id | Check name |
| --- | --- |
| `main` | `Validate PR Title` |

### `.github/workflows/auto-merge-dependabot.yml`

Not a status check candidate for required checks—it merges PRs, it
doesn't gate them.

### CodeQL

No `codeql.yml` (or equivalent static-analysis) workflow currently exists
in this repository, so there is nothing to include here.

## Skip-vs-fail semantics

The `changes` / `changes-docs` gating jobs are unconditional (no `if:`), and
every downstream job in `go.yml` and `pages.yml` carries
`if: needs.changes(-docs).outputs.code/docs == 'true'`. When that condition
evaluates false, GitHub Actions reports the downstream job as **skipped**,
not failed. Per GitHub's documented behavior, a required status check that
is reported as "skipped" is treated as passing for merge purposes (unlike a
check that never runs at all, which blocks the merge button indefinitely).
This is standard behavior and was spot-checked against GitHub's current
documentation as part of this audit; no repo-side change is needed for it to
work correctly once required checks are configured.

## Recommendation

Configure a ruleset (or classic branch protection) on `main` that requires,
at minimum:

- `Validate PR Title`
- `Check Go Modules (stable)`, `Check Go Modules (oldstable)`
- `Build Go Code (ubuntu-latest, stable)` (and the other 5 matrix legs, or a
  narrower subset if full OS/version coverage isn't desired as a merge
  gate)
- `Lint Go Code (stable)`, `Lint Go Code (oldstable)`
- `Test Go Code (ubuntu-latest, stable)` (and the other 5 matrix legs)
- `build-md`, `lint-md` (only if docs-path PRs should be gated too)

This is a GitHub Settings change (Settings → Branches / Rulesets), not a
workflow-file change, and is out of scope for this audit; it's called out
here so the follow-up action is tracked.
