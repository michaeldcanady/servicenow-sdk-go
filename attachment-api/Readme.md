# Attachment API

The Attachment API provides endpoints that allow you to upload and query file attachments.
You can upload or retrieve a single file with each request.

## \[GET\] /now/attachment

Returns the metadata for multiple attachments.

### V1 client compatible

```golang
package main

import (
    "context"

    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
    }

    // Instantiate new attachment request builder.
    requestBuilder := attachmentapi.NewV1CompatibleAttachmentRequestBuilder2(pathParameters, client)

    // Call the get method, with or without AttachmentRequestBuilderGetQueryParameters.
    // Response is a AttachmentCollectionResponse.
    response, err := requestBuilder.Get(context.Background(), nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[POST\] /now/attachment/file

Upload file of any supported content type.

### V1 client compatible

```golang
package main

import (
    "context"

    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
    }

    // Instantiate new attachment request builder.
    requestBuilder := attachmentapi.NewV1CompatibleAttachmentFileRequestBuilder2(pathParameters, client)

    // The content type of the file you want to upload
    dataContentType := "text/plain"
    // the byte content of the file
    data := []byte("this is example data")

    media := attachmentapi.NewMedia(dataContentType, data)

    // Define the required query parameters
    requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
        QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
            TableSysID: "INC00000001",
            TableName:  "incident",
            FileName:   "example.txt",
        }
    }

    // Call the post method with your content type, data, and request configurations.
    // Response is the uploaded file.
    response, err := requestBuilder.Post(context.Background(), media, requestConfiguration)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[POST\] /now/attachment/upload

Upload file of any supported content type.

### V1 client compatible

```golang
package main

import (
    "context"

    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
    }

    // Instantiate new attachment request builder.
    requestBuilder := attachmentapi.NewV1CompatibleAttachmentUploadRequestBuilder(pathParameters, client)

    body := // TODO: how to make multipart body?

    // Call the post method with your content type, data, and request configurations.
    // Response is the uploaded file.
    response, err := requestBuilder.Post(context.Background(), body, nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[GET\] /now/attachment/\<sys_id\>

<!-- TODO: Write [GET] /now/attachment/\<sys_id\> -->
