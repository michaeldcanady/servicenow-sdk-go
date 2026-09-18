---
title: CI/CD architecture blueprint
description: Target architecture and phased migration plan for the GitHub Actions fleet — domain pods, canonical quality gate, canonical release pipeline, and event-model conventions.
---

# Technical Blueprint: CI/CD Architecture (`github/workflows`)

This blueprint records the audit and target architecture for the repository's
GitHub Actions fleet. It is the reference for how CI/CD workflows are
structured, named, triggered, and enforced. All migrations described here are
`chore(ci)` changes that land by PR in the sequence below.

## 1. Current state

The fleet contains 35 workflows in `.github/workflows/`. They cluster into
five concerns:

| Domain | Workflows | Trigger model |
| :--- | :--- | :--- |
| Quality gates | `quality-ci.yml`, `quality-weekly.yml`, `quality-gate.yml`, `reusable-check-go-deps.yml`, `reusable-build-go.yml`, `reusable-lint-go.yml`, `reusable-test-go.yml`, `reusable-govulncheck.yml`, `quality-license.yml`, `quality-license-backfill.yml`, `pr-branch.yml`, `pr-linked-issue.yml`, `pr-title.yml`, `pr-dependabot.yml`, `pr-labeler.yml`, `pr-opencode-review.yml`, `docs-preview.yml` | `pull_request`, `push`, `schedule` |
| Release path | `rel-stable.yml`, `rel-weekly.yml`, `rel-pipeline.yml`, `rel-verify.yml`, `docs-version.yml`, `rel-backport.yml`, `rel-forward-port.yml`, `rel-stamp-deprecations.yml`, `rel-maintenance-label.yml` | `push`, `workflow_call`, `workflow_dispatch`, `pull_request_target` |
| Scheduled bots | `e2e-nightly.yml`, `issues-stale.yml`, `issues-project-sync.yml`, `security-scorecard.yml` | `schedule`, `workflow_dispatch` |
| Issue automation | `issues-similarity.yml`, `issues-status-sync.yml` | `issues`, `pull_request_target` |
| Security | `security-codeql.yml`, `security-zizmor.yml` | `pull_request`, `schedule` |

Shared building-block actions live in the org repository
`NerdIT-Tech/.github`, pinned to tagged commits. The quality gate and the
release pipeline are reusable workflows owned locally (`quality-gate.yml`,
`reusable-*.yml`, `rel-pipeline.yml`). The release pipeline
runs in-band (software bill of materials, SBOM, signing, and verification
invoked by the releasing workflow, not by a `release` event) because GitHub
does not start workflow runs from events that `GITHUB_TOKEN` created.

Deliberate strengths to preserve:

- `permissions: {}` workflow defaults with least-privilege job scoping.
- SHA-pinned actions with tag comments everywhere.
- Metadata-only `pull_request_target` use with documented zizmor ignores.
- Environment-gated secrets for live instance e2e tests.
- ADR 011-aligned backport, forward-port, stamp, and maintenance-label flows.

## 2. Audit findings

### Correctness

| Severity | Finding | Location |
| :--- | :--- | :--- |
| High | `workflow_dispatch` on `ci.yml` runs no jobs: the `changes` job skips dispatch, so every gated job inherits an empty output. | `ci.yml` |
| High | Trigger paths and the internal change filter disagree: `.golangci.yml` appears only in the PR trigger, so a `.golangci.yml`-only PR runs zero checks and the lint gate never re-runs under new config. | `ci.yml` |
| High | Dead paths from the org migration: trigger and filter still watch `.github/workflows/reusable-*.yaml` and `.github/actions/*` that were deleted when shared logic moved to `NerdIT-Tech/.github`. Cross-repo changes cannot trigger this repo, so these entries can never fire. | `ci.yml` |
| High | `release-verify.yml` re-implements the quality gate inline with its own pinned tool versions, so the tag gate and the PR gate are two divergent copies. | `release-verify.yml` |

### Consistency

| Severity | Finding | Location |
| :--- | :--- | :--- |
| Medium | `stable-release.yml` and `weekly-release.yml` duplicate the `release-please` → SBOM → sign → verify pipeline. | both files |
| Medium | The two docs deploy workflows write `gh-pages` under different concurrency groups, so a stable deploy and a preview deploy can interleave on the same branch. | `docs-stable.yml`, `docs-preview.yml` |
| Medium | The `docs-preview.yml` trigger lists `.github/actions/check-snippet-regions/**` twice. | `docs-preview.yml` |
| Medium | Three copies of idempotent label provisioning and two copies of the secret-presence check repeat tactical bash. | `stale-issues.yml`, `maintenance-label.yml`, `forward-port-tracker.yml`, `e2e-nightly.yml`, `sync-project-status.yml` |
| Medium | Two change-detection mechanisms coexist: `tj-actions/changed-files` in `ci.yml` and the org `detect-changes` action elsewhere. | `ci.yml`, `docs-preview.yml`, `weekly-release.yml` |
| Medium | Concurrency group formulas and `cancel-in-progress` values vary per file without a written convention. | all files |
| Medium | Workflow and file names are inconsistent: generic `pr.yml` and `ci.yml`, Title Case `name:` values, and emoji in the `zizmor.yml` name. Job and workflow names feed required-check names in branch-protection rulesets, so renames have merge-blocking consequences. | all files |
| Low | `codeql.yml` and `zizmor.yml` run on `push: main` plus PR plus schedule; the push leg rescans a ref the PR leg just scanned. | both files |

