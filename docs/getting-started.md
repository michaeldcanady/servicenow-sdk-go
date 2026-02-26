# Getting started

Welcome to the ServiceNow SDK for Go! This guide helps you set up the SDK and
make your first successful API call in just a few minutes.

## Prerequisites

Before you begin, ensure your development environment meets these requirements:

- **Go:** Version 1.23 or higher installed.
- **ServiceNow Instance:** Access to a ServiceNow instance (a
  [Personal Developer Instance](https://developer.servicenow.com/) works
  perfectly).

## Install the SDK

Add the SDK to your existing Go module using the `go get` command:

```bash
go get github.com/michaeldcanady/servicenow-sdk-go
```

## Initialize the client

To interact with ServiceNow, you must first configure your credentials and
initialize a client. The following example demonstrates a basic setup using
username and password authentication.

```go
{% include-markdown 'snippets/auth.go' start='// [START auth_imports]' end='// [END auth_imports]' comments=false trailing-newlines=false dedent=true %}

func main() {
    // 1. Configure your credentials
    {% include-markdown 'snippets/auth.go' start='// [START auth_basic_admin]' end='// [END auth_basic_admin]' comments=false trailing-newlines=false dedent=true %}

    // 2. Initialize the client for your instance
    {% include-markdown 'snippets/auth.go' start='// [START client_init_panic]' end='// [END client_init_panic]' comments=false trailing-newlines=false dedent=true %}
}
```

## Make your first request

Once your client is initialized, you can perform operations like retrieving
records from a table. See the [Table Operations](user-guide/tables.md) guide
for more details.

## Next steps

Now that you have the basic setup, explore these topics to dive deeper:

- **[Authentication Guide](user-guide/authentication.md):** Learn about OAuth2
  and alternative credential types.
- **[Table Operations](user-guide/tables.md):** Master CRUD operations for
  ServiceNow records.
- **[Attachments](user-guide/attachments.md):** Manage files associated with
  your records.
- **[API Reference](apis/index.md):** Consult the detailed documentation for
  specific API modules.
