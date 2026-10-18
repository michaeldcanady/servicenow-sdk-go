---
title: 'ADR 018: `.Table()` becomes a collection builder with a `.ByName()` escape hatch; `Table2` bridges v2'
description: >-
  The table entry becomes a collection builder exposing typed tables and a raw
  ByName escape hatch; because Go cannot overload, the new shape ships as
  Table2 in v2 and renames to Table at the next major.
---

# ADR 018: `.Table()` becomes a collection builder with a `.ByName()` escape hatch; `Table2` bridges v2

## Status

Proposed. Accepted on merge; this record and its rule file
(`docs/rules/R-SURF-018.md`) ship together with the design-session work.

## Context

The raw-table escape hatch today is `client.Now().Table(name string)`, which
hardcodes a `*TableRecord` return. The hybrid surface (ADR 017) needs the entry
point to do more: expose built-in typed tables, and expose a raw escape hatch
for an arbitrary table — i.e. a **collection builder**, not a single hardcoded
call.

Two constraints collide:

- **Path-segment naming** (the third axis of ADR-008) says an accessor should
  mirror its literal URL segment. The segment here is `table`, so the accessor
  wants to be named `Table`.
- **Go has no method overloading.** The new zero-arg collection builder and the
  legacy `Table(name string)` cannot both be named `Table` in v2.

Renaming the entry outright would break every existing caller the moment v2
ships, and the legacy form must keep working for the rest of v2 (it is removed
only at the next major). Alternatives considered: overloading (not possible in
Go), dropping the legacy form early (breaks callers mid-major), or naming the
new shape something unrelated to the URL segment (violates ADR-008's axis).

## Decision

- In v2, the table entry becomes a **collection builder** that exposes built-in
  typed tables (literal chaining) and a raw escape hatch **`.ByName(name)`**
  that reaches `/now/table/{name}`.
- Because the legacy `Table(name string)` must coexist until the next major, the
  new collection-builder method ships as **`Table2()`** in v2 — the naming rule
  appends a number on a collision — and is renamed to **`Table()`** at the next
  major once the legacy form is removed.
- Path segments map to method calls. Scoped APIs (caseapi, cmdbinstanceapi,
  appointmentbookingapi, …) keep semantic names with a documented URL mapping
  rather than strict pathing.
- The legacy `Now().Table(name string) *TableRecord` keeps working for all of v2.

## Consequences

- **Pros:** the entry mirrors the URL segment per ADR-008; the raw `.ByName`
  escape hatch is preserved; there is no overloading hack and no premature
  break.
- **Cons:** `Table2` is temporarily surprising and must be renamed at the next
  major (a scheduled follow-up, not optional); two entry shapes coexist across
  v2, so docs must be explicit about which to use; scoped APIs deliberately do
  not follow strict pathing, which needs the mapping documented.
- **Rule for future naming on this entry:** `Table2` is a bridge, not a name —
  rename it to `Table` when the legacy `Table(name string)` is removed at the
  next major, and do not introduce a third spelling in the meantime.

See also: ADR 008 (naming axes), ADR 017 (hybrid surface), ADR 019 (typed tables).
