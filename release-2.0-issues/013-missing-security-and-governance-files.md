# Missing SECURITY.md and CODEOWNERS; ancient dependabot-merge action

- **Priority:** P2
- **Raised by:** Senior DevOps (with Senior Product Manager)
- **Area:** Governance / supply chain

## Problem

1. **No `SECURITY.md`.** An SDK that handles ServiceNow credentials (OAuth2 flows, JWT
   bearer, ROPC in `credentials/`) should tell researchers how to report a
   vulnerability privately. GitHub surfaces the absence on the repo's Security tab.
2. **No `CODEOWNERS`.** With auto-merge enabled for Dependabot PRs (`pr.yml`), there is
   no required-reviewer backstop on workflow or credential-path changes.
3. **`dependabot/fetch-metadata@v3.1.0`** in `pr.yml` is several major versions old
   (current is v2.x under the new versioning... verify — the action renumbered; v3.1.0
   predates the consolidation). While auditing: the auto-merge rule merges anything
   non-semver-major, including GitHub Actions bumps, with **no passing-checks
   requirement expressed in the workflow** — it relies entirely on branch protection
   requiring status checks. Confirm branch protection actually requires the CI jobs,
   otherwise a red Dependabot PR can auto-merge (and note issue 001 means the test
   check was never a real gate anyway).

## Recommendation

1. Add `SECURITY.md` with a private-reporting channel (GitHub private vulnerability
   reporting is one click to enable) and the supported-versions table (v2.x supported;
   v1.x critical-fixes-only per issue 008).
2. Add `.github/CODEOWNERS` covering at minimum `/.github/`, `/credentials/`, and
   `go.mod`.
3. Bump `fetch-metadata`, and verify branch-protection required checks match the
   current CI job names (they changed names during the v2 rework — a renamed job
   silently stops being "required").
