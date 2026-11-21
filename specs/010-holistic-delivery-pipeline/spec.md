# A holistic delivery pipeline

Status: draft. Decisions settled in a four-round interview with the
maintainer on Nov 21, 2026. The decision records are ADRs 024 to 028:

- [ADR 024](../../website/docs/contributing/adrs/024-no-long-lived-ci-secrets.md):
  no long-lived secrets in CI, with a list of exceptions that expire.
- [ADR 025](../../website/docs/contributing/adrs/025-delivery-system-boundary-and-merge-gate.md):
  the delivery system's boundary, its owners, and the merge gate.
- [ADR 026](../../website/docs/contributing/adrs/026-tool-version-pins-and-drift-check.md):
  tool-version pins and the drift check.
- [ADR 027](../../website/docs/contributing/adrs/027-release-verification-without-signing-or-sbom.md):
  releases without signing or SBOM. Replaces ADR 012 and withdraws ADR 013.
- [ADR 028](../../website/docs/contributing/adrs/028-api-compatibility-check.md):
  the API-compatibility check.

This spec folds in the PR pipeline plan from the `principal-devops-engineer`
agent. [Items dropped from the DevOps plan](#items-dropped-from-the-devops-plan)
lists the parts that no longer apply.

## Problem

The repo's CI and delivery workflows grew one change at a time. Each change
made sense alone, but nothing ties them to a shared goal, so two problems
follow.

**Maintenance churn.** Since Aug 1, 2026, about 154 of 343 commits were
`chore(ci)`, `chore(tooling)`, `chore(docs)`, or `chore(deps)`. Many of them
fix an earlier pipeline change. For example, #960 and #962 fix the
docs-archive change in #959, and several commits fix zizmor check names and
version comments. Before the cleanup in #976 to #985, there were 42 workflow
files and 8 scheduled jobs. Pipeline logic lives in seven places:
`mise.toml`, `.mise/tasks/`, `scripts/`, `.github/scripts/`,
`.github/actions/`, `.github/policies/`, and `.githooks/`.

**Merge gates you can't trust.** A green PR doesn't mean the PR is safe to
merge:

- `main` has no branch rules. `gh api .../rules/branches/main` returns `[]`.
  The runbook and `.github/policies/` describe rulesets that aren't live.
- Even the documented rules require no Go check. Build, tests, lint,
  `govulncheck`, the docs build, Vale, and license headers are all
  path-filtered, so a PR with failing unit tests can merge.
- PRs that release-please opens with `GITHUB_TOKEN` show no Actions checks
  unless someone approves their runs.
- Dependabot auto-merge is off, because without rules nothing makes it wait
  for CI.

Recent commits on `main` already removed the weekly preview release (#976),
the branch-name check (#981), the backport and forward-port workflows (#982),
SBOM (#983), project sync and its personal access token (#984), and signing
(#985). This spec records why, and sets the target state that the remaining
work moves toward.

## Goals and non-goals

### Goals

- A green required check suite is enough evidence that a PR is safe to merge.
  Human review covers only the code.
- Merging the release-please PR is the only human step in a release. Approving
  the workflow runs on that PR is part of the same step.
- Each piece of pipeline logic has one home, one owner, and one stated reason
  to exist.
- PR checks finish in under 10 minutes at the median.
- CI stores no long-lived secret except a documented exception that expires.
- The pipeline serves one maintainer working with AI agents. External
  contributors can still open PRs from forks without extra work from the
  maintainer.

### Non-goals

- Signing release artifacts or generating an SBOM. ADR 027 records why.
- Weekly preview releases. #976 removed them.
- Automated backports or a forward-port tracker. ADR 023 makes both manual.
- Merge queue. GitHub offers it only for repositories that an organization
  owns, and ADR 010 keeps this repo under a personal account through v3.
- Required approving reviews. A sole maintainer can't approve their own PRs.
- Changing systems outside this repo from this repo. This spec documents
  their owners and interfaces only.
- Moving CI tool steps onto `mise` tasks. ADR 026 records that choice.

## Design

### Principles

- **The gate is the contract.** Three required checks decide whether a PR can
  merge: `PR Gate`, `PR Policy`, and `CodeQL`. A job that matters is part of
  `PR Gate`, or it doesn't block a merge.
- **Every workflow earns its place.** A workflow's header comment states its
  purpose, its trigger, and what it catches that no other check catches. A
  scheduled job must catch something that PR and release checks can't.
- **No stored credentials by default.** Use `GITHUB_TOKEN`, OpenID Connect
  (OIDC), or GitHub's built-in approval flow. A stored secret needs an entry
  in ADR 024's exceptions list.
- **Small surface over clever automation.** When automation doesn't work or
  costs more than it saves, remove it and write down the manual step. The
  backport, forward-port, project sync, preview, SBOM, and signing removals
  follow this principle.

### When a PR is done

A PR can merge when all of the following are true:

- `PR Gate` passes. It runs every content check that applies to the change:
  Go build and tests, the race detector, lint, `go mod` checks,
  `govulncheck`, API compatibility, the docs build, Vale, snippet checks,
  license headers, workflow lint, zizmor, dependency review, and the tool-pin
  drift check.
- `PR Policy` passes. It checks the PR title, the linked work item, and, for
  PRs into `release/*`, the forward-port assessment that ADR 023 requires.
- `CodeQL` passes.
- The branch is up to date with its base, because the ruleset uses strict
  status checks.

No approving review is required. CodeRabbit reviews every PR, and its review
is advice. Agents may review PRs, but they never push to or approve a PR in
CI.

An admin can bypass the rules in an emergency. Each bypass gets a line in the
bypass log on the release-branch-protection page: the date, the PR, and the
reason.

### When a release is done

1. Commits land on `main` (or on `release/vX.Y`) through gated PRs.
2. release-please opens or updates its release PR. GitHub creates the
   workflow runs on that PR in an approval-required state, because `GITHUB_TOKEN` opened
   the PR.
3. The maintainer selects **Approve workflows to run**, waits for the gate,
   and merges.
4. `stable-release.yml` creates the tag and the GitHub release. It then sends
   the version to the Jira release webhook (ADR 022, decision 10), and
   `docs-stable` and `docs-archive` publish the docs.

`stamp-deprecations.yml` pushes a commit to the release PR, which also
creates an approval-required run. Approve it the same way.

`release-verify.yml` stays until both of these are true. Then it's removed
(ADR 027):

- `PR Gate` is required on release PRs, and strict status checks are on. The
  merged tree then equals the tested tree.
- A tag ruleset blocks hand-pushed `v*` tags.

### Credentials

Today's stored secrets and what happens to each one:

| Secret or variable | Used by | Fate |
| --- | --- | --- |
| `SYNC_PROJECT_SYNC_TOKEN_PROVIDER_PEM` (GitHub App private key) and the `SYNC_PROJECT_SYNC_TOKEN_PROVIDER_APP_ID` variable | Nothing since #984 | Delete after a final usage check. Uninstall the project-sync App. |
| `TOKEN` (from 2023) | Nothing on `main` | Delete after a final usage check. |
| `CODECOV_TOKEN` | Nothing. Coverage upload uses OIDC. | Delete after a final usage check. |
| `SN_INSTANCE`, `SN_USERNAME`, `SN_PASSWORD`, `SN_CLIENT_ID`, and `SN_CLIENT_SECRET` in the `e2e-tests` environment | `e2e-nightly.yml` | Exception 1 in ADR 024. |
| Jira Automation webhook secret (name not chosen) | The release step in ADR 022, decision 10 | Exception 2 in ADR 024, once the step exists. |

Two spikes look for ways to retire the exceptions:

- **e2e tests over OIDC.** ADR 014 lets the SDK present an external OIDC
  token to ServiceNow verbatim. Test whether the e2e instance can trust the
  GitHub Actions OIDC issuer, so the nightly run authenticates with no stored
  credential. If it works, the SDK also gets a live test of its own feature.
- **Jira releases from deployment events.** The GitHub for Jira app reports
  GitHub Actions deployment events. Test whether a Jira Automation rule can
  release a Fix version from one, which removes the webhook secret.

For bot PRs, the design uses GitHub's built-in approval instead of a new
identity. A later spike tests a fully automatic path: `stable-release.yml`
and `stamp-deprecations.yml` start `PR Gate` with `workflow_dispatch` on the
release-please branch. GitHub documents that `workflow_dispatch` events from
`GITHUB_TOKEN` always create runs. It doesn't document whether a check from a
dispatched run satisfies a required check on the PR, so test that on a scratch PR
first. ADR 024 records the rejected options: octo-sts and a stored App key.

### Tool versions

`mise.toml` is the one source of tool versions. CI keeps its native actions,
because Dependabot keeps their pins current and they add PR annotations
(ADR 026).

- Where an action reads a version file that `mise` also reads, point the
  action at that file. `release-verify.yml` and `vet-go-snippets` already use
  `go-version-file: go.mod`. Node is a candidate. See
  [Open questions](#open-questions).
- For the remaining duplicated pins, a drift check fails the PR when a
  workflow's pin differs from `mise.toml`. Today these are golangci-lint
  2.14.0, Vale 3.18.0, zizmor 1.29.0, and Node 20.
- The CI Go matrix uses the `stable` and `oldstable` aliases on purpose. It
  stays out of the drift check. The `go` pin in `mise.toml` tracks `stable`.
- `govulncheck` and `tparse` install `latest` in CI but are pinned in
  `mise.toml`. Either pin them in CI as well and add them to the drift check,
  or record in `mise.toml` that CI tracks `latest` on purpose.

### Workflow layout

- Use one entry workflow per job to be done. Here is the target set, with
  about 13 entries against a soft cap of 12. Add an entry only with a reason
  that no existing entry covers:

  | Entry workflow | Job to be done |
  | --- | --- |
  | `pr-gate.yml` | The content gate on every PR |
  | `pr-policy.yml` | PR metadata: title, work item, and forward-port assessment |
  | `codeql.yml` | Code scanning |
  | `ci.yml` | Push and scheduled runs of the full matrix |
  | `stable-release.yml` | release-please for `main` and `release/v*` |
  | `stamp-deprecations.yml` | Stamp `{unreleased}` on the release PR |
  | `docs-preview.yml` | PR doc previews on `gh-pages` |
  | `docs-stable.yml` | The production docs deploy |
  | `docs-archive.yml` | Per-release docs archives from tags |
  | `e2e-nightly.yml` | Live e2e tests |
  | `scorecard.yml` | OpenSSF Scorecard |
  | `zizmor.yml` | Scheduled workflow security analysis |
  | `labeler.yml` | PR labels by path |

- Add a `reusable-*.yaml` file only when two or more workflows call it. Fold
  each single-caller reusable workflow back into its caller.
- Start every workflow with a header comment that states its purpose, its
  trigger, and what it catches that no other check catches.
- `labeler.yml` stays.

### Docs publishing

All three flows stay: PR previews on `gh-pages`, the stable deploy, and the
archive builds from tags. The docs build check moves into `PR Gate`. It builds
with the same `archived-versions` input that the stable deploy uses, so the PR
build matches production. When the Vale config or styles change, the gate runs
Vale on the whole site. For other changes, it lints only the changed files.

### Pre-commit hook

The hook stays optional. It runs only fast steps that fix the problem
themselves: license headers (#980) and `gofmt -s` on staged Go files. Slow or
noisy checks, such as Vale and golangci-lint, stay in `mise run check` and in
CI. The devcontainer already runs `mise run hooks:install` on creation, so
agents that commit inside it get the hook.

### Maintenance lines

ADR 023 stands. Backports are manual cherry-pick PRs into `release/vX.Y`, and
each PR into `release/*` carries a forward-port assessment. No required
review exists to enforce it, so `PR Policy` enforces it instead. A PR into
`release/*` must contain one of these lines:

- `Backport of #N`
- `Won't port:` followed by a reason

For the existing branches:

- Delete `release/v2.0`. Version 2.1 is the latest minor, so under ADR 011
  rule 7, 2.0 receives no further fixes.
- Keep `release/v1.0` until the v1 end-of-life announcement. Until then, it
  receives critical and security fixes only.
- Cut `release/v2.1` lazily, only when a fix first needs it.

### External systems

The delivery system includes pieces that live outside this repo. This repo
documents each one's owner and interface, and changes each one only through
its own channel.

| System | Owner | Interface with this repo |
| --- | --- | --- |
| Branch and tag rulesets, merge settings | Maintainer, as code in `NerdIT-Tech/infra-terraform` | Required checks `PR Gate`, `PR Policy`, and `CodeQL`. Strict status checks. Squash title set to the PR title. Rebase merging off. Admin bypass. A tag ruleset for `v*`. |
| Jira project `SNSFG1` and its Automation rules | Maintainer | The release webhook (ADR 022, decision 10). The work-item key in `PR Policy` after cutover. |
| GitHub for Jira app | Maintainer | PR and deployment events into Jira |
| CodeRabbit | Maintainer, through `.coderabbit.yaml` | Advisory PR review |
| GitGuardian | Maintainer | PR secret scan. Optional required check. |
| Codecov | Maintainer, through `.codecov.yml` | Coverage report over OIDC. Informational only. |
| The e2e ServiceNow instance | Maintainer | `SN_*` credentials, ADR 024 exception 1 |

The release-branch-protection runbook and
`.github/policies/servicenow-sdk-go-branch-protection.yml` shrink to a
pointer to the Terraform source, plus the bypass log.

## Phased plan

Effort: S is under half a day, M is one to two days, and L is more than two
days. The `G` numbers refer to gaps in the DevOps plan.

### Phase 0: make enforcement real (start now)

| # | Change | Where | Effort |
| --- | --- | --- | --- |
| 0.1 | Restore the `main` and `release/v*` rulesets with today's checks. Keep strict status checks on. Keep the admin bypass. Require no approvals. Verify with `gh api repos/michaeldcanady/servicenow-sdk-go/rules/branches/main`. (G1, G24) | infra-terraform | S |
| 0.2 | Set the squash commit title to the PR title and turn off rebase merging. (G14) | infra-terraform | S |
| 0.3 | Shrink the runbook and the policy file to a pointer and a bypass log. (G3) | This repo | S |
| 0.4 | Rely on the built-in approval for release-please and stamp runs. Correct the `ci.yml` comment that says release PRs already pass the same gates. On the next release PR, confirm that the **Approve workflows to run** banner appears. (G4, G5) | This repo | S |
| 0.5 | Delete `SYNC_PROJECT_SYNC_TOKEN_PROVIDER_PEM`, its `_APP_ID` variable, `TOKEN`, and `CODECOV_TOKEN` after a final usage check. Uninstall the project-sync App. | Settings | S |
| 0.6 | Delete `release/v2.0` after you confirm it has no open PRs. | Settings | S |

### Phase 1: close coverage gaps and cut churn (start now)

Phase 1 adds jobs to the workflows that exist today. Phase 3 moves them into
`PR Gate`.

| # | Change | Gap | Effort |
| --- | --- | --- | --- |
| 1.1 | Use one change detector, `.github/actions/detect-changes`, that fails open. Drop `tj-actions/changed-files`. Add the missing inputs: `.golangci.yml`, `.codecov.yml`, `.github/actions/**`, `scripts/**`, `mise.toml`, and `.github/styles/**`. Fix the stale `docs/snippets` path. | G6, G7, G8 | S |
| 1.2 | Add `go test -race` on Linux stable for both modules. Triage any races before the job blocks merges. | G9 | S |
| 1.3 | Lint and run `govulncheck` on both modules. Lint the `integration`, `e2e`, and `snippets` build tags. | G10, G11 | S |
| 1.4 | Add actionlint and shellcheck to CI. | G15 | S |
| 1.5 | Add `actions/dependency-review-action` for changes to `go.mod`, `go.sum`, or `website/package-lock.json`. | G13 | S |
| 1.6 | Add the API-compatibility job from ADR 028: blocking on `release/*`, advisory on `main`. | G12 | M |
| 1.7 | Run Vale on the whole site when the Vale config, styles, or version change. Build PR docs with the same `archived-versions` input as the stable deploy. | G17, G18 | S |
| 1.8 | Add the tool-pin drift check from ADR 026. Use version-file inputs where an action supports a file that `mise` also reads. | ADR 026 | S |
| 1.9 | Make fork PRs pass: skip sticky-comment steps on forks, and make the Codecov upload non-fatal. Verify the SARIF upload on a fork PR. | G20, G21 | S |
| 1.10 | Replace Codecov as a gate with an in-repo patch-coverage floor in the test job. Keep Codecov as an informational report. | G16 | M |
| 1.11 | Remove `build-go` where `test-go` covers the same OS and Go pair. Switch golangci-lint to `install-mode: binary`. | G26, G27 | S |
| 1.12 | Make `govulncheck` failures explain that a new advisory, not the PR, caused them. | G22 | S |
| 1.13 | Add a weekly flakiness report from the Actions API. If macOS cancel rates stay high, move macOS to push and scheduled runs only. | G19 | M |
| 1.14 | Add header comments to every workflow. Fold single-caller `reusable-*.yaml` files into their callers. | Layout | M |
| 1.15 | Add `gofmt -s` on staged Go files to the pre-commit hook. | Hook | S |
| 1.16 | Add the forward-port assessment rule for PRs into `release/*` to `branch-policy.yml`. | ADR 023 | S |
| 1.17 | Register ADR 024 exception 1 (`SN_*`) with an owner and an expiry date. | ADR 024 | S |

Spikes in phase 1:

- The e2e tests over OIDC (ADR 014).
- Jira releases from GitHub for Jira deployment events.
- A `workflow_dispatch` run of the gate on a scratch release-please branch.

### Phase 2: Jira cutover (on its own schedule)

Spec 009 and ADR 022 own this phase. This spec depends on only three of its
outcomes:

- The issue workflows retire: `issue-similarity-check`, `issue-status-sync`,
  `stale-issues`, and their reusable workflows.
- The `PR Policy` work-item rule switches from `#N` to an `SNSFG1` key. Until
  cutover, the `#N` rule and the `no-issue-required` label stay.
- The Jira release webhook lands as ADR 024 exception 2, unless the
  deployment-events spike removes the need for it.

### Phase 3: one gate and one policy check (waits on phase 2)

| # | Change | Gap | Effort |
| --- | --- | --- | --- |
| 3.1 | Add `pr-gate.yml` with a plain top-level `PR Gate` job that passes only when every needed job passed or was skipped on purpose. It has no workflow-level `paths:` filter. Move the Phase 1 jobs, `license-check.yml`, and the docs check jobs into it. `ci.yml` keeps only push and scheduled runs. Add `merge_group` to the triggers. | G2, G6, G7 | M |
| 3.2 | Merge `pr.yml` and `branch-policy.yml` into `pr-policy.yml`, with one aggregate `PR Policy` job. Keep Dependabot auto-merge outside the aggregate. | G23 | S |
| 3.3 | Cherry-pick 3.1 and 3.2 to `release/v1.0` by hand (ADR 023). Wait until both checks report there. | G25 | S |
| 3.4 | Switch both rulesets to exactly `PR Gate`, `PR Policy`, and `CodeQL`, and remove the old contexts in the same change. Rehearse with a throwaway PR on each branch, because a wrong context blocks every merge (#917). | G2 | S |
| 3.5 | Add a tag ruleset that blocks hand-pushed `v*` tags. Then remove `release-verify.yml` and its call in `stable-release.yml`. | ADR 027 | S |

### Items dropped from the DevOps plan

| DevOps plan item | Why it's dropped |
| --- | --- |
| 0.2: GitHub App token for release-please, backport, and stamp | Replaced by the built-in approval (ADR 024). The backport workflow is gone (#982). |
| 1.3 (part): branch-name check in `PR Policy` | Removed in #981. |
| 1.5 and G25: backport the gate through the backport label | The label flow is gone. Phase 3.3 cherry-picks by hand. |
| G4 (part): `weekly-release.yml` | Removed in #976. |
| G5 (part): backport PRs run no CI | The backport workflow is gone (#982). |
| 3.7: run CI tools through `jdx/mise-action` | Rejected in favor of native actions with a drift check (ADR 026). |
| 4.2: merge queue | Not available on a personal-account repo. `merge_group` stays in 3.1 at no cost. |
| Forward-port tracker | Removed in #982. ADR 023 replaced it. |

## Alternatives considered

- **Keep adding workflows as needs arise.** This is the status quo. Rejected,
  because it produced the churn this spec addresses.
- **Require each check by job name.** Rejected: path filters leave skipped
  jobs pending, and every matrix change forces a ruleset change. One aggregate
  `PR Gate` check avoids both.
- **A gate job that polls the Checks API for other workflows.** Rejected: it
  depends on timing, uses runner minutes while it waits, and can't tell a
  job that hasn't started from one that never starts.
- **A GitHub App key or octo-sts for bot PRs.** Rejected in ADR 024.
- **CI runs `mise` tasks.** Rejected in ADR 026.
- **Remove PR doc previews.** Considered to shrink the PR-time write surface.
  Rejected: the maintainer uses the previews.
- **Remove `labeler.yml`.** Considered, because release-please reads commits,
  not labels. Rejected by the maintainer.
- **Drop maintenance releases.** Considered as a way to cut ADR 011's surface.
  Rejected: ADR 023 keeps manual backports.

## Open questions

1. **Reason for dropping signing and SBOM (for review).** ADR 027 uses this
   wording, which the maintainer hasn't reviewed yet: "The upkeep cost more
   than the value, for a library that people get through the Go module proxy,
   where `sum.golang.org` already gives checksum integrity, and that has no
   binaries to sign."
2. **Node version file.** `setup-node` can read a `.node-version` or `.nvmrc`
   file. Does the pinned `mise` version read either file by default, or does
   it need configuration? If neither works cleanly, Node stays in the drift
   check.
3. **`govulncheck` and `tparse` in CI.** Pin them in CI and add them to the
   drift check, or keep `latest` in CI on purpose?
4. **Name of the Jira webhook secret.** ADR 022 leaves it open. ADR 024's
   `Verify` command needs the name before the release step lands.
5. **GitGuardian as a required check.** Should `GitGuardian Security Checks`
   become a fourth required check? It reports on every PR, but it's a
   third-party app outside the gate.
6. **Coverage floor.** Which patch-coverage floor does phase 1.10 start
   with? Pick it from current data, so the gate doesn't fail on day one.
7. **`release/v1.0` end of life.** When does v1 reach end of life, so that
   the branch can be deleted under ADR 011 rule 7?
