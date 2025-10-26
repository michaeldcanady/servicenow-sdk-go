# Retrieve multiple records

## Overview

Retrieves multiple records from the specified table.

## Path parameters

| Name      | Description                                                                                                                       |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------|
| `baseurl` | The absolute base URI for the request (this is the same for all requests) including: the schema, the domain, and a path fragment. |
| `table`   | The table name of for the operation.                                                                                              |

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

        params := &tableapi.TableRequestBuilderGetQueryParameters{
            // Optional query parameters
        }

        ctx := context.Background()

        response, err := client.Now2().Table2("{TableName}").Get2(ctx, params)
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
        }

        requestBuilder := tableapi.NewTableRequestBuilder2(client, pathParameters)

        params := &tableapi.TableRequestBuilderGetQueryParameters{
            // Optional query parameters
        }

        ctx := context.Background()

        response, err := requestBuilder.Get2(ctx, params)
        if err != nil {
            log.Fatal(err)
        }

        // Process response
    }
    ```
