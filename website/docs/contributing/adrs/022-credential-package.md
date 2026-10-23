---
title: 'ADR 022: The credential package replaces credentials with typed providers and one token lifecycle'
description: >-
  A new credential package implements every inbound ServiceNow authentication
  flow as typed providers, keeps the instance URL on the client, and shares one
  token lifecycle across OAuth flows.
---

# ADR 022: The credential package replaces credentials with typed providers and one token lifecycle

## Status

Proposed. Accepted on merge; this record, its rule file
(`docs/rules/R-AUTH-022.md`), and the
[credential package blueprint](../blueprints/credential-package-blueprint.md)
ship before the package implementation. The existing `credentials/` package
stays until v3.

## Context

The SDK now supports, or needs to support, every inbound REST API authentication
method ServiceNow documents:

- Basic authentication.
- Authorization code, for public and confidential clients.
- Client credentials.
- Resource owner password credentials (ROPC).
- JWT bearer.
- Third-party token grant, which this record calls external OIDC
  authentication.
- API key.
- HMAC.
- Mutual TLS.

Legacy implicit grant, browser-based SAML/SSO, and cookie sessions are outside
the scope of an API client. ServiceNow does not document device authorization
for inbound REST API authentication.

`credentials/` works for the flows it covers, but its structure does not scale
to the full set:

- It exposes two parallel public surfaces: `NewXProvider` constructors and
  `NewXCredential` constructors. The credential constructors take unexported
  interfaces, so callers outside the package cannot use them.
  `BaseAccessTokenProvider` and `Authority` are exported with no external
  consumer.
- It uses one `AuthConfig` and one `AuthOption` bag for every flow. Invalid
  combinations are possible, and the option names do not say which flow they
  belong to.
- It lets a caller name the ServiceNow instance on both the credential and the
  client, contrary to ADR 015.
- Each flow owns its own cache, refresh, persistence, and revocation logic, so
  behavior drifts.
- It keys persisted tokens only by client ID, so tokens for different
  instances, flows, scopes, or users can collide.
- It requires a client secret for JWT bearer and accepts only RS256, while
  ServiceNow documents RS, ES, and HS algorithms.
- It fetches external OIDC discovery metadata in the constructor, so
  construction performs network I/O and can fail during a transient IdP
  outage.
- It can put the client secret in the authorization URL. ServiceNow's
  private-client prose says to do that, but its parameter table omits the
  secret and OAuth 2.1 forbids it.
- It has no API key, HMAC, or mutual TLS support.
- `NewROPCProvider` does not carry the `Not recommended:` guidance that
  ADR 016 requires.

ADR 015 already establishes the provider-only surface and client-owned URL for
`credentials/`. ADR 014 makes external OIDC tokens verbatim and puts validation
in ServiceNow. ADR 016 keeps ROPC and discourages it. This record carries those
constraints into a new package instead of retrofitting the old one.

Alternatives considered:

1. **Refactor `credentials/` in place** — rejected: it would mix a public break
   with an internal rewrite, and it would leave no migration path for callers
   who need the old surface during v2.
2. **One shared option bag for the new package** — rejected: it reproduces the
   invalid-combination problem. Typed per-flow configs make required fields
   explicit and let `Validate` reject combinations that cannot work.
3. **One token lifecycle per flow** — rejected: the duplication is the current
   package's main source of drift. The flows differ in how they acquire,
   refresh, and revoke a token, not in how they cache, persist, or serialize
   access to it.
4. **Export credential types as extension points** — rejected: ADR 015. No
   external caller needs to construct the lifecycle directly; providers and
   capability interfaces cover the use cases.
5. **Keep the URL or instance on the credential** — rejected: ADR 015, and it
   gives the instance two sources of truth. The client already owns the URL and
   passes it to the provider through initialization.
6. **Make mutual TLS an `AuthenticationProvider`** — rejected: mutual TLS is a
   transport configuration, not a token or request-signing strategy. Modeling
   it as a provider would force it into a lifecycle it does not have.
7. **Wrap Kiota's `ApiKeyAuthenticationProvider` directly** — rejected:
   ServiceNow deployments also use query placement, a custom parameter name,
   and a key prefix. A thin ServiceNow provider can support those and share the
   package's sentinels without duplicating Kiota's host checks.
8. **Include implicit grant, device authorization, SAML/SSO, and cookie
   sessions** — rejected: implicit is legacy, ServiceNow does not document
   device authorization for inbound REST API authentication, and SAML/SSO and
   cookie sessions are browser flows.
9. **Split the design across one ADR per flow** — rejected: the package
   boundary, URL ownership, and shared lifecycle are one decision. Per-flow
   behavior belongs in the blueprint and tests; one rule file keeps enforcement
   in one place.
