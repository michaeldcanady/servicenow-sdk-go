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

	// Step 5: Configure query parameters
	params := &tableapi.TableItemRequestBuilderGetQueryParameters{
		// Optional configurations
	}

	// Step 6: Prepare your data
	data := map[string]string{
		"short_description": "example incident",
		"description": "incident created by servicenow-sdk-go",
	}

	// Step 7: Execute request
	record, err := requestBuilder.Put3(context.Background(), data, params)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}