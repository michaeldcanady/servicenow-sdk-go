---
id: R-AUTH-015
applies_to: ["credentials/**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: credentials surface, providers, credential internal, BaseAccessTokenProvider,
Authority, WithInstance, WithURL, token store, provider constructors -->

MUST expose only provider constructors (`NewXProvider`), `TokenStore`, per-flow option types, and shared error sentinels from `credentials/`.
MUST keep the credential types, `BaseAccessTokenProvider`, and `Authority` unexported.
MUST name the ServiceNow URL once, on the client, passed to the provider via `Initialize`.
MUST NOT export credential types whose constructors take unexported interfaces no external caller can implement.
MUST NOT add `WithInstance` or `WithURL` back to credential options (the external OIDC issuer stays a credential option because it names the IdP, not the instance).

## Verify

```bash
# Fails (non-zero) if WithInstance/WithURL is reintroduced on credentials.
! grep -RInE 'func With(Instance|URL)\(' --include='*.go' credentials/
```

## Rationale

The exported credential types were API surface no consumer could construct, and the instance was set twice (credential + client); exposing providers only keeps the public surface usable and the URL single-sourced (see ADR 015).
