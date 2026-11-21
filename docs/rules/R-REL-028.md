---
id: R-REL-028
applies_to: [".github/workflows/*.yml", ".github/workflows/*.yaml", "**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# An API-compatibility check blocks incompatible changes on release branches

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- Verify fails until the job lands, which is the intended signal for a proposed rule. -->

MUST: Run `gorelease` or `apidiff` in `PR Gate` against the latest tag on the same release line as the PR when Go code changes.
MUST: Fail the job on an incompatible exported-API change in a PR into `release/*`.
MUST NOT: Fail the job on PRs into `main`; report each incompatible change instead.
SHOULD: Add `!` to the PR title or a `BREAKING CHANGE:` footer when the report flags a change on `main`.
MUST: Pin the tool under R-CI-026.

## Verify

```bash
# Fails when no workflow runs an API-compatibility tool.
grep -RqE 'gorelease|apidiff' .github/workflows
```

## Rationale

An incompatible change on a maintenance line breaks every consumer pinned to that major. See [ADR 028](../../website/docs/contributing/adrs/028-api-compatibility-check.md).
