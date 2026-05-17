# Go Coding Conventions for ServiceNow SDK

## Core Principles
- **Idiomatic Go**: Follow standard Go practices (Effective Go).
- **Error Handling**: Use explicit error handling. Avoid `panic` unless truly exceptional.
    - Prefer `if err != nil { return fmt.Errorf("context: %w", err) }`.
    - Use `errors.Is` and `errors.As` for checking error types.
- **Formatting**: Always use `gofmt` (or `goimports`).
- **Naming**: 
    - Use `camelCase` for private symbols, `PascalCase` for public symbols.
    - Keep names concise but descriptive.
    - Receiver names should be short (1-3 letters, e.g., `rb` for `RequestBuilder`).

## Structural Patterns
- **Early Returns**: Reduce nesting by returning early on errors or special cases.
- **Composition over Inheritance**: Use embedding and interfaces to achieve polymorphism.
- **Interfaces**: Define interfaces where they are used (consumer side), keeping them small.

## Testing Standards
- **Testify**: Use `github.com/stretchr/testify/assert` and `github.com/stretchr/testify/require` for assertions.
- **Table-Driven Tests**: Use table-driven tests for multiple test cases of the same logic.
- **Mocking**: Use `testify/mock` and generated mocks when necessary.

## ServiceNow SDK Specifics
- **Kiota Alignment**: Follow the patterns in `internal/new/` for RequestBuilders and Parsables.
- **Core Types**: Use `core.ServiceNowError` and `core.ApiError` for API interactions.
