# Table API

## Overview

The `Table API` provides endpoints that allow you to perform create, read, update, and delete (CRUD) operations on existing tables.

## \[GET\] <code>/now/table/<var>{tableName}</var></code>

Retrieves multiple records for the specified table.

=== "Fluent"

    ``` golang
    package main

    import (
        "context"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        //Implement credential and client.
        //...

        // define parameters you wish to (optional)
        params := &tableapi.TableRequestBuilderGetQueryParameters{
        ...
        }

        ctx := context.Background()
        // Modify your context as desired.
        //...

        response, err := client.Now2().Table2("{TableName}").Get2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }
        // Handle response
        //...
    }
    ```

=== "Standard"

    ``` golang
    package main

    import (
        "context"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        //Implement credential and client.
        //...
        pathParameters := {
            "baseurl":"https://www.{instance}.service-now.com/api/now",
            "table": "incident",
        }

        // Instantiate new TableItemRequestBuilder.
        requestBuilder := tableapi.NewTableRequestBuilder2(client, pathParameters)

        // define parameters you wish to (optional)
        params := &tableapi.TableRequestBuilderGetQueryParameters{
        //...
        }

        ctx := context.Background()
        // Modify your context as desired.
        //...

        // Call the get method, with or without TableRequestBuilderGetQueryParameters.
        // Response is a TableCollectionResponse.
        response, err := requestBuilder.Get2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }
        // Handle response
        //...
    }
    ```

## \[POST\] <code>/now/table/<var>{tableName}</var></code>

Inserts one record in the specified table. 
> Note: Multiple record insertion is **not** supported by this method.

=== "Fluent"

    ``` golang
    package main

    import (
        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        //Implement credential and client.
        //...

        // define parameters you wish to (optional)
        params := &tableapi.TableRequestBuilderPostQueryParameters{
        //...
        }

        ctx := context.Background()
        // Modify your context as desired.
        //...

        // data map of information you want to use for the new record
        data := map[string]string{
            "short_description": "example incident",
            "description": "incident created by servicenow-sdk-go",
        }

        response, err := client.Now2().Table2("{TableName}").Post4(ctx, data, params)
        if err != nil {
            log.Fatal(err)
        }
        // Handle response
        //...
    }
    ```

=== "Standard"

    ``` golang
    package main

    import (
        "context"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        //Implement credential and client.
        //...

        pathParameters := {
            "baseurl":"https://www.{instance}.service-now.com/api/now",
            "table": "incident",
        }

        // define parameters you wish to (optional)
        params := &tableapi.TableRequestBuilderPostQueryParameters{
        //...
        }

        ctx := context.Background()
        // Modify your context as desired.
        //...

        // data map of information you want to use for the new record
        data := map[string]string{
            "short_description": "example incident",
            "description": "incident created by servicenow-sdk-go",
        }

        // Instantiate new TableItemRequestBuilder.
        requestBuilder := tableapi.NewTableRequestBuilder2(client, pathParameters)

        // Call the get method, with or without TableRequestBuilderPostQueryParamters.
        // Make sure you include the data parameter
        // Response is a TableItemResponse.
        response, err := requestBuilder.Post4(ctx, data, params)
        if err != nil {
            log.Fatal(err)
        }
        // Handle response
        //...
    }
    ```

## \[DELETE\] <code>/now/table/<var>{tableName}</var>/<var>{sys_id}</var></code>

Deletes the specified record from the specified table.

=== "Fluent"

    ``` golang
    package main

    import (
        "context"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        //Implement credential and client.
        //...

        // define parameters you wish to (optional)
        params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
        //...
        }

        ctx := context.Background()
        // Modify your context as desired.
        //...

        response, err := client.Now2().Table2("{TableName}").ByID("{sysID}").Delete2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }
        // Handle response
        //...
    }
    ```

=== "Standard"
        
    ``` golang
    package main

    import (
        "context"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        //Implement credential and client.
        //...

        pathParameters := {
            "baseurl":"https://www.{instance}.service-now.com/api/now",
            "table": "incident",
            "sysId": "INC00000000",
        }

        // define parameters you wish to (optional)
        params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
        //...
        }

        ctx := context.Background()
        // Modify your context as desired.
        //...

        // Instantiate new TableItemRequestBuilder.
        requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)

        // Call the delete method, with or without TableItemRequestBuilderDeleteQueryParameters.
        err := requestBuilder.Delete2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }
        // Handle response
        //...
    }
    ```

## \[GET\] <code>/now/table/<var>{tableName}</var>/<var>{sys_id}</var></code>

Retrieves the record identified by the specified sys_id from the specified table.

=== "Fluent"

    ``` golang
    package main

    import (
        "context"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        //Implement credential and client.
        //...

        // define parameters you wish to (optional)
        params := &tableapi.TableItemRequestBuilderGetQueryParameters{
        //...
        }

        ctx := context.Background()
        // Modify your context as desired.
        //...

        response, err := client.Now2().Table2("{TableName}").ByID("{sysID}").Get2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }
        // Handle response
        //...
    }
    ```

=== "Standard"

    ``` golang
    package main

    import (
        "context"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        //Implement credential and client.
        //...

        pathParameters := {
            "baseurl":"https://www.{instance}.service-now.com/api/now",
            "table": "incident",
            "sysId": "INC00000000",
        }

        // Instantiate new TableItemRequestBuilder.
        requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)

        // define parameters you wish to (optional)
        params := &tableapi.TableItemRequestBuilderGetQueryParameters{
        //...
        }

        ctx := context.Background()
        // Modify your context as desired.
        //...

        // Call the get method, with or without TableItemRequestBuilderGetQueryParameters.
        // record is of type TableItemResponse
        record, err := requestBuilder.Get2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }
        // Handle response
        //...
    }
    ```

## \[PUT\] <code>/now/table/<var>{tableName}</var>/<var>{sys_id}</var></code>

=== "Fluent"

    ``` golang
    package main

        import (
            "context"

            tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
        )

        func main() {
            //Implement credential and client.
            //...

            // define parameters you wish to (optional)
            params := &tableapi.TableItemRequestBuilderGetQueryParameters{
            //...
            }

            ctx := context.Background()
            // Modify your context as desired.
            //...

            // data map of information you want to use for the new record
            data := map[string]string{
                "short_description": "example incident",
                "description": "incident created by servicenow-sdk-go",
            }

            response, err := client.Now2().Table2("{TableName}").Put3(ctx, data, params)
            if err != nil {
                log.Fatal(err)
            }
            // Handle response
            //...
        }
    ```

=== "Standard"

    ``` golang
        package main

        import (
            "context"

            tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
        )

        func main() {
            //Implement credential and client.
            //...

            pathParameters := {
                "baseurl":"https://www.{instance}.service-now.com/api/now",
                "table": "incident",
                "sysId": "INC00000000",
            }

            // Instantiate new TableItemRequestBuilder.
            requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)

            // define parameters you wish to (optional)
            params := &tableapi.TableItemRequestBuilderGetQueryParameters{
            //...
            }

            ctx := context.Background()
            // Modify your context as desired.
            //...

            // data map of information you want to use for the new record
            data := map[string]string{
                "short_description": "example incident",
                "description": "incident created by servicenow-sdk-go",
            }

            // Call the get method, with or without TableItemRequestBuilderGetQueryParameters.
            // record is of type TableItemResponse
            record, err := requestBuilder.Put3(ctx, data, params)
            if err != nil {
                log.Fatal(err)
            }
            // Handle response
            //...
        }
    ```