### Hygiene

| Severity | Finding | Location |
| :--- | :--- | :--- |
| Low | `sync-project-status.yml` polls hourly with a classic PAT. Projects v2 has no webhook, so polling is forced; document the least-privilege and rotation expectation. | `sync-project-status.yml` |
| Low | The local `ensure-backport-label` job lacks `timeout-minutes`. | `maintenance-label.yml` |
| Low | No workflow-schema lint at PR time. zizmor audits security; `actionlint` catches YAML and event-model errors. | quality gate |

## 3. Target architecture

### 3.1. Naming

GitHub does not support subdirectories under `.github/workflows/` for
trigger-scanned files, so organization uses filename prefixes. Use `pr-`,
`quality-`, `security-`, `docs-`, `rel-`, `issues-`, and `e2e-` prefixes. Use
`rel-` rather than `release-` so the file namespace never reads as the ADR 011
`release/vX.Y` branch namespace. Write `name:` values in sentence case.

| Pod | Example files |
| :--- | :--- |
| `pr-` | `pr-title`, `pr-branch`, `pr-linked-issue`, `pr-dependabot`, `pr-labeler`, `pr-opencode-review` |
| `quality-` | `quality-ci`, `quality-weekly`, `quality-license` |
| `security-` | `security-codeql`, `security-zizmor`, `security-scorecard` |
| `docs-` | `docs-preview`, `docs-stable`, `docs-version` |
| `rel-` | `rel-stable`, `rel-weekly`, `rel-verify`, `rel-backport`, `rel-forward-port`, `rel-sbom`, `rel-sign`, `rel-stamp-deprecations`, `rel-maintenance-label` |
| `issues-` | `issues-similarity`, `issues-status-sync`, `issues-stale`, `issues-project-sync` |
| `e2e-` | `e2e-nightly` |

### 3.2. Event model

Each pod reacts to the minimal event set that expresses its intent:

| Pod | Events |
| :--- | :--- |
| `pr-*` | `pull_request`, with `pull_request_target` only for metadata-only issue and PR writes |
| `quality-*` | `pull_request` and `push` with path filters, plus `schedule` for the weekly matrix |
| `security-*` | `pull_request` and `schedule`; `branch_protection_rule` for scorecard; no `push` leg |
| `docs-*` | `pull_request` and `push` scoped to `website/**`, plus `push` on `v*` tags |
| `rel-*` | `push` on `main` and `release/v*`, `schedule` for weekly, `workflow_dispatch`; `pull_request_target` on `closed` for backport and forward-port |
| `issues-*` | `issues`, `pull_request_target`, `schedule` |
| `e2e-*` | `schedule` and `workflow_dispatch` |

### 3.3. Rules

1. **Filter at the edge.** Put path intent in trigger `paths`. A job exists to
   run, not to re-decide. Only events without path metadata (`schedule`,
   `workflow_dispatch`) get a single gate job.
2. **One canonical quality gate.** A `quality-gate` workflow owned locally
   encodes the Go build, lint, test, vulnerability, and module-check matrices.
   The PR gate, the weekly matrix, and release verification all call it, so the
   tagged ref and the merge commit pass literally the same gate.
3. **One canonical release pipeline.** A `rel-pipeline` workflow composes SBOM,
   signing, and verification. Both the stable and weekly release entry points
   call it.
4. **Serialize shared physical resources globally.** Deploys to `gh-pages`
   share one named concurrency group across workflows. Tag writes share a
   `release-publish` group. Last-write races disappear by construction.
5. **Cron is a first-class event.** Scheduled pods stay scheduled. Keep the
   cron budget in one table in the docs so new slots stay provably
   collision-free.
6. **Machine-enforced conventions.** `actionlint` runs in the quality gate next
   to zizmor. Both are required checks on PRs.
7. **Concurrency is formulaic.** Match the group and `cancel-in-progress` to the
   event profile. Multi-event check pods serialize per PR, per branch, and per
   run: group `${{ github.workflow }}-${{ github.head_ref || github.ref ||
   github.run_id }}` with `cancel-in-progress: true` — keying by `github.ref`
   alone collapses every PR run onto the base branch and lets concurrent PRs
   cancel each other. PR-only pods key by `github.event.pull_request.number`.
   Push and release pods key by `ref` with `cancel-in-progress: false`. Mounted
   pods keep a static per-entity key (for example `stale-issues`,
   `issue-similarity-${{ github.event.issue.number }}`). Shared physical
   resources get named groups (`gh-pages-deploy` for gh-pages writers,
   `release-publish` for tag and release writers), applied to the job that
   writes and never to a workflow that delegates the write to a reusable, so a
   caller and its callee cannot deadlock on the same group.

## 4. Migration plan

