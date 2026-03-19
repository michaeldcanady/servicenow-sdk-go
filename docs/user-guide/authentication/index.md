# Authentication

To interact with ServiceNow using the SDK, you must first authenticate your
client. The SDK supports multiple authentication methods to accommodate
different security requirements and deployment environments.

## Choose an authentication method

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

- **[Table Operations](../tables.md):** Interact with ServiceNow tables.
- **[Attachments](../attachments.md):** Upload and download file attachments.
- **[Batch API](../batch.md):** Execute multiple operations in a single request.
