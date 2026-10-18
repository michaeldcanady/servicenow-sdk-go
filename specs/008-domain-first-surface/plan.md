# Plan: Domain-first surface (design-session decisions)

Status: planning artifact. Captures the consolidated design tree from the
"users develop without knowing ServiceNow's API architecture" design session,
and the ADRs it produced. The ADRs (017–021) are the primary source; this file
is the working plan that ties them together.

## North star

A ServiceNow developer should be able to say *"I want to do X"* and reach for
the SDK, not for ServiceNow's REST documentation. The SDK presents a
domain-oriented surface for the common cases and offers honest escape hatches
to the raw ServiceNow surface for everything else.

This is **hybrid**, not thin-wrapper, not fully-opinionated: domain APIs for
common workflows, explicit raw access when the domain view doesn't cover you.

## The design tree

### Surface shape

- **Entry point:** `client.Now()` returns a collection builder. The raw-table
  escape hatch is `.ByName(name)` (reaches `/now/table/{name}`). Typed access
  is exposed as built-in typed tables that chain literally.
- **Coexistence:** the old `client.Now().Table(name string)` is renamed for
  v2 (becomes a collection builder), and a `Table2()` temporary alias lets new
  and old call shapes coexist until the next major, when the old form is
  removed and `Table2` is renamed back to `Table`.
- **Path-segment naming:** a builder method name maps to a URL path segment;
  on a name collision, append a number (`Table2`).

### Typed tables

- **Built-in typed tables** (Incident, Case first, growing) ship with the SDK;
  consumers cannot modify them.
- **User-defined typed tables** wrap the exported collection builder via a
  wrapper *variable* (`tables := mycool.New(client.Now().Table2())`), not a
  literal chain — Go's static typing requires a named type to attach methods.
  Built-ins chain literally; user tables are reached through the wrapper var.
- **Field descriptors** (key + typed accessor) drive both built-in and
  user-defined typed fields; typed query uses `inc.Priority().Eq(IncidentPriority.High)`.
- **Exported foundation:** the `internal/store` and `internal/serialization`
  helpers are exported (or a convenience base type wraps them) so a user table
  is as capable as a built-in — **full parity**, no second-class citizen.
- **Uniformity:** typed tables stay uniform. Domain *actions* (e.g. Case
  add-comment) live in scoped API packages, not on the table type.
- **References** are a typed wrapper (sys_id + display + link); they do **not**
  auto-navigate to the referenced record.

### Query

Two coexisting styles: the existing fluent builder (`query.Where(...)`) and
typed domain fields (`inc.Priority().Eq(...)`).

### Errors

- Best-effort categorization of ServiceNow errors into SDK error types; the raw
  ServiceNow details (`detail`/`message`/`status`) stay attached so callers can
  unwrap to the original.
- Fallback: an un-categorizable error degrades to a generic SDK error that
  still carries the raw payload.
- Add `ConflictError` (409) and `ValidationError` to
  `core.DefaultErrorMapping()`; `UnauthorizedError.Is` already matches
  `snerrors.ErrNotAuthenticated`.

### Record-scoped resources (attachments)

- The ServiceNow attachment API is **flat** (`/api/now/attachment`, reached by
  a `table_name^table_sys_id` filter), not nested under a table path. So
  record-scoped helpers take explicit arguments:
  `attachmentapi.ForRecord(parent, tableName, sysID)`.
- Attachments-only for now; the seam is **not** generalized until a second
  record-scoped consumer exists (consistent with ADR-007).
- The helper lives in `attachmentapi`, importing neither `tableapi` nor the
  typed-table packages.

### Shipping strategy

New implementation ships **alongside** the old for the rest of v2; the old is
removed at the next major version (with `Table2` renamed to `Table`).

## ADRs produced by this session

| ADR | Rule ID | Decision |
| --- | ------- | -------- |
| 017 | SURF-017 | Hybrid domain-first surface — the north star |
| 018 | SURF-018 | `.Table()` becomes a collection builder; `.ByName()` escape; `Table2` coexistence |
| 019 | TBL-019 | Typed tables: built-in + user-defined wrappers, field descriptors, full parity, uniformity |
| 020 | ATT-020 | Record-scoped resources are explicit helpers, not nested builders |
| 021 | ERR-021 | Best-effort error categorization with raw details attached; `ConflictError` + `ValidationError` |

## Decision log (the grilling that produced this)

- Q1 hybrid; Q3 core workflows first (CRUD + query + pagination + attachments);
  Q4 two query styles; Q5/Q7 best-effort errors with raw details; Q11 add
  Conflict/Validation; Q6/Q14/Q18/Q19/Q23 export store/serialization + base
  type, built-in typed tables, full parity; Q9/Q15/Q16/Q29 wrapper-var
  chaining; Q12/Q13/Q17/Q24 collection builder + `ByName` + `Table2`; Q20
  coexist-then-remove; Q21/Q22 typed tables uniform, references don't navigate;
  Q25–Q28 attachments as explicit-args helpers.

## Convention change

ADRs are now maintained in two formats (see the ADR index's "Proposing a new
ADR" section):

- **Human narrative ADR** — `website/docs/contributing/adrs/NNN-slug.md`
  (unchanged repo convention).
- **Agent rule file** — `docs/rules/R-{AREA}-{NNN}.md`, RFC-2119 MUST lines
  with a runnable `## Verify` command and a ≤3-sentence `## Rationale`.

Both carry the same ADR number. The `NNN` is the shared key; the `AREA` token
scopes the rule for `applies_to` filtering.
