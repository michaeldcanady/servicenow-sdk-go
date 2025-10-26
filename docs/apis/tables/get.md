# Retrieve a Single Record

## Overview

Retrieves a specific record by `sys_id` from the specified table.

## Path parameters

| Name      | Description                                                                                                                                 |
|-----------|---------------------------------------------------------------------------------------------------------------------------------------------|
| `baseurl` | The absolute base URI for the request (generally this is the same for all requests) including: the schema, the domain, and a path fragment. |
| `table`   | The table name of for the operation.                                                                                                        |
| `sysId`   | The `sys_id` of the record.                                                                                                                 |

## Optional query parameters

| Name                   | Type                  | Possible values                                           | Description                                                                                                               |
|------------------------|-----------------------|-----------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| `DisplayValue`         | tableapi.DisplayValue | `tableapi.TRUE`, `tableapi.FALSE`, or `tableapi.ALL`      | Determines the type of data returned, either the actual values from the database or the display values of the fields.     |
| `ExcludeReferenceLink` | bool                  |                                                           | Flag that indicates whether to exclude Table API links for reference fields.                                              |
| `Fields`               | []string              |                                                           | List of fields to include in the response.                                                                                |
| `QueryNoDomain`        | bool                  |                                                           | Flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured. |
| `View`                 | tableapi.View         | `tableapi.DESKTOP`, `tableapi.MOBILE`, or `tableapi.BOTH` | UI view for which to render the data.                                                                                     |

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
