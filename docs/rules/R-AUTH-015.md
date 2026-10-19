---
id: R-AUTH-015
applies_to: ["credentials/**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# The credentials package exports providers, not credential types

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Export from `credentials/` only provider constructors (`NewXProvider`), `TokenStore`, per-flow option types, and shared error sentinels.
MUST: Keep credential types, `BaseAccessTokenProvider`, and `Authority` unexported.
MUST: Name the ServiceNow URL once, on the client, and pass it to each provider through `Initialize`.
MUST NOT: Export a credential type whose constructor takes an unexported interface that no external caller can implement.
MUST NOT: Define `WithInstance*` or `WithURL*` options on credentials.
MAY: Keep the external OIDC issuer as a credential option, because it names the IdP, not the instance.

## Verify

```bash
# grep exit 1 = no match = pass; exit 2 fails closed.
grep -RInE 'func With(Instance|URL)' --include='*.go' credentials/; [ $? -eq 1 ]
```

## Rationale

Exported credential types were API surface no consumer could construct, and the instance was set twice (credential and client). Exposing providers only keeps the public surface usable and the URL single-sourced.
