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
	// Step 1: Create credentials
    {%
		include-markdown 'assets/snippets/credentials.go'
	%}

    // Step 2: Initialize client
    {%
		include-markdown 'assets/snippets/client.go'
	%}

	// Step 3: Configure query parameters
	params := &tableapi.TableRequestBuilderPostQueryParameters{
		// Optional configurations
	}

	// Step 4: Build request body
	data := map[string]string{
		"short_description": "example incident",
		"description":       "incident created by servicenow-sdk-go",
	}

	// Step 5: Execute request
	response, err := client.Now2().Table2("xSDK_SN_TABLEx").Post4(context.Background(), data, params)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}