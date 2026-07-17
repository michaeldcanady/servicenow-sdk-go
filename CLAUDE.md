# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A Go SDK (`github.com/michaeldcanady/servicenow-sdk-go`) that exposes ServiceNow's REST APIs as a
fluent, typed client. It is built directly on Microsoft's Kiota abstractions
(`kiota-abstractions-go`, `kiota-http-go`, `kiota-serialization-*-go`) rather than a generated
Kiota client — the request-builder/parsable/backing-store pattern is hand-written to match Kiota's
conventions so the SDK "feels" like other Kiota-generated SDKs (e.g. msgraph-sdk-go).

## Commands

```bash
go build ./...                      # build everything
go test ./...                       # unit tests only (integration/e2e are build-tag gated, see below)
go test ./tableapi/... -run TestXxx -v   # single package / single test
golangci-lint run ./...             # lint (config: .golangci.yml)
gofmt -s -w .                       # format (also: `just fmt`)
just build | just lint | just fmt   # justfile wraps the above
./scripts/test.sh --report          # unit tests + HTML coverage report (coverage.html)
./scripts/test.sh --md-report       # unit tests + Markdown test report
```

Integration and e2e tests are excluded from a plain `go test ./...` by build tags:
```bash
go test -tags integration ./tests/integration/...   # godog/BDD, uses httpmock — no live instance needed
go test -tags e2e ./tests/e2e/...                   # hits a real ServiceNow instance via .env credentials
```

Docs (mkdocs, in `docs/`) are built via `just build-docs` / `just serve-docs` (requires podman/docker).

## Architecture

### Client entry point and fluent chaining

`NewServiceNowServiceClient(opts...)` (`servicenow_service_client.go`) builds a
`ServiceNowServiceClient` wrapping a `core.RequestBuilder` rooted at the instance base URL.
Every API surface hangs off it via chained builder methods, e.g.:

```go
client.Now().Table("incident").ByID(sysID).Get(ctx, nil)
client.Now().Attachment().File().Post(ctx, media, nil)
client.CustomerService()...
client.Cdm()...
client.AppointmentBooking()...
```

Each `<Name>()` method on a builder clones the parent's path parameters (`maps.Clone`), adds/overrides
one entry, and constructs the child builder with the same `RequestAdapter`. This is how the URL
template's path parameters accumulate as you chain deeper (see `now_request_builder.go` for the
`Now()` namespace's own set of children).

### `core/` — shared request/response machinery

