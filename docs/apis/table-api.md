# Table API

## Overview

The `Table API` provides endpoints that allow you to perform create, read, update, and delete (CRUD) operations on existing tables.

## \[GET\] /now/table/{tableName}

Retrieves multiple records for the specified table.

### Fluent implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    // define parameters you wish to (optional)
    params := *tableapi.TableRequestBuilderGetQueryParameters{
    ...
    }

    response, err := client.Now.Table("{TableName}").Get(params)
    if err != nil {
        log.Fatal(err)
    }
    // Handle response
    ...
}
```

### Standard implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.
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
        log.Fatal(err)
    }
}
```

## \[POST\] /now/table/{tableName}

Inserts one record in the specified table. 
> Note: Multiple record insertion is **not** supported by this method.

### Fluent implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    // define parameters you wish to (optional)
    params := *tableapi.TableRequestBuilderPostQueryParameters{
    ...
    }

    // data map of information you want to use for the new record
    data := map[string]string{
        "short_description": "example incident",
        "description": "incident created by servicenow-sdk-go",
    }

    response, err := client.Now.Table("{TableName}").Post(data, params)
    if err != nil {
        log.Fatal(err)
    }
    // Handle response
    ...
}
```

### Standard implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
        "table": "incident",
    }

     // define parameters you wish to (optional)
    params := *tableapi.TableRequestBuilderPostQueryParameters{
    ...
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
    response, err := requestBuilder.Post3(data, nil)

    // Test err, should be nil
    if err != nil {
        log.Fatal(err)
    }

    // Handle item response
    ...
}
```

## \[DELETE\] /now/table/{tableName}/{sys_id}

Deletes the specified record from the specified table.

### Fluent implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    // define parameters you wish to (optional)
    params := *tableapi.TableItemRequestBuilderDeleteQueryParameters{
    ...
    }

    response, err := client.Now.Table("{TableName}").Delete(params)
    if err != nil {
        log.Fatal(err)
    }
    // Handle response
    ...
}
```

### Standard Implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
        "table": "incident",
        "sysId": "INC00000000",
    }

    // define parameters you wish to (optional)
    params := *tableapi.TableItemRequestBuilderDeleteQueryParameters{
        ...
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := tableapi.NewTableItemRequestBuilder(client, pathParameters)

    // Call the delete method, with or without TableItemRequestBuilderDeleteQueryParameters.
    err := requestBuilder.Delete(params)

    // Since there is no record, the delete method only returns an error.
    if err != nil {
        log.Fatalf(err)
    }
    ...
}
```

## \[GET\] /now/table/{tableName}/{sys_id}

Retrieves the record identified by the specified sys_id from the specified table.

### Fluent implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    // define parameters you wish to (optional)
    params := *tableapi.TableItemRequestBuilderGetQueryParameters{
    ...
    }

    response, err := client.Now.Table("{TableName}").Get(params)
    if err != nil {
        log.Fatal(err)
    }
    // Handle response
    ...
}
```

### Standard implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
        "table": "incident",
        "sysId": "INC00000000",
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := tableapi.NewTableItemRequestBuilder(client, pathParameters)

   // define parameters you wish to (optional)
   params := *tableapi.TableItemRequestBuilderGetQueryParameters{
   ...
   }

    // Call the get method, with or without TableItemRequestBuilderGetQueryParameters.
    // record is of type TableItemResponse
    record, err := requestBuilder.Get(params)

    // evaluate err, should be nil
    if err != nil {
        log.Fatal(err)
    }
    // Handle response
    ...
}
```

## \[PUT\] /now/table/{tableName}/{sys_id}

### Fluent implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    // define parameters you wish to (optional)
    params := *tableapi.TableItemRequestBuilderGetQueryParameters{
    ...
    }

    // data map of information you want to use for the new record
    data := map[string]string{
        "short_description": "example incident",
        "description": "incident created by servicenow-sdk-go",
    }

    response, err := client.Now.Table("{TableName}").Put2(data, params)
    if err != nil {
        log.Fatal(err)
    }
    // Handle response
    ...
}
```

### Standard implementation

```golang
package main

import (
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    ... //Implement credential and client.

    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
        "table": "incident",
        "sysId": "INC00000000",
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := tableapi.NewTableItemRequestBuilder(client, pathParameters)

   // define parameters you wish to (optional)
   params := *tableapi.TableItemRequestBuilderPutQueryParameters{
   ...
   }

   // data map of information you want to use for the new record
   data := map[string]string{
       "short_description": "example incident",
       "description": "incident created by servicenow-sdk-go",
   }

   // Call the get method, with or without TableItemRequestBuilderGetQueryParameters.
   // record is of type TableItemResponse
   record, err := requestBuilder.Put2(data, params)

   // evaluate err, should be nil
   if err != nil {
       log.Fatal(err)
   }
   // Handle response
   ...
}
```