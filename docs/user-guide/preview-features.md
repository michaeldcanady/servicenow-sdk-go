# Preview Features

The ServiceNow SDK for Go includes experimental features that are currently in "Preview" status. These features are available for early testing and feedback, but they are not yet considered stable and may undergo breaking changes in future releases.

To prevent accidental use in production environments, preview features are protected by **Go Build Tags**.

## How to use Preview Features

To enable a preview feature, you must include the corresponding build tag when compiling or running your Go application.

### Using `go build`

```bash
go build -tags <tag_name> .
```

### Using `go run`

```bash
go run -tags <tag_name> main.go
```

## Available Preview Features

| Feature | Build Tag | Description |
|---------|-----------|-------------|
| **Fluent Query Builder** | `preview.query` | A type-safe, fluent API for constructing complex ServiceNow queries. |

---

!!! warning "Stability Notice"
    Preview features are subject to change. We do not recommend using them in critical production workflows until they reach "Supported" status. We welcome your feedback and bug reports on these features!
