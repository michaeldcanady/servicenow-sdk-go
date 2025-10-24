package tableapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// NewTableItemRequestBuilder2 creates a new instance of the TableItemRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableItemRequestBuilder.
func NewTableItemRequestBuilder2(client core.Client2, pathParameters map[string]string) *TableItemRequestBuilder {
	requestBuilder := core.NewRequestBuilder2(
		client,
		tableItemURLTemplate,
		pathParameters,
	)
	return &TableItemRequestBuilder{
		*requestBuilder,
	}
}

// Get2 sends an HTTP GET request using the specified query parameters and returns a TableItemResponse.
func (rB *TableItemRequestBuilder) Get2(ctx context.Context, params *TableItemRequestBuilderGetQueryParameters) (*TableItemResponse, error) {
	config := &tableItemGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	err := rB.SendGet3(ctx, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Delete2 sends an HTTP DELETE request using the specified query parameters and returns an error if the request or response encounters any issues.
func (rB *TableItemRequestBuilder) Delete2(ctx context.Context, params *TableItemRequestBuilderDeleteQueryParameters) error {
	config := &tableItemDeleteRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: nil,
	}

	return rB.SendDelete3(ctx, config.toConfiguration())
}

// Put3 updates a table item using an HTTP PUT request.
// It takes a map of table entry data and optional query parameters to send in the request.
// The method returns a TableItemResponse representing the updated item or an error if the request fails.
func (rB *TableItemRequestBuilder) Put3(ctx context.Context, tableEntry interface{}, params *TableItemRequestBuilderPutQueryParameters) (*TableItemResponse, error) {
	tableEntry, err := convertFromTableEntry(tableEntry)
	if err != nil {
		return nil, err
	}

	config := &tableItemPutRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     tableEntry,
		response: &TableItemResponse2[TableEntry]{},
	}

	err = rB.SendPut3(ctx, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}
