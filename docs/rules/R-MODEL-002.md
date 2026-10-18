---
id: R-MODEL-002
applies_to: ["**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: models, backing store, backingstore, absent vs zero, dirty
tracking, dirty set, GetX setX, ServiceNowItem, struct fields -->

MUST embed `core.BaseModel` and keep property data in a Kiota `BackingStore` instead of plain struct fields.
MUST expose each property as a `GetX() (T, error)` / `setX(T) error` accessor pair.
MUST return pointers from getters so `nil` ("not sent") stays distinct from a zero value.
MUST serialize only the backing store's dirty set so request bodies carry only the properties the caller set.
MUST constrain generic model code with `model.ServiceNowItem`.
MUST NOT declare exported model data as plain struct fields.
SHOULD build accessors from the store/serialization helper idiom rather than hand-rolling each one.

## Verify

```bash
# Fails (non-zero) if any model still uses encoding/json struct tags
# (a sign of a plain-field, non-backed model).
! grep -RInE 'json:"' --include='*.go' model/
```

## Rationale

A struct field cannot tell "the instance sent an empty value" from "the instance sent nothing," which caused absent/zero confusion and accidental whole-object overwrites; a backing store with change tracking fixes both (see ADR 002).
