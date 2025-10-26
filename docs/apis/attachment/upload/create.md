# Create attachment file

## Overview

Upload file of any supported content type. Requires you to provide the table sys id, table name, and file name via **the multipart form**.

## Path parameters

N/A - doesn't support standard format.

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
        //Implement credential and client.
        ...

        body := // TODO: how to make multipart body?

        // define the configurations you wish to (optional)
        config := &attachmentapi.AttachmentUploadRequestBuilderPostRequestConfiguration{
            //...
        }

        // Call the post method with your content type, data, and request configurations.
        // Response is the uploaded file.
        response, err := client.Now().Attachment2().Upload().Post(context.Background(), body, config)

        // Test err, should be nil
        if err != nil {
            log.Fatal(err)
        }

        // Handle response
        ...
    }
    ```
