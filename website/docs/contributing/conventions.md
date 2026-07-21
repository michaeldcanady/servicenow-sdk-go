---
title: Conventions reference
description: >-
  The field guide for seasoned contributors — module anatomy, the error
  taxonomy, model rules, helper composition, and the git/release conventions,
  in one dense page.
---

# Conventions reference

This is the dense page: every convention reviewers hold a PR to, in one
place, organized for lookup rather than reading order. If you're new, the
[first PR walkthrough](first-pr.mdx) and [architecture](architecture.md)
pages teach these gradually — come back here when you want the whole contract.

Each convention exists for a reason; where the reason isn't obvious, the link
goes to the [design decision](design-decisions.md) behind it.

## Module anatomy

One package per ServiceNow API surface (`tableapi/`, `attachmentapi/`, …),
all structurally identical — the uniformity is the hand-written substitute
for generated code ([why?](design-hand-written-kiota.md)).

| Rule | Detail |
| ---- | ------ |
| **File per concern** | `<resource>_request_builder.go` holds the builder; each verb's query parameters and request configuration live in their own files (`..._<verb>_query_parameters.go`, `..._<verb>_request_configuration.go`). Never inline them. |
| **Constructor triad** | Exactly three constructors per builder: `New<X>RequestBuilderInternal(pathParameters, adapter, …)`, `NewDefault<X>RequestBuilder(rawURL, adapter)`, `New<X>RequestBuilder(rawURL, adapter, factory)`. |
| **Verb methods** | One method per HTTP verb (`Get`/`Post`/`Patch`/`Put`/`Delete`), each with a matching `To<Verb>RequestInformation`. |
| **URL templates** | Unexported `const`s in Kiota URI-template syntax: `{+baseurl}/api/now/...{?sysparm_query,sysparm_limit}`. |
| **Generics** | Builders and responses are constrained by `model.ServiceNowItem` (`store.BackedModel` + `serialization.Parsable` + `GetSysID()`). |
| **Chaining** | Each child-builder method clones the parent's path parameters (`maps.Clone`), adds its own, and passes the same `RequestAdapter`. |
| **Package readme** | Every module carries a `Readme.md` summarizing the endpoint, mirroring `tableapi/Readme.md`. |

**Reference implementations:** `tableapi/` is the canonical, fullest example
(every verb, paging, generics); `policyapi/` is the minimal one. When in
doubt, look there — and if msgraph-sdk-go does it a certain way, that's the
default answer here too.

## The nil-guard prologue

Every verb method opens with the same two guards, returning **shared**
sentinels — never a fresh `errors.New`:

```go
if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
    return nil, snerrors.ErrNilRequestBuilder
}
if conversion.IsNil(rB.GetRequestAdapter()) {
    return nil, snerrors.ErrNilRequestAdapter
}
```

`internal/conversion.IsNil` is the standard nil check everywhere (it's
reflect-safe for typed-nil interfaces); a bare `x == nil` on an interface is
a review comment waiting to happen.

## The error taxonomy

Sentinels live in **three** places that look similar but aren't
interchangeable ([why?](design-error-handling.mdx)):

| Location | Import | What belongs there |
| -------- | ------ | ------------------ |
| `errors/errors.go` | `snerrors` | The shared cross-package sentinels (`ErrNilRequestBuilder`, `ErrNilRequestAdapter`, `ErrNilResponse`, `ErrNilConfig`, `ErrNilBody`, …). **Check here first** — almost every nil-guard uses one of these. |
| Root `errors.go` | `servicenowsdkgo` | A couple of client-configuration sentinels only. |
| `<module>api/errors.go` | package-local | Conditions genuinely specific to that module. |

Rules:

- Reuse a sentinel **by identity**, not by matching text — duplicating the
  message as a fresh `errors.New` breaks `errors.Is` for callers, which is
  exactly the v1 bug the standardization fixed.
- Message phrasing: `"[parameter] cannot be nil"` for nil checks,
  `"[parameter] is required"` for missing inputs; no contractions.
