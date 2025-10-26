# Update table record

## Overview

Updates the specified record by `sys_id` in the specified table.

## Path parameters

| Name      | Description                                                                                                                       |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------|
| `baseurl` | The absolute base URI for the request (this is the same for all requests) including: the schema, the domain, and a path fragment. |
| `table`   | The table name of for the operation.                                                                                              |
| `sysId`   | The sys id of the table record.                                                                                                   |

## Optional query parameters

| Name                       | Type                    | Possible values                | Description                                                                                                           |
|----------------------------|-------------------------|--------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| `DisplayValue`             | `tableapi.DisplayValue` | `TRUE`, `FALSE`, or `ALL`      | Determines the type of data returned, either the actual values from the database or the display values of the fields. |
| `ExcludeReferenceLink`     | `bool`                  | N/A                            | Flag that indicates whether to exclude Table API links for reference fields.                                          |
| `Fields`                   | `[]string `             | N/A                            | List of fields to include in the response.                                                                            |
| `InputDisplayValue`        | `bool`                  | N/A                            | Flag that indicates whether to set field values using the display value or the actual value.                          |
| `QueryNoDomain`            | `bool`                  | N/A                            | Flag that indicates whether to restrict the record search to only the domains configured for the logged in user.      |
| `View`                     | `tableapi.View`         | `DESKTOP`, `MOBILE`, or `BOTH` | UI view for which to render the data.                                                                                 |

# Required query parameters

N/A

## Examples

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
