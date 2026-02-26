# Create attachment file

## Overview

Upload file of any supported content type. Requires you to provide the table sys id, table name, and file name via **the request query parameters**.

## Path parameters

| Name      | Description                                                                                                                       |
| --------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `baseurl` | The absolute base URI for the request (this is the same for all requests) including: the schema, the domain, and a path fragment. |

## Optional query parameters

| Name                | Type      | Possible values | Description                                                  |
|---------------------|-----------|-----------------|--------------------------------------------------------------|
| `EncryptionContext` | `*string` | N/A             | `sys_id` of an encryption context record.                      |

## Required query parameters

| Name                | Type      | Possible values | Description                                                                       |
|---------------------|-----------|-----------------|-----------------------------------------------------------------------------------|
| `FileName`          | `*string` | N/A             | Name to provided file.                                                            |
| `TableName`         | `*string` | N/A             | Name of the designated table which contains the record to attach the file to.     |
| `TableSysID`        | `*string` | N/A             | Specifies the `sys_id` of the record in the designated table to attach the file to. |

## Examples

=== "Fluent"

    ```go
    package main

    {% include-markdown '../../../snippets/attachments.go' start='// [START attachment_imports]' end='// [END attachment_imports]' comments=false trailing-newlines=false dedent=true %}

    func main() {
        // Step 1: Create credentials
        {% include-markdown '../../../snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}

        // Step 2: Initialize client
        {% include-markdown '../../../snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}

        {% include-markdown '../../../snippets/attachments.go' start='// [START attachment_file_create]' end='// [END attachment_file_create]' comments=false trailing-newlines=false dedent=true %}
    }
    ```

=== "Standard"

    ```go
    package main

    {% include-markdown '../../../snippets/attachments.go' start='// [START attachment_imports]' end='// [END attachment_imports]' comments=false trailing-newlines=false dedent=true %}

    func main() {
        // Step 1: Create credentials
        {% include-markdown '../../../snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}

        // Step 2: Initialize client
        {% include-markdown '../../../snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}

        {% include-markdown '../../../snippets/attachments.go' start='// [START attachment_std_file_create]' end='// [END attachment_std_file_create]' comments=false trailing-newlines=false dedent=true %}
    }
    ```
