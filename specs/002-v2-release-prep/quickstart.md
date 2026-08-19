# Migration Validation Guide: v1 to v2

## Prerequisites
- A project using `servicenow-sdk-go` v1.x.

## Validation Scenarios

### 1. Client Initialization
- **Action**: Update `NewServiceNowClient` to `NewServiceNowServiceClient`.
- **Expected Outcome**: Project compiles and successfully authenticates.

### 2. Query Building
- **Action**: Replace `query` package imports with `query2`.
- **Expected Outcome**: Query strings generated match expected ServiceNow encoded query format.

### 3. Request Execution
- **Action**: Update `Get()` calls to use the new `Get()` (which was previously `Get2()` or `Get3()`).
- **Expected Outcome**: API responses are correctly parsed into the new model types (e.g., `TableRecord`).

## Regression Testing
- Run `go test ./...` and ensure 100% pass rate.
- Run integration tests in `tests/` to verify end-to-end functionality with a live ServiceNow instance.
