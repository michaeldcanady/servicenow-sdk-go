package servicenowsdkgo

import (
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

type NowRequestBuilder2 struct {
	intCore.RequestBuilder2
}

const (
	nowURLTemplate = "{+baseurl}/api/now"
)

// NewNowRequestBuilder creates a new instance of the NowRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created NowRequestBuiabstraction
func NewNowRequestBuilder2(client intCore.Client2, pathParameters map[string]string) *NowRequestBuilder2 {
	return &NowRequestBuilder2{
		intCore.NewRequestBuilder2(client, nowURLTemplate, pathParameters),
	}
}

// Table returns a TableRequestBuilder associated with the NowRequestBuilder.
// It accepts a table name as a parameter and constructs the URL for table-related requests.
// The returned TableRequestBuilder can be used to build and execute table-related requests.
func (rB *NowRequestBuilder2) Table(tableName string) *tableapi.TableRequestBuilder2 {
	pathParameters := rB.GetPathParameters()
	pathParameters["table"] = tableName
	client := rB.GetClient()
	requestBuilder, _ := tableapi.NewTableRequestBuilder2(client, pathParameters)
	return requestBuilder
}
