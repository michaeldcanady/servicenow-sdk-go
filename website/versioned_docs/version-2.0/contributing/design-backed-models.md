---
title: Why aren't models plain structs?
description: >-
  Models keep their data in a change-tracking backing store instead of struct
  fields — so "absent" and "zero" stay distinguishable and Patch bodies stay
  minimal.
---

# Why aren't models plain structs?

*Primary source: [ADR 002 — Backing-store-backed models](https://github.com/michaeldcanady/servicenow-sdk-go/blob/main/docs/adr/002-backing-store-models.md).*

## The decision

Models aren't plain structs. Every model embeds `core.BaseModel` and stores
its data in a Kiota `BackingStore` — a key/value store with change tracking —
rather than in struct fields. Accessors are generated from the helpers in
`internal/store` (`DefaultBackedModelAccessorFunc` /
`DefaultBackedModelMutatorFunc`), so every property is a `GetX() (T, error)` /
`setX(T) error` pair around the store.

This replaced plain struct fields during the v2 rework
([PR #474](https://github.com/michaeldcanady/servicenow-sdk-go/pull/474)).

## Why

- **Absent ≠ zero.** ServiceNow omits fields freely (and `sysparm_fields`
  makes it routine). A struct field can't distinguish "the instance sent an
  empty string" from "the instance didn't send this field"; a store lookup
  can, which is why getters return pointers and errors.
- **Dirty tracking.** The store records which properties changed, so a
  `Put`/`Patch` body serializes only the fields the caller actually set —
  not a full struct with zero values that would overwrite server data.
- **Kiota parity.** The SDK deliberately matches the conventions of
  Kiota-generated SDKs (msgraph-sdk-go and friends). Backed models are how
  those SDKs behave; deviating would make this the one Kiota-style SDK whose
  models act differently.

## What to do in new code

- Embed `core.BaseModel`; never add plain data fields to a model.
- Build each accessor pair on `internal/store` helpers, and each
  `Serialize`/`GetFieldDeserializers` on the `internal/serialization`
  generators — no hand-rolled property plumbing.
- Generic code should constrain models with `model.ServiceNowItem`
  (`store.BackedModel` + `serialization.Parsable` + `GetSysID()`).

The user-facing consequence of this design is the pointer-and-error getter
pattern explained in [Core Concepts](../core-concepts.mdx#4-backed-models-and-the-pointer-getter-pattern).
