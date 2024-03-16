# Table API

The `Table API` is a powerful interface provided by ServiceNow, allowing developers to perform Create, Read, Update, and Delete (CRUD) operations on existing tables within a ServiceNow instance. This document provides a detailed guide on how to use the API endpoints with examples written in Go.

## Common Setup

Before you can interact with the API, you need to set up your environment. Here’s the common setup code that will be used in all the examples:

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    // Define the base URL of your ServiceNow instance
    baseURL := "https://www.{instance}.service-now.com/api/now"

    // Implement your credential and client.
    client := // Your client setup here
...

```

## \[DELETE\] Removes a Record

This method allows you to delete a specific record from a specified table.

[Try on Playground](https://go.dev/play/p/iPn4PiVi1gm)

```golang
...
    // Define the table and sysId of the record you want to delete.
    pathParameters := map[string]string{
        "baseurl": baseURL,
        "table":   "{tableName}",
        "sysId":   "{sysId}",
    }

    // Create a new TableItemRequestBuilder with your client and path parameters.
    requestBuilder := tableapi.NewTableItemRequestBuilder(client, pathParameters)

    // Optional Query Parameters
    params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
        QueryNoDomain: true // Default false
    }

    // Call the Delete method. You can pass nil if you don't have any query parameters.
    err := requestBuilder.Delete(params)

    // The Delete method only returns an error. If there's an error, it will be handled here.
    if err != nil {
        panic(err)
    }
}

```

## \[GET\] Multiple Records

This method retrieves multiple records from a specified table.

[Try on Playground](https://go.dev/play/p/Pn4npKdCGvU)

```golang
...
    // Define the table from which you want to retrieve records.
    pathParameters := map[string]string{
        "baseurl": baseURL,
        "table":   "{tableName}",
    }

    // Create a new TableRequestBuilder with your client and path parameters.
    requestBuilder := tableapi.NewTableRequestBuilder(client, pathParameters)

    // Optional Query Parameters
    params := &tableapi.TableRequestBuilderGetQueryParameters{
        DisplayValue:             true,
        ExcludeReferenceLink:     false,
        Fields:                   []string{},
        QueryNoDomain:            false,
        View:
        Limit:                    10,
        NoCount:                  true,
        Offset:                   20,
        Query:                    "...",
        QueryCategory:            "...",
        SuppressPaginationHeader: false,
    }

    // Call the Get method. You can pass nil if you don't have any query parameters.
    // The response will be a TableCollectionResponse.
    response, err := requestBuilder.Get(params)

    // Handle any errors.
    if err != nil {
        panic(err)
    }
}
```

## \[GET\] A Specific Record

This method retrieves a specific record identified by its sys_id from a specified table.

[Try on Playground](https://go.dev/play/p/gFlzIvA01ld)

```golang
...
    // Define the table and the sysId of the record you want to retrieve.
    pathParameters := map[string]string{
        "baseurl": baseURL,
        "table":   "{tableName}",
        "sysId":   "{sysId}",
    }

    // Create a new TableItemRequestBuilder with your client and path parameters.
    requestBuilder := tableapi.NewTableItemRequestBuilder(client, pathParameters)

    params := &tableapi.TableItemRequestBuilderGetQueryParameters{

    }

    // Call the Get method. You can pass nil if you don't have any query parameters.
    // The response will be a TableItemResponse.
    record, err := requestBuilder.Get(params)

    // Handle any errors.
    if err != nil {
        panic(err)
    }
}
```

## \[POST\] Create a Record

This method inserts a new record into a specified table. Note that this method does not support the insertion of multiple records.

[Try on Playground](https://go.dev/play/p/gbkVOBVivqr)

```golang
...
    // Define the table where you want to insert a new record.
    pathParameters := map[string]string{
        "baseurl": baseURL,
        "table":   "{tableName}",
    }

    // Define the data for the new record.
    data := map[string]string{
        "short_description": "example incident",
        "description":       "incident created by servicenow-sdk-go",
    }

    // Create a new TableRequestBuilder with your client and path parameters.
    requestBuilder := tableapi.NewTableRequestBuilder(client, pathParameters)

    param := &tableapi.TableRequestBuilderPostQueryParameters{

    }

    // Call the Post2 method with the data for the new record. You can pass nil if you don't have any query parameters.
    // The response will be a TableItemResponse.
    response, err := requestBuilder.Post2(data, param)

    // Handle any errors.
    if err != nil {
        panic(err)
    }
}
```

## \[PUT\] Update a Record

Update one record in the specified table.
*Note: Make sure only the fields you intend on updating are included*

[Try on Playground](https://go.dev/play/p/ZrGrIVfWd9I)

```golang
...
    pathParameters := {
        "baseurl": baseURL,
        "table": "{tableName}",
        "sysId": "{sysId}",
    }

    // data map of information you want to use for the new record
    data := map[string]string{
        "short_description": "example incident",
        "description": "incident created by servicenow-sdk-go",
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := tableapi.NewTableRequestBuilder(client, pathParameters)

    // Optional query parameters
    params := &tableapi.TableItemRequestBuilderPutQueryParameters{

    }

    // Call the get method, with or without TableRequestBuilderPutQueryParameters.
    // Make sure you include the data parameter
    // Response is a TableItemResponse.
    response, err := requestBuilder.Put2(data, params)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```
