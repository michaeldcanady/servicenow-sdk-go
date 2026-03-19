# Delete attachments older than X days
As attachments accumulate over time, they can consume storage and clutter records. Many teams implement retention policies to automatically remove older files—especially logs, exports, or temporary documents. This task shows you how to identify attachments older than a specified number of days and delete them.

This workflow pairs with listing and downloading attachments.

## When to use this pattern

Use this pattern when you need to:

- Enforce retention policies (for example, delete logs older than 30 days)  
- Clean up temporary or auto‑generated files  
- Reduce storage usage on large tables  
- Prepare a record for archival or migration

## Required values

## Example

```golang
package main

import (
    "context"
    "fmt"
    "log"
    "time"

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

    // Step 3: Determine the cutoff date
    retentionDays := 30
    cutoff := time.Now().AddDate(0, 0, -retentionDays)

    // Step 4: List attachments for a specific record
    config := &attachmentapi.AttachmentRequestBuilder2GetRequestConfiguration{
        QueryParameters: &attachmentapi.AttachmentRequestBuilder2GetQueryParameters{
            SysparmQuery: query.And(
                query.String("table_sys_id").Is("{table entry's sys_id}"),
                query.String("table_name").Is("{name of table}"),
                query.Date("sys_created_on").Before(query.NewDateTimeValue(cutoff))
            ).String(),
        },
    }

    resp, err := client.Now().Attachment().Get(context.Background(), config)
    if err != nil {
        log.Fatalf("unable to send/receive request: %v", err)
    }

    attachments, err := resp.GetResult()
    if err != nil {
        log.Fatalf("unable to retrieve result(s): %v", err)
    }

    if len(attachments) == 0 {
        log.Println("No attachments found.")
        return
    }

    // Step 5: Delete attachments older than the cutoff
    for _, attachment := range attachments {
        createdOn, err := attachment.GetSysCreatedOn()
        if err != nil {
            log.Printf("unable to retrieve creation date: %v", err)
            continue
        }

        sysID, err := attachment.GetSysID()
        if err != nil {
            log.Printf("unable to retrieve sys_id: %v", err)
            continue
        }

        deleteConfig := &attachmentapi.AttachmentItemRequestBuilderDeleteRequestConfiguration{}

        if err := client.
            Now().
            Attachment().
            ByID(*sysID).
            Delete(context.Background(), deleteConfig); err != nil {

            log.Printf("failed to delete attachment %s: %v", *sysID, err)
            continue
        }

        fmt.Printf("Deleted attachment %s (created %s)\n", *sysID, createdOn.String())
    }
}
```

## Tips

- Always **log deletions** for auditing.  
- Consider performing a **dry run** first—list what you would delete without deleting.  
- Download attachments before deleting them if you need to archive them.  
- Use `sysparm_limit` and pagination for tables with many attachments.  
- Combine this with the **Download All Attachments** task to archive before cleanup.
