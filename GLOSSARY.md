# ServiceNow SDK

A typed Go client for ServiceNow's REST APIs, built on Kiota abstractions.

## Language

**Provider**:
The exported Kiota-facing object that authenticates a ServiceNow request. A provider either presents a token from an internal credential or authenticates each request directly.
_Avoid_: auth handler, credential (for the exported object)

**Credential**:
The internal token lifecycle for one OAuth flow. It acquires, caches, refreshes, persists, and revokes tokens; callers construct a provider, not a credential.
_Avoid_: provider (for the internal object), token provider

**Direct provider**:
A provider that authenticates each request without a token lifecycle: Basic, static Bearer, API key, or HMAC. Mutual TLS is a transport helper, not a provider.
_Avoid_: tokenless credential, stateless provider

**OIDP**:
ServiceNow's Application Registry record that validates an ID token or user access token issued by an external IdP on inbound API calls, using the provider's JWKS endpoint.
_Avoid_: OIDC provider, OAuth OIDC entity, external OAuth entity

**External IdP**:
The third-party identity provider, such as Microsoft Entra ID, that issues the token the SDK presents to ServiceNow.
_Avoid_: identity provider, third-party provider

**Token authority**:
The entity that issues and controls the token for a flow. For external OIDC authentication the token authority is the external IdP, not ServiceNow.
_Avoid_: authority, authorization server (reserve for ServiceNow's own OAuth server)

**Third-party token grant**:
ServiceNow's inbound OAuth mechanism, configured as an OIDP record, that authenticates API calls using an ID token or user access token from a trusted external IdP for a user account.
_Avoid_: external ID token auth, third-party OIDC record

**External OIDC authentication**:
Authenticating to ServiceNow by presenting a token issued by an external IdP — an ID token by default, or an audience-scoped user access token when configured — rather than a token issued by ServiceNow's own OAuth server.
_Avoid_: federated auth, third-party OIDC login