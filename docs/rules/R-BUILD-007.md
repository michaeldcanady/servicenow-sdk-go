---
id: R-BUILD-007
applies_to: ["*api/**/*_builder.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Builders ship together with their verbs

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST NOT: Add a request-builder accessor before the verb methods it leads to are implemented.
MUST: Ship a builder and its verb methods in the same release.
MUST: Remove an unreleased builder-chain segment that has no reachable verb, instead of leaving a stub.
MUST NOT: Remove a builder accessor that has shipped, outside a major version bump (see R-NAME-010, R-REL-025).

## Verify

```bash
# Heuristic: catches common stub markers only; allowlist legitimate error text explicitly.
# grep exit 1 = no match = pass; exit 2 fails closed.
grep -RInE 'coming soon|not implemented|unimplemented|TODO.*implement' --include='*.go' --exclude='*_test.go' --exclude-dir=.git .; [ $? -eq 1 ]
```

## Rationale

A fluent chain that compiles and navigates but silently does nothing is worse than a compile error, and it contradicts the fail-loud principle (see R-ERR-006). Builders are added when their verbs are ready, not to reserve a URL segment.
