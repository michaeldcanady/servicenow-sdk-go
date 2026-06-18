# Spec: v2.0 Final Polish & Production Readiness

**Status**: Draft | **Date**: 2026-06-17 | **Version**: 2.0.0-rc

## 1. Overview
As the ServiceNow SDK for Go approaches its v2.0 release, it is essential to ensure that the public API and internal architecture are consistent, idiomatic, and free of "experimental" artifacts. This specification covers the final refinements required to graduate the SDK from "v2-prep" to a stable production release.

## 2. Goals
- **Architectural Purity**: Remove all "2" suffixes from internal packages and types (excluding protocol-specific names like `oauth2`).
- **Dry Error Handling**: Centralize the repetitive `ErrorMappings` logic found across all RequestBuilders.
- **Robustness**: Ensure `ServicenowError` is safe and informative, even when underlying data is missing.
- **Debt Elimination**: Resolve remaining `TODO` markers and misleading documentation.

## 3. Requirements

### 3.1 Package & Type Refinement
- Rename the `internal/ast2` package to `internal/ast`.
- Ensure all factory methods (e.g., `New...RequestBuilderInternal`) follow a consistent naming and parameter pattern.
- Verify no public-facing types retain "2" suffixes (e.g., `BatchRequestBuilder2` tests).

### 3.2 Error Handling Centralization
- Create a centralized `DefaultErrorMapping()` function in the `internal` package that returns the standard `abstractions.ErrorMappings`.
- Update all `RequestBuilder` methods to use this central function instead of local map literals.

### 3.3 Error Model Safety
- Update `ServicenowError.Error()` in `internal/service_now_error.go` to safely handle cases where `GetError()` or its components are nil.
- Ensure the error message includes as much context as possible (Message, Detail, Status).

### 3.4 Technical Debt & Documentation
- Resolve all `TODO: add tests` markers by implementing the missing unit tests.
- Fix misleading comments in RequestBuilders (e.g., references to wrong URL templates or package names).
- Implement foundation for Multipart body helpers in `attachmentapi`.

## 4. Non-Goals
- Updating `go.mod` to include `/v2` (this is deferred until release day).
- Renaming the `oauth2` package or related types.

## 5. Success Criteria
- [ ] `go test ./...` passes with 0 failures.
- [ ] No `ast2` references remain in the codebase.
- [ ] `ErrorMappings` is centralized and used everywhere.
- [ ] `ServicenowError.Error()` does not panic on nil data.
- [ ] 0 `TODO` markers remaining in the core SDK packages.
