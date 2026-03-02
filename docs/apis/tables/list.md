# List table records

## Overview

Retrieves multiple records from the specified table.

## Path parameters

| Name      | Description                                                                                                                       |
| --------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `baseurl` | The absolute base URI for the request (this is the same for all requests) including: the schema, the domain, and a path fragment. |
| `table`   | The table name of for the operation.                                                                                              |

## Optional query parameters

| Name                       | Type                     | Possible values                                                                             | Description                                                                                                           |
| -------------------------- | ------------------------ | ------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `DisplayValue`             | `tableapi.DisplayValue2` | `tableapi.DisplayValue2True`, `tableapi.DisplayValue2False`, or `tableapi.DisplayValue2All` | Determines the type of data returned, either the actual values from the database or the display values of the fields. |
| `ExcludeReferenceLink`     | `bool`                   | N/A                                                                                         | Flag that indicates whether to exclude Table API links for reference fields.                                          |
| `Fields`                   | `[]string`               | N/A                                                                                         | List of fields to include in the response.                                                                            |
| `QueryNoDomain`            | `bool`                   | N/A                                                                                         | Flag that indicates whether to restrict the record search to only the domains configured for the logged in user.      |
| `View`                     | `tableapi.View2`         | `View2Desktop`, `View2Mobile`, or `View2Both`                                               | UI view for which to render the data.                                                                                 |
| `Limit`                    | `int`                    | N/A                                                                                         | Maximum number of records to return.                                                                                  |
| `NoCount`                  | `bool`                   | N/A                                                                                         | Flag that indicates whether to return the number of rows in the associated table.                                     |
| `Offset`                   | `int`                    | N/A                                                                                         | Starting record index for which to begin retrieving records.                                                          |
| `Query`                    | `string`                 | N/A                                                                                         | Encoded query used to filter the result set.                                                                          |
| `QueryCategory`            | `string`                 | N/A                                                                                         | Name of the category to use for queries.                                                                              |
| `SuppressPaginationHeader` | `bool`                   | N/A                                                                                         | Flag that indicates whether to remove the Link header from the response.                                              |

## Required query parameters

N/A

## Examples

=== "Fluent"

    ```go
    package main

    {% include-markdown '../../snippets/tables.go' start='// [START table_imports]' end='// [END table_imports]' comments=false trailing-newlines=false dedent=true %}

    func main() {
        // Step 1: Create credentials
        {% include-markdown '../../snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}
        // Step 2: Initialize client
        {% include-markdown '../../snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}
        {% include-markdown '../../snippets/tables.go' start='// [START table_list_fluent]' end='// [END table_list_fluent]' comments=false trailing-newlines=false dedent=true %}
    }
    ```

=== "Standard"

    ```go
    package main

    {% include-markdown '../../snippets/tables.go' start='// [START table_imports]' end='// [END table_imports]' comments=false trailing-newlines=false dedent=true %}

    func main() {
        // Step 1: Create credentials
        {% include-markdown '../../snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}
        // Step 2: Initialize client
        {% include-markdown '../../snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}
        {% include-markdown '../../snippets/tables.go' start='// [START table_collection_standard_setup]' end='// [END table_collection_standard_setup]' comments=false trailing-newlines=false dedent=true %}
        {% include-markdown '../../snippets/tables.go' start='// [START table_list_standard]' end='// [END table_list_standard]' comments=false trailing-newlines=false dedent=true %}
    }
    ```
