---
title: 'ADR 024: CI stores no long-lived secrets, except listed exceptions that expire'
description: >-
  CI authenticates with GITHUB_TOKEN, OIDC, or GitHub's built-in approval for
  bot PRs. A stored secret needs an entry in this record's exceptions list,
  with an owner and an expiry date.
---

# ADR 024: CI stores no long-lived secrets, except listed exceptions that expire

## Status

Proposed (2026-11-21). The plan is in
`specs/010-holistic-delivery-pipeline/spec.md`.

## Context

A stored secret in GitHub Actions is a standing risk. Anyone who can change a
workflow that reads the secret can leak it, and the secret stays valid until
someone rotates it. The repo has already paid for this:

- The hourly project-sync job needed a classic personal access token (PAT),
  because fine-grained tokens don't support Projects v2. #984 removed the job.
- After #984, these secrets are left with no workflow that uses them:
  `SYNC_PROJECT_SYNC_TOKEN_PROVIDER_PEM` (a GitHub App private key) and its
  `_APP_ID` variable, `TOKEN` (from 2023), and `CODECOV_TOKEN`. Coverage
  upload already uses OIDC.

One real need remains for an elevated identity. release-please opens its
release PR with `GITHUB_TOKEN`, and `stamp-deprecations.yml` pushes a commit
to that PR with the same token. GitHub doesn't start ordinary workflow runs
from events that `GITHUB_TOKEN` causes. The
[`GITHUB_TOKEN` concepts page](https://docs.github.com/en/actions/concepts/security/github_token)
lists two exceptions:

- `workflow_dispatch` and `repository_dispatch` events always create runs.
- When a workflow creates or updates a PR with `GITHUB_TOKEN`, the
  `opened`, `synchronize`, and `reopened` events create runs in an
  approval-required state. A user with write access starts them by selecting
  **Approve workflows to run** in the merge box.

Alternatives considered for bot PRs:

1. **A GitHub App installation token, from an App private key stored as a
   secret.** The installation token lasts an hour, but the private key is a
   long-lived secret. Rejected.
2. **octo-sts ([octo-sts/app](https://github.com/octo-sts/app)).** A GitHub
   App that exchanges an Actions OIDC token for a short-lived, scoped GitHub
   token, based on a trust policy in `.github/chainguard/<name>.sts.yaml`.
   This repo would store no secret, but the trust moves. The hosted service
   at `octo-sts.dev` would hold write access to this repo, and its project
   page doesn't name who runs it. Self-hosting needs a separate App key in a
   cloud key service, plus a service to run. Rejected.
3. **A PAT.** It's long-lived, tied to a person, and broadly scoped.
   Rejected.
4. **Post success statuses for the required checks without running them.**
   This fakes the gate. Rejected.
5. **GitHub's built-in approval for bot-PR runs** (chosen). It needs no new
   identity and no secret. The cost is one click on the release PR, which the
   maintainer already opens to merge.

The research for this record found no first-party GitHub feature that
exchanges an Actions OIDC token for a GitHub API token. That gap is the
reason octo-sts exists.

## Decision

1. **CI authenticates without stored secrets.** Workflows use
   `GITHUB_TOKEN` with the least `permissions:` they need, OIDC for services
   that accept it (for example, Codecov), or GitHub's built-in approval for
   bot-PR runs.
2. **A stored secret needs an exception.** Each exception is listed below with
   its owner, its expiry date, and how it could be removed. When an exception
   expires, the owner renews it with a new date or removes the secret.
3. **Bot PRs use the built-in approval.** The maintainer approves the runs on
   the release-please PR and on the stamp commit, then merges. A later spike
   tests a fully automatic path: the release workflow starts `PR Gate` with
   `workflow_dispatch` on the release-please branch. Adopt it only after a
   scratch PR shows that the dispatched run satisfies the required check.
4. **Unused secrets are deleted.** Delete
   `SYNC_PROJECT_SYNC_TOKEN_PROVIDER_PEM`, its `_APP_ID` variable, `TOKEN`,
   and `CODECOV_TOKEN` after a final usage check, and uninstall the
   project-sync App.

### Exceptions

| # | Secret | Used by | Owner | Expires | Way out |
| --- | --- | --- | --- | --- | --- |
| 1 | `SN_INSTANCE`, `SN_USERNAME`, `SN_PASSWORD`, `SN_CLIENT_ID`, and `SN_CLIENT_SECRET` in the `e2e-tests` environment | `e2e-nightly.yml` | Maintainer (`michaeldcanady`) | Nov 21, 2027 | A spike tests whether the e2e instance trusts the GitHub Actions OIDC issuer, using the SDK's external-OIDC support ([ADR 014](014-external-oidc-verbatim-token.md)). |
| 2 | The Jira Automation webhook secret (name not chosen yet) | The release step in ADR 022 (proposed, in progress), decision 10 | Maintainer (`michaeldcanady`) | 12 months after the secret is created | A spike tests whether a Jira rule can release a Fix version from a GitHub for Jira deployment event. |

Exception 2 takes effect only when the webhook step lands.

## Consequences

- **Pros:**
  - No PAT and no App key sits in the repo.
  - A leaked workflow can expose only the listed exceptions, and each one has
    an owner and a date.
  - The rule gives every new integration the same first question: can it run
    on OIDC?
- **Cons:**
  - Each release needs one extra click to approve the bot-PR runs.
  - The stamp commit needs a second approval when it lands after the first.
  - Integrations that support only static tokens need an exception, or they
    can't run in CI.
- **Rule for future CI work:** don't add a `secrets.*` reference without
  adding its row to the exceptions table in the same PR. Don't add a GitHub
  App key, a PAT, or an octo-sts trust policy without a new ADR that
  supersedes decision 3.
