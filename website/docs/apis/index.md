# API reference overview

This section provides detailed reference for each supported ServiceNow API module in the SDK. Each submodule corresponds to a ServiceNow REST API of the same name, providing a Go-idiomatic way to interact with the platform.

Whether you're automating incident management, synchronizing user records, or handling file attachments, these modules provide the necessary tools for building robust ServiceNow integrations.

## Supported APIs

### Core platform

- [**Table API**](tables/index.md): Create, read, update, and delete (CRUD) operations for ServiceNow tables.
- [**Attachment API**](attachment/index.md): Manage files and raw file content attached to records.
- [**Batch API**](batch/index.md): Group multiple requests into a single efficient call.
- [**Stats API**](stats/index.md): Aggregate statistics (count, sum, avg, min, max) over table records.
- [**Documents API**](documents/index.md): Document Management — create, explore, version, and stream documents.
- [**Activity Subscriptions API**](activity-subscriptions/index.md): Follow objects and manage activity-stream subscriptions.

### CMDB & CSDM

- [**CMDB Instance API**](cmdb-instance/index.md): CRUD on configuration items by CMDB class, plus CI relationships.
- [**Application Service API**](app-service/index.md): Register and query CSDM application services.

### Customer Service

- [**Case API**](case/index.md): Create, read, and update CSM cases.
- [**Account API**](account/index.md): Read customer account records.
- [**Appointment Booking API**](appointment-booking/index.md): Check availability and book service appointments.

### Configuration Data Management

- [**CDM Applications API**](cdm-applications/index.md): Deployables, shared components, uploads, and exports.
- [**CDM Changesets API**](cdm-changesets/index.md): Changeset activity, commit status, and impact analysis.
- [**CDM Editor API**](cdm-editor/index.md): Edit and validate configuration data nodes.
- [**Policy API**](policy/index.md): Policy input mappings.

## Design philosophy

- **Idiomatic Go**: Uses standard Go patterns for error handling, context support, and naming.
- **Fluent & Standard Builders**: Offers both a fluent API for ease of use and standard builders for more manual control.
- **Internal consistency** within each submodule takes precedence, though cross-module consistency remains a broader goal.

Explore the following subsections for implementation details, code snippets, and specific endpoint documentation.
