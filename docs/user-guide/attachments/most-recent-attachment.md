# Find the most recent attachment

When working with attachments, you often don’t need every attachment—you need the latest log, screenshot, export, or document. This task shows you how to list attachments for a specific record, sort them by creation time, and select the most recent one.

This is a foundational workflow that pairs with downloading, processing, or deleting attachments.

## When to use this pattern

Use this pattern when you need to:

- Retrieve the newest log file for debugging
- Identify the latest user‑submitted document
- Prepare for a follow‑up action (download, delete, archive)
- Build automation workflows that react to new attachments

## Required values

| Value    | Description              |
| -------- | ------------------------ |
| Instance | the Service-Now instance |

## Example

```golang
package main

import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
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

	// Step 3: List attachments (filter using query params as needed)
	config := &attachmentapi.AttachmentRequestBuilder2GetRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentRequestBuilder2GetQueryParameters{
			SysparmLimit: 1,
			SysparmQuery: "ORDERBYDESCsys_created_on",
		},
	}

	resp, err := client.Now().Attachment().Get(
		context.Background(),
		config,
	)
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

	// Step 5: Select the most recent attachment (already sorted by ORDERBYDESC)
	latest := attachments[0]

	fileName, err := latest.GetFileName()
	if err != nil {
		log.Fatalf("unable to retrieve attachment's name: %v", err)
	}

	sysID, err := latest.GetSysID()
	if err != nil {
		log.Fatalf("unable to retrieve attachment's sys_id: %v", err)
	}

	fmt.Printf("Most recent attachment: %s (sys_id: %s)\n", *fileName, *sysID)
}
```

## Tips

- Use `ORDERBYDESCsys_created_on` to make sure the newest attachment appears first.
- Limit your query (`sysparm_limit`) to reduce response size and improve performance.
- Filter by table and `sys_id` when you only care about attachments for a specific record.
- Always check for empty results—not all records have attachments.
