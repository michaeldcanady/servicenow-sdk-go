# Testing guide

The SDK maintains high standards for code quality by requiring comprehensive tests for
every new feature and bug fix. This guide explains how to write and run tests
within the SDK.

## Test categories

The SDK uses two main categories of tests: unit tests and integration tests.

### Unit tests

Unit tests verify individual components in isolation. They're located in the
same directory as the code they test and use the `_test.go` suffix. These tests
must be fast and not require external network access.

To run all unit tests:

```bash
go test ./...
```

### Integration tests

Integration tests verify the SDK's behavior with a real ServiceNow instance.
They're located in the `integration/` directory. These tests are essential for
validating that the SDK correctly handles actual API responses.

To run integration tests:

1.  **Configure environment variables** in a `.env` file as described in the
    [setup guide](setup.md).
2.  **Execute the tests:**
    ```bash
    go test -v ./integration/...
    ```

## Mocking HTTP requests

The project uses the `httpmock` library to simulate ServiceNow API responses in
unit tests. This lets you verify request construction and response parsing
without a live instance.

```go
func TestExample(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    // Register a mock responder
    httpmock.RegisterResponder("GET", "https://instance.service-now.com/api/now/table/incident",
        httpmock.NewStringResponder(200, `{"result": []}`))

    // Execute your test logic...
}
```

## Writing effective tests

When contributing tests, follow these best practices:

- **Test for success and failure:** Make sure you test both happy paths and
  many error conditions, for example, 401 Unauthorized, 404 Not Found.
- **Use meaningful data:** Avoid using "foo" or "bar." Use data that resembles
  real ServiceNow records, for example, "INC0010001."
- **Verify request details:** In unit tests, check that the correct headers,
  query parameters, and body reach the server.
- **Check for regressions:** If you're fixing a bug, add a test case that
  reproduces the bug to make sure it doesn't return.

## Coverage reporting

The project uses Codecov to track test coverage. You can generate a local coverage report
using the standard Go tools:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Aim to maintain or increase the current coverage level with your contributions.
