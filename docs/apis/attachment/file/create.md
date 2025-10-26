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
        // Initialize credentials and client

        //build media type
        dataContentType := "text/plain"
        data := []byte("this is example data")
        media := attachmentapi.NewMedia(dataContentType, data)

        config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
            QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
                TableSysID: "INC00000001", // required
                TableName:  "incident", // required
                FileName:   "example.txt", // required
            }
            // Optional configurations
        }

        response, err := client.Now2().Attachment2().File().Post(context.Background(), media, config)
        if err != nil {
            log.Fatal(err)
        }
        
        // Process response
    }
    ```
