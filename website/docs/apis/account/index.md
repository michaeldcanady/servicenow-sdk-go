# Account API overview

The Account API provides read access to customer account records
(`customer_account`) used by Customer Service Management.

## Basic usage

Access it through the `Now()` namespace:

```go
// List accounts
accounts, err := client.Now().Account().Get(context.Background(), nil)

// Get a single account by its account ID
account, err := client.Now().Account().ByID("{accountID}").Get(context.Background(), nil)
```

## Available operations

- **List accounts** — `Account().Get(ctx, config)` returns an account collection.
- **Get account** — `Account().ByID(accountID).Get(ctx, config)` returns a single account.
