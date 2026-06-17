# Data Model: Robust Testing Framework

## Mock Objects

### MockRequestAdapter
- **Interface**: Implements `github.com/microsoft/kiota-abstractions-go.RequestAdapter`
- **Behavior**: Stores `RequestInformation` for verification; returns `MockResponse` or error.

### MockResponseHandler
- **Interface**: Implements `github.com/microsoft/kiota-abstractions-go.ResponseHandler`
- **Behavior**: Intercepts raw HTTP responses and maps them to models or errors.

## Test Configuration

### Environment Variables
- `SN_INSTANCE`: The ServiceNow instance URL (e.g., `https://dev12345.service-now.com`).
- `SN_USERNAME`: Admin username for E2E tests.
- `SN_PASSWORD`: Admin password for E2E tests.
- `SN_TEST_TABLE`: A target table for CRUD operations (default: `incident`).

### Mock JSON Payloads
- **Location**: `testdata/*.json`
- **Format**: Standard JSON matching ServiceNow API responses.

## Observability
- **Failure Log Format**: Structured JSON containing test name, phase (unit/integration/e2e), error details, and timestamp.
