# Release branch protection runbook

Workflow-level hygiene for `release/v*` PRs is enforced in-repo
(`.github/workflows/branch-policy.yml` now targets `release/**` too), but
the hard guarantees — no direct pushes, required reviews, protected history —
are repository rulesets, which are admin settings outside Actions. They are
declared in this repo as code, following the same schema Microsoft's Kiota
repositories use: `.github/policies/servicenow-sdk-go-branch-protection.yml`
(#658, ADR 011 rule 3). That file is the single source of truth — reviewed
in PRs like any other code — so the setup survives maintainer turnover.

## Ruleset definition (declared in `.github/policies/`)

Both require only universally reporting status checks ("Check Branch Name",
"Check Linked Issue", "Validate PR Title", "CodeQL") — path-filtered jobs
would sit pending forever on PRs whose paths they skip.

Apply the maintenance-line declaration with
`scripts/apply-maintenance-ruleset.sh`, which is idempotent
(create-or-update) and encodes exactly the spec below — prefer it over
clicking through the UI:

```bash
gh auth status  # must be an admin

# Bootstrap: create without status checks (none exist yet on a fresh branch)
scripts/apply-maintenance-ruleset.sh --dry-run   # inspect payload first
scripts/apply-maintenance-ruleset.sh

# After the bootstrap PR's CI has reported once, enforce the declared checks:
scripts/apply-maintenance-ruleset.sh \
  "Check Branch Name" "Check Linked Issue" "Validate PR Title" "CodeQL"
```

The table below is the human-readable spec both files encode; consult it to
review changes or recreate by hand if ever needed.

| Setting | Value |
| --- | --- |
| Targets | Branches matching pattern `release/v*` |
| Bypass actors | Repository admin (emergency use only) |
| Require pull request before merging | On |
| Required approvals | Match whatever `main` requires today |
| Dismiss stale reviews on new commits | Match `main` |
| Require status checks to pass | On (see below) |
| Require branches to be up to date | Match `main` |
| Require linear history | On |
| Block force pushes | On |
| Block deletions | On |

### Required status checks

Do not hand-enumerate check names at ruleset-creation time — a fresh
`release/vX.Y` branch has no runs yet, and GitHub only offers checks that
have already reported on the branch. Instead:

1. Cut the branch and open one throwaway PR against it.
2. Let `ci.yml`, `branch-policy.yml`, and `pr.yml` report once.
3. Edit the ruleset and require the checks that appeared:
   the two Branch Policy checks ("Check Branch Name", "Check Linked Issue"),
   "Validate PR Title", and the CI test/lint jobs relevant to Go changes.

## What workflows already cover (no admin action needed)

- `branch-policy.yml`: branch naming (`backport/…` heads are exempt as
  automation) and linked-issue checks now run on `release/**` PRs.
- `pr.yml`: PR title lint already targeted `main` and `release/**`.
- `ci.yml`: build/test/lint run for any PR touching Go paths, regardless of
  base branch; pushes to `release/**` were already covered.
- `labeler.yml` / `CODEOWNERS`: path-based and base-agnostic; the #658
  audit required no changes.

## Known gap

Direct pushes to `release/v*` bypass PR-event workflows entirely (nothing
runs, nothing posts) until the ruleset above blocks them. Apply the ruleset
**before** cutting the first maintenance branch — ADR 011 rule 3 assumes it.
