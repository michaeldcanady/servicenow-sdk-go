# Batch Operations

The Batch API allows you to group multiple REST API calls into a single request. This is highly efficient for performing multiple operations while reducing network overhead.

## Creating a Batch Request

To perform multiple operations in one go:

```go
import (
    "context"
    "fmt"
    "log"
    "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
)

func main() {
    // ... client initialization ...

    ctx := context.Background()

    // Create a batch request
    batchRequest := batchapi.NewBatchRequest()

    // Add multiple requests to the batch
    // Each request needs a unique ID within the batch
    batchRequest.AddRequest("req-1", "GET", "/api/now/v1/table/incident?sysparm_limit=1", nil, nil)
    batchRequest.AddRequest("req-2", "GET", "/api/now/v1/table/task?sysparm_limit=1", nil, nil)

    response, err := client.Now2().Batch().Post(ctx, batchRequest, nil)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    // Process batch responses
    for _, resp := range response.GetServicedResponses() {
        fmt.Printf("Request ID: %s, Status: %d
", resp.GetId(), resp.GetStatus())
    }
}
```

## Batch Request Limits

Please note that ServiceNow instances usually have limits on the number of requests allowed in a single batch (default is 50). Check your instance's `glide.rest.batch.max_requests` property.
