# List table records

## Overview

Retrieves multiple records from the specified table.

## Path parameters

| Name      | Description                                                                                                                       |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------|
| `baseurl` | The absolute base URI for the request (this is the same for all requests) including: the schema, the domain, and a path fragment. |
| `table`   | The table name of for the operation.                                                                                              |

## Optional query parameters

| Name                       | Type                    | Possible values                | Description                                                                                                           |
|----------------------------|-------------------------|--------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| `DisplayValue`             | `tableapi.DisplayValue` | `TRUE`, `FALSE`, or `ALL`      | Determines the type of data returned, either the actual values from the database or the display values of the fields. |
| `ExcludeReferenceLink`     | `bool`                  | N/A                            | Flag that indicates whether to exclude Table API links for reference fields.                                          |
| `Fields`                   | `[]string `             | N/A                            | List of fields to include in the response.                                                                            |
| `QueryNoDomain`            | `bool`                  | N/A                            | Flag that indicates whether to restrict the record search to only the domains configured for the logged in user.      |
| `View`                     | `tableapi.View`         | `DESKTOP`, `MOBILE`, or `BOTH` | UI view for which to render the data.                                                                                 |
| `Limit`                    | `int`                   | N/A                            | Maximum number of records to return.                                                                                  |
| `NoCount`                  | `bool`                  | N/A                            | Flag that indicates whether to return the number of rows in the associated table.                                     |
| `Offset`                   | `int`                   | N/A                            | Starting record index for which to begin retrieving records.                                                          |
| `Query`                    | `string`                | N/A                            | Encoded query used to filter the result set.                                                                          |
| `QueryCategory`            | `string`                | N/A                            | Name of the category to use for queries.                                                                              |
| `SuppressPaginationHeader` | `bool`                  | N/A                            | Flag that indicates whether to remove the Link header from the response.                                              |

## Required query parameters

N/A

## Examples

=== "Fluent"

    ```go
    {%
    include-markdown 'assets/snippets/table-list-fluent.go'
    %}
    ```

=== "Standard"

    ```go
    {%
    include-markdown 'assets/snippets/table-list-standard.go'
    %}
    ```