- HTTP error mapping always goes through `core.DefaultErrorMapping()` —
  never a bespoke error struct. Per-module mappings register via
  `internal.GetErrorRegistryInstance()`.

## Model rules

Models are backing-store-backed, not plain structs
([why?](design-backed-models.md)):

- Embed `core.BaseModel`; **never add plain data fields** to a model.
- Every property is a `GetX() (T, error)` / `setX(T) error` pair built on
  `internal/store` (`DefaultBackedModelAccessorFunc` /
  `DefaultBackedModelMutatorFunc`).
- `Serialize()` / `GetFieldDeserializers()` are built from the
  `internal/serialization` generators — no hand-rolled property plumbing.

## Compose `internal/`, don't reinvent

If you're writing a nil check, a header string, an accessor, or serializer
plumbing inside a module, stop — the helper almost certainly exists:

| Package | Use it for |
| ------- | ---------- |
| `internal/conversion` | `IsNil`, `As2`, collection casts, string→primitive converters |
| `internal/store` | Backing-store accessor/mutator generators |
| `internal/serialization` | `Serialize`/`SerializeXFunc`/`DeserializeXFunc` generators |
| `internal/http` | `RequestHeader`/`HTTPHeader`/`ContentType` enums, default middleware/client |
| `internal/ast` + `query/` | The fluent query-condition builder and its `sysparm_query` renderer |

`internal/` is never imported by consumers — it's implementation surface,
so helpers can evolve freely.

## Testing conventions

The full treatment is the [testing guide](testing.md); the reviewable
contract in brief:

- **Every exported type and method ships with tests, in the same PR** —
  co-located `_test.go`, table-driven with `testify`, HTTP stubbed with
  `httpmock` (plus `internal/mocking` doubles).
- Tables include the **failure rows**: nil-guard sentinels, at least one
  mapped API error, and the success path.
- Test data looks like ServiceNow (`INC0010001`, not `"foo"`), and the
  *request* is asserted (headers, query, body), not just the response.
- Bug fixes carry the test that fails before the fix.

## Git, commits, and releases

- **Branches:** `type/kebab-description` off `main` —
  `fix/tableapi-nil-pointer`, `docs/error-sentinel-notes`.
- **Commits and PR titles:** [Conventional Commits](https://www.conventionalcommits.org/)
  (`feat(scope): …`, `fix: …`, `BREAKING CHANGE:` footer for majors). CI
  lints the PR title; it becomes the squash commit.
- **CI/workflow changes are always `chore`** — anything touching
  `.github/workflows/`, `scripts/` CI helpers, or other pipeline plumbing
  is `chore:`, never `fix:`/`feat:`, even if it fixes a broken run. These
  changes don't affect the published SDK, so they must not surface in
  `CHANGELOG.md` as a fix or feature.
- **Never edit `VERSION` or `CHANGELOG.md`** — `release-please` generates
  both from commit messages.
- **Local gate before review:** `gofmt -s -w .`, `golangci-lint run ./...`,
  `go test ./...` (config: `.golangci.yml`; `just build|lint|fmt` wrap the
  same commands).

## Documentation conventions

- A PR that changes exported API surface updates the docs site (`website/`)
  in the same PR — or says why not in the description.
- Go samples are single-sourced from `website/snippets/*.go` behind
  `// [START x]` / `// [END x]` region markers and rendered with the
  `GoSnippet`/`GoExample` components. CI compiles every snippet
  (`go vet -tags snippets ./website/snippets/`), so samples can't rot.
- New pages register in `website/sidebars.ts`; preview with
  `just serve-docs` (Node 20+, `just setup-docs` first).
- Design changes get an ADR under `docs/adr/` plus a summary page here —
  see [Why it's built this way](design-decisions.md). New modules get a
  blueprint under `docs/blueprints/` first.

## The playbook, when you need it

Adding a whole new API surface pulls all of the above together in order:
[Add a new API module](add-api-module.md).
