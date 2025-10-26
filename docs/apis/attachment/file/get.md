# Get attachment file

## Overview

Retrieves an attachment's file by sys id.

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
        // Initialize credentials and client

        config := &attachmentapi.AttachmentItemRequestBuilderGetRequestConfiguration{
            // Optional configurations
        }

        response, err := client.Now().Attachment2().ByID("{sys_id}").Get(context.Background(), config)
        if err != nil {
            panic(err)
        }

        // Process response
    }
    ```
