---
id: R-AUTH-016
applies_to: ["credentials/**/*.go", "website/docs/**/*auth*"]
status: accepted
supersedes: []
conflicts_checked: true
---

# ROPC stays available and is documented as not recommended

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Keep `NewROPCProvider` and its credential available.
MUST NOT: Remove the ROPC grant.
MUST: Give `NewROPCProvider` a godoc paragraph beginning `Not recommended:` that points to client credentials for machine identity and the authorization code flow for user identity.
MUST: Repeat the same recommendation in the authentication guide.
MUST NOT: Mark ROPC with a `// Deprecated:` marker (R-REL-025).

## Verify

```bash
# The doc comment must carry a Not recommended: paragraph and no Deprecated: marker. Fails closed if the function is missing.
doc=$(grep -RIn -B8 'func NewROPCProvider' --include='*.go' credentials/) && [ -n "$doc" ] \
  && printf '%s\n' "$doc" | grep -q 'Not recommended:' \
  && { printf '%s\n' "$doc" | grep -q 'Deprecated:'; [ $? -eq 1 ]; }
```

## Rationale

OAuth 2.1 drops ROPC, but some legacy ServiceNow setups have no substitute and the SDK has no removal plan. This SDK reserves `Deprecated:` for APIs being replaced or removed (R-REL-025), so ROPC gets an explicit not-recommended paragraph instead.
