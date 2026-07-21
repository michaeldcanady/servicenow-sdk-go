package snippets

import (
	"context"
	"fmt"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func _() {
	var client *servicenowsdkgo.ServiceNowClient
	ctx := context.Background()

	// [START pagination_table_basic]
	// 1. Execute a list request
	listResponse, err := client.Now2().TableV2("incident").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Create the iterator
	iterator, err := tableapi.NewDefaultTablePageIterator(listResponse, client.RequestAdapter)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Iterate over all records across all pages
	err = iterator.Iterate(ctx, false, func(record *tableapi.TableRecord) bool {
		// Process the record
		sysId, _ := record.GetSysID()
		fmt.Printf("Incident ID: %s\n", *sysId)
		return true // Continue to the next record
	})

	if err != nil {
		log.Fatal(err)
	}
	// [END pagination_table_basic]

	// [START pagination_table_manual]
	// Fetch the next page of results manually
	nextPage, err := iterator.Next(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Process the items on the next page
	results := nextPage.Result
	for _, item := range results {
		fmt.Println(item)
	}
	// [END pagination_table_manual]

	// [START pagination_attachment]
	// 1. Execute an attachment list request
	attachmentResponse, err := client.Now2().Attachment2().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Create the iterator
	attachmentIterator, err := attachmentapi.NewAttachmentPageIterator(attachmentResponse, client.RequestAdapter)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Iterate over attachments
	if err := attachmentIterator.Iterate(ctx, false, func(attachment attachmentapi.Attachment2) bool {
		fileName, _ := attachment.GetFileName()
		fmt.Printf("Attachment: %s\n", *fileName)
		return true
	}); err != nil {
		log.Fatal(err)
	}
	// [END pagination_attachment]

	// [START pagination_item_by_item]
	// Iterate item by item using NextItem
	for iterator.HasNext() {
		item, err := iterator.NextItem(ctx)
		if err != nil {
			break
		}
		fmt.Println(item)
	}
	// [END pagination_item_by_item]

	// [START pagination_state_management]
	// Reset the iterator to the beginning
	iterator.Reset()

	// Restart iteration of the current page
	iterator.ResetPage()
	// [END pagination_state_management]
}
