---
title: Architecture decision records
description: >-
  The catalog of accepted and proposed ADRs — every load-bearing design trade-off
  in the SDK, what was rejected, and the rules each decision imposes on new code.
---

# Architecture decision records

Architecture Decision Records (ADRs) are this repo's memory of *why* it's
built the way it is. Each record states the context, the decision, the
alternatives that were rejected, and the consequences — including the rules a
change to that area must respect.

The three decisions most contributors bump into first have plain-language
summaries under [Why it's built this way](../design-decisions.md); the records
below are the primary sources.

| ADR | Decision | Status |
| --- | -------- | ------ |
| [001](001-error-standardization.md) | Standardizing error handling and messaging — shared sentinels in `errors/`, strict phrasing | Accepted |
| [002](002-backing-store-models.md) | Backing-store-backed models — absent ≠ zero, dirty-tracked writes | Accepted |
| [003](003-hand-written-on-kiota.md) | Hand-written client on Kiota runtime abstractions — no generator, no from-scratch | Accepted |
| [004](004-vanilla-kiota-request-configuration.md) | Vanilla Kiota request-configuration pattern for query parameters — pointer fields, no go-querystring | Accepted |
| [005](005-generic-page-iterator-only.md) | One generic `core.PageIterator`, no per-module wrappers | Accepted |
| [006](006-nil-receiver-sentinel-error.md) | Nil-receiver guards return a sentinel error, never `(nil, nil)` | Accepted |
| [007](007-no-speculative-builder-chains.md) | No builder-chain scaffolding ahead of an implemented operation | Accepted |
| [008](008-package-symbol-url-naming-independence.md) | Package names, exported symbol names, and URL segments are independent naming axes | Accepted |
| [009](009-requestbuilder-requestinformation-naming.md) | Keep `RequestBuilder`/`RequestInformation` naming (Kiota parity) — settled through v3 | Accepted |
| [010](010-no-nerdit-tech-migration.md) | Stay at `github.com/michaeldcanady/servicenow-sdk-go` — settled through v3 | Accepted |
| [011](011-release-branches-and-cross-major-flow.md) | Release branches are lazily cut, downstream-only, and never silently diverge | Accepted |
| [012](012-release-provenance-in-band.md) | Release provenance (SBOM, signatures, verification) is generated in-band, not event-driven | Accepted |
| [013](013-event-driven-release-dispatch.md) | Proposed direction: decouple release tasks via `repository_dispatch` events with reconciliation, ending at a GitHub App for native triggers | Proposed |
| [014](014-external-oidc-verbatim-token.md) | External OIDC tokens are presented verbatim; ServiceNow's OIDP validates identity | Accepted |
| [015](015-credentials-provider-surface.md) | The credentials package exposes providers; a credential is internal | Accepted |
| [016](016-keep-ropc-discouraged.md) | Keep the ROPC grant, but steer callers away from it | Accepted |
| [017](017-hybrid-domain-first-surface.md) | Hybrid domain-first surface — domain APIs for common cases, honest escape hatches for the rest | Proposed |
| [018](018-table-collection-builder.md) | `.Table()` becomes a collection builder with a `.ByName()` escape hatch; `Table2` bridges v2 | Proposed |
| [019](019-typed-tables.md) | Typed tables — built-in and user-defined, full parity, uniform shape | Proposed |
| [020](020-record-scoped-resources.md) | Record-scoped resources are explicit-args helpers, not nested builders | Proposed |
| [021](021-best-effort-error-categorization.md) | Best-effort error categorization with raw ServiceNow details attached | Proposed |
| [022](022-credential-package.md) | The credential package replaces credentials with typed providers and one shared token lifecycle | Proposed |

## Proposing a new ADR

Every ADR is maintained in **two formats** carrying the same number: a human
narrative record (here) and an agent rule file (in `docs/rules/`). The
narrative is the primary source — it holds the context, the rejected
alternatives, and the trade-offs. The rule file is the machine-scannable
distillation an agent consults before touching code in that area.

1. Check the highest existing number above; yours is **that plus one**.
   Numbering is unconditional — never reserve or skip numbers. Three-digit
   prefix, matching the existing filenames.
2. **Write the narrative record** at
   `website/docs/contributing/adrs/NNN-<short-title>.md` following the
   Status / Context / Decision / Consequences shape used by
   [002](002-backing-store-models.md) and
   [003](003-hand-written-on-kiota.md), including the alternatives considered
   inside Context. Add Docusaurus frontmatter (`title`, `description`) like
   the other pages here.
3. **Write the matching rule file** at `docs/rules/R-{AREA}-{NNN}.md`, where
   `NNN` is the same number and `{AREA}` is a short token for the concern. The
   rule file uses this shape:

   ```markdown
   ---
   id: R-{AREA}-{NNN}
   applies_to: ["{glob or path}"]
   status: accepted
   supersedes: []
   conflicts_checked: true
   ---

   <!-- Keywords: alternative names an agent might search for -->

   MUST: ...
   MUST NOT: ...
   SHOULD: ...
   SHOULD NOT: ...
   MAY: ...

   ## Verify

       ```bash
       # a command that exits non-zero when the rule is violated
       ```

   ## Rationale

   One to three sentences, or a link to the narrative ADR.
   ```

   The `MUST`/`MUST NOT`/`SHOULD`/`SHOULD NOT`/`MAY` lines use RFC 2119
   keywords. Every rule carries a `## Verify` command that is genuinely
   runnable and exits non-zero on a violation; agents (and later, CI) use it
   to check a change without reading the whole ADR. Keep `## Rationale` to
   three sentences or fewer — link to the narrative ADR for the full why.

   The `AREA` token scopes the rule for `applies_to` filtering. Existing
   tokens: `ERR` (errors), `MODEL` (models/backing store), `BUILD`
   (construction/builders), `PAGE` (pagination), `NAME` (naming/identity),
   `REL` (release/branch flow), `AUTH` (credentials/auth), `SURF` (public
   surface/entry points), `TBL` (typed tables), `ATT` (record-scoped
   resources). Add a new token when a genuinely new concern appears.
4. Rule files live under root `docs/rules/`, **outside** `website/docs/`, so
   Docusaurus does not publish them as pages — they are for agents reading the
   filesystem, not the human docs site. Do not register them in
   `website/sidebars.ts`.
5. Add the new record to the table above and to the "Architecture decision
   records" category in `website/sidebars.ts` (the narrative ADR only).
6. If your ADR supersedes an earlier one, don't edit the old decision — mark
   it `Superseded by ADR-NNN` in its Status line, set the old rule file's
   `status: superseded` and list the new ID under `supersedes`, and note the
   supersession in the new record, so the history of why it changed survives.

Routine fixes and anything fully explained by the diff don't need an ADR.
If you're unsure a change rises to ADR level, open an issue and ask.
