---
title: Add a new API module
description: >-
  The playbook for adding SDK support for a new ServiceNow REST API — from
  pinning down the real API surface to shipping the docs page.
---

# Add a new API module

This is the playbook for adding SDK support for a ServiceNow REST API that
doesn't have a package yet. Every `*api` package in the repository follows the
same shape; a new one should be indistinguishable in structure from the
existing ones.

Two reference implementations to keep open while you work:

- **`tableapi/`** — the canonical, fullest example: every verb, paging,
  generics.
- **`policyapi/`** — the minimal example: a small resource tree with few
  verbs.

## 1. Pin down the real API surface

Don't scaffold from memory or from marketing docs. Confirm, against a live
instance (a free [PDI](https://developer.servicenow.com/) works):

- the base path (e.g. `/api/now/documents`) and which sub-paths exist,
- which HTTP verbs each path accepts,
- which query parameters are honored,
- the exact response shape from a real 200 response — nesting, arrays vs.
  objects, field naming.

ServiceNow error responses are informative: `"Requested URI does not
represent any resource"` on *every* path guess usually means the owning
plugin isn't active on your instance, not that your path is wrong. If the
plugin can't be activated, stop — don't build against a guessed schema.

## 2. Write a blueprint

Add a short design doc under `docs/blueprints/` (see
`documents_api_blueprint.md` for the format): package name, base path, a
table of paths → request builders → verbs, and the models with their fields.
This is what reviewers check the implementation against.

## 3. Create the package

Create `<name>api/` at the repo root. For **each resource**, create a
request-builder file; for **each verb on that resource**, create its
companion files:

| File | Contents |
| ---- | -------- |
| `<resource>_request_builder.go` | Builder struct, constructor triad, one method per verb + `To<Verb>RequestInformation` |
| `<resource>_request_builder_<verb>_query_parameters.go` | Query-parameter struct with `url:"..."` tags |
| `<resource>_request_builder_<verb>_request_configuration.go` | Wraps headers + query parameters |
| `<resource>_request_builder_test.go` | Table-driven tests per verb |

One type or verb per file — never inline query parameters or configurations
into the builder file.

### The constructor triad

Every request builder exposes exactly three constructors:

```go
// Raw entry point: path parameters + adapter.
func NewXRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *XRequestBuilder

// From a raw URL, with the default parsable factory.
func NewDefaultXRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *XRequestBuilder

// From a raw URL, with a caller-supplied parsable factory.
func NewXRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter, factory serialization.ParsableFactory) *XRequestBuilder
```

### The nil-guard prologue

Every verb method starts with the same two guards, returning the **shared**
sentinels (never a fresh `errors.New`):

```go
if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
    return nil, snerrors.ErrNilRequestBuilder
}
if conversion.IsNil(rB.GetRequestAdapter()) {
    return nil, snerrors.ErrNilRequestAdapter
}
```

See [Error handling design](design-error-handling.mdx) for why, and which
sentinel lives where.

### Other conventions

- URL templates are unexported `const`s in Kiota URI-template syntax:
  `{+baseurl}/api/now/...{?sysparm_query,sysparm_limit}`.
- Generic models are constrained by `model.ServiceNowItem`:
  `type XRequestBuilder[T model.ServiceNowItem] struct { ... }`.
- Error mapping always goes through `core.DefaultErrorMapping()` — never a
  bespoke error struct.
- Compose `internal/conversion`, `internal/http`, and `internal/serialization`
  helpers; don't reimplement nil checks, header enums, or serializer plumbing.
- Models embed `core.BaseModel` and read/write through backing-store
  accessors, not struct fields — see
  [Backing-store models](design-backed-models.md).
- Add a `Readme.md` in the package summarizing the endpoint, mirroring
  `tableapi/Readme.md`.

## 4. Test it

- **Unit tests** are required for every exported type and method: co-located,
  table-driven with `testify`, HTTP stubbed with `httpmock`. Cover the
  nil-guard errors, the success path, and at least one mapped API error.
- **Integration tests** (optional but encouraged): a Gherkin `.feature` file
  under `tests/integration/features/` plus step definitions, build-tagged
  `//go:build integration`. See the [Testing guide](testing.md).
- **Verify against a live instance** before calling it done. Unit tests only
  prove your mocks round-trip through your own serializer — they can't catch
  a mismatch with what ServiceNow actually sends. Point a scratch program at
  a populated table, inspect every field, fix what the real response
  surfaces, and add a regression test for it.

## 5. Document it

A new module ships with its docs page:

1. Add `website/docs/apis/<name>/index.md` with an overview and the available
   operations, and register it in `website/sidebars.ts`.
2. Put every code sample in `website/snippets/<name>.go` behind
   `// [START region]` / `// [END region]` markers and render it with the
   `GoSnippet` component — CI compiles all snippets, so samples can't rot.
3. Update the Supported APIs table in the root `Readme.md`.

## 6. Validate

```bash
gofmt -s -w <name>api/
go build ./...
go vet ./...
golangci-lint run ./<name>api/...
go test ./<name>api/...
go vet -tags snippets ./website/snippets/
```

Then follow the [contribution workflow](index.mdx#submitting-your-changes) —
branch naming, Conventional Commits (`feat(<name>api): add <API> support`),
and the PR checklist.
