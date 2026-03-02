# Welcome to ServiceNow SDK for Go

[![GoDoc](https://img.shields.io/static/v1?style=plastic&label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go)
![GitHub Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub release](https://img.shields.io/github/v/release/michaeldcanady/servicenow-sdk-go?style=plastic)

The **ServiceNow SDK for Go** is a powerful, type-safe, and intuitive client
library for interacting with ServiceNow REST APIs. Built on the modern
Microsoft Kiota framework, it provides a fluent development experience tailored
specifically for Go developers.

## Why use this SDK?

Building integrations with ServiceNow can be complex. This SDK simplifies that
process by providing:

- **Fluent API:** Write readable and discoverable code that matches the API
  hierarchy.
- **Strong typing:** Leverage Go generics (in V2 modules) for compile-time
  checks and improved IDE support.
- **Resilience:** Benefit from automatic retries and modular
  middleware for logging and error handling.
- **Extensibility:** Integrate custom authentication methods or
  specialized table models.

## Quick start

Get up and running with just a few lines of code:

```go
{% include-markdown 'snippets/tables.go' start='// [START table_imports]' end='// [END table_imports]' comments=false trailing-newlines=false dedent=true %}

func main() {
    {% include-markdown 'snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}
    {% include-markdown 'snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}
    ctx := context.Background()
    {% include-markdown 'snippets/tables.go' start='// [START table_list_guide]' end='// [END table_list_guide]' comments=false trailing-newlines=false dedent=true %}
}
```

## Explore the documentation

- **[Getting Started](getting-started.md):** Follow the quick start guide to
  install the SDK and make your first call.
- **[User Guide](user-guide/authentication.md):** Deep dive into core features
  like authentication, table operations, and file attachments.
- **[API Reference](apis/index.md):** Browse detailed technical documentation
  for every supported ServiceNow API module.
- **[Contributor Guide](contributing/index.md):** Learn how to help
  improve the SDK and become part of the community.

---

*This project is community-driven and isn't an official ServiceNow product.*
