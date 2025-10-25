# Attachment api

## Overview

The Attachment api provides endpoints that allows you to upload and query file attachments.
You can upload or retrieve a single file with each request.

## \[Get\] <code>/now/attachment</code>

Returns the metadata for multiple attachments.

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

## \[Post\] <code>/now/attachment/file</code>

Upload file of any supported content type. Requires you to provide the table sys id, table name, and file name via **the request headers**.

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

## \[Post\] <code>/now/attachment/upload</code>

Upload file of any supported content type. Requires you to provide the table sys id, table name, and file name via **the multipart form**.

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

## \[Get\] <code>/now/attachment/<var>{sys_id}</var></code>

Retrieve specific attachment metadata using the sys id.

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

## \[Delete\] <code>/now/attachment/<var>{sys_id}</var></code>

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

## \[Get\] <code>/now/attachment/<var>{sys_id}</var>/file</code>

Retrieves a file with content using provided parameters.

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

        // Call the delete method with/without request configurations.
        // response is the file with its metadata
        response, err := client.Now().Attachment2().ByID("{sys_id}").Get(context.Background(), config)

        // Test err, should be nil
        if err != nil {
            panic(err)
        }

        // Handle response
        ...
    }
    ```