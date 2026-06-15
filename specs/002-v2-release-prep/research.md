# Research: v2 Release Prep

## Deprecated Code Removal

### Findings
The codebase contains over 100 instances of the `// Deprecated:` tag. These are spread across multiple packages:
- `attachment-api/`
- `core/`
- `credentials/`
- `table-api/`
- `servicenow_client.go`
- `now_request_builder.go`

Most deprecated items have newer alternatives (e.g., `SendGet2`, `NewServiceNowServiceClient`).

### Decision
All items marked with `// Deprecated:` will be removed in v2.0. This is a breaking change, which is consistent with a major version release.

## Query System Selection: query vs query2

### Comparison
- **query**: Uses a traditional builder pattern with an internal AST. It is more verbose and less type-safe in its API surface.
- **query2**: Uses a fluent API based on field types (`StringField`, `NumberField`, etc.). It is more intuitive, type-safe, and immutable.

### Decision
Select **query2** as the standard query building system.

### Rationale
`query2` represents a modern redesign focused on usability. It aligns better with the SDK's goal of providing a developer-friendly experience.

### Alternatives Considered
- Keeping both: Rejected to avoid confusion and maintain a clean health state.
- Merging both: Rejected as `query2` is a complete redesign.

### Migration Strategy
1. Delete the `query/` directory.
2. Keep the `query2/` directory for now.
3. Update any internal references or documentation that points to `query/` to use `query2/` instead.
4. Note: `core.Query` is still used by internal request builders and will be maintained as the internal representation for now, but `query2` will be the recommended public API for building complex queries.
