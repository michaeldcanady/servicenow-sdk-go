# CMDB Instance API overview

The CMDB Instance API performs CRUD operations on configuration items (CIs)
by CMDB class, and manages relationships between CIs.

## Basic usage

Access it through the CMDB namespace:

```go
instance := client.Now().Cmdb().Instance()

// Query CIs of a class
servers, err := instance.ByClass("cmdb_ci_linux_server").Get(context.Background(), nil)

// Get / update a specific CI
ci, err := instance.ByClass("cmdb_ci_linux_server").ByID("{sysID}").Get(context.Background(), nil)
updated, err := instance.ByClass("cmdb_ci_linux_server").ByID("{sysID}").Patch(context.Background(), ciBody, nil)
```

## Available operations

- **Query class instances** — `ByClass(className).Get(ctx, config)`.
- **Create instance** — `ByClass(className).Post(ctx, body, config)`.
- **Get instance** — `ByClass(className).ByID(sysID).Get(ctx, config)`.
- **Replace / update instance** — `...ByID(sysID).Put(ctx, body, config)` and `...Patch(ctx, body, config)`.
- **Create relation** — `...ByID(sysID).Relation().Post(ctx, body, config)`.
- **Delete relation** — `...Relation().ByID(relSysID).Delete(ctx, config)`.
