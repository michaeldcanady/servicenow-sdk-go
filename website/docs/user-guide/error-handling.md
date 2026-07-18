# Handling errors

Every SDK operation returns an `error` as its last value. Errors fall into two
groups: **API errors** returned by ServiceNow (wrong credentials, missing
record, throttling) and **usage errors** raised by the SDK before a request is
sent (nil configuration, missing adapter).

## API errors

When ServiceNow responds with a failure status code, the SDK deserializes the
response body into a typed error that wraps the platform's `error.message`,
`error.detail`, and `status` fields. Match the type with `errors.As`:

```go
import (
    "errors"

    "github.com/michaeldcanady/servicenow-sdk-go/core"
)

response, err := client.Now().Table("incident").ByID("{SysID}").Get(ctx, nil)
if err != nil {
    var notFound *core.NotFoundError
    var unauthorized *core.UnauthorizedError

    switch {
    case errors.As(err, &notFound):
        // The record does not exist — treat as absent, not fatal.
    case errors.As(err, &unauthorized):
        // Credentials rejected — re-authenticate or fail fast.
    default:
        return err
    }
}
```

The mapped types, one per status class:

| Type | Status |
| ---- | ------ |
| `core.BadRequestError` | 400 |
| `core.UnauthorizedError` | 401 |
| `core.ForbiddenError` | 403 |
| `core.NotFoundError` | 404 |
| `core.TooManyRequestsError` | 429 |
| `core.ServerError` | 5XX |
| `core.ServiceNowError` | any other error status |

All of them embed `core.ServiceNowError`, so a single
`errors.As(err, &snErr)` with `var snErr *core.ServiceNowError` catches any
API error when you don't care which class it was.

### Reading the platform's error detail

`Error()` returns the platform's message, and the full payload is available
through `GetError()`:

```go
var snErr *core.ServiceNowError
if errors.As(err, &snErr) {
    mainErr, _ := snErr.GetError()
    message, _ := mainErr.GetMessage()
    detail, _ := mainErr.GetDetail()
    status, _ := mainErr.GetStatus()
    log.Printf("ServiceNow error: message=%v detail=%v status=%v",
        message, detail, status)
}
```

The getters return `*string`; check for `nil` before dereferencing — the
platform does not populate every field for every failure.

## Usage errors (sentinels)

Misusing the SDK — calling a method on a nil builder, passing a nil body —
returns a sentinel from the
[`errors` package](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go/errors).
Sentinels compare with `errors.Is`:

```go
import snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

if errors.Is(err, snerrors.ErrNilRequestBuilder) {
    // The builder chain was constructed from a nil client.
}
```

Commonly encountered sentinels:

- `ErrNilRequestBuilder` — a builder method was called on a nil builder,
  usually because client construction failed and the error was ignored.
- `ErrNilRequestAdapter` — the client has no request adapter; check the
  options passed to `NewServiceNowServiceClient`.
- `ErrNilBody` — a `Post`/`Put`/`Patch` was given a nil body.
- `ErrNilContext` — a nil `context.Context` was passed.

Compare sentinels by identity (`errors.Is`), never by matching the message
string — wording is not part of the compatibility contract.

## Retries and 429

The default HTTP pipeline already retries transient failures with backoff
(see [Configuring the client](configuration.md)). If you handle
`core.TooManyRequestsError` yourself, you are seeing a request that exhausted
those retries — back off at the workflow level rather than retrying
immediately.

## Next steps

- **[Configuring the client](configuration.md):** Middleware, retries, and logging.
- **[Table Operations](tables.mdx):** The requests these errors come from.
