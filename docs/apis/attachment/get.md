# Get attachment

## Overview

Retrieve specific attachment metadata using the sys id.

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

        // define the configurations you wish to (optional)
        config := &attachmentapi.AttachmentItemRequestBuilderGetRequestConfiguration{
            //...
        }

        // Call the get method with/without request configurations.
        // Response is the attachment item.
        response, err := client.Now().Attachment2().ByID("{sys_id}").Get(context.Background(), config)

        // Test err, should be nil
        if err != nil {
            log.Fatal(err)
        }

        // Handle response
        ...
    }
    ```
