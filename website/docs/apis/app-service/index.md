# Application Service API overview

The Application Service API manages CSDM application services: register or
create services and query service details from the CMDB.

## Basic usage

Access it through the CMDB namespace:

```go
appService := client.Now().Cmdb().AppService()

// Register an existing application service
resp, err := appService.Csdm().RegisterService().Post(context.Background(), registerRequest, nil)

// Find a service
found, err := appService.Csdm().FindService().Get(context.Background(), nil)
```

## Available operations

- **Create service** — `AppService().Create().Post(ctx, body, config)`.
- **Register service** — `Csdm().RegisterService().Post(ctx, body, config)`.
- **Find service** — `Csdm().FindService().Get(ctx, config)`.
- **Service details** — `Csdm().ByID(sysID).ServiceDetails()` reads details for a service.
- **Populate service** — `Csdm().ByID(sysID).PopulateService()` populates a service's mapping.
