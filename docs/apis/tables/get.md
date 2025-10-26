# Retrieve a Single Record

## Overview

Retrieves a specific record by `sys_id` from the specified table.

## Optional query parameters

## Examples

=== "Fluent"

    ```go
    package main

    import (
        "context"
        "log"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        // Initialize credentials and client
        // ...

        params := &tableapi.TableItemRequestBuilderGetQueryParameters{
            // Optional query parameters
        }

        ctx := context.Background()

        response, err := client.Now2().Table2("{TableName}").ByID("{sysID}").Get2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }

        // Process response
    }
    ```

=== "Standard"

    ```go
    package main

    import (
        "context"
        "log"

        tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
    )

    func main() {
        // Initialize credentials and client
        // ...

        pathParameters := map[string]string{
            "baseurl": "https://www.{instance}.service-now.com/api/now",
            "table":   "incident",
            "sysId":   "INC00000000",
        }

        requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)

        params := &tableapi.TableItemRequestBuilderGetQueryParameters{
            // Optional query parameters
        }

        ctx := context.Background()

        record, err := requestBuilder.Get2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }

        // Process record
    }
    ```
