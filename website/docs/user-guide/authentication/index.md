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
| A user signs in interactively, and the app cannot keep a secret (CLI tool, SPA) | [Authorization Code — Public](./auth-code-public.mdx) | Same flow secured with PKCE instead of a secret. |
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

Each page provides a **playbook‑style guide** describing:

- Required ServiceNow configuration  
- Required values for your application  
- How to initialize the SDK  

## Next steps

Once authenticated, you can begin performing operations:

- **[Table Operations](../tables/index.mdx):** Interact with ServiceNow tables.
- **[Attachments](../attachments/index.mdx):** Upload and download file attachments.
- **[Batch API](../batch.mdx):** Execute multiple operations in a single request.
