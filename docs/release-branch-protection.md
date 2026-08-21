# Release branch protection runbook

Workflow-level hygiene for `release/v*` PRs is enforced in-repo
(`.github/workflows/branch-policy.yml` now targets `release/**` too), but
the hard guarantees — no direct pushes, required reviews, protected history —
are repository rulesets, which are admin settings outside Actions. This
runbook captures the exact setup so it survives maintainer turnover
(#658, ADR 011 rule 3).

## One-time ruleset setup (admin: Settings → Rules → Rulesets)

The ruleset is codified in
`scripts/apply-maintenance-ruleset.sh` — prefer it over clicking through the
UI, since it is idempotent (create-or-update) and reviewable:

```bash
# Bootstrap: create without status checks (none exist yet on a fresh branch)
gh auth status  # must be an admin
scripts/apply-maintenance-ruleset.sh --dry-run   # inspect payload first
scripts/apply-maintenance-ruleset.sh

# After the bootstrap PR's CI has reported once, enforce checks:
scripts/apply-maintenance-ruleset.sh \
  "Check Branch Name" "Check Linked Issue" "Validate PR Title" \
  "test (stable)" "lint"
```

The table below is the human-readable spec the script encodes; consult it to
review changes or recreate by hand if ever needed.

Create a ruleset named `maintenance-lines`:

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
