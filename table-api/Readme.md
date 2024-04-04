# Table API

The `Table API` provides endpoints that allow you to perform create, read, update, and delete (CRUD) operations on existing tables.

## \[DELETE\] /now/table/{tableName}/{sys_id}

Deletes the specified record from the specified table.

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
        "table": "incident",
        "sysId": "INC00000000",
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := tableapi.NewTableItemRequestBuilder(client, pathParameters)

    // Call the delete method, with or without TableItemRequestBuilderDeleteQueryParameters.
    err := requestBuilder.Delete(nil)

    // Since there is no record, the delete method only returns an error.
    if err != nil {
        panic(err)
    }
}
```

## \[GET\] /now/table/{tableName}

Retrieves multiple records for the specified table.

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
        "table": "incident",
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := tableapi.NewTableRequestBuilder(client, pathParameters)

    // Call the get method, with or without TableRequestBuilderGetQueryParameters.
    // Response is a TableCollectionResponse.
    response, err := requestBuilder.Get(nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[GET\] /now/table/{tableName}/{sys_id}

Retrieves the record identified by the specified sys_id from the specified table.

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
        "table": "incident",
        "sysId": "INC00000000",
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := tableapi.NewTableItemRequestBuilder(client, pathParameters)

    // Call the get method, with or without TableItemRequestBuilderGetQueryParameters.
    // record is of type TableItemResponse
    record, err := requestBuilder.Get(nil)

    // evaluate err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[POST\] /now/table/{tableName}

Inserts one record in the specified table. Multiple record insertion is not supported by this method.

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
        "table": "incident",
    }

    // data map of information you want to use for the new record
    data := map[string]string{
        "short_description": "example incident",
        "description": "incident created by servicenow-sdk-go",
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := tableapi.NewTableRequestBuilder(client, pathParameters)

    // Call the get method, with or without TableRequestBuilderPostQueryParamters.
    // Make sure you include the data paramter
    // Response is a TableItemResponse.
    response, err := requestBuilder.Post(data, nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```
