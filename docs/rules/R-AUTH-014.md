---
id: R-AUTH-014
applies_to: ["credentials/**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# OIDC tokens are presented verbatim and validated only by ServiceNow

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- TODO(owner): decide behavior for opaque (non-JWT) IdP tokens, which have no parseable `exp`. -->

MUST: Send the IdP-issued token verbatim as a Bearer token.
MUST NOT: Validate the token's signature or claims in the SDK; ServiceNow's OIDC provider (OIDP) configuration record is the validator.
MUST: Parse `exp`, and no other claim, from the token to compute the cache TTL.
MUST: Return an error when the token has no parseable `exp`.
MUST NOT: Handle machine identity in the OIDC flow; the native client-credentials grant covers it.
MUST NOT: Add token validation unless a concrete IdP-compatibility problem requires it, recorded in a superseding rule.

## Verify

```bash
# `jwt.ParseUnverified` (exp extraction) is allowed; verifying parsers and JWKS libraries are not.
# grep exit 1 = no match = pass; exit 2 fails closed.
{ grep -RInE 'jwks|VerifySignature|jose\.Verify|jwt\.Parse\(|ParseWithClaims' --include='*.go' credentials/; [ $? -eq 1 ]; } \
  && { grep -E 'lestrrat-go/jwx|coreos/go-oidc|MicahParks/keyfunc' go.mod; [ $? -eq 1 ]; }
```

## Rationale

IdP claim layouts vary, so SDK-side validation either false-rejects configurations ServiceNow accepts or duplicates ServiceNow's rules and adds a JWKS dependency. The OIDP record is the single validator.
