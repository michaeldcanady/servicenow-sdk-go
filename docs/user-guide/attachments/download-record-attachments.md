# Download All Attachments for a Record

Once you’ve identified which attachments belong to a record, the next step is often to download them - for backups, migrations, processing, or exporting to another system. This task shows you how to retrieve the file content for each attachment and save it locally.

This workflow pairs naturally with listing attachments and with cleanup tasks like deleting older files.

## When to Use This Pattern

- Use this pattern when you need to:
- Export all attachments for a record
- Back up files before deleting or archiving them
- Process attachments locally (e.g., parse logs, read CSVs, extract text)
- Migrate attachments to another system

## Required Values

| Value             | Description              |
| ----------------- | ------------------------ |
| Instance          | the Service-Now instance |
| Table Entry SysID | the table entry's SysID  |

## Example

```golang
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "path/filepath"

    servicenow "github.com/michaeldcanady/servicenow-sdk-go"
    attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
    "github.com/michaeldcanady/servicenow-sdk-go/query"
)

func main() {
    // Step 1: Authenticate with your ServiceNow instance
    cred := credentials.NewBasicAuthenticationProvider("{username}", "{password}")

    clientOpts := []servicenow.ServiceNowServiceClientOption{
        servicenow.WithAuthenticationProvider(cred),
        servicenow.WithInstance("{instance}"),
    }

    // Step 2: Initialize the ServiceNow client
    client, err := servicenow.NewServiceNowServiceClient(clientOpts...)
    if err != nil {
        log.Fatalf("failed to initialize client: %v", err)
    }

    // Step 3: List attachments for a specific record
    config := &attachmentapi.AttachmentRequestBuilder2GetRequestConfiguration{
        QueryParameters: &attachmentapi.AttachmentRequestBuilder2GetQueryParameters{
            SysparmQuery: query.And(
                query.String("table_sys_id").Is("{table entry's sys_id}"),
                query.String("table_name").Is("{name of table}"),
            ).String(),
        },
    }

    resp, err := client.Now().Attachment().Get(context.Background(), config)
    if err != nil {
        log.Fatalf("unable to send/receive request: %v", err)
    }

    // Step 4: Retrieve the attachment results
    attachments, err := resp.GetResult()
    if err != nil {
        log.Fatalf("unable to retrieve result(s): %v", err)
    }

    if len(attachments) == 0 {
        log.Println("No attachments found.")
        return
    }

    // Step 5: Download each attachment
    for _, attachment := range attachments {
        sysID, err := attachment.GetSysID()
        if err != nil {
            log.Printf("unable to retrieve attachment sys_id: %v", err)
            continue
        }

        fileName, err := attachment.GetFileName()
        if err != nil {
            log.Printf("unable to retrieve attachment name: %v", err)
            continue
        }

        // Download the file content
        fileResp, err := client.
            Now().
            Attachment().
            ByID(*sysID).
            File().
            Get(context.Background(), nil)

        if err != nil {
            log.Printf("unable to download attachment %s: %v", *sysID, err)
            continue
        }

        content, err := fileResp.GetContent()
        if err != nil {
            log.Printf("unable to read attachment content: %v", err)
            continue
        }

        // Write the file to disk
        outputPath := filepath.Join("./downloads", *fileName)
        if err := os.WriteFile(outputPath, content, 0644); err != nil {
            log.Printf("failed to write file %s: %v", outputPath, err)
            continue
        }

        fmt.Printf("Downloaded attachment: %s → %s\n", *fileName, outputPath)
    }
}
```
