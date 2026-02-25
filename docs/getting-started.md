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
import (
    "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func main() {
    // 1. Setup credentials (Basic Auth example)
    cred := credentials.NewUsernamePasswordCredential("admin", "password")

    // 2. Initialize the client for your instance
    client, err := servicenowsdkgo.NewServiceNowClient2(cred, "your-instance")
    if err != nil {
        panic(err)
    }
}
```

## 3. What's Next?

Now that you have a client, you can start interacting with ServiceNow:

- [**Authentication Guide**](user-guide/authentication.md): Learn about OAuth2 and custom credentials.
- [**Table Operations**](user-guide/tables.md): Learn how to CRUD records in tables.
- [**Working with Attachments**](user-guide/attachments.md): Learn how to upload and download files.
- [**Batch Operations**](user-guide/batch.md): Learn how to group multiple requests.
- [**API Reference**](apis/index.md): Explore detailed documentation for each module.
