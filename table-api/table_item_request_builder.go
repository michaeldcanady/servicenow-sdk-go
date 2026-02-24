package tableapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// Deprecated: deprecated since v{unreleased}. Please use [NewTableItemRequestBuilder3] or [NewTableItemRequestBuilder2Internal]
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

// Deprecated: deprecated since v{unreleased}. Please use [TableItemRequestBuilder2.Get]
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

// Deprecated: deprecated since v{unreleased}. Please use [TableItemRequestBuilder2.Delete]
func (rB *TableItemRequestBuilder) Delete2(ctx context.Context, params *TableItemRequestBuilderDeleteQueryParameters) error {
	config := &tableItemDeleteRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: nil,
	}

	return rB.SendDelete3(ctx, config.toConfiguration())
}

// Deprecated: deprecated since v{unreleased}. Please use [TableItemRequestBuilder2.Put]
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
