---
id: R-ERR-001
applies_to: ["errors/**/*.go", "*api/**/*.go", "core/**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Shared sentinel errors live in the errors package

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Declare each shared sentinel error once, in the `errors` package (imported as `snerrors`).
MUST: Return or wrap an existing shared sentinel when it already describes the failure.
MUST NOT: Declare a package-level `errors.New` sentinel in any package other than `errors`.
MUST NOT: Repeat an error message literal that a shared sentinel already provides.
SHOULD: Route repeated validation failures (nil parameter, missing input) through shared helpers so wording stays uniform.

## Verify

```bash
# grep exit 1 = no match = pass; exit 2 (bad path) fails closed.
grep -RInE '^\s*(var\s+)?Err[A-Za-z0-9_]*\s*=\s*errors\.New\(' --include='*.go' --exclude='*_test.go' --exclude-dir=errors --exclude-dir=.git .; [ $? -eq 1 ]
```

## Rationale

v1 accumulated hundreds of similar-but-different inline error strings, which broke `errors.Is` for every consumer. One home for sentinels restores a stable error identity. Message phrasing conventions moved to CONTRIBUTING (see LEDGER.md).
