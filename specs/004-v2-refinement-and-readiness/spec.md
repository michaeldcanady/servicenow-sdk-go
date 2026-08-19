# Specification: v2.0 Refinement & Readiness

## 1. Introduction
The v2.0 release of the ServiceNow SDK for Go is a major milestone. While the primary goal is removing deprecated code, this specification addresses the secondary goal of ensuring the resulting codebase is of production-grade quality, idiomatic, and fully featured.

## 2. Technical Rationale

### 2.1 API Suffixes (`*2`)
During the v1.x lifecycle, many new features were introduced with a `2` suffix (e.g., `TableRequestBuilder2`) to avoid breaking existing users. For v2.0, these suffixes must be removed.
- **Goal**: Clean public API surface.
- **Impact**: High (Breaking change).

### 2.2 Missing HTTP Methods
ServiceNow APIs often support `HEAD` for metadata discovery. Currently, several RequestBuilders (e.g., `TableRequestBuilder`) have `TODO` markers for `HEAD`.
- **Implementation**: Utilize Kiota's `ToHeadRequestInformation` and `Send` patterns.

### 2.3 Error Mapping
Currently, most RequestBuilders use a catch-all error mapping:
```go
errorMapping := abstractions.ErrorMappings{
    "XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
}
```
This is insufficient for production usage where users need to distinguish between Auth (401), Permissions (403), and Not Found (404) programmatically.

### 2.4 Internal Consolidation
The `internal/new` directory was created to stage Kiota-aligned versions of core utilities. Maintaining two internal structures increases cognitive load.
- **Action**: Merge `internal/new` into `internal`.

## 3. Implementation Details

### 3.1 Suffix Removal Strategy
A repository-wide search and replace will be required, followed by targeted file renames.
1. Identify all `*2` files and types.
2. Rename types in code.
3. Update imports.
4. Rename files.

### 3.2 Testing Strategy
Every `TODO: add tests` marker represents a potential bug or untested edge case.
- **Requirement**: Each resolved `TODO` must be accompanied by a unit test using the established `testify` and `mocking` patterns.

## 4. Documentation Strategy
The `CONTRIBUTING.md` is the first point of entry for new contributors. It must be updated to reflect the move to Go 1.25.0 and the automated CI/CD pipeline using `release-please`.
