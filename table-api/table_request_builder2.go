package tableapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInt "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/michaeldcanady/servicenow-sdk-go/models"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	tableURLTemplate   = "{+baseurl}/api/now/v2/table/{table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
	firstLinkHeaderKey = "first"
	prevLinkHeaderKey  = "prev"
	nextLinkHeaderKey  = "next"
	lastLinkHeaderKey  = "last"
)

// TableRequestBuilder2 provides operations to manage Service-Now tables.
type TableRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
	factory serialization.ParsableFactory
}

// NewAPIV1CompatibleDefaultTableRequestBuilder2Internal converts api v1 compatible elements into api v2 compatible elements
func NewAPIV1CompatibleDefaultTableRequestBuilder2Internal(
	pathParameters map[string]string,
	client core.Client, //nolint: staticcheck
) *TableRequestBuilder2 {
	reqAdapter, _ := internal.NewServiceNowRequestAdapterBase(core.NewAPIV1ClientAdapter(client))

	return NewDefaultTableRequestBuilder2Internal(
		pathParameters,
		reqAdapter,
	)
}

// NewDefaultTableRequestBuilder2Internal instantiates a new TableRequestBuilder2 and sets the default values.
func NewDefaultTableRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder2 {
	return newTableRequestBuilder2Internal(pathParameters, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// newTableRequestBuilder2Internal instantiates a new TableRequestBuilder2 with custom parsable for table entries.
func newTableRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder2 {
	m := &TableRequestBuilder2{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, tableURLTemplate, pathParameters),
		factory:            factory,
	}
	return m
}

// NewDefaultTableRequestBuilder2 instantiates a new TableRequestBuilder2 and sets the default values.
func NewDefaultTableRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder2 {
	return newTableRequestBuilder2(rawURL, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// newTableRequestBuilder2 instantiates a new TableRequestBuilder2 with custom parsable for table entries.
func newTableRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams[intCore.RawURLKey] = rawURL
	return newTableRequestBuilder2Internal(urlParams, requestAdapter, factory)
}

// ByID instantiates a new TableItemRequestBuilder2 for the specific record sysID.
func (rB *TableRequestBuilder2) ByID(sysID string) *TableItemRequestBuilder2 {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.BaseRequestBuilder.PathParameters)
	pathParameters["sysid"] = sysID

	return newTableItemRequestBuilder2Internal(pathParameters, rB.BaseRequestBuilder.RequestAdapter, rB.factory)
}

// Get Fetches a response containing Table Entry resources.
func (rB *TableRequestBuilder2) Get(ctx context.Context, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (*newInt.ServiceNowCollectionResponse[serialization.Parsable], error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &TableRequestBuilder2GetRequestConfiguration{}
	}

	opts := nethttplibrary.NewHeadersInspectionOptions()
	opts.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, opts)

	requestInfo, err := rB.ToGetRequestInformation(ctx, nil, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"400": models.NewServiceNowErrorFromDiscriminatorValue,
		"401": models.NewServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, newInt.CreateServiceNowCollectionResponseFromDiscriminatorValue(rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(*newInt.ServiceNowCollectionResponse[serialization.Parsable])
	if !ok {
		return nil, errors.New("res is not ServiceNowResponse")
	}

	if err := newInt.ParseNavLinkHeaders(opts.ResponseHeaders.Get("Link"), snRes); err != nil {
		return nil, err
	}

	return snRes, nil
}

// Post Creates a new Table Record resource.
func (rB *TableRequestBuilder2) Post(ctx context.Context, body TableRecord, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (TableRecord, error) { //nolint:dupl
	if internal.IsNil(rB) {
		return nil, nil
	}

	// TODO: make changes it body based on sysparm_input_display_value

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, newInt.CreateServiceNowItemResponseFromDiscriminatorValue[serialization.Parsable](rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(*newInt.ServiceNowItemResponse[serialization.Parsable])
	if !ok {
		return nil, errors.New("res is not ServiceNowResponse")
	}

	result, err := snRes.GetResult()
	if err != nil {
		return nil, err
	}

	record, ok := result.(TableRecord)
	if !ok {
		return nil, errors.New("result is not TableRecord")
	}

	return record, nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *TableRequestBuilder2) ToGetRequestInformation(_ context.Context, _ TableRecord, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	kiotaRequestInfo.Headers.TryAdd(headerAccept, contentTypeApplicationJSON)

	return &kiotaRequestInfo.RequestInformation, nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *TableRequestBuilder2) ToPostRequestInformation(ctx context.Context, body TableRecord, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:dupl
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	kiotaRequestInfo.Headers.TryAdd(headerAccept, contentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.BaseRequestBuilder.RequestAdapter, contentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}

	return &kiotaRequestInfo.RequestInformation, nil
}
