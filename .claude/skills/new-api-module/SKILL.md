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

   ServiceNow's docs site (`docs.servicenow.com` / `www.servicenow.com/docs`) is a
   JS-rendered SPA — `WebFetch`/`WebSearch` will only return nav-menu content, never
   the actual endpoint/param/schema tables. Don't guess a spec from memory and scaffold
   against it. Instead, discover the real shape from a live instance:

   - Check whether `.env` (or `tests/e2e` env vars: `SN_INSTANCE`, `SN_USERNAME`,
     `SN_PASSWORD`) is available. **Never `Read` or `cat`/`grep` `.env` directly** —
     it's permission-blocked by design since it holds credentials. Instead source it
     inline within a single Bash call so secrets never enter the conversation:
     ```bash
     set -a; source .env; set +a
     curl -s -u "$SN_USERNAME:$SN_PASSWORD" -H "Accept: application/json" \
       "https://${SN_INSTANCE}.service-now.com/api/now/<path>"
     ```
   - Probe the base path and plausible sub-paths/verbs with `curl`. ServiceNow's error
     responses are informative: `"Requested URI does not represent any resource"` on
     *every* sub-path guess (not just a wrong one) usually means the owning plugin
     isn't installed/active on that instance, not that you picked the wrong path.
     Confirm by checking for the API's backing tables via the Table API
     (`/api/now/table/<table>?sysparm_limit=1` — `"Invalid table <x>"` means it doesn't
     exist) or, for custom/scripted APIs, `sys_ws_definition`.
   - Vary query parameters against a real, populated table (`incident` is always
     present on a dev instance) to discover accepted `sysparm_*` params and observe
     the exact response shape (nesting, arrays vs. objects, field naming) directly
     from a 200 response, rather than assuming it matches another API's shape.
   - If the target API/plugin isn't installed and can't be activated, **don't
     scaffold against a guessed schema**. Tell the user, and offer to probe for a
     different active-but-uncovered API to scaffold as a stand-in (compare against
     the existing `<name>api/` directories at the repo root to avoid duplicating one
     that's already covered).

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
         return nil, snerrors.ErrNilRequestBuilder
     }
     if conversion.IsNil(rB.GetRequestAdapter()) {
         return nil, snerrors.ErrNilRequestAdapter
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

8. **Verify against the live instance, not just unit tests**, when `.env` credentials
   are available. Unit tests only prove the mock round-trips through your own
   `Serialize`/`GetFieldDeserializers` — they can't catch a mismatch between your
   model and what ServiceNow actually sends. Write a throwaway `cmd/<name>check/main.go`
   that builds a real client (`servicenowsdkgo.NewServiceNowServiceClient` with
   `credentials.NewBasicProvider` + `WithInstance`) and calls the new builder against a
   populated table, then print/inspect every field. Run it with the same inline
   `set -a; source .env; set +a` pattern from step 1. Fix whatever the real response
   surfaces (e.g. kiota deserializes raw JSON object leaves as `*string`, not `string`
   — a naive `map[string]any` walk that only handles `string` will silently stringify
   pointer addresses instead of values), add a regression test for it, then **delete
   the throwaway `cmd/` directory** before handing back.

## Notes

- Do not invent abstractions beyond what `core.RequestBuilder` already provides —
  new modules should compose existing `internal/*` helpers, not duplicate them.
- Keep each generated file focused on one type/verb, matching the granular
  file-per-concern layout already used across every existing `*api` package.
