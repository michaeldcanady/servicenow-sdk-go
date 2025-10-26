# Create attachment file

## Overview

Upload file of any supported content type. Requires you to provide the table sys id, table name, and file name via **the request headers**.

## Path parameters

N/A - doesn't support standard format.

## Optional query parameters

| Name                | Type      | Possible values | Description                                                  |
|---------------------|-----------|-----------------|--------------------------------------------------------------|
| `EncryptionContext` | `*string` | N/A             | Sys_id of an encryption context record.                      |

## Required query parameters

| Name                | Type      | Possible values | Description                                                                       |
|---------------------|-----------|-----------------|-----------------------------------------------------------------------------------|
| `FileName`          | `*string` | N/A             | Name to provided file.                                                            |
| `TableName`         | `*string` | N/A             | Name of the designated table which contains the record to attach the file to.     |
| `TableSysID`        | `*string` | N/A             | Specifies the sys_id of the record in the designated table to attach the file to. |

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

        // The content type of the file you want to upload
        dataContentType := "text/plain"
        // the byte content of the file
        data := []byte("this is example data")

        media := attachmentapi.NewMedia(dataContentType, data)

        // Define the required and optional configurations
        config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
            ...
            QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
                TableSysID: "INC00000001", // required
                TableName:  "incident", // required
                FileName:   "example.txt", // required
            }
            ...
        }

        // Call the post method with your content type, data, and request configurations.
        // Response is the uploaded file.
        response, err := client.Now().Attachment2().File().Post(context.Background(), media, config)

        // Test err, should be nil
        if err != nil {
            log.Fatal(err)
        }
        
        // Handle response
        ...
    }
    ```