The migration proceeds in four phases. Each phase is independently shippable
and observable.

### Phase 0 — Stabilize

Fix correctness findings without renames or new abstractions:

- Make `workflow_dispatch` on `quality-ci.yml` run the full matrix.
- Unify trigger paths, the change filter, and `.golangci.yml` coverage; drop
  the dead shared-org paths.
- Deduplicate the `docs-preview.yml` trigger glob.
- Add missing `timeout-minutes`.

Exit criteria: zizmor and `actionlint` clean, CI green.

### Phase 1 — One gate, one pipeline

Deduplicate behavior without visible check-run changes:

- Add the gate: a local `quality-gate.yml` orchestrator over local
  `reusable-*.yml` legs, each accepting a `checkout-ref` input; point
  `quality-ci.yml` and `rel-verify.yml` at it while keeping the
  `Verify Tagged Ref` job name.
- Add the `rel-pipeline` workflow; rewrite both release orchestrations over it
  and drop the now-superseded `sbom.yml` and `sign-release.yml`.
- Add the shared `gh-pages` concurrency group to both docs deploy workflows.
- Replace `tj-actions/changed-files` with the org `detect-changes` action.

Exit criteria: one source of truth for the gate and the pipeline, no required
check renamed.

### Phase 2 — Event model and rename wave

Apply the pod prefixes and sentence-case names. Drop the redundant `push: main`
legs on `security-codeql.yml` and `security-zizmor.yml`. Replace local `changes`
jobs with trigger-level paths plus a dedicated weekly matrix workflow. Codify
concurrency formulas per event class. Move the repeated label-provisioning and
secret-presence code into shared composite actions.

This phase changes one thing the branch-protection rulesets can see: the
reported name of the zizmor job. For reusable workflows, GitHub reports a
check as `caller job / callee job`, so dropping the emoji from the
`security-zizmor.yml` job name changes the reported contexts like this:

| Required check | Reported today | Reported after the wave |
| :--- | :--- | :--- |
| `security-zizmor.yml` zizmor job (blocking) | `Run zizmor 🌈 / Run zizmor -- blocking` | `Run zizmor / Run zizmor -- blocking` |
| `security-zizmor.yml` zizmor job (SARIF) | `Run zizmor 🌈 / Run zizmor -- SARIF upload` | `Run zizmor / Run zizmor -- SARIF upload` |

The live `main` and `release/v*` rulesets today require the context `Run
zizmor 🌈 — auditor · all inputs · fail on any finding`, which no check has
reported since the 2026-09 reusable-workflow move; merges stay unblocked
because the repo is administered under the `admin(always)` bypass. Because of
that drift, the wave itself needs no coordinated ruleset edit to keep merges
working.

Decision: fix the ruleset context to the real post-wave name. Update both
rulesets to require `Run zizmor / Run zizmor -- blocking` in the same change
as the wave, which makes zizmor gate merges again instead of leaving a phantom
required check in place.

Items already landed ahead of the wave (they rename nothing): the redundant
`push: main` legs are gone from `security-codeql.yml` and
`security-zizmor.yml`; the concurrency formulas from rule 7 are applied; label
provisioning is shared through the local `.github/actions/ensure-label`
composite; secret-presence checks already delegate to the org `check-secret`
composite.

The wave itself is implemented on the `draft/gate-orchestrator` branch: files
now use the pod prefixes (`quality-ci.yml`, `quality-weekly.yml`,
`rel-stable.yml`, `rel-verify.yml`, `security-codeql.yml`, `issues-stale.yml`,
and so on); `branch-policy.yml` split into `pr-branch.yml` and
`pr-linked-issue.yml`, `pr.yml` into `pr-title.yml` and `pr-dependabot.yml`;
`name:` values are sentence case; the `changes` job is gone from `quality-ci`
(two-lived changed-files detection replaced by trigger-level paths), and its
`schedule` leg moved to a dedicated `quality-weekly.yml` matrix workflow. The
other four required contexts (`Check Branch Name`, `Check Linked Issue`,
`Validate PR Title`, `CodeQL`) survive unchanged; the `pr-branch`,
`pr-linked-issue`, and `pr-title` splits preserve the job names byte for byte
so GitHub keeps reporting the same checks from the renamed files.

Exit criteria: merges work, all checks map cleanly, dispatch and schedule run
real pipelines.

### Phase 3 — Enforce and document

Add `actionlint` to the quality gate. Write the CI conventions reference
(naming, event model, concurrency, secrets pattern, cron budget) and file the
design decisions that cross the ADR bar with the product-manager agent. Update
`CLAUDE.md` if any relevant command changes.

Exit criteria: a new workflow can be added from the conventions doc alone.

## 5. Risks

- **Required-check renames** block merges if the ruleset is not updated in
  lockstep. That is why Phase 2 is isolated as its own phase.
- **`NerdIT-Tech/.github` is org-owned.** Prefer additive reusable inputs that
  preserve old behavior, and pin consumer upgrades to tags so each repo
  migrates at its own pace.
- **Reuse chains harden, they do not weaken, guarantees.** Release
  verification must run the exact PR gate, which is the point of Phase 1.