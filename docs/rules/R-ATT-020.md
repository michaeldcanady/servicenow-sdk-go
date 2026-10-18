---
id: R-ATT-020
applies_to: ["attachmentapi/**/*.go", "**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

<!-- Keywords: attachments, ForRecord, record scoped, flat api, table_name,
table_sys_id, nested builder, explicit args, parent record -->

MUST reach a record's attachments through an explicit-args helper `attachmentapi.ForRecord(parent, tableName, sysID)`.
MUST pass the parent table name explicitly — do not infer it or couple it from a path parameter.
MUST keep the helper in `attachmentapi`, importing neither `tableapi` nor any typed-table package.
MUST build the flat attachment query (`table_name=...^table_sys_id=...`) rather than a nested table path.
MUST NOT nest a record-scoped builder under a table record as if the API were `/now/table/{t}/{id}/attachment` — the real API is flat.
MUST NOT generalize the record-scoped helper into a cross-module seam until a second record-scoped consumer exists.
SHOULD add further record-scoped helpers as flat-API entry points when their endpoints are actually flat, following this pattern.

## Verify

```bash
# Fails (non-zero) if attachmentapi starts importing tableapi (coupling).
! grep -RInE '"github.com/.*/tableapi"' attachmentapi/
```

## Rationale

ServiceNow's attachment API is flat (`/now/attachment` filtered by `table_name^table_sys_id`), so a nested record builder would fake a URL that doesn't exist; an explicit-args `ForRecord` helper matches the real API and keeps `attachmentapi` decoupled from `tableapi` (see ADRs 007, 008, 020).
