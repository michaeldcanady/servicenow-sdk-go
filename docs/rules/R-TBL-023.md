---
id: R-TBL-023
applies_to: ["tables/**/*.go", "model/**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Typed references never auto-navigate

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- TODO(owner): replace the type name in Verify with the real one. -->

MUST: Represent a typed reference as a wrapper carrying `sys_id`, display value, and link.
MUST NOT: Auto-navigate a typed reference to the referenced record.
MUST: Resolve a reference only through an explicit call by the caller.

## Verify

```bash
# Fails closed (exit 1) if no generic reference wrapper is declared.
grep -RqE 'type Reference[A-Za-z]*\[' --include='*.go' .
```

## Rationale

ServiceNow reference fields arrive as a link plus value, and following one costs a network request. An implicit follow inside a getter would hide that cost, those failures, and access-control surprises behind what looks like a field read.
