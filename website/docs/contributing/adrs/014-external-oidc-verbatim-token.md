---
title: 'ADR 014: External OIDC tokens are presented verbatim; ServiceNow validates identity'
description: >-
  The external OIDC authentication flow presents an IdP-issued identity token
  to ServiceNow without SDK-side validation; the OIDP record is the validator.
---

# ADR 014: External OIDC tokens are presented verbatim; ServiceNow validates identity

## Status

Accepted

## Context

The external OIDC authentication flow lets a user authenticate to ServiceNow
through a third-party identity provider (for example Microsoft Entra ID). The
SDK acts as an OIDC client: it resolves the IdP's discovery document, runs the
authorization code + PKCE grant, and sends the resulting token as a Bearer
token on API calls. ServiceNow validates it against an OIDP record
("Configure an OIDC provider to verify ID tokens") — signature via the IdP's
JWKS endpoint, expiry, the rule that `aud` equals the record's Client ID, and
a user claim that resolves to a `sys_user` record.

Because the OIDP is the validator, it is tempting to validate tokens in the
SDK too. That temptation was considered and rejected.

Alternatives considered:

1. **SDK validates tokens client-side** — verify signature, `aud`, and the
   user claim before sending. Rejected: IdP claim layouts vary (Okta's
   `sub` is provider-internal; Auth0 adds custom claims only with Actions;
   Google's *access tokens are opaque, not JWTs*), so strict validation
   either false-rejects valid configurations or ends up with per-IdP rules
   that duplicate ServiceNow's. It also adds a JWKS dependency and a fetch
   per token to the SDK.
2. **Machine identity through the same path** (client-credentials access
   tokens from the external IdP). Rejected after verification: ServiceNow's
   OIDP always resolves a user claim to `sys_user`, and a client-credentials
   token carries no user claims, so the documented outcome is
   "User Not Authenticated". Machine-to-machine stays on ServiceNow's native
   client-credentials grant (its own OAuth server), which the SDK already
   implements.
3. **Present the IdP token verbatim with no SDK-side validation** —
   Chosen. The SDK sanity-checks only that the token has a parseable `exp`
   (so the shared token cache can compute TTL), then passes it through.

## Decision

External OIDC authentication presents the IdP-issued token **verbatim** as a
Bearer token. The SDK owns the protocol (discovery, PKCE, redirect, exchange,
refresh, revoke), the token lifecycle (cache until `exp`, reuse, persist),
and safety (allowed hosts, no token logging, errors that distinguish IdP
acquisition failure from ServiceNow rejection) — and performs **no**
signature or claims validation. ServiceNow's OIDP record is the validator.
v1 runs the authorization code + PKCE grant as a public client and defaults
to the ID token (`openid` scope, `aud` = client ID by construction);
audience-scoped user access tokens are supported through the scopes option
for IdPs and setups that require them (for example Entra ID's v1
`api://{clientId}/.default` path).

## Consequences

- **Pros:** the provider stays generic across standards-compliant issuers;
  no JWKS fetching or crypto in the SDK; no risk of false-rejecting an IdP
  configuration ServiceNow accepts; the external surface is one constructor
  plus a small set of options.
- **Cons:** a misconfigured OIDP surfaces only at request time as a live
  "User Not Authenticated" instead of failing fast in the SDK; there is no
  client-side defense in depth; security depends on the ServiceNow OIDP
  configuration (user-claim mapping to `sys_user`, `aud` match, JTI policy).
  Mitigation is documentation and a troubleshooting guide, deliberately not
  code.
- **Rule for future auth work:** do not add token validation to the SDK
  unless a concrete IdP-compatibility problem demands it, and never treat
  machine identity as an OIDP concern — that is the native client-credentials
  grant's job.