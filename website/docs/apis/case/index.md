# Case API overview

The Case API manages Customer Service Management (CSM) cases: create cases,
retrieve or update them, and read case activities and field values.

## Basic usage

This module hangs off the client root via the `CustomerService()` namespace:

```go
cases := client.CustomerService().Case()

// List cases
list, err := cases.Get(context.Background(), nil)

// Create a case
created, err := cases.Post(context.Background(), newCase, nil)

// Get and update a specific case
item, err := cases.ByID("{caseSysID}").Get(context.Background(), nil)
updated, err := cases.ByID("{caseSysID}").Put(context.Background(), caseBody, nil)
```

## Available operations

- **List cases** — `Case().Get(ctx, config)`.
- **Create case** — `Case().Post(ctx, body, config)`.
- **Get case** — `Case().ByID(id).Get(ctx, config)`.
- **Update case** — `Case().ByID(id).Put(ctx, body, config)`.
- **Case activities** — `Case().ByID(id).Activities()` reads a case's activity stream.
- **Field values** — `Case().FieldValues(fieldName).Get` and `Case().ByID(id).FieldValues(fieldName).Get` list valid values for a case field.
