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
    config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
    	QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
    		// Optional configurations
    	},
	}

	// Step 4: Execute request
	response, err := client.Now2().TableV2("xSDK_SN_TABLEx").Get(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
}