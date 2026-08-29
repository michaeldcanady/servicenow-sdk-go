---
name: design-decisions
description: >
  Points to this repo's (servicenow-sdk-go) Architecture Decision Records,
  published on the docs site at website/docs/contributing/adrs/ -- the
  deliberate trade-offs behind hand-writing the client on top of Kiota's
  runtime abstractions instead of generating it from OpenAPI or writing it
  fully from scratch (ADR 003), storing model properties in a Kiota
  BackingStore instead of plain struct fields so "absent" and "zero" stay
  distinguishable (ADR 002), and standardizing error sentinels/messaging
  across the three separate error-sentinel locations (ADR 001). Also covers
  the v2.0-era conventions: pointer-typed query parameters wired through
  vanilla Kiota request-configuration (ADR 004), one generic core.PageIterator
  instead of per-module wrappers (ADR 005), nil-receiver guards returning a
  shared sentinel error instead of (nil, nil) (ADR 006), never scaffolding a
  builder chain ahead of an implemented operation (ADR 007), package/
  symbol/URL names as independent naming axes (ADR 008), keeping the
  RequestBuilder/RequestInformation names despite their imprecise roles, for
  Kiota/msgraph parity (ADR 009), and staying at the michaeldcanady GitHub
  org/module path for the v2 lifecycle instead of migrating to a NerdIT-Tech
  org (ADR 010) -- plus release branches that are lazily cut, downstream-only,
  and never silently diverge, with tracked forward-ports across majors (ADR
  011). Consult this BEFORE proposing or making changes to the
  request-builder/model architecture, error handling conventions, pagination,
  nil-guard behavior, module naming, branching/release flow, or anything that
  would make this SDK diverge from Kiota-generated SDK conventions (e.g.
  msgraph-sdk-go) -- these are not accidents or defaults, they were chosen
  over real alternatives for stated reasons, and changing them without knowing
  why risks silently undoing a considered trade-off. Also consult it before
  answering any "why does this repo do X" / "why not just do Y instead"
  question. Use it proactively even when not asked explicitly -- e.g. before
  suggesting plain-struct models, a generated-from-OpenAPI client, a
  per-module page-iterator wrapper, a bare (nil, nil) nil-guard return, a
  preemptive release/* branch, or a bespoke error type instead of the shared
  sentinels.
---

# Design decisions

This repo records its architectural trade-offs as ADRs, published on the
docs site at
[`website/docs/contributing/adrs/`](../../../website/docs/contributing/adrs/)
(index page: [`index.md`](../../../website/docs/contributing/adrs/index.md)) —
**not** in `Readme.md` or `CLAUDE.md`. `Readme.md`/`CLAUDE.md` cover what the
SDK does and how it's structured; the ADRs cover why, what was rejected, and
what a change here has to respect. Keeping them separate is deliberate:
cramming rationale into the top-level docs turns them into a decision log
nobody can navigate, and duplicating the same reasoning in two places means
it drifts out of sync the first time only one copy gets updated.

The catalog is eleven files numbered `001-` through `011-` (three digits, not
four), each with Docusaurus frontmatter and an entry in the index table:

1. [`001-error-standardization.md`](../../../website/docs/contributing/adrs/001-error-standardization.md)
   — centralizes sentinel errors and message phrasing in the `/errors`
   package. Note this interacts with the "three separate error-sentinel
   locations" gotcha documented in the root `CLAUDE.md` (root `errors.go`,
   `errors/errors.go`, and a few package-local `errors.go` files) — always
   reuse an existing sentinel by identity, don't create a fresh
   `errors.New(...)` with matching text, or `errors.Is` breaks for callers.
2. [`002-backing-store-models.md`](../../../website/docs/contributing/adrs/002-backing-store-models.md)
   — every model embeds `core.BaseModel` and stores properties in a Kiota
   `BackingStore` instead of plain struct fields, specifically so "the field
   was never sent" and "the field was sent as zero/empty" stay distinguishable,
   and so writes only serialize what the caller actually touched.
3. [`003-hand-written-on-kiota.md`](../../../website/docs/contributing/adrs/003-hand-written-on-kiota.md)
   — the founding decision. ServiceNow doesn't publish usable OpenAPI specs,
   so generating from Kiota's CLI was rejected; writing everything from
   scratch was rejected as reinventing undifferentiated plumbing (URI
   templates, auth, serialization, retries). Instead the SDK hand-writes
   request builders/models on Kiota's runtime libraries
   (`kiota-abstractions-go`, `kiota-http-go`, `kiota-serialization-*-go`),
   deliberately mirroring the conventions of Kiota-*generated* SDKs like
   msgraph-sdk-go: request-builder chaining, `RequestConfiguration` shapes,
   parsable factories, backed models. This is why "just match the
   Kiota-generated pattern exactly" is usually the right call when a piece of
   this SDK looks hand-rolled and slightly different from upstream Kiota
   conventions — the divergence is either a considered ADR-003 trade-off or
   drift worth fixing, not a free design choice.
4. [`004-vanilla-kiota-request-configuration.md`](../../../website/docs/contributing/adrs/004-vanilla-kiota-request-configuration.md)
   — every `*QueryParameters` struct uses pointer fields with
   `uriparametername` tags through `abstractions.ConfigureRequestInformation`,
   replacing a bespoke go-querystring wrapper, to match msgraph-sdk-go and let
   "unset" and "zero" stay distinguishable. Sharp edge: new integer query
   params must be `*int32`, not `*int` — the native encoder silently drops a
   bare `*int`.
5. [`005-generic-page-iterator-only.md`](../../../website/docs/contributing/adrs/005-generic-page-iterator-only.md)
   — `core.NewPageIterator[T]` is the one documented pagination pattern;
   per-module wrapper constructors (the old `tableapi`/`attachmentapi` ones)
   were removed and shouldn't be re-added for new modules, even for symmetry.
6. [`006-nil-receiver-sentinel-error.md`](../../../website/docs/contributing/adrs/006-nil-receiver-sentinel-error.md)
   — nil-receiver guards on verb methods return `snerrors.ErrNilRequestBuilder`,
   never a bare `(nil, nil)`/`nil`, so a nil builder fails loud at the call
   site instead of silently succeeding. Enforced by the
   `principal-software-engineer` agent via new-module design reviews.
7. [`007-no-speculative-builder-chains.md`](../../../website/docs/contributing/adrs/007-no-speculative-builder-chains.md)
   — don't add a request-builder accessor for a URL segment until the
   operation(s) behind it are actually implemented; a navigable chain with no
   working verb method is worse than not having the chain at all.
8. [`008-package-symbol-url-naming-independence.md`](../../../website/docs/contributing/adrs/008-package-symbol-url-naming-independence.md)
   — package names, exported symbol names, and accessor/URL-segment names are
   independent naming axes (see `aggregationapi` package vs. `Stats*` types vs.
   `Now().Stats()`); an inconsistency in one doesn't mean the other two need to
   change too.
9. [`009-requestbuilder-requestinformation-naming.md`](../../../website/docs/contributing/adrs/009-requestbuilder-requestinformation-naming.md)
   — keeps `RequestBuilder`/`RequestInformation` named as-is despite the
   names being technically imprecise (`RequestBuilder` is really a facade,
   `RequestInformation` is the actual builder), because they mirror
   `kiota-abstractions-go`/msgraph-sdk-go conventions. Settled through v3 —
   don't revisit without a new major-version boundary and a concrete
   consumer-facing reason.
10. [`010-no-nerdit-tech-migration.md`](../../../website/docs/contributing/adrs/010-no-nerdit-tech-migration.md)
    — declines a proposed GitHub org migration to NerdIT-Tech at v2; the
    module path stays `github.com/michaeldcanady/servicenow-sdk-go` (only
    the `/v2` semver suffix changes). Settled through v3 — don't revisit
    without a new major-version boundary and a concrete, non-speculative
    reason to move.
11. [`011-release-branches-and-cross-major-flow.md`](../../../website/docs/contributing/adrs/011-release-branches-and-cross-major-flow.md)
    — trunk-first development with lazily-cut `release/vX.Y` maintenance
    branches: fixes land on `main` first and backport down via label;
    anything landing on a maintenance branch without backport provenance
    triggers a `needs-forward-port` issue so divergence is tracked debt, never
    silent. Never create preemptive `release/*` branches or treat a
    maintenance branch as a development trunk.

Plain-language contributor summaries of ADRs 001–003 live alongside them on
the site under [Why it's built this way](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/design-decisions)
(`design-decisions.md`, `design-backed-models.md`, `design-hand-written-kiota.md`,
`design-error-handling.mdx`) — those link into these primary sources.

Before touching an area covered by an ADR, or answering a "why"/"should we
change this" question:

1. Read the specific ADR(s) that cover the area you're about to touch.
2. If what you're about to do would contradict an ADR, say so explicitly and
   confirm with the user first — don't just quietly change course. A few of
   these look like the "wrong" choice in isolation (hand-writing instead of
   generating; pointer-returning getters instead of plain fields) and are
   only correct in the context of the trade-off the ADR records.

## Keeping the catalog current

This is a living document, not a one-time snapshot. When a **real
architectural trade-off** gets decided in conversation — something with a
rejected alternative and a reason, the kind of thing a new contributor would
otherwise have to reverse-engineer from git blame — add a new ADR:

1. Check the highest existing number in
   `website/docs/contributing/adrs/`; yours is **that plus one**, full stop.
   Create `NNN-<short-title>.md` there (three-digit prefix, matching the
   existing files — don't switch to four digits).
2. Follow the shape already used by `001`–`003`: `Status`, `Context`,
   `Decision`, `Consequences` (`002`/`003` also include an `Alternatives
   considered`-style discussion inside Context — keep that). Add Docusaurus
   frontmatter (`title` matching the H1, one-line `description`) like the
   existing pages.
3. Register it: add a row to the table in
   `website/docs/contributing/adrs/index.md` and an entry to the
   "Architecture decision records" category in `website/sidebars.ts`. An
   unregistered ADR is invisible on the site.
4. If the new ADR changes or replaces an earlier one, don't edit the old
   file's decision — mark it `Superseded by ADR-00N` in its Status line and
   note the supersession in the new file, so the history of *why it changed*
   is kept rather than overwritten.

Routine bug fixes, formatting, or anything already fully explained by the
diff/commit message don't need an ADR. And if a change you're about to make
would only be correctly understood by *also* updating `CLAUDE.md`'s factual
"what/how" description (not its reasoning), update that too — ADRs and
`CLAUDE.md` should never describe two different realities.
