# Implementation Plan: v2.0 Comprehensive Refinement

## Objective
Finalize the ServiceNow SDK for Go v2.0 release with a production-grade architecture, idiomatic Go patterns, and robust error handling. This plan focus on architectural cleanup and explicit error mapping while deferring CI-breaking changes.

## Key Files & Context
- `go.mod`: Module path update (deferred).
- `internal/`: Consolidation of `internal/new` and unification of core utilities.
- `query2/`: Promotion to `query`.
- `*/request_builder.go`: Removal of `*2` suffixes and implementation of enhanced error mapping.
- `internal/service_now_error.go`: Centralized error handling.

## Proposed Changes

### Phase 1: Package & Internal Consolidation
1. **Directory Renaming**:
   - Rename `*-api` directories to single-word names (e.g., `table-api` -> `tableapi`).
   - Move `query2/` to `query/` (after deleting the old `query/` if it exists).
2. **Internal Unification**:
   - Move contents of `internal/new/mocking` to `internal/mocking`.
   - Move contents of `internal/new/testutils` to `internal/testutils`.
   - Remove `internal/new/` directory.
3. **Internal Import Updates**: Update all internal imports to reflect the new directory structure (keeping the current module path).

### Phase 2: API Surface Cleanup
1. **Suffix Removal**:
   - Identify and rename all types and methods ending in `2` (e.g., `TableRequestBuilder2` -> `TableRequestBuilder`).
   - Consolidate redundant builders.
2. **HEAD Method Implementation**:
   - Ensure all primary RequestBuilders (`Table`, `Attachment`, `Batch`, etc.) implement the `Head()` method using Kiota's `SendNoContent` pattern.

### Phase 3: Enhanced Error Mapping
We will implement explicit error mapping for ServiceNow's common HTTP status codes.

| HTTP Code | Error Type | Mapping Key | Description |
|-----------|------------|-------------|-------------|
| **400** | `BadRequestError` | `"400"` | Validation failed or malformed request. |
| **401** | `UnauthorizedError` | `"401"` | Missing or invalid credentials. |
| **403** | `ForbiddenError` | `"403"` | Authenticated but lacks permissions for the resource. |
| **404** | `NotFoundError` | `"404"` | The requested resource or record does not exist. |
| **429** | `TooManyRequestsError` | `"429"` | Rate limit exceeded. |
| **5XX** | `ServerError` | `"5XX"` | ServiceNow internal platform error. |
| **XXX** | `ServiceNowError` | `"XXX"` | Default fallback for other error codes. |

#### Implementation Steps:
1. Define specific error structs in `internal/service_now_error.go` that embed `ServiceNowError`.
2. Update `ErrorMappings` in all `RequestBuilders`:
   ```go
   errorMapping := abstractions.ErrorMappings{
       "400": internal.CreateBadRequestErrorFromDiscriminatorValue,
       "401": internal.CreateUnauthorizedErrorFromDiscriminatorValue,
       "403": internal.CreateForbiddenErrorFromDiscriminatorValue,
       "404": internal.CreateNotFoundErrorFromDiscriminatorValue,
       "429": internal.CreateTooManyRequestsErrorFromDiscriminatorValue,
       "5XX": internal.CreateServerErrorFromDiscriminatorValue,
       "XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
   }
   ```

### Phase 4: Pluggable Logging
1. Define a `Logger` interface in `internal`:
   ```go
   type Logger interface {
       Log(message string, args ...interface{})
   }
   ```
2. Add `Logger` to `ServiceNowServiceClientConfig`.
3. Implement a default "noop" logger and allow users to provide their own via `WithLogger` option.

### Deferred Phase: Major Version Module Path Update
*Note: This phase is deferred to avoid breaking the CI pipeline prematurely.*
1. **Module Path Update**: Update `go.mod` to `github.com/michaeldcanady/servicenow-sdk-go/v2`.
2. **Global Import Update**: Update all internal and example imports to include the `/v2` suffix.

## Verification & Testing
1. **Structural Integrity**: Run `go test ./...` to ensure package renames are correct and imports are consistent.
2. **Error Mapping Tests**:
   - Create unit tests for each error code using `httpmock` to return specific status codes.
   - Verify that the returned error can be type-asserted to the specific error type (e.g., `errors.As(err, &NotFoundError{})`).
3. **HEAD Tests**: Verify `Head()` calls return headers without bodies and handle errors correctly.
4. **Migration Validation**: Follow the updated migration guide to ensure a clean upgrade path.
