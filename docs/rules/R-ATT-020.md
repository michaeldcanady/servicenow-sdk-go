---
id: R-ATT-020
applies_to: ["attachmentapi/**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Record attachments use a flat, explicit-args helper

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- TODO(owner): define `parent` in `ForRecord(parent, tableName, sysID)`. List any typed-table packages to ban in Verify beyond tableapi. -->

MUST: Reach a record's attachments through the helper `attachmentapi.ForRecord(parent, tableName, sysID)`.
MUST: Pass the parent table name explicitly; never infer it from a path parameter.
MUST: Keep the helper in `attachmentapi`.
MUST NOT: Import `tableapi` or any typed-table package from `attachmentapi`.
MUST: Build the flat attachment query (`table_name=...^table_sys_id=...`) against `/now/attachment`.
MUST NOT: Nest a record-scoped builder under a table record; `/now/table/{t}/{id}/attachment` does not exist.
MUST NOT: Generalize the record-scoped helper into a cross-module seam until a second record-scoped consumer exists.
SHOULD: Add further record-scoped helpers only where the endpoint is flat, following this pattern.

## Verify

```bash
# grep exit 1 = no match = pass; exit 2 fails closed. Also requires ForRecord to exist.
{ grep -RInE '"github\.com/[^"]*/tableapi"' --include='*.go' attachmentapi/; [ $? -eq 1 ]; } \
  && grep -RqE 'func ForRecord\(' --include='*.go' attachmentapi/
```

## Rationale

ServiceNow's attachment API is flat (`/now/attachment` filtered by `table_name` and `table_sys_id`), so a nested record builder would fake a URL that does not exist. An explicit-args helper matches the real API and keeps `attachmentapi` decoupled from `tableapi` (see R-BUILD-007, R-SURF-019).
