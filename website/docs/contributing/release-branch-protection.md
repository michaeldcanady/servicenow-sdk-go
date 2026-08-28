---
title: Release branch protection runbook
description: How release branch protection is declared as code in this repo and applied by hand when needed.
---

# Release branch protection runbook

Workflow-level hygiene for `release/v*` PRs is enforced in-repo
(`.github/workflows/branch-policy.yml` now targets `release/**` too), but
the hard guarantees — no direct pushes, required reviews, protected history —
are admin settings outside Actions. They're declared in this repo as code
(`.github/policies/servicenow-sdk-go-branch-protection.yml`, #658, ADR 011
rule 3) **and** applied in GitHub as live settings so OpenSSF Scorecard can
detect them. The file is the reviewed source of truth; the live settings are
what Scorecard evaluates — both must stay aligned.

## How protection is applied (rulesets, Scorecard-compliant)

Scorecard v5.5.0 (`ossf/scorecard-action@v2.4.4` in `.github/workflows/scorecard.yml`)
detects **both** classic branch protection (REST `branches/*/protection`) and
repository rulesets (GraphQL `rulesets`) — see
`clients/githubrepo/branches.go` (`rulesets` + `branchProtectionRules`
queries; rulesets are matched via `fnmatch` on `refs/heads/<branch>`). The
repo now uses **rulesets only** (classic was removed after migration) — v5.5.0
correctly detects them; older Scorecard would have required classic.

Live state as of #725 (verify with commands under Verification):

| Branch / pattern | Ruleset | Id | Where to edit |
| --- | --- | --- | --- |
| `main` | `servicenow-sdk-go-branch-protection-main` | `21693268` (renamed from `main-2`) | Settings → Rules → Rulesets |
| `release/v*` (pattern, covers `release/v1.0` and future `release/vX.Y`) | `servicenow-sdk-go-branch-protection` (original `servicenow*` base, id `21692927`) | `21692927` | Settings → Rules → Rulesets |

Both are the consolidated successors of the fragmented import (`main-1`, `release/v1.0-*` removed). The `release/v*` wildcard covers existing `release/v1.0` without a per-branch ruleset. Keep live rulesets in sync with the declarative file when it changes.

## Ruleset definition (declared in `.github/policies/`)

`servicenow-sdk-go-branch-protection.yml` declares two branch-protection
rules using the `resource: repository` / `configuration.branchProtectionRules`
schema:

| Rule | Branches | Guarantees |
| --- | --- | --- |
| trunk | `main` | PR + 1 CODEOWNERS approval, stale-review dismissal, conversation resolution, no deletion/force-push |
| maintenance lines | `release/v*` | same as trunk **plus required linear history** |

Both require only universally reporting status checks (`Check Branch Name`,
`Check Linked Issue`, `Validate PR Title`, `CodeQL`) — path-filtered jobs
would sit pending forever on PRs whose paths they skip.

The table below is the human-readable spec the declaration encodes; consult
it to review changes or apply the settings by hand. For `main`, the Settings
→ Branches UI maps one-to-one; for `release/v*`, use Settings → Rules →
Rulesets.

| Setting | Value |
| --- | --- |
| Targets | `main` (`refs/heads/main`) and `release/v*` (`refs/heads/release/v*`) |
| Bypass actors | Repository admin (`actor_id: 5`, `RepositoryRole`, `bypass_mode: always`) — allows single maintainer to merge without a second approval (Scorecard EnforceAdmins then `false`, hence 8/10 not 10) |
| Require pull request before merging | On |
| Required approvals | 1 (admin bypass means admin can merge with 0; non-admin needs 1) |
| Dismiss stale reviews on new commits | On |
| Require status checks to pass | On (see below) |
| Require branches to be up to date | **On** (`strict_required_status_checks_policy: true`) — Scorecard-compliant (Tier 2); admin bypass avoids blocking emergency merges |
| Require last push approval | **On** (`require_last_push_approval: true`) — Scorecard-compliant (Tier 2) |
| Require linear history | Off for `main`, **On** for `release/v*` (cherry-pick backports must stay one-to-one) |
| Block force pushes | On (`non_fast_forward`) |
| Block deletions | On (`deletion`) |
| Require conversation resolution | On |
| Require code owners review | On |

### Required status checks

Don't hand-enumerate check names at ruleset-creation time — a fresh
`release/vX.Y` branch has no runs yet, and GitHub only offers checks that
have already reported on the branch. Instead:

