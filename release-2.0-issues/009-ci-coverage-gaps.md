# CI coverage gaps: integration tests, vuln scanning, pinned lint version

- **Priority:** P1
- **Raised by:** Senior DevOps
- **Area:** CI/CD

## Problem

Several classes of verification exist in the repo but never run in CI, and some CI
inputs are nondeterministic:

1. **Integration tests never run in CI.** `tests/integration/` (godog + httpmock, no
   live instance needed — explicitly designed to be CI-safe) is behind
   `//go:build integration`, and no workflow passes `-tags integration`. The entire BDD
   suite is dead weight unless someone runs it locally.
2. **The `preview.query` build tag is linted but never tested.** `.golangci.yml` sets
   `build-tags: [preview.query]`, but `go test` in CI runs without tags, so
   tag-gated code is compiled by the linter and never exercised by tests.
3. **No vulnerability scanning for Go deps.** CodeQL runs, but `govulncheck` (which
   checks the actual call graph against the Go vuln DB) does not. For an SDK whose
   consumers inherit its dependency tree, this is table stakes.
4. **`golangci-lint-action` uses `version: latest`** — lint results change under your
   feet when a new linter version releases; PRs go red for reasons unrelated to their
   diff. Same for `go install github.com/mfridman/tparse@latest` in the test job.
5. **Stale lint exclusions**: `.golangci.yml` excludes `dupl` for
   `(actsub-api|documents-api|table-api)/.*` — paths that no longer exist after the
   directory renames, so the exclusion is dead config.

## Recommendation

1. Add an `integration-test` job: `go test -tags integration ./tests/integration/...`
   (Linux-only is fine; it's mocked).
2. Add a unit-test matrix leg or extra step with `-tags preview.query`.
3. Add a `govulncheck ./...` job (weekly schedule + per-PR).
4. Pin `golangci-lint` to a released version and bump via Dependabot/renovate;
   pin `tparse` with a version suffix.
5. Delete the dead `dupl` path exclusion or update it to the current directory names.
