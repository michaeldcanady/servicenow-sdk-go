---
title: 'ADR 015: The credentials package exposes providers; a credential is internal'
description: >-
  The credentials package exposes only provider constructors, option types,
  the token store, and sentinels; the credential and its base provider are
  internal, and the ServiceNow URL is named once, on the client.
---

# ADR 015: The credentials package exposes providers; a credential is internal

## Status

Accepted

## Context

`credentials/` grew two parallel ways to name the same thing. Each flow has a
constructor returning a Kiota `authentication.AuthenticationProvider`
(`NewROPCProvider`, `NewClientCredentialsProvider`, `NewJWTProvider`,
`NewAuthorizationCodeProvider`) and a constructor returning an exported
credential type (`NewROPCCredential`, `NewClientCredentialsCredential`,
`NewJWTCredential`, `NewAuthorizationCodeCredential`).

The credential constructors take unexported interfaces (`ropcClient`,
`clientCredentialsClient`, `jwtClient`, `authorizationCodeClient`) whose method
sets are partly unexported. No package outside `credentials` can implement those
interfaces, so an external caller can pass only `nil`. The exported credential
types and their constructors are therefore API surface no consumer can use;
in-package tests are their only real caller. The package also exported
`BaseAccessTokenProvider` and `Authority`, whose only callers are in-package.

The interface between the two was also inconsistent: a caller set the
ServiceNow instance twice — once as a credential option and once as a client
option — even though the client requires a URL and passes it to the provider
through `Initialize` regardless.

Alternatives considered:

1. **Keep both surfaces and document the credential types as an extension
   point** — rejected: no external caller can construct one, so the extension
   point does not exist.
2. **Keep the credential types exported but change their constructors to take
   exported interfaces** — rejected: it publishes internal protocol-client
   seams as public API with no consumer need.
3. **Collapse to one public register** — chosen.

## Decision

The public surface of `credentials/` is the constructors that return
`authentication.AuthenticationProvider` (the `NewXProvider` functions and the
Basic and Bearer providers), plus `TokenStore`, the per-flow option types, and
the shared error sentinels.

**Credential** names the unexported token-lifecycle type that acquires, caches,
refreshes, and revokes tokens for one flow. **Provider** names the Kiota-facing
object a caller passes to the client. The credential types,
`BaseAccessTokenProvider`, and `Authority` are unexported. Each credential
receives its acquisition, refresh, and revocation hooks at construction instead
of patching unexported fields afterward.

Naming an instance is a client concern. A caller sets the ServiceNow URL once
on `ServiceNowServiceClient`, and the client passes it to the provider through
`Initialize`. The credential options no longer carry `WithInstance` or
`WithURL`. The external OIDC issuer stays a credential option because it names
the IdP, not the instance.

## Consequences

- **Pros:** the public surface matches what callers can actually use; the
  provider/credential vocabulary is unambiguous; the ServiceNow URL has one
  source, so callers stop setting it twice; and no consumer-facing break comes
  from unexporting types they could not construct.
- **Cons:** renaming exported symbols is a breaking change, so it ships in a
  major release, and the documented `credentials.WithInstance(...)` call style
  disappears from the auth snippets.
- **Related:** the naming axes stay independent per
  [ADR-008](008-package-symbol-url-naming-independence.md); this record fixes
  which register a symbol belongs to, not how it is spelled.
