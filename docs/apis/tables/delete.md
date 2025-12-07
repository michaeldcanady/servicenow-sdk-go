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

        params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
            // Optional configurations
        }

        err := client.Now2().Table2("xSDK_SN_TABLEx").ByID("xSDK_SN_TABLE_SYS_IDx").Delete2(context.Background(), params)
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

        pathParameters := map[string]string{
            "baseurl": "https://www.xSDK_SN_INSTANCEx.service-now.com/api/now",
            "table":   "xSDK_SN_TABLEx",
            "sysId":   "xSDK_SN_TABLE_SYS_IDx",
        }

        requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)

        params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
            // Optional configurations
        }

        err := requestBuilder.Delete2(context.Background(), params)
        if err != nil {
            log.Fatal(err)
        }

        // Process result
    }
    ```
