---
id: R-BUILD-007
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: speculative builder chain, dead-end chain, stub, coming soon,
unimplemented, reserve url segment, scaffold -->

MUST NOT add a request-builder accessor before the verb methods it leads to are implemented and shipping together.
MUST remove builder-chain segments that have no reachable verb operation rather than leaving them as stubs.
SHOULD add a builder only once the operations it exposes are ready to ship together.

## Verify

```bash
# Fails (non-zero) if an API package contains a stub / not-implemented marker,
# the tell-tale of a speculatively scaffolded chain.
! grep -RInE 'coming soon|not implemented|unimplemented' --include='*.go' *api/
```

## Rationale

A fluent chain that compiles and navigates but silently does nothing is worse than a compile error and contradicts ADR-006's fail-loud principle; builders are added when their verbs are ready, not to reserve a URL segment (see ADR 007).