- `core.RequestBuilder` / `core.BaseRequestBuilder` — the base every `*api` request builder embeds.
- `core.BaseModel` — the base every model embeds; models are **backing-store-backed**, not plain
  structs — properties are read/written through `internal/store` accessor/mutator funcs against a
  `kiotaStore.BackingStore`, not struct fields (this changed from plain fields to store-backed
  accessors during the v2 rework — see PR #474 / commit `95ee5a9`).
- `core.ServiceNowItemResponse[T]` / `core.ServiceNowCollectionResponse[T]` — generic response
  envelopes; `T` is constrained by `internal/model.ServiceNowItem` (`store.BackedModel` +
  `serialization.Parsable` + `GetSysID()`).
- `core.DefaultErrorMapping()` + `internal.GetErrorRegistryInstance()` — every request builder's
  `Send`/`SendPrimitive` call passes `core.DefaultErrorMapping()`, which maps HTTP status codes
  (`"400"`, `"401"`, `"404"`, `"5XX"`, `"XXX"`, ...) to `core.ServiceNowError` discriminator
  factories. New API-error handling should go through this, never a bespoke error struct.
  `internal.GetErrorRegistryInstance()` is a singleton `Dictionary` used to register/look up per-module
  error mappings.
- `core.PageIterator` / `core/paging.go` — cursor-based pagination via the response's `Link` headers
  (`core.ParseHeaders`). Only paginated endpoints (e.g. `tableapi`) use this.

### `internal/` — implementation helpers, never imported by consumers

- `internal/conversion` — `IsNil` (reflect-safe nil check used everywhere as the standard nil-guard),
  `As2`, `CollectionApply`/`CastCollection`, string→primitive converters.
- `internal/store` — `DefaultBackedModelAccessorFunc[S,T]` / `DefaultBackedModelMutatorFunc[S,T]` and
  the lower-level `DefaultStoreAccessorFunc`/`DefaultStoreMutatorFunc` — this is *the* mechanism every
  model's `GetX()`/`setX()` pair is built on.
- `internal/serialization` — `Serialize`/`SerializeXFunc`/`DeserializeXFunc` helper generators used in
  every model's `Serialize()` / `GetFieldDeserializers()`.
- `internal/http` — `RequestHeader`/`HTTPHeader`/`ContentType` enums (`.String()` methods), default
  middleware/client construction.
- `internal/ast` + `query/` — the fluent query-condition builder (`query.Field(...).Equals(...)`, etc.)
  compiles to an `internal/ast` expression tree with an `Operator` enum and a `StringerVisitor` that
  renders ServiceNow's `sysparm_query` encoded-query syntax.
- `internal/oauth2` — token/credential plumbing used by `credentials/`.
- `internal/kiota_request_information.go` / `ConfigureRequestInformation` — the one place that applies
  `Headers`/`Options`/`QueryParameters` from a per-verb `*RequestConfiguration` onto a
  `RequestInformation`. Request builders that also carry a request body do that step manually via
  `SetContentFromParsable` (see the `*_request_builder.go` `ToXRequestInformation` methods) because
  body content isn't part of the generic `RequestConfiguration[T]` shape.

### API modules (`accountapi/`, `tableapi/`, `attachmentapi/`, `batchapi/`, `caseapi/`, `documentsapi/`, `policyapi/`, `cdm*api/`, etc.)

One package per ServiceNow API surface, all structurally identical — `tableapi/` is the canonical/
fullest reference (paging + all verbs), `policyapi/` is the minimal single-resource reference. Every
package follows:

- **Constructor triad**: `New<X>RequestBuilderInternal(pathParameters, requestAdapter, ...)` (raw
  entry point), `NewDefault<X>RequestBuilder(rawURL, requestAdapter)` (default parsable),
  `New<X>RequestBuilder(rawURL, requestAdapter, factory)` (custom parsable).
- **One method per HTTP verb** (`Get`/`Post`/`Patch`/`Put`/`Delete`), each starting with:
  ```go
  if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
      return nil, nil
  }
  if conversion.IsNil(rB.GetRequestAdapter()) {
      return nil, snerrors.ErrNilRequestAdapter   // see error-sentinel note below
  }
  ```
- A matching `ToXRequestInformation` method building the `abstractions.RequestInformation`.
- Query parameters and request configuration types live in their **own files per verb**
  (`<resource>_request_builder_<verb>_query_parameters.go`,
  `..._request_configuration.go`), never inlined into the request builder file.
- A `Readme.md` per package describing the endpoint.
- A `.claude/skills/new-api-module` skill and `.claude/agents/api-module-consistency-reviewer`
  agent exist in this repo specifically to scaffold/review new modules against this pattern — use
  them instead of re-deriving the shape by hand.

### Error-sentinel layout (non-obvious — three separate places)

There are **three** distinct sentinel-error locations that look similar but are not interchangeable:
1. Root package `errors.go` (`package servicenowsdkgo`) — a couple of client-config-level sentinels.
2. `errors/errors.go` (imported as `snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"`) —
   the shared, cross-package sentinels (`ErrNilRequestAdapter`, `ErrNilResponse`, `ErrNilConfig`,
   `ErrNilBody`, etc.) that most `internal/*`, `core/*`, and `*api` nil-guards should return.
3. Some individual API packages (e.g. `tableapi/errors.go`) additionally define their **own**
   package-local sentinels with similar-but-different wording. When adding a nil-guard, check
   `errors/errors.go` first and reuse an existing sentinel by exact identity (not just matching text)
   — many call sites historically duplicated the same message as a fresh `errors.New(...)` instead of
   sharing one sentinel, which breaks `errors.Is` for callers.

### `credentials/` — authentication

Pluggable credential/auth-provider implementations (authorization code, client credentials, ROPC,
JWT bearer) satisfying Kiota's `authentication.AuthenticationProvider`, wired in via
`servicenowsdkgo.WithAuthenticationProvider(...)` or the higher-level `WithRequestAdapter(...)` client
option.

## Testing conventions

- **Unit tests**: co-located `_test.go` per file, table-driven with `testify` (`assert`/`require`),
  HTTP mocked via `httpmock` and internal `testify/mock`-based mocks in `internal/mocking`. Every new
  exported type/method needs a test — this is enforced by convention (and by the
  `api-module-consistency-reviewer` agent), not tooling.
- **Integration tests** (`tests/integration/`, `//go:build integration`): Gherkin `.feature` files
  under `tests/integration/features/` + step definitions using `godog`, `httpmock`, and `godotenv` for
  `.env`-based config — no live ServiceNow instance required.
- **E2E tests** (`tests/e2e/`, `//go:build e2e`): hit a real ServiceNow instance using credentials
  from `.env`; run manually, not part of the default suite.

## Versioning & commits

`VERSION` and `CHANGELOG.md` are generated by `release-please` from Conventional Commits
(`feat:`, `fix:`, `refactor:`, `docs:`, `perf:`, `chore:`, `test:`, `BREAKING CHANGE:` footer for
majors) — **never edit `VERSION` or `CHANGELOG.md` by hand.**
