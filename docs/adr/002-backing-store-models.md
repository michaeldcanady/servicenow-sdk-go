# ADR 002: Backing-store-backed models

## Status

Accepted (implemented in the v2 rework, PR #474)

## Context

v1 models were plain structs (`TableEntry` was literally a `map[string]interface{}`).
Two problems followed:

1. **Absent vs. zero was indistinguishable.** ServiceNow omits fields freely
   (and `sysparm_fields` makes partial records routine). A struct field cannot
   distinguish "the instance sent an empty string" from "the instance did not
   send this field", which forced callers into guesswork and produced
   accidental data loss on write.
2. **Writes serialized whole objects.** Without change tracking, a `PUT`/`PATCH`
   body contained every field — including zero values that overwrote server-side
   data the caller never touched.

The SDK also aims to feel like Kiota-generated SDKs (see ADR 003), whose
models are store-backed.

## Decision

Every model embeds `core.BaseModel` and stores property data in a Kiota
`BackingStore` — a key/value store with change tracking — instead of struct
fields.

- Accessors are generated from `internal/store` helpers
  (`DefaultBackedModelAccessorFunc` / `DefaultBackedModelMutatorFunc`), giving
  every property a `GetX() (T, error)` / `setX(T) error` pair.
- Getters return **pointers**: `nil` means "not sent", distinct from a zero
  value.
- Serialization walks the store's dirty set, so request bodies contain only
  properties the caller set.
- Generic code constrains models with `model.ServiceNowItem`
  (`store.BackedModel` + `serialization.Parsable` + `GetSysID()`).

## Consequences

- **Pros:** absent/zero distinction; minimal request bodies with no
  read-modify-write races on untouched fields; parity with msgraph-sdk-go and
  other Kiota SDKs; a single mechanical pattern for every model.
- **Cons:** accessors are more verbose than field access — reads are a
  `(value, error)` pair and a pointer dereference; models cannot be
  constructed as struct literals; contributors must learn the
  `internal/store` / `internal/serialization` helper idiom.
- User-facing explanation lives in the docs site's Core Concepts page; the
  contributor-facing summary is the "Backing-store models" design-decision
  page.
