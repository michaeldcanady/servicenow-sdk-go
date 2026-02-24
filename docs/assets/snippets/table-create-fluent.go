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

	// Step 3: Configure request
	config := &tableapi.TableRequestBuilder2PostRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2PostQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Build request body
	data := tableapi.NewTableRecord()
	data.SetValue("short_description", "example incident")
	data.SetValue("description", "incident created by servicenow-sdk-go")

	// Step 5: Execute request
	response, err := client.Now2().TableV2("xSDK_SN_TABLEx").Post(context.Background(), data, config)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}