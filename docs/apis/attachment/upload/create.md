# Create attachment file

## Overview

Upload file of any supported content type. Requires you to provide the table sys id, table name, and file name via **the multipart form**.

## Path parameters

N/A - doesn't support standard implementation.

## Optional query parameters

N/A

## Required query parameters

N/A


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

        {% include-markdown '../../../snippets/attachments.go' start='// [START attachment_upload_create]' end='// [END attachment_upload_create]' comments=false trailing-newlines=false dedent=true %}
    }
    ```
