# Authentication

To interact with ServiceNow using the SDK, you must first authenticate your
client. The SDK supports multiple authentication methods to accommodate
different security requirements and deployment environments.

## Choose an authentication method

Pick the flow that matches how your application runs:

| Your situation | Use | Why |
| -------------- | --- | --- |
| Server-to-server integration, no user present | [Client Credentials](./client-credentials.md) | The app authenticates as itself with a client ID and secret. |
| Server-to-server with key-based trust | [JWT Token](./jwt-token.md) | Signed assertions instead of a shared secret; best for non-interactive services with strict credential policies. |
| A user signs in interactively, and the app can keep a secret (server-side service, desktop app) | [Authorization Code — Private](./auth-code-private.md) | Standard OAuth2 with a confidential client. |
| A user signs in interactively, and the app cannot keep a secret (CLI tool, SPA) | [Authorization Code — Public](./auth-code-public.md) | Same flow secured with PKCE instead of a secret. |
| Legacy or tightly controlled environment where interactive login is impossible | [ROPC](./ropc.md) | Exchanges a username/password for OAuth tokens; use only when the flows above are unavailable. |
| Local development, testing, quick scripts | [Basic](./basic.md) | Simplest to set up; avoid for production integrations. |

ServiceNow provides multiple OAuth and non‑OAuth authentication models.  
The SDK currently supports:

- [Basic Authentication](./basic.md)
- [Authorization Code (Private)](./auth-code-private.md)
- [Authorization Code (Public)](./auth-code-public.md)
- [Client Credentials](./client-credentials.md)
- [JWT Token](./jwt-token.md)
- [Resource Owner Password Credential (ROPC)](./ropc.md)

Each page provides a **playbook‑style guide** describing:

- Required ServiceNow configuration  
- Required values for your application  
- How to initialize the SDK  

## Next steps

Once authenticated, you can begin performing operations:

- **[Table Operations](../tables.mdx):** Interact with ServiceNow tables.
- **[Attachments](../attachments/index.mdx):** Upload and download file attachments.
- **[Batch API](../batch.mdx):** Execute multiple operations in a single request.
