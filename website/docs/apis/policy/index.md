# Policy API overview

The Policy API manages Configuration Data Management policy input mappings:
create, resolve, and delete the mappings that bind policy inputs to
configuration data.

## Basic usage

Access it through the `Cdm()` namespace:

```go
mappings := client.Cdm().Policies().Mappings()

// Create a policy mapping
created, err := mappings.Post(context.Background(), nil)

// Resolve policy inputs
resolved := mappings.Inputs().Resolved()

// Delete a mapping
err = mappings.Delete(context.Background(), nil)
```

## Available operations

- **Create mapping** — `Policies().Mappings().Post(ctx, config)`.
- **Delete mapping** — `Policies().Mappings().Delete(ctx, config)`.
- **Resolved inputs** — `Policies().Mappings().Inputs().Resolved()` resolves policy input mappings.
