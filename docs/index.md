# Welcome to ServiceNow SDK for Go

[![GoDoc](https://img.shields.io/static/v1?style=plastic&label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go)
![GitHub Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub release](https://img.shields.io/github/v/release/michaeldcanady/servicenow-sdk-go?style=plastic)

The **ServiceNow SDK for Go** is a powerful, type-safe, and intuitive client library for interacting with ServiceNow REST APIs. Built on the modern Microsoft Kiota framework, it provides a fluent development experience tailored for Go developers.

## Key Features

- **Fluent API**: Discoverable and readable code structure.
- **Type Safety**: Leverages Go generics (V2) for compile-time checks.
- **Modular**: Only use the parts of the SDK you need (Table, Attachment, Batch).
- **Extensible**: Easy to add custom authentication or custom table models.
- **Middleware Support**: Built-in support for retries, logging, and more via Kiota.

## Quick Start

```go
{% include-markdown 'snippets/tables.go' start='// [START table_imports]' end='// [END table_imports]' comments=false trailing-newlines=false dedent=true %}

func main() {
    {% include-markdown 'snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}
    {% include-markdown 'snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}
    ctx := context.Background()
    {% include-markdown 'snippets/tables.go' start='// [START table_list_guide]' end='// [END table_list_guide]' comments=false trailing-newlines=false dedent=true %}
}

```

## How the Docs are Organized

- [**User Guide**](user-guide/getting-started.md): Practical tutorials for common tasks like authentication and CRUD operations.
- [**Preview Features**](user-guide/preview-features.md): Documentation for experimental features requiring build tags (e.g., Fluent Query Builder).
- [**API Reference**](apis/index.md): Detailed documentation of every supported ServiceNow API module.
- [**Contributor Guide**](contributing/index.md): Information for those looking to help improve the SDK.

---

*This project is community-driven and not an official ServiceNow product.*
