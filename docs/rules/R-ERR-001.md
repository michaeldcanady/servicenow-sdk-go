---
id: R-ERR-001
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: error handling, error standardization, sentinel errors, error
messages, phrasing, centralize errors, errors package -->

MUST add new shared sentinel errors to the `errors` package rather than declaring ad-hoc `errors.New` values inline in API packages.
MUST phrase a nil-parameter failure as `"<parameter> cannot be nil"`.
MUST phrase a missing-input failure as `"<parameter> is required"`.
MUST NOT use contractions in error message strings — write `cannot`, not `can't`.
MUST NOT scatter hard-coded error message literals across API packages when a shared sentinel already covers the failure.
SHOULD centralize repeated validation logic (such as parameter nil checks) in the `errors` package where feasible.

## Verify

```bash
# Fails (non-zero) if a Go string literal still contains a contraction.
! grep -RInE "\"[^\"]*can't[^\"]*\"" --include='*.go' .
```

## Rationale

v1 accumulated hundreds of similar-but-different error strings, which broke `errors.Is` for every consumer; central sentinels and one phrasing convention restore a stable error identity. A fresh inline `errors.New` reintroduces the bug (see ADR 001, ADR 006).
