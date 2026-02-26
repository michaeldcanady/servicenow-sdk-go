# Batch API Overview

The Batch API allows you to submit multiple REST requests in a single HTTP call. This is highly efficient for performing numerous operations while minimizing network overhead and reducing the number of round-trips to the ServiceNow instance.

Commonly used for:
- Executing multiple unrelated requests (e.g., creating an incident and updating a user).
- Reducing latency when performing many small operations.
- Ensuring a group of operations are sent to the server together.

## Basic Usage

Batching involves collecting multiple `RequestInformation` objects (from other APIs like Table or Attachment), wrapping them in a `BatchRequestModel`, and sending them via the Batch endpoint.

```go
client, _ := servicenowsdkgo.NewServiceNowClient2(credential, instance)

// 1. Prepare individual requests using ToRequestInformation methods
request1, _ := client.Now2().TableV2("incident").ToGetRequestInformation(context.Background(), nil)
request2, _ := client.Now2().TableV2("sys_user").ToGetRequestInformation(context.Background(), nil)

// 2. Create a batch request model
body := batchapi.NewBatchRequestModel()
// (Add requests to body - usually via a helper function)

// 3. Execute the batch
response, err := client.Now2().Batch().Post(context.Background(), body, nil)
```

## Available Operations

- [**Submit Batch Request**](create.md): Send a collection of requests to the batch processing endpoint.
