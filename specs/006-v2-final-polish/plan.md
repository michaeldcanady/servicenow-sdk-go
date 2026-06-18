# Implementation Plan: v2.0 Final Polish

**Branch**: `feat/v2-final-polish` | **Date**: 2026-06-17 | **Spec**: [specs/006-v2-final-polish/spec.md](spec.md)

## Summary
Finalize the ServiceNow SDK for Go v2.0 release by cleaning up architectural artifacts, centralizing error handling logic, and resolving remaining technical debt.

## Technical Context
- **Language**: Go 1.25.0
- **Primary Deps**: Kiota Abstractions & HTTP
- **Testing**: Testify (Unit), Godog (Integration)

## Phases

### Phase 1: Internal Architectural Cleanup
- **Rename `internal/ast2`**: Move `internal/ast2` to `internal/ast` and update all imports.
- **Standardize Factories**: Audit all RequestBuilder packages (`tableapi`, `attachmentapi`, etc.) and ensure factory methods follow the `New[Name]RequestBuilderInternal` pattern.
- **Remove Suffixes**: Scan for any remaining `*2` types or test names and rename them.

### Phase 2: Error Handling Refinement
- **Centralize Mappings**:
    - Add `func DefaultErrorMapping() abstractions.ErrorMappings` to `internal/error.go` (or `internal/service_now_error.go`).
    - Replace all inline `ErrorMappings` literals in RequestBuilders with a call to this function.
- **Stringer Safety**:
    - Update `ServicenowError.Error()` to include nil checks for every level of the error model.
    - Improve the error string format: `ServiceNow Error: [Status] [Message] - [Detail]`.

### Phase 3: Debt & Feature Gaps
- **TODO Resolution**:
    - Address the 19 identified `TODO: add tests` markers.
    - Fix the `int64` type mismatch TODO in `attachmentapi`.
- **Documentation Audit**:
    - Fix copy-pasted comments in `tableapi` and others.
    - Update `Readme.md` files in sub-packages to reflect v2 patterns.
- **Multipart Helpers**:
    - Implement a `MultipartBody` helper in `internal` or `attachmentapi` to simplify file uploads.

### Phase 4: Validation
- **Unit Testing**: Run `go test ./...` to ensure no regressions.
- **Static Analysis**: Run `golangci-lint` to ensure code quality.

## Constitution Check
- [x] Principle I: Library-First - Keeps core logic lean and reusable.
- [x] Principle IV: Surgical Changes - Minimizes impact by centralizing repetitive logic.
- [x] Principle VI: Testing - Resolves all remaining testing gaps.

## Complexity Tracking
N/A
