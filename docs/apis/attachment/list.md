# List attachments

## Overview

Returns the metadata for multiple attachments.

## Path parameters

| Name      | Description                                                                                                                       |
| --------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `baseurl` | The absolute base URI for the request (this is the same for all requests) including: the schema, the domain, and a path fragment. |

## Optional query parameters

| Name            | Type     | Possible values | Description                                                  |
| --------------- | -------- | --------------- | ------------------------------------------------------------ |
| `SysParmLimit`  | `int`    | N/A             | Maximum number of records to return.                         |
| `SysParmOffset` | `int`    | N/A             | Starting record index for which to begin retrieving records. |
| `SysParmQuery`  | `string` | N/A             | Encoded query used to filter the result set.                 |

## Required query parameters

N/A

## Examples

=== "Fluent"

    ```go
    package main

    {% include-markdown '../../snippets/attachments.go' start='// [START attachment_imports]' end='// [END attachment_imports]' comments=false trailing-newlines=false dedent=true %}

    func main() {
        // Step 1: Create credentials
        {% include-markdown '../../snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}

        // Step 2: Initialize client
        {% include-markdown '../../snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}

        {% include-markdown '../../snippets/attachments.go' start='// [START attachment_list]' end='// [END attachment_list]' comments=false trailing-newlines=false dedent=true %}
    }
    ```

=== "Standard"

    ```go
    package main

    {% include-markdown '../../snippets/attachments.go' start='// [START attachment_imports]' end='// [END attachment_imports]' comments=false trailing-newlines=false dedent=true %}

    func main() {
        // Step 1: Create credentials
        {% include-markdown '../../snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}

        // Step 2: Initialize client
        {% include-markdown '../../snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}

        {% include-markdown '../../snippets/attachments.go' start='// [START attachment_std_list]' end='// [END attachment_std_list]' comments=false trailing-newlines=false dedent=true %}
    }
    ```
