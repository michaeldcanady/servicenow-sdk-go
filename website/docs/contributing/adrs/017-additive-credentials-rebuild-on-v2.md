---
title: 'ADR 017: Ship the credentials rebuild additively on v2, not as a new major'
description: >-
  The credentials rebuild lands as an additive v2 minor with deprecation shims;
  the old surface is retained and removed only in the next major.
---

# ADR 017: Ship the credentials rebuild additively on v2, not as a new major

## Status

Accepted (maintainer decision, 2026-10-13). Scopes an exception to
[ADR-011](011-release-branches-and-cross-major-flow.md) for this change;
ADR-011 otherwise still governs release branches and the cross-major flow.

## Context

The credentials rebuild replaces the per-flow credential types with a single
provider surface and adds external OIDC. The original plan landed it as a
breaking v3 change on `main`, following ADR-011's trunk-first model, where
`main` is always the tip of the next major.

Two problems surfaced with that plan:

1. The change is additive in substance. The rebuild introduces new
   constructors, option structs, and a new flow. Existing call sites would
   keep working unchanged if the old symbols stayed in place. Forcing every
   consumer through a module-path bump for an additive improvement is a large
   cost for little benefit.
2. ADR-010 makes the `/vN` module-path bump a release-day runbook item, not a
   preemptive edit. A v3 bump on `main` before release would churn every
   import suffix for a major consumers do not yet need.

Alternatives considered:

1. **Break in the next major on `main`** — land the rebuild as v3. Rejected for
   now: it pays the module-path cost immediately and forces a migration for an
   API that can stay source-compatible.
2. **Additive surface with deprecation shims, removed in the next major** —
   chosen.
3. **Ship the new surface only, with no shims** — rejected: it breaks v2
   consumers on upgrade, which the additive goal exists to avoid.

## Decision

Ship the credentials rebuild additively on v2:

1. The new provider surface is additive: new constructors, per-flow option
   structs (`ClientCredentialsOptions`, `ROPCOptions`, `JWTOptions`,
   `AuthorizationCodeOptions`, `ExternalOIDCOptions`, `BasicOptions`,
   `BearerOptions`), external OIDC, `KeyedTokenStore`, and shared sentinels in
   the `errors` package.
2. Every v2 symbol stays in place, marked deprecated in godoc: the `WithX`
   functional options, `AuthConfig`, `Authority` and its constructors and URL
   methods, the exported credential types, the `New*Credential` constructors,
   `GetAuthentication`, the client-ID `TokenStore` and `SetTokenStore`, and the
   `CredentialError` and `Oauth2Error` types and sentinels.
3. `AuthOption` becomes an interface satisfied by both the deprecated `WithX`
   helpers and the new option structs, so the six colliding constructors
   (`NewClientCredentialsProvider`, `NewJWTProvider`, `NewROPCProvider`,
   `NewAuthorizationCodeProvider`, `NewPrivateAuthorizationCodeProvider`,
   `NewPublicAuthorizationCodeProvider`) accept either call style under the
   same names.
4. The module path stays `.../v2`. The rebuild lands as a v2 minor; no new
   major is opened.
5. Removing the deprecated surface is a breaking change and waits for the next
   major release.

This is a scoped exception to ADR-011. Because the rebuild does not open a new
major, ADR-011's maintenance-branch and forward-port rules are not triggered
for it. ADR-011 still governs how release branches are cut and how cross-major
work is ported.

One source-compatibility break is deliberate and narrow: a bare
`credentials.AuthOption(func(*credentials.AuthConfig) {})` literal no longer
compiles, because `AuthOption` is now an interface. Callers use a `WithX`
helper or an option struct instead.

## Consequences

- **Pros:** existing consumers upgrade without code changes; adoption of the
  new surface is incremental; the rebuild ships without a module-path event.
- **Cons:** the deprecated surface must be carried and tested until the next
  major; the constructors offer two call styles at once; the shim types widen
  the public API.
- **Rule for future decisions:** an additive improvement does not need a new
  major just because ADR-011 treats `main` as next-major trunk. Prefer
  additivity with deprecation, and reserve a major bump for changes that cannot
  stay source-compatible.
