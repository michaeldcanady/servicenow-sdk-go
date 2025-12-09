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
    params := &tableapi.TableRequestBuilderGetQueryParameters{
        // Optional configurations
    }

    // Step 6: Execute request
    response, err := requestBuilder.Get2(context.Background(), params)
    if err != nil {
        log.Fatal(err)
    }

    // Process response
}
