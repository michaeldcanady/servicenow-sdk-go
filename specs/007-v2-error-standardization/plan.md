# Implementation Plan: v2.0 Error Standardization & Casing Alignment

**Branch**: `v2-error-standardization` | **Spec**: [specs/007-v2-error-standardization/spec.md](spec.md)

## Summary
Migrate core error structures to a public `errors` package and standardize the brand casing of "ServiceNow" across the SDK to ensure consistency and a professional developer experience for the v2.0 release.

## Technical Context
- **Language**: Go 1.25+
- **Major Dependency**: Microsoft Kiota Abstractions
- **Scope**: Repository-wide breaking change.

## Phases

### Phase 1: Initialize Public `errors` Package
- [ ] Create `/errors` directory.
- [ ] Implement `errors.go` with sentinel errors (`ErrNilContext`, etc.).
- [ ] Define the `ServiceNowError` interface and concrete types in `errors/service_now_error.go`.
- [ ] Move `DefaultErrorMapping` to the `errors` package.

### Phase 2: Casing Standardization
- [ ] Perform a global search for `ServiceNow` and replace with `ServiceNow` in all Go files (names, comments, methods).
- [ ] Ensure all constructors (e.g., `NewServiceNowServiceClient`) follow the correct casing.

### Phase 3: SDK-wide Migration
- [ ] Update all `RequestBuilder` implementations to import and use the new `errors` package.
- [ ] Update `internal/page_iterator.go` and other internal utilities to use `errors.DefaultErrorMapping()`.
- [ ] Remove deprecated or redundant error mappings in `internal/`.

### Phase 4: Validation & Cleanup
- [ ] Update all unit tests to reflect the new package structure.
- [ ] Ensure `errors.As` works correctly with the new public types.
- [ ] Run `go test ./...` and `golangci-lint run ./...`.

## Success Criteria
- No `ServiceNow` (lowercase 'n') casing remains in the public API.
- Users can import `github.com/michaeldcanady/servicenow-sdk-go/errors` and handle API failures type-safely.
- All tests pass.
