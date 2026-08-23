---
title: Architecture decision records
description: >-
  The catalog of accepted ADRs — every load-bearing design trade-off in the
  SDK, what was rejected, and the rules each decision imposes on new code.
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

## Proposing a new ADR

1. Check the highest existing number above; yours is **that plus one**.
   Numbering is unconditional — never reserve or skip numbers. Three-digit
   prefix, matching the existing filenames.
2. Create `website/docs/contributing/adrs/NNN-<short-title>.md` following the
   Status / Context / Decision / Consequences shape used by
   [002](002-backing-store-models.md) and
   [003](003-hand-written-on-kiota.md), including the alternatives considered
   inside Context. Add Docusaurus frontmatter (`title`, `description`) like
   the other pages here.
3. Add the new record to the table above and to the "Architecture decision
   records" category in `website/sidebars.ts`.
4. If your ADR supersedes an earlier one, don't edit the old decision — mark
   it `Superseded by ADR-NNN` in its Status line and note the supersession in
   the new record, so the history of why it changed survives.

Routine fixes and anything fully explained by the diff don't need an ADR.
If you're unsure a change rises to ADR level, open an issue and ask.
