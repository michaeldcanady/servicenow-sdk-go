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
		include-markdown 'assets/snippets/table-item-request-standard.go'
	%}

	// Step 5: Configure request
	config := &tableapi.TableItemRequestBuilder2PutRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2PutQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Prepare your data
	data := tableapi.NewTableRecord()
	data.SetValue("short_description", "updated incident description")

	// Step 7: Execute request
	record, err := requestBuilder.Put(context.Background(), data, config)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}