10. **Add a Go submodule** — rejected: the package belongs to the existing
    module. A submodule would add release and dependency-version overhead for
    no consumer benefit.
11. **Use functional options instead of typed configs** — rejected: functional
    options hide required inputs and allow invalid combinations. The typed
    config is the validation boundary.

## Decision

Add a new `credential/` package to the existing module. It becomes the supported
authentication surface. `credentials/` stays for compatibility, receives only
security fixes after parity, and is removed in v3 with a `BREAKING CHANGE`.

The package name is `credential`, singular, and it lives in
`github.com/michaeldcanady/servicenow-sdk-go/v2`. Protocol code stays in
`internal/oauth2`; the new package does not expose a second OAuth client.

### Surface and vocabulary

**Provider** names the Kiota-facing object a caller passes to the client. The
new package does not add its own provider interface; constructors return
concrete provider types that implement `authentication.AuthenticationProvider`.

**Credential** names the unexported token lifecycle. It acquires, caches,
refreshes, persists, and revokes tokens for one OAuth flow. Callers construct a
provider, never a credential.

**Direct provider** names a provider that authenticates each request without a
token lifecycle: Basic, static Bearer, API key, and HMAC. Mutual TLS is a
transport helper, not a provider.

Provider constructors take one typed config struct per flow and return concrete
provider types; the client accepts the `authentication.AuthenticationProvider`
interface. Constructors validate the config through `Validate` and return an
error. They perform no network I/O, start no callback server, and open no
browser.

`Flow` is an integer-backed enum. `FlowUnknown` is the zero value and is
invalid for a token key; a provider sets a known flow before it keys a token or
labels an error. The numeric values are stable because they key persisted
tokens: append new flows, and never reorder or renumber existing values.
`Flow.String()` returns stable diagnostic names.

The package exports only:

- Constructors and config structs.
- Concrete provider types.
- `TokenStore`, `MemoryTokenStore`, `Token`, `TokenKey`, and `Flow`.
- Supporting enums for token kind, authentication location, client
  authentication method, and signing algorithms.
- `OAuthError`.
- Optional capability interfaces (`Initializer`, `Revoker`) and signer
  interfaces.
- `NewMemoryTokenStore` and `NewMTLSHTTPClient`.

It does not export a credential type, `BaseAccessTokenProvider`, `Authority`, or
a provider interface of its own.

`ProviderConfig` contains transport settings only: `HTTPClient` and
`AllowedHosts`. Flow configs embed it. It is composition for shared transport
settings, not a shared option bag.

Providers that need the token endpoint implement
`Initialize(baseURL string) error`. Providers that do not need it implement a
no-op. `Initializer` and `Revoker` are optional capability interfaces; callers
type-assert for them.

### URL ownership and initialization

No credential config contains a ServiceNow URL or instance. The external OIDC
`Issuer` stays because it names the external IdP, not the instance.

The caller sets the instance once on the client with `WithURL` or
`WithInstance`. The client already owns the base URL and passes it to the
provider.

The client calls `Initialize(baseURL string) error` first. For old providers it
falls back to the legacy `Initialize(string)` method. It returns initialization
errors from `NewServiceNowServiceClient`, which already returns an error. If
only the legacy method exists and the URL is empty, the client returns
`ErrEmptyURI`.

A provider that needs the token endpoint returns `ErrEmptyURI` from
`Initialize("")`. A direct provider returns nil. The client uses `ErrEmptyURI`
for an empty base URL instead of a new `errors.New` message. `Initialize` is
idempotent: a second call with the same URL does nothing, and a different URL
returns `ErrAlreadyInitialized`. A provider that is used before initialization
and needs a URL returns `ErrNotInitialized`.

Direct providers require HTTPS and honor `AllowedHosts` through Kiota's
allowed-host validator. When `AllowedHosts` is empty, that validator allows any
HTTPS host. Set `AllowedHosts` to the instance host when you need a strict
allowlist. `AllowedHosts` is an allowlist, not a second instance setting; the
client still owns the URL.

### Shared OAuth token lifecycle

All OAuth flows use one unexported lifecycle. The lifecycle owns cache, expiry
check, refresh, persistence, revocation, and single-flight. Flows supply hooks
for initial acquisition, refresh, revocation, and token normalization. They do
not implement their own cache or locking.

The cache key is a `TokenKey` with the base URL, client ID, flow, sorted scopes,
and an optional session key. User-specific flows require a session key when
persistence is enabled. The default store is in memory and per provider.
Persistence is opt-in through `TokenStore`.

A token is expired when `exp` is within a small configurable skew. A refresh
response that omits a refresh token keeps the existing one. Refresh happens only
when the authority issued a refresh token, except client credentials, which
re-acquire when no refresh token exists.

