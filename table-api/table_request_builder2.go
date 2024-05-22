package tableapi

import (
	"errors"
	"strconv"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intTable "github.com/michaeldcanady/servicenow-sdk-go/table-api/internal"
)

const (
	tableURLTemplate = "{+baseurl}/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder is responsible for building requests for table operations.
type TableRequestBuilder2 struct {
	intTable.RequestBuilder
}

// NewTableRequestBuilder initializes a new TableRequestBuilder with the given client and path parameters.
func NewTableRequestBuilder2(client core.Client, pathParameters map[string]string) (*TableRequestBuilder2, error) {
	if internal.IsNil(client) {
		return nil, ErrNilClient
	}

	_, basePathOk := pathParameters[internal.BasePathParameter]
	if !basePathOk {
		return nil, core.ErrMissingBasePathParam
	}
	_, tableOk := pathParameters["table"]
	if !tableOk {
		return nil, errors.New("missing \"table\" parameter")
	}

	return &TableRequestBuilder2{
		core.NewRequestBuilder( //nolint:staticcheck
			client,
			tableURLTemplate,
			pathParameters,
		),
	}, nil
}

// TableService defines the operations available for the table as a whole.
type TableService interface {
	Get(params *TableRequestBuilderGetQueryParameters) (*TableCollectionResponse, error)
	Post(data interface{}, params *TableRequestBuilderPostQueryParameters) (*TableItemResponse, error)
	Count() (int, error)
}

// Ensure that TableRequestBuilder implements TableService.
var _ TableService = (*TableRequestBuilder2)(nil)

// ByID creates a TableItemRequestBuilder for a specific record in the table identified by sysID.
func (rB *TableRequestBuilder2) ByID(sysID string) (*TableItemRequestBuilder2, error) {
	pathParameters := rB.RequestBuilder.(*core.RequestBuilder).PathParameters //nolint:staticcheck
	client := rB.RequestBuilder.(*core.RequestBuilder).Client                 //nolint:staticcheck
	pathParameters["sysId"] = sysID
	return NewTableItemRequestBuilder2(client, pathParameters)
}

// Get retrieves a collection of table items based on the provided query parameters.
func (rB *TableRequestBuilder2) Get(params *TableRequestBuilderGetQueryParameters) (*TableCollectionResponse, error) {
	config := &tableGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		mapping:  nil,
		response: &TableCollectionResponse2[TableEntry]{},
	}

	err := rB.SendGet2(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Post creates a new table item with the provided data and query parameters.
func (rB *TableRequestBuilder2) Post(data interface{}, params *TableRequestBuilderPostQueryParameters) (*TableItemResponse, error) {
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

	err = rB.SendPost3(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Count retrieves the total count of items in the table.
func (rB *TableRequestBuilder2) Count() (int, error) {
	requestInfo, err := rB.RequestBuilder.ToHeadRequestInformation()
	if err != nil {
		return -1, err
	}

	errorMapping := core.ErrorMapping{"4XX": "hi"}

	response, err := rB.RequestBuilder.(*core.RequestBuilder).Client.Send(requestInfo, errorMapping) //nolint:staticcheck
	if err != nil {
		return -1, err
	}

	count, err := strconv.Atoi(response.Header.Get("X-Total-Count"))
	if err != nil {
		count = 0
	}

	return count, nil
}
