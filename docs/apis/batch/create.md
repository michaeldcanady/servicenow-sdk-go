# Create batch request

## Overview

Submits a `BatchRequest` containing all desired requests.

## Path parameters

N/A

## Optional query parameters

N/A

# Required query parameters

N/A

## Examples

=== "Fluent"

    ```go
    package main

    {% include-markdown '../../snippets/batch.go' start='// [START batch_imports]' end='// [END batch_imports]' comments=false trailing-newlines=false dedent=true %}

    {% include-markdown '../../snippets/batch.go' start='// [START batch_helper]' end='// [END batch_helper]' comments=false trailing-newlines=false dedent=true %}

    func main() {
        // Step 1: Create credentials
        {% include-markdown '../../snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}
        // Step 2: Initialize client
        {% include-markdown '../../snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}
        {% include-markdown '../../snippets/batch.go' start='// [START batch_create]' end='// [END batch_create]' comments=false trailing-newlines=false dedent=true %}
    }
    ```

=== "Standard"
    ```go
    package main

    {% include-markdown '../../snippets/batch.go' start='// [START batch_imports]' end='// [END batch_imports]' comments=false trailing-newlines=false dedent=true %}

    {% include-markdown '../../snippets/batch.go' start='// [START batch_helper]' end='// [END batch_helper]' comments=false trailing-newlines=false dedent=true %}

    func main() {
        // Step 1: Create credentials
        {% include-markdown '../../snippets/auth.go' start='// [START auth_basic]' end='// [END auth_basic]' comments=false trailing-newlines=false dedent=true %}
        // Step 2: Initialize client
        {% include-markdown '../../snippets/auth.go' start='// [START client_init]' end='// [END client_init]' comments=false trailing-newlines=false dedent=true %}
        {% include-markdown '../../snippets/batch.go' start='// [START batch_std_create]' end='// [END batch_std_create]' comments=false trailing-newlines=false dedent=true %}
    }
    ```
