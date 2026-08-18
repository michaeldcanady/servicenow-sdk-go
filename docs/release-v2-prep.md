# v2.0 Release Preparation

Status: **Code complete** — build, tests, and lint are green on `release/2.0`.
Last verified: 2026-08-18

---

## Release Blockers

These must be executed before the v2.0.0 tag is cut. Both are sequencing/automation
tasks, not code changes.

### 1. No path from release/2.0 to a v2.0.0 tag

**Issue:** [#556](https://github.com/michaeldcanady/servicenow-sdk-go/issues/556)

`stable-release.yml` runs release-please only on pushes to `main`. Nothing on
`release/2.0` knows how to produce a tag. Without intervention, release-please will
compute a **minor** bump (v1.13.0) instead of v2.0.0.

**Required action:**
1. Merge `release/2.0` into `main` via a **single squash-merge PR** whose title is
   conventional with a breaking-change footer (e.g. `feat!: v2 rework`).
2. The PR body must carry a `BREAKING CHANGE:` footer summarizing the v2 changes.
3. Verify the resulting release-please PR proposes **2.0.0** before merging it.
4. Extend `pr.yml` branch filter to include `release/*` so future release branches
   get title linting.

**Hazards:**
- Non-conventional commit messages on the branch will parse wrong in release-please.
- No `BREAKING CHANGE:` footer → release-please computes v1.13.0.
- Tags can't be re-pointed once the Go module proxy caches them.

### 2. Module path must change to /v2

**Issue:** [#557](https://github.com/michaeldcanady/servicenow-sdk-go/issues/557)

Go's semantic import versioning requires the module path to end in `/v2` for any
`v2.x.y` tag. `go.mod` currently declares
`module github.com/michaeldcanady/servicenow-sdk-go`.

**Required action (immediately before the release merge, as one commit):**
1. `go.mod`: change to `module github.com/michaeldcanady/servicenow-sdk-go/v2`
2. Rewrite every internal self-import:
   ```bash
   grep -rl 'michaeldcanady/servicenow-sdk-go' --include='*.go' . \
     | xargs sed -i 's|michaeldcanady/servicenow-sdk-go|michaeldcanady/servicenow-sdk-go/v2|g'
   go build ./... && go test ./...
   ```
   Guard against double-applying `/v2/v2`.
3. Update `release-please-config.json` `package-name` to the `/v2` path.
4. Update install instructions in `Readme.md` and `docs/` to `go get .../v2`.

**Sequence:** This must happen **after** the squash-merge PR is prepared but
**before** it is merged to `main` and the tag is cut.

---

## Pre-GA Must-Fixes

These should be addressed before the general availability announcement but are not
strictly required for the tag cut itself.

### 3. CI: ci.yml missing types on pull_request

**Issue:** [#580](https://github.com/michaeldcanady/servicenow-sdk-go/issues/580)

Retargeting a PR never starts CI — green checks display with zero Go tests run.

**Fix:** Add `types: [opened, reopened, edited, synchronize]` to `ci.yml`'s
`pull_request` trigger, matching `pr.yml`.

### 4. Pre-GA public API surface audit

**Issue:** [#564](https://github.com/michaeldcanady/servicenow-sdk-go/issues/564)

The v2 rework was ~90 commits. No deliberate final pass over exported identifiers.

**Checklist:**
- [ ] Generate surface diff: `go doc -all ./...` or `apidiff` v1-tag → release/2.0
- [ ] Review `...Internal` constructors — demote or document why public
- [ ] Confirm no `internal/store` types leak through exported signatures
- [ ] Confirm no unintended `// Deprecated:` v1 surfaces remain
- [ ] Ensure every exported identifier has a doc comment
- [ ] Run `api-module-consistency-reviewer` across all `*api` packages

### 5. Migration guide and release notes

**Issue:** [#559](https://github.com/michaeldcanady/servicenow-sdk-go/issues/559)

No consumer-facing artifact answers "I'm on v1.12 — what do I do?"

**Required:**
- [ ] Migration guide (side-by-side v1→v2 examples for client construction, table
  CRUD, queries, attachments, pagination, error handling, authentication)
- [ ] v2.0.0 release notes hand-curated (auto-generated changelog for a 90-commit
  squash won't tell the story)
- [ ] Launch checklist: pkg.go.dev renders `/v2` docs, badges resolve, docs site
  reflects v2

### 6. CI coverage gaps

**Issue:** [#560](https://github.com/michaeldcanady/servicenow-sdk-go/issues/560)

- [ ] Integration tests (`tests/integration/v2`, godog + httpmock) never run in CI —
  add a job: `go test -tags integration ./tests/integration/v2/...`
- [ ] `preview.query`-tagged code linted but never tested — add a test matrix leg
  with `-tags preview.query`
- [ ] No `govulncheck` — add weekly + per-PR vulnerability scanning
- [ ] `golangci-lint` uses `version: latest` — pin to a released version
- [ ] Dead `dupl` exclusion for renamed directories — remove or correct

---

## Current State

| Check | Status |
|-------|--------|
| `go build ./...` | Pass |
| `go test ./...` (31 packages) | Pass |
| `go vet ./...` | Pass |
| `golangci-lint run ./...` | 0 issues |
| Integration tests (local, `-tags integration`) | Pass |

**Open issues:** 27 total — 2 urgent (blockers above), 4 high (pre-GA above),
remainder deferred.

---

## Release Sequence (Ordered Checklist)

1. ~~Fix critical code bugs (#592, #593)~~ — Done (532cd6bb)
2. Complete the public API surface audit (#564)
3. Fix CI trigger (#580) and add integration test job (#560)
4. Write migration guide + hand-curate release notes (#559)
5. Execute the `/v2` module path bump (#557) as a single commit
6. Squash-merge `release/2.0` → `main` with `feat!: v2 rework` title and
   `BREAKING CHANGE:` footer (#556)
7. Verify release-please proposes **2.0.0** (not 1.13.0)
8. Tag and publish
9. Verify pkg.go.dev renders the `/v2` module correctly
10. Announce

---

## Post-Release Backlog

Deferrable to after v2.0.0 GA:

| Issue | Description |
|-------|-------------|
| #563 | Move tests into nested go.mod |
| #568 | Collapse nil-guard preamble into shared helper |
| #585 | CDM comma-ok sentinel error question |
| #502 | Backfill ADRs for existing design decisions |
| #501 | Spike: Kiota raw URL vs path parameters |
| #500 | Remaining authentication flows (device code, certificate) |
| #499 | BYO logging via WithLogger client option |
| #498 | Testing rigor epic (BDD policy + coverage ratchet) |
| #497 | License headers on Go files + CI enforcement |
