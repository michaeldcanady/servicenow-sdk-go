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
		include-markdown 'assets/snippets/table-request-standard.go'
	%}

	// Step 5: Configure query parameters
	params := &tableapi.TableRequestBuilderPostQueryParameters{
		// Optional configurations
	}

	// Step 6: Build request body
	data := map[string]string{
		"short_description": "example incident",
		"description":       "incident created by servicenow-sdk-go",
	}

	// Step 7: Execute request
	response, err := requestBuilder.Post4(context.Background(), data, params)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}