# User guide

Task-oriented guides for the SDK's core features. If you're new, start with
[Getting Started](../getting-started.mdx) and then
[Core Concepts](../core-concepts.mdx) — the rules the guides below build on.

- **[Migrating from v1 to v2](migrate-v1-to-v2.mdx):** Coming from a 1.x
  release? Map every v1 construct to its v2 replacement.
- **[Support Policy](support-policy.md):** Which Go versions and SDK
  versions are supported, and for how long.
- **[Authentication](authentication/index.md):** Choose and configure a
  credential flow — Basic, OAuth2 authorization code, client credentials,
  JWT bearer, or ROPC.
- **[Table Operations](tables/index.mdx):** Create, read, update, and delete
  records in any ServiceNow table.
- **[Attachments](attachments/index.mdx):** Upload, download, list, and
  clean up files attached to records.
- **[Batch Operations](batch.mdx):** Combine multiple requests into a single
  HTTP call.
- **[Pagination](pagination.mdx):** Iterate over large result sets with page
  iterators.
- **More APIs:** Task guides for every other supported module — from
  [aggregate stats](apis/aggregation/aggregate-records.mdx) and
  [CMDB queries](apis/cmdb-instance/query-configuration-items.mdx) to
  [CSM cases](apis/case/read-cases.mdx) and the
  [CDM workflow](apis/cdm-changesets/track-changesets.mdx).
- **[Handling Errors](error-handling.mdx):** Match API and usage errors with
  `errors.As`/`errors.Is`.
- **[Configuring the Client](configuration.mdx):** Middleware, retries,
  logging, and transport options.
- **[Querying](query-builder.mdx):** Build encoded queries fluently
  (preview feature).
