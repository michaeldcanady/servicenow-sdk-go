# Table API

The `Table API` is an essential interface provided by ServiceNow, enabling developers to perform Create, Read, Update, and Delete (CRUD) operations on the tables within a ServiceNow instance. This guide offers comprehensive instructions on utilizing the API endpoints, with code examples in Go.

Please ensure that any strings enclosed in curly braces `{}` are replaced with actual values before using the examples. The documentation demonstrates two approaches for interacting with the API:

1. Fluent Interface Pattern: This approach allows for chaining methods in a single expression.
2. Building Request: This method involves constructing the request step by step.

## Initial Setup

To begin using the API, you must first configure your environment. Below is the standard setup code utilized across all examples:

```golang
package main

import (
    tableapi "github.com/RecoLabs/servicenow-sdk-go/table-api"
)

func main() {
    // Define the base URL of your ServiceNow instance
    baseURL := "https://www.{instance}.service-now.com/api/now"

    // Implement your credential and client.
    client := // Your client setup here
...

```

## \[DELETE\] Remove Record

To remove a specific record from a table:

### 1. Fluent Interface Pattern

```golang
...
    
    // Optional Query Parameters
    params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
        // Define Query Parameters
    }

    // Execute the delete operation
    err := client.Now().Table2("{tableName}").ByID("{sysId}").Delete(params)
    if err != nil {
        panic(err)  // Error handling
    }
}
```

### 2. Building Request

[Try on Playground](https://go.dev/play/p/kiIt77rWHn7)

```golang
...
    // Specify the path parameters
    pathParameters := map[string]string{
        "baseurl": baseURL,
        "table":   "{tableName}", // Table name for the record
        "sysId":   "{sysId}", // Sys Id of the record
    }

    // Initialize a new request builder
    requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)

    // Optional Query Parameters
    params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
        // Define Query Parameters
    }

    // Execute the delete operation
    err := requestBuilder.Delete(params)
    if err != nil {
        panic(err) // Error handling
    }
}

```

## \[GET\] Retrieve Records

To fetch multiple records from a table:

### 1. Fluent Interface Pattern

```golang
...

    // Optional Query Parameters
    params := &tableapi.TableRequestBuilderGetQueryParameters{
        // Define Query Parameters
    }

    // Execute the retrieval operation
    response, err := client.Now().Table2("{tableName}").Get(params)
    if err != nil {
        panic(err) // Error handling
    }
}
```

### 2. Building Request

[Try on Playground](https://go.dev/play/p/3OtdlSuaPEv)

```golang
...
    // Specify the path parameters
    pathParameters := map[string]string{
        "baseurl": baseURL,
        "table":   "{tableName}", // Table name for the records
    }

    // Initialize a new request builder
    requestBuilder := tableapi.NewTableRequestBuilder2(client, pathParameters)

    // Optional Query Parameters
    params := &tableapi.TableRequestBuilderGetQueryParameters{
        // Define Query Parameters
    }

    // Execute the retrieval operation
    response, err := requestBuilder.Get(params)
    if err != nil {
        panic(err)  // Error handling
    }
}
```

## \[GET\] Retrieve Record

This method retrieves a specific record identified by its sys_id from a specified table.

### 1. Fluent Interface Pattern

```golang
...

    // Optional Query Parameters
    params := &tableapi.TableItemRequestBuilderGetQueryParameters{
        // Define Query Parameters
    }

    // Execute the retrieval operation
    record, err := client.Now().Table2("{tableName}")ByID("{sysId}").Get()
    if err != nil {
        panic(err) // Error handling
    }
}
```

### 2. Building Request

[Try on Playground](https://go.dev/play/p/UdZVsnPcPPH)

```golang
...
    // Specify the path parameters
    pathParameters := map[string]string{
        "baseurl": baseURL,
        "table":   "{tableName}", // Table name for the record
        "sysId":   "{sysId}", // Sys Id of the record
    }

    // Create a new TableItemRequestBuilder with your client and path parameters.
    requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)

    // Optional Query Parameters
    params := &tableapi.TableItemRequestBuilderGetQueryParameters{
        // Define Query Parameters
    }

    // Execute the retrieval operation.
    record, err := requestBuilder.Get(params)
    if err != nil {
        panic(err) // Error handling
    }
}
```

## \[POST\] Create a Record

This method inserts a new record into a specified table.

> *Note: that this method does not support the insertion of multiple records.*

### 1. Fluent Interface Pattern

```golang
...
    // Define the data for the new record
    data := map[string]string{
        "short_description": "example incident",
        "description":       "incident created by servicenow-sdk-go",
    }

    // Optional Query Parameters
    param := &tableapi.TableRequestBuilderPostQueryParameters{
        // Define Query Parameters
    }

    // Execute the creation operation.
    response, err := client.Now().Table2("{tableName}").Post(data, param)
    if err != nil {
        panic(err) // Error handling
    }
}
```

### 2. Building Request

[Try on Playground](https://go.dev/play/p/gYvf6NE0oxB)

```golang
...
    // Specify the path parameters
    pathParameters := map[string]string{
        "baseurl": baseURL,
        "table":   "{tableName}", // Table name for the record
    }

    // Define the data for the new record
    data := map[string]string{
        "short_description": "example incident",
        "description":       "incident created by servicenow-sdk-go",
    }

    // Create a new TableRequestBuilder with your client and path parameters.
    requestBuilder := tableapi.NewTableRequestBuilder2(client, pathParameters)

    // Optional Query Parameters
    param := &tableapi.TableRequestBuilderPostQueryParameters{
        // Define Query Parameters
    }

    // Execute the creation operation.
    response, err := requestBuilder.Post(data, param)
    if err != nil {
        panic(err) // Error handling
    }
}
```

## \[PUT\] Update a Record

Update one record in the specified table.
> *Note: Make sure only the fields you intend on updating are included*

### 1. Fluent Interface Pattern

```golang
...

    // Define the data for the new record
    data := map[string]string{
        "short_description": "example incident",
        "description": "incident created by servicenow-sdk-go",
    }

    // Optional query parameters
    params := &tableapi.TableItemRequestBuilderPutQueryParameters{
        // Define Query Parameters
    }

    // Execute the update operation.
    response, err := client.Now().Table2("{tableName}").ByID("{sysID}").Put(data, params)
    if err != nil {
        panic(err) // Error handling
    }
}
```

### 2. Building Request

[Try on Playground](https://go.dev/play/p/d_gFYT6MjCn)

```golang
...

    // Specify the path parameters
    pathParameters := {
        "baseurl": baseURL,
        "table": "{tableName}",
        "sysId": "{sysID}",
    }

    // Define the data for the new record
    data := map[string]string{
        "short_description": "example incident",
        "description": "incident created by servicenow-sdk-go",
    }

    // Create a new TableRequestBuilder with your client and path parameters.
    requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)

    // Optional query parameters
    params := &tableapi.TableItemRequestBuilderPutQueryParameters{
        // Define Query Parameters
    }

    // Execute the update operation.
    response, err := requestBuilder.Put(data, params)
    if err != nil {
        panic(err) // Error handling
    }
}
```
