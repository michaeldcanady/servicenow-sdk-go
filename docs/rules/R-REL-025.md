---
id: R-REL-025
applies_to: ["**/*.go", "RELEASE.md"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Deprecation marks replacement and removal; discouragement is documented separately

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Mark an exported API that is being replaced or removed with a `// Deprecated:` godoc paragraph.
MUST: Name the replacement in the Deprecated paragraph and state when removal is planned.
MUST: Spell the marker exactly `Deprecated:` (capital D, in its own paragraph) so tooling recognizes it.
MUST NOT: Mark an API `Deprecated:` when it has no replacement and no planned removal.
MUST: Give a discouraged-but-kept API a godoc paragraph beginning `Not recommended:` that says why and names the preferred alternative (for example, R-AUTH-016).
MUST NOT: Remove or incompatibly change an exported identifier outside a major version bump (R-NAME-010).
SHOULD: Keep a deprecated API for at least the rest of the major version in which it was deprecated.

## Verify

```bash
# Flags malformed markers (lowercase or all-caps). grep exit 1 = no match = pass; exit 2 fails closed.
grep -RInE '^\s*//\s*(deprecated|DEPRECATED):' --include='*.go' --exclude-dir=.git .; [ $? -eq 1 ]
```

## Rationale

This SDK reserves `Deprecated:` for APIs on their way out, so a marker always means "migrate" and tooling warnings stay meaningful. APIs that are discouraged but staying get a documented `Not recommended:` paragraph, which keeps the two signals distinct.
