# Delete a record

## Overview

Deletes the specified record from the table.

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

        params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
            // Optional query parameters
        }

        ctx := context.Background()

        err := client.Now2().Table2("{TableName}").ByID("{sysID}").Delete2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }

        // Process result
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

        params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
            // Optional query parameters
        }

        ctx := context.Background()

        err := requestBuilder.Delete2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }

        // Process result
    }
    ```
