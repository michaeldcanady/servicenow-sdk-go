---
title: Roadmap from v2 to v3
description: >-
  Where the SDK and its repository are headed, the decisions that shape the
  route, and the small work units that get from today's v2 to v3.0.0.
---

# Roadmap from v2 to v3

This page describes where the SDK and its repository are headed and the order
of work to get there. It also defines the scope of v3.0.0.

The roadmap breaks the work into **units**. A unit is one small pull request
(one concern, ideally under about 300 changed lines) or one non-code action,
such as a settings change. Each unit becomes its own issue when work on it
starts. Issue and pull request numbers on this page refer to the
[servicenow-sdk-go issue tracker](https://github.com/michaeldcanady/servicenow-sdk-go/issues).

## Where the repository is today

An audit in November 2026 found that the code is healthy and the process
around it isn't:

- The build, `golangci-lint`, unit tests, and mocked integration tests all
  pass. Most packages have 85–99% statement coverage. The exception is
  `caseapi`, at 60%.
- Since August 2026, the history shows 169 `chore` commits, 29 `fix` commits,
  and 4 `feat` commits. Most of that effort went into CI.
- Merges need an admin bypass because two required status check names no
  longer match the names that the workflows report (#917).
- Documentation Stable, Sync Project Status, and the Dependabot Go updates
  fail on `main`.
- The domain-first design in ADRs 017–021 is still `Proposed`, and the pull
  requests that carry the planned credentials and logging work (#872 and #690)
  have stalled.
- Agent configuration is spread across `.agents/`, `.claude/`, `.opencode/`,
  `docs/agents/`, `docs/rules/`, and two near-identical instruction files.

## Where the repository is going

1. **v2.x is stable and acts as the bridge.** Each new surface ships next to
   the surface it replaces, and the old surface is marked deprecated.
2. **v3.0.0 is the domain-first SDK.** The new surfaces are the only surfaces,
   the deprecated code is gone, and the module path ends in `/v3`.
3. **CI stays green by default.** Every workflow has a stated purpose and a
   recent passing run. Required check names are treated as a contract.
4. **Agent configuration is vendor-neutral and has one source.**
5. **Documentation follows one style, and Vale enforces it.**

## Decisions that shape the route

| Topic | Decision |
|---|---|
| Deprecation | New and old surfaces coexist in v2.x, and the old surface carries a `Deprecated:` notice. v3.0.0 removes everything that's deprecated. |
| `Table2` | `Table2()` ships in v2.x as the bridge and is renamed to `Table()` in v3.0.0. Both `Table2` and the legacy `Table(name string)` go away in v3. See [ADR 018](adrs/018-table-collection-builder.md). |
| Module path | The module stays at `github.com/michaeldcanady/servicenow-sdk-go`. Only the major-version suffix changes. See [ADR 010](adrs/010-no-nerdit-tech-migration.md). |
| CI | First, contract CI to what one maintainer can keep green. Then expand it onto the shared `NerdIT-Tech/.github` workflows. |
| Preview releases | Weekly preview releases are dropped. Only stable releases ship. |
| Maintenance-branch tooling | The backport, forward-port tracker, and maintenance-label workflows are removed until v3 creates the first `release/v2.x` branch, and then restored. The policy in [ADR 011](adrs/011-release-branches-and-cross-major-flow.md) stays in force. |
| AI review | OpenCode Review stays in CI. |
| Issue automation | Issue status sync, project status sync, and the GitHub Projects board are dropped. The labeler, the stale bot, and the issue similarity check stay. |
| Agent configuration | `AGENTS.md` and `.agents/` are the only sources. `CLAUDE.md`, `.claude/`, and `.opencode/` are generated shims. |
| Agent skills | The repository uses only skills it authors. |
| Prose style | The Google developer documentation style guide, enforced by Vale. |
| Task runner | mise replaces `just`. It pins the tool versions that CI uses, and every CI check with a local equivalent has a mise task. |

## Guardrails

These rules apply to every unit:

- **Required check names are a contract.** If a pull request renames a job
  or moves it into a reusable workflow, update
  `.github/policies/servicenow-sdk-go-branch-protection.yml` in the same
  pull request.
- **Consult the ADRs first.** If a change contradicts an ADR, amend the ADR
  in the same pull request.
- **Preserve error identity in v2.** When a package-local sentinel is
  replaced by a shared one, keep the old name as a deprecated alias so that
  `errors.Is` still matches. New error categories keep
  `errors.As` with `*core.ServiceNowError` working.
- **Hold the v3 release pull request.** After the `/v3` module-path change
  lands, release-please proposes 3.0.0 right away. Don't merge that pull
  request until every v3 unit is done.

## Critical path

The credentials rebuild is the longest chain of dependent units. Its design
is already accepted in [ADR 015](adrs/015-credentials-provider-surface.md),
so it starts as soon as Phase 0 is done and runs in parallel with Phases 1
and 2:

> P0-15 → P3-24 through P3-39 → P3-59 and P3-60 → P3-65 (v2.2.0) → P4-01 →
> P4-02 → P4-07 → P4-10 through P4-13 → P4-18 → P4-21 (v3.0.0)

Step 0 comes first. It's short, and the tasks it adds are how you verify every
later unit before you push.

These tracks run in parallel with the credentials chain:

- The Phase 0 tracks.
- The Phase 1 removals.
- The Phase 2 chain P2-01 → P2-04 → P2-05 → P3-01, which gates the table
  work.
- Logging, `caseapi` coverage, and the spike on the sentinel for unexpected
  response types.

Phase 5 is off the critical path.

## Work units

Size key:

- **XS:** a few lines.
- **S:** under about 300 changed lines.
- **M:** the largest allowed size.

Units marked **Action** are non-code steps. A unit with "new" in the Issue
column gets a new issue when work starts.

### Step 0: Local tooling with mise

Goal: [mise](https://mise.jdx.dev/) replaces `just` as the task runner and
pins every tool that CI uses. You can run each CI check that has a local
equivalent before you push, with the same tool versions that CI uses. Every
later phase uses these tasks to verify its units.

Task names group by area, for example `go:test:integration` and `docs:lint`. A
short task becomes an entry in `mise.toml`. A task that needs a script becomes
a file task under `.mise/tasks/`.

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| T-01 | `chore(tooling)`: add `mise.toml` that pins Go, Node 20, golangci-lint v2.14.0, Vale 3.18.0, zizmor 1.29.0, actionlint, shellcheck, govulncheck, and tparse at the versions CI uses | XS | — | new |
| T-02 | `chore(tooling)`: port every `just` recipe to a mise task with the same behavior: `go:build`, `go:test`, `go:fmt`, `go:lint`, `docs:setup`, `docs:serve`, `docs:build`, `docs:clean`, `docs:snippets`, `license:add`, and `license:check` | S | T-01 | new |
| T-03 | `chore(tooling)`: add `go:test:integration` and `go:test:e2e` tasks for the tagged suites, and `go:test:coverage` and `go:test:report` tasks that wrap `scripts/test.sh` | XS | T-02 | new |
| T-04 | `chore(tooling)`: add a `go:test:oldstable` task that runs the unit tests on the `oldstable` Go version, matching the CI matrix | XS | T-02 | new |
| T-05 | `chore(tooling)`: add a `go:deps:check` task that checks `go mod tidy` drift and runs `go mod verify` in both Go modules, matching the `check-go-deps` action | S | T-02 | new |
| T-06 | `chore(tooling)`: add a `go:vuln` task that runs govulncheck on both Go modules | XS | T-02 | new |
| T-07 | `chore(tooling)`: add a `docs:lint` task that runs Vale with the repository configuration on the same paths as Documentation Stable | XS | T-02 | new |
| T-08 | `chore(tooling)`: add a `ci:lint` task that runs zizmor and actionlint on `.github/`, and shellcheck on `scripts/` | XS | T-02 | new |
| T-09 | `chore(tooling)`: add `pr:branch` and `pr:title` tasks that apply the branch-name rule and the Conventional Commits title rule that the PR checks enforce | S | T-02 | new |
| T-10 | `chore(tooling)`: add `check` (everything a pull request runs) and `check:all` (adds the merge-only checks, such as the full-site Vale run and the docs build), and list the checks that only CI can run | S | T-03 through T-09 | new |
| T-11 | `chore(devcontainer)`: install mise in the devcontainer, run `mise install` after creation, and stop installing `just` | XS | T-01 | new |
| T-12 | `docs(contributing)`: replace `just` with mise in the setup, first pull request, and conventions pages, and in `AGENTS.md` and `CLAUDE.md` | S | T-10 | new |
| T-13 | `chore(tooling)`: delete the `justfile` | XS | T-11, T-12 | new |
| T-14 | `chore(ci)`: install CI tools from `mise.toml` so that local and CI versions can't drift; required check names stay unchanged | S | T-01 | new |

Some checks have no local equivalent: CodeQL, OpenSSF Scorecard, OpenCode
Review, GitGuardian, and the linked-issue check. They stay CI-only, and T-10
documents them.

### Phase 0: Foundation

Goal: `main` is green, the shipped defects are fixed, and the backlog reflects
this roadmap.

P0-06 is on hold. GitHub doesn't apply the policy file in
`.github/policies/` on its own, so branch protection moves to Terraform in
`NerdIT-Tech/infra-terraform`. Until that migration lands, `main` and
`release/v*` have no rulesets, and no unit waits for P0-06.

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P0-01 | `chore(agents)`: commit the curated skill set, and keep `design-decisions` and `google-tech-writing` until P2-10 and P2-11 replace them | S | — | new |
| P0-02 | Action: review the uncommitted edits to the architect and product-manager agent definitions | XS | — | new |
| P0-03 | `chore(agents)`: track the `.claude/` symlink shims, and ignore local settings files | XS | P0-01 | new |
| P0-04 | `chore`: remove `specs/` after you file an issue for any spec that has no matching issue or ADR | XS | — | new |
| P0-05 | `chore(ci)`: align the required check names in both rulesets | XS | — | #917 |
| P0-06 | On hold: after the Terraform migration applies the rulesets, confirm that a pull request merges without a bypass | — | P0-05, Terraform migration | #917 |
| P0-07 | `chore(ci)`: rerun pull request workflows when the base branch changes | S | P0-05 | #580 |
| P0-08 | `chore(docs)`: switch Vale from the Microsoft package to the Google package | XS | — | new |
| P0-09 | `docs(contributing)`: fix the Vale errors under `contributing/` | S | P0-08, T-07 | new |
| P0-10 | `docs(apis)`: fix the Vale errors under `apis/` | S | P0-08, T-07 | new |
| P0-11 | `docs`: fix the Vale errors in the user guide and landing pages | S | P0-08, T-07 | new |
| P0-12 | Action: re-scope #881 to Go error strings | — | P0-08 | #881 |
| P0-13 | `chore(deps)`: keep both Go modules tidy when Dependabot updates a dependency | S | — | new |
| P0-14 | Action: recreate and merge the stalled Dependabot pull requests | — | P0-13 | #906, #907 |
| P0-15 | `fix(credentials)`: keep `client_secret` out of the authorization URL | S | — | #851 |
| P0-16 | `fix(caseapi)`: forward query parameters on the field values request | XS | — | #742 |
| P0-17 | Action: merge the release pull request for v2.1.1 | — | P0-15, P0-16 | #894 |
| P0-18 | Action: publish a GitHub security advisory for #851, with guidance on rotating secrets | — | P0-17 | #851 |
| P0-19 | Action: renew the OpenCode Review API key | — | — | #897 |
| P0-20 | Action: close milestone v2.0.0, and create milestones v2.2.0 and v3.0.0 | — | — | — |
| P0-21 | Action: assign issues to milestones as the issue map shows | — | P0-20 | — |
| P0-22 | Action: close #471 (duplicate of #642) and #500 (superseded by #832) | — | — | — |
| P0-23 | Action: close #560, or re-scope it to the gap that remains | — | — | #560 |
| P0-24 | Action: remove the `stale` label from planned issues | — | — | — |
| P0-25 | `chore(ci)`: exempt issues that have a milestone from the stale bot | XS | — | new |
| P0-26 | Action: close #812 and #813, and point to Phase 5 | — | — | — |
| P0-27 | Action: delete branches that are already merged into `main` | — | — | — |
| P0-28 | Action: review the unmerged branches in batches | M | P0-27 | — |
| P0-29 | Action: run the Phase 0 exit check | — | P0-11, P0-14, P0-17 | — |

### Phase 1: Contract CI

Goal: about 34 workflow files, down from 44, all green.

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P1-01 | `chore(ci)`: remove the weekly preview release workflow, its configuration, and every reference to it; amend ADRs 011 and 012 | S | — | new |
| P1-02 | `chore(ci)`: remove the backport, forward-port tracker, and maintenance-label workflows and their zizmor exceptions; add a note to ADR 011 | S | — | new |
| P1-03 | Action: file the v3.0.0 issue to restore the maintenance-branch tooling | — | P0-20, P1-02 | new |
| P1-04 | `chore(ci)`: remove issue status sync and its reusable workflow | XS | — | new |
| P1-05 | `chore(ci)`: remove project status sync and `project-label-sync.yml` | XS | — | new |
| P1-06 | `docs(contributing)`: remove the GitHub Projects steps from the triage docs, the issue templates, and the agent definitions | S | P1-04, P1-05 | new |
| P1-07 | `chore(ci)`: remove the one-time license backfill workflow | XS | — | new |
| P1-08 | `chore(ci)`: remove the reusable workflows that nothing calls | XS | — | new |
| P1-09 | Action: delete the secrets that only the removed workflows used | — | P1-01, P1-02, P1-04, P1-05, P1-07 | — |
| P1-10 | `docs(contributing)`: add a workflow inventory page | S | P1-01 through P1-08 | new |
| P1-11 | Action: run the Phase 1 exit check, and decide the status of ADR 013 | — | P1-01 through P1-10 | — |

### Phase 2: Vendor-neutral agent configuration

Goal: `AGENTS.md` and `.agents/` (with `agents/`, `skills/`, `docs/`, and
`rules/` subdirectories) are the only sources. `CLAUDE.md`, `.claude/`, and
`.opencode/` are generated shims, and the top-level `docs/` directory is
gone.

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P2-01 | `chore(agents)`: remove vendor names from `AGENTS.md` | XS | P0-01 | new |
| P2-02 | `chore(agents)`: reduce `CLAUDE.md` to an import of `AGENTS.md` | XS | P2-01 | new |
| P2-03 | `chore(agents)`: move `docs/agents/` to `.agents/docs/` | S | P2-01 | new |
| P2-04 | `chore(agents)`: move `docs/rules/` to `.agents/rules/`, and update every reference | S | P2-03 | new |
| P2-05 | `docs(adrs)`: reconcile the rule-file IDs with the ADR catalog | S | P2-04 | new |
| P2-06 | `chore(agents)`: limit agent frontmatter to `name`, `description`, and `tools` | XS | P0-02 | new |
| P2-07 | `chore(agents)`: add an `agents:sync` mise task that regenerates the shims | S | P0-03, P2-02, P2-03, T-02 | new |
| P2-08 | `chore(ci)`: fail CI when the shims drift, using a step in `ci.yml` | XS | P2-07 | new |
| P2-09 | `chore(agents)`: delete `skills-lock.json` | XS | P0-01 | new |
| P2-10 | `chore(agents)`: rewrite `google-tech-writing` as a repository-authored skill | S | P0-01 | new |
| P2-11 | `chore(agents)`: rewrite `design-decisions` so that it applies ADRs through `.agents/rules/` | S | P2-04, P2-05 | new |
| P2-12 | `chore(ci)`: load the OpenCode Review prompt from the reviewer agent definition on the base branch | S | P0-19, P2-01 | new |
| P2-13 | Action: run the Phase 2 exit check | — | P2-01 through P2-12 | — |

### Phase 3: v2.2, the bridges

Goal: every new surface ships next to the surface it replaces. Each unit adds
its `Deprecated:` notice in the same pull request, and all new code has
table-driven tests.

#### Design gate

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-01 | `docs(adrs)`: accept ADRs 017–021 | XS | P2-04, P2-05 | new |

#### Exported foundation

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-02 | `feat(core)`: export the backing-store accessor and mutator helpers | S | P3-01 | #883 |
| P3-03 | `feat(core)`: export the serialization helper generators | S | P3-01 | #883 |
| P3-04 | `docs`: add a snippet that shows a user-defined model | XS | P3-02, P3-03 | #883 |

#### Tables

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-05 | `feat(tableapi)`: add the `Table2().ByName()` builder, and deprecate `Table(name)` | S | P3-01 | #882 |
| P3-06 | `test(tableapi)`: add integration scenarios for `Table2` | S | P3-05 | new |
| P3-07 | `feat(tableapi)`: add the typed table shape and the Incident table | M | P3-05 | #885 |
| P3-08 | `feat(tableapi)`: add the Case typed table | S | P3-07 | #885 |
| P3-09 | `feat(tableapi)`: add typed query field descriptors | M | P3-07 | #885 |
| P3-10 | `feat(tableapi)`: support reads on user-defined typed tables | M | P3-02, P3-03, P3-09 | #886 |
| P3-11 | `feat(tableapi)`: support writes on user-defined typed tables | S | P3-10 | #886 |
| P3-12 | `test(tableapi)`: add integration scenarios for typed tables | S | P3-11 | new |

#### Attachments

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-13 | `feat(attachmentapi)`: add the `ForRecord` helper | S | P3-01 | #884 |

#### Errors

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-14 | `feat(core)`: add `ConflictError` for HTTP 409 | S | P3-01 | new |
| P3-15 | `feat(core)`: add `ValidationError` | S | P3-14 | new |
| P3-16 | `feat(core)`: attach the raw ServiceNow error details to every category | S | P3-15 | new |
| P3-17 | Spike: choose the sentinel for an unexpected response type | S | — | #585 |
| P3-18 | `feat(errors)`: add the unexpected-response-type sentinel | S | P3-17 | new |
| P3-19 to P3-22 | `refactor`: apply the standard error phrasing and shared sentinels, with aliases, one package group at a time: `errors/` and the root package, `core/` and `internal/`, `tableapi`, then the remaining API packages | S each | P0-12 (P3-20 also needs P3-18) | #881 |
| P3-23 | `refactor(credentials)`: apply the standard error phrasing | S | P3-39 | #881 |

#### Credentials

These units replace #872. Each one adds the matching `Deprecated:` notices.

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-24 | Action: map each commit in #872 to one of #834–#842 | — | P0-15 | — |
| P3-25 | Move the credential sentinels to the shared errors package, with aliases | S | P3-24 | #834 |
| P3-26 | Add the presented-token field, and make expiry nil-safe | S | P3-25 | #836 |
| P3-27 | Persist, reload, and revoke tokens through the token store | S | P3-26 | #835 |
| P3-28 | Coalesce concurrent token refreshes | S | P3-27 | #835 |
| P3-29 | `fix`: send no secret at the token endpoint for public clients | S | P3-28 | #837 |
| P3-30 | Use a separate PKCE verifier for each authorization | S | P3-29 | #837 |
| P3-31 | Add the shared provider options | S | P3-30 | #838 |
| P3-32 | Set the instance once on the client, and deprecate `WithInstance` and `WithURL` | S | P3-31 | #838 |
| P3-33 | Move ROPC onto the provider surface | S | P3-32 | #838 |
| P3-34 | Add the client-credentials provider | S | P3-33 | #839 |
| P3-35 | Add the JWT provider | S | P3-33 | #840 |
| P3-36 | Add the authorization-code provider | S–M | P3-33 | #841 |
| P3-37 | Align the Basic provider | XS | P3-34, P3-35, P3-36 | #842 |
| P3-38 | Align the Bearer provider | XS | P3-37 | #842 |
| P3-39 | Deprecate `BaseAccessTokenProvider` and `Authority` | S | P3-38 | #842 |
| P3-40 | Action: close #872 as superseded | — | P3-39 | — |

#### External OIDC

This work is additive, so it can ship in v2.2 or v2.3.

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-41 | `feat(credentials)`: resolve internal OIDC discovery documents | S | P3-33 | #844 |
| P3-42 | `feat(credentials)`: add the external OIDC provider | M | P3-36, P3-41 | #845 |
| P3-43 | Action: verify the provider against a live identity provider tenant | — | P3-42 | #829 |
| P3-44a to P3-44c | Action: verify the provider with ADFS, Auth0, and Google, one action each | — | P3-42 | #823 |

#### Logging

These units replace #690.

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-45 | Action: sort the files changed in #690 into core work, follow-up work, and changes to drop | — | — | — |
| P3-46 | `feat(core)`: add `LoggingMiddleware`, which does nothing by default | S | P3-45 | #499 |
| P3-47 | `feat(core)`: wire the `WithLogger` option into the default pipeline | S | P3-46 | #499 |
| P3-48 | `docs`: explain how to bring your own logger | XS | P3-47 | #499 |
| P3-49 | Action: close #690, and file its follow-up work | — | P3-47 | — |

#### Quality

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-50 to P3-54 | `test(caseapi)`: cover one builder family per unit, and reach at least 85% coverage | S each | — (the field values unit needs P0-16) | new |
| P3-55 | Action: file coverage issues for the next lowest-coverage packages | — | P3-54 | #636 |

#### Documentation

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-56 | Add glossary terms for the domain-first surface | XS | P3-05, P3-07, P3-13 | #887 |
| P3-57 | Add a page on domain-first conventions | S | P3-01 | #887 |
| P3-58 | Add a migration note for `Table2` and typed tables that shows the old and new forms | S | P3-11 | #887 |
| P3-59 | Write the authentication guide for the provider surface | S | P3-39 | #843 |
| P3-60 | Write the credentials migration guide that shows the old and new forms | S | P3-39 | #843 |
| P3-61 | Write the external OIDC guide | S | P3-42 | #846 |
| P3-62 | Add external OIDC troubleshooting | S | P3-43 | #846 |

#### Live tests and release

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P3-63 | `test(e2e)`: add live scenarios for `Table2` and typed tables | S | P3-12 | new |
| P3-64 | `test(e2e)`: add live scenarios for the new credential providers | S | P3-39 | new |
| P3-65 | Action: release v2.2.0 with every deprecation notice stamped | — | Every unit in milestone v2.2.0 | — |

### Phase 4: v3.0.0

#### Scope of v3.0.0

- `Table2()` becomes `Table()`. The legacy `Table(name string)` and
  `Table2()` are both removed
  ([ADR 018](adrs/018-table-collection-builder.md)).
- Typed tables and the exported foundation are the default documented path
  ([ADR 017](adrs/017-hybrid-domain-first-surface.md) and
  [ADR 019](adrs/019-typed-tables.md)).
- The credentials package exposes only providers, options, `TokenStore`, and
  sentinels, and the v1 adapter is removed
  ([ADR 015](adrs/015-credentials-provider-surface.md)).
- Categorized errors are the default error mapping
  ([ADR 021](adrs/021-best-effort-error-categorization.md)).
- Every symbol marked `Deprecated:` is removed.
- The module path ends in `/v3`.

#### Units

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P4-01 | Action: confirm that every bridge has shipped and every deprecation notice is stamped | — | P3-65 | — |
| P4-02 | Action: create the last v2 maintenance branch from the last v2 tag, and confirm its ruleset | — | P4-01 | — |
| P4-03 | `chore(ci)`: restore the backport workflow | XS | P4-02, P1-03 | new |
| P4-04 | `chore(ci)`: restore the forward-port tracker on `main` and on the v2 maintenance branch | XS | P4-02 | new |
| P4-05 | `chore(ci)`: restore the maintenance-label workflow | XS | P4-02 | new |
| P4-06 | Action: open a test backport pull request against the v2 maintenance branch | — | P4-03, P4-04 | — |
| P4-07 | `refactor!`: change the module path to `/v3`, and decide whether to rename `tests/integration/v2/`. This change is mechanical and atomic, so the size limit doesn't apply | M | P4-02 | new |
| P4-08 | `chore(release)`: update the release-please configuration for v3 | XS | P4-07 | new |
| P4-09 | `refactor(tableapi)!`: rename `Table2` to `Table`, and remove the legacy `Table` | S | P4-07 | new |
| P4-10 | `refactor(credentials)!`: remove the `NewXCredential` constructors | S | P4-07 | new |
| P4-11 | `refactor(credentials)!`: remove `BaseAccessTokenProvider` and `Authority` | S | P4-07 | new |
| P4-12 | `refactor(credentials)!`: remove `WithInstance` and `WithURL` | XS | P4-07 | new |
| P4-13 | `refactor(credentials)!`: remove the v1 `GetAuthentication` adapter | XS | P4-07 | #831 |
| P4-14 | `refactor!`: remove the sentinel aliases, one pull request per package | XS each | P4-07 | new |
| P4-15 | `feat(core)!`: make categorized errors the default error mapping | S | P4-07 | new |
| P4-16 | `refactor!`: remove the remaining deprecated symbols, one pull request per package | XS–S | P4-09 through P4-15 | new |
| P4-17 | `docs`: write the v2-to-v3 migration guide for the module path and tables | S | P4-09 | new |
| P4-18 | `docs`: write the v2-to-v3 migration guide for credentials and errors | S | P4-13, P4-15 | new |
| P4-19 | `docs`: make typed tables the default path in the getting started page and the user guide | S | P4-09 | new |
| P4-20 | Action: snapshot the v2 documentation | — | P4-02 | — |
| P4-21 | Action: merge the held release pull request to ship v3.0.0 | — | P4-06 through P4-20 | — |

### Phase 5: Expand CI

Goal: replace the local reusable workflows and composite actions with the
shared `NerdIT-Tech/.github` workflows. This phase starts after CI has stayed
green for 30 days following Phase 1.

| ID | Unit | Size | Depends on | Issue |
|---|---|---|---|---|
| P5-00 | Action: map each local reusable workflow to a shared workflow, or mark it as staying local | — | P1-11 | — |
| P5-ADR | `docs(adrs)`: record the decision to depend on shared organization workflows | S | P5-00 | new |
| P5-01 to P5-14 | `chore(ci)`: replace one local reusable workflow with its shared workflow, pinned by commit SHA, and delete the local copy. One unit for each of `build-go`, `check-go-deps`, `check-license`, `check-snippets`, `codeql`, `docs-build`, `docs-gate`, `govulncheck`, `lint-go`, `lint-pr-title`, `scorecard`, `stale-issues`, `test-go`, and `zizmor` | XS–S each | P5-ADR | new |
| P5-15 | Action: decide whether to rename the workflow files, keeping the required check names unchanged | — | P5-14 | — |
| P5-16 | Action: run the Phase 5 exit check | — | P5-01 through P5-15 | — |

## How existing issues split

Where a unit list splits an existing issue, the original issue becomes the
parent that tracks its units:

| Parent issue | Units |
|---|---|
| #883 | P3-02 to P3-04 |
| #885 | P3-07 to P3-09 |
| #886 | P3-10, P3-11 |
| #881 | P3-19 to P3-23 |
| #835 | P3-27, P3-28 |
| #837 | P3-29, P3-30 |
| #838 | P3-31 to P3-33 |
| #842 | P3-37 to P3-39 |
| #832 | #823, #829, #844, #845, #846 |
| #823 | P3-44a to P3-44c |
| #499 | P3-46 to P3-48 |
| #887 | P3-56 to P3-58 |
| #843 | P3-59, P3-60 |
| #846 | P3-61, P3-62 |

Two new parent issues track the v3 work. One covers removing the deprecated
surfaces (P4-09 to P4-16), and one covers restoring the maintenance-branch
tooling (P4-03 to P4-05).

## Issue map

| Destination | Issues |
|---|---|
| Milestone v2.2.0 | #499, #585, #636, #834–#842, #843, #881, #882, #883, #884, #885, #886, #887 |
| Milestone v2.2.0 or v2.3.0 | #823, #829, #832, #844, #845, #846 |
| Milestone v3.0.0 | #831, plus the two new v3 parent issues |
| No milestone | #464, #466, #468, #498, #642 |
| After v3 | #501, #568, #631, #632, #633 |
| To close | #471, #500, #560 (if covered), #812, #813, #872, #690, plus #580, #897, and #917 when their units are done |

## Exit checks

Each code unit passes these checks before it merges:

- `mise run check` passes locally. Until Step 0 is done, run `go build ./...`,
  `go test ./...`, `go test -tags integration ./tests/integration/...`, and
  `golangci-lint run ./...` directly.
- New exported API has table-driven tests.
- Documentation changes pass Vale.

Each phase is done when its checks pass:

- **Step 0:**
  - On a fresh devcontainer, `mise install` installs every pinned tool.
  - `mise run check` passes on `main`.
  - Each `just` recipe has a mise task that does the same thing.
  - The `justfile` is gone, and no page or instruction file mentions `just`.
  - CI and `mise.toml` pin the same tool versions.

- **Phase 0:**
  - A pull request merges without an admin bypass. This criterion waits for
    the Terraform migration (P0-06).
  - CI, Documentation Stable, CodeQL, and zizmor are green on `main`.
  - The Dependabot pull requests merge.
  - v2.1.1 and its security advisory are published.
- **Phase 1:**
  - The workflow directory matches the workflow inventory page.
  - Every remaining workflow is green or passes a manual dry run.
  - zizmor reports no findings.
- **Phase 2:**
  - The shim script produces no diff in CI.
  - `AGENTS.md` and `.agents/` name no vendor outside the shim notes.
  - Local agents and OpenCode Review read the same agent definitions.
- **Phase 3:**
  - Every replaced surface carries a `Deprecated:` notice.
  - The integration and nightly end-to-end runs cover `Table2`, typed tables,
    and the new credential providers.
- **Phase 4:**
  - A sample consumer module builds against `/v3`.
  - No Go file contains `Deprecated:`.
  - A backport pull request to the v2 maintenance branch merges.
- **Phase 5:**
  - No workflow references a local reusable workflow or composite action.
  - The required checks report under their expected names.
  - CI stays green for a week.
