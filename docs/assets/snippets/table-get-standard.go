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
	config := &tableapi.TableItemRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2GetQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Execute request
	record, err := requestBuilder.Get(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}