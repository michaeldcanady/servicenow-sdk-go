---
name: new-api-module
description: Scaffold a new ServiceNow API module package (request builders, query parameters, request configurations, and tests) following this repo's established Kiota-based pattern. Use when adding support for a new ServiceNow API endpoint/table.
disable-model-invocation: true
---

# New API Module

Scaffolds a new `<name>api` package matching the conventions used by the existing
API modules (`tableapi`, `policyapi`, `caseapi`, `documentsapi`, etc.).

## When to use

The user wants to add SDK support for a new ServiceNow REST endpoint that doesn't
yet have a package (e.g. "add support for the Change Management API").

## Reference implementation

Use `tableapi/` as the canonical example — it has the fullest set of verbs (GET
collection, GET item, POST, PATCH, PUT, DELETE) and paging support. For a
simpler module, `policyapi/` shows a smaller single-resource shape.

## Steps

1. **Confirm the API surface** with the user: base path (e.g. `/api/now/v1/table`),
   supported HTTP verbs, path/query parameters, and the response body shape.

2. **Create the package directory** `<name>api/` at the repo root (sibling to
   `tableapi/`, `policyapi/`).

3. **For each resource + verb combination**, generate matching pairs of files,
   mirroring `tableapi/table_item_request_builder*.go`:
   - `<resource>_request_builder.go` — the request builder struct + constructors
     (`New<X>RequestBuilderInternal`, `NewDefault<X>RequestBuilder`, `New<X>RequestBuilder`)
     and one method per HTTP verb (`Get`, `Post`, `Patch`, `Put`, `Delete`).
   - `<resource>_request_builder_<verb>_query_parameters.go` — query parameter struct
     with `url:"..."` struct tags (uses `github.com/google/go-querystring/query`).
   - `<resource>_request_builder_<verb>_request_configuration.go` — wraps the query
     parameters + headers per Kiota's `RequestConfiguration` convention.
   - `<resource>_request_builder_test.go` — table-driven `testify` tests per verb,
     using `httpmock` to stub HTTP responses (see `tableapi/table_item_request_builder_test.go`).

4. **Follow these repo conventions exactly**:
   - Import `core "github.com/michaeldcanady/servicenow-sdk-go/core"` for
     `core.RequestBuilder`, `core.ServiceNowItemResponse[T]`, `core.DefaultErrorMapping()`.
   - Every method starts with the nil-guard pattern:
     ```go
     if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
         return nil, nil
     }
     if conversion.IsNil(rB.GetRequestAdapter()) {
         return nil, errors.New("request adapter is nil")
     }
     ```
   - Use `internal/conversion`, `internal/http`, `internal/model` helpers rather
     than reimplementing nil-checks or type assertions.
   - Generic parsable types follow `model.ServiceNowItem` constraint, e.g.
     `type <X>RequestBuilder[T model.ServiceNowItem] struct { ... }`.
   - URL templates go in an unexported `const` block using Kiota's URI-template
     syntax (`{+baseurl}/api/now/...{?query,params}`).
   - Add a `Readme.md` in the new package directory summarizing the endpoint,
     mirroring `tableapi/Readme.md`.

5. **Write unit tests** for every new file — this repo requires unit tests
   alongside every new feature (see `GEMINI.md` engineering standards). Cover
   nil-adapter errors, success paths, and error-mapping paths using `httpmock`.

6. **Optionally add integration coverage**: if the endpoint needs BDD-level
   verification, add a `.feature` file under `tests/integration/features/`
   plus step definitions in `tests/integration/`, following the pattern in
   `tests/integration/table_steps_test.go` (build-tagged `//go:build integration`,
   uses `godog`, `httpmock`, `godotenv`).

7. **Run validation** before handing back:
   ```
   gofmt -s -w <name>api/
   go build ./...
   go vet ./...
   golangci-lint run ./<name>api/...
   go test ./<name>api/...
   ```

## Notes

- Do not invent abstractions beyond what `core.RequestBuilder` already provides —
  new modules should compose existing `internal/*` helpers, not duplicate them.
- Keep each generated file focused on one type/verb, matching the granular
  file-per-concern layout already used across every existing `*api` package.
