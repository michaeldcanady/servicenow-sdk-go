# 2.0 Release Readiness Review — Issue Tracker

Findings from a three-role review (Senior Product Manager, Senior Principal Engineer,
Senior DevOps) of the `release/2.0` branch, conducted 2026-07-17. Current state at
review time: branch is 90 commits ahead of `main`, `go build ./...` clean,
`go test ./...` passing locally, version at 1.12.0.

## Status (2026-07-18)

Fix PRs are open against `release/2.0`:
#481 (001), #482+part of 002 (009), #483 (010), #484 (006), #485 (007), #486 (013), #487 (004 + part of 005).
Issue 012's go-directive item is blocked by kiota-http-go's own `go 1.25.0` floor (see the issue file).
Still open: 002/003 (release-day runbooks), 005 (remaining sentinel consolidation), 008, 011 (decision), 012 (test-module split), 014.

## Status (2026-07-24)

Follow-up codebase-exploration audit against `release/2.0` @ `aca3942` (post
ADR-006 / PR #550):

- **006, 007, 013 confirmed done.** No action. (007 has one harmless cosmetic
  straggler — the Table API issue-label badge URL in `Readme.md:40` — not worth its
  own tracker item; note left here in case someone wants to sweep it opportunistically.)
- **004's original citations were stale** (the `tableapi` lines it named are fixed).
  Doc corrected in place to reflect that, and to redirect the still-open part of its
  scope — three CDM packages (`cdmapplicationsapi`, `cdmchangesetapi`,
  `cdmeditorapi`) that have **zero** nil-receiver guards at all, contradicting
  ADR-006's "applied everywhere" claim — to new issue
  [#551](https://github.com/michaeldcanady/servicenow-sdk-go/issues/551).
  **#551 blocks the v2.0.0 tag**: it's a code bug (nil-pointer panic instead of a
  documented sentinel error) in public API surface that freezes at GA, not a
  cosmetic/doc issue.
- **005 confirmed still open**, with two new concrete duplicate-sentinel examples
  added as evidence (`ErrNilContext`, `ErrNilResponse`) plus a fourth sentinel
  micro-location (`attachmentapi/errors.go`) folded into scope. Does not block the
  tag by itself — see reasoning in doc 005 and the audit notes — but should ship in
  the same window as #551 since both touch `snerrors.ErrNilRequestBuilder`/sentinel
  identity.
- 001, 002, 003, 014 re-confirmed still open as previously tracked; no new
  information this pass.

## Release blockers (P0)

| # | Issue | Owner role |
| --- | --- | --- |
| [001](001-ci-test-failures-never-fail-the-build.md) | CI test failures never fail the build (`report-tests.sh` always exits 0) | DevOps |
| [002](002-no-release-path-from-release-2.0-to-v2.0.0-tag.md) | No defined path from `release/2.0` to a v2.0.0 tag; non-conventional commits and missing `BREAKING CHANGE` footer would yield v1.13.0 | DevOps + PM |
| [003](003-v2-module-path-runbook.md) | `/v2` module path bump — deferred by design, needs a release-day runbook so the tag isn't burned | Engineering |
| [#551](https://github.com/michaeldcanady/servicenow-sdk-go/issues/551) | Three CDM `*api` packages have zero nil-receiver guards, panic instead of returning `snerrors.ErrNilRequestBuilder`, contradicting ADR-006 | Engineering |

## High priority (P1) — do before GA

| # | Issue | Owner role |
| --- | --- | --- |
| [004](004-nil-receiver-guards-return-nil-nil.md) | `return nil, nil` nil-receiver guards silently swallow bugs; original defect fixed, doc now points remaining CDM gap at #551 | Engineering |
| [005](005-consolidate-error-sentinels.md) | Three (now four) error-sentinel locations + inline `errors.New` duplicates break `errors.Is` | Engineering |
| [006](006-repo-hygiene-stray-working-files.md) | Scratch files (`coverage.html`, `fix_error_mappings.py`, `files_to_fix.txt`, ...) committed at repo root — **done** | PM + Engineering |
| [007](007-readme-links-all-broken.md) | Every Readme API-table link 404s (old hyphenated directory names) — **done**, one cosmetic badge-URL straggler remains | PM |
| [008](008-v2-migration-guide-and-launch-comms.md) | No v1→v2 migration guide, v1 support policy, or launch checklist | PM |
| [009](009-ci-coverage-gaps.md) | Integration tests never run in CI; no govulncheck; unpinned lint/tool versions; stale lint exclusions | DevOps |
| [014](014-public-api-surface-audit.md) | Final exported-surface audit before the API freezes at v2.0.0; confirmed `New<X>RequestBuilderInternal` constructors are exported (not internal) across `*api` packages | Engineering |

## Medium priority (P2) — decide before GA, fix opportunistically

| # | Issue | Owner role |
| --- | --- | --- |
| [010](010-docs-pipeline-inefficiencies.md) | Docs workflow: stale `pages.yml` filter, deploy rebuilds instead of using its artifact | DevOps |
| [011](011-dual-release-please-channels.md) | Weekly-preview and stable release-please channels share one changelog/tag namespace | DevOps |
| [012](012-go-directive-and-test-deps.md) | `go 1.25.0` floor limits adoption; godog/httpmock/godotenv pollute consumer-visible go.mod | Engineering |
| [013](013-missing-security-and-governance-files.md) | No SECURITY.md/CODEOWNERS; verify Dependabot auto-merge is gated by required checks — **done** | DevOps + PM |

## Suggested sequencing

1. **Now:** 001 (real test gate) → then everything else merges against honest CI.
2. **In parallel:** [#551](https://github.com/michaeldcanady/servicenow-sdk-go/issues/551) + 005 (one error-contract sweep, now that 004's original scope is done), 009, 014.
3. **Product track:** 008 (migration guide + support policy), 011 decision.
4. **Release day:** 003 then 002, in that order, per their runbooks.

## What we deliberately did not flag

- The Kiota-style hand-written builder pattern itself — it's consistent, tested, and
  matches the msgraph-sdk-go idiom the project intentionally targets.
- `VERSION`/`CHANGELOG.md` automation — release-please wiring is correct for the
  steady state; only the 2.0 transition (002/003) and channel overlap (011) need work.
- TODO/FIXME/HACK markers, stub bodies, ADR-007 speculative-builder-chain
  violations, skipped/empty tests — none found in a 2026-07-24 sweep.
- The missing `/v2` go.mod suffix — known, intentionally deferred to release day
  (issue 003), not a new finding.
- The NerdIT-Tech org question — resolved via ADR-010 / PR #550 (issue #496,
  closed).