On `invalid_grant`:

- Non-interactive grants that can acquire a new token without user interaction
  clear the token and re-acquire once. A caller-supplied JWT assertion cannot
  re-acquire and returns `ErrReauthRequired`.
- ROPC clears the token and returns `ErrReauthRequired`; it does not silently
  reuse or retain the password.
- Interactive grants, authorization code and external OIDC, clear the token and
  return `ErrReauthRequired`.

Transient refresh failures return an error and keep the token. If the token is
still valid, the request can use it; if it is expired, the caller can retry.

Acquisition and refresh are single-flight per token key. Concurrent requests
share one attempt. A canceled waiter does not cancel the shared attempt.

Revocation is best-effort and optional through `Revoker`. A missing revocation
endpoint clears the local token and succeeds. A protocol failure returns an
error and leaves the cached token in place.

Store failures are best-effort: a load error is a cache miss, and a save error
does not fail the request. Persistence never becomes a new outage.

### Errors and redaction

Authentication sentinels live in `errors/` with the other shared sentinels. This
record adds `ErrReauthRequired` and the config and initialization sentinels the
package needs. It reuses existing sentinels such as `ErrNilConfig`, `ErrEmptyKey`,
and `ErrEmptyURI` where they fit.

`ErrTokenAcquisition` remains the identity for an authority that fails to issue
a token. Add it to `errors/` for the new package; the old
`credentials.ErrTokenAcquisition` stays until v3. `ErrNotAuthenticated` remains
the identity for ServiceNow rejecting a request. Callers must be able to
distinguish the two with `errors.Is`.

`OAuthError` carries the operation, flow, protocol error code, description,
protocol error URI, HTTP status, and a wrapped sentinel. `URI` is the RFC 6749
`error_uri`, never the authorization URL.

Callers use `errors.Is` for sentinels and `errors.As` for `OAuthError`.

Errors, logs, and formatted strings never include access tokens, refresh tokens,
ID tokens, client secrets, passwords, API keys, HMAC secrets or signatures,
authorization codes, PKCE verifiers, or complete authorization URLs. They may
name the flow, the operation, the host, and the protocol error code.

### Flow decisions

OAuth configs default confidential clients to `client_secret_post` and support
`client_secret_basic`; public clients use `none`. External OIDC prefers the
IdP's advertised token endpoint authentication method, with an explicit
override.

**Basic.** A direct provider. It sets the `Authorization` header to
`Basic base64(username:password)` for allowed HTTPS hosts. It retains the
password because every request needs it, so the docs steer callers to client
credentials or mutual TLS for production.

**Static Bearer.** A direct provider for a token the caller acquired elsewhere.
It sets the `Authorization` header to `Bearer <token>` and does not cache or
refresh it.

**API key.** A direct provider. The default is the `x-sn-apikey` header. Query
placement, a custom parameter name, and a key prefix are explicit opt-ins.
Header placement is the default because query strings leak into logs and
proxies.

**HMAC.** A direct provider. The default signer signs the raw request body with
SHA-256, encodes the result as standard base64, and sets
`x-sn-hmac-signature-256` to `keyId=...,signature=...`. SHA-384 and SHA-512,
header or query placement, a custom signer, and an optional key ID are
supported. The provider computes the signature before it adds a query
authentication parameter. A GET request signs an empty body.

**Mutual TLS.** `NewMTLSHTTPClient` returns an `*http.Client` configured with
client certificates. The caller passes it through `WithHTTPClient`. It is not an
`AuthenticationProvider` and has no token lifecycle.

**Authorization code.** The provider supports public and confidential clients
through the same constructor. A public client has no client secret and uses
S256 PKCE. A confidential client sends its secret only to the token endpoint.
The default is a local callback server on loopback plus the system browser
opener. A callback override lets a web application own the redirect. State is
generated and validated. The client secret is absent from the authorization URL
by default; `ClientSecretInAuthURL` is an explicit opt-in for ServiceNow
instances whose documented behavior requires it.

**Client credentials.** A confidential provider for machine identity. It
acquires a token on first use, uses the shared lifecycle, and re-acquires when
the authority issued no refresh token. It has no session key because the token
is not user-specific.

**ROPC.** The provider stays available and keeps the `Not recommended:`
guidance from ADR 016. It uses the password for the first token and then drops
its reference. A caller-owned `PasswordSource` can supply a password for an
explicit re-acquisition. A refresh failure with `invalid_grant` clears the token
and returns `ErrReauthRequired`; the provider does not silently reuse a retained
password.

