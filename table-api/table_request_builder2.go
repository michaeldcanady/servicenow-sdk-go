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
	tableURLTemplate = "{+baseurl}/api/now/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder2 ...
type TableRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
	factory serialization.ParsableFactory
}

// NewDefaultTableRequestBuilder2Internal ...
func NewDefaultTableRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder2 {
	return newRequestBuilder2Internal(pathParameters, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// newRequestBuilder2Internal instantiates a new TableRequestBuilderKiota and sets the default values.
func newRequestBuilder2Internal(
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

// NewDefaultTableRequestBuilder2 ...
func NewDefaultTableRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder2 {
	return newRequestBuilderBuilder2(rawURL, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// newRequestBuilderBuilder2 instantiates a new TableRequestBuilderKiota and sets the default values.
func newRequestBuilderBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL
	return newRequestBuilder2Internal(urlParams, requestAdapter, factory)
}

func (rB *TableRequestBuilder2) ByID(sysID string) *TableItemRequestBuilder2 {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.BaseRequestBuilder.PathParameters)
	pathParameters["sysid"] = sysID

	return NewTableItemRequestBuilder2Internal(pathParameters, rB.BaseRequestBuilder.RequestAdapter, rB.factory)
}

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

	requestInfo, err := rB.toGetRequestInformation(ctx, requestConfiguration)
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

func (rB *TableRequestBuilder2) Post(ctx context.Context, body TableRecord, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (TableRecord, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.toPostRequestInformation(ctx, body, requestConfiguration)
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

func (rB *TableRequestBuilder2) toGetRequestInformation(ctx context.Context, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			kiotaRequestInfo.AddQueryParameters(*params)
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
	kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	return &kiotaRequestInfo.RequestInformation, nil
}

func (rB *TableRequestBuilder2) toPostRequestInformation(ctx context.Context, body TableRecord, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			kiotaRequestInfo.AddQueryParameters(*params)
		}
		kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
	kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
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
