# 2.0 Release Readiness Review — Issue Tracker

Findings from a three-role review (Senior Product Manager, Senior Principal Engineer,
Senior DevOps) of the `release/2.0` branch, conducted 2026-07-17. Current state at
review time: branch is 90 commits ahead of `main`, `go build ./...` clean,
`go test ./...` passing locally, version at 1.12.0.

## Release blockers (P0)

| # | Issue | Owner role |
| --- | --- | --- |
| [001](001-ci-test-failures-never-fail-the-build.md) | CI test failures never fail the build (`report-tests.sh` always exits 0) | DevOps |
| [002](002-no-release-path-from-release-2.0-to-v2.0.0-tag.md) | No defined path from `release/2.0` to a v2.0.0 tag; non-conventional commits and missing `BREAKING CHANGE` footer would yield v1.13.0 | DevOps + PM |
| [003](003-v2-module-path-runbook.md) | `/v2` module path bump — deferred by design, needs a release-day runbook so the tag isn't burned | Engineering |

## High priority (P1) — do before GA

| # | Issue | Owner role |
| --- | --- | --- |
| [004](004-nil-receiver-guards-return-nil-nil.md) | `return nil, nil` nil-receiver guards silently swallow bugs; fix the contract while breaking changes are free | Engineering |
| [005](005-consolidate-error-sentinels.md) | Three error-sentinel locations + inline `errors.New` duplicates break `errors.Is` | Engineering |
| [006](006-repo-hygiene-stray-working-files.md) | Scratch files (`coverage.html`, `fix_error_mappings.py`, `files_to_fix.txt`, ...) committed at repo root | PM + Engineering |
| [007](007-readme-links-all-broken.md) | Every Readme API-table link 404s (old hyphenated directory names) | PM |
| [008](008-v2-migration-guide-and-launch-comms.md) | No v1→v2 migration guide, v1 support policy, or launch checklist | PM |
| [009](009-ci-coverage-gaps.md) | Integration tests never run in CI; no govulncheck; unpinned lint/tool versions; stale lint exclusions | DevOps |
| [014](014-public-api-surface-audit.md) | Final exported-surface audit before the API freezes at v2.0.0 | Engineering |

## Medium priority (P2) — decide before GA, fix opportunistically

| # | Issue | Owner role |
| --- | --- | --- |
| [010](010-docs-pipeline-inefficiencies.md) | Docs workflow: stale `pages.yml` filter, deploy rebuilds instead of using its artifact | DevOps |
| [011](011-dual-release-please-channels.md) | Weekly-preview and stable release-please channels share one changelog/tag namespace | DevOps |
| [012](012-go-directive-and-test-deps.md) | `go 1.25.0` floor limits adoption; godog/httpmock/godotenv pollute consumer-visible go.mod | Engineering |
| [013](013-missing-security-and-governance-files.md) | No SECURITY.md/CODEOWNERS; verify Dependabot auto-merge is gated by required checks | DevOps + PM |

## Suggested sequencing

1. **Now:** 001 (real test gate) → then everything else merges against honest CI.
2. **In parallel:** 004+005 (one error-contract sweep), 006, 007, 009, 014.
3. **Product track:** 008 (migration guide + support policy), 011 decision, 013.
4. **Release day:** 003 then 002, in that order, per their runbooks.

## What we deliberately did not flag

- The Kiota-style hand-written builder pattern itself — it's consistent, tested, and
  matches the msgraph-sdk-go idiom the project intentionally targets.
- `VERSION`/`CHANGELOG.md` automation — release-please wiring is correct for the
  steady state; only the 2.0 transition (002/003) and channel overlap (011) need work.
