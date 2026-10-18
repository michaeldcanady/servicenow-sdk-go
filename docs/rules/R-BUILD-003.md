---
id: R-BUILD-003
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: kiota, hand-written, runtime abstractions, openapi, generator,
msgraph, tie-breaker, request builders, from scratch -->

MUST build on the Kiota runtime libraries (`kiota-abstractions-go`, `kiota-http-go`, `kiota-serialization-*-go`) instead of hand-rolling HTTP, auth, URI-template, or serialization plumbing.
MUST hand-write request builders, models, and response envelopes to match Kiota-generated SDK conventions (request-builder chaining, `RequestConfiguration` shapes, parsable factories, backed models).
MUST NOT generate the client from OpenAPI with the Kiota CLI.
MUST NOT reimplement generic plumbing the Kiota runtime already provides.
SHOULD treat msgraph-sdk-go as the tie-breaker when a structural question has no local precedent.

## Verify

```bash
# Fails (non-zero) if API/core packages import net/http directly,
# bypassing the Kiota request adapter.
! grep -RInE '"net/http"' --include='*.go' *api/ core/
```

## Rationale

ServiceNow publishes no reliable OpenAPI and its response shapes map poorly to generator output, so the surface is hand-written on Kiota's battle-tested runtime rather than generated or built from scratch; msgraph-sdk-go is the structural tie-breaker (see ADR 003).
