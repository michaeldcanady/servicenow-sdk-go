# API reference overview

This section provides detailed reference for each supported ServiceNow API module in the SDK. Each submodule corresponds to a ServiceNow REST API of the same name, providing a Go-idiomatic way to interact with the platform.

Whether you're automating incident management, synchronizing user records, or handling file attachments, these modules provide the necessary tools for building robust ServiceNow integrations.

## Supported APIs

### Core platform

- [**Table API**](tables/index.mdx): Create, read, update, and delete (CRUD) operations for ServiceNow tables.
- [**Attachment API**](attachment/index.mdx): Manage files and raw file content attached to records.
- [**Batch API**](batch/index.mdx): Group multiple requests into a single efficient call.
- [**Aggregation API**](stats/index.mdx): Aggregate statistics (count, sum, avg, min, max) over table records.
- [**Documents API**](documents/index.mdx): Document Management — create, explore, version, and stream documents.
- [**Activity Subscriptions API**](activity-subscriptions/index.mdx): Follow objects and manage activity-stream subscriptions.

### CMDB & CSDM

- [**CMDB Instance API**](cmdb-instance/index.mdx): CRUD on configuration items by CMDB class, plus CI relationships.
- [**Application Service API**](app-service/index.mdx): Register and query CSDM application services.

### Customer Service

- [**Case API**](case/index.mdx): Create, read, and update CSM cases.
- [**Account API**](account/index.mdx): Read customer account records.
- [**Appointment Booking API**](appointment-booking/index.mdx): Check availability and book service appointments.

### Configuration Data Management

- [**CDM Applications API**](cdm-applications/index.mdx): Deployables, shared components, uploads, and exports.
- [**CDM Changesets API**](cdm-changesets/index.mdx): Changeset activity, commit status, and impact analysis.
- [**CDM Editor API**](cdm-editor/index.mdx): Edit and validate configuration data nodes.
- [**Policy API**](policy/index.mdx): Policy input mappings.

## Design philosophy

- **Idiomatic Go**: Uses standard Go patterns for error handling, context support, and naming.
- **Fluent & Standard Builders**: Offers both a fluent API for ease of use and standard builders for more manual control — see [Core Concepts](../core-concepts.mdx#fluent-vs-standard) for when to use which.
- **Cross-module consistency**: Every module follows the same structural pattern (constructor triad, per-verb configurations, shared error mapping), so learning one module means learning them all. The pattern is documented in [Add a new API module](../contributing/add-api-module.md).

Explore the following subsections for implementation details, code snippets, and specific endpoint documentation.
