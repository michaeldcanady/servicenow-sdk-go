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

`servicenow-sdk-go-branch-protection.yml` declares two branch-protection
rules using the `resource: repository` / `configuration.branchProtectionRules`
schema:

| Rule | Branches | Guarantees |
| --- | --- | --- |
| trunk | `main` | PR + 1 CODEOWNERS approval, stale-review dismissal, conversation resolution, no deletion/force-push |
| maintenance lines | `release/v*` | same as trunk **plus required linear history** |

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
2. Let `codeql.yml`, `branch-policy.yml`, and `pr.yml` report once.
3. Edit the ruleset and require the checks that appeared — exactly the four
   declared in `.github/policies/servicenow-sdk-go-branch-protection.yml`
   ("Check Branch Name", "Check Linked Issue", "Validate PR Title", "CodeQL").
   Never add path-filtered jobs (the ci.yml build/test/lint matrix): they
   skip docs-only or workflow-only PRs, and a required check that never
   reports sticks "pending" forever, blocking every merge until an admin
   bypasses.

## Bypass actors

The human spec below lists a repository-admin bypass for emergency use, but
the `branchProtectionRules` schema has no field for it — bypass actors can
only be configured when applying the ruleset in repo settings. The declared
file is therefore the source of truth for everything except bypasses; record
any bypass grant here in this runbook when it happens.

## What workflows already cover (no admin action needed)

- `branch-policy.yml`: branch naming (`backport/…` heads are exempt as
  automation) and linked-issue checks now run on `release/**` PRs.
- `pr.yml`: PR title lint already targeted `main` and `release/**`.
- `codeql.yml`: analyzes `release/**` PRs too, so "CodeQL" can report on
  maintenance-line merges (a required check must first be able to appear).
- `ci.yml`: build/test/lint run for any PR touching Go paths, regardless of
  base branch; pushes to `release/**` were already covered.
- `labeler.yml` / `CODEOWNERS`: path-based and base-agnostic; the #658
  audit required no changes.

## Known gap

Direct pushes to `release/v*` bypass PR-event workflows entirely (nothing
runs, nothing posts) until the ruleset above blocks them. Apply the ruleset
**before** cutting the first maintenance branch — ADR 011 rule 3 assumes it.
