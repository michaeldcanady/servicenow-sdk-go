# Delete attachment

## Overview

Delete specific attachment using the sys id.

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
        config := &attachmentapi.AttachmentItemRequestBuilderDeleteRequestConfiguration{
            //...
        }

        // Call the delete method with/without request configurations.
        err := client.Now().Attachment2().ByID("{sys_id}").Delete(context.Background(), config)

        // Test err, should be nil
        if err != nil {
            log.Fatal(err)
        }

        ...
    }
    ```
