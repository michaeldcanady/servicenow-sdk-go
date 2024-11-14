# Batch API

The `Batch API` provides an endpoint to perform multiple requests simultaneously. This is particularly useful when you need to make multiple independent requests and want to optimize network usage.

## \[POST\] /now/batch

Sends multiple REST API requests in a single HTTP call. This can significantly improve the performance of your application by reducing the number of HTTP requests.

> **NOTE**: For performance reasons, avoid including long-running requests and requests that retrieve large amounts of data.

### Code Example

[Try on Playground](https://go.dev/play/p/ZED4jOzmNbE).

```golang
package main

import (
    batchapi "github.com/RecoLabs/servicenow-sdk-go/batch-api"
    core "github.com/RecoLabs/servicenow-sdk-go/core"
    tableapi "github.com/RecoLabs/servicenow-sdk-go/table-api"
    servicenow "github.com/RecoLabs/servicenow-sdk-go"
)

// createTableDeleteRequest creates a delete request for a given table and sysId.
func createTableDeleteRequest(client *servicenow.ServiceNowClient, tableName, sysId string) (*core.RequestInformation, error) {
    pathParameters := map[string]string{
        "baseurl": client.BaseUrl + "/now",
        "table":   tableName,
    }

    requestBuilder := tableapi.NewTableRequestBuilder(client, pathParameters)
    return requestBuilder.ToDeleteRequestInformation2(nil)
}

// buildBatchrequest builds a batch request from a list of request information.
func buildBatchrequest(client *servicenow.ServiceNowClient, requests ...*core.RequestInformation)  (*batchapi.BatchRequest, error) {
    batchRequest := batchapi.NewBatchRequest(client)

    for _, request := range requests {
        if err := batchRequest.AddRequest(request, False); err != nil {
            return nil, err
        }
    }
    return batchRequest, nil
}

func main() {
    cred := credentials.NewUsernamePasswordCredential(username, password)
    client, err := servicenow.NewServiceNowClient2(cred, instance)
    if err != nil {
        fmt.Println("Error creating ServiceNow client:", err)
        return
    }
    
    deleteRequest, err := createTableDeleteRequest(client, "tableName", "sysId")
    if err != nil {
        fmt.Println("Error creating table delete request:", err)
        return
    }

    pathParameters := map[string]string{
        "baseurl": client.BaseUrl + "/now",
    }

    batchRequest, err := buildBatchrequest(client, deleteRequest)
    if err != nil {
        fmt.Println("Error building batch request:", err)
        return
    }

    builder := batchapi.NewBatchRequestBuilder(client, pathParameters)
    response, err := builder.Post(client, batchRequest)
    if err != nil {
        fmt.Println("Error posting batch request:", err)
        return
    }
    // Handle response
}
```

### Code Explanation

The code example above demonstrates how to use the Batch API to send multiple requests in a single HTTP call. Here's a step-by-step explanation:

1. `createTableDeleteRequest`: This function creates a delete request for a given table and sysId.
2. `buildBatchrequest`: This function builds a batch request from a list of request information.
3. `main`: This function is the entry point of the program. It creates a ServiceNow client, builds a delete request, adds it to a batch request, and sends the batch request.

### Error Handling

The code example includes error handling. If an error occurs at any point (for example, when creating the ServiceNow client or sending the batch request), the error will be printed to the console and the program will exit. This prevents the program from continuing with invalid data.

### See Also

- [ServiceNow API Documentation](https://developer.servicenow.com/dev.do#!/reference/api/paris/rest/c_TableAPI)
- [Go Documentation](https://golang.org/doc/)