1. Cut the branch and open one throwaway PR against it.
2. Let `codeql.yml`, `branch-policy.yml`, and `pr.yml` report once.
3. Edit the ruleset and require the checks that appeared — exactly the four
   declared in `.github/policies/servicenow-sdk-go-branch-protection.yml`
   ("Check Branch Name," "Check Linked Issue," "Validate PR Title," "CodeQL.")
   Never add path-filtered jobs (the ci.yml build/test/lint matrix): they
   skip docs-only or workflow-only PRs, and a required check that never
   reports sticks "pending" forever, blocking every merge until an admin
   bypasses.

### Ruleset payload reference (live)

`main` (`servicenow-sdk-go-branch-protection-main`, `refs/heads/main`):

```json
{
  "conditions": {"ref_name": {"include": ["refs/heads/main"]}},
  "bypass_actors": [{"actor_id": 5, "actor_type": "RepositoryRole", "bypass_mode": "always"}],
  "rules": [
    {"type": "deletion"},
    {"type": "non_fast_forward"},
    {"type": "pull_request", "parameters": {"required_approving_review_count": 1, "dismiss_stale_reviews_on_push": true, "require_code_owner_review": true, "require_last_push_approval": true, "required_review_thread_resolution": true}},
    {"type": "required_status_checks", "parameters": {"strict_required_status_checks_policy": true, "required_status_checks": [{"context": "Check Branch Name"}, {"context": "Check Linked Issue"}, {"context": "Validate PR Title"}, {"context": "CodeQL"}]}}
  ]
}
```

`release/v*` (`servicenow-sdk-go-branch-protection`, `refs/heads/release/v*`, the original `servicenow*` base) — same as `main` **plus** `required_linear_history`:

```json
{
  "conditions": {"ref_name": {"include": ["refs/heads/release/v*"]}},
  "bypass_actors": [{"actor_id": 5, "actor_type": "RepositoryRole", "bypass_mode": "always"}],
  "rules": [
    {"type": "deletion"},
    {"type": "non_fast_forward"},
    {"type": "required_linear_history"},
    {"type": "pull_request", "parameters": {"required_approving_review_count": 1, "dismiss_stale_reviews_on_push": true, "require_code_owner_review": true, "require_last_push_approval": true, "required_review_thread_resolution": true}},
    {"type": "required_status_checks", "parameters": {"strict_required_status_checks_policy": true, "required_status_checks": [{"context": "Check Branch Name"}, {"context": "Check Linked Issue"}, {"context": "Validate PR Title"}, {"context": "CodeQL"}]}}
  ]
}
```

The declarative file is still the source of truth for review — the live payloads above are its applied form. The `bypass_actors` entry is why Scorecard reports `8/10` not `10`: `EnforceAdmins` is `false` whenever a bypass exists (`branches.go:applyRepoRules` sets `EnforceAdmins = len(bypass)==0`). Keeping `bypass:always` for `admin` lets a single maintainer merge without a second approval, at the cost of Tier 5.

## Scorecard alignment

- **v5.5.0 detection:** `BranchProtection` raw probe calls `GetDefaultBranch` + `ListReleases` to enumerate branches, then for each does a GraphQL `rulesets` query (available with `contents: read`) plus `branchProtectionRules`/`RefUpdateRule`. Matching rulesets are merged onto the `BranchRef` in `branches.go:applyRepoRules`. Scorecard therefore scores **the live GitHub settings, not the file** — the file alone leaves the check at 0. #725 verified v5.5.0 *does* detect rulesets (`gh api repos/.../rulesets` and `repository.rulesets`); classic is no longer needed.
- **Score achieved:** with both rulesets `strict:true` + `last_push:true`, Scorecard Branch-Protection is **8/10** on `v5.5.0` (`score is 8: branch protection is not maximal … required approving review count is 1 on branch 'main'; 'branch protection settings apply to administrators' is disabled`). Tier 1-3 pass. Reaching 9 needs `count:2`, 10 needs `count:2` **and** no bypass (`EnforceAdmins:true`). The repo stays at `count:1` with `admin(always)` bypass so a single maintainer can merge without a peer (Tier 4/5) — see #724 for two-person.
- **Why not 4/10:** before the #725 update, `strict:false` + `last_push:false` gave `score is 4` with additional warns for `up-to-date` and `last push` (Tier 2). Updating those two flags to `true` lifted the score to `8` while retaining the `1`-approval admin bypass.

