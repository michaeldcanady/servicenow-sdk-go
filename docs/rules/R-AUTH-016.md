---
id: R-AUTH-016
applies_to: ["credentials/**/*.go", "website/docs/**/*auth*"]
status: accepted
supersedes: []
conflicts_checked: true
---

<!-- Keywords: ROPC, resource owner password, deprecated, removal, legacy,
guidance, discouraged, client credentials, authorization code -->

MUST keep `NewROPCProvider` and its credential available.
MUST recommend against the grant in the `NewROPCProvider` godoc and the authentication guide, pointing to client credentials for machine identity and the authorization code flow for user identity.
MUST NOT remove the ROPC grant.
MUST NOT mark ROPC with a `// Deprecated:` godoc marker — it implies a removal that is not planned.

## Verify

```bash
# Fails (non-zero) if a `// Deprecated:` marker is added to the ROPC provider.
! grep -RInE '// Deprecated:' --include='*.go' credentials/ | grep -i ropc
```

## Rationale

OAuth 2.1 deprecates ROPC, but some legacy ServiceNow setups have no substitute and the SDK has no removal plan; keeping it with documented guidance warns callers without promising a break that will not come (see ADR 016).
