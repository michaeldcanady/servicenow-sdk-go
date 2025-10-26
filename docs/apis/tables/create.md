# Create table record

## Overview

Creates a single record in the specified table.
> **Note:** This endpoint doesn't support bulk insertion.

## Path parameters

| Name      | Description                                                                                                                       |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------|
| `baseurl` | The absolute base URI for the request (this is the same for all requests) including: the schema, the domain, and a path fragment. |
| `table`   | The table name of for the operation.                                                                                              |

## Optional query parameters

<!-- vale Vale.Spelling = NO -->
| Name                   | Type                    | Possible values                | Description                                                                                                           |
|------------------------|-------------------------|--------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| `DisplayValue`         | `tableapi.DisplayValue` | `TRUE`, `FALSE`, or `ALL`      | Determines the type of data returned, either the actual values from the database or the display values of the fields. |
| `ExcludeReferenceLink` | `bool`                  | N/A                            | Flag that indicates whether to exclude Table API links for reference fields.                                          |
| `Fields`               | `[]string `             | N/A                            | List of fields to include in the response.                                                                            |
| `InputDisplayValue`    | `bool`                  | N/A                            | Flag that indicates whether to set field values using the display value or the actual value.                          |
| `View`                 | `tableapi.View`         | `DESKTOP`, `MOBILE`, or `BOTH` | UI view for which to render the data.                                                                                 |
<!-- vale Vale.Spelling = YES -->

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

        params := &tableapi.TableRequestBuilderPostQueryParameters{
            // Optional configurations
        }

        data := map[string]string{
            "short_description": "example incident",
            "description":       "incident created by servicenow-sdk-go",
        }

        response, err := client.Now2().Table2("{TableName}").Post4(context.Background(), data, params)
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

        pathParameters := map[string]string{
            "baseurl": "https://www.{instance}.service-now.com/api/now",
            "table":   "incident",
        }

        requestBuilder := tableapi.NewTableRequestBuilder2(client, pathParameters)

        params := &tableapi.TableRequestBuilderPostQueryParameters{
            // Optional configurations
        }

        data := map[string]string{
            "short_description": "example incident",
            "description":       "incident created by servicenow-sdk-go",
        }

        response, err := requestBuilder.Post4(context.Background(), data, params)
        if err != nil {
            log.Fatal(err)
        }

        // Process response
    }
    ```