## Verification

Run after any change to the declarative file or live settings:

```bash
# Rulesets — both patterns
gh api repos/michaeldcanady/servicenow-sdk-go/rulesets \
  --jq '.[] | {id, name, enforcement, target, conditions, bypass_actors, rules: [.rules[].type]}'
gh api repos/michaeldcanady/servicenow-sdk-go/rulesets/21692927 \
  --jq '{name, conditions, rules, bypass_actors}'
gh api repos/michaeldcanady/servicenow-sdk-go/rulesets/21693268 \
  --jq '{name, conditions, rules, bypass_actors}'

# GraphQL view Scorecard actually uses (rulesets + classic)
gh api graphql -f query='{repository(owner:"michaeldcanady", name:"servicenow-sdk-go"){rulesets(first:10){nodes{name enforcement target conditions{refName{include}} rules(first:10){nodes{type}}}} branchProtectionRules(first:10){nodes{pattern}}}}'

# Classic should be 404 (migrated to rulesets)
gh api repos/michaeldcanady/servicenow-sdk-go/branches/main/protection --jq . 2>&1 | head -n 5

# Scorecard itself (workflow_dispatch, then check code-scanning alert 18)
gh workflow run scorecard.yml --repo michaeldcanady/servicenow-sdk-go
gh api repos/michaeldcanady/servicenow-sdk-go/code-scanning/alerts/18 \
  --jq '.most_recent_instance.message.text'
```

Expect `BranchProtectionID` to be `score is 8: branch protection is not maximal … required approving review count is 1` (or `10` after #724), never `score is 0: branch protection not enabled for branch 'main'`. Before the strict/lastPush fix it was `score is 4` with extra `up-to-date`/`last push` warns.

## Bypass actors

The human spec lists a repository-admin bypass for emergency use. The
`branchProtectionRules` schema in the declarative file has no field for it —
bypass actors can only be configured when applying the live rulesets:

- Both `servicenow-sdk-go-branch-protection-main` and `servicenow-sdk-go-branch-protection` use `bypass_actors: [{"actor_id": 5, "actor_type": "RepositoryRole", "bypass_mode": "always"}]` (role `5` = Repository admin). This lets a single maintainer merge without a second approval — at the cost of Scorecard `EnforceAdmins:false` (hence `8/10` not `10`). The declared file is therefore the source of truth for everything except bypasses; record any bypass grant here in this runbook when it happens.
- Previously classic used `enforce_admins: false` for the same effect; now rulesets carry the bypass.

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

## Release automation rehearsal (must still function under protection)

Protection must not break `release-please`, `dependabot`, or backport
automation. Rehearse after any live change:

1. **Dry-run the declarative file:** `python3 -m json.tool .github/policies/servicenow-sdk-go-branch-protection.yml` — confirms schema parse.
2. **Open a throwaway PR** against `main` and against `release/v1.0` (e.g. `chore/rehearse-branch-protection (#725)`) and verify the four required checks report and the PR is mergeable only with 1 CODEOWNERS approval and conversation resolution. With `strict:true`, the branch must be up-to-date before merging (rebase if behind). Close without merging.
3. **Backport label:** label a dummy `main` PR with `backport release/v1.0` and confirm `backport.yml` would fan out (check `Select existing target branches` log); no actual backport PR needed.
4. **Stable/maintenance release configs:** `stable-release.yml` (main) and `maintenance-label.yml`/`backport.yml` are branch-scoped and use `contents: write`/`pull-requests: write` — they push to `release-please--` branches and open PRs, never directly to `main`/`release/v*`, so the `admin(always)` bypass keeps them unblocked (admin actor bypasses the PR approval, not the branch creation). No `restrictions` (push allowlist) is set.

If any job would have been skipped as `pending` forever (path-filtered `ci.yml` matrix), it must **not** be a required status check — see Required status checks above.

## Known gap

Direct pushes to `release/v*` bypass PR-event workflows entirely (nothing
runs, nothing posts) until the ruleset above blocks them. Apply the ruleset
**before** cutting the first maintenance branch — ADR 011 rule 3 assumes it.
The `release/v*` wildcard ruleset (`21692927`) now covers any future line without extra API
calls at cut time; no per-branch classic protection remains.
