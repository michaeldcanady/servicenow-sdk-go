# Specification: v2.0 Error Standardization & Casing Alignment

## Status: Draft

## Summary
As part of the v2.0 release, the ServiceNow SDK for Go must provide a public, idiomatic way to handle API errors. This involves creating a dedicated `errors` package, standardizing the brand casing from "ServiceNow" to "ServiceNow", and migrating core error structures out of `internal/`.

## Goals
1. **Public Error API**: Expose `ServiceNowError` and specific HTTP errors (e.g., `BadRequestError`) in a public `errors` package.
2. **Brand Consistency**: Standardize all "ServiceNow" (lowercase 'n') symbols to "ServiceNow" across the entire codebase.
3. **Consolidated Sentinels**: Move common sentinel errors (e.g., `ErrNilContext`) to the new `errors` package.

## Proposed Changes

### 1. New Package: `errors`
- Location: `/errors`
- Contents:
    - `ServiceNowError` (renamed from `ServiceNowError`)
    - `BadRequestError`, `UnauthorizedError`, `ForbiddenError`, `NotFoundError`, `TooManyRequestsError`, `ServerError`
    - `DefaultErrorMapping()`
    - Common sentinel errors.

### 2. Casing Standardization
- Rename `ServiceNowError` -> `ServiceNowError`
- Rename `NewServiceNowError` -> `NewServiceNowError`
- Audit and update any other occurrences of "ServiceNow" to "ServiceNow" in public and internal symbols.

### 3. Internal Refactoring
- Update `internal/service_now_error.go` to use the new package or move its logic.
- Update all RequestBuilders to use `errors.DefaultErrorMapping()`.

## Success Criteria
- [ ] Users can perform `errors.As(err, &errors.ServiceNowError{})`.
- [ ] No public symbols contain the "ServiceNow" (lowercase 'n') casing.
- [ ] All API modules consistently use the same error mapping logic.
- [ ] Codebase compiles and all tests pass with the new package structure.
