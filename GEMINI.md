# ServiceNow SDK for Go - Project Context

This project is a ServiceNow API client for Go, designed to provide a simple and uniform way to interact with ServiceNow. It is built on a request/response pipeline architecture inspired by Microsoft Kiota, ensuring a consistent developer experience for multiple ServiceNow API modules.

## Project Overview
- **Purpose**: To provide a robust, type-safe, and scalable Go SDK for interacting with the ServiceNow platform.
- **Main Technologies**:
  - **Go (v1.24+)**: Core language.
  - **Microsoft Kiota**: Abstractions and HTTP library for building request/response flows.
  - **Gherkin/Godog**: Used for BDD-style integration testing.
  - **Testify**: Standard unit testing and mocking.
- **Architecture**:
  - **Modules**: Separate packages for `table-api`, `attachment-api`, and `batch-api`.
  - **Internal Core**: `internal/new/` contains shared Kiota-based request builders, serialization helpers, and response models.
  - **Client**: `ServiceNowClient` serves as the main entry point, providing access to the different API namespaces.

## Building and Running
The following commands are essential for developing and validating the SDK:

- **Testing**:
  - Run all tests: `go test ./...`
  - Run integration tests (requires environment variables): `go test ./tests/...`
- **Documentation**:
  - Build docs container: `just build-docs` (Requires `just` and `podman/docker`)
  - Serve docs locally: `just serve-docs`
  - Generate static site: `just generate-docs`
- **Cleaning**:
  - Remove doc artifacts: `just clean-docs`

## Core Mandates & Engineering Standards
- **Surgical Changes**: Always prefer minimal, targeted edits over large refactors unless explicitly requested.
- **Idiomatic Go**: Follow standard Go conventions (gofmt, linting, proper error handling).
- **Kiota Alignment**: Maintain consistency with Kiota's request builder and serialization patterns.
- **Testing Requirements**:
  - **Unit Tests**: Every new feature or bug fix must have accompanying unit tests in the same package (e.g., `*_test.go`).
  - **Integration Tests**: Significant API changes should be verified with Godog features in the `tests/` directory.
  - **Mocking**: Use `testify/mock` for internal mocks and `httpmock` for simulating ServiceNow API responses.
- **Error Handling**: Use the `core.ServiceNowError` and `core.ApiError` structures for API-related failures.

## Governance & Behavioral Rules
Gemini operates as an AI contributor and must strictly adhere to the following rules for all changes.

### 1. Semantic Versioning (SemVer 2.0.0)
Gemini must follow SemVer for all changes affecting code, schemas, workflows, or documentation.
- **MAJOR**: Backwards-incompatible changes.
- **MINOR**: Backwards-compatible feature additions.
- **PATCH**: Bug fixes, refactors, documentation, internal improvements.

**Required for every update:**
- Current and proposed version.
- Justification for the version increment.
- Categorized list of changes (MAJOR/MINOR/PATCH).
- Changelog-ready summary.
- **Never apply a version bump silently or infer it without explicit reasoning.**

### 2. Conventional Commits & Atomicity
All changes must be broken into **atomic, logically isolated commits**.
- **Atomic**: One conceptual change per commit (e.g., do not mix refactors with features).
- **Types**: `feat:`, `fix:`, `refactor:`, `docs:`, `perf:`, `chore:`, `test:`.
- **Breaking Changes**: Use `BREAKING CHANGE:` footer for MAJOR updates.

### 3. Commit-Ready Blocks
For every change, Gemini must output a "Commit-Ready Block" containing:
1. Proposed version bump and reason.
2. Code diff.
3. SemVer classification.
4. Conventional Commit message.
5. Changelog entry.
6. **Pause for user confirmation.**

## Project Structure
- `core/`: Base abstractions and common types.
- `credentials/`: Authentication and credential providers.
- `internal/new/`: Reusable Kiota-based components (RequestBuilder, PageIterator, etc.).
- `table-api/`, `attachment-api/`, `batch-api/`: Specific API implementations.
- `tests/`: BDD integration tests (Godog features).

## Development Workflows
### Research -> Strategy -> Execution
1. **Research**: Identify the relevant API package and Kiota abstractions. Verify existing patterns.
2. **Strategy**: Propose changes that align with the `RequestBuilder` and `Parsable` patterns. Determine SemVer impact.
3. **Execution**:
   - Update/Create models and request builders.
   - Add unit tests and run `go test ./...`.
   - Produce a **Commit-Ready Block** and wait for approval.
   - (Optional) Update integration tests in `tests/`.

## Key Files
- `servicenow_client.go`: Main entry point for the SDK.
- `now_request_builder.go`: Entry point for the "Now" API namespace.
- `internal/new/kiota_utils.go`: Utility functions for Kiota integration.
- `VERSION`: File containing the current project version.
- `CHANGELOG.md`: Record of all notable changes.
