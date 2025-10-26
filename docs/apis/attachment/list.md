# List attachments

## Overview

Returns the metadata for multiple attachments.

## Path parameters

N/A - doesn't support standard implementation.

## Optional query parameters

| Name            | Type     | Possible values | Description                                                  |
|-----------------|----------|-----------------|--------------------------------------------------------------|
| `SysParmLimit`  | `int`    | N/A             | Maximum number of records to return.                         |
| `SysParmOffset` | `int`    | N/A             | Starting record index for which to begin retrieving records. |
| `SysParmQuery`  | `string` | N/A             | Encoded query used to filter the result set.                 |

## Required query parameters

N/A

## Examples

=== "Fluent"

    ``` golang
    package main

    import (
        "context"

        attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
    )

    func main() {
        // Initialize credentials and client

        config := &attachmentapi.AttachmentRequestBuilder2GetRequestConfiguration{
           // Optional configurations
        }

        response, err := client.Now2().Attachment2().Get(context.Background(), config)
        if err != nil {
            log.Fatal(err)
        }

        // Process response
    }
    ```
