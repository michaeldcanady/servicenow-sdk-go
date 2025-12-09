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
	params := &tableapi.TableItemRequestBuilderDeleteQueryParameters{
		// Optional configurations
	}

	// Step 4: Execute request
	err := client.Now2().Table2("xSDK_SN_TABLEx").ByID("xSDK_SN_TABLE_SYS_IDx").Delete2(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}
}