# Attachment API

The `Attachment API` provides endpoints that allow you to upload and query file attachments.
You can upload or retrieve a single file with each request.

## \[GET\] Retrieve an attachment

Returns the metadata for multiple attachments.

[Try on Playground](https://go.dev/play/p/m4bOojq8ODB)

```golang
package main

import (
    attachmentapi "github.com/RecoLabs/servicenow-sdk-go/attachment-api"
)

func main() {
    
    //Implement credential and client.
    pathParameters := map[string]string{
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
