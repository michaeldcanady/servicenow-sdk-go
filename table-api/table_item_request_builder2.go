package tableapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	// tableItemURLTemplate is the URL template for constructing table item URLs.
	tableItemURLTemplate = "{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}"
)

// TableItemRequestBuilder2 represents a request builder for table items.
type TableItemRequestBuilder2 struct {
	intCore.RequestBuilder2
}

// NewTableItemRequestBuilder2 creates a new instance of TableItemRequestBuilder2.
func NewTableItemRequestBuilder2(client intCore.Client2, pathParameters map[string]string) (*TableItemRequestBuilder2, error) {
	if internal.IsNil(client) {
		return nil, ErrNilClient
	}

	_, basePathOk := pathParameters[internal.BasePathParameter]
	if !basePathOk {
		return nil, core.ErrMissingBasePathParam
	}

	_, tableOk := pathParameters["table"]
	if !tableOk {
		return nil, ErrNilParameterTable
	}

	_, sysIDOk := pathParameters["sysId"]
	if !sysIDOk {
		return nil, ErrNilParameterSysID
	}

	return &TableItemRequestBuilder2{
		intCore.NewRequestBuilder2(client, tableItemURLTemplate, pathParameters),
	}, nil
}

// Get sends an HTTP GET request using the specified query parameters and returns a TableItemResponse.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *TableItemResponse: The response data as a TableItemResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *TableItemRequestBuilder2) Get(ctx context.Context, opts ...intCore.RequestConfigurationOption) (*TableItemResponse, error) {
	opts = append(opts, intCore.WithResponse(&TableItemResponse2[TableEntry]{}))

	resp, err := rB.Send(ctx, intCore.GET, opts...)
	if err != nil {
		return nil, err
	}

	return resp.(*TableItemResponse), nil
}

// Delete sends an HTTP DELETE request using the specified query parameters and returns an error if the request or response encounters any issues.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderDeleteQueryParameters to include in the DELETE request.
//
// Returns:
//   - error: An error if there was an issue with the request or response, or nil if the request was successful.
func (rB *TableItemRequestBuilder2) Delete(ctx context.Context, opts ...intCore.RequestConfigurationOption) error {
	_, err := rB.Send(ctx, intCore.DELETE, opts...)
	return err
}

// Put updates a table item using an HTTP PUT request.
// It takes a map of table entry data and optional query parameters to send in the request.
// The method returns a TableItemResponse representing the updated item or an error if the request fails.
//
// Parameters:
//   - tableEntry: A map[string]string or TableEntry containing the data to update the table item.
//   - params: An optional pointer to TableItemRequestBuilderPutQueryParameters, which can be used to specify query parameters for the request.
//
// Returns:
//   - *TableItemResponse: A TableItemResponse containing the updated item data.
//   - error: An error, if the request fails at any point, such as request information creation or JSON deserialization.
func (rB *TableItemRequestBuilder2) Put(ctx context.Context, opts ...intCore.RequestConfigurationOption) (*TableItemResponse, error) {
	opts = append(opts, intCore.WithResponse(&TableItemResponse2[TableEntry]{}))

	resp, err := rB.Send(ctx, intCore.PUT, opts...)
	if err != nil {
		return nil, err
	}

	return resp.(*TableItemResponse), nil
}
