# CDM Editor API overview

The CDM Editor API edits Configuration Data Management data nodes and
validates configuration data.

## Basic usage

Access it through the `Cdm()` namespace:

```go
editor := client.Cdm().Editor()

// List nodes
nodes, err := editor.Nodes().Get(context.Background(), nil)

// Create, update, delete a node
created, err := editor.Nodes().Post(context.Background(), createRequest, nil)
updated, err := editor.Nodes().ByID("{nodeID}").Put(context.Background(), updateRequest, nil)
deleted, err := editor.Nodes().ByID("{nodeID}").Delete(context.Background(), nil)

// Validate configuration data
validation, err := editor.Validation().Get(context.Background(), nil)
```

## Available operations

- **List nodes** — `Nodes().Get(ctx, config)`.
- **Create node** — `Nodes().Post(ctx, body, config)`.
- **Update node** — `Nodes().ByID(id).Put(ctx, body, config)`.
- **Delete node** — `Nodes().ByID(id).Delete(ctx, config)`.
- **Validate** — `Validation().Get(ctx, config)`.
