# Spec: Internal Package Migration

## Objective
Audit the `internal` package and migrate components that should be publicly accessible to a new `core` package, ensuring the SDK is extensible and its error handling is accessible to users.

## Background
Currently, the `internal` package contains many fundamental types and interfaces (e.g., `RequestBuilder`, `PageIterator`, `ServiceNowError`) that are essential for using and extending the SDK. In Go, the `internal` directory prevents these components from being imported by code outside the module. Migrating these to a public package (e.g., `core`) will allow users to:
1. Handle specific API errors (e.g., type-asserting to `BadRequestError`).
2. Implement custom `RequestBuilder`s for unsupported endpoints.
3. Access core SDK abstractions without relying on internal-only paths.

## Audit Results

### Public Candidates (Move to `core` package)
These components are part of the SDK's public API surface or are essential for its extensibility.

| Component | Current File | Description | Benefit |
| :--- | :--- | :--- | :--- |
| `RequestBuilder` | `internal/base_request_builder.go` | Interface for all request builders. | Allows users to implement custom API builders. |
| `BaseRequestBuilder` | `internal/base_request_builder.go` | Base implementation of `RequestBuilder`. | Simplifies creation of custom request builders. |
| `PageIterator` | `internal/page_iterator.go` | Generic iterator for paginated collections. | Enables custom pagination logic and flow control. |
| `PageResult` | `internal/page_result.go` | Model for a single page of results. | Provides access to raw paginated data. |
| `ServiceNowCollectionResponse` | `internal/service_now_collection_response.go` | Base collection response model. | Standardizes collection response handling. |
| `ServiceNowItemResponse` | `internal/service_now_item_response.go` | Base item response model. | Standardizes single-item response handling. |
| `ServiceNowError` | `internal/service_now_error.go` | Base error type for ServiceNow APIs. | Enables base error handling for all API calls. |
| `BadRequestError` etc. | `internal/service_now_error.go` | Specific status code error types. | Enables programmatic error response (e.g., retries on 429). |
| `MainError`, `MainErrorable` | `internal/main_error.go` | Detailed error information models. | Provides access to the detailed "why" of an error. |
| `BackedModel` | `internal/backed_model.go` | Interface for models with a backing store. | Required for Kiota's dirty-tracking and delta support. |
| `BaseModel` | `internal/base_model.go` | Base implementation of `BackedModel`. | Standardizes model behavior across the SDK. |
| `Option` | `internal/option.go` | Generic functional option type. | Provides a consistent pattern for SDK configuration. |
| `ServiceNowRequestAdapter` | `internal/http/servicenow_request_adapter.go` | Custom Kiota Request Adapter. | Allows network-level customization (proxies, logging). |
| `ServiceNowClientConfig` | `internal/http/servicenow_client_config.go` | Configuration for the ServiceNow client. | Exposes client setup for advanced users. |

### Truly Internal (Remain in `internal` package)
These components are implementation details of the SDK's machinery and should not be exposed.

| Sub-package | Description |
| :--- | :--- |
| `internal/conversion` | Reflection-based type conversion and casting helpers. |
| `internal/serialization` | Kiota serialization and deserialization glue code. |
| `internal/store` | Backing store management for models. |
| `internal/ast` | Abstract Syntax Tree for OData filter expressions. |
| `internal/mocking` | Mock implementations for unit testing. |
| `internal/testutils` | Helper functions for testing. |

## Proposed Solution: `core` Package
Create a new package `github.com/michaeldcanady/servicenow-sdk-go/core` to house the public core abstractions.

### Why `core`?
- It is a standard name for base abstractions in Go SDKs.
- It clearly distinguishes fundamental types from feature-specific API packages (e.g., `tableapi`).

## Migration Plan

### Phase 1: Preparation
1. Create the `core/` directory in the project root.
2. Initialize it with a `doc.go` file explaining its purpose.

### Phase 2: Core Abstractions Migration
1. Move the following files from `internal/` to `core/`:
   - `backed_model.go`, `base_model.go`, `base_request_builder.go`
   - `main_error.go`, `option.go`, `page_iterator.go`, `page_result.go`
   - `service_now_collection_response.go`, `service_now_error.go`, `service_now_item_response.go`, `service_now_response.go`
2. Update the package name in these files to `core`.
3. Update internal imports within these files to point to their new locations or remaining `internal` sub-packages.

### Phase 3: HTTP & Client Migration
1. Create `core/http/` or move HTTP-related components directly into `core/`.
2. Move client configuration files from `internal/http/`:
   - `servicenow_request_adapter.go`, `servicenow_client_config.go`, `servicenow_client_option.go`

### Phase 4: Global Update
1. Update all imports in the repository from `.../internal` to `.../core` where applicable.
2. Update `tableapi`, `attachmentapi`, and other API packages to use the new `core` types.
3. Fix any type name collisions or visibility issues.

### Phase 5: Verification
1. Run all unit tests: `go test ./...`.
2. Run integration tests: `go test ./tests/...`.

## Migration Risks & Mitigation
- **Breaking Change**: This is a MAJOR breaking change.
- **Import Cycles**: Ensure `internal` sub-packages do NOT depend on `core`.
- **Refactoring Overhead**: Significant find-and-replace will be needed.

## Alternatives Considered
- **Keeping everything in `internal`**: Discourages extensibility and makes error handling difficult.
- **Moving to root package**: The root package would become too cluttered.
