# Welcome to ServiceNow SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/michaeldcanady/servicenow-sdk-go.svg)](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go)

The **ServiceNow SDK for Go** is a powerful, type-safe, and intuitive client library for interacting with ServiceNow REST APIs. Built on the modern Microsoft Kiota framework, it provides a fluent development experience tailored for Go developers.

## Key Features

- **Fluent API**: Discoverable and readable code structure.
- **Type Safety**: Leverages Go generics (V2) for compile-time checks.
- **Modular**: Only use the parts of the SDK you need (Table, Attachment, Batch).
- **Extensible**: Easy to add custom authentication or custom table models.
- **Middleware Support**: Built-in support for retries, logging, and more via Kiota.

## Quick Start

```go
import (
    "context"
    "fmt"
    "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func main() {
    cred := credentials.NewUsernamePasswordCredential("admin", "password")
    client, _ := servicenowsdkgo.NewServiceNowClient2(cred, "my-instance")

    ctx := context.Background()
    response, _ := client.Now2().TableV2("incident").Get(ctx, nil)

    for _, record := range response.GetValue() {
        fmt.Println(record.Get("number"))
    }
}
```

## How the Docs are Organized

- [**User Guide**](user-guide/getting-started.md): Practical tutorials for common tasks like authentication and CRUD operations.
- [**Preview Features**](user-guide/preview-features.md): Documentation for experimental features requiring build tags (e.g., Fluent Query Builder).
- [**API Reference**](apis/index.md): Detailed documentation of every supported ServiceNow API module.
- [**Contributor Guide**](contributing/index.md): Information for those looking to help improve the SDK.

---

*This project is community-driven and not an official ServiceNow product.*
