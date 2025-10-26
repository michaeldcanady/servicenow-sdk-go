# Delete table record

## Overview

Deletes the specified record from the table.

## Path parameters

| Name      | Description                                                                                                                       |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------|
| `baseurl` | The absolute base URI for the request (this is the same for all requests) including: the schema, the domain, and a path fragment. |
| `table`   | The table name of for the operation.                                                                                              |
| `sysId`   | The sys id of the table record.                                                                                                   |

## Optional query parameters

| Name                   | Type   | Possible values | Description                                                                                                               |
|------------------------|--------|-----------------|---------------------------------------------------------------------------------------------------------------------------|
| `QueryNoDomain`        | `bool` | N/A             | Flag that indicates whether to restrict the record search to only the domains configured for the logged in user.          |

# Required query parameters

N/A

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
