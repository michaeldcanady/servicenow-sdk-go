package tableapi

import (
	"context"
	"errors"
	"maps"
	"regexp"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	tableURLTemplate = "{+baseurl}/api/now/v2/table{/sysid}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder2 provides operations to manage Service-Now tables.
type TableRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
	factory serialization.ParsableFactory
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
	urlParams["request-raw-url"] = rawURL
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
func (rB *TableRequestBuilder2) Get(ctx context.Context, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (ServiceNowCollectionResponse, error) {
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

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceNowCollectionResponseFromDiscriminatorValue(rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(ServiceNowCollectionResponse)
	if !ok {
		return nil, errors.New("res is not ServiceNowResponse")
	}

	if err := parseNavLinkHeaders(opts.ResponseHeaders.Get("Link"), snRes); err != nil {
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

	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceNowResponseFromDiscriminatorValue(rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(ServiceNowResponse)
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
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			kiotaRequestInfo.AddQueryParameters(*params)
		}
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	return &kiotaRequestInfo.RequestInformation, nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *TableRequestBuilder2) ToPostRequestInformation(ctx context.Context, body TableRecord, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:dupl
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			kiotaRequestInfo.AddQueryParameters(*params)
		}
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}

	return &kiotaRequestInfo.RequestInformation, nil
}

const (
	firstLinkHeaderKey = "first"
	prevLinkHeaderKey  = "prev"
	nextLinkHeaderKey  = "next"
	lastLinkHeaderKey  = "last"
)

var (
	linkHeaderRegex = regexp.MustCompile(`<([^>]+)>;rel="([^"]+)"`)
)

// parseNavLinkHeaders parses navigational links and applies the to the provided response.
func parseNavLinkHeaders(hearderLinks []string, resp ServiceNowCollectionResponse) error {
	for _, header := range hearderLinks {
		linkMatches := linkHeaderRegex.FindAllStringSubmatch(header, -1)

		for _, match := range linkMatches {
			link := match[1]
			rel := match[2]

			var err error
			// Determine the type of link based on the 'rel' attribute
			switch rel {
			case firstLinkHeaderKey:
				err = resp.setFirstLink(&link)
			case prevLinkHeaderKey:
				err = resp.setPreviousLink(&link)
			case nextLinkHeaderKey:
				err = resp.setNextLink(&link)
			case lastLinkHeaderKey:
				err = resp.setLastLink(&link)
			}
			if err != nil {
				return err
			}
		}
	}

	return nil
}
