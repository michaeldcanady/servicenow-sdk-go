# Configuring the client

`NewServiceNowServiceClient` accepts functional options that control
authentication, the HTTP pipeline, logging, and model storage. This page
covers the options beyond authentication (for credential flows, see the
[Authentication guide](authentication/index.md)).

## Instance and URL

Point the client at your instance with one of:

```go
servicenowsdkgo.WithInstance("{instance}")
// or, for a full URL (e.g. behind a proxy):
servicenowsdkgo.WithURL("https://{instance}.service-now.com")
```

## The default HTTP pipeline

Out of the box the client uses Kiota's default middleware chain plus a
headers-inspection handler. That gives every request:

- **Automatic retries** with backoff for transient failures (including 429
  responses that carry `Retry-After`).
- **Redirect handling**, **compression**, and a standard **user-agent**.

You do not need to configure anything to get retry behavior.

## Custom middleware

`WithMiddleware` injects your own `nethttplibrary.Middleware`
implementations — for request logging, metrics, or custom headers:

```go
import nethttplibrary "github.com/microsoft/kiota-http-go"

middleware := append(
    nethttplibrary.GetDefaultMiddlewares(), // keep retries, redirects, compression
    &myLoggingMiddleware{},
)

client, err := servicenowsdkgo.NewServiceNowServiceClient(
    servicenowsdkgo.WithAuthenticationProvider(cred),
    servicenowsdkgo.WithInstance("{instance}"),
    servicenowsdkgo.WithMiddleware(middleware...),
)
```

:::warning
Supplying `WithMiddleware` **replaces** the default chain rather than
appending to it. If you still want retries, redirects, and compression,
start from `nethttplibrary.GetDefaultMiddlewares()` and append your own
handlers, as shown above.
:::

A middleware is any type implementing `Intercept`:

```go
type myLoggingMiddleware struct{}

func (m *myLoggingMiddleware) Intercept(
    pipeline nethttplibrary.Pipeline, middlewareIndex int, req *http.Request,
) (*http.Response, error) {
    start := time.Now()
    resp, err := pipeline.Next(req, middlewareIndex)
    log.Printf("%s %s -> %v (%s)", req.Method, req.URL.Path, err, time.Since(start))
    return resp, err
}
```

## Custom HTTP client

To control transport-level settings (timeouts, TLS, proxies), supply your own
`*http.Client`:

```go
servicenowsdkgo.WithHTTPClient(&http.Client{Timeout: 30 * time.Second})
```

## Logging

`WithLogger` attaches a logger the SDK uses for its internal diagnostics.
Pass any value with a `Log(message string, args ...interface{})` method:

```go
type stdLogger struct{}

func (stdLogger) Log(message string, args ...interface{}) {
    log.Printf(message, args...)
}

// ...
servicenowsdkgo.WithLogger(stdLogger{})
```

For request/response logging, prefer a custom middleware (above) — it sees
the actual HTTP traffic.

## Advanced options

- **`WithRequestAdapter`** — supply a fully custom Kiota
  `abstractions.RequestAdapter`, taking over serialization and transport
  entirely. The other pipeline options are ignored when you provide one.
- **`WithBackingStoreFactory`** — replace the backing store implementation
  used by models, for example to add change tracking.

## Next steps

- **[Handling errors](error-handling.md):** What comes back when a request fails.
- **[Authentication](authentication/index.md):** Credential flow options.
