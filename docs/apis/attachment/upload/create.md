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

    ``` golang
    package main

    import (
        "context"

        attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
    )

    func main() {
        // Initialize credentials and client

        body := // TODO: how to make multipart body?

        config := &attachmentapi.AttachmentUploadRequestBuilderPostRequestConfiguration{
            // Optional configurations
        }

        response, err := client.Now2().Attachment2().Upload().Post(context.Background(), body, config)
        if err != nil {
            log.Fatal(err)
        }

        // Process response
    }
    ```
