# Batch API

The Batch API lets you combine multiple REST API requests into a single HTTP
request. This is highly efficient for performing multiple operations while
minimizing network overhead and latency.

## Understand batch requests

A batch request consists of multiple individual sub-requests. ServiceNow
executes each sub-request and returns the results in a single combined response.

### Key benefits

- **Reduced Latency:** Fewer round-trips between your application and
  ServiceNow.
- **Improved Performance:** Faster execution of multiple related tasks.
- **Efficiency:** Better utilization of network resources.

## Create a batch request

To create a batch request, you first need to define the individual requests
using the SDK's request builders and then combine them into a
`BatchRequestModel`.

```go
{% include-markdown 'snippets/batch.go' start='// [START batch_create]' end='// [START batch_create]' comments=false trailing-newlines=false dedent=true %}
```

### Batch helper

Combining multiple requests can involve repetitive code. You can use a helper
function to streamline the process of adding requests to your batch model.

```go
{% include-markdown 'snippets/batch.go' start='// [START batch_helper]' end='// [START batch_helper]' comments=false trailing-newlines=false dedent=true %}
```

## Next steps

- **[Table Operations](tables.md):** Learn about the types of table requests
  you can batch.
- **[Attachments](attachments.md):** Learn about managing attachments, which
  can also be part of a batch request.
