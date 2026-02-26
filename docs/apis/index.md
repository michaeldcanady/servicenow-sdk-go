# API reference overview

This section provides detailed reference for each supported ServiceNow API module in the SDK. Each submodule corresponds to a ServiceNow REST API of the same name, providing a Go-idiomatic way to interact with the platform.

Whether you're automating incident management, synchronizing user records, or handling file attachments, these modules provide the necessary tools for building robust ServiceNow integrations.

## Supported APIs

- [**Table API**](tables/index.md): Create, read, update, and delete (CRUD) operations for ServiceNow tables.
- [**Attachment API**](attachment/index.md): Manage files and raw file content attached to records.
- [**Batch API**](batch/index.md): Group multiple requests into a single efficient call.
- [**Fluent Query Builder (Preview)**](../user-guide/query-builder.md): A type-safe and expressive way to build ServiceNow queries.

## Design Philosophy

- **Idiomatic Go**: Uses standard Go patterns for error handling, context support, and naming.
- **Fluent & Standard Builders**: Offers both a fluent API for ease of use and standard builders for more manual control.
- **Internal consistency** within each submodule takes precedence, though cross-module consistency remains a broader goal.

Explore the subsections above for implementation details, code snippets, and specific endpoint documentation.
