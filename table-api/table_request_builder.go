package tableapi

import (
	"context"
	"strconv"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

func New2TableRequestBuilder(client core.Client2, pathParameters map[string]string) *TableRequestBuilder {
	requestBuilder := core.NewRequestBuilder2(
		client,
		tableURLTemplate,
		pathParameters,
	)
	return &TableRequestBuilder{
		*requestBuilder,
	}
}

func (rB *TableRequestBuilder) ByID2(sysId string) *TableItemRequestBuilder {
	pathParameters := rB.PathParameters
	pathParameters["sysId"] = sysId
	return NewTableItemRequestBuilder2(rB.Client2, pathParameters)
}

func (rB *TableRequestBuilder) Get2(ctx context.Context, params *TableRequestBuilderGetQueryParameters) (*TableCollectionResponse, error) {
	errMapping := core.NewErrorMapping()
	errMapping.Set("XXX", "placeholder")
	config := &tableGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		mapping:  errMapping,
		response: &TableCollectionResponse2[TableEntry]{},
	}

	err := rB.SendGet3(ctx, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

func (rB *TableRequestBuilder) Post4(ctx context.Context, data interface{}, params *TableRequestBuilderPostQueryParameters) (*TableItemResponse, error) {
	data, err := convertFromTableEntry(data)
	if err != nil {
		return nil, err
	}

	config := &tablePostRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     data.(map[string]string),
		mapping:  nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	err = rB.SendPost4(ctx, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Count sends an HTTP HEAD request and retrieves the value of "X-Total-Count" from the response header, which represents the count of items.
//
// Returns:
//   - int: The count of items.
//   - error: An error if there was an issue with the request or response.
func (rB *TableRequestBuilder) Count() (int, error) {
	requestInfo, err := rB.ToHeadRequestInformation()
	if err != nil {
		return -1, err
	}

	errorMapping := core.ErrorMapping{"4XX": "hi"}

	response, err := rB.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return -1, err
	}

	count, err := strconv.Atoi(response.Header.Get("X-Total-Count"))
	if err != nil {
		count = 0
	}

	return count, nil
}
