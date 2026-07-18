# ADR 003: Hand-written client on Kiota runtime abstractions

## Status

Accepted (founding decision, documented retroactively)

## Context

The SDK needs a consistent, typed client surface for many ServiceNow REST
APIs. The obvious options:

1. **Generate** a client with the Kiota CLI from OpenAPI descriptions.
2. **Hand-write** everything, including HTTP/serialization plumbing.
3. **Hand-write** the client surface on top of Kiota's runtime libraries.

ServiceNow does not publish complete, accurate OpenAPI documents for its REST
APIs, so option 1 generates an incomplete SDK — or requires hand-maintaining
specs, which moves the hand-work one step earlier without removing it.
ServiceNow's response shapes (the `result` envelope, three-faceted fields with
value/display-value/link, table-generic endpoints serving every schema) also
map poorly onto generator output.

Option 2 reinvents URI-template expansion, auth plumbing, serialization
registries, retry middleware, and backing stores — none of which are
differentiating.

## Decision

Option 3: build directly on Microsoft's Kiota runtime
(`kiota-abstractions-go`, `kiota-http-go`, `kiota-serialization-*-go`), but
hand-write the request builders, models, and response envelopes, deliberately
matching the conventions of Kiota-*generated* SDKs (msgraph-sdk-go and
friends): request-builder chaining, `RequestConfiguration` shapes, parsable
factories, backed models.

## Consequences

- **Pros:** full control over ServiceNow-specific shapes
  (`RecordElement`, `ServiceNowCollectionResponse[T]`); the runtime's real
  work (middleware, auth, serialization, URI templates) is free and
  battle-tested; developers arriving from other Kiota SDKs find familiar
  idioms.
- **Cons:** the uniformity a generator would enforce mechanically must be
  enforced by convention instead. This is the root cause of the repo's strict
  module pattern — the constructor triad, one-file-per-verb layout, and the
  "Add a new API module" playbook are the hand-written substitute for
  generated code.
- **Tie-breaker rule for reviews:** when a structural question has no local
  precedent, the answer is whatever msgraph-sdk-go does.
