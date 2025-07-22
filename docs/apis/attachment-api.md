# Attachment API

## Overview

The Attachment API provides endpoints that allow you to upload and query file attachments.
You can upload or retrieve a single file with each request.

## \[GET\] /now/attachment

Returns the metadata for multiple attachments.

```golang
package main

import (
    "context"

    serviceNow "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
    }

    client := serviceNow.NewServiceNowClient2()

    // Call the get method, with or without AttachmentRequestBuilderGetQueryParameters.
    // Response is a AttachmentCollectionResponse.
    response, err := client.Now().Attachment2().Get(context.Background(), nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[POST\] /now/attachment/file

Upload file of any supported content type.

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

    // The content type of the file you want to upload
    dataContentType := "text/plain"
    // the byte content of the file
    data := []byte("this is example data")

    media := attachmentapi.NewMedia(dataContentType, data)

    // Define the required query parameters
    requestConfiguration := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
        QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
            TableSysID: "INC00000001",
            TableName:  "incident",
            FileName:   "example.txt",
        }
    }

    // Call the post method with your content type, data, and request configurations.
    // Response is the uploaded file.
    response, err := client.Now().Attachment2().Post(context.Background(), media, requestConfiguration)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[POST\] /now/attachment/upload

Upload file of any supported content type.

```golang
package main

import (
    "context"

    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    body := // TODO: how to make multipart body?

    // Call the post method with your content type, data, and request configurations.
    // Response is the uploaded file.
    response, err := client.Now().Attachment2().Post(context.Background(), body, nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[GET\] /now/attachment/\<sys_id\>

```golang
package main

import (
    "context"

    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    // Call the get method with/without request configurations.
    // Response is the attachment item.
    response, err := client.Now().Attachment2().ByID("sys id here").Get(context.Background(), nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[DELETE\] /now/attachment/\<sys_id\>

### V1 client compatible

```golang
package main

import (
    "context"

    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    // Call the delete method with/without request configurations.
    err := client.Now().Attachment2().ByID("sys id here").Delete(context.Background(), nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```

## \[GET\] /now/attachment/\<sys_id\>/file

```golang
package main

import (
    "context"

    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    // Call the delete method with/without request configurations.
    // response is the file with its metadata
    response, err := client.Now().Attachment2().ByID("sys id here").Get(context.Background(), nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```