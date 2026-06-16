# Quickstart: Robust Testing Framework

## Prerequisites
- Go 1.25 or 1.26 (Officially supported n, n-1)
- `godog` CLI (optional, can be run via `go test`)

## Unit Testing (Test Table Pattern)
1. Create a `testdata/` directory in your package.
2. Add a mock response JSON file (e.g., `testdata/get_record.json`).
3. In your `*_test.go` file, use the Test Table pattern:
```go
func TestMyFeature(t *testing.T) {
    tests := []struct {
        name     string
        mockData string
        expected string
    }{
        {"Success", "success.json", "expected_value"},
        {"Error", "error.json", "error_message"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Use testutils.LoadJSON(tt.mockData)
            // Use mocking.MockRequestAdapter
            // Assert with testify
        })
    }
}
```
4. Run tests: `go test ./path/to/package`.

## Integration Testing (BDD)
1. Add a `.feature` file to `tests/integration/features/`.
2. Implement step definitions in `tests/integration/steps/`.
3. Run integration tests: `go test ./tests/integration/...`.

## E2E Testing (Live Instance)
1. Create a `.env` file in the root with `SN_INSTANCE`, `SN_USERNAME`, and `SN_PASSWORD`.
2. Run E2E tests (Manual Trigger): `go test ./tests/e2e/...`.

## Coverage Report
Run the unified test runner:
```bash
./scripts/test.sh --report
```
This will run all tests and open the HTML coverage report in your browser.
