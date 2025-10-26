# Delete attachment

## Overview

Delete specific attachment using the sys id.

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

        config := &attachmentapi.AttachmentItemRequestBuilderDeleteRequestConfiguration{
            // Optional configurations
        }

        err := client.Now2().Attachment2().ByID("{sys_id}").Delete(context.Background(), config)
        if err != nil {
            log.Fatal(err)
        }

        // Process response
    }
    ```
