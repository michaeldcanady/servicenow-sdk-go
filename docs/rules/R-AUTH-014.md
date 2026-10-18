---
id: R-AUTH-014
applies_to: ["credentials/**/*.go", "**/*oidc*.go", "**/*auth*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: oidc, external identity, idp, token validation, jwks, claims,
aud, bearer, verbatim, oidp, pkce -->

MUST present the IdP-issued token verbatim as a Bearer token.
MUST perform no signature or claims validation in the SDK — ServiceNow's OIDP record is the validator.
MUST sanity-check only that the token has a parseable `exp` (so the cache can compute TTL).
MUST NOT treat machine identity as an OIDP concern — that is the native client-credentials grant's job.
MUST NOT add token validation unless a concrete IdP-compatibility problem demands it.

## Verify

```bash
# Fails (non-zero) if client-side token validation (JWKS / signature verify) is added.
! grep -RInE 'jwks|VerifySignature|jose\.Verify' --include='*.go' credentials/ auth/
```

## Rationale

IdP claim layouts vary, so SDK-side validation either false-rejects configurations ServiceNow accepts or duplicates ServiceNow's rules and adds a JWKS dependency; the OIDP record is the validator (see ADR 014).
