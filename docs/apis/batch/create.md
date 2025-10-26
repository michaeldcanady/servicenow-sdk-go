# Create batch request

## Overview

Submits a `BatchRequest` containing all desired requests.

## Path parameters

N/A

## Optional query parameters

N/A

# Required query parameters

N/A

## Examples

=== "Fluent"

    ``` golang
    package main

    import (
        context

        batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
    )

    // batchRequests A helper function to combine provided request information into a single `BatchRequest`.
    func batchRequests(excludeResponseHeaders bool, requests ...*abstractions.RequestInformation) (*batchapi.BatchRequestModel, error) {
        body := batchapi.NewBatchRequestModel()
        for _, request := range requests {
            restRequest, err := batchapi.CreateRestRequestFromRequestInformation(request, excludeResponseHeaders)
            if err != nil {
                return nil, err
            }
            if err := body.AddRequest(restRequest); err != nil {
                return nil, err
            }
        }
        return body, nil
    }

    func main() {
        // Initialize credentials and client

        // Build requests, using ToXXXRequestInformation method

        body := batchRequests(true, ...requests)
        
        // build your request
        response, err := client.Now().Batch().Post(context.Backgroud(), body)
        if err != nil {
            log.Fatal(err)
        }

        // Process response
    }
    ```