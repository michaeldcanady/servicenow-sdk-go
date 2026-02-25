# Working with Attachments

The Attachment API lets you manage attachments for ServiceNow records.

## Listing Attachments

To list attachments for a specific table or record:

```go
import (
    "context"
    "fmt"
    "log"
    "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    // ... client initialization ...

    ctx := context.Background()
    
    // List all attachments
    response, err := client.Now2().Attachment2().Get(ctx, nil)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    for _, attachment := range response.GetValue() {
        fmt.Printf("Attachment: %s
", attachment.GetFileName())
    }
}
```

## Creating an Attachment

To create an attachment by uploading a file:

```go
import (
    "context"
    "os"
    "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
)

func main() {
    // ... client initialization ...

    ctx := context.Background()
    
    file, err := os.Open("path/to/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Upload attachment for an incident
    // sys_id is the record you want to attach the file to
    // tableName is the table of the record
    params := &attachmentapi.AttachmentRequestBuilderFileQueryParameters{
        TableName: "incident",
        TableSysId: "your-incident-sys-id",
        FileName: "file.txt",
    }

    config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
        QueryParameters: params,
    }

    response, err := client.Now2().Attachment2().File().Post(ctx, file, config)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Printf("Created attachment with sys_id: %s
", response.GetId())
}
```

## Downloading an Attachment

To download the content of an attachment:

```go
sysId := "your-attachment-sys-id"

content, err := client.Now2().Attachment2().ById(sysId).File().Get(ctx, nil)
if err != nil {
    log.Fatal(err)
}

// content is a []byte containing the file data
```

## Deleting an Attachment

```go
sysId := "your-attachment-sys-id"
err := client.Now2().Attachment2().ById(sysId).Delete(ctx, nil)
```
