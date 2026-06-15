# Upload attachment from disk

Uploading a file from disk is the most common way developers attach documents, logs, screenshots, or exports to a ServiceNow record. This variation walks through the workflow end‑to‑end and highlights the key decisions you’ll make along the way.

## When to use this pattern

- You have a file already stored locally (for example, a .txt, .pdf, .png, .zip)
- You’re building a CLI tool, automation script, or integration service
- You want a simple, reliable way to attach files without streaming or buffering logic

## Required values

| Value         | Description                               |
| ------------- | ----------------------------------------- |
| Table Name    | the record’s table (for example, `incident`) |
| Record sys_id | the specific record to attach the file to |
| File name     | the name that will appear in Service-Now  |
| File content  | read directly from disk                   |

## Example

```golang
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func main() {
    // Step 1: Authenticate with your ServiceNow instance
    cred := credentials.NewBasicAuthenticationProvider(username, password)

    clientOpts := []credentials.ServiceNowServiceClientOption{
        servicenow.WithAuthenticationProvider(cred),
        servicenow.WithInstance("{instance}"),
    }

    // Step 2: Initialize the ServiceNow client
    client, err := servicenow.NewServiceNowServiceClient(clientOpts...)
    if err != nil {
        log.Fatalf("failed to initialize client: %v", err)
    }

    // Step 3: Read the file from disk
    filePath := "./example.txt"
    fileBytes, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatalf("failed to read file: %v", err)
    }

    // Step 4: Wrap the file content in a Media object
    mediaType := "text/plain" // adjust based on your file type
    media := attachmentapi.NewMedia(mediaType, fileBytes)

    // Step 5: Provide the required attachment parameters
    tableName := "{TableName}"
    tableSysID := "{SysID}"
    fileName := "example.txt"

    config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
        QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
            TableName:  &tableName,
            TableSysID: &tableSysID,
            FileName:   &fileName,
        },
    }

    // Step 6: Upload the file
    resp, err := client.Now().Attachment().File().Post(
        context.Background(),
        media,
        config,
    )
    if err != nil {
        log.Fatal(err)
    }

    sysID, err := attachment.GetSysID()
    if err != nil {
        log.Printf("unable to retrieve attachment sys_id: %v", err)
        continue
    }

    fmt.Printf("Uploaded attachment with sys_id: %s\n", *sysID)
}
```

## Tips

- Use absolute paths when running from `cron` jobs or containerized environments.
- Set the correct Multipurpose Internet Mail Extensions (MIME) type (`text/plain`, `application/pdf`, `image/png`, etc.) so ServiceNow handles previews correctly.
- Check for duplicates by listing attachments first if your workflow might upload the same file.
- Validate file size if your instance enforces attachment limits.
