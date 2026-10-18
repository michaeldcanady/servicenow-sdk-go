---
title: 'ADR 021: Best-effort error categorization with raw ServiceNow details attached'
description: >-
  ServiceNow errors are mapped into SDK error types on a best-effort basis, with
  the raw detail/message/status always attached for unwrap-based exact handling;
  409 and validation get first-class error types.
---

# ADR 021: Best-effort error categorization with raw ServiceNow details attached

## Status

Proposed. Accepted on merge; this record and its rule file
(`docs/rules/R-ERR-021.md`) ship together with the design-session work.

## Context

ADR 001 centralized error sentinels and a status-code-to-type mapping, and
ADR 006 made nil-receiver guards fail loud with a shared sentinel. What remains
is categorizing the errors ServiceNow itself returns. The error body has a
stable envelope — `{"error": {"detail", "message", "status"}}` — but the set of
situations it can describe is open-ended, so no fixed set of SDK error types can
exhaustively categorize it.

Callers want two things at once: a **category** they can switch on for the
common cases, and the **raw payload** so nothing is lost when they need exact
handling. A fixed, exhaustive taxonomy would drop detail; a pass-through-only
design would give no ergonomic category.

Options considered: exhaustive categorization (rejected — impossible to cover
open-ended ServiceNow errors); pure pass-through with no types (rejected — no
ergonomic `errors.As`/switch target); best-effort categorization with the raw
details always attached and a generic fallback (chosen). Feasibility was
confirmed by inspecting `core.DefaultErrorMapping()` and the existing typed
errors.

## Decision

- Categorize ServiceNow errors into SDK error types on a **best-effort** basis
  via `core.DefaultErrorMapping()`.
- **Always attach the raw ServiceNow `detail`, `message`, and `status`** to every
  categorized error so callers can unwrap to the original and inspect exact
  values.
- Add **`ConflictError`** (HTTP 409) and **`ValidationError`** to the default
  mapping.
- An **un-categorizable error degrades to a generic SDK error** that still
  carries the raw payload — never drop the raw details.
- Keep `UnauthorizedError.Is` matching `snerrors.ErrNotAuthenticated` (already in
  place).
- Map additional status codes incrementally as they are encountered.

## Consequences

- **Pros:** callers get a switchable category for the common cases and lossless
  raw access for everything else; new status codes can be added incrementally
  without a breaking change; no error detail is ever discarded.
- **Cons:** categorization is best-effort, so a category is a convenience, not a
  guarantee — callers doing exact handling must unwrap; the typed categories
  must not be over-trusted, which is a documentation responsibility.
- **Rule for future error work:** when adding an error category, attach the raw
  details and keep the generic fallback intact; a change that drops the raw
  payload or narrows unwrap access violates this decision. Prefer adding a
  mapping over expanding an existing type's meaning.

See also: ADR 001 (error standardization), ADR 006 (nil-receiver sentinel).
