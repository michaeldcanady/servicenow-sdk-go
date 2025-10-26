# Update a Record

## Overview

Updates the specified record by `sys_id` in the specified table.

## Optional query parameters

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
