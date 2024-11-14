package tableapi

import (
	"context"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/RecoLabs/servicenow-sdk-go/internal"
	intCore "github.com/RecoLabs/servicenow-sdk-go/internal/core"
)

const (
	tableItemURLTemplate = "{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}"
)

// TableItemRequestBuilder2 provides an interface for the service methods, adhering to the Interface Segregation Principle.
type TableItemRequestBuilder2[T TableRecord] interface {
	Get(context.Context, *TableItemRequestBuilderGetQueryParameters) (TableItemResponse3[T], error)
	Delete(context.Context, *TableItemRequestBuilderDeleteQueryParameters) error
	Put(context.Context, interface{}, *TableItemRequestBuilderPutQueryParameters) (TableItemResponse3[T], error)
}

type tableItemRequestBuilder2[T TableRecord] struct {
	intCore.Sendable
}

// newTableItemRequestBuilder2 creates a new instance of TableItemRequestBuilder2.
func newTableItemRequestBuilder2[T TableRecord](client intCore.ClientSendable, pathParameters map[string]string) (TableItemRequestBuilder2[T], error) {
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

	return &tableItemRequestBuilder2[T]{
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
func (rB *tableItemRequestBuilder2[T]) Get(ctx context.Context, params *TableItemRequestBuilderGetQueryParameters) (TableItemResponse3[T], error) {
	config := &intCore.RequestConfigurationImpl{
		Header:          nil,
		QueryParameters: interface{}(params),
		Data:            nil,
		ErrorMapping:    nil,
		Response:        &tableItemResponse3[T]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodGet, config)
	if err != nil {
		return nil, err
	}

	return resp.(TableItemResponse3[T]), nil
}

// Delete sends an HTTP DELETE request using the specified query parameters and returns an error if the request or response encounters any issues.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderDeleteQueryParameters to include in the DELETE request.
//
// Returns:
//   - error: An error if there was an issue with the request or response, or nil if the request was successful.
func (rB *tableItemRequestBuilder2[T]) Delete(ctx context.Context, params *TableItemRequestBuilderDeleteQueryParameters) error {
	config := &intCore.RequestConfigurationImpl{
		Header:          nil,
		QueryParameters: interface{}(params),
		Data:            nil,
		ErrorMapping:    nil,
		Response:        &tableItemResponse3[T]{},
	}

	_, err := rB.Send(ctx, intCore.MethodDelete, config)

	return err
}

// Put updates a table item using an HTTP PUT request.
// It takes a map of table entry data and optional query parameters to send in the request.
// The method returns a TableItemResponse representing the updated item or an error if the request fails.
//
// Parameters:
//   - entry: A map[string]string or TableEntry containing the data to update the table item.
//   - params: An optional pointer to TableItemRequestBuilderPutQueryParameters, which can be used to specify query parameters for the request.
//
// Returns:
//   - *TableItemResponse: A TableItemResponse containing the updated item data.
//   - error: An error, if the request fails at any point, such as request information creation or JSON deserialization.
func (rB *tableItemRequestBuilder2[T]) Put(ctx context.Context, entry interface{}, params *TableItemRequestBuilderPutQueryParameters) (TableItemResponse3[T], error) {
	entry, err := convertFromTableEntry(entry)
	if err != nil {
		return nil, err
	}

	config := &intCore.RequestConfigurationImpl{
		Header:          nil,
		QueryParameters: interface{}(params),
		Data:            entry,
		ErrorMapping:    nil,
		Response:        &tableItemResponse3[T]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodPut, config)
	if err != nil {
		return nil, err
	}

	return resp.(TableItemResponse3[T]), nil
}
