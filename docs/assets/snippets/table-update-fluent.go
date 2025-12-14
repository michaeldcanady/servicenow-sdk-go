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

	// Step 3: Configure query parameters
	params := &tableapi.TableItemRequestBuilderGetQueryParameters{
		// Optional configurations
	}

	ctx := context.Background()

	// Step 4: Prepare your data
	data := map[string]string{
		"short_description": "example incident",
		"description": "incident created by servicenow-sdk-go",
	}

	// Step 5: Execute request
	response, err := client.Now2().Table2("xSDK_SN_TABLEx").Put3(context.Background(), data, params)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}