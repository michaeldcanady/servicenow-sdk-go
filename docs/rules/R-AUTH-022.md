---
id: R-AUTH-022
applies_to: ["credential/**/*.go", "website/docs/user-guide/authentication/**/*.mdx", "website/snippets/auth.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# The credential package exposes typed providers and owns no instance URL

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- Keywords: Provider, Credential, direct provider, token store, initialization, authorization code, client credentials, ROPC, JWT bearer, external OIDC, API key, HMAC, mutual TLS -->
<!-- TODO(owner): add the exact ServiceNow HMAC canonical-string examples and test vectors when the implementation lands. -->

MUST: Implement each inbound ServiceNow REST API authentication flow in `credential/`: Basic, static Bearer, API key, HMAC, authorization code, client credentials, ROPC, JWT bearer, and external OIDC (third-party token grant) as providers, and mutual TLS as a transport helper.
MUST: Return a concrete provider type from every `NewXProvider` constructor; the concrete type must implement `authentication.AuthenticationProvider`.
MUST: Export the concrete provider types and keep their state unexported.
MUST: Accept one typed per-flow config struct in each provider constructor and validate it through `Validate` before returning.
MUST: Define `Flow` as an integer-backed enum with `FlowUnknown` as the zero value and never use `FlowUnknown` for a token key or a flow-specific error.
MUST NOT: Reorder or renumber existing `Flow` values; append new flows.
MUST: Keep `Credential` and the shared OAuth token lifecycle unexported.
MUST: Use the shared OAuth token lifecycle for authorization code, client credentials, ROPC, JWT bearer, and external OIDC.
MUST: Keep every ServiceNow URL and instance name off credential configs. Set the URL once on the client and pass the base URL through `Initialize`.
MUST: Implement `Initialize(baseURL string) error` on providers that need a token endpoint; implement a no-op on direct providers.
MUST: Use `ProviderConfig` only for transport settings (`HTTPClient` and `AllowedHosts`), never as a shared option bag.
MUST: Perform validation only in constructors. Do not perform network I/O, start a callback server, or open a browser.
MUST: Use shared sentinels from `errors/` and the typed `OAuthError`.
MUST: Keep `ErrTokenAcquisition` and `ErrNotAuthenticated` distinguishable under `errors.Is`.
MUST: Return `ErrReauthRequired` when an interactive flow or a caller-supplied JWT assertion cannot refresh and the caller must authenticate again.
MUST: Fetch external OIDC discovery metadata lazily and cache it.
MUST: Use an explicit external OIDC `TokenKind`, with ID token as the default and access token as an opt-in.
MUST: Support a caller-supplied external OIDC token that skips discovery, refresh, and revocation.
MUST: Support optional confidential-client authentication at the external IdP, with the secret sent only to the IdP token endpoint.
MUST: Present external OIDC tokens verbatim and parse only `exp` (R-AUTH-014). Reject an external OIDC token or caller-supplied JWT assertion without a parseable `exp`.
MUST: Send the client secret only to the token endpoint by default.
MUST: Default confidential ServiceNow native OAuth clients to `client_secret_post`, allow `client_secret_basic`, and use `none` for public clients.
MUST: Prefer the external IdP's advertised token endpoint authentication method, with an explicit override.
MUST: Use S256 PKCE for public authorization code clients.
MUST: Support the local callback server and browser opener by default and a callback override for web applications.
MUST: Generate and validate the authorization state.
MUST: Support RS256, RS384, RS512, ES256, ES384, ES512, HS256, HS384, and HS512 for JWT bearer.
MUST: Accept a caller-supplied JWT assertion or build one from claims.
MUST: Parse only `exp` from a caller-supplied JWT assertion; do not verify its signature or other claims.
MUST: Generate a `jti`, default `iss`, `sub`, and `aud` to the client ID, and validate an SDK-built assertion before sending it.
MUST: Default API key authentication to the `x-sn-apikey` header and support query placement, a custom parameter name, and a prefix.
MUST: Support HMAC SHA-256, SHA-384, and SHA-512, header or query placement, and a custom signer.
MUST: Default HMAC to SHA-256, standard base64, the matching `x-sn-hmac-signature-{256,384,512}` header, and `keyId=...,signature=...`.
MUST: Compute the HMAC before adding a query authentication parameter and sign an empty body for GET.
MUST: Keep ROPC available with a `Not recommended:` godoc paragraph and no `Deprecated:` marker (R-AUTH-016).
MUST: Discard the ROPC password reference after the first token. A caller that sets `PasswordSource` opts into explicit re-acquisition. Return `ErrReauthRequired` after an `invalid_grant` refresh failure.
MUST: Make mutual TLS a transport helper returned by `NewMTLSHTTPClient`, not an `AuthenticationProvider`.
MUST: Require HTTPS and honor `AllowedHosts` through Kiota's allowed-host validator.
MUST: Keep tokens, secrets, passwords, API keys, HMAC secrets and signatures, authorization codes, PKCE verifiers, and complete authorization URLs out of errors and logs.
MUST: Key persisted tokens by base URL, client ID, flow, sorted scopes, and session key.
MUST: Use single-flight acquisition and refresh per token key.
MUST: Keep transient refresh failures from discarding a possibly valid token.
MUST: Treat token-store failures as best-effort; a store failure never fails an API request.
MUST: Document every flow with a user-guide page and an entry in the migration guide.
MUST NOT: Export a credential type, `BaseAccessTokenProvider`, `Authority`, or a package-specific provider interface.
MUST NOT: Define `WithURL`, `WithInstance`, or a URL or instance field on a credential config.
MUST NOT: Add functional options or a shared option bag.
MUST NOT: Add a token cache, refresh loop, or persistence to a direct provider.
MUST NOT: Require a client secret for JWT bearer.
MUST NOT: Put the client secret in an authorization URL unless `ClientSecretInAuthURL` is true.
MUST NOT: Validate external OIDC token signatures or claims in the SDK (R-AUTH-014).
MUST NOT: Fetch discovery, start a browser, or start a callback server in a constructor.
MUST NOT: Add implicit grant, device authorization, SAML/SSO, or cookie-session support without a superseding ADR.
MUST NOT: Declare package-local sentinels with `errors.New` (R-ERR-001).
MUST NOT: Reuse types from `credentials/` in the new package.
MUST NOT: Mark ROPC with a `Deprecated:` marker.
SHOULD: Keep token persistence opt-in and use an in-memory store by default.
SHOULD: Name the package `credential` and keep it in the existing module; do not add a submodule.
SHOULD: Ship a working vertical slice before adding the next flow; do not scaffold stubs (R-BUILD-007).
SHOULD: Document the HMAC custom-signer escape hatch for deployments with a different canonical string.
MAY: Export additional signer or callback hooks when a caller must supply that behavior.

## Verify

```bash
# The package does not exist yet. Once it does, fail on credential-owned URL
# options, interface-returning provider constructors, string Flow enums,
# exported credential internals, package-local sentinels, and functional
# options. Exit 0 while the package is absent. Exit 2 is a check failure.
[ -d credential ] || exit 0

check() {
  local pattern="$1" out code
  out=$(grep -RInE "$pattern" --include='*.go' --exclude='*_test.go' credential/)
  code=$?
  if [ "$code" -eq 0 ]; then
    printf '%s\n' "$out"
    return 1
  fi
  [ "$code" -eq 1 ] && return 0
  return "$code"
}

status=0
check 'func With(Instance|URL)' || status=$?
check 'func New[A-Za-z0-9_]*Provider\([^)]*\) \(authentication\.AuthenticationProvider' || status=$?
check '^[[:space:]]*Flow[[:space:]]+string' || status=$?
check 'type (BaseAccessTokenProvider|Authority|Credential)\b' || status=$?
check '^[[:space:]]*(var[[:space:]]+)?Err[A-Za-z0-9_]*[[:space:]]*=[[:space:]]*errors\.New\(' || status=$?
check '^[[:space:]]*type[[:space:]]+[A-Za-z0-9_]+Option[[:space:]]+func' || status=$?
if ! grep -RqE '^type Flow int$' --include='*.go' --exclude='*_test.go' credential/; then
  printf 'Flow must be an integer-backed enum\n'
  status=1
fi
exit "$status"
```

## Rationale

The old package duplicated lifecycle code, exposed credential types no external
caller could construct, and let callers name the instance twice. One typed
provider surface with a shared lifecycle keeps every flow consistent and leaves
the URL on the client. See ADR 022.
