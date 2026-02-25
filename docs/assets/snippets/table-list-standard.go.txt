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
    config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
        QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
            // Optional configurations
        },
    }

    // Step 6: Execute request
    response, err := requestBuilder.Get(context.Background(), config)
    if err != nil {
        log.Fatal(err)
    }

    // Process response
}
