# Getting Started

## Requirements

- Installed [Go 1.23+](https://go.dev/doc/install).
- Active ServiceNow instance or [Developer Instance](https://developer.servicenow.com/dev.do).

## 1. Install the SDK

To add the SDK to your project:

```bash
go get github.com/michaeldcanady/servicenow-sdk-go
```

## 2. Initialize the Client

First, create a credential and then initialize the ServiceNow client.

```go
{% include-markdown 'snippets/auth.go' start='// [START auth_imports]' end='// [END auth_imports]' comments=false trailing-newlines=false dedent=true %}

func main() {
    // 1. Setup credentials (Basic Auth example)
    {% include-markdown 'snippets/auth.go' start='// [START auth_basic_admin]' end='// [END auth_basic_admin]' comments=false trailing-newlines=false dedent=true %}

    // 2. Initialize the client for your instance
    {% include-markdown 'snippets/auth.go' start='// [START client_init_panic]' end='// [END client_init_panic]' comments=false trailing-newlines=false dedent=true %}
}
```

## 3. What's Next?

Now that you have a client, you can start interacting with ServiceNow:

- [**Authentication Guide**](user-guide/authentication.md): Learn about OAuth2 and custom credentials.
- [**Table Operations**](user-guide/tables.md): Learn how to CRUD records in tables.
- [**Working with Attachments**](user-guide/attachments.md): Learn how to upload and download files.
- [**Batch Operations**](user-guide/batch.md): Learn how to group multiple requests.
- [**API Reference**](apis/index.md): Explore detailed documentation for each module.
