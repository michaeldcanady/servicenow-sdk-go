# ServiceNow SDK

A typed Go client for ServiceNow's REST APIs, built on Kiota abstractions.

## Language

**OIDP**:
ServiceNow's Application Registry record that accepts and validates identity tokens issued by an external identity provider on inbound API calls, using the provider's JWKS endpoint.
_Avoid_: OIDC provider, OAuth OIDC entity, external OAuth entity

**External IdP**:
The third-party identity provider, such as Microsoft Entra ID, that issues the identity token the SDK presents to ServiceNow.
_Avoid_: identity provider, third-party provider

**Token authority**:
The entity that issues and controls the token for a flow. For external OIDC authentication the token authority is the external IdP, not ServiceNow.
_Avoid_: authority, authorization server (reserve for ServiceNow's own OAuth server)

**Third-party token grant**:
ServiceNow's inbound OAuth mechanism, configured as an OIDP record, that authenticates API calls using an identity token from a trusted external IdP for a user account.
_Avoid_: external ID token auth, third-party OIDC record

**External OIDC authentication**:
Authenticating to ServiceNow by presenting an identity token signed and issued by an external IdP, rather than a token issued by ServiceNow's own OAuth server.
_Avoid_: federated auth, third-party OIDC login