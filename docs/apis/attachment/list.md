# List attachments

## Overview

Returns the metadata for multiple attachments.

## Path parameters

N/A - doesn't support standard format.

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
        //Implement credential and client.
        ...
    
        // define the configurations you wish to (optional)
        config := &attachmentapi.AttachmentRequestBuilder2GetRequestConfiguration{
           // ...
        }

        // Call the get method, with or without AttachmentRequestBuilderGetQueryParameters.
        // Response is a AttachmentCollectionResponse.
        response, err := client.Now().Attachment2().Get(context.Background(), config)

        // Test err, should be nil
        if err != nil {
            log.Fatal(err)
        }

        // Handle response
        ...
    }
    ```
