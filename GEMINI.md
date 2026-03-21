# ServiceNow SDK for Go - Project Context

This project is a Service-Now API client enabling Go programs to interact with Service-Now in a simple and uniform way. It is built on top of Microsoft Kiota abstractions and follows its request/response pipeline patterns.

## Project Overview

- **Purpose:** Provide a type-safe, idiomatic Go SDK for ServiceNow APIs.
- **Main Technologies:**
  - **Language:** Go 1.23+
  - **API Framework:** Microsoft Kiota (abstractions, serialization, http-go).
  - **Testing:** Standard `go test`, [Testify](https://github.com/stretchr/testify), [httpmock](https://github.com/jarcoal/httpmock), and [Godog](https://github.com/cucumber/godog) for BDD.
  - **Documentation:** MkDocs (run via Docker/Podman using `justfile`).

## Architecture and Structure

The SDK is organized by API modules, with a core client providing access to them.

- **`servicenow_service_client.go`**: Entry point for the SDK. It initializes the `RequestAdapter` and registers default serializers/deserializers.
- **`table-api/`**, **`attachment-api/`**, **`batch-api/`**, **`policy-api/`**: Module-specific implementations following the Kiota `RequestBuilder` pattern.
- **`internal/`**:
  - `kiota/`: Custom Kiota extensions and base RequestBuilder.
  - `http/`: HTTP-related utilities and constants.
  - `model/`: Shared data models and interfaces.
  - `serialization/`: Custom serialization logic.
- **`credentials/`**: Authentication providers and credential management (OAuth2, Basic Auth).

### RequestBuilder Pattern

The SDK extensively uses the Kiota `RequestBuilder` pattern. Each API endpoint is represented by a RequestBuilder that provides methods like `Get`, `Post`, `Put`, `Delete`, etc.
Endpoints with path parameters (e.g., `sys_id`) are accessed via methods like `ById(sysId)`.

## Key Commands

- **Build:** `go build ./...`
- **Test:** `go test ./...`
- **BDD Tests:** `go test ./tests/...` (requires `godog`)
- **Linting:** `golangci-lint run` (uses `.golangci.yml`)
- **Documentation:**
  - Build docs image: `just build-docs`
  - Serve docs locally: `just serve-docs`
  - Generate static site: `just generate-docs`

## Development Conventions

- **Kiota Compliance:** New API modules should follow the Kiota pattern:
  - Use `RequestBuilder` for defining endpoints.
  - Use `To[Method]RequestInformation` for request preparation.
  - Use `Parsable` interfaces for serialization.
- **Generics:** Use generics for type-safe API responses where applicable (see `TableRequestBuilder`).
- **Testing:**
  - Unit tests should use `httpmock` for mocking ServiceNow API responses.
  - BDD tests in `tests/` should use `godog` and Gherkin features in `tests/features/`.
- **Error Handling:** Use the custom error types and discriminator-based error creation (e.g., `CreateServiceNowErrorFromDiscriminatorValue`).
- **Formatting:** Code must be formatted with `gofmt` (enforced by `golangci-lint`).

## Documentation

Documentation source is in the `/docs` directory. It uses MkDocs with the Material theme. API documentation is often generated or manually updated based on ServiceNow's API specifications.
