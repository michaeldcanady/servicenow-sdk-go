---
title: 'ADR 020: Record-scoped resources are explicit-args helpers, not nested builders'
description: >-
  Because ServiceNow's attachment API is flat, a record's attachments are
  reached by an explicit-args helper (ForRecord) rather than a nested builder
  that would fake a URL that does not exist.
---

# ADR 020: Record-scoped resources are explicit-args helpers, not nested builders

## Status

Proposed. Accepted on merge; this record and its rule file
(`docs/rules/R-ATT-020.md`) ship together with the design-session work.

## Context

Callers naturally expect a record-scoped child builder —
`ByID(sysID).Attachments()` — because that is how nested resources read in
most SDKs, including Kiota-generated ones. But ServiceNow's attachment API is
**flat**: it lives at `/api/now/attachment` and targets a specific record via a
`table_name=...^table_sys_id=...` query filter, **not** at a nested path like
`/now/table/{table}/{sysID}/attachment`.

That mismatch creates two problems if we nest a builder:

- It would fake a URL structure that does not exist, breaking the URL-segment
  naming axis of ADR-008 (the accessor would imply a path that isn't there).
- It would couple `attachmentapi` to `tableapi` (and to the typed-table
  packages), when the two are independently useful.

Alternatives considered: a nested builder anyway (dishonest URL, unwanted
coupling); a generic "record-scoped resource" seam generalized now (speculative
— there is only one such resource today, which ADR-007 warns against); or an
explicit-args helper that builds the real flat filter (chosen).

## Decision

- Reach a record's attachments through an explicit-args helper:
  `attachmentapi.ForRecord(parent, tableName, sysID)`.
- Pass the parent **table name explicitly**; do not infer it or couple it from a
  path parameter.
- Keep the helper in `attachmentapi`, importing neither `tableapi` nor any
  typed-table package.
- Build the flat attachment query (`table_name=...^table_sys_id=...`), not a
  nested table path.
- **Attachments-only for now.** Do not generalize the record-scoped helper into
  a cross-module seam until a second record-scoped consumer actually exists
  (per ADR-007). When another flat record-scoped API appears, follow this
  pattern rather than pre-building a framework.

## Consequences

- **Pros:** the helper matches the real flat API instead of faking a nested URL;
  `attachmentapi` stays decoupled from `tableapi` and the typed tables; callers
  get an ergonomic record-scoped entry that reads correctly against the wire.
- **Cons:** `ByID(sysID).Attachments()` is not available, which surprises
  callers who expect nesting — the docs must explain *why* (the API is flat);
  generalizing the pattern waits for a second consumer, so the first instance is
  deliberately bespoke.
- **Rule for future record-scoped work:** model a record-scoped helper against
  the endpoint's real URL shape; if the API is flat, use explicit args; do not
  nest a builder to imply a path that does not exist, and do not generalize the
  seam until a second consumer justifies it.

See also: ADR 007 (no speculation), ADR 008 (naming axes), ADR 017 (hybrid
surface).
