# Attachment API

The Attachment API provides endpoints that allow you to upload and query file attachments.
You can upload or retrieve a single file with each request.

## \[GET\] /now/attachment

Returns the metadata for multiple attachments.

```golang
package main

import (
    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := {
        "baseurl":"https://www.{instance}.service-now.com/api/now",
    }

    // Instantiate new TableItemRequestBuilder.
    requestBuilder := attachmentapi.NewAttachmentRequestBuilder(client, pathParameters)

    // Call the get method, with or without AttachmentRequestBuilderGetQueryParameters.
    // Response is a AttachmentCollectionResponse.
    response, err := requestBuilder.Get(nil)

    // Test err, should be nil
    if err != nil {
        panic(err)
    }
}
```