**JWT bearer.** The provider accepts a caller-supplied assertion or builds one
from claims. It supports RS256, RS384, RS512, ES256, ES384, ES512, HS256,
HS384, and HS512. The client secret is optional and, when set, authenticates the
client at the token endpoint. A caller-supplied assertion gets only an `exp`
parse; the SDK does not verify its signature or other claims. An SDK-built
assertion generates a `jti`, defaults `iss`, `sub`, and `aud` to the client ID,
and validates its claims before sending.

**External OIDC.** The provider implements the third-party token grant. It
supports interactive authorization code plus PKCE and a caller-supplied token.
`TokenKind` is explicit: ID token by default and access token as an opt-in. The
provider fetches discovery metadata lazily and caches it; it does not fetch in
the constructor. A confidential client at the external IdP is optional; the
secret goes only to the IdP token endpoint. The token is presented verbatim.
The SDK parses only `exp` and performs no signature or claims validation, per
ADR 014. A token without a parseable `exp` is rejected. A caller-supplied token
skips discovery, refresh, and revocation.

ADR 014 remains authoritative for external OIDC. This record extends it in v2
with optional confidential-client authentication at the external IdP and a
caller-supplied static token path. Neither extension adds SDK-side validation.

### Exclusions

The package excludes:

- Implicit grant, which is legacy and has no refresh token.
- Device authorization, which ServiceNow does not document for inbound REST API
  authentication.
- SAML/SSO and cookie sessions, which are browser flows.
- Client-side validation of external OIDC signatures or claims, per ADR 014.
- Token introspection. ServiceNow validates the token; the SDK does not query
  an introspection endpoint.

### Migration and release

The package ships additively in v2. `credentials/` remains available. After the
new package reaches parity, mark `credentials/` deprecated with a
`// Deprecated:` paragraph that names `credential` and plans removal in v3.
Write the notice with the `{unreleased}` placeholder; the `stamp-deprecations`
workflow fills it. Migrate in-repo callers and documentation. Remove
`credentials/` in v3 with a `BREAKING CHANGE`.

ROPC is not deprecated. It keeps the `Not recommended:` paragraph and the
guidance in the authentication guide.

The new package is one public package in one Go module. There is no
`credential/go.mod`, and there is no second OAuth client.

### Documentation

Each flow ships with a user-guide page, a snippet, and an entry in the migration
guide. The authentication index lists the new provider and its use case. The
glossary defines Provider, Credential, and Direct provider. The
[blueprint](../blueprints/credential-package-blueprint.md) is the implementation
plan, not a shipped API reference. Do not document a flow before it ships.

## Consequences

- **Pros:** every OAuth flow shares one cache, refresh, persistence, and
  revocation implementation; typed configs reject invalid combinations at
  construction; the ServiceNow URL has one source; the package covers every
  inbound REST API authentication method; the additive path lets v2 callers
  migrate without a forced break; security defaults keep secrets out of URLs,
  errors, and logs.
- **Cons:** `credential/` and `credentials/` coexist until v3, so the repo
  carries duplicate code for a release line; there are more config structs than
  the old option bag; removing `credentials/` is a breaking change; Basic and
  the direct providers retain their credentials for every request; HMAC
  canonicalization varies by deployment, so the default signer cannot cover
  every custom setup.
- **Risks and mitigations:** ServiceNow's private-client prose conflicts with
  its parameter table and OAuth 2.1, so the default follows the specification
  and the opt-in is explicit. External OIDC token validation stays out of the
  SDK, so a misconfigured OIDP surfaces at request time; the mitigation is a
  troubleshooting guide and a typed rejection boundary. Interactive flows are
  harder to test; the callback override and fake code source make them
  testable without a browser. Adding a second package risks drift; the old
  package is frozen except for security fixes, and a parity checklist gates
  its deprecation.
- **Rule for future authentication work:** add a flow only when ServiceNow
  documents it for inbound REST API authentication, and add it to the
  blueprint before code. Keep the provider-only surface, typed configs, the
  client-owned URL, the shared lifecycle, shared sentinels, and redaction. Do
  not export credential types, add a URL to a credential config, add a token
  lifecycle to a direct provider, require a client secret for JWT bearer, or
  add SDK-side validation of external OIDC tokens without superseding
  ADR 014.

See also: [ADR 001](001-error-standardization.md) (shared sentinels),
[ADR 006](006-nil-receiver-sentinel-error.md) (nil guards),
[ADR 007](007-no-speculative-builder-chains.md) (no speculative builder
chains), [ADR 008](008-package-symbol-url-naming-independence.md) (naming
axes), [ADR 011](011-release-branches-and-cross-major-flow.md) (release flow),
[ADR 014](014-external-oidc-verbatim-token.md) (external OIDC verbatim
tokens), [ADR 015](015-credentials-provider-surface.md) (provider-only
surface), and [ADR 016](016-keep-ropc-discouraged.md) (ROPC guidance).
