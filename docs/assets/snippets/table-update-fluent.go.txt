package main

import (
	"context"
    "fmt"
    "log"
    "time"

    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
    tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
	{%
		include-markdown 'assets/snippets/table-item-request-fluent.go'
	%}

	// Step 3: Configure request
	config := &tableapi.TableItemRequestBuilder2PutRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2PutQueryParameters{
			// Optional configurations
		},
	}

	ctx := context.Background()

	// Step 4: Prepare your data
	data := tableapi.NewTableRecord()
	data.SetValue("short_description", "updated incident description")

	// Step 5: Execute request
	response, err := client.Now2().TableV2("xSDK_SN_TABLEx").ById("xSDK_SN_TABLE_SYS_IDx").Put(ctx, data, config)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}