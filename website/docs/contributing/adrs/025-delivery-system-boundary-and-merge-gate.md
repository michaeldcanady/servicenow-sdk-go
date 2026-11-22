---
title: 'ADR 025: Three required checks gate every merge, and external systems have named owners'
description: >-
  PR Gate, PR Policy, and CodeQL are the only required checks. No approving
  review is required, the admin bypass stays for logged emergencies, and
  rulesets live as code in infra-terraform.
---

# ADR 025: Three required checks gate every merge, and external systems have named owners

## Status

Proposed (2026-11-21). The plan is in
`specs/010-holistic-delivery-pipeline/spec.md`.

## Context

A green PR should be enough evidence that the PR is safe to merge, with human
review limited to the code. Today it isn't:

- `main` has no branch rules. `gh api .../rules/branches/main` returns `[]`,
  and classic branch protection returns 404. The runbook
  (`release-branch-protection.md`) and
  `.github/policies/servicenow-sdk-go-branch-protection.yml` describe
  rulesets that aren't live.
- Even the documented rules require no Go check. The Go, docs, and license
  jobs are path-filtered, so requiring them by name would leave skipped jobs
  pending forever. A PR with failing unit tests can merge.
- Required check names have broken merges before. In #917, a reusable
  workflow's `caller / called` name prefix changed and every merge blocked.

The repo has one maintainer, working with AI agents. A required approving
review would block every PR the maintainer opens, because GitHub doesn't let
authors approve their own PRs. CodeRabbit reviews every PR, but it has write
access that its config can't narrow.

Several parts of the delivery system live outside this repo: the rulesets and
merge settings (managed in `NerdIT-Tech/infra-terraform`), Jira Automation,
the GitHub for Jira app, CodeRabbit, GitGuardian, Codecov, and the e2e
ServiceNow instance. A guarantee that a green PR is safe depends on them, so
the repo needs a record of who owns each one and how it connects.

Alternatives considered:

1. **Require each job by name.** This is the documented state. Rejected:
   path filters make it impossible, and every matrix change becomes a ruleset
   change.
2. **A gate job that polls the Checks API until other workflows finish.**
   Rejected: it depends on timing, uses runner minutes while waiting, and
   can't tell a check that hasn't started from one that never starts.
3. **Require an approving review, and let CodeRabbit's approval count.**
   Rejected: an app with write access could then unblock merges by itself.
4. **Require a human approving review.** Rejected: it needs a second
   maintainer.
5. **Remove the admin bypass.** Rejected: a CI outage would then block every
   merge, and only the maintainer can fix it.
6. **Cover only the files in this repo.** Rejected: the gate that enforces
   the guarantee lives in another repo.

## Decision

1. **Three required checks.** The `main` and `release/v*` rulesets require
   exactly `PR Gate`, `PR Policy`, and `CodeQL`.
   - `PR Gate` is a plain top-level job in `pr-gate.yml`, not a
     reusable-workflow call, so its name has no prefix. The workflow has no
     workflow-level `paths:` filter. A `changes` job decides which checks
     run, and it turns every check on when detection fails. `PR Gate` passes
     only when every needed job passed or was skipped on purpose.
   - `PR Policy` checks the PR title, the linked work item, and, for PRs into
     `release/*`, the forward-port assessment from
     [ADR 023](023-manual-backports-and-forward-port-assessments.md). It runs
     on `edited`, `labeled`, and `unlabeled` events, separate from the gate.
     A title edit then doesn't cancel a test run.
   - `CodeQL` is reported by the code-scanning app.
2. **Strict status checks.** A branch must be up to date with its base before
   it merges, so the tested tree is the merged tree.
3. **No required approvals.** CodeRabbit's review is advice. Agents may
   review PRs, but they never push to or approve a PR in CI.
4. **The admin bypass stays, and each use is logged.** The log lives on the
   release-branch-protection page: the date, the PR, and the reason.
5. **Rulesets and merge settings live as code in
   `NerdIT-Tech/infra-terraform`.** That source sets the required checks,
   strict status checks, the squash commit title (the PR title), rebase
   merging (off), the admin bypass, and a tag ruleset for `v*`. The runbook
   and `.github/policies/` file in this repo shrink to a pointer to it, plus
   the bypass log.
6. **External systems have named owners.** The external-systems
   table in the spec lists each system's owner and its interface with this repo. A change
   to an external system goes through that system's own channel. This repo
   implements only what lives in it.

## Consequences

- **Pros:**
  - One stable set of check names. Adding or removing a job changes the gate,
    not the ruleset.
  - Squash titles equal the validated PR title, which is the commit message
    release-please reads.
  - The bypass log shows exactly where the guarantee didn't hold.
- **Cons:**
  - The merge guarantee depends on a repo that this repo's CI can't check.
    A ruleset that drifts in infra-terraform breaks the guarantee without a
    signal here. The rule file's `Verify` command reads the live ruleset to
    catch this.
  - Strict status checks re-run the gate after every merge to the base.
    Merge queue would avoid that, but GitHub offers it only for repositories
    that an organization owns
    ([ADR 010](010-no-nerdit-tech-migration.md)).
  - A live release branch reports `PR Gate` only after the workflow is
    cherry-picked to it, so the ruleset change for `release/v*` must wait for
    that.
- **Rule for future CI work:** don't add a required check outside these
  three. Put a new content check in `PR Gate`, and a new metadata check in
  `PR Policy`.
