---
title: 'ADR 019: Typed tables — built-in and user-defined, full parity, uniform shape'
description: >-
  Strongly-typed tables ship built-in and are user-extensible via a wrapper
  variable over the exported foundation; typed fields come from field
  descriptors, user tables match built-ins, and tables stay free of domain
  actions.
---

# ADR 019: Typed tables — built-in and user-defined, full parity, uniform shape

## Status

Proposed. Accepted on merge; this record and its rule file
(`docs/rules/R-TBL-019.md`) ship together with the design-session work.

## Context

The hybrid surface (ADR 017) needs strongly-typed tables — Incident, Case, and
whatever a consumer's own instance holds (`MyCoolTable`) — with typed fields
for both query and access. Several design questions had to be settled together
because they constrain each other:

- **How does a consumer define their own typed table?** Go has no method
  overloading and no way to attach methods to an anonymous chained value, so a
  user cannot add methods to `client.Now().Table2()` inline. A named type
  attached to a variable is the only way to extend the surface with methods.
- **Do built-in and user-defined tables share one foundation, or are user
  tables second-class?** The `internal/store` and `internal/serialization`
  helpers are currently unreachable by SDK consumers, so without change a user
  table could not match a built-in.
- **Do typed tables carry domain actions** (e.g. "close this incident", "add a
  comment to this case")? That pulls every module toward bespoke wrapper types,
  which ADR-005 already rejected for pagination.
- **Do typed references navigate?** A field referencing another record could
  lazily fetch it, or just carry its identity.

Alternatives considered for user extension: generic field accessors only (no
methods, so no ergonomic user API); mandatory codegen (heavyweight, and not
every consumer wants a build step); or a closed set of built-in tables only
(fails the "build your own" requirement).

## Decision

- **Built-in typed tables** (Incident, Case first, growing) ship with the SDK.
  Consumers use but cannot modify them; they chain literally off the collection
  builder (ADR 018).
- **User-defined typed tables** wrap the exported collection builder in a named
  **wrapper variable**: `tables := mycool.New(client.Now().Table2())`. The
  wrapper type carries the methods; built-ins chain literally, user tables go
  through the wrapper var. This is a direct consequence of Go's static typing.
- **Field descriptors** (a storage key plus a typed accessor) drive both typed
  fields and typed query (`inc.Priority().Eq(IncidentPriority.High)`), coexisting
  with the fluent builder (`query.Where(...)`).
- **Full parity.** The `internal/store` and `internal/serialization` helpers are
  exported — or a convenience base type wraps them — so a user table has the
  same accessor generation and backing-store behavior as a built-in. No
  second-class citizen.
- **Uniformity.** Typed tables expose record CRUD and query, **not** domain
  actions. Domain actions live in scoped API packages (echoing ADR-005's
  one-generic-pattern philosophy and ADR-007's restraint).
- **References** are a typed wrapper carrying sys_id, display value, and link;
  they do **not** auto-navigate to the referenced record.
- Optional codegen may generate field descriptors, but the exported helpers are
  the contract — codegen is convenience, not a requirement.

## Consequences

- **Pros:** consumers build first-class typed SDKs with the same power as the
  built-ins; one foundation to maintain; typed query is discoverable and
  strongly typed while the fluent builder remains; tables stay predictable
  instead of sprouting domain actions.
- **Cons:** the wrapper-var pattern is initially surprising (it needs a short
  explanation of Go static typing in the docs); exporting the store/serialization
  helpers widens the public surface and must be done deliberately; references
  that do not navigate may surprise callers expecting lazy loading.
- **Rule for future typed-table work:** any capability added to a built-in table
  must be reachable by a user table through the exported foundation; if a change
  would give built-ins something users cannot replicate, it violates this
  decision. Add domain actions in scoped packages, not on table types.

See also: ADR 002 (backing-store models), ADR 005 (one generic pattern), ADR
007 (no speculation), ADR 017 (hybrid surface), ADR 018 (entry surface).
