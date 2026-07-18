# `/v2` module path bump — release-day runbook item

- **Priority:** P0 — release blocker (deliberately deferred; must not be forgotten)
- **Raised by:** Senior Principal Engineer
- **Area:** Go module semantics

## Problem

Go's semantic import versioning requires the module path to end in `/v2` for any
`v2.x.y` tag. `go.mod` still declares
`module github.com/michaeldcanady/servicenow-sdk-go`. This deferral is **intentional**
(to keep the branch mergeable during development), but if the v2.0.0 tag is cut before
the path changes, `go get` of the tagged version will fail module verification and the
tag is burned permanently — tags can't be re-pointed once the proxy caches them.

## Recommendation

Execute immediately before the release merge (issue 002), as one commit:

1. `go.mod`: `module github.com/michaeldcanady/servicenow-sdk-go/v2`
2. Rewrite **every** internal self-import (all packages import siblings by full module
   path — `core`, `internal/...`, every `*api` package, `credentials`, `query`, tests):
   ```bash
   grep -rl 'michaeldcanady/servicenow-sdk-go' --include='*.go' . \
     | xargs sed -i 's|michaeldcanady/servicenow-sdk-go|michaeldcanady/servicenow-sdk-go/v2|g'
   go build ./... && go test ./...
   ```
   (Guard against double-applying `/v2/v2`; a `mod upgrade` tool like
   `github.com/marwan-at-work/mod` does this safely.)
3. Update `release-please-config.json` `package-name` to the `/v2` path (both configs —
   stable and weekly).
4. Update install instructions in `Readme.md` and `docs/` to `go get .../v2`.
5. Only then merge to `main` and let release-please tag `v2.0.0`.

## Cross-references

- Depends on / sequenced with: issue 002 (release path).
