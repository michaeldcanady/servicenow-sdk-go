# Research: Robust Testing Framework

## Decision 1: Kiota Mock Implementation
- **Decision**: Implement `MockRequestAdapter` and `MockResponseHandler` in `internal/new/mocking/`.
- **Rationale**: Provides a central place for mocks that all modules can use, ensuring consistency with Kiota's abstraction layer.
- **Alternatives considered**: Mocking at the HTTP client level with `httpmock`. While `httpmock` is great for external calls, mocking the Kiota abstractions allows testing the SDK's internal request building and response parsing logic more directly.

## Decision 2: Godog Structure
- **Decision**: Place feature files in `tests/integration/features/` and step definitions in `tests/integration/steps/`.
- **Rationale**: Keeps BDD tests separate from the library code, as recommended in `GEMINI.md`, while maintaining a clear organization.
- **Alternatives considered**: Putting `.feature` files alongside source code. Rejected to keep a cleaner separation of concerns and follow the project's established pattern.

## Decision 3: JSON Loading Helpers
- **Decision**: Use a simple helper function in `internal/new/testutils/` that reads files from a `testdata/` directory.
- **Rationale**: This is a standard Go pattern (`testdata/` directory is ignored by the toolchain) and allows for easy maintenance of mock payloads.
- **Alternatives considered**: Embedding JSON as strings in Go files. Rejected as it makes the code messy and harder to maintain for large payloads.

## Decision 4: Unified Coverage Reporting
- **Decision**: Use `go test -coverprofile=coverage.out ./...` followed by `go tool cover -html=coverage.out` in a centralized `scripts/test.sh`.
- **Rationale**: Uses native Go toolchain capabilities, which is idiomatic and requires no extra dependencies.
- **Alternatives considered**: Third-party coverage tools like `gocov`. Rejected to stick with the "Idiomatic Go" mandate and minimize dependencies.

## Decision 5: E2E Environment Configuration
- **Decision**: Use `.env` files (ignored by git) and environment variables for live instance credentials.
- **Rationale**: Secure and standard practice for CI/CD and local development.
- **Alternatives considered**: Hardcoding credentials or using a separate config file. Rejected for security and flexibility.

## Decision 6: Test Table Pattern
- **Decision**: Mandate Test Tables (subtests) for all unit tests.
- **Rationale**: Consistent organization, improved readability, and easier addition of new test cases. Aligns with standard Go best practices.
- **Alternatives considered**: Individual test functions per scenario. Rejected for lack of scalability and consistency.

## Decision 7: Go Version Support
- **Decision**: Target the two most recent major Go releases (n and n-1).
- **Rationale**: Aligns with the Go team's official support policy and ensures compatibility for the vast majority of users.
- **Alternatives considered**: Supporting older versions. Rejected to avoid maintenance overhead and leverage newer language features if needed.
