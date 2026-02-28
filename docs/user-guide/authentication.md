# Authentication

To interact with ServiceNow using the SDK, you must first authenticate your
client. The SDK supports multiple authentication methods to suit different
security requirements and environment configurations.

## Choose an authentication method

ServiceNow provides several ways to authenticate. The SDK currently supports
Basic Authentication (Username/Password) and Token-based authentication (OAuth
2.0).

### Basic authentication

Basic authentication's the simplest method, requiring only a username and
password. It's suitable for automated scripts and internal tools where complex
OAuth flows might be unnecessary.

To use basic authentication, create a `UsernamePasswordCredential` with your
ServiceNow credentials.

```go
{% include-markdown 'snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}
```

### Token-based authentication

Token-based authentication uses OAuth 2.0 to provide a more secure way to
authenticate without sharing user credentials directly with the application.
This method is recommended for production environments.

To use token-based authentication, you must provide your Client ID, Client
Secret, and the Base URL of your ServiceNow instance.

```go
{% include-markdown 'snippets/auth.go' start='// [START auth_token]' end='// [END auth_token]' comments=false trailing-newlines=false dedent=true %}
```

## Initialize the client

After configuring your credentials, you must initialize the `ServiceNowClient`
to start making requests to your instance.

```go
{% include-markdown 'snippets/auth.go' start='// [START client_init]' end='// [START client_init]' comments=false trailing-newlines=false dedent=true %}
```

## Next steps

Now that you've authenticated your client, you can begin performing
operations:

- **[Table Operations](tables.md):** Learn how to interact with ServiceNow
  tables.
- **[Attachments](attachments.md):** Manage files associated with your records.
- **[Batch API](batch.md):** Combine multiple operations into a single
  request.
