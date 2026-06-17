# Tasks: v2.0 Refinement & Readiness

## Phase 1: Architectural Standardization

- [ ] **T101: Remove "2" Suffixes (Breaking)**
    - Rename `table-api/table_request_builder2.go` -> `table-api/table_request_builder.go`
    - Rename `query2` package -> `query`
    - Update all types (e.g., `TableRequestBuilder2` -> `TableRequestBuilder`)
    - Update all internal imports.
- [ ] **T102: Consolidate Internal Directory**
    - Move contents of `internal/new/` to `internal/`
    - Update imports across the repository.
- [ ] **T103: Standardize RequestBuilder Patterns**
    - Ensure all RequestBuilders use consistent `New<Name>RequestBuilder` and `New<Name>RequestBuilderInternal` patterns.

## Phase 2: Functional Readiness

- [ ] **T201: Implement Missing HEAD Methods**
    - Add `Head` method to `TableRequestBuilder`.
    - Add `Head` method to `AttachmentRequestBuilder`.
    - Add `Head` method to `BatchRequestBuilder`.
- [ ] **T202: Refine Error Mapping**
    - Update `errorMapping` in all RequestBuilders to handle 400, 401, 403, 404, 500 specifically.
- [ ] **T203: Attachment Multipart Helpers**
    - Implement `CreateMultipartBody` helper in `attachment-api`.
    - Provide example usage in documentation.

## Phase 3: Testing & Quality Assurance

- [ ] **T301: Resolve Testing TODOs**
    - Add tests for `attachment-api/attachment.go` (L486).
    - Add tests for `batch-api/batch_request_test.go` (L138).
    - Add tests for `internal/kiota_request_information_test.go` (L47).
    - (Complete all 19 identified TODOs).
- [ ] **T302: Standardize Mocking**
    - Audit `internal/mocking` and ensure it follows a consistent pattern for RequestAdapters and Parsables.

## Phase 4: Documentation & Environment

- [ ] **T401: Update Contributing Guide**
    - Update Go version to 1.25.0.
    - Add section on `release-please` and Conventional Commits.
- [ ] **T402: GoDoc Examples**
    - Add `Example` functions for `ServiceNowServiceClient`.
    - Add `Example` for `TableRequestBuilder.Get`.
    - Add `Example` for `AttachmentRequestBuilder.Upload`.
- [ ] **T403: Dependency Audit**
    - Run `go mod tidy`.
    - Update all `github.com/microsoft/kiota-*` dependencies to latest.
