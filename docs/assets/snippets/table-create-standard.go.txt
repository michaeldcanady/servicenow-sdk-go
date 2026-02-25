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

	// Step 5: Configure request
	config := &tableapi.TableRequestBuilder2PostRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2PostQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Build request body
	data := tableapi.NewTableRecord()
	data.SetValue("short_description", "example incident")
	data.SetValue("description", "incident created by servicenow-sdk-go")

	// Step 7: Execute request
	response, err := requestBuilder.Post(context.Background(), data, config)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}