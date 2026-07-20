---
name: design-decisions
description: >
  Points to this repo's (servicenow-sdk-go) Architecture Decision Records in
  docs/adr/ -- the deliberate trade-offs behind hand-writing the client on
  top of Kiota's runtime abstractions instead of generating it from OpenAPI
  or writing it fully from scratch (ADR 003), storing model properties in a
  Kiota BackingStore instead of plain struct fields so "absent" and "zero"
  stay distinguishable (ADR 002), and standardizing error sentinels/messaging
  across the three separate error-sentinel locations (ADR 001). Also covers
  the v2.0-era conventions: pointer-typed query parameters wired through
  vanilla Kiota request-configuration (ADR 004), one generic core.PageIterator
  instead of per-module wrappers (ADR 005), nil-receiver guards returning a
  shared sentinel error instead of (nil, nil) (ADR 006), never scaffolding a
  builder chain ahead of an implemented operation (ADR 007), and package/
  symbol/URL names as independent naming axes (ADR 008). Consult this BEFORE
  proposing or making changes to the request-builder/model architecture,
  error handling conventions, pagination, nil-guard behavior, module naming,
  or anything that would make this SDK diverge from Kiota-generated SDK
  conventions (e.g. msgraph-sdk-go) -- these are not accidents or defaults,
  they were chosen over real alternatives for stated reasons, and changing
  them without knowing why risks silently undoing a considered trade-off.
  Also consult it before answering any "why does this repo do X" / "why not
  just do Y instead" question. Use it proactively even when not asked
  explicitly -- e.g. before suggesting plain-struct models, a
  generated-from-OpenAPI client, a per-module page-iterator wrapper, a bare
  (nil, nil) nil-guard return, or a bespoke error type instead of the shared
  sentinels.
---

# Design decisions

This repo records its architectural trade-offs as ADRs in
[`docs/adr/`](../../../docs/adr/) — **not** in `Readme.md` or `CLAUDE.md`.
`Readme.md`/`CLAUDE.md` cover what the SDK does and how it's structured;
`docs/adr/` covers why, what was rejected, and what a change here has to
respect. Keeping them separate is deliberate: cramming rationale into the
top-level docs turns them into a decision log nobody can navigate, and
duplicating the same reasoning in two places means it drifts out of sync the
first time only one copy gets updated.

There is no index file or template in `docs/adr/` yet — it's eight files,
numbered `001-` through `008-` (three digits, not four). Read them directly:

1. [`001-error-standardization.md`](../../../docs/adr/001-error-standardization.md)
   — centralizes sentinel errors and message phrasing in the `/errors`
   package. Note this interacts with the "three separate error-sentinel
   locations" gotcha documented in the root `CLAUDE.md` (root `errors.go`,
   `errors/errors.go`, and a few package-local `errors.go` files) — always
   reuse an existing sentinel by identity, don't create a fresh
   `errors.New(...)` with matching text, or `errors.Is` breaks for callers.
2. [`002-backing-store-models.md`](../../../docs/adr/002-backing-store-models.md)
   — every model embeds `core.BaseModel` and stores properties in a Kiota
   `BackingStore` instead of plain struct fields, specifically so "the field
   was never sent" and "the field was sent as zero/empty" stay distinguishable,
   and so writes only serialize what the caller actually touched.
3. [`003-hand-written-on-kiota.md`](../../../docs/adr/003-hand-written-on-kiota.md)
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
4. [`004-vanilla-kiota-request-configuration.md`](../../../docs/adr/004-vanilla-kiota-request-configuration.md)
   — every `*QueryParameters` struct uses pointer fields with
   `uriparametername` tags through `abstractions.ConfigureRequestInformation`,
   replacing a bespoke go-querystring wrapper, to match msgraph-sdk-go and let
   "unset" and "zero" stay distinguishable. Sharp edge: new integer query
   params must be `*int32`, not `*int` — the native encoder silently drops a
   bare `*int`.
5. [`005-generic-page-iterator-only.md`](../../../docs/adr/005-generic-page-iterator-only.md)
   — `core.NewPageIterator[T]` is the one documented pagination pattern;
   per-module wrapper constructors (the old `tableapi`/`attachmentapi` ones)
   were removed and shouldn't be re-added for new modules, even for symmetry.
6. [`006-nil-receiver-sentinel-error.md`](../../../docs/adr/006-nil-receiver-sentinel-error.md)
   — nil-receiver guards on verb methods return `snerrors.ErrNilRequestBuilder`,
   never a bare `(nil, nil)`/`nil`, so a nil builder fails loud at the call
   site instead of silently succeeding. Enforced by the
   `api-module-consistency-reviewer` agent and `new-api-module` skill.
7. [`007-no-speculative-builder-chains.md`](../../../docs/adr/007-no-speculative-builder-chains.md)
   — don't add a request-builder accessor for a URL segment until the
   operation(s) behind it are actually implemented; a navigable chain with no
   working verb method is worse than not having the chain at all.
8. [`008-package-symbol-url-naming-independence.md`](../../../docs/adr/008-package-symbol-url-naming-independence.md)
   — package names, exported symbol names, and accessor/URL-segment names are
   independent naming axes (see `aggregationapi` package vs. `Stats*` types vs.
   `Now().Stats()`); an inconsistency in one doesn't mean the other two need to
   change too.

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

1. Create `docs/adr/00N-<short-title>.md` (three-digit prefix, matching the
   existing files — don't switch to four digits).
2. Follow the shape already used by `001`–`003`: `Status`, `Context`,
   `Decision`, `Consequences` (`002`/`003` also include an `Alternatives
   considered`-style discussion inside Context — keep that).
3. If the new ADR changes or replaces an earlier one, don't edit the old
   file's decision — mark it `Superseded by ADR-00N` in its Status line and
   note the supersession in the new file, so the history of *why it changed*
   is kept rather than overwritten.

Routine bug fixes, formatting, or anything already fully explained by the
diff/commit message don't need an ADR. And if a change you're about to make
would only be correctly understood by *also* updating `CLAUDE.md`'s factual
"what/how" description (not its reasoning), update that too — ADRs and
`CLAUDE.md` should never describe two different realities.
