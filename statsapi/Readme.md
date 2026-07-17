# Stats API

The `Stats` API returns aggregate statistics (count, sum, average, minimum, maximum) computed over
the records of a Service-Now table, without returning the records themselves.

This module currently supports the **ungrouped** shape of the API only: a single aggregate result
per request. `sysparm_group_by` / `sysparm_having` are not supported — when set, the platform
returns an array of results under `result` instead of a single object, which `StatsRequestBuilder`
does not yet model.

## \[GET\] /now/stats/{tableName}

Retrieves aggregate statistics for the specified table.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/statsapi"
)

func main() {
	// Implement credential and client.
	var client *servicenowsdkgo.ServiceNowServiceClient

	requestConfiguration := &statsapi.StatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &statsapi.StatsRequestBuilderGetQueryParameters{
			Count:        true,
			SumFields:    []string{"reassignment_count"},
			Query:        "active=true",
			DisplayValue: statsapi.DisplayValueAll,
		},
	}

	response, err := client.Now().Stats("incident").Get(context.Background(), requestConfiguration)
	if err != nil {
		panic(err)
	}

	result, err := response.GetResult()
	if err != nil {
		panic(err)
	}

	stats, err := result.GetStats()
	if err != nil {
		panic(err)
	}

	count, _ := stats.GetCount()
	sum, _ := stats.GetSum()
	_ = count
	_ = sum
}
```
