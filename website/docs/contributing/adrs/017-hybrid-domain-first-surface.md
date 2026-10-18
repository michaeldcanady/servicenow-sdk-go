---
title: 'ADR 017: Hybrid domain-first surface — domain APIs for common cases, honest escape hatches for the rest'
description: >-
  The SDK presents domain-oriented APIs for core workflows and always keeps an
  escape hatch to the raw ServiceNow surface, so callers reach for intent, not
  ServiceNow REST paths.
---

# ADR 017: Hybrid domain-first surface — domain APIs for common cases, honest escape hatches for the rest

## Status

Proposed. Accepted on merge; this record and its rule file
(`docs/rules/R-SURF-017.md`) ship together with the design-session work.

## Context

The SDK's core premise should be: a ServiceNow developer says *"I want to do
X"* and reaches for the SDK — not for ServiceNow's REST documentation. Today
the surface is a faithful but low-level mirror of the raw API, so callers must
already think in ServiceNow terms (table names, sys_ids, encoded query strings,
`sysparm_*` parameters) to get anything done.

Three shapes were considered for fixing that:

1. **Thin wrapper over the raw REST surface** — a thin wrapper keeps the SDK
   honest but fails the premise: callers still need to know ServiceNow's URL
   structure and parameter conventions, just with an extra layer in front.
2. **Fully opinionated domain SDK** — model every real workflow (Incident,
   Case, Change, CMDB, …) as a first-class domain API. Rejected: ServiceNow's
   surface is too broad and too instance-specific to model exhaustively; the
   SDK would always lag and would paint itself into corners it can't cover.
3. **Hybrid** — ship domain-oriented APIs for the common workflows and always
   keep an honest escape hatch to the raw surface.

Strong typing is a stated primary goal, and users must be able to build their
own strongly-typed table SDKs (Incidents, Cases, MyCoolTable) on the exported
foundation — the built-ins are conveniences, not a walled garden. New
implementation ships alongside the old for the rest of v2; the old is removed
at the next major.

## Decision

Adopt the **hybrid** surface.

- Present a domain-oriented API for each core workflow (CRUD, query,
  pagination, attachments) so a caller can express intent without knowing
  ServiceNow's REST paths.
- Always keep an escape hatch from every domain affordance to the raw ServiceNow
  surface — raw table access (`ByName`), raw query parameters, raw response
  envelopes. No dead ends.
- Let consumers build their own strongly-typed table SDKs on the exported
  foundation with full parity to the built-in tables (see ADR 019).
- Ship the core workflows first and grow outward; do not speculate broad domain
  coverage ahead of a concrete workflow (consistent with ADR 007).

## Consequences

- **Pros:** callers reach for the SDK to express intent, and the common cases
  become one-liners; escape hatches mean the raw surface is always one step
  away, so the domain layer never becomes a cage; user-defined typed tables can
  match the built-ins.
- **Cons:** two surfaces (domain and raw) to maintain and keep coherent; the
  domain layer can lag the raw surface, so the escape hatch must stay genuinely
  reachable; there is a constant pull to over-speculate domain APIs, which this
  record restrains by deferring to ADR-007.
- **Rule for future surface work:** when adding to the surface, ask whether the
  common case can be expressed as intent and whether the raw path is still one
  step away — a change that closes the escape hatch or forces ServiceNow
  knowledge for a core workflow is a regression against this decision.

See also: ADR 018 (entry surface), ADR 019 (typed tables), ADR 020
(record-scoped helpers), ADR 021 (error categorization).
