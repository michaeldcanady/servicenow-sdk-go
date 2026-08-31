---
name: design-decisions
description: >-
  Carries this repo's load-bearing architecture decisions (servicenow-sdk-go) as
  ready-to-apply rules, so design work and review don't accidentally undo a
  considered trade-off. Hand-written-on-Kiota client that mirrors Kiota-generated
  SDK conventions (msgraph-sdk-go); backing-store models not plain structs;
  shared error sentinels; nil-receiver guards returning sentinels not (nil, nil);
  pointer-typed query parameters; one generic page iterator; no speculative
  builder chains; three independent naming axes; trunk-first with lazily-cut
  release/vX.Y branches. USE when the user asks "why does this repo do X"/"why
  not Y instead", wants a design review/RFC, or proposes anything touching the
  request-builder/model architecture, error handling, pagination, nil-guard
  behavior, naming, module path, or branch/release flow. USE proactively when
  you're about to suggest a plain-struct model, a generated-from-OpenAPI
  client, a per-module page-iterator wrapper, a fresh errors.New where a shared
  sentinel exists, a bare (nil, nil) nil-guard return, a speculative builder
  chain, a rename of RequestBuilder/RequestInformation, or a preemptive
  release/* branch. Full records: website/docs/contributing/adrs/. Not for
  routine bug fixes or formatting, which don't need ADR backing.
---

# Design decisions

This repo's simplest-looking code is its most deliberately chosen. Nearly
every "why is this odd?" has an Architecture Decision Record (ADR) behind it,
and the ADRs weren't accidents — they were chosen over real alternatives. This
skill distills those decisions into rules you can apply *to the code you're
touching right now*. Where a rule needs the full reasoning — alternatives
considered, rejected options, consequences — the linked ADR has it.

The rulebook is grouped by **situation**, not by document. Find your situation,
apply the rule. Deep dives live at
[`website/docs/contributing/adrs/`](../../../website/docs/contributing/adrs/)
(index: [`index.md`](../../../website/docs/contributing/adrs/index.md)),
with plain-language summaries in
`website/docs/contributing/design-decisions.md`.

## First principles

1. **Mirror Kiota-generated conventions.** This SDK is hand-written on
   Microsoft's Kiota runtime (`kiota-abstractions-go`, `kiota-http-go`,
   `kiota-serialization-*-go`) and deliberately conforms to what a
   *generated* Kiota SDK would produce (msgraph-sdk-go is the reference
   shape): request-builder chaining, `RequestConfiguration` types,
   parsable factories, backed models. ServiceNow publishes no usable OpenAPI
   specs, so generating was rejected; writing from scratch was rejected as
   reinventing Kiota's plumbing. **Rule:** when a piece of this SDK looks
   hand-rolled and differs from the Kiota-generated shape, decide whether the
   divergence is a considered ADR trade-off or drift to fix — don't accept it
   as a free design choice, and don't "improve" it away from Kiota parity.
   ([ADR 003](../../../website/docs/contributing/adrs/003-hand-written-on-kiota.md))
2. **Distinguish absent from zero.** Backing stores exist so "the caller never
   set this" and "the caller set zero/empty" stay distinguishable, and so
   `Patch` bodies only serialize what was actually set. **Rule:** models embed
   `core.BaseModel`, properties go through `internal/store` accessor/mutator
   pairs (`GetX() (T, error)` / `setX(T) error`), getters return pointers
   (`nil` = not sent). Never reach for a plain struct while writing a model.
   ([ADR 002](../../../website/docs/contributing/adrs/002-backing-store-models.md))

## Rules by situation

### Error handling — adding or triaging an error path

- **Reuse sentinels by identity, never by text.** Shared sentinels live in
  `errors/errors.go` (`snerrors.ErrNilRequestBuilder`, `ErrNilRequestAdapter`,
  `ErrNilResponse`, `ErrNilConfig`, `ErrNilBody`, ...). A fresh
  `errors.New("... cannot be nil")` with matching text breaks `errors.Is` for
  every caller — the exact bug v2 reworked away. Message phrasing is
  standardized: `"[param] cannot be nil"`, `"[param] is required"`, no
  contractions.
- **Route HTTP errors through the shared mapping.** Request builders pass
  `core.DefaultErrorMapping()` to every `Send`/`SendPrimitive` call, never a
  bespoke error struct. There are *three* error-sentinel locations that look
  alike (root `errors.go`, `errors/errors.go`, package-local ones) — check
  `errors/errors.go` first for an identity to reuse.
- ([ADR 001](../../../website/docs/contributing/adrs/001-error-standardization.md))

### Nil-guard behavior — writing a verb method

- **A nil receiver returns a sentinel, never `(nil, nil)`.** Every verb method
  opens with:
  ```go
  if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
      return nil, snerrors.ErrNilRequestBuilder
  }
  if conversion.IsNil(rB.GetRequestAdapter()) {
      return nil, snerrors.ErrNilRequestAdapter
  }
  ```
  (single-error-return verbs return the sentinel alone). A bare `(nil, nil)` /
  `nil` is a silent-failure trap the ADR explicitly rejects — it's a
  consistency-review finding, not a style nit. Navigation methods returning a
  bare `*Builder` and other nil-response semantics are out of scope.
- ([ADR 006](../../../website/docs/contributing/adrs/006-nil-receiver-sentinel-error.md))

### Query parameters — adding a param to a request

- **Pointer fields, `uriparametername` tags, wired through
  `abstractions.ConfigureRequestInformation`.** Unset ≠ zero is the whole
  point, exactly like msgraph-sdk-go. **Sharp edge:** integer params must be
  `*int32`, not `*int` — the native encoder silently drops a bare `*int` with
  no error. Also: request builders that carry a body skip the generic
  `ConfigureRequestInformation` path for content and set it manually via
  `SetContentFromParsable` in `ToXRequestInformation`.
- ([ADR 004](../../../website/docs/contributing/adrs/004-vanilla-kiota-request-configuration.md))

### Pagination — adding a collection endpoint

- **One pattern: `core.NewPageIterator[T]`.** Per-module wrappers were removed
  deliberately (three near-identical constructors, and they implied `Link`-
  header pagination endpoints may not support). Don't add a
  `New<X>PageIterator` for a new module, "even for symmetry." Document
  per-endpoint whether it actually emits `Link` headers.
- ([ADR 005](../../../website/docs/contributing/adrs/005-generic-page-iterator-only.md))

### Builder chains — adding an accessor

- **No speculative chains.** A request-builder accessor exists only when the
  operations it leads to are implemented. A navigable chain with no working
  verb method is worse than no chain at all.
- ([ADR 007](../../../website/docs/contributing/adrs/007-no-speculative-builder-chains.md))

### Naming — package, symbol, accessor

- **Three independent axes.** Package names follow ServiceNow's API-surface
  name; exported symbols follow the wire format; accessor methods follow the
  URL segment. They don't have to agree (`aggregationapi` / `Stats*` /
  `Now().Stats()`). An inconsistency in one axis doesn't make the others
  wrong.
- **`RequestBuilder` / `RequestInformation` keep their names.** Technically
  imprecise (`RequestBuilder` is a facade, `RequestInformation` is the real
  builder), but they mirror kiota-abstractions-go/msgraph-sdk-go. Settled
  through v3 — don't propose a rename.
- ([ADR 008](../../../website/docs/contributing/adrs/008-package-symbol-url-naming-independence.md),
  [ADR 009](../../../website/docs/contributing/adrs/009-requestbuilder-requestinformation-naming.md))

### Module path / org — versioning or org questions

- **Stays `github.com/michaeldcanady/servicenow-sdk-go`.** Only the `/vN`
  semantic-version suffix changes across majors (a release-day runbook item).
  No GitHub org transfer is planned through v3 — don't propose one.
- ([ADR 010](../../../website/docs/contributing/adrs/010-no-nerdit-tech-migration.md))

### Branch & release flow — where work should land

- **Trunk-first.** `main` is always the tip of the next major's development.
  All changes arrive by PR; no in-place commits to `main` or `release/*`.
- **Maintenance branches are `release/vX.Y`** (no patch segment), **created
  lazily** at first need, never preemptively. Never cut a `release/*` branch
  ahead of need, and never treat one as a development trunk.
- **Downstream-only within a major.** Fixes land on `main` first, then move
  down via a label-triggered backport (`backport release/vX.Y`). Direct pushes
  to `release/*` are blocked.
- **Drift is tracked debt, never silent.** A merge into `release/*` without
  backport provenance opens a `needs-forward-port` issue — it closes by
  porting up or recording an explicit "won't port."
- ([ADR 011](../../../website/docs/contributing/adrs/011-release-branches-and-cross-major-flow.md))

## Handling a conflict with a rule

If what you're about to do would contradict one of these decisions, **say so
explicitly and confirm with the user before changing course** — don't quietly
diverge. Several of these look like the "wrong" choice in isolation
(hand-writing instead of generating; backed models instead of plain structs)
and are only correct in light of the trade-off the ADR records. When a real
new trade-off gets decided in conversation (a rejected alternative plus a
reason), note it as an ADR candidate for the product-manager agent rather than
silently codifying it in code.

## When a new ADR is warranted

A real architectural trade-off — rejected alternative, stated reason, the kind
of thing a new contributor would otherwise reverse-engineer from `git blame` —
deserves an ADR:

1. **Number:** highest existing number in
   `website/docs/contributing/adrs/` plus one (`NNN-<short-title>.md`,
   three-digit prefix).
2. **Shape:** `Status`, `Context` (incl. alternatives considered), `Decision`,
   `Consequences`, with Docusaurus frontmatter — match `001`–`003`.
3. **Register:** add a row to `index.md` and an entry in `website/sidebars.ts`
   — an unregistered ADR is invisible on the site.
4. **Supersession:** if it replaces an earlier decision, mark the old file
   `Superseded by ADR-00N` in its Status; keep the history, don't overwrite.
5. **Keep CLAUDE.md honest:** if the change alters a "what/how" fact in
   `CLAUDE.md`, update it too — ADR and CLAUDE.md must never describe two
   different realities.

Routine bug fixes, formatting, and behavior fully explained by the diff don't
need an ADR.