# Testing Guide

Testing is a critical part of the ServiceNow SDK for Go. We aim for high coverage and reliable tests.

## Unit Tests

Unit tests are located alongside the code they test (e.g., `table_request_builder_test.go`). We use the standard `testing` package along with `testify` for assertions.

To run all unit tests:
```bash
go test ./...
```

### Mocking HTTP Requests

We use `github.com/jarcoal/httpmock` to mock ServiceNow API responses. This allows us to test request construction and response parsing without requiring a real ServiceNow instance.

```go
{% include-markdown '../snippets/testing.go' start='// [START testing_mocking]' end='// [END testing_mocking]' comments=false trailing-newlines=false dedent=true %}
```

## Integration Tests

Integration tests are located in the `integration/` directory. These tests run against a real ServiceNow instance and require credentials.

To run integration tests, you need to set environment variables:
- `SN_INSTANCE`
- `SN_USERNAME`
- `SN_PASSWORD`

```bash
go test -v ./integration/...
```

## Continuous Integration

We use GitHub Actions to run tests automatically on every Pull Request. The CI workflow runs:
- `golangci-lint`
- Unit tests on multiple Go versions
- Coverage reporting via Codecov
