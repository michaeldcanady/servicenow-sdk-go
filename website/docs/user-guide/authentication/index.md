# Authentication

To interact with ServiceNow using the SDK, you must first authenticate your
client. The SDK supports multiple authentication methods to accommodate
different security requirements and deployment environments.

## Choose an authentication method

Pick the flow that matches how your application runs:

| Your situation | Use | Why |
| -------------- | --- | --- |
| Server-to-server integration, no user present | [Client Credentials](./client-credentials.mdx) | The app authenticates as itself with a client ID and secret. |
| Server-to-server with key-based trust | [JWT Token](./jwt-token.mdx) | Signed assertions instead of a shared secret; best for non-interactive services with strict credential policies. |
| A user signs in interactively, and the app can keep a secret (server-side service, desktop app) | [Authorization Code — Private](./auth-code-private.mdx) | Standard OAuth2 with a confidential client. |
| A user signs in interactively, and the app can't keep a secret (CLI tool, SPA) | [Authorization Code — Public](./auth-code-public.mdx) | Same flow secured with PKCE instead of a secret. |
| External identity provider (Entra ID, Okta, Auth0, ADFS) | [External OIDC](./external-oidc.mdx) | Federated identity via OIDC discovery; Authorization Code + PKCE against the IdP. |
| Legacy or tightly controlled environment where interactive login is impossible | [ROPC](./ropc.mdx) | Exchanges a username/password for OAuth tokens; use only when the flows above are unavailable. |
| Local development, testing, quick scripts | [Basic](./basic.mdx) | Simplest to set up; avoid for production integrations. |

ServiceNow provides multiple OAuth and non‑OAuth authentication models.  
The SDK currently supports:

- [Basic Authentication](./basic.mdx)
- [Authorization Code (Private)](./auth-code-private.mdx)
- [Authorization Code (Public)](./auth-code-public.mdx)
- [Client Credentials](./client-credentials.mdx)
- [JWT Token](./jwt-token.mdx)
- [Resource Owner Password Credential (ROPC)](./ropc.mdx)
- [External OIDC](./external-oidc.mdx)

Each page provides a **playbook‑style guide** describing:

- Required ServiceNow configuration  
- Required values for your application  
- How to initialize the SDK  

## Instance URL is set once on the client

The ServiceNow instance URL is configured **once on the client**, not on the
authentication provider. All authentication providers are initialized lazily
when the client makes its first request:

```go
import (
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2"
    "github.com/michaeldcanady/servicenow-sdk-go/v2/credentials"
)

cred, err := credentials.NewClientCredentialsProvider(
    "client-id",
    "client-secret",
    credentials.ClientCredentialsOptions{
        CommonOptions: credentials.CommonOptions{
            AllowedHosts: []string{"myinstance.service-now.com"},
        },
    },
)
if err != nil {
    log.Fatal(err)
}

client, err := servicenowsdkgo.NewServiceNowServiceClient(
    servicenowsdkgo.WithAuthenticationProvider(cred),
    servicenowsdkgo.WithInstance("myinstance.service-now.com"),
)
if err != nil {
    log.Fatal(err)
}
```

The instance URL is passed **once** via `WithInstance` (or `WithURL`) on the
client. All authentication providers are initialized lazily when the first
request is made.

## Token persistence and lifecycle

The SDK supports optional token persistence through the `TokenStore` interface.
When a token store is configured, tokens are loaded before acquisition, saved
after successful acquisition or refresh, and deleted on revocation. This
enables token reuse across process restarts without re-authentication.

```go
cred, err := credentials.NewClientCredentialsProvider(
    "client-id",
    "client-secret",
    credentials.ClientCredentialsOptions{
        CommonOptions: credentials.CommonOptions{
            AllowedHosts: []string{"myinstance.service-now.com"},
            TokenStore:   myTokenStore, // implements credentials.TokenStore
        },
    },
)
```

## ROPC is available but discouraged

The ROPC flow is available but strongly discouraged. Per ADR-016, ROPC sends
user credentials directly to the client application, which expands the attack
surface, defeats MFA, and violates modern OAuth2 best practices. Prefer
Authorization Code with PKCE (public client) or Client Credentials for
machine-to-machine scenarios.

## Next steps

Once authenticated, you can begin performing operations:

- **[Table Operations](../tables/index.mdx):** Interact with ServiceNow tables.
- **[Attachments](../attachments/index.mdx):** Upload and download file attachments.
- **[Batch API](../batch.mdx):** Execute multiple operations in a single request.